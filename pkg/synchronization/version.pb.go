// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: synchronization/version.proto

package synchronization

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Version specifies a session version, providing default behavior that can vary
// without affecting existing sessions.
type Version int32

const (
	// Invalid is the default session version and represents an unspecfied and
	// invalid version. It is used as a sanity check to ensure that version is
	// set for a session.
	Version_Invalid Version = 0
	// Version1 represents session version 1.
	Version_Version1 Version = 1
)

// Enum value maps for Version.
var (
	Version_name = map[int32]string{
		0: "Invalid",
		1: "Version1",
	}
	Version_value = map[string]int32{
		"Invalid":  0,
		"Version1": 1,
	}
)

func (x Version) Enum() *Version {
	p := new(Version)
	*p = x
	return p
}

func (x Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version) Descriptor() protoreflect.EnumDescriptor {
	return file_synchronization_version_proto_enumTypes[0].Descriptor()
}

func (Version) Type() protoreflect.EnumType {
	return &file_synchronization_version_proto_enumTypes[0]
}

func (x Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version.Descriptor instead.
func (Version) EnumDescriptor() ([]byte, []int) {
	return file_synchronization_version_proto_rawDescGZIP(), []int{0}
}

var File_synchronization_version_proto protoreflect.FileDescriptor

var file_synchronization_version_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2a, 0x24, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x31, 0x10, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f,
	0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63,
	0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_synchronization_version_proto_rawDescOnce sync.Once
	file_synchronization_version_proto_rawDescData = file_synchronization_version_proto_rawDesc
)

func file_synchronization_version_proto_rawDescGZIP() []byte {
	file_synchronization_version_proto_rawDescOnce.Do(func() {
		file_synchronization_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_version_proto_rawDescData)
	})
	return file_synchronization_version_proto_rawDescData
}

var file_synchronization_version_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synchronization_version_proto_goTypes = []interface{}{
	(Version)(0), // 0: synchronization.Version
}
var file_synchronization_version_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_synchronization_version_proto_init() }
func file_synchronization_version_proto_init() {
	if File_synchronization_version_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synchronization_version_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_version_proto_goTypes,
		DependencyIndexes: file_synchronization_version_proto_depIdxs,
		EnumInfos:         file_synchronization_version_proto_enumTypes,
	}.Build()
	File_synchronization_version_proto = out.File
	file_synchronization_version_proto_rawDesc = nil
	file_synchronization_version_proto_goTypes = nil
	file_synchronization_version_proto_depIdxs = nil
}
