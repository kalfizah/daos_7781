// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ctl/storage_scm.proto

package ctl

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// ScmModule represent Storage Class Memory modules installed.
type ScmModule struct {
	Channelid            uint32   `protobuf:"varint,1,opt,name=channelid,proto3" json:"channelid,omitempty"`
	Channelposition      uint32   `protobuf:"varint,2,opt,name=channelposition,proto3" json:"channelposition,omitempty"`
	Controllerid         uint32   `protobuf:"varint,3,opt,name=controllerid,proto3" json:"controllerid,omitempty"`
	Socketid             uint32   `protobuf:"varint,4,opt,name=socketid,proto3" json:"socketid,omitempty"`
	Physicalid           uint32   `protobuf:"varint,5,opt,name=physicalid,proto3" json:"physicalid,omitempty"`
	Capacity             uint64   `protobuf:"varint,6,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Uid                  string   `protobuf:"bytes,7,opt,name=uid,proto3" json:"uid,omitempty"`
	PartNumber           string   `protobuf:"bytes,8,opt,name=partNumber,proto3" json:"partNumber,omitempty"`
	FirmwareRevision     string   `protobuf:"bytes,9,opt,name=firmwareRevision,proto3" json:"firmwareRevision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScmModule) Reset()         { *m = ScmModule{} }
func (m *ScmModule) String() string { return proto.CompactTextString(m) }
func (*ScmModule) ProtoMessage()    {}
func (*ScmModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{0}
}

func (m *ScmModule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmModule.Unmarshal(m, b)
}
func (m *ScmModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmModule.Marshal(b, m, deterministic)
}
func (m *ScmModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmModule.Merge(m, src)
}
func (m *ScmModule) XXX_Size() int {
	return xxx_messageInfo_ScmModule.Size(m)
}
func (m *ScmModule) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmModule.DiscardUnknown(m)
}

var xxx_messageInfo_ScmModule proto.InternalMessageInfo

func (m *ScmModule) GetChannelid() uint32 {
	if m != nil {
		return m.Channelid
	}
	return 0
}

func (m *ScmModule) GetChannelposition() uint32 {
	if m != nil {
		return m.Channelposition
	}
	return 0
}

func (m *ScmModule) GetControllerid() uint32 {
	if m != nil {
		return m.Controllerid
	}
	return 0
}

func (m *ScmModule) GetSocketid() uint32 {
	if m != nil {
		return m.Socketid
	}
	return 0
}

func (m *ScmModule) GetPhysicalid() uint32 {
	if m != nil {
		return m.Physicalid
	}
	return 0
}

func (m *ScmModule) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *ScmModule) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ScmModule) GetPartNumber() string {
	if m != nil {
		return m.PartNumber
	}
	return ""
}

func (m *ScmModule) GetFirmwareRevision() string {
	if m != nil {
		return m.FirmwareRevision
	}
	return ""
}

