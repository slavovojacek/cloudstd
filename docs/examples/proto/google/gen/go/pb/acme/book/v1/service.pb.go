// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: acme/book/v1/service.proto

package bookv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_acme_book_v1_service_proto protoreflect.FileDescriptor

var file_acme_book_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x63,
	0x6d, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x63,
	0x6d, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xef, 0x02, 0x0a, 0x0b, 0x42, 0x6f,
	0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1f, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1f, 0x2e, 0x61, 0x63,
	0x6d, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61,
	0x63, 0x6d, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x1f, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x8a, 0x01, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x16, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76,
	0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x42, 0x58, 0xaa, 0x02,
	0x0c, 0x41, 0x63, 0x6d, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c,
	0x41, 0x63, 0x6d, 0x65, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x41,
	0x63, 0x6d, 0x65, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x41, 0x63, 0x6d, 0x65, 0x3a, 0x3a,
	0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_acme_book_v1_service_proto_goTypes = []interface{}{
	(*ListBooksRequest)(nil),  // 0: acme.book.v1.ListBooksRequest
	(*GetBookRequest)(nil),    // 1: acme.book.v1.GetBookRequest
	(*CreateBookRequest)(nil), // 2: acme.book.v1.CreateBookRequest
	(*UpdateBookRequest)(nil), // 3: acme.book.v1.UpdateBookRequest
	(*DeleteBookRequest)(nil), // 4: acme.book.v1.DeleteBookRequest
	(*ListBooksResponse)(nil), // 5: acme.book.v1.ListBooksResponse
	(*Book)(nil),              // 6: acme.book.v1.Book
	(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_acme_book_v1_service_proto_depIdxs = []int32{
	0, // 0: acme.book.v1.BookService.ListBooks:input_type -> acme.book.v1.ListBooksRequest
	1, // 1: acme.book.v1.BookService.GetBook:input_type -> acme.book.v1.GetBookRequest
	2, // 2: acme.book.v1.BookService.CreateBook:input_type -> acme.book.v1.CreateBookRequest
	3, // 3: acme.book.v1.BookService.UpdateBook:input_type -> acme.book.v1.UpdateBookRequest
	4, // 4: acme.book.v1.BookService.DeleteBook:input_type -> acme.book.v1.DeleteBookRequest
	5, // 5: acme.book.v1.BookService.ListBooks:output_type -> acme.book.v1.ListBooksResponse
	6, // 6: acme.book.v1.BookService.GetBook:output_type -> acme.book.v1.Book
	6, // 7: acme.book.v1.BookService.CreateBook:output_type -> acme.book.v1.Book
	6, // 8: acme.book.v1.BookService.UpdateBook:output_type -> acme.book.v1.Book
	7, // 9: acme.book.v1.BookService.DeleteBook:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_acme_book_v1_service_proto_init() }
func file_acme_book_v1_service_proto_init() {
	if File_acme_book_v1_service_proto != nil {
		return
	}
	file_acme_book_v1_request_response_proto_init()
	file_acme_book_v1_resource_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_acme_book_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_acme_book_v1_service_proto_goTypes,
		DependencyIndexes: file_acme_book_v1_service_proto_depIdxs,
	}.Build()
	File_acme_book_v1_service_proto = out.File
	file_acme_book_v1_service_proto_rawDesc = nil
	file_acme_book_v1_service_proto_goTypes = nil
	file_acme_book_v1_service_proto_depIdxs = nil
}
