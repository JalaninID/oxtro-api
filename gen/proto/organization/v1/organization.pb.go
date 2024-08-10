// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/organization/v1/organization.proto

package orgranizationv1

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

type RequestOrganization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *RequestOrganization) Reset() {
	*x = RequestOrganization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_organization_v1_organization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestOrganization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestOrganization) ProtoMessage() {}

func (x *RequestOrganization) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_v1_organization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestOrganization.ProtoReflect.Descriptor instead.
func (*RequestOrganization) Descriptor() ([]byte, []int) {
	return file_proto_organization_v1_organization_proto_rawDescGZIP(), []int{0}
}

func (x *RequestOrganization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestOrganization) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type ResponseOrganization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *ResponseOrganization) Reset() {
	*x = ResponseOrganization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_organization_v1_organization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseOrganization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseOrganization) ProtoMessage() {}

func (x *ResponseOrganization) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_v1_organization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseOrganization.ProtoReflect.Descriptor instead.
func (*ResponseOrganization) Descriptor() ([]byte, []int) {
	return file_proto_organization_v1_organization_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseOrganization) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseOrganization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseOrganization) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

var File_proto_organization_v1_organization_proto protoreflect.FileDescriptor

var file_proto_organization_v1_organization_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x22, 0x41, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x52, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x32, 0x61, 0x0a, 0x0c, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1d,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2f, 0x5a,
	0x2d, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6f,
	0x72, 0x67, 0x72, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_organization_v1_organization_proto_rawDescOnce sync.Once
	file_proto_organization_v1_organization_proto_rawDescData = file_proto_organization_v1_organization_proto_rawDesc
)

func file_proto_organization_v1_organization_proto_rawDescGZIP() []byte {
	file_proto_organization_v1_organization_proto_rawDescOnce.Do(func() {
		file_proto_organization_v1_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_organization_v1_organization_proto_rawDescData)
	})
	return file_proto_organization_v1_organization_proto_rawDescData
}

var file_proto_organization_v1_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_organization_v1_organization_proto_goTypes = []interface{}{
	(*RequestOrganization)(nil),  // 0: auth.v1.RequestOrganization
	(*ResponseOrganization)(nil), // 1: auth.v1.ResponseOrganization
}
var file_proto_organization_v1_organization_proto_depIdxs = []int32{
	0, // 0: auth.v1.Organization.CreateOrganization:input_type -> auth.v1.RequestOrganization
	1, // 1: auth.v1.Organization.CreateOrganization:output_type -> auth.v1.ResponseOrganization
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_organization_v1_organization_proto_init() }
func file_proto_organization_v1_organization_proto_init() {
	if File_proto_organization_v1_organization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_organization_v1_organization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestOrganization); i {
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
		file_proto_organization_v1_organization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseOrganization); i {
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
			RawDescriptor: file_proto_organization_v1_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_organization_v1_organization_proto_goTypes,
		DependencyIndexes: file_proto_organization_v1_organization_proto_depIdxs,
		MessageInfos:      file_proto_organization_v1_organization_proto_msgTypes,
	}.Build()
	File_proto_organization_v1_organization_proto = out.File
	file_proto_organization_v1_organization_proto_rawDesc = nil
	file_proto_organization_v1_organization_proto_goTypes = nil
	file_proto_organization_v1_organization_proto_depIdxs = nil
}
