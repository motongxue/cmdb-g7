// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: apps/host/pb/host.proto

package host

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
	resource "github.com/motongxue/cmdb-g7/apps/resource"
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

type DescribeBy int32

const (
	DescribeBy_HOST_ID DescribeBy = 0
)

// Enum value maps for DescribeBy.
var (
	DescribeBy_name = map[int32]string{
		0: "HOST_ID",
	}
	DescribeBy_value = map[string]int32{
		"HOST_ID": 0,
	}
)

func (x DescribeBy) Enum() *DescribeBy {
	p := new(DescribeBy)
	*p = x
	return p
}

func (x DescribeBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DescribeBy) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_host_pb_host_proto_enumTypes[0].Descriptor()
}

func (DescribeBy) Type() protoreflect.EnumType {
	return &file_apps_host_pb_host_proto_enumTypes[0]
}

func (x DescribeBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DescribeBy.Descriptor instead.
func (DescribeBy) EnumDescriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{0}
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 固有基础属性
	// @gotags: json:"base"
	Base *resource.Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base"`
	// 非固有的基础属性
	// @gotags: json:"information"
	Information *resource.Information `protobuf:"bytes,2,opt,name=information,proto3" json:"information"`
	// 主机资源 固有的属性
	// @gotags: json:"describe"
	Describe *Describe `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetBase() *resource.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Host) GetInformation() *resource.Information {
	if x != nil {
		return x.Information
	}
	return nil
}

func (x *Host) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 核数
	// @gotags: json:"cpu"
	Cpu int64 `protobuf:"varint,1,opt,name=cpu,proto3" json:"cpu"`
	// 内存
	// @gotags: json:"memory"
	Memory int64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory"`
	// GPU数量
	// @gotags: json:"gpu_amount"
	GpuAmount int32 `protobuf:"varint,3,opt,name=gpu_amount,json=gpuAmount,proto3" json:"gpu_amount"`
	// GPU类型
	// @gotags: json:"gpu_spec"
	GpuSpec string `protobuf:"bytes,4,opt,name=gpu_spec,json=gpuSpec,proto3" json:"gpu_spec"`
	// 操作系统类型，分为Windows和Linux
	// @gotags: json:"os_type"
	OsType string `protobuf:"bytes,5,opt,name=os_type,json=osType,proto3" json:"os_type"`
	// 操作系统名称
	// @gotags: json:"os_name"
	OsName string `protobuf:"bytes,6,opt,name=os_name,json=osName,proto3" json:"os_name"`
	// 序列号
	// @gotags: json:"serial_number"
	SerialNumber string `protobuf:"bytes,7,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number"`
	// 镜像ID
	// @gotags: json:"image_id"
	ImageId string `protobuf:"bytes,8,opt,name=image_id,json=imageId,proto3" json:"image_id"`
	// 公网出带宽最大值，单位为 Mbps
	// @gotags: json:"internet_max_bandwidth_out"
	InternetMaxBandwidthOut int64 `protobuf:"varint,9,opt,name=internet_max_bandwidth_out,json=internetMaxBandwidthOut,proto3" json:"internet_max_bandwidth_out"`
	// 公网入带宽最大值，单位为 Mbps
	// @gotags: json:"internet_max_bandwidth_in"
	InternetMaxBandwidthIn int64 `protobuf:"varint,10,opt,name=internet_max_bandwidth_in,json=internetMaxBandwidthIn,proto3" json:"internet_max_bandwidth_in"`
	// 秘钥对名称
	// @gotags: json:"key_pair_name"
	KeyPairName []string `protobuf:"bytes,11,rep,name=key_pair_name,json=keyPairName,proto3" json:"key_pair_name"`
	// 安全组  采用逗号分隔
	// @gotags: json:"security_groups"
	SecurityGroups []string `protobuf:"bytes,12,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups"`
	// 额外的无法通用的一些属性, 比如只有腾讯云独有的一些属性
	// @gotags: json:"extra"
	Extra map[string]string `protobuf:"bytes,13,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Describe.ProtoReflect.Descriptor instead.
func (*Describe) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{1}
}

func (x *Describe) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Describe) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Describe) GetGpuAmount() int32 {
	if x != nil {
		return x.GpuAmount
	}
	return 0
}

func (x *Describe) GetGpuSpec() string {
	if x != nil {
		return x.GpuSpec
	}
	return ""
}

func (x *Describe) GetOsType() string {
	if x != nil {
		return x.OsType
	}
	return ""
}

func (x *Describe) GetOsName() string {
	if x != nil {
		return x.OsName
	}
	return ""
}

func (x *Describe) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Describe) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *Describe) GetInternetMaxBandwidthOut() int64 {
	if x != nil {
		return x.InternetMaxBandwidthOut
	}
	return 0
}

func (x *Describe) GetInternetMaxBandwidthIn() int64 {
	if x != nil {
		return x.InternetMaxBandwidthIn
	}
	return 0
}

func (x *Describe) GetKeyPairName() []string {
	if x != nil {
		return x.KeyPairName
	}
	return nil
}

func (x *Describe) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

func (x *Describe) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type QueryHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 关键字参数
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,2,opt,name=keywords,proto3" json:"keywords"`
}

