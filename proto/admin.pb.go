// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: admin.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DataSource struct {
	FileOffset           uint64   `protobuf:"varint,1,opt,name=FileOffset,proto3" json:"FileOffset,omitempty"`
	PartitionID          uint64   `protobuf:"varint,2,opt,name=PartitionID,proto3" json:"PartitionID,omitempty"`
	ExtentID             uint64   `protobuf:"varint,3,opt,name=ExtentID,proto3" json:"ExtentID,omitempty"`
	ExtentOffset         uint64   `protobuf:"varint,4,opt,name=ExtentOffset,proto3" json:"ExtentOffset,omitempty"`
	Size_                uint64   `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`
	Hosts                []string `protobuf:"bytes,6,rep,name=Hosts,proto3" json:"Hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataSource) Reset()      { *m = DataSource{} }
func (*DataSource) ProtoMessage() {}
func (*DataSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{0}
}
func (m *DataSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSource.Merge(m, src)
}
func (m *DataSource) XXX_Size() int {
	return m.Size()
}
func (m *DataSource) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSource.DiscardUnknown(m)
}

var xxx_messageInfo_DataSource proto.InternalMessageInfo

func (m *DataSource) GetFileOffset() uint64 {
	if m != nil {
		return m.FileOffset
	}
	return 0
}

func (m *DataSource) GetPartitionID() uint64 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

func (m *DataSource) GetExtentID() uint64 {
	if m != nil {
		return m.ExtentID
	}
	return 0
}

func (m *DataSource) GetExtentOffset() uint64 {
	if m != nil {
		return m.ExtentOffset
	}
	return 0
}

func (m *DataSource) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *DataSource) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type CacheRequest struct {
	Volume               string        `protobuf:"bytes,1,opt,name=Volume,proto3" json:"Volume,omitempty"`
	Inode                uint64        `protobuf:"varint,2,opt,name=Inode,proto3" json:"Inode,omitempty"`
	FixedFileOffset      uint64        `protobuf:"varint,3,opt,name=FixedFileOffset,proto3" json:"FixedFileOffset,omitempty"`
	Version              uint32        `protobuf:"varint,4,opt,name=Version,proto3" json:"Version,omitempty"`
	Sources              []*DataSource `protobuf:"bytes,5,rep,name=Sources,proto3" json:"Sources,omitempty"`
	TTL                  int64         `protobuf:"varint,6,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CacheRequest) Reset()      { *m = CacheRequest{} }
func (*CacheRequest) ProtoMessage() {}
func (*CacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{1}
}
func (m *CacheRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CacheRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheRequest.Merge(m, src)
}
func (m *CacheRequest) XXX_Size() int {
	return m.Size()
}
func (m *CacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CacheRequest proto.InternalMessageInfo

func (m *CacheRequest) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *CacheRequest) GetInode() uint64 {
	if m != nil {
		return m.Inode
	}
	return 0
}

func (m *CacheRequest) GetFixedFileOffset() uint64 {
	if m != nil {
		return m.FixedFileOffset
	}
	return 0
}

