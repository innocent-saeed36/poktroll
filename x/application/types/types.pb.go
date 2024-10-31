// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/application/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_sortkeys "github.com/cosmos/gogoproto/sortkeys"
	types1 "github.com/pokt-network/poktroll/x/shared/types"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Application defines the type used to store an on-chain definition and state for an application
type Application struct {
	Address string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Stake   *types.Coin `protobuf:"bytes,2,opt,name=stake,proto3" json:"stake,omitempty"`
	// CRITICAL_DEV_NOTE: The number of service_configs must be EXACTLY ONE.
	// This prevents applications from over-servicing.
	// The field is kept repeated (a list) for both legacy and future logic reaosns.
	// References:
	//  - https://github.com/pokt-network/poktroll/pull/750#discussion_r1735025033
	//  - https://www.notion.so/buildwithgrove/Off-chain-Application-Stake-Tracking-6a8bebb107db4f7f9dc62cbe7ba555f7
	ServiceConfigs []*types1.ApplicationServiceConfig `protobuf:"bytes,3,rep,name=service_configs,json=serviceConfigs,proto3" json:"service_configs,omitempty"`
	// TODO_BETA: Rename `delegatee_gateway_addresses` to `gateway_addresses_delegated_to`.
	// Ensure to rename all relevant configs, comments, variables, function names, etc as well.
	DelegateeGatewayAddresses []string `protobuf:"bytes,4,rep,name=delegatee_gateway_addresses,json=delegateeGatewayAddresses,proto3" json:"delegatee_gateway_addresses,omitempty"`
	// A map from sessionEndHeights to a list of Gateways.
	// The key is the height of the last block of the session during which the
	// respective undelegation was committed.
	// The value is a list of gateways being undelegated from.
	// TODO_DOCUMENT(@red-0ne): Need to document the flow from this comment
	// so its clear to everyone why this is necessary; https://github.com/pokt-network/poktroll/issues/476#issuecomment-2052639906.
	PendingUndelegations map[uint64]UndelegatingGatewayList `protobuf:"bytes,5,rep,name=pending_undelegations,json=pendingUndelegations,proto3" json:"pending_undelegations" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The end height of the session at which an application initiated its unstaking process.
	// If the application did not unstake, this value will be 0.
	UnstakeSessionEndHeight uint64                      `protobuf:"varint,6,opt,name=unstake_session_end_height,json=unstakeSessionEndHeight,proto3" json:"unstake_session_end_height,omitempty"`
	PendingTransfer         *PendingApplicationTransfer `protobuf:"bytes,7,opt,name=pending_transfer,json=pendingTransfer,proto3" json:"pending_transfer,omitempty"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_1899440439257283, []int{0}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return m.Size()
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Application) GetStake() *types.Coin {
	if m != nil {
		return m.Stake
	}
	return nil
}

func (m *Application) GetServiceConfigs() []*types1.ApplicationServiceConfig {
	if m != nil {
		return m.ServiceConfigs
	}
	return nil
}

func (m *Application) GetDelegateeGatewayAddresses() []string {
	if m != nil {
		return m.DelegateeGatewayAddresses
	}
	return nil
}

func (m *Application) GetPendingUndelegations() map[uint64]UndelegatingGatewayList {
	if m != nil {
		return m.PendingUndelegations
	}
	return nil
}

func (m *Application) GetUnstakeSessionEndHeight() uint64 {
	if m != nil {
		return m.UnstakeSessionEndHeight
	}
	return 0
}

func (m *Application) GetPendingTransfer() *PendingApplicationTransfer {
	if m != nil {
		return m.PendingTransfer
	}
	return nil
}

// UndelegatingGatewayList is used as the Value of `pending_undelegations`.
// It is required to store a repeated list of strings as a map value.
type UndelegatingGatewayList struct {
	GatewayAddresses []string `protobuf:"bytes,2,rep,name=gateway_addresses,json=gatewayAddresses,proto3" json:"gateway_addresses,omitempty"`
}

func (m *UndelegatingGatewayList) Reset()         { *m = UndelegatingGatewayList{} }
func (m *UndelegatingGatewayList) String() string { return proto.CompactTextString(m) }
func (*UndelegatingGatewayList) ProtoMessage()    {}
func (*UndelegatingGatewayList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1899440439257283, []int{1}
}
func (m *UndelegatingGatewayList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UndelegatingGatewayList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *UndelegatingGatewayList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UndelegatingGatewayList.Merge(m, src)
}
func (m *UndelegatingGatewayList) XXX_Size() int {
	return m.Size()
}
func (m *UndelegatingGatewayList) XXX_DiscardUnknown() {
	xxx_messageInfo_UndelegatingGatewayList.DiscardUnknown(m)
}