func (x *QueryHostRequest) Reset() {
	*x = QueryHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryHostRequest) ProtoMessage() {}

func (x *QueryHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryHostRequest.ProtoReflect.Descriptor instead.
func (*QueryHostRequest) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{2}
}

func (x *QueryHostRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryHostRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type HostSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// @gotags: json:"items"
	Items []*Host `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *HostSet) Reset() {
	*x = HostSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostSet) ProtoMessage() {}

func (x *HostSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostSet.ProtoReflect.Descriptor instead.
func (*HostSet) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{3}
}

func (x *HostSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *HostSet) GetItems() []*Host {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required"`
	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,2,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// @gotags: json:"data" validate:"required"
	UpdateHostData *UpdateHostData `protobuf:"bytes,3,opt,name=update_host_data,json=updateHostData,proto3" json:"data" validate:"required"`
}

func (x *UpdateHostRequest) Reset() {
	*x = UpdateHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHostRequest) ProtoMessage() {}

func (x *UpdateHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHostRequest.ProtoReflect.Descriptor instead.
func (*UpdateHostRequest) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateHostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateHostRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdateHostRequest) GetUpdateHostData() *UpdateHostData {
	if x != nil {
		return x.UpdateHostData
	}
	return nil
}

type UpdateHostData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"information"
	Information *resource.Information `protobuf:"bytes,1,opt,name=information,proto3" json:"information"`
	// @gotags: json:"describe"
	Describe *Describe `protobuf:"bytes,2,opt,name=describe,proto3" json:"describe"`
}

func (x *UpdateHostData) Reset() {
	*x = UpdateHostData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHostData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHostData) ProtoMessage() {}

func (x *UpdateHostData) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHostData.ProtoReflect.Descriptor instead.
func (*UpdateHostData) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateHostData) GetInformation() *resource.Information {
	if x != nil {
		return x.Information
	}
	return nil
}

func (x *UpdateHostData) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type DescribeHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"describe_by"
	DescribeBy DescribeBy `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=course.cmdb.host.DescribeBy" json:"describe_by"`
	// @gotags: json:"value" validate:"required"
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value" validate:"required"`
}

func (x *DescribeHostRequest) Reset() {
	*x = DescribeHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeHostRequest) ProtoMessage() {}

func (x *DescribeHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeHostRequest.ProtoReflect.Descriptor instead.
func (*DescribeHostRequest) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeHostRequest) GetDescribeBy() DescribeBy {
	if x != nil {
		return x.DescribeBy
	}
	return DescribeBy_HOST_ID
}

func (x *DescribeHostRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ReleaseHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主机Id
	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required"`
}

func (x *ReleaseHostRequest) Reset() {
	*x = ReleaseHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseHostRequest) ProtoMessage() {}

func (x *ReleaseHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseHostRequest.ProtoReflect.Descriptor instead.
func (*ReleaseHostRequest) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{7}
}

func (x *ReleaseHostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apps_host_pb_host_proto protoreflect.FileDescriptor

var file_apps_host_pb_host_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x68,
	0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x1a, 0x1f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x36, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x08,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x9c, 0x04, 0x0a, 0x08, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x70, 0x75, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x70, 0x75, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x70, 0x75, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x70, 0x75, 0x53, 0x70, 0x65, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x4f, 0x75, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x5f, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x49, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x50,
	0x61, 0x69, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x3b, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a,
	0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x66, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22,
	0x4d, 0x0a, 0x07, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xb6,
	0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x0b, 0x69, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x36, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x68, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x08, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x6a, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42,
	0x79, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x48, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x19, 0x0a, 0x0a, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x4f, 0x53, 0x54, 0x5f,
	0x49, 0x44, 0x10, 0x00, 0x32, 0xf8, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3a, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x09,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x74, 0x12, 0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x68, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f,
	0x74, 0x6f, 0x6e, 0x67, 0x78, 0x75, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2d, 0x67, 0x37, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_host_pb_host_proto_rawDescOnce sync.Once
	file_apps_host_pb_host_proto_rawDescData = file_apps_host_pb_host_proto_rawDesc
)

func file_apps_host_pb_host_proto_rawDescGZIP() []byte {
	file_apps_host_pb_host_proto_rawDescOnce.Do(func() {
		file_apps_host_pb_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_host_pb_host_proto_rawDescData)
	})
	return file_apps_host_pb_host_proto_rawDescData
}

