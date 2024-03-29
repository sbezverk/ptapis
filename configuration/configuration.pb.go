// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: configuration.proto

package configuration

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

type Function int32

const (
	Function_PTAGENT     Function = 0
	Function_PTCOLLECTOR Function = 1
	Function_PTRECORDER  Function = 2
)

// Enum value maps for Function.
var (
	Function_name = map[int32]string{
		0: "PTAGENT",
		1: "PTCOLLECTOR",
		2: "PTRECORDER",
	}
	Function_value = map[string]int32{
		"PTAGENT":     0,
		"PTCOLLECTOR": 1,
		"PTRECORDER":  2,
	}
)

func (x Function) Enum() *Function {
	p := new(Function)
	*p = x
	return p
}

func (x Function) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Function) Descriptor() protoreflect.EnumDescriptor {
	return file_configuration_proto_enumTypes[0].Descriptor()
}

func (Function) Type() protoreflect.EnumType {
	return &file_configuration_proto_enumTypes[0]
}

func (x Function) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Function.Descriptor instead.
func (Function) EnumDescriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{0}
}

type ReturnCode int32

const (
	ReturnCode_OK      ReturnCode = 0
	ReturnCode_ERRID   ReturnCode = 1
	ReturnCode_ERRFUNC ReturnCode = 2
	ReturnCode_ERRNAME ReturnCode = 3
)

// Enum value maps for ReturnCode.
var (
	ReturnCode_name = map[int32]string{
		0: "OK",
		1: "ERRID",
		2: "ERRFUNC",
		3: "ERRNAME",
	}
	ReturnCode_value = map[string]int32{
		"OK":      0,
		"ERRID":   1,
		"ERRFUNC": 2,
		"ERRNAME": 3,
	}
)

func (x ReturnCode) Enum() *ReturnCode {
	p := new(ReturnCode)
	*p = x
	return p
}

func (x ReturnCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReturnCode) Descriptor() protoreflect.EnumDescriptor {
	return file_configuration_proto_enumTypes[1].Descriptor()
}

func (ReturnCode) Type() protoreflect.EnumType {
	return &file_configuration_proto_enumTypes[1]
}

func (x ReturnCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReturnCode.Descriptor instead.
func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{1}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id       []byte   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Function Function `protobuf:"varint,3,opt,name=function,proto3,enum=configuration.Function" json:"function,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Request) GetFunction() Function {
	if x != nil {
		return x.Function
	}
	return Function_PTAGENT
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []byte            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReturnCode  ReturnCode        `protobuf:"varint,2,opt,name=return_code,json=returnCode,proto3,enum=configuration.ReturnCode" json:"return_code,omitempty"`
	ErrorDetail string            `protobuf:"bytes,3,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	Port        map[string]uint32 `protobuf:"bytes,5,rep,name=port,proto3" json:"port,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Reply) GetReturnCode() ReturnCode {
	if x != nil {
		return x.ReturnCode
	}
	return ReturnCode_OK
}

func (x *Reply) GetErrorDetail() string {
	if x != nil {
		return x.ErrorDetail
	}
	return ""
}

func (x *Reply) GetPort() map[string]uint32 {
	if x != nil {
		return x.Port
	}
	return nil
}

var File_configuration_proto protoreflect.FileDescriptor

var file_configuration_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe3, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x32, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x37, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x38,
	0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x54,
	0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x54, 0x43, 0x4f, 0x4c,
	0x4c, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x54, 0x52, 0x45,
	0x43, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x02, 0x2a, 0x39, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x52, 0x52, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x52, 0x52,
	0x46, 0x55, 0x4e, 0x43, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x52, 0x52, 0x4e, 0x41, 0x4d,
	0x45, 0x10, 0x03, 0x32, 0x56, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x46, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x2e,
	0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configuration_proto_rawDescOnce sync.Once
	file_configuration_proto_rawDescData = file_configuration_proto_rawDesc
)

func file_configuration_proto_rawDescGZIP() []byte {
	file_configuration_proto_rawDescOnce.Do(func() {
		file_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_configuration_proto_rawDescData)
	})
	return file_configuration_proto_rawDescData
}

var file_configuration_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_configuration_proto_goTypes = []interface{}{
	(Function)(0),   // 0: configuration.Function
	(ReturnCode)(0), // 1: configuration.ReturnCode
	(*Request)(nil), // 2: configuration.Request
	(*Reply)(nil),   // 3: configuration.Reply
	nil,             // 4: configuration.Reply.PortEntry
}
var file_configuration_proto_depIdxs = []int32{
	0, // 0: configuration.Request.function:type_name -> configuration.Function
	1, // 1: configuration.Reply.return_code:type_name -> configuration.ReturnCode
	4, // 2: configuration.Reply.port:type_name -> configuration.Reply.PortEntry
	2, // 3: configuration.Configurator.ConfigurationRequest:input_type -> configuration.Request
	3, // 4: configuration.Configurator.ConfigurationRequest:output_type -> configuration.Reply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_configuration_proto_init() }
func file_configuration_proto_init() {
	if File_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_configuration_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_configuration_proto_goTypes,
		DependencyIndexes: file_configuration_proto_depIdxs,
		EnumInfos:         file_configuration_proto_enumTypes,
		MessageInfos:      file_configuration_proto_msgTypes,
	}.Build()
	File_configuration_proto = out.File
	file_configuration_proto_rawDesc = nil
	file_configuration_proto_goTypes = nil
	file_configuration_proto_depIdxs = nil
}