func (m *CacheRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *CacheRequest) GetSources() []*DataSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *CacheRequest) GetTTL() int64 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type CacheReadRequest struct {
	CacheRequest         *CacheRequest `protobuf:"bytes,1,opt,name=CacheRequest,proto3" json:"CacheRequest,omitempty"`
	Offset               uint64        `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Size_                uint64        `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	Data                 []byte        `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CacheReadRequest) Reset()      { *m = CacheReadRequest{} }
func (*CacheReadRequest) ProtoMessage() {}
func (*CacheReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{2}
}
func (m *CacheReadRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CacheReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CacheReadRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CacheReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheReadRequest.Merge(m, src)
}
func (m *CacheReadRequest) XXX_Size() int {
	return m.Size()
}
func (m *CacheReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CacheReadRequest proto.InternalMessageInfo

func (m *CacheReadRequest) GetCacheRequest() *CacheRequest {
	if m != nil {
		return m.CacheRequest
	}
	return nil
}

func (m *CacheReadRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *CacheReadRequest) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

// Deprecated: Do not use.
func (m *CacheReadRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CachePrepareRequest struct {
	CacheRequest         *CacheRequest `protobuf:"bytes,1,opt,name=CacheRequest,proto3" json:"CacheRequest,omitempty"`
	FlashNodes           []string      `protobuf:"bytes,2,rep,name=FlashNodes,proto3" json:"FlashNodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CachePrepareRequest) Reset()      { *m = CachePrepareRequest{} }
func (*CachePrepareRequest) ProtoMessage() {}
func (*CachePrepareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{3}
}
func (m *CachePrepareRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CachePrepareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CachePrepareRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CachePrepareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CachePrepareRequest.Merge(m, src)
}
func (m *CachePrepareRequest) XXX_Size() int {
	return m.Size()
}
func (m *CachePrepareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CachePrepareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CachePrepareRequest proto.InternalMessageInfo

func (m *CachePrepareRequest) GetCacheRequest() *CacheRequest {
	if m != nil {
		return m.CacheRequest
	}
	return nil
}

func (m *CachePrepareRequest) GetFlashNodes() []string {
	if m != nil {
		return m.FlashNodes
	}
	return nil
}

func init() {
	proto.RegisterType((*DataSource)(nil), "proto.DataSource")
	proto.RegisterType((*CacheRequest)(nil), "proto.CacheRequest")
	proto.RegisterType((*CacheReadRequest)(nil), "proto.CacheReadRequest")
	proto.RegisterType((*CachePrepareRequest)(nil), "proto.CachePrepareRequest")
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor_73a7fc70dcc2027c) }

var fileDescriptor_73a7fc70dcc2027c = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4f, 0x8b, 0xd3, 0x40,
	0x1c, 0xdd, 0x69, 0xfe, 0xac, 0xfb, 0x4b, 0xc5, 0x3a, 0x2b, 0xcb, 0xb0, 0x87, 0x18, 0x7a, 0x0a,
	0x88, 0x5d, 0x58, 0x0f, 0x82, 0xc7, 0xb5, 0x2e, 0x16, 0x44, 0x97, 0xd9, 0x65, 0xef, 0xd3, 0x66,
	0xda, 0x0e, 0xb4, 0x99, 0x9a, 0x99, 0x40, 0xf1, 0x7b, 0x08, 0x3d, 0xfa, 0x49, 0xbc, 0x09, 0x1e,
	0xfd, 0x08, 0x52, 0xbf, 0x88, 0xe4, 0x37, 0x53, 0x4d, 0xeb, 0xd1, 0x53, 0x7e, 0xef, 0xfd, 0x26,
	0x2f, 0xef, 0xbd, 0x0c, 0x24, 0xa2, 0x58, 0xaa, 0x72, 0xb0, 0xaa, 0xb4, 0xd5, 0x34, 0xc2, 0xc7,
	0xf9, 0xf3, 0x99, 0xb2, 0xf3, 0x7a, 0x3c, 0x98, 0xe8, 0xe5, 0xc5, 0x4c, 0xcf, 0xf4, 0x05, 0xd2,
	0xe3, 0x7a, 0x8a, 0x08, 0x01, 0x4e, 0xee, 0xad, 0xfe, 0x57, 0x02, 0x30, 0x14, 0x56, 0xdc, 0xea,
	0xba, 0x9a, 0x48, 0x9a, 0x02, 0x5c, 0xab, 0x85, 0xfc, 0x30, 0x9d, 0x1a, 0x69, 0x19, 0xc9, 0x48,
	0x1e, 0xf2, 0x16, 0x43, 0x33, 0x48, 0x6e, 0x44, 0x65, 0x95, 0x55, 0xba, 0x1c, 0x0d, 0x59, 0x07,
	0x0f, 0xb4, 0x29, 0x7a, 0x0e, 0x0f, 0xde, 0xac, 0xad, 0x2c, 0xed, 0x68, 0xc8, 0x02, 0x5c, 0xff,
	0xc1, 0xb4, 0x0f, 0x5d, 0x37, 0x7b, 0xfd, 0x10, 0xf7, 0x7b, 0x1c, 0xa5, 0x10, 0xde, 0xaa, 0x4f,
	0x92, 0x45, 0xb8, 0xc3, 0x99, 0x3e, 0x81, 0xe8, 0xad, 0x36, 0xd6, 0xb0, 0x38, 0x0b, 0xf2, 0x13,
	0xee, 0xc0, 0xab, 0x70, 0xf3, 0xe5, 0xe9, 0x51, 0xff, 0x1b, 0x81, 0xee, 0x6b, 0x31, 0x99, 0x4b,
	0x2e, 0x3f, 0xd6, 0xd2, 0x58, 0x7a, 0x06, 0xf1, 0xbd, 0x5e, 0xd4, 0x4b, 0x89, 0xf6, 0x4f, 0xb8,
	0x47, 0x8d, 0xc8, 0xa8, 0xd4, 0x85, 0xf4, 0xa6, 0x1d, 0xa0, 0x39, 0x3c, 0xba, 0x56, 0x6b, 0x59,
	0xb4, 0x52, 0x3b, 0xd7, 0x87, 0x34, 0x65, 0x70, 0x7c, 0x2f, 0x2b, 0xa3, 0x74, 0x89, 0xbe, 0x1f,
	0xf2, 0x1d, 0xa4, 0xcf, 0xe0, 0xd8, 0xd5, 0x67, 0x58, 0x94, 0x05, 0x79, 0x72, 0xf9, 0xd8, 0x95,
	0x3b, 0xf8, 0x5b, 0x2c, 0xdf, 0x9d, 0xa0, 0x3d, 0x08, 0xee, 0xee, 0xde, 0xb1, 0x38, 0x23, 0x79,
	0xc0, 0x9b, 0xd1, 0xe7, 0xf8, 0x4c, 0xa0, 0xe7, 0x73, 0x88, 0x62, 0x97, 0xe5, 0xe5, 0x7e, 0x36,
	0x4c, 0x94, 0x5c, 0x9e, 0x7a, 0xf9, 0xf6, 0x8a, 0xff, 0x53, 0x82, 0x4f, 0xe3, 0xd2, 0xc6, 0x07,
	0xed, 0x06, 0xad, 0x76, 0xcf, 0x20, 0x6c, 0x8c, 0x62, 0xaa, 0xee, 0x55, 0x87, 0x11, 0x8e, 0xd8,
	0xfb, 0xb2, 0x70, 0x8a, 0xca, 0x37, 0x95, 0x5c, 0x89, 0x4a, 0xfe, 0xb7, 0xb3, 0xe6, 0x86, 0x2d,
	0x84, 0x99, 0xbf, 0xd7, 0x85, 0x34, 0xac, 0x83, 0x3f, 0xb4, 0xc5, 0xb8, 0xaf, 0x5e, 0xf5, 0xbe,
	0x6f, 0x53, 0xf2, 0x63, 0x9b, 0x92, 0x9f, 0xdb, 0x94, 0x6c, 0x7e, 0xa5, 0x47, 0xe3, 0x18, 0x95,
	0x5f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x28, 0xe1, 0xd4, 0xdf, 0xf4, 0x02, 0x00, 0x00,
}

