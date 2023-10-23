// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: proto/seckill/seckill.proto

package shop_product_srv

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

type DelSecKillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelSecKillRequest) Reset() {
	*x = DelSecKillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_seckill_seckill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelSecKillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelSecKillRequest) ProtoMessage() {}

func (x *DelSecKillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_seckill_seckill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelSecKillRequest.ProtoReflect.Descriptor instead.
func (*DelSecKillRequest) Descriptor() ([]byte, []int) {
	return file_proto_seckill_seckill_proto_rawDescGZIP(), []int{0}
}

func (x *DelSecKillRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SecKillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPage int32 `protobuf:"varint,1,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	PageSize    int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *SecKillsRequest) Reset() {
	*x = SecKillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_seckill_seckill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecKillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecKillsRequest) ProtoMessage() {}

func (x *SecKillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_seckill_seckill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecKillsRequest.ProtoReflect.Descriptor instead.
func (*SecKillsRequest) Descriptor() ([]byte, []int) {
	return file_proto_seckill_seckill_proto_rawDescGZIP(), []int{1}
}

func (x *SecKillsRequest) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *SecKillsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type SecKillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Seckills []*SecKill `protobuf:"bytes,3,rep,name=seckills,proto3" json:"seckills,omitempty"`
	Total    int32      `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	Current  int32      `protobuf:"varint,5,opt,name=current,proto3" json:"current,omitempty"`
	Pagesize int32      `protobuf:"varint,6,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
}

func (x *SecKillsResponse) Reset() {
	*x = SecKillsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_seckill_seckill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecKillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecKillsResponse) ProtoMessage() {}

func (x *SecKillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_seckill_seckill_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecKillsResponse.ProtoReflect.Descriptor instead.
func (*SecKillsResponse) Descriptor() ([]byte, []int) {
	return file_proto_seckill_seckill_proto_rawDescGZIP(), []int{2}
}

func (x *SecKillsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SecKillsResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SecKillsResponse) GetSeckills() []*SecKill {
	if x != nil {
		return x.Seckills
	}
	return nil
}

func (x *SecKillsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SecKillsResponse) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *SecKillsResponse) GetPagesize() int32 {
	if x != nil {
		return x.Pagesize
	}
	return 0
}

type SecKill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price      float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Num        int32   `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
	Pid        int32   `protobuf:"varint,5,opt,name=pid,proto3" json:"pid,omitempty"`
	StartTime  string  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    string  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	CreateTime string  `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Pname      string  `protobuf:"bytes,9,opt,name=pname,proto3" json:"pname,omitempty"`
}

func (x *SecKill) Reset() {
	*x = SecKill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_seckill_seckill_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecKill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecKill) ProtoMessage() {}

func (x *SecKill) ProtoReflect() protoreflect.Message {
	mi := &file_proto_seckill_seckill_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecKill.ProtoReflect.Descriptor instead.
func (*SecKill) Descriptor() ([]byte, []int) {
	return file_proto_seckill_seckill_proto_rawDescGZIP(), []int{3}
}

func (x *SecKill) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SecKill) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecKill) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SecKill) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *SecKill) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SecKill) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *SecKill) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *SecKill) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *SecKill) GetPname() string {
	if x != nil {
		return x.Pname
	}
	return ""
}

type AddSecKillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price      float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Num        int32   `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
	Pid        int32   `protobuf:"varint,5,opt,name=pid,proto3" json:"pid,omitempty"`
	StartTime  string  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    string  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	CreateTime string  `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Pname      string  `protobuf:"bytes,9,opt,name=pname,proto3" json:"pname,omitempty"`
}

func (x *AddSecKillRequest) Reset() {
	*x = AddSecKillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_seckill_seckill_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSecKillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSecKillRequest) ProtoMessage() {}

func (x *AddSecKillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_seckill_seckill_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSecKillRequest.ProtoReflect.Descriptor instead.
func (*AddSecKillRequest) Descriptor() ([]byte, []int) {
	return file_proto_seckill_seckill_proto_rawDescGZIP(), []int{4}
}

func (x *AddSecKillRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddSecKillRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddSecKillRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AddSecKillRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *AddSecKillRequest) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *AddSecKillRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *AddSecKillRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *AddSecKillRequest) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *AddSecKillRequest) GetPname() string {
	if x != nil {
		return x.Pname
	}
	return ""
}

type AddSecKillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AddSecKillResponse) Reset() {
	*x = AddSecKillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_seckill_seckill_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSecKillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSecKillResponse) ProtoMessage() {}

func (x *AddSecKillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_seckill_seckill_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSecKillResponse.ProtoReflect.Descriptor instead.
func (*AddSecKillResponse) Descriptor() ([]byte, []int) {
	return file_proto_seckill_seckill_proto_rawDescGZIP(), []int{5}
}

func (x *AddSecKillResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddSecKillResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SecKillEditData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price     float32      `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Num       int32        `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
	Pid       int32        `protobuf:"varint,5,opt,name=pid,proto3" json:"pid,omitempty"`
	Pname     string       `protobuf:"bytes,6,opt,name=pname,proto3" json:"pname,omitempty"`
	Prono     []*Productno `protobuf:"bytes,7,rep,name=prono,proto3" json:"prono,omitempty"`
	StartTime string       `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   string       `protobuf:"bytes,9,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Code      int32        `protobuf:"varint,10,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string       `protobuf:"bytes,11,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SecKillEditData) Reset() {
	*x = SecKillEditData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_seckill_seckill_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecKillEditData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecKillEditData) ProtoMessage() {}

func (x *SecKillEditData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_seckill_seckill_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecKillEditData.ProtoReflect.Descriptor instead.
func (*SecKillEditData) Descriptor() ([]byte, []int) {
	return file_proto_seckill_seckill_proto_rawDescGZIP(), []int{6}
}

func (x *SecKillEditData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SecKillEditData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecKillEditData) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SecKillEditData) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *SecKillEditData) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SecKillEditData) GetPname() string {
	if x != nil {
		return x.Pname
	}
	return ""
}

