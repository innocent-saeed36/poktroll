// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pocket/gateway/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgStakeGateway struct {
	Address string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Stake   *types.Coin `protobuf:"bytes,2,opt,name=stake,proto3" json:"stake,omitempty"`
}

func (m *MsgStakeGateway) Reset()         { *m = MsgStakeGateway{} }
func (m *MsgStakeGateway) String() string { return proto.CompactTextString(m) }
func (*MsgStakeGateway) ProtoMessage()    {}
func (*MsgStakeGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5d84c597736301, []int{0}
}
func (m *MsgStakeGateway) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStakeGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStakeGateway.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStakeGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStakeGateway.Merge(m, src)
}
func (m *MsgStakeGateway) XXX_Size() int {
	return m.Size()
}
func (m *MsgStakeGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStakeGateway.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStakeGateway proto.InternalMessageInfo

func (m *MsgStakeGateway) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgStakeGateway) GetStake() *types.Coin {
	if m != nil {
		return m.Stake
	}
	return nil
}

type MsgStakeGatewayResponse struct {
}

func (m *MsgStakeGatewayResponse) Reset()         { *m = MsgStakeGatewayResponse{} }
func (m *MsgStakeGatewayResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStakeGatewayResponse) ProtoMessage()    {}
func (*MsgStakeGatewayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5d84c597736301, []int{1}
}
func (m *MsgStakeGatewayResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStakeGatewayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStakeGatewayResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStakeGatewayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStakeGatewayResponse.Merge(m, src)
}
func (m *MsgStakeGatewayResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStakeGatewayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStakeGatewayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStakeGatewayResponse proto.InternalMessageInfo

type MsgUnstakeGateway struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgUnstakeGateway) Reset()         { *m = MsgUnstakeGateway{} }
func (m *MsgUnstakeGateway) String() string { return proto.CompactTextString(m) }
func (*MsgUnstakeGateway) ProtoMessage()    {}
func (*MsgUnstakeGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5d84c597736301, []int{2}
}
func (m *MsgUnstakeGateway) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnstakeGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnstakeGateway.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnstakeGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnstakeGateway.Merge(m, src)
}
func (m *MsgUnstakeGateway) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnstakeGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnstakeGateway.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnstakeGateway proto.InternalMessageInfo

func (m *MsgUnstakeGateway) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgUnstakeGatewayResponse struct {
}

func (m *MsgUnstakeGatewayResponse) Reset()         { *m = MsgUnstakeGatewayResponse{} }
func (m *MsgUnstakeGatewayResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUnstakeGatewayResponse) ProtoMessage()    {}
func (*MsgUnstakeGatewayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc5d84c597736301, []int{3}
}
func (m *MsgUnstakeGatewayResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnstakeGatewayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnstakeGatewayResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnstakeGatewayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnstakeGatewayResponse.Merge(m, src)
}
func (m *MsgUnstakeGatewayResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnstakeGatewayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnstakeGatewayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnstakeGatewayResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgStakeGateway)(nil), "pocket.gateway.MsgStakeGateway")
	proto.RegisterType((*MsgStakeGatewayResponse)(nil), "pocket.gateway.MsgStakeGatewayResponse")
	proto.RegisterType((*MsgUnstakeGateway)(nil), "pocket.gateway.MsgUnstakeGateway")
	proto.RegisterType((*MsgUnstakeGatewayResponse)(nil), "pocket.gateway.MsgUnstakeGatewayResponse")
}

func init() { proto.RegisterFile("pocket/gateway/tx.proto", fileDescriptor_bc5d84c597736301) }

