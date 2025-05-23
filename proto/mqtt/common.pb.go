// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/mqtt/common.proto

package omo_msp_mqtt

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

type RequestInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 string   `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	User                 string   `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ac4d88d62e40be, []int{0}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestInfo) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type RequestFilter struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Field                string   `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32   `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Values               []string `protobuf:"bytes,7,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFilter) Reset()         { *m = RequestFilter{} }
func (m *RequestFilter) String() string { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()    {}
func (*RequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ac4d88d62e40be, []int{1}
}

func (m *RequestFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFilter.Unmarshal(m, b)
}
func (m *RequestFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFilter.Marshal(b, m, deterministic)
}
func (m *RequestFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFilter.Merge(m, src)
}
func (m *RequestFilter) XXX_Size() int {
	return xxx_messageInfo_RequestFilter.Size(m)
}
func (m *RequestFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFilter proto.InternalMessageInfo

func (m *RequestFilter) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestFilter) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *RequestFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestFilter) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RequestFilter) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RequestFilter) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestFilter) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type RequestUpdate struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Field                string   `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Values               []string `protobuf:"bytes,6,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUpdate) Reset()         { *m = RequestUpdate{} }
func (m *RequestUpdate) String() string { return proto.CompactTextString(m) }
func (*RequestUpdate) ProtoMessage()    {}
func (*RequestUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ac4d88d62e40be, []int{2}
}

func (m *RequestUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUpdate.Unmarshal(m, b)
}
func (m *RequestUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUpdate.Marshal(b, m, deterministic)
}
func (m *RequestUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUpdate.Merge(m, src)
}
func (m *RequestUpdate) XXX_Size() int {
	return xxx_messageInfo_RequestUpdate.Size(m)
}
func (m *RequestUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUpdate proto.InternalMessageInfo

func (m *RequestUpdate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestUpdate) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *RequestUpdate) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestUpdate) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type ReplyStatus struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyStatus) Reset()         { *m = ReplyStatus{} }
func (m *ReplyStatus) String() string { return proto.CompactTextString(m) }
func (*ReplyStatus) ProtoMessage()    {}
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ac4d88d62e40be, []int{3}
}

func (m *ReplyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatus.Unmarshal(m, b)
}
func (m *ReplyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatus.Marshal(b, m, deterministic)
}
func (m *ReplyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatus.Merge(m, src)
}
func (m *ReplyStatus) XXX_Size() int {
	return xxx_messageInfo_ReplyStatus.Size(m)
}
func (m *ReplyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatus proto.InternalMessageInfo

func (m *ReplyStatus) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReplyStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ReplyInfo struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyInfo) Reset()         { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()    {}
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ac4d88d62e40be, []int{4}
}

func (m *ReplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInfo.Unmarshal(m, b)
}
func (m *ReplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInfo.Merge(m, src)
}
func (m *ReplyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyInfo.Size(m)
}
func (m *ReplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInfo proto.InternalMessageInfo

func (m *ReplyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type RequestIntFlag struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 int32    `protobuf:"varint,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestIntFlag) Reset()         { *m = RequestIntFlag{} }
func (m *RequestIntFlag) String() string { return proto.CompactTextString(m) }
func (*RequestIntFlag) ProtoMessage()    {}
func (*RequestIntFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ac4d88d62e40be, []int{5}
}

func (m *RequestIntFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestIntFlag.Unmarshal(m, b)
}
func (m *RequestIntFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestIntFlag.Marshal(b, m, deterministic)
}
func (m *RequestIntFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestIntFlag.Merge(m, src)
}
func (m *RequestIntFlag) XXX_Size() int {
	return xxx_messageInfo_RequestIntFlag.Size(m)
}
func (m *RequestIntFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestIntFlag.DiscardUnknown(m)
}

var xxx_messageInfo_RequestIntFlag proto.InternalMessageInfo

