// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: analyze.proto

package AnalysisByNLP

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

type RpcComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content   []string `protobuf:"bytes,1,rep,name=Content,proto3" json:"Content,omitempty"`
	ProductId string   `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *RpcComment) Reset() {
	*x = RpcComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyze_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcComment) ProtoMessage() {}

func (x *RpcComment) ProtoReflect() protoreflect.Message {
	mi := &file_analyze_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcComment.ProtoReflect.Descriptor instead.
func (*RpcComment) Descriptor() ([]byte, []int) {
	return file_analyze_proto_rawDescGZIP(), []int{0}
}

func (x *RpcComment) GetContent() []string {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *RpcComment) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type AnalyzeComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fraction int32   `protobuf:"varint,1,opt,name=Fraction,proto3" json:"Fraction,omitempty"`
	Interval []int32 `protobuf:"varint,2,rep,packed,name=Interval,proto3" json:"Interval,omitempty"`
}

func (x *AnalyzeComment) Reset() {
	*x = AnalyzeComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyze_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzeComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeComment) ProtoMessage() {}

func (x *AnalyzeComment) ProtoReflect() protoreflect.Message {
	mi := &file_analyze_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeComment.ProtoReflect.Descriptor instead.
func (*AnalyzeComment) Descriptor() ([]byte, []int) {
	return file_analyze_proto_rawDescGZIP(), []int{1}
}

func (x *AnalyzeComment) GetFraction() int32 {
	if x != nil {
		return x.Fraction
	}
	return 0
}

func (x *AnalyzeComment) GetInterval() []int32 {
	if x != nil {
		return x.Interval
	}
	return nil
}

var File_analyze_proto protoreflect.FileDescriptor

var file_analyze_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x22, 0x44, 0x0a, 0x0a, 0x72, 0x70, 0x63, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x48,
	0x0a, 0x0e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x32, 0x52, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x15, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4e, 0x4c, 0x50, 0x12, 0x13, 0x2e, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0x17, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x2e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14,
	0x67, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x2f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x42,
	0x79, 0x4e, 0x4c, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analyze_proto_rawDescOnce sync.Once
	file_analyze_proto_rawDescData = file_analyze_proto_rawDesc
)

func file_analyze_proto_rawDescGZIP() []byte {
	file_analyze_proto_rawDescOnce.Do(func() {
		file_analyze_proto_rawDescData = protoimpl.X.CompressGZIP(file_analyze_proto_rawDescData)
	})
	return file_analyze_proto_rawDescData
}

var file_analyze_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_analyze_proto_goTypes = []interface{}{
	(*RpcComment)(nil),     // 0: analyze.rpcComment
	(*AnalyzeComment)(nil), // 1: analyze.AnalyzeComment
}
var file_analyze_proto_depIdxs = []int32{
	0, // 0: analyze.Greeter.AnalysisCommentsByNLP:input_type -> analyze.rpcComment
	1, // 1: analyze.Greeter.AnalysisCommentsByNLP:output_type -> analyze.AnalyzeComment
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_analyze_proto_init() }
func file_analyze_proto_init() {
	if File_analyze_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analyze_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcComment); i {
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
		file_analyze_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzeComment); i {
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
			RawDescriptor: file_analyze_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analyze_proto_goTypes,
		DependencyIndexes: file_analyze_proto_depIdxs,
		MessageInfos:      file_analyze_proto_msgTypes,
	}.Build()
	File_analyze_proto = out.File
	file_analyze_proto_rawDesc = nil
	file_analyze_proto_goTypes = nil
	file_analyze_proto_depIdxs = nil
}
