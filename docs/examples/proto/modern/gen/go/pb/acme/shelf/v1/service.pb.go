// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: acme/shelf/v1/service.proto

package shelfv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_acme_shelf_v1_service_proto protoreflect.FileDescriptor

var file_acme_shelf_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61,
	0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x61, 0x63,
	0x6d, 0x65, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xa1, 0x04, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76,
	0x65, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0f, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x12, 0x25,
	0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x1e, 0x2e, 0x61, 0x63,
	0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x63,
	0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x21, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x21, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e,
	0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x21, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x91, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x63, 0x6d, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x18, 0x70, 0x62,
	0x2f, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x68, 0x65, 0x6c, 0x66, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x58, 0xaa, 0x02, 0x0d, 0x41,
	0x63, 0x6d, 0x65, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x41,
	0x63, 0x6d, 0x65, 0x5c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x41,
	0x63, 0x6d, 0x65, 0x5c, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x41, 0x63, 0x6d, 0x65, 0x3a,
	0x3a, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_acme_shelf_v1_service_proto_goTypes = []interface{}{
	(*ListShelvesRequest)(nil),      // 0: acme.shelf.v1.ListShelvesRequest
	(*BatchGetShelvesRequest)(nil),  // 1: acme.shelf.v1.BatchGetShelvesRequest
	(*GetShelfRequest)(nil),         // 2: acme.shelf.v1.GetShelfRequest
	(*CreateShelfRequest)(nil),      // 3: acme.shelf.v1.CreateShelfRequest
	(*UpdateShelfRequest)(nil),      // 4: acme.shelf.v1.UpdateShelfRequest
	(*DeleteShelfRequest)(nil),      // 5: acme.shelf.v1.DeleteShelfRequest
	(*ListShelvesResponse)(nil),     // 6: acme.shelf.v1.ListShelvesResponse
	(*BatchGetShelvesResponse)(nil), // 7: acme.shelf.v1.BatchGetShelvesResponse
	(*GetShelfResponse)(nil),        // 8: acme.shelf.v1.GetShelfResponse
	(*CreateShelfResponse)(nil),     // 9: acme.shelf.v1.CreateShelfResponse
	(*UpdateShelfResponse)(nil),     // 10: acme.shelf.v1.UpdateShelfResponse
	(*DeleteShelfResponse)(nil),     // 11: acme.shelf.v1.DeleteShelfResponse
}
var file_acme_shelf_v1_service_proto_depIdxs = []int32{
	0,  // 0: acme.shelf.v1.ShelfService.ListShelves:input_type -> acme.shelf.v1.ListShelvesRequest
	1,  // 1: acme.shelf.v1.ShelfService.BatchGetShelves:input_type -> acme.shelf.v1.BatchGetShelvesRequest
	2,  // 2: acme.shelf.v1.ShelfService.GetShelf:input_type -> acme.shelf.v1.GetShelfRequest
	3,  // 3: acme.shelf.v1.ShelfService.CreateShelf:input_type -> acme.shelf.v1.CreateShelfRequest
	4,  // 4: acme.shelf.v1.ShelfService.UpdateShelf:input_type -> acme.shelf.v1.UpdateShelfRequest
	5,  // 5: acme.shelf.v1.ShelfService.DeleteShelf:input_type -> acme.shelf.v1.DeleteShelfRequest
	6,  // 6: acme.shelf.v1.ShelfService.ListShelves:output_type -> acme.shelf.v1.ListShelvesResponse
	7,  // 7: acme.shelf.v1.ShelfService.BatchGetShelves:output_type -> acme.shelf.v1.BatchGetShelvesResponse
	8,  // 8: acme.shelf.v1.ShelfService.GetShelf:output_type -> acme.shelf.v1.GetShelfResponse
	9,  // 9: acme.shelf.v1.ShelfService.CreateShelf:output_type -> acme.shelf.v1.CreateShelfResponse
	10, // 10: acme.shelf.v1.ShelfService.UpdateShelf:output_type -> acme.shelf.v1.UpdateShelfResponse
	11, // 11: acme.shelf.v1.ShelfService.DeleteShelf:output_type -> acme.shelf.v1.DeleteShelfResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_acme_shelf_v1_service_proto_init() }
func file_acme_shelf_v1_service_proto_init() {
	if File_acme_shelf_v1_service_proto != nil {
		return
	}
	file_acme_shelf_v1_request_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_acme_shelf_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_acme_shelf_v1_service_proto_goTypes,
		DependencyIndexes: file_acme_shelf_v1_service_proto_depIdxs,
	}.Build()
	File_acme_shelf_v1_service_proto = out.File
	file_acme_shelf_v1_service_proto_rawDesc = nil
	file_acme_shelf_v1_service_proto_goTypes = nil
	file_acme_shelf_v1_service_proto_depIdxs = nil
}
