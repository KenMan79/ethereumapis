// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: eth/v1/events_service.proto

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	gateway "github.com/grpc-ecosystem/grpc-gateway/v2/proto/gateway"
	github_com_prysmaticlabs_eth2_types "github.com/prysmaticlabs/eth2-types"
	_ "github.com/prysmaticlabs/ethereumapis/eth/ext"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StreamEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *StreamEventsRequest) Reset() {
	*x = StreamEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_events_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEventsRequest) ProtoMessage() {}

func (x *StreamEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_events_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEventsRequest.ProtoReflect.Descriptor instead.
func (*StreamEventsRequest) Descriptor() ([]byte, []int) {
	return file_eth_v1_events_service_proto_rawDescGZIP(), []int{0}
}

func (x *StreamEventsRequest) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type EventHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot                      github_com_prysmaticlabs_eth2_types.Slot `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	Block                     []byte                                   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty" ssz-size:"32"`
	State                     []byte                                   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty" ssz-size:"32"`
	EpochTransition           bool                                     `protobuf:"varint,4,opt,name=epoch_transition,json=epochTransition,proto3" json:"epoch_transition,omitempty"`
	PreviousDutyDependentRoot []byte                                   `protobuf:"bytes,5,opt,name=previous_duty_dependent_root,json=previousDutyDependentRoot,proto3" json:"previous_duty_dependent_root,omitempty" ssz-size:"32"`
	CurrentDutyDependentRoot  []byte                                   `protobuf:"bytes,6,opt,name=current_duty_dependent_root,json=currentDutyDependentRoot,proto3" json:"current_duty_dependent_root,omitempty" ssz-size:"32"`
}

func (x *EventHead) Reset() {
	*x = EventHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_events_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHead) ProtoMessage() {}

func (x *EventHead) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_events_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHead.ProtoReflect.Descriptor instead.
func (*EventHead) Descriptor() ([]byte, []int) {
	return file_eth_v1_events_service_proto_rawDescGZIP(), []int{1}
}

func (x *EventHead) GetSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *EventHead) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *EventHead) GetState() []byte {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *EventHead) GetEpochTransition() bool {
	if x != nil {
		return x.EpochTransition
	}
	return false
}

func (x *EventHead) GetPreviousDutyDependentRoot() []byte {
	if x != nil {
		return x.PreviousDutyDependentRoot
	}
	return nil
}

func (x *EventHead) GetCurrentDutyDependentRoot() []byte {
	if x != nil {
		return x.CurrentDutyDependentRoot
	}
	return nil
}

type EventBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot  github_com_prysmaticlabs_eth2_types.Slot `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	Block []byte                                   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty" ssz-size:"32"`
}

func (x *EventBlock) Reset() {
	*x = EventBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_events_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventBlock) ProtoMessage() {}

func (x *EventBlock) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_events_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventBlock.ProtoReflect.Descriptor instead.
func (*EventBlock) Descriptor() ([]byte, []int) {
	return file_eth_v1_events_service_proto_rawDescGZIP(), []int{2}
}

func (x *EventBlock) GetSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *EventBlock) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

