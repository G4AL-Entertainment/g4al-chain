// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: g4alchain/game/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreateProject struct {
	Creator     string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Symbol      string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Delegate    []string `protobuf:"bytes,5,rep,name=delegate,proto3" json:"delegate,omitempty"`
}

func (m *MsgCreateProject) Reset()         { *m = MsgCreateProject{} }
func (m *MsgCreateProject) String() string { return proto.CompactTextString(m) }
func (*MsgCreateProject) ProtoMessage()    {}
func (*MsgCreateProject) Descriptor() ([]byte, []int) {
	return fileDescriptor_72f587bd2183429d, []int{0}
}
func (m *MsgCreateProject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateProject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateProject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateProject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateProject.Merge(m, src)
}
func (m *MsgCreateProject) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateProject) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateProject.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateProject proto.InternalMessageInfo

func (m *MsgCreateProject) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateProject) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MsgCreateProject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCreateProject) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgCreateProject) GetDelegate() []string {
	if m != nil {
		return m.Delegate
	}
	return nil
}

type MsgCreateProjectResponse struct {
}

func (m *MsgCreateProjectResponse) Reset()         { *m = MsgCreateProjectResponse{} }
func (m *MsgCreateProjectResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateProjectResponse) ProtoMessage()    {}
func (*MsgCreateProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72f587bd2183429d, []int{1}
}
func (m *MsgCreateProjectResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateProjectResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateProjectResponse.Merge(m, src)
}
func (m *MsgCreateProjectResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateProjectResponse proto.InternalMessageInfo

type MsgUpdateProject struct {
	Creator     string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Symbol      string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Delegate    []string `protobuf:"bytes,5,rep,name=delegate,proto3" json:"delegate,omitempty"`
}

func (m *MsgUpdateProject) Reset()         { *m = MsgUpdateProject{} }
func (m *MsgUpdateProject) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateProject) ProtoMessage()    {}
func (*MsgUpdateProject) Descriptor() ([]byte, []int) {
	return fileDescriptor_72f587bd2183429d, []int{2}
}
func (m *MsgUpdateProject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateProject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateProject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateProject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateProject.Merge(m, src)
}
func (m *MsgUpdateProject) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateProject) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateProject.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateProject proto.InternalMessageInfo

func (m *MsgUpdateProject) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateProject) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MsgUpdateProject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgUpdateProject) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgUpdateProject) GetDelegate() []string {
	if m != nil {
		return m.Delegate
	}
	return nil
}

type MsgUpdateProjectResponse struct {
}

func (m *MsgUpdateProjectResponse) Reset()         { *m = MsgUpdateProjectResponse{} }
func (m *MsgUpdateProjectResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateProjectResponse) ProtoMessage()    {}
func (*MsgUpdateProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72f587bd2183429d, []int{3}
}
func (m *MsgUpdateProjectResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateProjectResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateProjectResponse.Merge(m, src)
}
func (m *MsgUpdateProjectResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateProjectResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateProject)(nil), "g4alentertainment.g4alchain.game.MsgCreateProject")
	proto.RegisterType((*MsgCreateProjectResponse)(nil), "g4alentertainment.g4alchain.game.MsgCreateProjectResponse")
	proto.RegisterType((*MsgUpdateProject)(nil), "g4alentertainment.g4alchain.game.MsgUpdateProject")
	proto.RegisterType((*MsgUpdateProjectResponse)(nil), "g4alentertainment.g4alchain.game.MsgUpdateProjectResponse")
}

func init() { proto.RegisterFile("g4alchain/game/tx.proto", fileDescriptor_72f587bd2183429d) }

