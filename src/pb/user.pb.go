// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 登录请求
type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return m.Size()
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// 响应token
type TokenResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *TokenResponse) Reset()         { *m = TokenResponse{} }
func (m *TokenResponse) String() string { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()    {}
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}
func (m *TokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenResponse.Merge(m, src)
}
func (m *TokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *TokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenResponse proto.InternalMessageInfo

func (m *TokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// 列表请求
type ListRequest struct {
	Columns  []string `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Page     int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32    `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *ListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// 列表响应
type ListResponse struct {
	Rows  []*UserListModel `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	Total int32            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetRows() []*UserListModel {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *ListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

// 用户列表字段
type UserListModel struct {
	Id         int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,string" pg:"id"`
	Username   string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty" pg:"username"`
	CreateTime int64  `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty" pg:"create_time"`
}

func (m *UserListModel) Reset()         { *m = UserListModel{} }
func (m *UserListModel) String() string { return proto.CompactTextString(m) }
func (*UserListModel) ProtoMessage()    {}
func (*UserListModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}
func (m *UserListModel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserListModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserListModel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserListModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListModel.Merge(m, src)
}
func (m *UserListModel) XXX_Size() int {
	return m.Size()
}
func (m *UserListModel) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListModel.DiscardUnknown(m)
}

var xxx_messageInfo_UserListModel proto.InternalMessageInfo

func (m *UserListModel) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserListModel) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserListModel) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*TokenResponse)(nil), "pb.TokenResponse")
	proto.RegisterType((*ListRequest)(nil), "pb.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "pb.ListResponse")
	proto.RegisterType((*UserListModel)(nil), "pb.UserListModel")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0x6c, 0x36, 0xd6, 0xbc, 0x24, 0x98, 0x0e, 0x45, 0x96, 0xa5, 0x6c, 0xca, 0x40, 0x21,
	0x88, 0xcd, 0x42, 0xf5, 0x54, 0xf0, 0x12, 0xc4, 0x8b, 0xf5, 0xb2, 0x56, 0x3c, 0xca, 0x6e, 0x76,
	0x1c, 0x07, 0x37, 0x33, 0xeb, 0xce, 0x2c, 0x05, 0x8f, 0x5e, 0xbd, 0x08, 0x5e, 0x3c, 0xf9, 0x7b,
	0x3c, 0x16, 0xbc, 0x78, 0x0a, 0x92, 0x78, 0xf2, 0x98, 0x5f, 0x50, 0x66, 0x66, 0xd3, 0x6e, 0x6f,
	0xf3, 0x7d, 0xef, 0x9b, 0xef, 0x7b, 0x6f, 0xde, 0x00, 0xd4, 0x8a, 0x56, 0xb3, 0xb2, 0x92, 0x5a,
	0x62, 0xaf, 0xcc, 0xc2, 0x43, 0x26, 0x25, 0x2b, 0x68, 0x9c, 0x96, 0x3c, 0x4e, 0x85, 0x90, 0x3a,
	0xd5, 0x5c, 0x0a, 0xe5, 0x14, 0xe1, 0x09, 0xe3, 0xfa, 0x43, 0x9d, 0xcd, 0x16, 0x72, 0x19, 0x33,
	0xc9, 0x64, 0x6c, 0xe9, 0xac, 0x7e, 0x6f, 0x91, 0x05, 0xf6, 0xe4, 0xe4, 0xe4, 0x05, 0x0c, 0xcf,
	0x25, 0xe3, 0x22, 0xa1, 0x9f, 0x6a, 0xaa, 0x34, 0x0e, 0xe1, 0xbe, 0x89, 0x13, 0xe9, 0x92, 0x06,
	0xe8, 0x08, 0x4d, 0xfb, 0xc9, 0x0d, 0x36, 0xb5, 0x32, 0x55, 0xea, 0x52, 0x56, 0x79, 0xe0, 0xb9,
	0xda, 0x0e, 0x93, 0x63, 0x18, 0x5d, 0xc8, 0x8f, 0x54, 0x24, 0x54, 0x95, 0x52, 0x28, 0x8a, 0x0f,
	0xa0, 0xa7, 0x0d, 0xd1, 0xb8, 0x38, 0x40, 0xde, 0xc2, 0xe0, 0x9c, 0x2b, 0xbd, 0x4b, 0x0b, 0x60,
	0x6f, 0x21, 0x8b, 0x7a, 0x29, 0x54, 0x80, 0x8e, 0xba, 0xd3, 0x7e, 0xb2, 0x83, 0x18, 0x83, 0x5f,
	0xa6, 0x8c, 0xda, 0x9c, 0x5e, 0x62, 0xcf, 0x2e, 0x9f, 0xd1, 0xd7, 0xfc, 0x33, 0x0d, 0xba, 0x96,
	0xbf, 0xc1, 0xe4, 0x25, 0x0c, 0x9d, 0x71, 0x13, 0x7f, 0x0c, 0x7e, 0x25, 0x2f, 0x9d, 0xed, 0xe0,
	0x74, 0x7f, 0x56, 0x66, 0xb3, 0x37, 0x8a, 0x56, 0x46, 0xf3, 0x4a, 0xe6, 0xb4, 0x48, 0x6c, 0xd9,
	0x75, 0xa9, 0xd3, 0xa2, 0xc9, 0x71, 0x80, 0xfc, 0x44, 0x30, 0xba, 0xa3, 0xc6, 0x53, 0xf0, 0xb8,
	0x1b, 0xba, 0x3b, 0x0f, 0xfe, 0xaf, 0x26, 0x7d, 0x9e, 0x3f, 0x56, 0xba, 0xe2, 0x82, 0x6d, 0x57,
	0x93, 0xbd, 0x92, 0x9d, 0x11, 0x9e, 0x93, 0xc4, 0xe3, 0x39, 0x3e, 0x69, 0x3d, 0xa0, 0x69, 0xb2,
	0x3f, 0xdf, 0xdf, 0xae, 0x26, 0x23, 0x23, 0xd9, 0xf1, 0xa4, 0xf5, 0xa6, 0x4f, 0x01, 0x16, 0x15,
	0x4d, 0x35, 0xbd, 0xe0, 0x4b, 0x1a, 0xf8, 0x36, 0xe0, 0x60, 0xbb, 0x9a, 0x8c, 0xcd, 0x05, 0x57,
	0x79, 0xa7, 0xb9, 0xb9, 0xd3, 0xd2, 0x9d, 0x7e, 0x45, 0xe0, 0x9b, 0x06, 0xf1, 0x73, 0xe8, 0xd9,
	0xf5, 0xe1, 0xb1, 0x99, 0xb0, 0xbd, 0xc9, 0xd0, 0xce, 0x7c, 0x67, 0x27, 0xe4, 0xe1, 0x97, 0xdf,
	0xff, 0xbe, 0x7b, 0x63, 0x32, 0x88, 0x4d, 0x7e, 0x5c, 0x18, 0xf9, 0x19, 0x7a, 0x84, 0x9f, 0x81,
	0x6f, 0x46, 0xc5, 0x0f, 0xac, 0xc9, 0xed, 0x7e, 0xc2, 0xf1, 0x2d, 0xd1, 0x58, 0x60, 0x6b, 0x31,
	0xc4, 0xd0, 0x58, 0x70, 0xa5, 0xe7, 0x87, 0xbf, 0xd6, 0x11, 0xba, 0x5a, 0x47, 0xe8, 0xef, 0x3a,
	0x42, 0xdf, 0x36, 0x51, 0xe7, 0xc7, 0x26, 0xea, 0x5c, 0x6d, 0xa2, 0xce, 0x9f, 0x4d, 0xd4, 0xc9,
	0xee, 0xd9, 0x8f, 0xf6, 0xe4, 0x3a, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xc4, 0x51, 0xe5, 0xc7, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	// 定义rpc方法
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/pb.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/pb.User/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	// 定义rpc方法
	Login(context.Context, *LoginRequest) (*TokenResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Login(ctx context.Context, req *LoginRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "List",
			Handler:    _User_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func (m *LoginRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PageSize != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.PageSize))
		i--
		dAtA[i] = 0x18
	}
	if m.Page != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Columns) > 0 {
		for iNdEx := len(m.Columns) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Columns[iNdEx])
			copy(dAtA[i:], m.Columns[iNdEx])
			i = encodeVarintUser(dAtA, i, uint64(len(m.Columns[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Total != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Rows) > 0 {
		for iNdEx := len(m.Rows) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rows[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUser(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UserListModel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserListModel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserListModel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreateTime != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.CreateTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LoginRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func (m *TokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func (m *ListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Columns) > 0 {
		for _, s := range m.Columns {
			l = len(s)
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if m.Page != 0 {
		n += 1 + sovUser(uint64(m.Page))
	}
	if m.PageSize != 0 {
		n += 1 + sovUser(uint64(m.PageSize))
	}
	return n
}

func (m *ListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rows) > 0 {
		for _, e := range m.Rows {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if m.Total != 0 {
		n += 1 + sovUser(uint64(m.Total))
	}
	return n
}

func (m *UserListModel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUser(uint64(m.Id))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.CreateTime != 0 {
		n += 1 + sovUser(uint64(m.CreateTime))
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: LoginRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
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
func (m *TokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: TokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
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
func (m *ListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: ListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageSize", wireType)
			}
			m.PageSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageSize |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
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
func (m *ListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: ListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rows = append(m.Rows, &UserListModel{})
			if err := m.Rows[len(m.Rows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
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
func (m *UserListModel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserListModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserListModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			m.CreateTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)