type EventChainReorg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot         github_com_prysmaticlabs_eth2_types.Slot  `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	Depth        uint64                                    `protobuf:"varint,2,opt,name=depth,proto3" json:"depth,omitempty"`
	OldHeadBlock []byte                                    `protobuf:"bytes,3,opt,name=old_head_block,json=oldHeadBlock,proto3" json:"old_head_block,omitempty" ssz-size:"32"`
	NewHeadBlock []byte                                    `protobuf:"bytes,4,opt,name=new_head_block,json=newHeadBlock,proto3" json:"new_head_block,omitempty" ssz-size:"32"`
	OldHeadState []byte                                    `protobuf:"bytes,5,opt,name=old_head_state,json=oldHeadState,proto3" json:"old_head_state,omitempty" ssz-size:"32"`
	NewHeadState []byte                                    `protobuf:"bytes,6,opt,name=new_head_state,json=newHeadState,proto3" json:"new_head_state,omitempty" ssz-size:"32"`
	Epoch        github_com_prysmaticlabs_eth2_types.Epoch `protobuf:"varint,7,opt,name=epoch,proto3" json:"epoch,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Epoch"`
}

func (x *EventChainReorg) Reset() {
	*x = EventChainReorg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_events_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventChainReorg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventChainReorg) ProtoMessage() {}

func (x *EventChainReorg) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_events_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventChainReorg.ProtoReflect.Descriptor instead.
func (*EventChainReorg) Descriptor() ([]byte, []int) {
	return file_eth_v1_events_service_proto_rawDescGZIP(), []int{3}
}

func (x *EventChainReorg) GetSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *EventChainReorg) GetDepth() uint64 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *EventChainReorg) GetOldHeadBlock() []byte {
	if x != nil {
		return x.OldHeadBlock
	}
	return nil
}

func (x *EventChainReorg) GetNewHeadBlock() []byte {
	if x != nil {
		return x.NewHeadBlock
	}
	return nil
}

func (x *EventChainReorg) GetOldHeadState() []byte {
	if x != nil {
		return x.OldHeadState
	}
	return nil
}

func (x *EventChainReorg) GetNewHeadState() []byte {
	if x != nil {
		return x.NewHeadState
	}
	return nil
}

func (x *EventChainReorg) GetEpoch() github_com_prysmaticlabs_eth2_types.Epoch {
	if x != nil {
		return x.Epoch
	}
	return github_com_prysmaticlabs_eth2_types.Epoch(0)
}

type EventFinalizedCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block []byte                                    `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty" ssz-size:"32"`
	State []byte                                    `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty" ssz-size:"32"`
	Epoch github_com_prysmaticlabs_eth2_types.Epoch `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Epoch"`
}

func (x *EventFinalizedCheckpoint) Reset() {
	*x = EventFinalizedCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_events_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventFinalizedCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventFinalizedCheckpoint) ProtoMessage() {}

func (x *EventFinalizedCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_events_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventFinalizedCheckpoint.ProtoReflect.Descriptor instead.
func (*EventFinalizedCheckpoint) Descriptor() ([]byte, []int) {
	return file_eth_v1_events_service_proto_rawDescGZIP(), []int{4}
}

func (x *EventFinalizedCheckpoint) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *EventFinalizedCheckpoint) GetState() []byte {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *EventFinalizedCheckpoint) GetEpoch() github_com_prysmaticlabs_eth2_types.Epoch {
	if x != nil {
		return x.Epoch
	}
	return github_com_prysmaticlabs_eth2_types.Epoch(0)
}

var File_eth_v1_events_service_proto protoreflect.FileDescriptor

var file_eth_v1_events_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f,
	0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x22, 0xc4, 0x02, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x12, 0x40, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2c,
	0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68,
	0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c,
	0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x1c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x1c, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x64, 0x75, 0x74, 0x79, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x19, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x44, 0x75, 0x74, 0x79, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x45, 0x0a, 0x1b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x75,
	0x74, 0x79, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52,
	0x18, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x75, 0x74, 0x79, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x22, 0x6c, 0x0a, 0x0a, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x40, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53,
	0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32,
	0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xe6, 0x02, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x6f, 0x72, 0x67, 0x12, 0x40, 0x0a, 0x04, 0x73,
	0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x65,
	0x70, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x0e, 0x6f, 0x6c, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18,
	0x02, 0x33, 0x32, 0x52, 0x0c, 0x6f, 0x6c, 0x64, 0x48, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x2c, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33,
	0x32, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x48, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x2c, 0x0a, 0x0e, 0x6f, 0x6c, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52,
	0x0c, 0x6f, 0x6c, 0x64, 0x48, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a,
	0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0c, 0x6e,
	0x65, 0x77, 0x48, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x65,
	0x70, 0x6f, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2d, 0x82, 0xb5, 0x18, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x22, 0x9b, 0x01, 0x0a, 0x18, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5,
	0x18, 0x02, 0x33, 0x32, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02,
	0x33, 0x32, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2d, 0x82, 0xb5, 0x18, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x32, 0x6e,
	0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x64, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x65,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x30, 0x01, 0x42, 0x7b,
	0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x76, 0x31, 0xca, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eth_v1_events_service_proto_rawDescOnce sync.Once
	file_eth_v1_events_service_proto_rawDescData = file_eth_v1_events_service_proto_rawDesc
)