func (m *DataSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataSource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataSource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Hosts) > 0 {
		for iNdEx := len(m.Hosts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Hosts[iNdEx])
			copy(dAtA[i:], m.Hosts[iNdEx])
			i = encodeVarintAdmin(dAtA, i, uint64(len(m.Hosts[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Size_ != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x28
	}
	if m.ExtentOffset != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.ExtentOffset))
		i--
		dAtA[i] = 0x20
	}
	if m.ExtentID != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.ExtentID))
		i--
		dAtA[i] = 0x18
	}
	if m.PartitionID != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.PartitionID))
		i--
		dAtA[i] = 0x10
	}
	if m.FileOffset != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.FileOffset))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CacheRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CacheRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TTL != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.TTL))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Sources) > 0 {
		for iNdEx := len(m.Sources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Sources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAdmin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Version != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x20
	}
	if m.FixedFileOffset != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.FixedFileOffset))
		i--
		dAtA[i] = 0x18
	}
	if m.Inode != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.Inode))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Volume) > 0 {
		i -= len(m.Volume)
		copy(dAtA[i:], m.Volume)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Volume)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CacheReadRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheReadRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CacheReadRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if m.Size_ != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x18
	}
	if m.Offset != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x10
	}
	if m.CacheRequest != nil {
		{
			size, err := m.CacheRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdmin(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CachePrepareRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CachePrepareRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CachePrepareRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.FlashNodes) > 0 {
		for iNdEx := len(m.FlashNodes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FlashNodes[iNdEx])
			copy(dAtA[i:], m.FlashNodes[iNdEx])
			i = encodeVarintAdmin(dAtA, i, uint64(len(m.FlashNodes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CacheRequest != nil {
		{
			size, err := m.CacheRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdmin(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdmin(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdmin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataSource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FileOffset != 0 {
		n += 1 + sovAdmin(uint64(m.FileOffset))
	}
	if m.PartitionID != 0 {
		n += 1 + sovAdmin(uint64(m.PartitionID))
	}
	if m.ExtentID != 0 {
		n += 1 + sovAdmin(uint64(m.ExtentID))
	}
	if m.ExtentOffset != 0 {
		n += 1 + sovAdmin(uint64(m.ExtentOffset))
	}
	if m.Size_ != 0 {
		n += 1 + sovAdmin(uint64(m.Size_))
	}
	if len(m.Hosts) > 0 {
		for _, s := range m.Hosts {
			l = len(s)
			n += 1 + l + sovAdmin(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CacheRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Volume)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.Inode != 0 {
		n += 1 + sovAdmin(uint64(m.Inode))
	}
	if m.FixedFileOffset != 0 {
		n += 1 + sovAdmin(uint64(m.FixedFileOffset))
	}
	if m.Version != 0 {
		n += 1 + sovAdmin(uint64(m.Version))
	}
	if len(m.Sources) > 0 {
		for _, e := range m.Sources {
			l = e.Size()
			n += 1 + l + sovAdmin(uint64(l))
		}
	}
	if m.TTL != 0 {
		n += 1 + sovAdmin(uint64(m.TTL))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CacheReadRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CacheRequest != nil {
		l = m.CacheRequest.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovAdmin(uint64(m.Offset))
	}
	if m.Size_ != 0 {
		n += 1 + sovAdmin(uint64(m.Size_))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CachePrepareRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CacheRequest != nil {
		l = m.CacheRequest.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	if len(m.FlashNodes) > 0 {
		for _, s := range m.FlashNodes {
			l = len(s)
			n += 1 + l + sovAdmin(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAdmin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdmin(x uint64) (n int) {
	return sovAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileOffset", wireType)
			}
			m.FileOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FileOffset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartitionID", wireType)
			}
			m.PartitionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PartitionID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtentID", wireType)
			}
			m.ExtentID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExtentID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtentOffset", wireType)
			}
			m.ExtentOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExtentOffset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hosts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hosts = append(m.Hosts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CacheRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CacheRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volume", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Volume = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inode", wireType)
			}
			m.Inode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Inode |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedFileOffset", wireType)
			}
			m.FixedFileOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FixedFileOffset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sources = append(m.Sources, &DataSource{})
			if err := m.Sources[len(m.Sources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CacheReadRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CacheReadRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheReadRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CacheRequest == nil {
				m.CacheRequest = &CacheRequest{}
			}
			if err := m.CacheRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CachePrepareRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CachePrepareRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CachePrepareRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CacheRequest == nil {
				m.CacheRequest = &CacheRequest{}
			}
			if err := m.CacheRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlashNodes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FlashNodes = append(m.FlashNodes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdmin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAdmin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdmin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdmin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdmin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdmin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdmin = fmt.Errorf("proto: unexpected end of group")
)