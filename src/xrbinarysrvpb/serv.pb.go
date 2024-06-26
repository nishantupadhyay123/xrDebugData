// ----------------------------------------------------------------------------
// xr_debug_data.proto - XR debug data protobuf definitions
// Arpil 2024,Copyright (c) 2016 by Cisco Systems, Inc.
// ----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.5.0
// source: serv.proto

package xrbinarysrvpb

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

type XrDebugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reqid   int64  `protobuf:"varint,1,opt,name=reqid,proto3" json:"reqid,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Errors  string `protobuf:"bytes,3,opt,name=errors,proto3" json:"errors,omitempty"`
	Decoder string `protobuf:"bytes,4,opt,name=decoder,proto3" json:"decoder,omitempty"`
}

func (x *XrDebugRequest) Reset() {
	*x = XrDebugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XrDebugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XrDebugRequest) ProtoMessage() {}

func (x *XrDebugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XrDebugRequest.ProtoReflect.Descriptor instead.
func (*XrDebugRequest) Descriptor() ([]byte, []int) {
	return file_serv_proto_rawDescGZIP(), []int{0}
}

func (x *XrDebugRequest) GetReqid() int64 {
	if x != nil {
		return x.Reqid
	}
	return 0
}

func (x *XrDebugRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *XrDebugRequest) GetErrors() string {
	if x != nil {
		return x.Errors
	}
	return ""
}

func (x *XrDebugRequest) GetDecoder() string {
	if x != nil {
		return x.Decoder
	}
	return ""
}

type XrDebugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *XrDebugResponse) Reset() {
	*x = XrDebugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XrDebugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XrDebugResponse) ProtoMessage() {}

func (x *XrDebugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_serv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XrDebugResponse.ProtoReflect.Descriptor instead.
func (*XrDebugResponse) Descriptor() ([]byte, []int) {
	return file_serv_proto_rawDescGZIP(), []int{1}
}

func (x *XrDebugResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *XrDebugResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_serv_proto protoreflect.FileDescriptor

var file_serv_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x78, 0x72,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x73, 0x72, 0x76, 0x22, 0x6c, 0x0a, 0x0e, 0x58, 0x72, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x71, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x65, 0x71, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x0f, 0x58, 0x72, 0x44, 0x65, 0x62,
	0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x5f, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x78, 0x72, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x73, 0x72, 0x76, 0x2e, 0x58, 0x72, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x78, 0x72, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x73, 0x72, 0x76, 0x2e, 0x58, 0x72, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2e, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x78, 0x72, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x73, 0x72, 0x76, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serv_proto_rawDescOnce sync.Once
	file_serv_proto_rawDescData = file_serv_proto_rawDesc
)

func file_serv_proto_rawDescGZIP() []byte {
	file_serv_proto_rawDescOnce.Do(func() {
		file_serv_proto_rawDescData = protoimpl.X.CompressGZIP(file_serv_proto_rawDescData)
	})
	return file_serv_proto_rawDescData
}

var file_serv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_serv_proto_goTypes = []interface{}{
	(*XrDebugRequest)(nil),  // 0: xrbinarysrv.XrDebugRequest
	(*XrDebugResponse)(nil), // 1: xrbinarysrv.XrDebugResponse
}
var file_serv_proto_depIdxs = []int32{
	0, // 0: xrbinarysrv.UploadService.UploadRequest:input_type -> xrbinarysrv.XrDebugRequest
	1, // 1: xrbinarysrv.UploadService.UploadRequest:output_type -> xrbinarysrv.XrDebugResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_serv_proto_init() }
func file_serv_proto_init() {
	if File_serv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XrDebugRequest); i {
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
		file_serv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XrDebugResponse); i {
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
			RawDescriptor: file_serv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serv_proto_goTypes,
		DependencyIndexes: file_serv_proto_depIdxs,
		MessageInfos:      file_serv_proto_msgTypes,
	}.Build()
	File_serv_proto = out.File
	file_serv_proto_rawDesc = nil
	file_serv_proto_goTypes = nil
	file_serv_proto_depIdxs = nil
}
