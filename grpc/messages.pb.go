// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: messages.proto

package messages

import (
	actor "github.com/asynkron/protoactor-go/actor"
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

type GRPCWeights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HiddenWeights []float32 `protobuf:"fixed32,1,rep,packed,name=hidden_weights,json=hiddenWeights,proto3" json:"hidden_weights,omitempty"`
	OutputWeights []float32 `protobuf:"fixed32,2,rep,packed,name=output_weights,json=outputWeights,proto3" json:"output_weights,omitempty"`
}

func (x *GRPCWeights) Reset() {
	*x = GRPCWeights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCWeights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCWeights) ProtoMessage() {}

func (x *GRPCWeights) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCWeights.ProtoReflect.Descriptor instead.
func (*GRPCWeights) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GRPCWeights) GetHiddenWeights() []float32 {
	if x != nil {
		return x.HiddenWeights
	}
	return nil
}

func (x *GRPCWeights) GetOutputWeights() []float32 {
	if x != nil {
		return x.OutputWeights
	}
	return nil
}

type GRPCExit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoordinatorPID *actor.PID `protobuf:"bytes,1,opt,name=coordinatorPID,proto3" json:"coordinatorPID,omitempty"`
	Address        string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port           int32      `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *GRPCExit) Reset() {
	*x = GRPCExit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCExit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCExit) ProtoMessage() {}

func (x *GRPCExit) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCExit.ProtoReflect.Descriptor instead.
func (*GRPCExit) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GRPCExit) GetCoordinatorPID() *actor.PID {
	if x != nil {
		return x.CoordinatorPID
	}
	return nil
}

func (x *GRPCExit) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GRPCExit) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type GRPCCollect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weights *GRPCWeights `protobuf:"bytes,1,opt,name=weights,proto3" json:"weights,omitempty"`
	Peers   int32        `protobuf:"varint,2,opt,name=peers,proto3" json:"peers,omitempty"`
}

func (x *GRPCCollect) Reset() {
	*x = GRPCCollect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCCollect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCCollect) ProtoMessage() {}

func (x *GRPCCollect) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCCollect.ProtoReflect.Descriptor instead.
func (*GRPCCollect) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GRPCCollect) GetWeights() *GRPCWeights {
	if x != nil {
		return x.Weights
	}
	return nil
}

func (x *GRPCCollect) GetPeers() int32 {
	if x != nil {
		return x.Peers
	}
	return 0
}

type GRPCAllPeersDone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GRPCAllPeersDone) Reset() {
	*x = GRPCAllPeersDone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCAllPeersDone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCAllPeersDone) ProtoMessage() {}

func (x *GRPCAllPeersDone) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCAllPeersDone.ProtoReflect.Descriptor instead.
func (*GRPCAllPeersDone) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x0b, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0b, 0x47, 0x52, 0x50, 0x43, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0d,
	0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x22, 0x6c, 0x0a, 0x08, 0x47, 0x52, 0x50, 0x43, 0x45, 0x78, 0x69, 0x74,
	0x12, 0x32, 0x0a, 0x0e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x50,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x50, 0x49, 0x44, 0x52, 0x0e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x50, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x54, 0x0a, 0x0b, 0x47, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x12, 0x2f, 0x0a, 0x07, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x47, 0x52,
	0x50, 0x43, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x07, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x52, 0x50, 0x43,
	0x41, 0x6c, 0x6c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x42, 0x0f, 0x5a, 0x0d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messages_proto_goTypes = []interface{}{
	(*GRPCWeights)(nil),      // 0: messages.GRPCWeights
	(*GRPCExit)(nil),         // 1: messages.GRPCExit
	(*GRPCCollect)(nil),      // 2: messages.GRPCCollect
	(*GRPCAllPeersDone)(nil), // 3: messages.GRPCAllPeersDone
	(*actor.PID)(nil),        // 4: actor.PID
}
var file_messages_proto_depIdxs = []int32{
	4, // 0: messages.GRPCExit.coordinatorPID:type_name -> actor.PID
	0, // 1: messages.GRPCCollect.weights:type_name -> messages.GRPCWeights
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCWeights); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCExit); i {
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
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCCollect); i {
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
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCAllPeersDone); i {
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
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
