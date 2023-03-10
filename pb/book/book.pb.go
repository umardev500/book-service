// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: pb/book/book.proto

package book

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId      string        `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Title       string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author      string        `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Publisher   string        `protobuf:"bytes,4,opt,name=publisher,proto3" json:"publisher,omitempty"`
	PagesTotal  int64         `protobuf:"varint,5,opt,name=pages_total,json=pagesTotal,proto3" json:"pages_total,omitempty"`
	Qty         int64         `protobuf:"varint,6,opt,name=qty,proto3" json:"qty,omitempty"`
	Cover       string        `protobuf:"bytes,7,opt,name=cover,proto3" json:"cover,omitempty"`
	Description string        `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	LocationId  string        `protobuf:"bytes,9,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Uploader    *BookUploader `protobuf:"bytes,10,opt,name=uploader,proto3" json:"uploader,omitempty"`
	Editors     []*BookEditor `protobuf:"bytes,11,rep,name=editors,proto3" json:"editors,omitempty"`
	CreatedAt   int64         `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   int64         `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   int64         `protobuf:"varint,14,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_book_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_pb_book_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_pb_book_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *Book) GetPagesTotal() int64 {
	if x != nil {
		return x.PagesTotal
	}
	return 0
}

func (x *Book) GetQty() int64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *Book) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *Book) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Book) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *Book) GetUploader() *BookUploader {
	if x != nil {
		return x.Uploader
	}
	return nil
}

func (x *Book) GetEditors() []*BookEditor {
	if x != nil {
		return x.Editors
	}
	return nil
}

func (x *Book) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Book) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Book) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_pb_book_book_proto protoreflect.FileDescriptor

var file_pb_book_book_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x62, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x2d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x70, 0x62, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2d,
	0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x03, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x71,
	0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52,
	0x07, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x62, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_book_book_proto_rawDescOnce sync.Once
	file_pb_book_book_proto_rawDescData = file_pb_book_book_proto_rawDesc
)

func file_pb_book_book_proto_rawDescGZIP() []byte {
	file_pb_book_book_proto_rawDescOnce.Do(func() {
		file_pb_book_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_book_book_proto_rawDescData)
	})
	return file_pb_book_book_proto_rawDescData
}

var file_pb_book_book_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_book_book_proto_goTypes = []interface{}{
	(*Book)(nil),         // 0: Book
	(*BookUploader)(nil), // 1: BookUploader
	(*BookEditor)(nil),   // 2: BookEditor
}
var file_pb_book_book_proto_depIdxs = []int32{
	1, // 0: Book.uploader:type_name -> BookUploader
	2, // 1: Book.editors:type_name -> BookEditor
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_book_book_proto_init() }
func file_pb_book_book_proto_init() {
	if File_pb_book_book_proto != nil {
		return
	}
	file_pb_book_book_uploader_proto_init()
	file_pb_book_book_editor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_book_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
			RawDescriptor: file_pb_book_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_book_book_proto_goTypes,
		DependencyIndexes: file_pb_book_book_proto_depIdxs,
		MessageInfos:      file_pb_book_book_proto_msgTypes,
	}.Build()
	File_pb_book_book_proto = out.File
	file_pb_book_book_proto_rawDesc = nil
	file_pb_book_book_proto_goTypes = nil
	file_pb_book_book_proto_depIdxs = nil
}
