// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: feature-server-proto/feature.proto

package featureserverpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FeatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FeatureRequest) Reset() {
	*x = FeatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_server_proto_feature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureRequest) ProtoMessage() {}

func (x *FeatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feature_server_proto_feature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureRequest.ProtoReflect.Descriptor instead.
func (*FeatureRequest) Descriptor() ([]byte, []int) {
	return file_feature_server_proto_feature_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FeatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Enabled bool   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *FeatureResponse) Reset() {
	*x = FeatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_server_proto_feature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureResponse) ProtoMessage() {}

func (x *FeatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feature_server_proto_feature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureResponse.ProtoReflect.Descriptor instead.
func (*FeatureResponse) Descriptor() ([]byte, []int) {
	return file_feature_server_proto_feature_proto_rawDescGZIP(), []int{1}
}

func (x *FeatureResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeatureResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureResponse) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_feature_server_proto_feature_proto protoreflect.FileDescriptor

var file_feature_server_proto_feature_proto_rawDesc = []byte{
	0x0a, 0x22, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x0e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x0f, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x32, 0x57, 0x0a, 0x07, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x49, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1d, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x6f, 0x73, 0x74, 0x68, 0x75, 0x73, 0x2f, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feature_server_proto_feature_proto_rawDescOnce sync.Once
	file_feature_server_proto_feature_proto_rawDescData = file_feature_server_proto_feature_proto_rawDesc
)

func file_feature_server_proto_feature_proto_rawDescGZIP() []byte {
	file_feature_server_proto_feature_proto_rawDescOnce.Do(func() {
		file_feature_server_proto_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_feature_server_proto_feature_proto_rawDescData)
	})
	return file_feature_server_proto_feature_proto_rawDescData
}

var file_feature_server_proto_feature_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_feature_server_proto_feature_proto_goTypes = []interface{}{
	(*FeatureRequest)(nil),  // 0: featureserver.FeatureRequest
	(*FeatureResponse)(nil), // 1: featureserver.FeatureResponse
}
var file_feature_server_proto_feature_proto_depIdxs = []int32{
	0, // 0: featureserver.Feature.IsEnabled:input_type -> featureserver.FeatureRequest
	1, // 1: featureserver.Feature.IsEnabled:output_type -> featureserver.FeatureResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feature_server_proto_feature_proto_init() }
func file_feature_server_proto_feature_proto_init() {
	if File_feature_server_proto_feature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feature_server_proto_feature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feature_server_proto_feature_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feature_server_proto_feature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feature_server_proto_feature_proto_goTypes,
		DependencyIndexes: file_feature_server_proto_feature_proto_depIdxs,
		MessageInfos:      file_feature_server_proto_feature_proto_msgTypes,
	}.Build()
	File_feature_server_proto_feature_proto = out.File
	file_feature_server_proto_feature_proto_rawDesc = nil
	file_feature_server_proto_feature_proto_goTypes = nil
	file_feature_server_proto_feature_proto_depIdxs = nil
}
