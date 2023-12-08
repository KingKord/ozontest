// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.25.0
// source: short.proto

package short

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

type URLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *URLRequest) Reset() {
	*x = URLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_short_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLRequest) ProtoMessage() {}

func (x *URLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_short_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLRequest.ProtoReflect.Descriptor instead.
func (*URLRequest) Descriptor() ([]byte, []int) {
	return file_short_proto_rawDescGZIP(), []int{0}
}

func (x *URLRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type URLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *URLResponse) Reset() {
	*x = URLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_short_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLResponse) ProtoMessage() {}

func (x *URLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_short_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLResponse.ProtoReflect.Descriptor instead.
func (*URLResponse) Descriptor() ([]byte, []int) {
	return file_short_proto_rawDescGZIP(), []int{1}
}

func (x *URLResponse) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

var File_short_proto protoreflect.FileDescriptor

var file_short_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x22, 0x1e, 0x0a, 0x0a, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x52, 0x4c, 0x22, 0x1f, 0x0a, 0x0b, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x55, 0x52, 0x4c, 0x32, 0x7c, 0x0a, 0x0c, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x55, 0x52, 0x4c, 0x12, 0x11, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2e, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2e, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x11, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_short_proto_rawDescOnce sync.Once
	file_short_proto_rawDescData = file_short_proto_rawDesc
)

func file_short_proto_rawDescGZIP() []byte {
	file_short_proto_rawDescOnce.Do(func() {
		file_short_proto_rawDescData = protoimpl.X.CompressGZIP(file_short_proto_rawDescData)
	})
	return file_short_proto_rawDescData
}

var file_short_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_short_proto_goTypes = []interface{}{
	(*URLRequest)(nil),  // 0: short.URLRequest
	(*URLResponse)(nil), // 1: short.URLResponse
}
var file_short_proto_depIdxs = []int32{
	0, // 0: short.UrlShortener.ShortenURL:input_type -> short.URLRequest
	0, // 1: short.UrlShortener.GetOriginalURL:input_type -> short.URLRequest
	1, // 2: short.UrlShortener.ShortenURL:output_type -> short.URLResponse
	1, // 3: short.UrlShortener.GetOriginalURL:output_type -> short.URLResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_short_proto_init() }
func file_short_proto_init() {
	if File_short_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_short_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLRequest); i {
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
		file_short_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLResponse); i {
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
			RawDescriptor: file_short_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_short_proto_goTypes,
		DependencyIndexes: file_short_proto_depIdxs,
		MessageInfos:      file_short_proto_msgTypes,
	}.Build()
	File_short_proto = out.File
	file_short_proto_rawDesc = nil
	file_short_proto_goTypes = nil
	file_short_proto_depIdxs = nil
}