func file_eth_v1_events_service_proto_rawDescGZIP() []byte {
	file_eth_v1_events_service_proto_rawDescOnce.Do(func() {
		file_eth_v1_events_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v1_events_service_proto_rawDescData)
	})
	return file_eth_v1_events_service_proto_rawDescData
}

var file_eth_v1_events_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eth_v1_events_service_proto_goTypes = []interface{}{
	(*StreamEventsRequest)(nil),      // 0: ethereum.eth.v1.StreamEventsRequest
	(*EventHead)(nil),                // 1: ethereum.eth.v1.EventHead
	(*EventBlock)(nil),               // 2: ethereum.eth.v1.EventBlock
	(*EventChainReorg)(nil),          // 3: ethereum.eth.v1.EventChainReorg
	(*EventFinalizedCheckpoint)(nil), // 4: ethereum.eth.v1.EventFinalizedCheckpoint
	(*gateway.EventSource)(nil),      // 5: gateway.EventSource
}
var file_eth_v1_events_service_proto_depIdxs = []int32{
	0, // 0: ethereum.eth.v1.Events.StreamEvents:input_type -> ethereum.eth.v1.StreamEventsRequest
	5, // 1: ethereum.eth.v1.Events.StreamEvents:output_type -> gateway.EventSource
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eth_v1_events_service_proto_init() }
func file_eth_v1_events_service_proto_init() {
	if File_eth_v1_events_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eth_v1_events_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEventsRequest); i {
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
		file_eth_v1_events_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHead); i {
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
		file_eth_v1_events_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventBlock); i {
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
		file_eth_v1_events_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventChainReorg); i {
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
		file_eth_v1_events_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventFinalizedCheckpoint); i {
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
			RawDescriptor: file_eth_v1_events_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eth_v1_events_service_proto_goTypes,
		DependencyIndexes: file_eth_v1_events_service_proto_depIdxs,
		MessageInfos:      file_eth_v1_events_service_proto_msgTypes,
	}.Build()
	File_eth_v1_events_service_proto = out.File
	file_eth_v1_events_service_proto_rawDesc = nil
	file_eth_v1_events_service_proto_goTypes = nil
	file_eth_v1_events_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (Events_StreamEventsClient, error)
}

type eventsClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsClient(cc grpc.ClientConnInterface) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (Events_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Events_serviceDesc.Streams[0], "/ethereum.eth.v1.Events/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Events_StreamEventsClient interface {
	Recv() (*gateway.EventSource, error)
	grpc.ClientStream
}

type eventsStreamEventsClient struct {
	grpc.ClientStream
}

func (x *eventsStreamEventsClient) Recv() (*gateway.EventSource, error) {
	m := new(gateway.EventSource)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	StreamEvents(*StreamEventsRequest, Events_StreamEventsServer) error
}

// UnimplementedEventsServer can be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (*UnimplementedEventsServer) StreamEvents(*StreamEventsRequest, Events_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServer).StreamEvents(m, &eventsStreamEventsServer{stream})
}

type Events_StreamEventsServer interface {
	Send(*gateway.EventSource) error
	grpc.ServerStream
}

type eventsStreamEventsServer struct {
	grpc.ServerStream
}

func (x *eventsStreamEventsServer) Send(m *gateway.EventSource) error {
	return x.ServerStream.SendMsg(m)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1.Events",
	HandlerType: (*EventsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _Events_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eth/v1/events_service.proto",
}
