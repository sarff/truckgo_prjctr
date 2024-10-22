// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: api/user.proto

package user

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

type NewDriverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Login     string `protobuf:"bytes,2,opt,name=Login,proto3" json:"Login,omitempty"`
	FullName  string `protobuf:"bytes,3,opt,name=FullName,proto3" json:"FullName,omitempty"`
	Phone     string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	License   string `protobuf:"bytes,5,opt,name=License,proto3" json:"License,omitempty"`
	CarModel  string `protobuf:"bytes,6,opt,name=CarModel,proto3" json:"CarModel,omitempty"`
	CarNumber string `protobuf:"bytes,7,opt,name=CarNumber,proto3" json:"CarNumber,omitempty"`
}

func (x *NewDriverRequest) Reset() {
	*x = NewDriverRequest{}
	mi := &file_api_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewDriverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewDriverRequest) ProtoMessage() {}

func (x *NewDriverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewDriverRequest.ProtoReflect.Descriptor instead.
func (*NewDriverRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{0}
}

func (x *NewDriverRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewDriverRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *NewDriverRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *NewDriverRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *NewDriverRequest) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *NewDriverRequest) GetCarModel() string {
	if x != nil {
		return x.CarModel
	}
	return ""
}

func (x *NewDriverRequest) GetCarNumber() string {
	if x != nil {
		return x.CarNumber
	}
	return ""
}

type NewDriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NewDriverResponse) Reset() {
	*x = NewDriverResponse{}
	mi := &file_api_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewDriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewDriverResponse) ProtoMessage() {}

func (x *NewDriverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewDriverResponse.ProtoReflect.Descriptor instead.
func (*NewDriverResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{1}
}

func (x *NewDriverResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NewCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Login    string `protobuf:"bytes,2,opt,name=Login,proto3" json:"Login,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=FullName,proto3" json:"FullName,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *NewCustomerRequest) Reset() {
	*x = NewCustomerRequest{}
	mi := &file_api_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCustomerRequest) ProtoMessage() {}

func (x *NewCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCustomerRequest.ProtoReflect.Descriptor instead.
func (*NewCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{2}
}

func (x *NewCustomerRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewCustomerRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *NewCustomerRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *NewCustomerRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type NewCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NewCustomerResponse) Reset() {
	*x = NewCustomerResponse{}
	mi := &file_api_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCustomerResponse) ProtoMessage() {}

func (x *NewCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCustomerResponse.ProtoReflect.Descriptor instead.
func (*NewCustomerResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{3}
}

func (x *NewCustomerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListDriverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ListDriverRequest) Reset() {
	*x = ListDriverRequest{}
	mi := &file_api_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDriverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDriverRequest) ProtoMessage() {}

func (x *ListDriverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDriverRequest.ProtoReflect.Descriptor instead.
func (*ListDriverRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{4}
}

func (x *ListDriverRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	mi := &file_api_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{5}
}

func (x *Driver) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Driver) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Driver) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type ListDriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drivers []*Driver `protobuf:"bytes,1,rep,name=drivers,proto3" json:"drivers,omitempty"`
}

func (x *ListDriverResponse) Reset() {
	*x = ListDriverResponse{}
	mi := &file_api_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDriverResponse) ProtoMessage() {}

func (x *ListDriverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDriverResponse.ProtoReflect.Descriptor instead.
func (*ListDriverResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{6}
}

func (x *ListDriverResponse) GetDrivers() []*Driver {
	if x != nil {
		return x.Drivers
	}
	return nil
}

type ListCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ListCustomerRequest) Reset() {
	*x = ListCustomerRequest{}
	mi := &file_api_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomerRequest) ProtoMessage() {}

func (x *ListCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomerRequest.ProtoReflect.Descriptor instead.
func (*ListCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{7}
}

func (x *ListCustomerRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	mi := &file_api_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{8}
}

func (x *Customer) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Customer) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Customer) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type ListCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customers []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
}

func (x *ListCustomerResponse) Reset() {
	*x = ListCustomerResponse{}
	mi := &file_api_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomerResponse) ProtoMessage() {}

func (x *ListCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomerResponse.ProtoReflect.Descriptor instead.
func (*ListCustomerResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{9}
}

func (x *ListCustomerResponse) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FullName  string  `protobuf:"bytes,2,opt,name=FullName,proto3" json:"FullName,omitempty"`
	Phone     string  `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Status    bool    `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
	Rating    float32 `protobuf:"fixed32,5,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Latitude  float64 `protobuf:"fixed64,6,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,7,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_api_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateUserRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateUserRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *UpdateUserRequest) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *UpdateUserRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UpdateUserRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	mi := &file_api_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *TypeRequest) Reset() {
	*x = TypeRequest{}
	mi := &file_api_user_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeRequest) ProtoMessage() {}

func (x *TypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeRequest.ProtoReflect.Descriptor instead.
func (*TypeRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{12}
}

func (x *TypeRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type TypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TypeResponse) Reset() {
	*x = TypeResponse{}
	mi := &file_api_user_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeResponse) ProtoMessage() {}

func (x *TypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeResponse.ProtoReflect.Descriptor instead.
func (*TypeResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{13}
}

func (x *TypeResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TypeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=Login,proto3" json:"Login,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_api_user_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{14}
}

func (x *UserRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Login      string  `protobuf:"bytes,2,opt,name=Login,proto3" json:"Login,omitempty"`
	FullName   string  `protobuf:"bytes,3,opt,name=FullName,proto3" json:"FullName,omitempty"`
	TypeUserID uint32  `protobuf:"varint,4,opt,name=TypeUserID,proto3" json:"TypeUserID,omitempty"`
	Status     bool    `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Phone      string  `protobuf:"bytes,6,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Rating     float32 `protobuf:"fixed32,7,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Latitude   float64 `protobuf:"fixed64,8,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude  float64 `protobuf:"fixed64,9,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_api_user_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{15}
}

func (x *UserResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserResponse) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserResponse) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserResponse) GetTypeUserID() uint32 {
	if x != nil {
		return x.TypeUserID
	}
	return 0
}

