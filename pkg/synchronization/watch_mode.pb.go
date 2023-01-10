// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: synchronization/watch_mode.proto

package synchronization

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

// WatchMode specifies the mode for filesystem watching.
type WatchMode int32

const (
	// WatchMode_WatchModeDefault represents an unspecified watch mode. It
	// should be converted to one of the following values based on the desired
	// default behavior.
	WatchMode_WatchModeDefault WatchMode = 0
	// WatchMode_WatchModePortable specifies that native recursive watching
	// should be used to monitor paths on systems that support it if those paths
	// fall under the home directory. In these cases, a watch on the entire home
	// directory is established and filtered for events pertaining to the
	// specified path. On all other systems and for all other paths, poll-based
	// watching is used.
	WatchMode_WatchModePortable WatchMode = 1
	// WatchMode_WatchModeForcePoll specifies that only poll-based watching
	// should be used.
	WatchMode_WatchModeForcePoll WatchMode = 2
	// WatchMode_WatchModeNoWatch specifies that no watching should be used
	// (i.e. no events should be generated).
	WatchMode_WatchModeNoWatch WatchMode = 3
)

// Enum value maps for WatchMode.
var (
	WatchMode_name = map[int32]string{
		0: "WatchModeDefault",
		1: "WatchModePortable",
		2: "WatchModeForcePoll",
		3: "WatchModeNoWatch",
	}
	WatchMode_value = map[string]int32{
		"WatchModeDefault":   0,
		"WatchModePortable":  1,
		"WatchModeForcePoll": 2,
		"WatchModeNoWatch":   3,
	}
)

func (x WatchMode) Enum() *WatchMode {
	p := new(WatchMode)
	*p = x
	return p
}

func (x WatchMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WatchMode) Descriptor() protoreflect.EnumDescriptor {
	return file_synchronization_watch_mode_proto_enumTypes[0].Descriptor()
}

func (WatchMode) Type() protoreflect.EnumType {
	return &file_synchronization_watch_mode_proto_enumTypes[0]
}

func (x WatchMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WatchMode.Descriptor instead.
func (WatchMode) EnumDescriptor() ([]byte, []int) {
	return file_synchronization_watch_mode_proto_rawDescGZIP(), []int{0}
}

var File_synchronization_watch_mode_proto protoreflect.FileDescriptor

var file_synchronization_watch_mode_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2a, 0x66, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d,
	0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x50,
	0x6f, 0x6c, 0x6c, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f,
	0x64, 0x65, 0x4e, 0x6f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x10, 0x03, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65,
	0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_watch_mode_proto_rawDescOnce sync.Once
	file_synchronization_watch_mode_proto_rawDescData = file_synchronization_watch_mode_proto_rawDesc
)

func file_synchronization_watch_mode_proto_rawDescGZIP() []byte {
	file_synchronization_watch_mode_proto_rawDescOnce.Do(func() {
		file_synchronization_watch_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_watch_mode_proto_rawDescData)
	})
	return file_synchronization_watch_mode_proto_rawDescData
}

var file_synchronization_watch_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synchronization_watch_mode_proto_goTypes = []interface{}{
	(WatchMode)(0), // 0: synchronization.WatchMode
}
var file_synchronization_watch_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_synchronization_watch_mode_proto_init() }
func file_synchronization_watch_mode_proto_init() {
	if File_synchronization_watch_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synchronization_watch_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_watch_mode_proto_goTypes,
		DependencyIndexes: file_synchronization_watch_mode_proto_depIdxs,
		EnumInfos:         file_synchronization_watch_mode_proto_enumTypes,
	}.Build()
	File_synchronization_watch_mode_proto = out.File
	file_synchronization_watch_mode_proto_rawDesc = nil
	file_synchronization_watch_mode_proto_goTypes = nil
	file_synchronization_watch_mode_proto_depIdxs = nil
}
