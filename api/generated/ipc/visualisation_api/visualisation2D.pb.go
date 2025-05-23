// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.6
// source: ipc/visualisation/visualisation2D.proto

package visualisation_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Particle2D struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PosX          float64                `protobuf:"fixed64,1,opt,name=pos_x,json=posX,proto3" json:"pos_x,omitempty"`
	PosY          float64                `protobuf:"fixed64,2,opt,name=pos_y,json=posY,proto3" json:"pos_y,omitempty"`
	Index         int64                  `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Particle2D) Reset() {
	*x = Particle2D{}
	mi := &file_ipc_visualisation_visualisation2D_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Particle2D) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Particle2D) ProtoMessage() {}

func (x *Particle2D) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_visualisation_visualisation2D_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Particle2D.ProtoReflect.Descriptor instead.
func (*Particle2D) Descriptor() ([]byte, []int) {
	return file_ipc_visualisation_visualisation2D_proto_rawDescGZIP(), []int{0}
}

func (x *Particle2D) GetPosX() float64 {
	if x != nil {
		return x.PosX
	}
	return 0
}

func (x *Particle2D) GetPosY() float64 {
	if x != nil {
		return x.PosY
	}
	return 0
}

func (x *Particle2D) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_ipc_visualisation_visualisation2D_proto protoreflect.FileDescriptor

const file_ipc_visualisation_visualisation2D_proto_rawDesc = "" +
	"\n" +
	"'ipc/visualisation/visualisation2D.proto\x12\x11ipc.visualisation\x1a\x1eipc/visualisation/common.proto\"L\n" +
	"\n" +
	"Particle2D\x12\x13\n" +
	"\x05pos_x\x18\x01 \x01(\x01R\x04posX\x12\x13\n" +
	"\x05pos_y\x18\x02 \x01(\x01R\x04posY\x12\x14\n" +
	"\x05index\x18\x03 \x01(\x03R\x05index2\xa9\x01\n" +
	"\x12Particle2DObserver\x12J\n" +
	"\x0fObserveParticle\x12\x1d.ipc.visualisation.Particle2D\x1a\x18.ipc.visualisation.Empty\x12G\n" +
	"\tCollision\x12 .ipc.visualisation.ParticleIndex\x1a\x18.ipc.visualisation.EmptyB$Z\"./generated/ipc/visualisation_api/b\x06proto3"

var (
	file_ipc_visualisation_visualisation2D_proto_rawDescOnce sync.Once
	file_ipc_visualisation_visualisation2D_proto_rawDescData []byte
)

func file_ipc_visualisation_visualisation2D_proto_rawDescGZIP() []byte {
	file_ipc_visualisation_visualisation2D_proto_rawDescOnce.Do(func() {
		file_ipc_visualisation_visualisation2D_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ipc_visualisation_visualisation2D_proto_rawDesc), len(file_ipc_visualisation_visualisation2D_proto_rawDesc)))
	})
	return file_ipc_visualisation_visualisation2D_proto_rawDescData
}

var file_ipc_visualisation_visualisation2D_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ipc_visualisation_visualisation2D_proto_goTypes = []any{
	(*Particle2D)(nil),    // 0: ipc.visualisation.Particle2D
	(*ParticleIndex)(nil), // 1: ipc.visualisation.ParticleIndex
	(*Empty)(nil),         // 2: ipc.visualisation.Empty
}
var file_ipc_visualisation_visualisation2D_proto_depIdxs = []int32{
	0, // 0: ipc.visualisation.Particle2DObserver.ObserveParticle:input_type -> ipc.visualisation.Particle2D
	1, // 1: ipc.visualisation.Particle2DObserver.Collision:input_type -> ipc.visualisation.ParticleIndex
	2, // 2: ipc.visualisation.Particle2DObserver.ObserveParticle:output_type -> ipc.visualisation.Empty
	2, // 3: ipc.visualisation.Particle2DObserver.Collision:output_type -> ipc.visualisation.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ipc_visualisation_visualisation2D_proto_init() }
func file_ipc_visualisation_visualisation2D_proto_init() {
	if File_ipc_visualisation_visualisation2D_proto != nil {
		return
	}
	file_ipc_visualisation_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ipc_visualisation_visualisation2D_proto_rawDesc), len(file_ipc_visualisation_visualisation2D_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ipc_visualisation_visualisation2D_proto_goTypes,
		DependencyIndexes: file_ipc_visualisation_visualisation2D_proto_depIdxs,
		MessageInfos:      file_ipc_visualisation_visualisation2D_proto_msgTypes,
	}.Build()
	File_ipc_visualisation_visualisation2D_proto = out.File
	file_ipc_visualisation_visualisation2D_proto_goTypes = nil
	file_ipc_visualisation_visualisation2D_proto_depIdxs = nil
}