var fileDescriptor_72f587bd2183429d = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x37, 0x49, 0xcc,
	0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4f, 0xcc, 0x4d, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0x00, 0x49, 0xa4, 0xe6, 0x95, 0xa4, 0x16, 0x95, 0x24, 0x66, 0xe6,
	0xe5, 0xa6, 0xe6, 0x95, 0xe8, 0xc1, 0x95, 0xea, 0x81, 0x94, 0x4a, 0xc9, 0xa0, 0x69, 0x2d, 0x28,
	0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x81, 0xe8, 0x57, 0x9a, 0xc6, 0xc8, 0x25, 0xe0, 0x5b, 0x9c, 0xee,
	0x5c, 0x94, 0x9a, 0x58, 0x92, 0x1a, 0x00, 0x91, 0x12, 0x92, 0xe0, 0x62, 0x4f, 0x06, 0x09, 0xe4,
	0x17, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0x95,
	0xb9, 0x49, 0xf9, 0x39, 0x12, 0x4c, 0x60, 0x09, 0x28, 0x4f, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31,
	0x37, 0x55, 0x82, 0x19, 0x2c, 0x0a, 0x66, 0x0b, 0x29, 0x70, 0x71, 0xa7, 0xa4, 0x16, 0x27, 0x17,
	0x65, 0x16, 0x94, 0x64, 0xe6, 0xe7, 0x49, 0xb0, 0x80, 0xa5, 0x90, 0x85, 0x84, 0xa4, 0xb8, 0x38,
	0x52, 0x52, 0x73, 0x52, 0xd3, 0x13, 0x4b, 0x52, 0x25, 0x58, 0x15, 0x98, 0x35, 0x38, 0x83, 0xe0,
	0x7c, 0x25, 0x29, 0x2e, 0x09, 0x74, 0x77, 0x05, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0xc2,
	0x1c, 0x1d, 0x5a, 0x90, 0x32, 0x28, 0x1d, 0x8d, 0xe2, 0x2e, 0x98, 0xa3, 0x8d, 0xda, 0x99, 0xb8,
	0x98, 0x7d, 0x8b, 0xd3, 0x85, 0xea, 0xb9, 0x78, 0x51, 0x43, 0xdb, 0x48, 0x8f, 0x50, 0x1c, 0xea,
	0xa1, 0x87, 0x84, 0x94, 0x15, 0xe9, 0x7a, 0x60, 0x0e, 0x01, 0x39, 0x00, 0x35, 0xe4, 0x88, 0x73,
	0x00, 0x8a, 0x1e, 0x22, 0x1d, 0x80, 0x35, 0x24, 0x9c, 0xfc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0xca, 0x34, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0xdf, 0xdd, 0xc4, 0xd1, 0x47, 0xd7, 0x15, 0xd9, 0x02, 0x7d, 0x90, 0x05, 0xba, 0x90, 0xa4, 0x5c,
	0x01, 0xcd, 0x07, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xb4, 0x6c, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x2a, 0x11, 0xf9, 0xfc, 0x26, 0x03, 0x00, 0x00,
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
	CreateProject(ctx context.Context, in *MsgCreateProject, opts ...grpc.CallOption) (*MsgCreateProjectResponse, error)
	UpdateProject(ctx context.Context, in *MsgUpdateProject, opts ...grpc.CallOption) (*MsgUpdateProjectResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateProject(ctx context.Context, in *MsgCreateProject, opts ...grpc.CallOption) (*MsgCreateProjectResponse, error) {
	out := new(MsgCreateProjectResponse)
	err := c.cc.Invoke(ctx, "/g4alentertainment.g4alchain.game.Msg/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateProject(ctx context.Context, in *MsgUpdateProject, opts ...grpc.CallOption) (*MsgUpdateProjectResponse, error) {
	out := new(MsgUpdateProjectResponse)
	err := c.cc.Invoke(ctx, "/g4alentertainment.g4alchain.game.Msg/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateProject(context.Context, *MsgCreateProject) (*MsgCreateProjectResponse, error)
	UpdateProject(context.Context, *MsgUpdateProject) (*MsgUpdateProjectResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateProject(ctx context.Context, req *MsgCreateProject) (*MsgCreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (*UnimplementedMsgServer) UpdateProject(ctx context.Context, req *MsgUpdateProject) (*MsgUpdateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g4alentertainment.g4alchain.game.Msg/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateProject(ctx, req.(*MsgCreateProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g4alentertainment.g4alchain.game.Msg/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateProject(ctx, req.(*MsgUpdateProject))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "g4alentertainment.g4alchain.game.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Msg_CreateProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _Msg_UpdateProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "g4alchain/game/tx.proto",
}

func (m *MsgCreateProject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateProject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateProject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegate) > 0 {
		for iNdEx := len(m.Delegate) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Delegate[iNdEx])
			copy(dAtA[i:], m.Delegate[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Delegate[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateProjectResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateProjectResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateProjectResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateProject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateProject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateProject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegate) > 0 {
		for iNdEx := len(m.Delegate) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Delegate[iNdEx])
			copy(dAtA[i:], m.Delegate[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Delegate[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateProjectResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateProjectResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateProjectResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateProject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Delegate) > 0 {
		for _, s := range m.Delegate {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateProjectResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateProject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Delegate) > 0 {
		for _, s := range m.Delegate {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgUpdateProjectResponse) Size() (n int) {
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
func (m *MsgCreateProject) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateProject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateProject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
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
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegate", wireType)
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
			m.Delegate = append(m.Delegate, string(dAtA[iNdEx:postIndex]))
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
func (m *MsgCreateProjectResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateProjectResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateProjectResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateProject) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateProject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateProject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
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
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegate", wireType)
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
			m.Delegate = append(m.Delegate, string(dAtA[iNdEx:postIndex]))
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
func (m *MsgUpdateProjectResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateProjectResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateProjectResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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