var file_apps_host_pb_host_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_host_pb_host_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_host_pb_host_proto_goTypes = []interface{}{
	(DescribeBy)(0),              // 0: course.cmdb.host.DescribeBy
	(*Host)(nil),                 // 1: course.cmdb.host.Host
	(*Describe)(nil),             // 2: course.cmdb.host.Describe
	(*QueryHostRequest)(nil),     // 3: course.cmdb.host.QueryHostRequest
	(*HostSet)(nil),              // 4: course.cmdb.host.HostSet
	(*UpdateHostRequest)(nil),    // 5: course.cmdb.host.UpdateHostRequest
	(*UpdateHostData)(nil),       // 6: course.cmdb.host.UpdateHostData
	(*DescribeHostRequest)(nil),  // 7: course.cmdb.host.DescribeHostRequest
	(*ReleaseHostRequest)(nil),   // 8: course.cmdb.host.ReleaseHostRequest
	nil,                          // 9: course.cmdb.host.Describe.ExtraEntry
	(*resource.Base)(nil),        // 10: course.cmdb.resource.Base
	(*resource.Information)(nil), // 11: course.cmdb.resource.Information
	(*request.PageRequest)(nil),  // 12: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),     // 13: infraboard.mcube.request.UpdateMode
}
var file_apps_host_pb_host_proto_depIdxs = []int32{
	10, // 0: course.cmdb.host.Host.base:type_name -> course.cmdb.resource.Base
	11, // 1: course.cmdb.host.Host.information:type_name -> course.cmdb.resource.Information
	2,  // 2: course.cmdb.host.Host.describe:type_name -> course.cmdb.host.Describe
	9,  // 3: course.cmdb.host.Describe.extra:type_name -> course.cmdb.host.Describe.ExtraEntry
	12, // 4: course.cmdb.host.QueryHostRequest.page:type_name -> infraboard.mcube.page.PageRequest
	1,  // 5: course.cmdb.host.HostSet.items:type_name -> course.cmdb.host.Host
	13, // 6: course.cmdb.host.UpdateHostRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	6,  // 7: course.cmdb.host.UpdateHostRequest.update_host_data:type_name -> course.cmdb.host.UpdateHostData
	11, // 8: course.cmdb.host.UpdateHostData.information:type_name -> course.cmdb.resource.Information
	2,  // 9: course.cmdb.host.UpdateHostData.describe:type_name -> course.cmdb.host.Describe
	0,  // 10: course.cmdb.host.DescribeHostRequest.describe_by:type_name -> course.cmdb.host.DescribeBy
	1,  // 11: course.cmdb.host.Service.SyncHost:input_type -> course.cmdb.host.Host
	3,  // 12: course.cmdb.host.Service.QueryHost:input_type -> course.cmdb.host.QueryHostRequest
	7,  // 13: course.cmdb.host.Service.DescribeHost:input_type -> course.cmdb.host.DescribeHostRequest
	5,  // 14: course.cmdb.host.Service.UpdateHost:input_type -> course.cmdb.host.UpdateHostRequest
	8,  // 15: course.cmdb.host.Service.ReleaseHost:input_type -> course.cmdb.host.ReleaseHostRequest
	1,  // 16: course.cmdb.host.Service.SyncHost:output_type -> course.cmdb.host.Host
	4,  // 17: course.cmdb.host.Service.QueryHost:output_type -> course.cmdb.host.HostSet
	1,  // 18: course.cmdb.host.Service.DescribeHost:output_type -> course.cmdb.host.Host
	1,  // 19: course.cmdb.host.Service.UpdateHost:output_type -> course.cmdb.host.Host
	1,  // 20: course.cmdb.host.Service.ReleaseHost:output_type -> course.cmdb.host.Host
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_apps_host_pb_host_proto_init() }
func file_apps_host_pb_host_proto_init() {
	if File_apps_host_pb_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_host_pb_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_apps_host_pb_host_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Describe); i {
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
		file_apps_host_pb_host_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryHostRequest); i {
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
		file_apps_host_pb_host_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostSet); i {
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
		file_apps_host_pb_host_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHostRequest); i {
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
		file_apps_host_pb_host_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHostData); i {
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
		file_apps_host_pb_host_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeHostRequest); i {
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
		file_apps_host_pb_host_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseHostRequest); i {
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
			RawDescriptor: file_apps_host_pb_host_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_host_pb_host_proto_goTypes,
		DependencyIndexes: file_apps_host_pb_host_proto_depIdxs,
		EnumInfos:         file_apps_host_pb_host_proto_enumTypes,
		MessageInfos:      file_apps_host_pb_host_proto_msgTypes,
	}.Build()
	File_apps_host_pb_host_proto = out.File
	file_apps_host_pb_host_proto_rawDesc = nil
	file_apps_host_pb_host_proto_goTypes = nil
	file_apps_host_pb_host_proto_depIdxs = nil
}