func (x *UserResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *UserResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserResponse) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *UserResponse) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UserResponse) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_api_user_proto protoreflect.FileDescriptor

var file_api_user_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x10, 0x4e, 0x65,
	0x77, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x11, 0x4e, 0x65,
	0x77, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6c, 0x0a, 0x12, 0x4e, 0x65, 0x77,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x52, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x3f, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x52, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x22, 0x2f, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x54, 0x0a,
	0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x4c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x22, 0x47, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x22, 0xbf, 0x01, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x2e,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26,
	0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0xf0, 0x01, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x54, 0x79, 0x70, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x32, 0xe8, 0x03, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x09,
	0x4e, 0x65, 0x77, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x74, 0x72, 0x75, 0x63,
	0x6b, 0x67, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x4e,
	0x65, 0x77, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x1b, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74,
	0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x75, 0x63,
	0x6b, 0x67, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x75,
	0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b,
	0x67, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x74, 0x72, 0x75,
	0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x61, 0x70, 0x69, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_proto_rawDescOnce sync.Once
	file_api_user_proto_rawDescData = file_api_user_proto_rawDesc
)

func file_api_user_proto_rawDescGZIP() []byte {
	file_api_user_proto_rawDescOnce.Do(func() {
		file_api_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_proto_rawDescData)
	})
	return file_api_user_proto_rawDescData
}

var file_api_user_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_user_proto_goTypes = []any{
	(*NewDriverRequest)(nil),     // 0: truckgo.NewDriverRequest
	(*NewDriverResponse)(nil),    // 1: truckgo.NewDriverResponse
	(*NewCustomerRequest)(nil),   // 2: truckgo.NewCustomerRequest
	(*NewCustomerResponse)(nil),  // 3: truckgo.NewCustomerResponse
	(*ListDriverRequest)(nil),    // 4: truckgo.ListDriverRequest
	(*Driver)(nil),               // 5: truckgo.Driver
	(*ListDriverResponse)(nil),   // 6: truckgo.ListDriverResponse
	(*ListCustomerRequest)(nil),  // 7: truckgo.ListCustomerRequest
	(*Customer)(nil),             // 8: truckgo.Customer
	(*ListCustomerResponse)(nil), // 9: truckgo.ListCustomerResponse
	(*UpdateUserRequest)(nil),    // 10: truckgo.UpdateUserRequest
	(*UpdateUserResponse)(nil),   // 11: truckgo.UpdateUserResponse
	(*TypeRequest)(nil),          // 12: truckgo.TypeRequest
	(*TypeResponse)(nil),         // 13: truckgo.TypeResponse
	(*UserRequest)(nil),          // 14: truckgo.UserRequest
	(*UserResponse)(nil),         // 15: truckgo.UserResponse
}
var file_api_user_proto_depIdxs = []int32{
	5,  // 0: truckgo.ListDriverResponse.drivers:type_name -> truckgo.Driver
	8,  // 1: truckgo.ListCustomerResponse.customers:type_name -> truckgo.Customer
	0,  // 2: truckgo.UserService.NewDriver:input_type -> truckgo.NewDriverRequest
	2,  // 3: truckgo.UserService.NewCustomer:input_type -> truckgo.NewCustomerRequest
	4,  // 4: truckgo.UserService.ListDrivers:input_type -> truckgo.ListDriverRequest
	7,  // 5: truckgo.UserService.ListCustomers:input_type -> truckgo.ListCustomerRequest
	10, // 6: truckgo.UserService.UpdateUser:input_type -> truckgo.UpdateUserRequest
	12, // 7: truckgo.UserService.GetType:input_type -> truckgo.TypeRequest
	14, // 8: truckgo.UserService.GetUser:input_type -> truckgo.UserRequest
	1,  // 9: truckgo.UserService.NewDriver:output_type -> truckgo.NewDriverResponse
	3,  // 10: truckgo.UserService.NewCustomer:output_type -> truckgo.NewCustomerResponse
	6,  // 11: truckgo.UserService.ListDrivers:output_type -> truckgo.ListDriverResponse
	9,  // 12: truckgo.UserService.ListCustomers:output_type -> truckgo.ListCustomerResponse
	11, // 13: truckgo.UserService.UpdateUser:output_type -> truckgo.UpdateUserResponse
	13, // 14: truckgo.UserService.GetType:output_type -> truckgo.TypeResponse
	15, // 15: truckgo.UserService.GetUser:output_type -> truckgo.UserResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_user_proto_init() }
func file_api_user_proto_init() {
	if File_api_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_proto_goTypes,
		DependencyIndexes: file_api_user_proto_depIdxs,
		MessageInfos:      file_api_user_proto_msgTypes,
	}.Build()
	File_api_user_proto = out.File
	file_api_user_proto_rawDesc = nil
	file_api_user_proto_goTypes = nil
	file_api_user_proto_depIdxs = nil
}
