// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tools/v1/pagination.proto

package toolsv1

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

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage   int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Total     int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	TotalPage int32 `protobuf:"varint,4,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_v1_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_tools_v1_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_tools_v1_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *Pagination) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Pagination) GetTotalPage() int32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

var File_tools_v1_pagination_proto protoreflect.FileDescriptor

var file_tools_v1_pagination_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x70, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x42, 0x7a, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x18, 0x61, 0x70, 0x70, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x14, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tools_v1_pagination_proto_rawDescOnce sync.Once
	file_tools_v1_pagination_proto_rawDescData = file_tools_v1_pagination_proto_rawDesc
)

func file_tools_v1_pagination_proto_rawDescGZIP() []byte {
	file_tools_v1_pagination_proto_rawDescOnce.Do(func() {
		file_tools_v1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_tools_v1_pagination_proto_rawDescData)
	})
	return file_tools_v1_pagination_proto_rawDescData
}

var file_tools_v1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tools_v1_pagination_proto_goTypes = []interface{}{
	(*Pagination)(nil), // 0: tools.v1.Pagination
}
var file_tools_v1_pagination_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tools_v1_pagination_proto_init() }
func file_tools_v1_pagination_proto_init() {
	if File_tools_v1_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tools_v1_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
			RawDescriptor: file_tools_v1_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tools_v1_pagination_proto_goTypes,
		DependencyIndexes: file_tools_v1_pagination_proto_depIdxs,
		MessageInfos:      file_tools_v1_pagination_proto_msgTypes,
	}.Build()
	File_tools_v1_pagination_proto = out.File
	file_tools_v1_pagination_proto_rawDesc = nil
	file_tools_v1_pagination_proto_goTypes = nil
	file_tools_v1_pagination_proto_depIdxs = nil
}