// ScmNamespace represents SCM namespace as pmem device files created on a ScmRegion.
type ScmNamespace struct {
	Uuid                 string              `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Blockdev             string              `protobuf:"bytes,2,opt,name=blockdev,proto3" json:"blockdev,omitempty"`
	Dev                  string              `protobuf:"bytes,3,opt,name=dev,proto3" json:"dev,omitempty"`
	NumaNode             uint32              `protobuf:"varint,4,opt,name=numa_node,json=numaNode,proto3" json:"numa_node,omitempty"`
	Size                 uint64              `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Mount                *ScmNamespace_Mount `protobuf:"bytes,6,opt,name=mount,proto3" json:"mount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ScmNamespace) Reset()         { *m = ScmNamespace{} }
func (m *ScmNamespace) String() string { return proto.CompactTextString(m) }
func (*ScmNamespace) ProtoMessage()    {}
func (*ScmNamespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{1}
}

func (m *ScmNamespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmNamespace.Unmarshal(m, b)
}
func (m *ScmNamespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmNamespace.Marshal(b, m, deterministic)
}
func (m *ScmNamespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmNamespace.Merge(m, src)
}
func (m *ScmNamespace) XXX_Size() int {
	return xxx_messageInfo_ScmNamespace.Size(m)
}
func (m *ScmNamespace) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmNamespace.DiscardUnknown(m)
}

var xxx_messageInfo_ScmNamespace proto.InternalMessageInfo

func (m *ScmNamespace) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ScmNamespace) GetBlockdev() string {
	if m != nil {
		return m.Blockdev
	}
	return ""
}

func (m *ScmNamespace) GetDev() string {
	if m != nil {
		return m.Dev
	}
	return ""
}

func (m *ScmNamespace) GetNumaNode() uint32 {
	if m != nil {
		return m.NumaNode
	}
	return 0
}

func (m *ScmNamespace) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ScmNamespace) GetMount() *ScmNamespace_Mount {
	if m != nil {
		return m.Mount
	}
	return nil
}

// Mount represents a mounted pmem block device.
type ScmNamespace_Mount struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	TotalBytes           uint64   `protobuf:"varint,2,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	AvailBytes           uint64   `protobuf:"varint,3,opt,name=avail_bytes,json=availBytes,proto3" json:"avail_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScmNamespace_Mount) Reset()         { *m = ScmNamespace_Mount{} }
func (m *ScmNamespace_Mount) String() string { return proto.CompactTextString(m) }
func (*ScmNamespace_Mount) ProtoMessage()    {}
func (*ScmNamespace_Mount) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{1, 0}
}

func (m *ScmNamespace_Mount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmNamespace_Mount.Unmarshal(m, b)
}
func (m *ScmNamespace_Mount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmNamespace_Mount.Marshal(b, m, deterministic)
}
func (m *ScmNamespace_Mount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmNamespace_Mount.Merge(m, src)
}
func (m *ScmNamespace_Mount) XXX_Size() int {
	return xxx_messageInfo_ScmNamespace_Mount.Size(m)
}
func (m *ScmNamespace_Mount) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmNamespace_Mount.DiscardUnknown(m)
}

var xxx_messageInfo_ScmNamespace_Mount proto.InternalMessageInfo

func (m *ScmNamespace_Mount) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ScmNamespace_Mount) GetTotalBytes() uint64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

func (m *ScmNamespace_Mount) GetAvailBytes() uint64 {
	if m != nil {
		return m.AvailBytes
	}
	return 0
}

// ScmModuleResult represents operation state for specific SCM/PM module.
//
// TODO: replace identifier with serial when returned in scan
type ScmModuleResult struct {
	Physicalid           uint32         `protobuf:"varint,1,opt,name=physicalid,proto3" json:"physicalid,omitempty"`
	State                *ResponseState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ScmModuleResult) Reset()         { *m = ScmModuleResult{} }
func (m *ScmModuleResult) String() string { return proto.CompactTextString(m) }
func (*ScmModuleResult) ProtoMessage()    {}
func (*ScmModuleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{2}
}

func (m *ScmModuleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmModuleResult.Unmarshal(m, b)
}
func (m *ScmModuleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmModuleResult.Marshal(b, m, deterministic)
}
func (m *ScmModuleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmModuleResult.Merge(m, src)
}
func (m *ScmModuleResult) XXX_Size() int {
	return xxx_messageInfo_ScmModuleResult.Size(m)
}
func (m *ScmModuleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmModuleResult.DiscardUnknown(m)
}

var xxx_messageInfo_ScmModuleResult proto.InternalMessageInfo

func (m *ScmModuleResult) GetPhysicalid() uint32 {
	if m != nil {
		return m.Physicalid
	}
	return 0
}

func (m *ScmModuleResult) GetState() *ResponseState {
	if m != nil {
		return m.State
	}
	return nil
}