var fileDescriptor_bc5d84c597736301 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xc8, 0x4f, 0xce,
	0x4e, 0x2d, 0xd1, 0x4f, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x48, 0xe8, 0x41, 0x25, 0xa4, 0xc4, 0x93, 0xf3, 0x8b, 0x73,
	0xf3, 0x8b, 0xf5, 0x73, 0x8b, 0xd3, 0xf5, 0xcb, 0x0c, 0x41, 0x14, 0x44, 0xa1, 0x94, 0x24, 0x44,
	0x22, 0x1e, 0xcc, 0xd3, 0x87, 0x70, 0xa0, 0x52, 0x72, 0x50, 0x3d, 0x49, 0x89, 0xc5, 0xa9, 0xfa,
	0x65, 0x86, 0x49, 0xa9, 0x25, 0x89, 0x86, 0xfa, 0xc9, 0xf9, 0x99, 0x79, 0x10, 0x79, 0xa5, 0x16,
	0x46, 0x2e, 0x7e, 0xdf, 0xe2, 0xf4, 0xe0, 0x92, 0xc4, 0xec, 0x54, 0x77, 0x88, 0x3d, 0x42, 0x46,
	0x5c, 0xec, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e,
	0x12, 0x97, 0xb6, 0xe8, 0x8a, 0x40, 0x8d, 0x75, 0x84, 0xc8, 0x04, 0x97, 0x14, 0x65, 0xe6, 0xa5,
	0x07, 0xc1, 0x14, 0x0a, 0xe9, 0x73, 0xb1, 0x16, 0x83, 0xcc, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x36, 0x92, 0xd4, 0x83, 0x2a, 0x07, 0xd9, 0xab, 0x07, 0xb5, 0x57, 0xcf, 0x39, 0x3f, 0x33, 0x2f,
	0x08, 0xa2, 0xce, 0x8a, 0xa7, 0xe9, 0xf9, 0x06, 0x2d, 0x98, 0x76, 0x25, 0x49, 0x2e, 0x71, 0x34,
	0x57, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0x59, 0x73, 0x09, 0xfa, 0x16, 0xa7,
	0x87, 0xe6, 0x15, 0x23, 0x3b, 0x51, 0x02, 0xcd, 0x89, 0x70, 0x87, 0xa0, 0x99, 0x2b, 0xcd, 0x25,
	0x89, 0xa1, 0x19, 0x66, 0xb2, 0xd1, 0x7e, 0x46, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x08, 0x2e,
	0x1e, 0x14, 0xff, 0xcb, 0xeb, 0xa1, 0x06, 0xbc, 0x1e, 0x9a, 0xd3, 0xa4, 0xd4, 0x09, 0x28, 0x80,
	0xd9, 0x20, 0x14, 0xc7, 0xc5, 0x87, 0xe6, 0x70, 0x45, 0x2c, 0x5a, 0x51, 0x95, 0x48, 0x69, 0x12,
	0x54, 0x02, 0x33, 0xdf, 0xc9, 0xeb, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0x0c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x0b, 0xf2, 0xb3, 0x4b,
	0x74, 0xf3, 0x52, 0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0xc1, 0x9c, 0xa2, 0xfc, 0x9c, 0x1c, 0xfd, 0x0a,
	0x44, 0x82, 0xab, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x27, 0x08, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0x03, 0x5b, 0x7d, 0x8f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	StakeGateway(ctx context.Context, in *MsgStakeGateway, opts ...grpc.CallOption) (*MsgStakeGatewayResponse, error)
	UnstakeGateway(ctx context.Context, in *MsgUnstakeGateway, opts ...grpc.CallOption) (*MsgUnstakeGatewayResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) StakeGateway(ctx context.Context, in *MsgStakeGateway, opts ...grpc.CallOption) (*MsgStakeGatewayResponse, error) {
	out := new(MsgStakeGatewayResponse)
	err := c.cc.Invoke(ctx, "/pocket.gateway.Msg/StakeGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnstakeGateway(ctx context.Context, in *MsgUnstakeGateway, opts ...grpc.CallOption) (*MsgUnstakeGatewayResponse, error) {
	out := new(MsgUnstakeGatewayResponse)
	err := c.cc.Invoke(ctx, "/pocket.gateway.Msg/UnstakeGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	StakeGateway(context.Context, *MsgStakeGateway) (*MsgStakeGatewayResponse, error)
	UnstakeGateway(context.Context, *MsgUnstakeGateway) (*MsgUnstakeGatewayResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) StakeGateway(ctx context.Context, req *MsgStakeGateway) (*MsgStakeGatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakeGateway not implemented")
}
func (*UnimplementedMsgServer) UnstakeGateway(ctx context.Context, req *MsgUnstakeGateway) (*MsgUnstakeGatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnstakeGateway not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_StakeGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStakeGateway)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StakeGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pocket.gateway.Msg/StakeGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StakeGateway(ctx, req.(*MsgStakeGateway))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnstakeGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnstakeGateway)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnstakeGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pocket.gateway.Msg/UnstakeGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnstakeGateway(ctx, req.(*MsgUnstakeGateway))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pocket.gateway.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StakeGateway",
			Handler:    _Msg_StakeGateway_Handler,
		},
		{
			MethodName: "UnstakeGateway",
			Handler:    _Msg_UnstakeGateway_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pocket/gateway/tx.proto",
}

func (m *MsgStakeGateway) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStakeGateway) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStakeGateway) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Stake != nil {
		{
			size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStakeGatewayResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStakeGatewayResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStakeGatewayResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUnstakeGateway) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnstakeGateway) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnstakeGateway) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnstakeGatewayResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnstakeGatewayResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnstakeGatewayResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgStakeGateway) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Stake != nil {
		l = m.Stake.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgStakeGatewayResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUnstakeGateway) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUnstakeGatewayResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgStakeGateway) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStakeGateway: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStakeGateway: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStakeGatewayResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStakeGatewayResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStakeGatewayResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUnstakeGateway) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUnstakeGateway: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnstakeGateway: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUnstakeGatewayResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUnstakeGatewayResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnstakeGatewayResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
