// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: vince/v1/store.proto

package v1

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

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns []*Metadata_Column `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Min     uint64             `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max     uint64             `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Id      string             `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_v1_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_vince_v1_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_vince_v1_store_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetColumns() []*Metadata_Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Metadata) GetMin() uint64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Metadata) GetMax() uint64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Metadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PrimaryIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources map[string]*PrimaryIndex_Resource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PrimaryIndex) Reset() {
	*x = PrimaryIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_v1_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryIndex) ProtoMessage() {}

func (x *PrimaryIndex) ProtoReflect() protoreflect.Message {
	mi := &file_vince_v1_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryIndex.ProtoReflect.Descriptor instead.
func (*PrimaryIndex) Descriptor() ([]byte, []int) {
	return file_vince_v1_store_proto_rawDescGZIP(), []int{1}
}

func (x *PrimaryIndex) GetResources() map[string]*PrimaryIndex_Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type Granule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Min  int64  `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max  int64  `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Size uint64 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Rows uint64 `protobuf:"varint,5,opt,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Granule) Reset() {
	*x = Granule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_v1_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Granule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Granule) ProtoMessage() {}

func (x *Granule) ProtoReflect() protoreflect.Message {
	mi := &file_vince_v1_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Granule.ProtoReflect.Descriptor instead.
func (*Granule) Descriptor() ([]byte, []int) {
	return file_vince_v1_store_proto_rawDescGZIP(), []int{2}
}

func (x *Granule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Granule) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Granule) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Granule) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Granule) GetRows() uint64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

type Metadata_Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NumRows   uint32 `protobuf:"varint,2,opt,name=num_rows,json=numRows,proto3" json:"num_rows,omitempty"`
	FstOffset uint32 `protobuf:"varint,4,opt,name=fst_offset,json=fstOffset,proto3" json:"fst_offset,omitempty"`
	Offset    uint64 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	Size      uint32 `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	RawSize   uint32 `protobuf:"varint,7,opt,name=raw_size,json=rawSize,proto3" json:"raw_size,omitempty"`
}

func (x *Metadata_Column) Reset() {
	*x = Metadata_Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_v1_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata_Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata_Column) ProtoMessage() {}

func (x *Metadata_Column) ProtoReflect() protoreflect.Message {
	mi := &file_vince_v1_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata_Column.ProtoReflect.Descriptor instead.
func (*Metadata_Column) Descriptor() ([]byte, []int) {
	return file_vince_v1_store_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Metadata_Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata_Column) GetNumRows() uint32 {
	if x != nil {
		return x.NumRows
	}
	return 0
}

func (x *Metadata_Column) GetFstOffset() uint32 {
	if x != nil {
		return x.FstOffset
	}
	return 0
}

func (x *Metadata_Column) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Metadata_Column) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Metadata_Column) GetRawSize() uint32 {
	if x != nil {
		return x.RawSize
	}
	return 0
}

type Metadata_Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Metadata_Chunk) Reset() {
	*x = Metadata_Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_v1_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata_Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata_Chunk) ProtoMessage() {}

func (x *Metadata_Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_vince_v1_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata_Chunk.ProtoReflect.Descriptor instead.
func (*Metadata_Chunk) Descriptor() ([]byte, []int) {
	return file_vince_v1_store_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Metadata_Chunk) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Metadata_Chunk) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type PrimaryIndex_Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Granules map[string]*Granule `protobuf:"bytes,2,rep,name=granules,proto3" json:"granules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PrimaryIndex_Resource) Reset() {
	*x = PrimaryIndex_Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_v1_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryIndex_Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryIndex_Resource) ProtoMessage() {}

func (x *PrimaryIndex_Resource) ProtoReflect() protoreflect.Message {
	mi := &file_vince_v1_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryIndex_Resource.ProtoReflect.Descriptor instead.
func (*PrimaryIndex_Resource) Descriptor() ([]byte, []int) {
	return file_vince_v1_store_proto_rawDescGZIP(), []int{1, 1}
}

func (x *PrimaryIndex_Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrimaryIndex_Resource) GetGranules() map[string]*Granule {
	if x != nil {
		return x.Granules
	}
	return nil
}

var File_vince_v1_store_proto protoreflect.FileDescriptor

var file_vince_v1_store_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0xc2, 0x02, 0x0a, 0x08, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x9d, 0x01, 0x0a, 0x06, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75, 0x6d,
	0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x75, 0x6d,
	0x52, 0x6f, 0x77, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x73, 0x74, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x61, 0x77, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x33, 0x0a, 0x05, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0xd6, 0x02, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x3d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x57, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xad, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x67, 0x72, 0x61,
	0x6e, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x48,
	0x0a, 0x0d, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x65, 0x0a, 0x07, 0x47, 0x72, 0x61, 0x6e,
	0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x42,
	0x6b, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x74, 0x73, 0x75, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56,
	0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vince_v1_store_proto_rawDescOnce sync.Once
	file_vince_v1_store_proto_rawDescData = file_vince_v1_store_proto_rawDesc
)

func file_vince_v1_store_proto_rawDescGZIP() []byte {
	file_vince_v1_store_proto_rawDescOnce.Do(func() {
		file_vince_v1_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_vince_v1_store_proto_rawDescData)
	})
	return file_vince_v1_store_proto_rawDescData
}

var file_vince_v1_store_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_vince_v1_store_proto_goTypes = []interface{}{
	(*Metadata)(nil),              // 0: v1.Metadata
	(*PrimaryIndex)(nil),          // 1: v1.PrimaryIndex
	(*Granule)(nil),               // 2: v1.Granule
	(*Metadata_Column)(nil),       // 3: v1.Metadata.Column
	(*Metadata_Chunk)(nil),        // 4: v1.Metadata.Chunk
	nil,                           // 5: v1.PrimaryIndex.ResourcesEntry
	(*PrimaryIndex_Resource)(nil), // 6: v1.PrimaryIndex.Resource
	nil,                           // 7: v1.PrimaryIndex.Resource.GranulesEntry
}
var file_vince_v1_store_proto_depIdxs = []int32{
	3, // 0: v1.Metadata.columns:type_name -> v1.Metadata.Column
	5, // 1: v1.PrimaryIndex.resources:type_name -> v1.PrimaryIndex.ResourcesEntry
	6, // 2: v1.PrimaryIndex.ResourcesEntry.value:type_name -> v1.PrimaryIndex.Resource
	7, // 3: v1.PrimaryIndex.Resource.granules:type_name -> v1.PrimaryIndex.Resource.GranulesEntry
	2, // 4: v1.PrimaryIndex.Resource.GranulesEntry.value:type_name -> v1.Granule
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_vince_v1_store_proto_init() }
func file_vince_v1_store_proto_init() {
	if File_vince_v1_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vince_v1_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_vince_v1_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryIndex); i {
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
		file_vince_v1_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Granule); i {
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
		file_vince_v1_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata_Column); i {
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
		file_vince_v1_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata_Chunk); i {
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
		file_vince_v1_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryIndex_Resource); i {
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
			RawDescriptor: file_vince_v1_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vince_v1_store_proto_goTypes,
		DependencyIndexes: file_vince_v1_store_proto_depIdxs,
		MessageInfos:      file_vince_v1_store_proto_msgTypes,
	}.Build()
	File_vince_v1_store_proto = out.File
	file_vince_v1_store_proto_rawDesc = nil
	file_vince_v1_store_proto_goTypes = nil
	file_vince_v1_store_proto_depIdxs = nil
}