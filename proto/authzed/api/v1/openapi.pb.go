// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: authzed/api/v1/openapi.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_authzed_api_v1_openapi_proto protoreflect.FileDescriptor

const file_authzed_api_v1_openapi_proto_rawDesc = "" +
	"\n" +
	"\x1cauthzed/api/v1/openapi.proto\x12\x0eauthzed.api.v1\x1a.protoc-gen-openapiv2/options/annotations.protoB\xd7\x03\x92A\x89\x03\x12\x9c\x01\n" +
	"\aAuthzed\"D\n" +
	"\rAuthzed, Inc.\x12\x1ehttps://github.com/authzed/api\x1a\x13support@authzed.com*F\n" +
	"\x12Apache 2.0 License\x120https://github.com/authzed/api/blob/main/LICENSE2\x031.0*\x03\x01\x02\x042\x10application/json:\x10application/jsonZf\n" +
	"d\n" +
	"\n" +
	"ApiKeyAuth\x12V\b\x02\x12ASpiceDB preshared-key, prefixed by Bearer: Bearer <preshared-key>\x1a\rAuthorization \x02b\x10\n" +
	"\x0e\n" +
	"\n" +
	"ApiKeyAuth\x12\x00rE\n" +
	"\x1bMore about the Authzed API.\x12&https://docs.authzed.com/reference/api\n" +
	"\x12com.authzed.api.v1P\x01Z2github.com/authzed/authzed-go/proto/authzed/api/v1b\x06proto3"

var file_authzed_api_v1_openapi_proto_goTypes = []any{}
var file_authzed_api_v1_openapi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authzed_api_v1_openapi_proto_init() }
func file_authzed_api_v1_openapi_proto_init() {
	if File_authzed_api_v1_openapi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_authzed_api_v1_openapi_proto_rawDesc), len(file_authzed_api_v1_openapi_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authzed_api_v1_openapi_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1_openapi_proto_depIdxs,
	}.Build()
	File_authzed_api_v1_openapi_proto = out.File
	file_authzed_api_v1_openapi_proto_goTypes = nil
	file_authzed_api_v1_openapi_proto_depIdxs = nil
}
