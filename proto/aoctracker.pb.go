// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: aoctracker.proto

package proto

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

type TrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentYear int32 `protobuf:"varint,1,opt,name=current_year,json=currentYear,proto3" json:"current_year,omitempty"`
	CurrentDay  int32 `protobuf:"varint,2,opt,name=current_day,json=currentDay,proto3" json:"current_day,omitempty"`
	CurrentPart int32 `protobuf:"varint,3,opt,name=current_part,json=currentPart,proto3" json:"current_part,omitempty"`
}

func (x *TrackRequest) Reset() {
	*x = TrackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aoctracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackRequest) ProtoMessage() {}

func (x *TrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aoctracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackRequest.ProtoReflect.Descriptor instead.
func (*TrackRequest) Descriptor() ([]byte, []int) {
	return file_aoctracker_proto_rawDescGZIP(), []int{0}
}

func (x *TrackRequest) GetCurrentYear() int32 {
	if x != nil {
		return x.CurrentYear
	}
	return 0
}

func (x *TrackRequest) GetCurrentDay() int32 {
	if x != nil {
		return x.CurrentDay
	}
	return 0
}

func (x *TrackRequest) GetCurrentPart() int32 {
	if x != nil {
		return x.CurrentPart
	}
	return 0
}

type TrackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TrackResponse) Reset() {
	*x = TrackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aoctracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackResponse) ProtoMessage() {}

func (x *TrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aoctracker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackResponse.ProtoReflect.Descriptor instead.
func (*TrackResponse) Descriptor() ([]byte, []int) {
	return file_aoctracker_proto_rawDescGZIP(), []int{1}
}

var File_aoctracker_proto protoreflect.FileDescriptor

var file_aoctracker_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x6f, 0x63, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x6f, 0x63, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x22, 0x75,
	0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x72, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x53, 0x0a, 0x11, 0x41, 0x4f, 0x43, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x61, 0x6f, 0x63, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x6f, 0x63, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x61, 0x6f, 0x63, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aoctracker_proto_rawDescOnce sync.Once
	file_aoctracker_proto_rawDescData = file_aoctracker_proto_rawDesc
)

func file_aoctracker_proto_rawDescGZIP() []byte {
	file_aoctracker_proto_rawDescOnce.Do(func() {
		file_aoctracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_aoctracker_proto_rawDescData)
	})
	return file_aoctracker_proto_rawDescData
}

var file_aoctracker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aoctracker_proto_goTypes = []interface{}{
	(*TrackRequest)(nil),  // 0: aoctracker.TrackRequest
	(*TrackResponse)(nil), // 1: aoctracker.TrackResponse
}
var file_aoctracker_proto_depIdxs = []int32{
	0, // 0: aoctracker.AOCTrackerService.Track:input_type -> aoctracker.TrackRequest
	1, // 1: aoctracker.AOCTrackerService.Track:output_type -> aoctracker.TrackResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aoctracker_proto_init() }
func file_aoctracker_proto_init() {
	if File_aoctracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aoctracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackRequest); i {
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
		file_aoctracker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackResponse); i {
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
			RawDescriptor: file_aoctracker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aoctracker_proto_goTypes,
		DependencyIndexes: file_aoctracker_proto_depIdxs,
		MessageInfos:      file_aoctracker_proto_msgTypes,
	}.Build()
	File_aoctracker_proto = out.File
	file_aoctracker_proto_rawDesc = nil
	file_aoctracker_proto_goTypes = nil
	file_aoctracker_proto_depIdxs = nil
}
