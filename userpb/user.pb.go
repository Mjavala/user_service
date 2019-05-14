// Code generated by protoc-gen-go. DO NOT EDIT.
// source: User/userpb/user.proto

package userpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebaacd1c7fe14cd1, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
}

func init() { proto.RegisterFile("User/userpb/user.proto", fileDescriptor_ebaacd1c7fe14cd1) }

var fileDescriptor_ebaacd1c7fe14cd1 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0b, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2f, 0x2d, 0x4e, 0x2d, 0x2a, 0x48, 0x02, 0x53, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x2c, 0x20, 0xb6, 0x52, 0x0c, 0x17, 0x0b, 0x48, 0x5e, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37,
	0x55, 0x82, 0x09, 0x2c, 0x02, 0x66, 0x0b, 0x89, 0x70, 0xb1, 0xa6, 0xe6, 0x26, 0x66, 0xe6, 0x48,
	0x30, 0x83, 0x05, 0x21, 0x1c, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0xc4, 0xe2, 0xe2, 0xf2, 0xfc, 0xa2,
	0x14, 0x09, 0x16, 0xb0, 0x04, 0x9c, 0x6f, 0xc4, 0xcb, 0xc5, 0x0d, 0x32, 0x3d, 0x38, 0xb5, 0xa8,
	0x2c, 0x33, 0x39, 0xd5, 0x89, 0x23, 0x8a, 0x0d, 0xe2, 0x8e, 0x24, 0x36, 0xb0, 0x1b, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0xf3, 0x93, 0x8a, 0x9d, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "User/userpb/user.proto",
}
