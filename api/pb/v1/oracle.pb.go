// Copyright © 2024 Luther Systems, Ltd. All right reserved.

// API Models and Documentation.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: pb/v1/oracle.proto

package v1

import (
	v1 "buf.build/gen/go/luthersystems/protos/protocolbuffers/go/common/v1"
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

type ClaimState int32

const (
	ClaimState_CLAIM_STATE_UNSPECIFIED                ClaimState = 0
	ClaimState_CLAIM_STATE_NEW                        ClaimState = 1
	ClaimState_CLAIM_STATE_LOECLAIM_COLLECTED_DETAILS ClaimState = 2
	ClaimState_CLAIM_STATE_LOECLAIM_ID_VERIFIED       ClaimState = 3
	ClaimState_CLAIM_STATE_OOECLAIM_REVIEWED          ClaimState = 4
	ClaimState_CLAIM_STATE_OOECLAIM_VALIDATED         ClaimState = 5
	ClaimState_CLAIM_STATE_LOEFIN_INVOICE_ISSUED      ClaimState = 6
	ClaimState_CLAIM_STATE_OOEFIN_INVOICE_REVIEWED    ClaimState = 7
	ClaimState_CLAIM_STATE_OOEFIN_INVOICE_APPROVED    ClaimState = 8
	ClaimState_CLAIM_STATE_OOEPAY_PAYMENT_TRIGGERED   ClaimState = 9
)

// Enum value maps for ClaimState.
var (
	ClaimState_name = map[int32]string{
		0: "CLAIM_STATE_UNSPECIFIED",
		1: "CLAIM_STATE_NEW",
		2: "CLAIM_STATE_LOECLAIM_COLLECTED_DETAILS",
		3: "CLAIM_STATE_LOECLAIM_ID_VERIFIED",
		4: "CLAIM_STATE_OOECLAIM_REVIEWED",
		5: "CLAIM_STATE_OOECLAIM_VALIDATED",
		6: "CLAIM_STATE_LOEFIN_INVOICE_ISSUED",
		7: "CLAIM_STATE_OOEFIN_INVOICE_REVIEWED",
		8: "CLAIM_STATE_OOEFIN_INVOICE_APPROVED",
		9: "CLAIM_STATE_OOEPAY_PAYMENT_TRIGGERED",
	}
	ClaimState_value = map[string]int32{
		"CLAIM_STATE_UNSPECIFIED":                0,
		"CLAIM_STATE_NEW":                        1,
		"CLAIM_STATE_LOECLAIM_COLLECTED_DETAILS": 2,
		"CLAIM_STATE_LOECLAIM_ID_VERIFIED":       3,
		"CLAIM_STATE_OOECLAIM_REVIEWED":          4,
		"CLAIM_STATE_OOECLAIM_VALIDATED":         5,
		"CLAIM_STATE_LOEFIN_INVOICE_ISSUED":      6,
		"CLAIM_STATE_OOEFIN_INVOICE_REVIEWED":    7,
		"CLAIM_STATE_OOEFIN_INVOICE_APPROVED":    8,
		"CLAIM_STATE_OOEPAY_PAYMENT_TRIGGERED":   9,
	}
)

func (x ClaimState) Enum() *ClaimState {
	p := new(ClaimState)
	*p = x
	return p
}

func (x ClaimState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClaimState) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_v1_oracle_proto_enumTypes[0].Descriptor()
}

func (ClaimState) Type() protoreflect.EnumType {
	return &file_pb_v1_oracle_proto_enumTypes[0]
}

