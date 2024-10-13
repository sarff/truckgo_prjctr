// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: shipping.proto

package grpcapi

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

type PriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *PriceRequest) Reset() {
	*x = PriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceRequest) ProtoMessage() {}

func (x *PriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceRequest.ProtoReflect.Descriptor instead.
func (*PriceRequest) Descriptor() ([]byte, []int) {
	return file_shipping_proto_rawDescGZIP(), []int{0}
}

func (x *PriceRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *PriceRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type PriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Price    float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Distance float64 `protobuf:"fixed64,3,opt,name=distance,proto3" json:"distance,omitempty"`
	Time     float64 `protobuf:"fixed64,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *PriceResponse) Reset() {
	*x = PriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceResponse) ProtoMessage() {}

func (x *PriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceResponse.ProtoReflect.Descriptor instead.
func (*PriceResponse) Descriptor() ([]byte, []int) {
	return file_shipping_proto_rawDescGZIP(), []int{1}
}

func (x *PriceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PriceResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PriceResponse) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *PriceResponse) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type RouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *RouteRequest) Reset() {
	*x = RouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteRequest) ProtoMessage() {}

func (x *RouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteRequest.ProtoReflect.Descriptor instead.
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return file_shipping_proto_rawDescGZIP(), []int{2}
}

func (x *RouteRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *RouteRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type RouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Steps    []*Step `protobuf:"bytes,2,rep,name=steps,proto3" json:"steps,omitempty"`
	Distance float64 `protobuf:"fixed64,3,opt,name=distance,proto3" json:"distance,omitempty"`
	Time     float64 `protobuf:"fixed64,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *RouteResponse) Reset() {
	*x = RouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteResponse) ProtoMessage() {}

func (x *RouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteResponse.ProtoReflect.Descriptor instead.
func (*RouteResponse) Descriptor() ([]byte, []int) {
	return file_shipping_proto_rawDescGZIP(), []int{3}
}

func (x *RouteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RouteResponse) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *RouteResponse) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *RouteResponse) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instruction string  `protobuf:"bytes,1,opt,name=instruction,proto3" json:"instruction,omitempty"`
	Distance    float64 `protobuf:"fixed64,2,opt,name=distance,proto3" json:"distance,omitempty"`
	Duration    float64 `protobuf:"fixed64,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_shipping_proto_rawDescGZIP(), []int{4}
}

func (x *Step) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

func (x *Step) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Step) GetDuration() float64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_shipping_proto_rawDescGZIP(), []int{5}
}

func (x *TestRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_shipping_proto_rawDescGZIP(), []int{6}
}

func (x *TestResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_shipping_proto protoreflect.FileDescriptor

var file_shipping_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x22, 0x48, 0x0a, 0x0c, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x6f, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7e,
	0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74, 0x65,
	0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b,
	0x67, 0x6f, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x60,
	0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x27, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xcc, 0x01, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x75, 0x63,
	0x6b, 0x67, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x75,
	0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x54, 0x65, 0x73,
	0x74, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x14, 0x2e, 0x74, 0x72, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x72,
	0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x65, 0x61, 0x72, 0x2f, 0x74, 0x72, 0x75, 0x63,
	0x6b, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_shipping_proto_rawDescOnce sync.Once
	file_shipping_proto_rawDescData = file_shipping_proto_rawDesc
)

func file_shipping_proto_rawDescGZIP() []byte {
	file_shipping_proto_rawDescOnce.Do(func() {
		file_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipping_proto_rawDescData)
	})
	return file_shipping_proto_rawDescData
}

var file_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_shipping_proto_goTypes = []any{
	(*PriceRequest)(nil),  // 0: truckgo.PriceRequest
	(*PriceResponse)(nil), // 1: truckgo.PriceResponse
	(*RouteRequest)(nil),  // 2: truckgo.RouteRequest
	(*RouteResponse)(nil), // 3: truckgo.RouteResponse
	(*Step)(nil),          // 4: truckgo.Step
	(*TestRequest)(nil),   // 5: truckgo.TestRequest
	(*TestResponse)(nil),  // 6: truckgo.TestResponse
}
var file_shipping_proto_depIdxs = []int32{
	4, // 0: truckgo.RouteResponse.steps:type_name -> truckgo.Step
	0, // 1: truckgo.ShippingService.CalculatePrice:input_type -> truckgo.PriceRequest
	2, // 2: truckgo.ShippingService.CalculateRoute:input_type -> truckgo.RouteRequest
	5, // 3: truckgo.ShippingService.TestFunc:input_type -> truckgo.TestRequest
	1, // 4: truckgo.ShippingService.CalculatePrice:output_type -> truckgo.PriceResponse
	3, // 5: truckgo.ShippingService.CalculateRoute:output_type -> truckgo.RouteResponse
	6, // 6: truckgo.ShippingService.TestFunc:output_type -> truckgo.TestResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shipping_proto_init() }
func file_shipping_proto_init() {
	if File_shipping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shipping_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PriceRequest); i {
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
		file_shipping_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PriceResponse); i {
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
		file_shipping_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RouteRequest); i {
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
		file_shipping_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RouteResponse); i {
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
		file_shipping_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Step); i {
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
		file_shipping_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TestRequest); i {
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
		file_shipping_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*TestResponse); i {
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
			RawDescriptor: file_shipping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipping_proto_goTypes,
		DependencyIndexes: file_shipping_proto_depIdxs,
		MessageInfos:      file_shipping_proto_msgTypes,
	}.Build()
	File_shipping_proto = out.File
	file_shipping_proto_rawDesc = nil
	file_shipping_proto_goTypes = nil
	file_shipping_proto_depIdxs = nil
}