var xxx_messageInfo_UndelegatingGatewayList proto.InternalMessageInfo

func (m *UndelegatingGatewayList) GetGatewayAddresses() []string {
	if m != nil {
		return m.GatewayAddresses
	}
	return nil
}

// PendingTransfer is used to store the details of a pending transfer.
// It is only intended to be used inside of an Application object.
type PendingApplicationTransfer struct {
	DestinationAddress string `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	SessionEndHeight   uint64 `protobuf:"varint,2,opt,name=session_end_height,json=sessionEndHeight,proto3" json:"session_end_height,omitempty"`
}

func (m *PendingApplicationTransfer) Reset()         { *m = PendingApplicationTransfer{} }
func (m *PendingApplicationTransfer) String() string { return proto.CompactTextString(m) }
func (*PendingApplicationTransfer) ProtoMessage()    {}
func (*PendingApplicationTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1899440439257283, []int{2}
}
func (m *PendingApplicationTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingApplicationTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PendingApplicationTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingApplicationTransfer.Merge(m, src)
}
func (m *PendingApplicationTransfer) XXX_Size() int {
	return m.Size()
}
func (m *PendingApplicationTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingApplicationTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_PendingApplicationTransfer proto.InternalMessageInfo

func (m *PendingApplicationTransfer) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *PendingApplicationTransfer) GetSessionEndHeight() uint64 {
	if m != nil {
		return m.SessionEndHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Application)(nil), "poktroll.application.Application")
	proto.RegisterMapType((map[uint64]UndelegatingGatewayList)(nil), "poktroll.application.Application.PendingUndelegationsEntry")
	proto.RegisterType((*UndelegatingGatewayList)(nil), "poktroll.application.UndelegatingGatewayList")
	proto.RegisterType((*PendingApplicationTransfer)(nil), "poktroll.application.PendingApplicationTransfer")
}

func init() { proto.RegisterFile("poktroll/application/types.proto", fileDescriptor_1899440439257283) }

var fileDescriptor_1899440439257283 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0xfe, 0xd9, 0x34, 0x57, 0x62, 0xc5, 0x14, 0x2d, 0x2d, 0x10, 0xa2, 0x9e, 0x8a,
	0x44, 0x13, 0x56, 0x38, 0x20, 0x76, 0x6a, 0xab, 0x09, 0x26, 0x71, 0x40, 0x29, 0x5c, 0x00, 0x29,
	0x4a, 0x13, 0x2f, 0xb5, 0xda, 0xd9, 0x51, 0x5e, 0xb7, 0xa3, 0xdf, 0x82, 0x0b, 0x5f, 0x83, 0x13,
	0x1f, 0x62, 0xc7, 0x89, 0xd3, 0x4e, 0x08, 0xb5, 0x5f, 0x64, 0x4a, 0xec, 0x74, 0xd1, 0xd6, 0x4a,
	0xbb, 0x54, 0x89, 0xdf, 0xe7, 0x79, 0xfb, 0xf3, 0xfb, 0x38, 0x46, 0x66, 0xc4, 0x27, 0x22, 0xe6,
	0xd3, 0xa9, 0xed, 0x45, 0xd1, 0x94, 0xfa, 0x9e, 0xa0, 0x9c, 0xd9, 0x62, 0x11, 0x11, 0xb0, 0xa2,
	0x98, 0x0b, 0x8e, 0xeb, 0x99, 0xc2, 0xca, 0x29, 0x9a, 0xf5, 0x90, 0x87, 0x3c, 0x15, 0xd8, 0xc9,
	0x93, 0xd4, 0x36, 0x0d, 0x9f, 0xc3, 0x19, 0x07, 0x7b, 0xe4, 0x01, 0xb1, 0xe7, 0x87, 0x23, 0x22,
	0xbc, 0x43, 0xdb, 0xe7, 0x94, 0xa9, 0x7a, 0x43, 0xd6, 0x5d, 0x69, 0x94, 0x2f, 0xaa, 0xf4, 0x6c,
	0x0d, 0x02, 0x63, 0x2f, 0x26, 0x81, 0x0d, 0x24, 0x9e, 0x53, 0x9f, 0xc8, 0x72, 0xeb, 0x77, 0x05,
	0x55, 0x7b, 0x37, 0xff, 0x8f, 0xbb, 0x68, 0xd7, 0x0b, 0x82, 0x98, 0x00, 0xe8, 0x9a, 0xa9, 0xb5,
	0xf7, 0xfa, 0xfa, 0xdf, 0x3f, 0x9d, 0xba, 0xea, 0xd8, 0x93, 0x95, 0xa1, 0x88, 0x29, 0x0b, 0x9d,
	0x4c, 0x88, 0x6d, 0x54, 0x01, 0xe1, 0x4d, 0x88, 0x5e, 0x34, 0xb5, 0x76, 0xb5, 0xdb, 0xb0, 0x94,
	0x3c, 0xa1, 0xb5, 0x14, 0xad, 0x35, 0xe0, 0x94, 0x39, 0x52, 0x87, 0x1d, 0xb4, 0xaf, 0x28, 0x5c,
	0x9f, 0xb3, 0x53, 0x1a, 0x82, 0x5e, 0x32, 0x4b, 0xed, 0x6a, 0xf7, 0x85, 0xb5, 0x1e, 0x8a, 0xa4,
	0xb5, 0x72, 0x6c, 0x43, 0x69, 0x19, 0xa4, 0x0e, 0xe7, 0x01, 0xe4, 0x5f, 0x01, 0x7f, 0x47, 0x4f,
	0x02, 0x32, 0x25, 0xa1, 0x27, 0x08, 0x71, 0x93, 0xdf, 0x73, 0x6f, 0xe1, 0x2a, 0x42, 0x02, 0x7a,
	0xd9, 0x2c, 0xb5, 0xf7, 0xfa, 0x4f, 0x2f, 0xfe, 0x3d, 0x2f, 0x6c, 0xdd, 0x50, 0x63, 0xdd, 0xe0,
	0xbd, 0xf4, 0xf7, 0x32, 0x3b, 0x9e, 0xa3, 0xc7, 0x11, 0x61, 0x01, 0x65, 0xa1, 0x3b, 0x63, 0x4a,
	0x46, 0x39, 0x03, 0xbd, 0x92, 0x72, 0x1f, 0x59, 0x9b, 0xc2, 0xcc, 0xc3, 0x5b, 0x9f, 0xa4, 0xfd,
	0x4b, 0xde, 0x7d, 0xcc, 0x44, 0xbc, 0xe8, 0x97, 0x13, 0x28, 0xa7, 0x1e, 0x6d, 0x10, 0xe0, 0x23,
	0xd4, 0x9c, 0xb1, 0x74, 0x68, 0x2e, 0x10, 0x00, 0xca, 0x99, 0x4b, 0x58, 0xe0, 0x8e, 0x09, 0x0d,
	0xc7, 0x42, 0xdf, 0x31, 0xb5, 0x76, 0xd9, 0x39, 0x50, 0x8a, 0xa1, 0x14, 0x1c, 0xb3, 0xe0, 0x43,
	0x5a, 0xc6, 0xdf, 0x50, 0x2d, 0x83, 0x16, 0xb1, 0xc7, 0xe0, 0x94, 0xc4, 0xfa, 0x6e, 0x1a, 0xd1,
	0xab, 0xcd, 0xbc, 0x8a, 0x31, 0x87, 0xfd, 0x59, 0xf9, 0x9c, 0x7d, 0xd5, 0x29, 0x5b, 0x68, 0xce,
	0x51, 0x63, 0xeb, 0x96, 0x70, 0x0d, 0x95, 0x26, 0x64, 0x91, 0x9e, 0xa0, 0xb2, 0x93, 0x3c, 0xe2,
	0x01, 0xaa, 0xcc, 0xbd, 0xe9, 0x2c, 0x3b, 0x23, 0x9d, 0xcd, 0x00, 0x37, 0xad, 0x58, 0xa8, 0x32,
	0xf8, 0x48, 0x41, 0x38, 0xd2, 0xfb, 0xae, 0xf8, 0x56, 0x6b, 0x05, 0xe8, 0x60, 0x8b, 0x0a, 0x9f,
	0xa0, 0x87, 0x77, 0x83, 0x2f, 0xde, 0x23, 0xf8, 0x5a, 0x78, 0x2b, 0xef, 0xd6, 0x2f, 0x0d, 0x35,
	0xb7, 0x4f, 0x03, 0x9f, 0xa0, 0x47, 0x01, 0x01, 0x41, 0x59, 0xba, 0xec, 0xde, 0xf7, 0x8b, 0xc1,
	0x39, 0x93, 0xaa, 0xe0, 0x97, 0x08, 0x6f, 0x48, 0xb6, 0x98, 0x4e, 0xae, 0x06, 0xb7, 0x22, 0xed,
	0x3b, 0x17, 0x4b, 0x43, 0xbb, 0x5c, 0x1a, 0xda, 0xd5, 0xd2, 0xd0, 0xfe, 0x2f, 0x0d, 0xed, 0xe7,
	0xca, 0x28, 0x5c, 0xae, 0x8c, 0xc2, 0xd5, 0xca, 0x28, 0x7c, 0x7d, 0x13, 0x52, 0x31, 0x9e, 0x8d,
	0x2c, 0x9f, 0x9f, 0xd9, 0xc9, 0x7c, 0x3b, 0x8c, 0x88, 0x73, 0x1e, 0x4f, 0xec, 0xf5, 0x1d, 0xf0,
	0xe3, 0xee, 0x75, 0x34, 0xda, 0x49, 0x6f, 0x82, 0xd7, 0xd7, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8c,
	0x2e, 0xfc, 0xb7, 0xb3, 0x04, 0x00, 0x00,
}

func (m *Application) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Application) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Application) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PendingTransfer != nil {
		{
			size, err := m.PendingTransfer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.UnstakeSessionEndHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.UnstakeSessionEndHeight))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PendingUndelegations) > 0 {
		keysForPendingUndelegations := make([]uint64, 0, len(m.PendingUndelegations))
		for k := range m.PendingUndelegations {
			keysForPendingUndelegations = append(keysForPendingUndelegations, uint64(k))
		}
		github_com_cosmos_gogoproto_sortkeys.Uint64s(keysForPendingUndelegations)
		for iNdEx := len(keysForPendingUndelegations) - 1; iNdEx >= 0; iNdEx-- {
			v := m.PendingUndelegations[uint64(keysForPendingUndelegations[iNdEx])]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i = encodeVarintTypes(dAtA, i, uint64(keysForPendingUndelegations[iNdEx]))
			i--
			dAtA[i] = 0x8
			i = encodeVarintTypes(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.DelegateeGatewayAddresses) > 0 {
		for iNdEx := len(m.DelegateeGatewayAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DelegateeGatewayAddresses[iNdEx])
			copy(dAtA[i:], m.DelegateeGatewayAddresses[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.DelegateeGatewayAddresses[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ServiceConfigs) > 0 {
		for iNdEx := len(m.ServiceConfigs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceConfigs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Stake != nil {
		{
			size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UndelegatingGatewayList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UndelegatingGatewayList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UndelegatingGatewayList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GatewayAddresses) > 0 {
		for iNdEx := len(m.GatewayAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.GatewayAddresses[iNdEx])
			copy(dAtA[i:], m.GatewayAddresses[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.GatewayAddresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *PendingApplicationTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingApplicationTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingApplicationTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SessionEndHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SessionEndHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Application) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Stake != nil {
		l = m.Stake.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.ServiceConfigs) > 0 {
		for _, e := range m.ServiceConfigs {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.DelegateeGatewayAddresses) > 0 {
		for _, s := range m.DelegateeGatewayAddresses {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.PendingUndelegations) > 0 {
		for k, v := range m.PendingUndelegations {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + sovTypes(uint64(k)) + 1 + l + sovTypes(uint64(l))
			n += mapEntrySize + 1 + sovTypes(uint64(mapEntrySize))
		}
	}
	if m.UnstakeSessionEndHeight != 0 {
		n += 1 + sovTypes(uint64(m.UnstakeSessionEndHeight))
	}
	if m.PendingTransfer != nil {
		l = m.PendingTransfer.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *UndelegatingGatewayList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GatewayAddresses) > 0 {
		for _, s := range m.GatewayAddresses {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *PendingApplicationTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.SessionEndHeight != 0 {
		n += 1 + sovTypes(uint64(m.SessionEndHeight))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Application) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Application: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Application: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stake == nil {
				m.Stake = &types.Coin{}
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceConfigs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceConfigs = append(m.ServiceConfigs, &types1.ApplicationServiceConfig{})
			if err := m.ServiceConfigs[len(m.ServiceConfigs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateeGatewayAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegateeGatewayAddresses = append(m.DelegateeGatewayAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingUndelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PendingUndelegations == nil {
				m.PendingUndelegations = make(map[uint64]UndelegatingGatewayList)
			}
			var mapkey uint64
			mapvalue := &UndelegatingGatewayList{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTypes
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTypes
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &UndelegatingGatewayList{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTypes(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTypes
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.PendingUndelegations[mapkey] = *mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakeSessionEndHeight", wireType)
			}
			m.UnstakeSessionEndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnstakeSessionEndHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingTransfer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PendingTransfer == nil {
				m.PendingTransfer = &PendingApplicationTransfer{}
			}
			if err := m.PendingTransfer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UndelegatingGatewayList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: UndelegatingGatewayList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UndelegatingGatewayList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GatewayAddresses = append(m.GatewayAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PendingApplicationTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PendingApplicationTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingApplicationTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionEndHeight", wireType)
			}
			m.SessionEndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionEndHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