func (m *RequestIntFlag) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestIntFlag) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *RequestIntFlag) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestList struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []string `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestList) Reset()         { *m = RequestList{} }
func (m *RequestList) String() string { return proto.CompactTextString(m) }
func (*RequestList) ProtoMessage()    {}
func (*RequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ac4d88d62e40be, []int{6}
}

func (m *RequestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestList.Unmarshal(m, b)
}
func (m *RequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestList.Marshal(b, m, deterministic)
}
func (m *RequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestList.Merge(m, src)
}
func (m *RequestList) XXX_Size() int {
	return xxx_messageInfo_RequestList.Size(m)
}
func (m *RequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestList proto.InternalMessageInfo

func (m *RequestList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestList) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyList struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	List                 []string     `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyList) Reset()         { *m = ReplyList{} }
func (m *ReplyList) String() string { return proto.CompactTextString(m) }
func (*ReplyList) ProtoMessage()    {}
func (*ReplyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ac4d88d62e40be, []int{7}
}

func (m *ReplyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyList.Unmarshal(m, b)
}
func (m *ReplyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyList.Marshal(b, m, deterministic)
}
func (m *ReplyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyList.Merge(m, src)
}
func (m *ReplyList) XXX_Size() int {
	return xxx_messageInfo_ReplyList.Size(m)
}
func (m *ReplyList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyList proto.InternalMessageInfo

func (m *ReplyList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyStatistic struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Key                  string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Owner                string       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Count                uint32       `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyStatistic) Reset()         { *m = ReplyStatistic{} }
func (m *ReplyStatistic) String() string { return proto.CompactTextString(m) }
func (*ReplyStatistic) ProtoMessage()    {}
func (*ReplyStatistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_39ac4d88d62e40be, []int{8}
}

func (m *ReplyStatistic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatistic.Unmarshal(m, b)
}
func (m *ReplyStatistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatistic.Marshal(b, m, deterministic)
}
func (m *ReplyStatistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatistic.Merge(m, src)
}
func (m *ReplyStatistic) XXX_Size() int {
	return xxx_messageInfo_ReplyStatistic.Size(m)
}
func (m *ReplyStatistic) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatistic.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatistic proto.InternalMessageInfo

func (m *ReplyStatistic) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyStatistic) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReplyStatistic) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReplyStatistic) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.mqtt.RequestInfo")
	proto.RegisterType((*RequestFilter)(nil), "omo.msp.mqtt.RequestFilter")
	proto.RegisterType((*RequestUpdate)(nil), "omo.msp.mqtt.RequestUpdate")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.mqtt.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.mqtt.ReplyInfo")
	proto.RegisterType((*RequestIntFlag)(nil), "omo.msp.mqtt.RequestIntFlag")
	proto.RegisterType((*RequestList)(nil), "omo.msp.mqtt.RequestList")
	proto.RegisterType((*ReplyList)(nil), "omo.msp.mqtt.ReplyList")
	proto.RegisterType((*ReplyStatistic)(nil), "omo.msp.mqtt.ReplyStatistic")
}

func init() {
	proto.RegisterFile("proto/mqtt/common.proto", fileDescriptor_39ac4d88d62e40be)
}

var fileDescriptor_39ac4d88d62e40be = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x8e, 0x9b, 0x30,
	0x10, 0x15, 0x10, 0x68, 0xe3, 0x94, 0xaa, 0x42, 0x55, 0x4a, 0x7a, 0x8a, 0x38, 0xe5, 0x44, 0xd4,
	0xf6, 0xd0, 0x43, 0x6f, 0x3d, 0x44, 0xaa, 0x54, 0xa9, 0x95, 0x57, 0x7b, 0xd9, 0x4b, 0x44, 0xc0,
	0x89, 0x50, 0x30, 0x26, 0xb6, 0xd9, 0x28, 0xc7, 0xfd, 0x8a, 0xfd, 0x8e, 0xfd, 0xc3, 0xd5, 0x18,
	0x07, 0x9c, 0x6c, 0xd8, 0xcd, 0x6d, 0xde, 0xd8, 0x33, 0xf3, 0xe6, 0xf9, 0x01, 0xfa, 0x52, 0x71,
	0x26, 0xd9, 0x9c, 0xee, 0xa4, 0x9c, 0xa7, 0x8c, 0x52, 0x56, 0xc6, 0x2a, 0x13, 0x7c, 0x60, 0x94,
	0xc5, 0x54, 0x54, 0x31, 0x1c, 0x45, 0x7b, 0x34, 0xc2, 0x64, 0x57, 0x13, 0x21, 0xff, 0x94, 0x6b,
	0x16, 0x7c, 0x42, 0x4e, 0x9d, 0x67, 0xa1, 0x35, 0xb5, 0x66, 0x43, 0x0c, 0x61, 0x10, 0xa0, 0xc1,
	0xba, 0x48, 0x36, 0xa1, 0xad, 0x52, 0x2a, 0x0e, 0xbe, 0xa2, 0xf7, 0xac, 0x22, 0x3c, 0x91, 0x8c,
	0x87, 0x8e, 0xca, 0xb7, 0x18, 0xee, 0x97, 0x09, 0x25, 0xe1, 0xa0, 0xb9, 0x0f, 0x31, 0xe4, 0x6a,
	0x41, 0x78, 0xe8, 0x36, 0x39, 0x88, 0xa3, 0x27, 0x0b, 0xf9, 0x7a, 0xf2, 0x22, 0x2f, 0x24, 0xe1,
	0xc1, 0x67, 0xe4, 0xb2, 0x7d, 0x49, 0xb8, 0x9e, 0xde, 0x00, 0xc8, 0xae, 0x73, 0x52, 0x64, 0x9a,
	0x40, 0x03, 0x20, 0x7b, 0x9f, 0x14, 0x35, 0xd1, 0xe3, 0x1b, 0x00, 0x73, 0xaa, 0x64, 0xd3, 0xcc,
	0xf6, 0xb1, 0x8a, 0x83, 0x31, 0xf2, 0xca, 0x9a, 0xae, 0xf4, 0x74, 0x1f, 0x6b, 0x74, 0xb2, 0x83,
	0x77, 0xb6, 0xc3, 0x18, 0x79, 0xaa, 0xa1, 0x08, 0xdf, 0x4d, 0x9d, 0xd9, 0x10, 0x6b, 0x14, 0x3d,
	0x76, 0x9c, 0x6f, 0xab, 0x2c, 0x91, 0xa4, 0x87, 0xb3, 0x56, 0xd1, 0xee, 0x54, 0x6c, 0xb7, 0x70,
	0x2e, 0x6e, 0x31, 0x30, 0xb7, 0x30, 0x99, 0xb9, 0xbd, 0xcc, 0xbc, 0x13, 0x66, 0x3f, 0xe1, 0x19,
	0xab, 0xe2, 0x70, 0x23, 0x13, 0x59, 0x0b, 0x10, 0x22, 0x65, 0x19, 0x51, 0xac, 0x7c, 0xac, 0x62,
	0x18, 0x46, 0x38, 0x67, 0xfc, 0x28, 0xa4, 0x02, 0xd1, 0x7f, 0x34, 0x54, 0x85, 0x3d, 0xaf, 0xff,
	0x0d, 0x79, 0x42, 0xb5, 0x54, 0x55, 0xa3, 0xef, 0x93, 0xd8, 0x74, 0x4f, 0x6c, 0xcc, 0xc4, 0xfa,
	0x62, 0x84, 0xd1, 0xc7, 0xd6, 0x51, 0x72, 0x01, 0x76, 0x79, 0xdd, 0x54, 0xee, 0xdb, 0xa6, 0x8a,
	0xfe, 0xb5, 0x2e, 0xfd, 0x9b, 0x0b, 0x79, 0xa1, 0xa1, 0x59, 0x6c, 0xbf, 0x74, 0x64, 0x91, 0x0b,
	0x19, 0x3a, 0x4a, 0x31, 0x15, 0x47, 0x99, 0x5e, 0xbb, 0xa7, 0xdd, 0xb1, 0xc4, 0xee, 0x4a, 0x0c,
	0x29, 0x9c, 0x6b, 0xa5, 0x78, 0xb0, 0x40, 0x0b, 0x9d, 0xcf, 0x85, 0xcc, 0x53, 0xa3, 0x8b, 0x75,
	0x65, 0x17, 0xa0, 0xb7, 0x25, 0x87, 0xa3, 0x9b, 0xb6, 0xe4, 0xd0, 0xb9, 0xce, 0x39, 0xfb, 0x52,
	0x52, 0x56, 0x97, 0x52, 0xdb, 0xbf, 0x01, 0xbf, 0x27, 0x77, 0xc6, 0x9f, 0xe0, 0x17, 0xa3, 0x6c,
	0x49, 0x45, 0xb5, 0x04, 0xb0, 0xf2, 0xd4, 0xc1, 0x8f, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d,
	0x3d, 0xf8, 0x24, 0x2b, 0x04, 0x00, 0x00,
}