func (x ClaimState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClaimState.Descriptor instead.
func (ClaimState) EnumDescriptor() ([]byte, []int) {
	return file_pb_v1_oracle_proto_rawDescGZIP(), []int{0}
}

type CreateClaimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Claim *Claim `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (x *CreateClaimRequest) Reset() {
	*x = CreateClaimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_oracle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClaimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClaimRequest) ProtoMessage() {}

func (x *CreateClaimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_oracle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClaimRequest.ProtoReflect.Descriptor instead.
func (*CreateClaimRequest) Descriptor() ([]byte, []int) {
	return file_pb_v1_oracle_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClaimRequest) GetClaim() *Claim {
	if x != nil {
		return x.Claim
	}
	return nil
}

type CreateClaimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exception *v1.Exception `protobuf:"bytes,1,opt,name=exception,proto3" json:"exception,omitempty"`
	Claim     *Claim        `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (x *CreateClaimResponse) Reset() {
	*x = CreateClaimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_oracle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClaimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClaimResponse) ProtoMessage() {}

func (x *CreateClaimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_oracle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClaimResponse.ProtoReflect.Descriptor instead.
func (*CreateClaimResponse) Descriptor() ([]byte, []int) {
	return file_pb_v1_oracle_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClaimResponse) GetException() *v1.Exception {
	if x != nil {
		return x.Exception
	}
	return nil
}

func (x *CreateClaimResponse) GetClaim() *Claim {
	if x != nil {
		return x.Claim
	}
	return nil
}

type GetClaimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClaimId string `protobuf:"bytes,1,opt,name=claim_id,json=claimId,proto3" json:"claim_id,omitempty"`
}

func (x *GetClaimRequest) Reset() {
	*x = GetClaimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_oracle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClaimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClaimRequest) ProtoMessage() {}

func (x *GetClaimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_oracle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClaimRequest.ProtoReflect.Descriptor instead.
func (*GetClaimRequest) Descriptor() ([]byte, []int) {
	return file_pb_v1_oracle_proto_rawDescGZIP(), []int{2}
}

func (x *GetClaimRequest) GetClaimId() string {
	if x != nil {
		return x.ClaimId
	}
	return ""
}

type GetClaimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exception *v1.Exception `protobuf:"bytes,1,opt,name=exception,proto3" json:"exception,omitempty"`
	Claim     *Claim        `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (x *GetClaimResponse) Reset() {
	*x = GetClaimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_oracle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClaimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClaimResponse) ProtoMessage() {}

func (x *GetClaimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_oracle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClaimResponse.ProtoReflect.Descriptor instead.
func (*GetClaimResponse) Descriptor() ([]byte, []int) {
	return file_pb_v1_oracle_proto_rawDescGZIP(), []int{3}
}

func (x *GetClaimResponse) GetException() *v1.Exception {
	if x != nil {
		return x.Exception
	}
	return nil
}

func (x *GetClaimResponse) GetClaim() *Claim {
	if x != nil {
		return x.Claim
	}
	return nil
}

// Claim represents an insurance claim.
type Claim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID for the claim. Immutable. Set by backend.
	ClaimId string `protobuf:"bytes,1,opt,name=claim_id,json=claimId,proto3" json:"claim_id,omitempty"`
	// State of the claim.
	State ClaimState `protobuf:"varint,2,opt,name=state,proto3,enum=pb.v1.ClaimState" json:"state,omitempty"`
}

func (x *Claim) Reset() {
	*x = Claim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_oracle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claim) ProtoMessage() {}

func (x *Claim) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_oracle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claim.ProtoReflect.Descriptor instead.
func (*Claim) Descriptor() ([]byte, []int) {
	return file_pb_v1_oracle_proto_rawDescGZIP(), []int{4}
}

func (x *Claim) GetClaimId() string {
	if x != nil {
		return x.ClaimId
	}
	return ""
}

func (x *Claim) GetState() ClaimState {
	if x != nil {
		return x.State
	}
	return ClaimState_CLAIM_STATE_UNSPECIFIED
}

type TestTriggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *TestTriggerRequest) Reset() {
	*x = TestTriggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_oracle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestTriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestTriggerRequest) ProtoMessage() {}

func (x *TestTriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_oracle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestTriggerRequest.ProtoReflect.Descriptor instead.
func (*TestTriggerRequest) Descriptor() ([]byte, []int) {
	return file_pb_v1_oracle_proto_rawDescGZIP(), []int{5}
}

func (x *TestTriggerRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type TestTriggerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exception *v1.Exception `protobuf:"bytes,1,opt,name=exception,proto3" json:"exception,omitempty"`
}

func (x *TestTriggerResponse) Reset() {
	*x = TestTriggerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_oracle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestTriggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestTriggerResponse) ProtoMessage() {}

func (x *TestTriggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_oracle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestTriggerResponse.ProtoReflect.Descriptor instead.
func (*TestTriggerResponse) Descriptor() ([]byte, []int) {
	return file_pb_v1_oracle_proto_rawDescGZIP(), []int{6}
}

func (x *TestTriggerResponse) GetException() *v1.Exception {
	if x != nil {
		return x.Exception
	}
	return nil
}

var File_pb_v1_oracle_proto protoreflect.FileDescriptor

var file_pb_v1_oracle_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x22, 0x6d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x05, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x22,
	0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x49, 0x64, 0x22, 0x6a, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x22, 0x4b, 0x0a, 0x05, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x22, 0x49, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0xfa, 0x02, 0x0a, 0x0a,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4c,
	0x41, 0x49, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x41, 0x49, 0x4d,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x2a, 0x0a, 0x26,
	0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x45, 0x43,
	0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x44,
	0x45, 0x54, 0x41, 0x49, 0x4c, 0x53, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x4c, 0x41, 0x49,
	0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x45, 0x43, 0x4c, 0x41, 0x49, 0x4d,
	0x5f, 0x49, 0x44, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x03, 0x12, 0x21,
	0x0a, 0x1d, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4f,
	0x45, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x45, 0x44, 0x10,
	0x04, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x4f, 0x4f, 0x45, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x45, 0x46, 0x49, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x4f,
	0x49, 0x43, 0x45, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x44, 0x10, 0x06, 0x12, 0x27, 0x0a, 0x23,
	0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4f, 0x45, 0x46,
	0x49, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45,
	0x57, 0x45, 0x44, 0x10, 0x07, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4f, 0x45, 0x46, 0x49, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x4f,
	0x49, 0x43, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x08, 0x12, 0x28,
	0x0a, 0x24, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4f,
	0x45, 0x50, 0x41, 0x59, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x49,
	0x47, 0x47, 0x45, 0x52, 0x45, 0x44, 0x10, 0x09, 0x42, 0x79, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x75, 0x74, 0x68, 0x65, 0x72, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x50, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x05, 0x50, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x11, 0x50, 0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x50, 0x62, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_v1_oracle_proto_rawDescOnce sync.Once
	file_pb_v1_oracle_proto_rawDescData = file_pb_v1_oracle_proto_rawDesc
)

func file_pb_v1_oracle_proto_rawDescGZIP() []byte {
	file_pb_v1_oracle_proto_rawDescOnce.Do(func() {
		file_pb_v1_oracle_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_v1_oracle_proto_rawDescData)
	})
	return file_pb_v1_oracle_proto_rawDescData
}

var file_pb_v1_oracle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_v1_oracle_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_v1_oracle_proto_goTypes = []any{
	(ClaimState)(0),             // 0: pb.v1.ClaimState
	(*CreateClaimRequest)(nil),  // 1: pb.v1.CreateClaimRequest
	(*CreateClaimResponse)(nil), // 2: pb.v1.CreateClaimResponse
	(*GetClaimRequest)(nil),     // 3: pb.v1.GetClaimRequest
	(*GetClaimResponse)(nil),    // 4: pb.v1.GetClaimResponse
	(*Claim)(nil),               // 5: pb.v1.Claim
	(*TestTriggerRequest)(nil),  // 6: pb.v1.TestTriggerRequest
	(*TestTriggerResponse)(nil), // 7: pb.v1.TestTriggerResponse
	(*v1.Exception)(nil),        // 8: common.v1.Exception
}
var file_pb_v1_oracle_proto_depIdxs = []int32{
	5, // 0: pb.v1.CreateClaimRequest.claim:type_name -> pb.v1.Claim
	8, // 1: pb.v1.CreateClaimResponse.exception:type_name -> common.v1.Exception
	5, // 2: pb.v1.CreateClaimResponse.claim:type_name -> pb.v1.Claim
	8, // 3: pb.v1.GetClaimResponse.exception:type_name -> common.v1.Exception
	5, // 4: pb.v1.GetClaimResponse.claim:type_name -> pb.v1.Claim
	0, // 5: pb.v1.Claim.state:type_name -> pb.v1.ClaimState
	8, // 6: pb.v1.TestTriggerResponse.exception:type_name -> common.v1.Exception
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pb_v1_oracle_proto_init() }
func file_pb_v1_oracle_proto_init() {
	if File_pb_v1_oracle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_v1_oracle_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateClaimRequest); i {
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
		file_pb_v1_oracle_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateClaimResponse); i {
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
		file_pb_v1_oracle_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetClaimRequest); i {
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
		file_pb_v1_oracle_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetClaimResponse); i {
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
		file_pb_v1_oracle_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Claim); i {
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
		file_pb_v1_oracle_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TestTriggerRequest); i {
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
		file_pb_v1_oracle_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*TestTriggerResponse); i {
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
			RawDescriptor: file_pb_v1_oracle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_v1_oracle_proto_goTypes,
		DependencyIndexes: file_pb_v1_oracle_proto_depIdxs,
		EnumInfos:         file_pb_v1_oracle_proto_enumTypes,
		MessageInfos:      file_pb_v1_oracle_proto_msgTypes,
	}.Build()
	File_pb_v1_oracle_proto = out.File
	file_pb_v1_oracle_proto_rawDesc = nil
	file_pb_v1_oracle_proto_goTypes = nil
	file_pb_v1_oracle_proto_depIdxs = nil
}