func (x *SecKillEditData) GetProno() []*Productno {
	if x != nil {
		return x.Prono
	}
	return nil
}

func (x *SecKillEditData) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *SecKillEditData) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *SecKillEditData) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SecKillEditData) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Productno struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Productno) Reset() {
	*x = Productno{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_seckill_seckill_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Productno) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Productno) ProtoMessage() {}

func (x *Productno) ProtoReflect() protoreflect.Message {
	mi := &file_proto_seckill_seckill_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Productno.ProtoReflect.Descriptor instead.
func (*Productno) Descriptor() ([]byte, []int) {
	return file_proto_seckill_seckill_proto_rawDescGZIP(), []int{7}
}

func (x *Productno) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Productno) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SecKillEditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price     float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Num       int32   `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
	Pid       int32   `protobuf:"varint,5,opt,name=pid,proto3" json:"pid,omitempty"`
	StartTime string  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   string  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *SecKillEditRequest) Reset() {
	*x = SecKillEditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_seckill_seckill_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecKillEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecKillEditRequest) ProtoMessage() {}

func (x *SecKillEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_seckill_seckill_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecKillEditRequest.ProtoReflect.Descriptor instead.
func (*SecKillEditRequest) Descriptor() ([]byte, []int) {
	return file_proto_seckill_seckill_proto_rawDescGZIP(), []int{8}
}

func (x *SecKillEditRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SecKillEditRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecKillEditRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SecKillEditRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *SecKillEditRequest) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SecKillEditRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *SecKillEditRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

var File_proto_seckill_seckill_proto protoreflect.FileDescriptor

var file_proto_seckill_seckill_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x2f,
	0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73,
	0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x22,
	0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x63, 0x4b,
	0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52,
	0x08, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xe2, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x53, 0x65, 0x63, 0x4b,
	0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x98, 0x02, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x45, 0x64, 0x69,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x6e, 0x6f, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x6e, 0x6f, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2f, 0x0a, 0x09,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6e, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xac, 0x01,
	0x0a, 0x12, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xce, 0x03, 0x0a,
	0x08, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x56, 0x0a, 0x0b, 0x53, 0x65, 0x63,
	0x4b, 0x69, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x65, 0x63, 0x4b,
	0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53,
	0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x41, 0x64, 0x64, 0x12,
	0x23, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x63, 0x4b, 0x69,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a,
	0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x65,
	0x6c, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x4b, 0x69,
	0x6c, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x65, 0x6c, 0x53, 0x65, 0x63,
	0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53,
	0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x45, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00,
	0x12, 0x5b, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x45, 0x64, 0x69, 0x74, 0x12,
	0x24, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x72, 0x76, 0x2e, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x63, 0x4b,
	0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a,
	0x20, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c,
	0x3b, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x72,
	0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_seckill_seckill_proto_rawDescOnce sync.Once
	file_proto_seckill_seckill_proto_rawDescData = file_proto_seckill_seckill_proto_rawDesc
)

func file_proto_seckill_seckill_proto_rawDescGZIP() []byte {
	file_proto_seckill_seckill_proto_rawDescOnce.Do(func() {
		file_proto_seckill_seckill_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_seckill_seckill_proto_rawDescData)
	})
	return file_proto_seckill_seckill_proto_rawDescData
}

var file_proto_seckill_seckill_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_seckill_seckill_proto_goTypes = []interface{}{
	(*DelSecKillRequest)(nil),  // 0: shop_product_srv.DelSecKillRequest
	(*SecKillsRequest)(nil),    // 1: shop_product_srv.SecKillsRequest
	(*SecKillsResponse)(nil),   // 2: shop_product_srv.SecKillsResponse
	(*SecKill)(nil),            // 3: shop_product_srv.SecKill
	(*AddSecKillRequest)(nil),  // 4: shop_product_srv.AddSecKillRequest
	(*AddSecKillResponse)(nil), // 5: shop_product_srv.AddSecKillResponse
	(*SecKillEditData)(nil),    // 6: shop_product_srv.SecKillEditData
	(*Productno)(nil),          // 7: shop_product_srv.Productno
	(*SecKillEditRequest)(nil), // 8: shop_product_srv.SecKillEditRequest
}
var file_proto_seckill_seckill_proto_depIdxs = []int32{
	3, // 0: shop_product_srv.SecKillsResponse.seckills:type_name -> shop_product_srv.SecKill
	7, // 1: shop_product_srv.SecKillEditData.prono:type_name -> shop_product_srv.Productno
	1, // 2: shop_product_srv.SecKills.SecKillList:input_type -> shop_product_srv.SecKillsRequest
	4, // 3: shop_product_srv.SecKills.SecKillAdd:input_type -> shop_product_srv.AddSecKillRequest
	0, // 4: shop_product_srv.SecKills.SecKillDel:input_type -> shop_product_srv.DelSecKillRequest
	0, // 5: shop_product_srv.SecKills.SecKillById:input_type -> shop_product_srv.DelSecKillRequest
	8, // 6: shop_product_srv.SecKills.SecKillEdit:input_type -> shop_product_srv.SecKillEditRequest
	2, // 7: shop_product_srv.SecKills.SecKillList:output_type -> shop_product_srv.SecKillsResponse
	5, // 8: shop_product_srv.SecKills.SecKillAdd:output_type -> shop_product_srv.AddSecKillResponse
	5, // 9: shop_product_srv.SecKills.SecKillDel:output_type -> shop_product_srv.AddSecKillResponse
	6, // 10: shop_product_srv.SecKills.SecKillById:output_type -> shop_product_srv.SecKillEditData
	5, // 11: shop_product_srv.SecKills.SecKillEdit:output_type -> shop_product_srv.AddSecKillResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_seckill_seckill_proto_init() }
func file_proto_seckill_seckill_proto_init() {
	if File_proto_seckill_seckill_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_seckill_seckill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelSecKillRequest); i {
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
		file_proto_seckill_seckill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecKillsRequest); i {
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
		file_proto_seckill_seckill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecKillsResponse); i {
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
		file_proto_seckill_seckill_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecKill); i {
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
		file_proto_seckill_seckill_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSecKillRequest); i {
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
		file_proto_seckill_seckill_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSecKillResponse); i {
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
		file_proto_seckill_seckill_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecKillEditData); i {
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
		file_proto_seckill_seckill_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Productno); i {
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
		file_proto_seckill_seckill_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecKillEditRequest); i {
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
			RawDescriptor: file_proto_seckill_seckill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_seckill_seckill_proto_goTypes,
		DependencyIndexes: file_proto_seckill_seckill_proto_depIdxs,
		MessageInfos:      file_proto_seckill_seckill_proto_msgTypes,
	}.Build()
	File_proto_seckill_seckill_proto = out.File
	file_proto_seckill_seckill_proto_rawDesc = nil
	file_proto_seckill_seckill_proto_goTypes = nil
	file_proto_seckill_seckill_proto_depIdxs = nil
}