// ScmMountResult represents operation state for specific SCM mount point.
type ScmMountResult struct {
	Mntpoint             string         `protobuf:"bytes,1,opt,name=mntpoint,proto3" json:"mntpoint,omitempty"`
	State                *ResponseState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Instanceidx          uint32         `protobuf:"varint,3,opt,name=instanceidx,proto3" json:"instanceidx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ScmMountResult) Reset()         { *m = ScmMountResult{} }
func (m *ScmMountResult) String() string { return proto.CompactTextString(m) }
func (*ScmMountResult) ProtoMessage()    {}
func (*ScmMountResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{3}
}

func (m *ScmMountResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmMountResult.Unmarshal(m, b)
}
func (m *ScmMountResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmMountResult.Marshal(b, m, deterministic)
}
func (m *ScmMountResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmMountResult.Merge(m, src)
}
func (m *ScmMountResult) XXX_Size() int {
	return xxx_messageInfo_ScmMountResult.Size(m)
}
func (m *ScmMountResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmMountResult.DiscardUnknown(m)
}

var xxx_messageInfo_ScmMountResult proto.InternalMessageInfo

func (m *ScmMountResult) GetMntpoint() string {
	if m != nil {
		return m.Mntpoint
	}
	return ""
}

func (m *ScmMountResult) GetState() *ResponseState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ScmMountResult) GetInstanceidx() uint32 {
	if m != nil {
		return m.Instanceidx
	}
	return 0
}

type PrepareScmReq struct {
	Reset_               bool     `protobuf:"varint,1,opt,name=reset,proto3" json:"reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareScmReq) Reset()         { *m = PrepareScmReq{} }
func (m *PrepareScmReq) String() string { return proto.CompactTextString(m) }
func (*PrepareScmReq) ProtoMessage()    {}
func (*PrepareScmReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{4}
}

func (m *PrepareScmReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareScmReq.Unmarshal(m, b)
}
func (m *PrepareScmReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareScmReq.Marshal(b, m, deterministic)
}
func (m *PrepareScmReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareScmReq.Merge(m, src)
}
func (m *PrepareScmReq) XXX_Size() int {
	return xxx_messageInfo_PrepareScmReq.Size(m)
}
func (m *PrepareScmReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareScmReq.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareScmReq proto.InternalMessageInfo

func (m *PrepareScmReq) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

type PrepareScmResp struct {
	Namespaces           []*ScmNamespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	State                *ResponseState  `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Rebootrequired       bool            `protobuf:"varint,3,opt,name=rebootrequired,proto3" json:"rebootrequired,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PrepareScmResp) Reset()         { *m = PrepareScmResp{} }
func (m *PrepareScmResp) String() string { return proto.CompactTextString(m) }
func (*PrepareScmResp) ProtoMessage()    {}
func (*PrepareScmResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{5}
}

func (m *PrepareScmResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareScmResp.Unmarshal(m, b)
}
func (m *PrepareScmResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareScmResp.Marshal(b, m, deterministic)
}
func (m *PrepareScmResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareScmResp.Merge(m, src)
}
func (m *PrepareScmResp) XXX_Size() int {
	return xxx_messageInfo_PrepareScmResp.Size(m)
}
func (m *PrepareScmResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareScmResp.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareScmResp proto.InternalMessageInfo

func (m *PrepareScmResp) GetNamespaces() []*ScmNamespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *PrepareScmResp) GetState() *ResponseState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *PrepareScmResp) GetRebootrequired() bool {
	if m != nil {
		return m.Rebootrequired
	}
	return false
}

type ScanScmReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanScmReq) Reset()         { *m = ScanScmReq{} }
func (m *ScanScmReq) String() string { return proto.CompactTextString(m) }
func (*ScanScmReq) ProtoMessage()    {}
func (*ScanScmReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{6}
}

func (m *ScanScmReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanScmReq.Unmarshal(m, b)
}
func (m *ScanScmReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanScmReq.Marshal(b, m, deterministic)
}
func (m *ScanScmReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanScmReq.Merge(m, src)
}
func (m *ScanScmReq) XXX_Size() int {
	return xxx_messageInfo_ScanScmReq.Size(m)
}
func (m *ScanScmReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanScmReq.DiscardUnknown(m)
}

var xxx_messageInfo_ScanScmReq proto.InternalMessageInfo

type ScanScmResp struct {
	Modules              []*ScmModule    `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
	Namespaces           []*ScmNamespace `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	State                *ResponseState  `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ScanScmResp) Reset()         { *m = ScanScmResp{} }
func (m *ScanScmResp) String() string { return proto.CompactTextString(m) }
func (*ScanScmResp) ProtoMessage()    {}
func (*ScanScmResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{7}
}

func (m *ScanScmResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanScmResp.Unmarshal(m, b)
}
func (m *ScanScmResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanScmResp.Marshal(b, m, deterministic)
}
func (m *ScanScmResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanScmResp.Merge(m, src)
}
func (m *ScanScmResp) XXX_Size() int {
	return xxx_messageInfo_ScanScmResp.Size(m)
}
func (m *ScanScmResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanScmResp.DiscardUnknown(m)
}

var xxx_messageInfo_ScanScmResp proto.InternalMessageInfo

func (m *ScanScmResp) GetModules() []*ScmModule {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *ScanScmResp) GetNamespaces() []*ScmNamespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *ScanScmResp) GetState() *ResponseState {
	if m != nil {
		return m.State
	}
	return nil
}

type FormatScmReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FormatScmReq) Reset()         { *m = FormatScmReq{} }
func (m *FormatScmReq) String() string { return proto.CompactTextString(m) }
func (*FormatScmReq) ProtoMessage()    {}
func (*FormatScmReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd3a540aea514928, []int{8}
}

func (m *FormatScmReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FormatScmReq.Unmarshal(m, b)
}
func (m *FormatScmReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FormatScmReq.Marshal(b, m, deterministic)
}
func (m *FormatScmReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FormatScmReq.Merge(m, src)
}
func (m *FormatScmReq) XXX_Size() int {
	return xxx_messageInfo_FormatScmReq.Size(m)
}
func (m *FormatScmReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FormatScmReq.DiscardUnknown(m)
}

var xxx_messageInfo_FormatScmReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScmModule)(nil), "ctl.ScmModule")
	proto.RegisterType((*ScmNamespace)(nil), "ctl.ScmNamespace")
	proto.RegisterType((*ScmNamespace_Mount)(nil), "ctl.ScmNamespace.Mount")
	proto.RegisterType((*ScmModuleResult)(nil), "ctl.ScmModuleResult")
	proto.RegisterType((*ScmMountResult)(nil), "ctl.ScmMountResult")
	proto.RegisterType((*PrepareScmReq)(nil), "ctl.PrepareScmReq")
	proto.RegisterType((*PrepareScmResp)(nil), "ctl.PrepareScmResp")
	proto.RegisterType((*ScanScmReq)(nil), "ctl.ScanScmReq")
	proto.RegisterType((*ScanScmResp)(nil), "ctl.ScanScmResp")
	proto.RegisterType((*FormatScmReq)(nil), "ctl.FormatScmReq")
}

func init() {
	proto.RegisterFile("ctl/storage_scm.proto", fileDescriptor_dd3a540aea514928)
}

var fileDescriptor_dd3a540aea514928 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6f, 0x13, 0x3b,
	0x10, 0xd6, 0x66, 0x93, 0x36, 0x3b, 0x49, 0xd3, 0x3e, 0xeb, 0x3d, 0xbd, 0x55, 0x40, 0x10, 0xad,
	0x04, 0x5a, 0x21, 0x35, 0x11, 0xe5, 0x80, 0xb8, 0xf6, 0xc0, 0xad, 0x15, 0x72, 0x6e, 0x20, 0x54,
	0x39, 0xde, 0xa1, 0xb1, 0xba, 0xb6, 0xb7, 0xb6, 0xb7, 0xb4, 0xfc, 0x09, 0x2e, 0xf0, 0x33, 0xf8,
	0x8f, 0xc8, 0xde, 0xdd, 0x10, 0xd2, 0x03, 0xf4, 0x36, 0xdf, 0x37, 0xdf, 0x8e, 0xbf, 0x99, 0xf1,
	0x1a, 0xfe, 0xe3, 0xae, 0x5c, 0x58, 0xa7, 0x0d, 0xbb, 0xc4, 0x0b, 0xcb, 0xe5, 0xbc, 0x32, 0xda,
	0x69, 0x12, 0x73, 0x57, 0x4e, 0x8f, 0x7c, 0x8e, 0x6b, 0x29, 0xb5, 0x6a, 0xe8, 0xec, 0x47, 0x0f,
	0x92, 0x25, 0x97, 0x67, 0xba, 0xa8, 0x4b, 0x24, 0x8f, 0x21, 0xe1, 0x6b, 0xa6, 0x14, 0x96, 0xa2,
	0x48, 0xa3, 0x59, 0x94, 0x1f, 0xd0, 0x5f, 0x04, 0xc9, 0xe1, 0xb0, 0x05, 0x95, 0xb6, 0xc2, 0x09,
	0xad, 0xd2, 0x5e, 0xd0, 0xec, 0xd2, 0x24, 0x83, 0x31, 0xd7, 0xca, 0x19, 0x5d, 0x96, 0x68, 0x44,
	0x91, 0xc6, 0x41, 0xf6, 0x1b, 0x47, 0xa6, 0x30, 0xb4, 0x9a, 0x5f, 0xa1, 0x13, 0x45, 0xda, 0x0f,
	0xf9, 0x0d, 0x26, 0x4f, 0x00, 0xaa, 0xf5, 0x9d, 0x15, 0x9c, 0x79, 0x23, 0x83, 0x90, 0xdd, 0x62,
	0xfc, 0xb7, 0x9c, 0x55, 0x8c, 0x0b, 0x77, 0x97, 0xee, 0xcd, 0xa2, 0xbc, 0x4f, 0x37, 0x98, 0x1c,
	0x41, 0x5c, 0x8b, 0x22, 0xdd, 0x9f, 0x45, 0x79, 0x42, 0x7d, 0x18, 0xaa, 0x31, 0xe3, 0xce, 0x6b,
	0xb9, 0x42, 0x93, 0x0e, 0x43, 0x62, 0x8b, 0x21, 0x2f, 0xe0, 0xe8, 0x93, 0x30, 0xf2, 0x33, 0x33,
	0x48, 0xf1, 0x46, 0x58, 0xdf, 0x58, 0x12, 0x54, 0xf7, 0xf8, 0xec, 0x6b, 0x0f, 0xc6, 0x4b, 0x2e,
	0xcf, 0x99, 0x44, 0x5b, 0x31, 0x8e, 0x84, 0x40, 0xbf, 0xae, 0xdb, 0x69, 0x25, 0x34, 0xc4, 0xde,
	0xde, 0xaa, 0xd4, 0xfc, 0xaa, 0xc0, 0x9b, 0x30, 0xa1, 0x84, 0x6e, 0xb0, 0xb7, 0xe7, 0xe9, 0xb8,
	0xb1, 0xe7, 0x99, 0x47, 0x90, 0xa8, 0x5a, 0xb2, 0x0b, 0xa5, 0x0b, 0xec, 0x26, 0xe1, 0x89, 0x73,
	0x5d, 0x84, 0xf2, 0x56, 0x7c, 0xc1, 0x30, 0x83, 0x3e, 0x0d, 0x31, 0x39, 0x86, 0x81, 0xd4, 0xb5,
	0x72, 0xa1, 0xf5, 0xd1, 0xc9, 0xff, 0x73, 0xee, 0xca, 0xf9, 0xb6, 0xa9, 0xf9, 0x99, 0x4f, 0xd3,
	0x46, 0x35, 0xfd, 0x08, 0x83, 0x80, 0x7d, 0xad, 0x8a, 0xb9, 0x75, 0x67, 0xd5, 0xc7, 0xe4, 0x29,
	0x8c, 0x9c, 0x76, 0xac, 0xbc, 0x58, 0xdd, 0x39, 0xb4, 0xc1, 0x6d, 0x9f, 0x42, 0xa0, 0x4e, 0x3d,
	0xe3, 0x05, 0xec, 0x86, 0x89, 0x4e, 0x10, 0x37, 0x82, 0x40, 0x05, 0x41, 0xf6, 0x01, 0x0e, 0x37,
	0x17, 0x88, 0xa2, 0xad, 0x4b, 0xb7, 0xb3, 0xbe, 0xe8, 0xde, 0xfa, 0x72, 0x18, 0x58, 0xc7, 0x1c,
	0x86, 0xe3, 0x46, 0x27, 0x24, 0x34, 0x40, 0xd1, 0x56, 0x5a, 0x59, 0x5c, 0xfa, 0x0c, 0x6d, 0x04,
	0xd9, 0x2d, 0x4c, 0x42, 0x71, 0xdf, 0x4e, 0x53, 0x7b, 0x0a, 0x43, 0xa9, 0x5c, 0xa5, 0x85, 0x72,
	0x6d, 0x23, 0x1b, 0xfc, 0xf7, 0x75, 0xc9, 0x0c, 0x46, 0x42, 0x59, 0xc7, 0x14, 0x47, 0x51, 0xdc,
	0xb6, 0xf7, 0x73, 0x9b, 0xca, 0x9e, 0xc1, 0xc1, 0x3b, 0x83, 0x15, 0x33, 0xb8, 0xe4, 0x92, 0xe2,
	0x35, 0xf9, 0x17, 0x06, 0x06, 0x2d, 0x36, 0xa7, 0x0e, 0x69, 0x03, 0xb2, 0xef, 0x11, 0x4c, 0xb6,
	0x75, 0xb6, 0x22, 0x2f, 0x01, 0x54, 0xb7, 0x09, 0x9b, 0x46, 0xb3, 0x38, 0x1f, 0x9d, 0xfc, 0x73,
	0x6f, 0x47, 0x74, 0x4b, 0xf4, 0x00, 0xe3, 0xcf, 0x61, 0x62, 0x70, 0xa5, 0xb5, 0x33, 0x78, 0x5d,
	0x0b, 0x83, 0xcd, 0xbf, 0x35, 0xa4, 0x3b, 0x6c, 0x36, 0x06, 0x58, 0x72, 0xa6, 0x1a, 0xef, 0xd9,
	0xb7, 0x08, 0x46, 0x1b, 0x68, 0x2b, 0x92, 0xc3, 0xbe, 0x0c, 0x0b, 0xeb, 0xfc, 0x4d, 0x3a, 0x7f,
	0xed, 0x1e, 0xbb, 0xf4, 0x4e, 0x33, 0xbd, 0x07, 0x35, 0x13, 0xff, 0x69, 0xbb, 0x13, 0x18, 0xbf,
	0xd5, 0x46, 0x32, 0xd7, 0xd8, 0x3c, 0x7d, 0xf3, 0xfe, 0xf5, 0xa5, 0x70, 0xeb, 0x7a, 0x35, 0xe7,
	0x5a, 0x2e, 0x0a, 0xa6, 0xed, 0xb1, 0x75, 0x8c, 0x5f, 0x85, 0x70, 0x61, 0x0d, 0x5f, 0xb4, 0x2f,
	0x48, 0xfb, 0x86, 0x2d, 0xc2, 0x1b, 0xb6, 0xe0, 0xae, 0x5c, 0xed, 0x85, 0xf0, 0xd5, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa1, 0xa1, 0xfc, 0x29, 0xfe, 0x04, 0x00, 0x00,
}
