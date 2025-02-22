// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: synchronization/rsync/receive.proto

package rsync

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

// ReceiverState encodes that status of an rsync receiver. It should be
// considered immutable.
type ReceiverState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path is the path currently being received.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// ReceivedSize is the number of bytes that have been received for the
	// current path from both block and data operations.
	ReceivedSize uint64 `protobuf:"varint,2,opt,name=receivedSize,proto3" json:"receivedSize,omitempty"`
	// ExpectedSize is the number of bytes expected for the current path.
	ExpectedSize uint64 `protobuf:"varint,3,opt,name=expectedSize,proto3" json:"expectedSize,omitempty"`
	// ReceivedFiles is the number of files that have already been received.
	ReceivedFiles uint64 `protobuf:"varint,4,opt,name=receivedFiles,proto3" json:"receivedFiles,omitempty"`
	// ExpectedFiles is the total number of files expected.
	ExpectedFiles uint64 `protobuf:"varint,5,opt,name=expectedFiles,proto3" json:"expectedFiles,omitempty"`
	// TotalReceivedSize is the total number of bytes that have been received
	// for all files from both block and data operations.
	TotalReceivedSize uint64 `protobuf:"varint,6,opt,name=totalReceivedSize,proto3" json:"totalReceivedSize,omitempty"`
}

func (x *ReceiverState) Reset() {
	*x = ReceiverState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_rsync_receive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiverState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiverState) ProtoMessage() {}

func (x *ReceiverState) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_rsync_receive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiverState.ProtoReflect.Descriptor instead.
func (*ReceiverState) Descriptor() ([]byte, []int) {
	return file_synchronization_rsync_receive_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiverState) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ReceiverState) GetReceivedSize() uint64 {
	if x != nil {
		return x.ReceivedSize
	}
	return 0
}

func (x *ReceiverState) GetExpectedSize() uint64 {
	if x != nil {
		return x.ExpectedSize
	}
	return 0
}

func (x *ReceiverState) GetReceivedFiles() uint64 {
	if x != nil {
		return x.ReceivedFiles
	}
	return 0
}

func (x *ReceiverState) GetExpectedFiles() uint64 {
	if x != nil {
		return x.ExpectedFiles
	}
	return 0
}

func (x *ReceiverState) GetTotalReceivedSize() uint64 {
	if x != nil {
		return x.TotalReceivedSize
	}
	return 0
}

var File_synchronization_rsync_receive_proto protoreflect.FileDescriptor

var file_synchronization_rsync_receive_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x22, 0xe5, 0x01, 0x0a,
	0x0d, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x53, 0x69, 0x7a, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75,
	0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_rsync_receive_proto_rawDescOnce sync.Once
	file_synchronization_rsync_receive_proto_rawDescData = file_synchronization_rsync_receive_proto_rawDesc
)

func file_synchronization_rsync_receive_proto_rawDescGZIP() []byte {
	file_synchronization_rsync_receive_proto_rawDescOnce.Do(func() {
		file_synchronization_rsync_receive_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_rsync_receive_proto_rawDescData)
	})
	return file_synchronization_rsync_receive_proto_rawDescData
}

var file_synchronization_rsync_receive_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_synchronization_rsync_receive_proto_goTypes = []interface{}{
	(*ReceiverState)(nil), // 0: rsync.ReceiverState
}
var file_synchronization_rsync_receive_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_synchronization_rsync_receive_proto_init() }
func file_synchronization_rsync_receive_proto_init() {
	if File_synchronization_rsync_receive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_synchronization_rsync_receive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiverState); i {
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
			RawDescriptor: file_synchronization_rsync_receive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_rsync_receive_proto_goTypes,
		DependencyIndexes: file_synchronization_rsync_receive_proto_depIdxs,
		MessageInfos:      file_synchronization_rsync_receive_proto_msgTypes,
	}.Build()
	File_synchronization_rsync_receive_proto = out.File
	file_synchronization_rsync_receive_proto_rawDesc = nil
	file_synchronization_rsync_receive_proto_goTypes = nil
	file_synchronization_rsync_receive_proto_depIdxs = nil
}
