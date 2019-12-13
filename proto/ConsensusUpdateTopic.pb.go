// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ConsensusUpdateTopic.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// All fields left null will not be updated.
// See [ConsensusService.updateTopic()](#proto.ConsensusService)
type ConsensusUpdateTopicTransactionBody struct {
	TopicID *TopicID `protobuf:"bytes,1,opt,name=topicID,proto3" json:"topicID,omitempty"`
	// Short publicly visible memo about the topic. No guarantee of uniqueness. Null for "do not update".
	Memo           *wrappers.StringValue `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	ValidStartTime *Timestamp            `protobuf:"bytes,3,opt,name=validStartTime,proto3" json:"validStartTime,omitempty"` // Deprecated: Do not use.
	// Effective consensus timestamp at (and after) which all consensus transactions and queries will fail.
	// The expirationTime may be no longer than 90 days from the consensus timestamp of this transaction.
	// If unspecified, no change.
	ExpirationTime *Timestamp `protobuf:"bytes,4,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	// Access control for update/delete of the topic.
	// If unspecified, no change.
	// If empty keyList - the adminKey is cleared.
	AdminKey *Key `protobuf:"bytes,6,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	// Access control for ConsensusService.submitMessage.
	// If unspecified, no change.
	// If empty keyList - the submitKey is cleared.
	SubmitKey *Key `protobuf:"bytes,7,opt,name=submitKey,proto3" json:"submitKey,omitempty"`
	// The amount of time to extend the topic's lifetime automatically at expirationTime if the autoRenewAccount is
	// configured and has funds.
	// Limited to a maximum of 90 days (server-side configuration which may change).
	// If unspecified, no change.
	AutoRenewPeriod *Duration `protobuf:"bytes,8,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	// Optional account to be used at the topic's expirationTime to extend the life of the topic.
	// The topic lifetime will be extended up to a maximum of the autoRenewPeriod or however long the topic
	// can be extended using all funds on the account (whichever is the smaller duration/amount).
	// If specified as the default value (0.0.0), the autoRenewAccount will be removed.
	AutoRenewAccount     *AccountID `protobuf:"bytes,9,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConsensusUpdateTopicTransactionBody) Reset()         { *m = ConsensusUpdateTopicTransactionBody{} }
func (m *ConsensusUpdateTopicTransactionBody) String() string { return proto.CompactTextString(m) }
func (*ConsensusUpdateTopicTransactionBody) ProtoMessage()    {}
func (*ConsensusUpdateTopicTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d59c9b7c204517, []int{0}
}

func (m *ConsensusUpdateTopicTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusUpdateTopicTransactionBody.Unmarshal(m, b)
}
func (m *ConsensusUpdateTopicTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusUpdateTopicTransactionBody.Marshal(b, m, deterministic)
}
func (m *ConsensusUpdateTopicTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusUpdateTopicTransactionBody.Merge(m, src)
}
func (m *ConsensusUpdateTopicTransactionBody) XXX_Size() int {
	return xxx_messageInfo_ConsensusUpdateTopicTransactionBody.Size(m)
}
func (m *ConsensusUpdateTopicTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusUpdateTopicTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusUpdateTopicTransactionBody proto.InternalMessageInfo

func (m *ConsensusUpdateTopicTransactionBody) GetTopicID() *TopicID {
	if m != nil {
		return m.TopicID
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetMemo() *wrappers.StringValue {
	if m != nil {
		return m.Memo
	}
	return nil
}

// Deprecated: Do not use.
func (m *ConsensusUpdateTopicTransactionBody) GetValidStartTime() *Timestamp {
	if m != nil {
		return m.ValidStartTime
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetExpirationTime() *Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetAdminKey() *Key {
	if m != nil {
		return m.AdminKey
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetSubmitKey() *Key {
	if m != nil {
		return m.SubmitKey
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetAutoRenewPeriod() *Duration {
	if m != nil {
		return m.AutoRenewPeriod
	}
	return nil
}

func (m *ConsensusUpdateTopicTransactionBody) GetAutoRenewAccount() *AccountID {
	if m != nil {
		return m.AutoRenewAccount
	}
	return nil
}

func init() {
	proto.RegisterType((*ConsensusUpdateTopicTransactionBody)(nil), "proto.ConsensusUpdateTopicTransactionBody")
}

func init() { proto.RegisterFile("ConsensusUpdateTopic.proto", fileDescriptor_18d59c9b7c204517) }

var fileDescriptor_18d59c9b7c204517 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x5f, 0x6b, 0xf2, 0x30,
	0x14, 0xc6, 0xa9, 0xfa, 0xfa, 0x27, 0x2f, 0x54, 0xc9, 0x55, 0x29, 0x43, 0xc6, 0x06, 0xc3, 0xab,
	0x3a, 0xb6, 0x9b, 0x0d, 0xbc, 0x59, 0xe7, 0x8d, 0x78, 0x23, 0xb1, 0xdb, 0x7d, 0x6c, 0xcf, 0xda,
	0x80, 0x4d, 0x42, 0x92, 0xce, 0xf5, 0xfb, 0xed, 0x83, 0x8d, 0xa6, 0xad, 0x63, 0xea, 0x55, 0x72,
	0x9e, 0xf3, 0x7b, 0x72, 0x9e, 0x43, 0x90, 0xff, 0x2a, 0xb8, 0x06, 0xae, 0x0b, 0xfd, 0x26, 0x13,
	0x6a, 0x20, 0x12, 0x92, 0xc5, 0x81, 0x54, 0xc2, 0x08, 0xfc, 0xcf, 0x1e, 0xfe, 0x34, 0x15, 0x22,
	0xdd, 0xc3, 0xdc, 0x56, 0xbb, 0xe2, 0x63, 0x7e, 0x50, 0x54, 0x4a, 0x50, 0xba, 0xc6, 0xfc, 0x49,
	0x48, 0x35, 0x8b, 0xa3, 0x52, 0x42, 0xab, 0xb8, 0xcb, 0x42, 0x51, 0xc3, 0x04, 0x6f, 0xea, 0x71,
	0xc4, 0x72, 0xd0, 0x86, 0xe6, 0xb2, 0x16, 0x6e, 0xbe, 0xbb, 0xe8, 0xf6, 0xd2, 0xe0, 0x48, 0x51,
	0xae, 0x69, 0x5c, 0x59, 0x43, 0x91, 0x94, 0x78, 0x86, 0x06, 0xa6, 0xd2, 0x57, 0x4b, 0xcf, 0xb9,
	0x76, 0x66, 0xff, 0x1f, 0xdc, 0xfa, 0x81, 0x20, 0xaa, 0x55, 0xd2, 0xb6, 0xf1, 0x3d, 0xea, 0xe5,
	0x90, 0x0b, 0xaf, 0x63, 0xb1, 0xab, 0xa0, 0xce, 0x1c, 0xb4, 0x99, 0x83, 0xad, 0x51, 0x8c, 0xa7,
	0xef, 0x74, 0x5f, 0x00, 0xb1, 0x24, 0x5e, 0x20, 0xf7, 0x93, 0xee, 0x59, 0xb2, 0x35, 0x54, 0x99,
	0x2a, 0xa0, 0xd7, 0xb5, 0xde, 0x49, 0x3b, 0xa2, 0xcd, 0x1c, 0x76, 0x3c, 0x87, 0x9c, 0xb0, 0xf8,
	0x09, 0xb9, 0xf0, 0x25, 0x59, 0xbd, 0xa6, 0x75, 0xf7, 0x2e, 0xbb, 0xc9, 0x09, 0x87, 0xef, 0xd0,
	0x90, 0x26, 0x39, 0xe3, 0x6b, 0x28, 0xbd, 0xbe, 0xf5, 0xa0, 0xc6, 0xb3, 0x86, 0x92, 0x1c, 0x7b,
	0x78, 0x86, 0x46, 0xba, 0xd8, 0xe5, 0xcc, 0x54, 0xe0, 0xe0, 0x0c, 0xfc, 0x6d, 0xe2, 0x67, 0x34,
	0xa6, 0x85, 0x11, 0x04, 0x38, 0x1c, 0x36, 0xa0, 0x98, 0x48, 0xbc, 0xa1, 0xe5, 0xc7, 0x0d, 0xdf,
	0x7e, 0x07, 0x39, 0xe5, 0xf0, 0x02, 0x4d, 0x8e, 0xd2, 0x4b, 0x1c, 0x8b, 0x82, 0x1b, 0x6f, 0xf4,
	0x67, 0x91, 0x46, 0x5d, 0x2d, 0xc9, 0x19, 0x19, 0x4e, 0x91, 0x1f, 0x8b, 0x3c, 0xc8, 0x20, 0x01,
	0x45, 0x83, 0x8c, 0xea, 0x2c, 0x55, 0x54, 0x66, 0xb5, 0x73, 0xe3, 0xec, 0xfa, 0xf6, 0xf2, 0xf8,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x65, 0x79, 0xd1, 0xf4, 0x65, 0x02, 0x00, 0x00,
}
