package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/andersosthus/feature-server/pkg/services/sqlstore/migrations"
	"github.com/grafana/grafana/pkg/services/sqlstore/migrator"
	"github.com/mailgun/groupcache"
	"k8s.io/klog/v2"

	"k8s.io/sample-controller/pkg/signals"
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()

	port := flag.Int("port", 9000, "port")
	otherPort := flag.Int("otherPort", 9001, "otherPort")

	flag.Parse()

	pool := groupcache.NewHTTPPoolOpts(fmt.Sprintf("http://localhost:%d", *port), &groupcache.HTTPPoolOptions{})
	pool.Set(fmt.Sprintf("http://localhost:%d", *port), fmt.Sprintf("http://localhost:%d", *otherPort))

	server := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", *port),
		Handler: pool,
	}

	go func() {
		klog.InfoS("serving cache...", "port", *port)
		if err := server.ListenAndServe(); err != nil {
			klog.ErrorS(err, "error serving cache")
		}
	}()
	defer server.Shutdown(context.Background())

	// group := groupcache.NewGroup("users", 3000000, groupcache.GetterFunc(
	// 	func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
	// 		klog.InfoS("loading user from store", "userid", key)
	// 		user, err := fetchUser(context.TODO(), key)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		if err := dest.SetProto(&user, time.Now().Add(time.Minute*5)); err != nil {
	// 			return err
	// 		}

	// 		return nil
	// 	},
	// ))

	ctx, cancel := context.WithCancel(context.Background())

	// if *port != 9001 {
	// 	go func() {
	// 		tCtx, tCancel := context.WithTimeout(ctx, 100*time.Millisecond)
	// 		defer tCancel()
	// 		var user User

	// 		for ok := true; ok; ok = true {
	// 			id := rand.Intn(100)

	// 			klog.InfoS("loading user", "userid", id)
	// 			if err := group.Get(tCtx, strconv.Itoa(id), groupcache.ProtoSink(&user)); err != nil {
	// 				klog.ErrorS(err, "error getting from cache", "key", id)
	// 			}

	// 			klog.InfoS("got user", "user", id, "name", user.Name)

	// 			time.Sleep(time.Duration(id) * time.Millisecond)
	// 		}
	// 	}()
	// } else {
	// 	go func() {
	// 		for ok := true; ok; ok = true {
	// 			stats := group.CacheStats(groupcache.MainCache)
	// 			klog.InfoS("cachedata", "hits", stats.Hits, "gets", stats.Gets, "items", stats.Items)
	// 			time.Sleep(1 * time.Second)
	// 		}
	// 	}()
	// }

	mig := migrator.NewMigrator(nil)
	migrations.AddMigrations(mig)

	stopCh := signals.SetupSignalHandler()

	select {
	case <-stopCh:
		klog.Infof("Received SIGTERM, exiting gracefully...")
	case <-ctx.Done():
	}

	cancel()
}
