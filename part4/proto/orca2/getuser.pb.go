// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: orca2/getuser.proto

package proto3

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

type ValidUserName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Statuscode string `protobuf:"bytes,2,opt,name=statuscode,proto3" json:"statuscode,omitempty"`
}

func (x *ValidUserName) Reset() {
	*x = ValidUserName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orca2_getuser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidUserName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidUserName) ProtoMessage() {}

func (x *ValidUserName) ProtoReflect() protoreflect.Message {
	mi := &file_orca2_getuser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidUserName.ProtoReflect.Descriptor instead.
func (*ValidUserName) Descriptor() ([]byte, []int) {
	return file_orca2_getuser_proto_rawDescGZIP(), []int{0}
}

func (x *ValidUserName) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ValidUserName) GetStatuscode() string {
	if x != nil {
		return x.Statuscode
	}
	return ""
}

type StatusTwo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Class      string `protobuf:"bytes,3,opt,name=class,proto3" json:"class,omitempty"`
	Roll       int64  `protobuf:"varint,4,opt,name=roll,proto3" json:"roll,omitempty"`
}

func (x *StatusTwo) Reset() {
	*x = StatusTwo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orca2_getuser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusTwo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusTwo) ProtoMessage() {}

func (x *StatusTwo) ProtoReflect() protoreflect.Message {
	mi := &file_orca2_getuser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusTwo.ProtoReflect.Descriptor instead.
func (*StatusTwo) Descriptor() ([]byte, []int) {
	return file_orca2_getuser_proto_rawDescGZIP(), []int{1}
}

func (x *StatusTwo) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *StatusTwo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StatusTwo) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *StatusTwo) GetRoll() int64 {
	if x != nil {
		return x.Roll
	}
	return 0
}

var File_orca2_getuser_proto protoreflect.FileDescriptor

var file_orca2_getuser_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x72, 0x63, 0x61, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0d,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x6b, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x74, 0x77,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x6c,
	0x32, 0x43, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x12, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x74, 0x77, 0x6f, 0x42, 0x12, 0x5a, 0x10, 0x76, 0x61, 0x6c, 0x6e, 0x6f, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_orca2_getuser_proto_rawDescOnce sync.Once
	file_orca2_getuser_proto_rawDescData = file_orca2_getuser_proto_rawDesc
)

func file_orca2_getuser_proto_rawDescGZIP() []byte {
	file_orca2_getuser_proto_rawDescOnce.Do(func() {
		file_orca2_getuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_orca2_getuser_proto_rawDescData)
	})
	return file_orca2_getuser_proto_rawDescData
}

var file_orca2_getuser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_orca2_getuser_proto_goTypes = []interface{}{
	(*ValidUserName)(nil), // 0: proto.ValidUserName
	(*StatusTwo)(nil),     // 1: proto.Status_two
}
var file_orca2_getuser_proto_depIdxs = []int32{
	0, // 0: proto.StatusUpdater.GetUser:input_type -> proto.ValidUserName
	1, // 1: proto.StatusUpdater.GetUser:output_type -> proto.Status_two
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orca2_getuser_proto_init() }
func file_orca2_getuser_proto_init() {
	if File_orca2_getuser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orca2_getuser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidUserName); i {
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
		file_orca2_getuser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusTwo); i {
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
			RawDescriptor: file_orca2_getuser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orca2_getuser_proto_goTypes,
		DependencyIndexes: file_orca2_getuser_proto_depIdxs,
		MessageInfos:      file_orca2_getuser_proto_msgTypes,
	}.Build()
	File_orca2_getuser_proto = out.File
	file_orca2_getuser_proto_rawDesc = nil
	file_orca2_getuser_proto_goTypes = nil
	file_orca2_getuser_proto_depIdxs = nil
}