package migrations

import (
	. "github.com/grafana/grafana/pkg/services/sqlstore/migrator"
)

func addFeatureTableMigration(mg *Migrator) {
	featureTable := Table{
		Name: "features",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "name", Type: DB_NVarchar, Length: 255, Nullable: false},
			{Name: "enabled", Type: DB_Bool, Nullable: false},
		},
		Indices: []*Index{
			{Cols: []string{"name"}},
		},
	}

	mg.AddMigration("create feature table", NewAddTableMigration(featureTable))
	mg.AddMigration("add index feature.name", NewAddIndexMigration(featureTable, featureTable.Indices[0]))
}
