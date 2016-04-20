// Code generated by protoc-gen-go.
// source: filesystem.proto
// DO NOT EDIT!

/*
Package filesystem is a generated protocol buffer package.

It is generated from these files:
	filesystem.proto

It has these top-level messages:
	ModFS
	CreateFSRequest
	CreateFSResponse
	ListFSRequest
	ListFSResponse
	ShowFSRequest
	ShowFSResponse
	DeleteFSRequest
	DeleteFSResponse
	UpdateFSRequest
	UpdateFSResponse
	GrantAddrFSRequest
	GrantAddrFSResponse
	RevokeAddrFSRequest
	RevokeAddrFSResponse
	LookupAddrFSRequest
	LookupAddrFSResponse
*/
package filesystem

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// ModFS ...
type ModFS struct {
	Name   string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=Status,json=status" json:"Status,omitempty"`
}

func (m *ModFS) Reset()                    { *m = ModFS{} }
func (m *ModFS) String() string            { return proto.CompactTextString(m) }
func (*ModFS) ProtoMessage()               {}
func (*ModFS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Request to create a new filesystem
type CreateFSRequest struct {
	Acctnum   string `protobuf:"bytes,1,opt,name=Acctnum,json=acctnum" json:"Acctnum,omitempty"`
	FSName    string `protobuf:"bytes,2,opt,name=FSName,json=fSName" json:"FSName,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=Token,json=token" json:"Token,omitempty"`
	IPAddress string `protobuf:"bytes,4,opt,name=IPAddress,json=iPAddress" json:"IPAddress,omitempty"`
}

func (m *CreateFSRequest) Reset()                    { *m = CreateFSRequest{} }
func (m *CreateFSRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFSRequest) ProtoMessage()               {}
func (*CreateFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Response from creating a new filesystem
type CreateFSResponse struct {
	Status  string `protobuf:"bytes,1,opt,name=Status,json=status" json:"Status,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=Payload,json=payload" json:"Payload,omitempty"`
}

func (m *CreateFSResponse) Reset()                    { *m = CreateFSResponse{} }
func (m *CreateFSResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateFSResponse) ProtoMessage()               {}
func (*CreateFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Request a list of all file systems for a given account
type ListFSRequest struct {
	Acctnum string `protobuf:"bytes,1,opt,name=Acctnum,json=acctnum" json:"Acctnum,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=Token,json=token" json:"Token,omitempty"`
}

func (m *ListFSRequest) Reset()                    { *m = ListFSRequest{} }
func (m *ListFSRequest) String() string            { return proto.CompactTextString(m) }
func (*ListFSRequest) ProtoMessage()               {}
func (*ListFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Response for displaying a list of all an accounts file systems.
type ListFSResponse struct {
	Status  string `protobuf:"bytes,1,opt,name=Status,json=status" json:"Status,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=Payload,json=payload" json:"Payload,omitempty"`
}

func (m *ListFSResponse) Reset()                    { *m = ListFSResponse{} }
func (m *ListFSResponse) String() string            { return proto.CompactTextString(m) }
func (*ListFSResponse) ProtoMessage()               {}
func (*ListFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Request to show the specific details about a file system
type ShowFSRequest struct {
	Acctnum string `protobuf:"bytes,1,opt,name=Acctnum,json=acctnum" json:"Acctnum,omitempty"`
	FSid    string `protobuf:"bytes,2,opt,name=FSid,json=fSid" json:"FSid,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=Token,json=token" json:"Token,omitempty"`
}

func (m *ShowFSRequest) Reset()                    { *m = ShowFSRequest{} }
func (m *ShowFSRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowFSRequest) ProtoMessage()               {}
func (*ShowFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Response for a specific file system for an account.
type ShowFSResponse struct {
	Status  string `protobuf:"bytes,1,opt,name=Status,json=status" json:"Status,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=Payload,json=payload" json:"Payload,omitempty"`
}

func (m *ShowFSResponse) Reset()                    { *m = ShowFSResponse{} }
func (m *ShowFSResponse) String() string            { return proto.CompactTextString(m) }
func (*ShowFSResponse) ProtoMessage()               {}
func (*ShowFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Request to delete a specific file system
type DeleteFSRequest struct {
	Acctnum string `protobuf:"bytes,1,opt,name=Acctnum,json=acctnum" json:"Acctnum,omitempty"`
	FSid    string `protobuf:"bytes,2,opt,name=FSid,json=fSid" json:"FSid,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=Token,json=token" json:"Token,omitempty"`
}

func (m *DeleteFSRequest) Reset()                    { *m = DeleteFSRequest{} }
func (m *DeleteFSRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteFSRequest) ProtoMessage()               {}
func (*DeleteFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// Response from deleting a file system
type DeleteFSResponse struct {
	Status  string `protobuf:"bytes,1,opt,name=Status,json=status" json:"Status,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=Payload,json=payload" json:"Payload,omitempty"`
}

func (m *DeleteFSResponse) Reset()                    { *m = DeleteFSResponse{} }
func (m *DeleteFSResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteFSResponse) ProtoMessage()               {}
func (*DeleteFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// Request to update a specific file system's information
type UpdateFSRequest struct {
	Acctnum string `protobuf:"bytes,1,opt,name=Acctnum,json=acctnum" json:"Acctnum,omitempty"`
	FSid    string `protobuf:"bytes,2,opt,name=FSid,json=fSid" json:"FSid,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=Token,json=token" json:"Token,omitempty"`
	Filesys *ModFS `protobuf:"bytes,4,opt,name=Filesys,json=filesys" json:"Filesys,omitempty"`
}

func (m *UpdateFSRequest) Reset()                    { *m = UpdateFSRequest{} }
func (m *UpdateFSRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateFSRequest) ProtoMessage()               {}
func (*UpdateFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateFSRequest) GetFilesys() *ModFS {
	if m != nil {
		return m.Filesys
	}
	return nil
}

// Response from an update operation
type UpdateFSResponse struct {
	Status  string `protobuf:"bytes,1,opt,name=Status,json=status" json:"Status,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=Payload,json=payload" json:"Payload,omitempty"`
}

func (m *UpdateFSResponse) Reset()                    { *m = UpdateFSResponse{} }
func (m *UpdateFSResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateFSResponse) ProtoMessage()               {}
func (*UpdateFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// Request grant an ip address access to a file system
type GrantAddrFSRequest struct {
	Acctnum string `protobuf:"bytes,1,opt,name=Acctnum,json=acctnum" json:"Acctnum,omitempty"`
	FSid    string `protobuf:"bytes,2,opt,name=FSid,json=fSid" json:"FSid,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=Token,json=token" json:"Token,omitempty"`
	Addr    string `protobuf:"bytes,4,opt,name=Addr,json=addr" json:"Addr,omitempty"`
}

func (m *GrantAddrFSRequest) Reset()                    { *m = GrantAddrFSRequest{} }
func (m *GrantAddrFSRequest) String() string            { return proto.CompactTextString(m) }
func (*GrantAddrFSRequest) ProtoMessage()               {}
func (*GrantAddrFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// Response from granting ip address access to a file system
type GrantAddrFSResponse struct {
	Status string `protobuf:"bytes,1,opt,name=Status,json=status" json:"Status,omitempty"`
}

func (m *GrantAddrFSResponse) Reset()                    { *m = GrantAddrFSResponse{} }
func (m *GrantAddrFSResponse) String() string            { return proto.CompactTextString(m) }
func (*GrantAddrFSResponse) ProtoMessage()               {}
func (*GrantAddrFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// Request revoke an ip address access to a file system
type RevokeAddrFSRequest struct {
	Acctnum string `protobuf:"bytes,1,opt,name=Acctnum,json=acctnum" json:"Acctnum,omitempty"`
	FSid    string `protobuf:"bytes,2,opt,name=FSid,json=fSid" json:"FSid,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=Token,json=token" json:"Token,omitempty"`
	Addr    string `protobuf:"bytes,4,opt,name=Addr,json=addr" json:"Addr,omitempty"`
}

func (m *RevokeAddrFSRequest) Reset()                    { *m = RevokeAddrFSRequest{} }
func (m *RevokeAddrFSRequest) String() string            { return proto.CompactTextString(m) }
func (*RevokeAddrFSRequest) ProtoMessage()               {}
func (*RevokeAddrFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// Response from revoking ip address access to a file system
type RevokeAddrFSResponse struct {
	Status string `protobuf:"bytes,1,opt,name=Status,json=status" json:"Status,omitempty"`
}

func (m *RevokeAddrFSResponse) Reset()                    { *m = RevokeAddrFSResponse{} }
func (m *RevokeAddrFSResponse) String() string            { return proto.CompactTextString(m) }
func (*RevokeAddrFSResponse) ProtoMessage()               {}
func (*RevokeAddrFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

// Request filesystem lookup
type LookupAddrFSRequest struct {
	FSid string `protobuf:"bytes,1,opt,name=FSid,json=fSid" json:"FSid,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=Addr,json=addr" json:"Addr,omitempty"`
}

func (m *LookupAddrFSRequest) Reset()                    { *m = LookupAddrFSRequest{} }
func (m *LookupAddrFSRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupAddrFSRequest) ProtoMessage()               {}
func (*LookupAddrFSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type LookupAddrFSResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *LookupAddrFSResponse) Reset()                    { *m = LookupAddrFSResponse{} }
func (m *LookupAddrFSResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupAddrFSResponse) ProtoMessage()               {}
func (*LookupAddrFSResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func init() {
	proto.RegisterType((*ModFS)(nil), "filesystem.ModFS")
	proto.RegisterType((*CreateFSRequest)(nil), "filesystem.CreateFSRequest")
	proto.RegisterType((*CreateFSResponse)(nil), "filesystem.CreateFSResponse")
	proto.RegisterType((*ListFSRequest)(nil), "filesystem.ListFSRequest")
	proto.RegisterType((*ListFSResponse)(nil), "filesystem.ListFSResponse")
	proto.RegisterType((*ShowFSRequest)(nil), "filesystem.ShowFSRequest")
	proto.RegisterType((*ShowFSResponse)(nil), "filesystem.ShowFSResponse")
	proto.RegisterType((*DeleteFSRequest)(nil), "filesystem.DeleteFSRequest")
	proto.RegisterType((*DeleteFSResponse)(nil), "filesystem.DeleteFSResponse")
	proto.RegisterType((*UpdateFSRequest)(nil), "filesystem.UpdateFSRequest")
	proto.RegisterType((*UpdateFSResponse)(nil), "filesystem.UpdateFSResponse")
	proto.RegisterType((*GrantAddrFSRequest)(nil), "filesystem.GrantAddrFSRequest")
	proto.RegisterType((*GrantAddrFSResponse)(nil), "filesystem.GrantAddrFSResponse")
	proto.RegisterType((*RevokeAddrFSRequest)(nil), "filesystem.RevokeAddrFSRequest")
	proto.RegisterType((*RevokeAddrFSResponse)(nil), "filesystem.RevokeAddrFSResponse")
	proto.RegisterType((*LookupAddrFSRequest)(nil), "filesystem.LookupAddrFSRequest")
	proto.RegisterType((*LookupAddrFSResponse)(nil), "filesystem.LookupAddrFSResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for FileSystemAPI service

type FileSystemAPIClient interface {
	CreateFS(ctx context.Context, in *CreateFSRequest, opts ...grpc.CallOption) (*CreateFSResponse, error)
	ListFS(ctx context.Context, in *ListFSRequest, opts ...grpc.CallOption) (*ListFSResponse, error)
	ShowFS(ctx context.Context, in *ShowFSRequest, opts ...grpc.CallOption) (*ShowFSResponse, error)
	DeleteFS(ctx context.Context, in *DeleteFSRequest, opts ...grpc.CallOption) (*DeleteFSResponse, error)
	UpdateFS(ctx context.Context, in *UpdateFSRequest, opts ...grpc.CallOption) (*UpdateFSResponse, error)
	GrantAddrFS(ctx context.Context, in *GrantAddrFSRequest, opts ...grpc.CallOption) (*GrantAddrFSResponse, error)
	RevokeAddrFS(ctx context.Context, in *RevokeAddrFSRequest, opts ...grpc.CallOption) (*RevokeAddrFSResponse, error)
	LookupAddrFS(ctx context.Context, in *LookupAddrFSRequest, opts ...grpc.CallOption) (*LookupAddrFSResponse, error)
}

type fileSystemAPIClient struct {
	cc *grpc.ClientConn
}

func NewFileSystemAPIClient(cc *grpc.ClientConn) FileSystemAPIClient {
	return &fileSystemAPIClient{cc}
}

func (c *fileSystemAPIClient) CreateFS(ctx context.Context, in *CreateFSRequest, opts ...grpc.CallOption) (*CreateFSResponse, error) {
	out := new(CreateFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemAPI/CreateFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemAPIClient) ListFS(ctx context.Context, in *ListFSRequest, opts ...grpc.CallOption) (*ListFSResponse, error) {
	out := new(ListFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemAPI/ListFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemAPIClient) ShowFS(ctx context.Context, in *ShowFSRequest, opts ...grpc.CallOption) (*ShowFSResponse, error) {
	out := new(ShowFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemAPI/ShowFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemAPIClient) DeleteFS(ctx context.Context, in *DeleteFSRequest, opts ...grpc.CallOption) (*DeleteFSResponse, error) {
	out := new(DeleteFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemAPI/DeleteFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemAPIClient) UpdateFS(ctx context.Context, in *UpdateFSRequest, opts ...grpc.CallOption) (*UpdateFSResponse, error) {
	out := new(UpdateFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemAPI/UpdateFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemAPIClient) GrantAddrFS(ctx context.Context, in *GrantAddrFSRequest, opts ...grpc.CallOption) (*GrantAddrFSResponse, error) {
	out := new(GrantAddrFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemAPI/GrantAddrFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemAPIClient) RevokeAddrFS(ctx context.Context, in *RevokeAddrFSRequest, opts ...grpc.CallOption) (*RevokeAddrFSResponse, error) {
	out := new(RevokeAddrFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemAPI/RevokeAddrFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemAPIClient) LookupAddrFS(ctx context.Context, in *LookupAddrFSRequest, opts ...grpc.CallOption) (*LookupAddrFSResponse, error) {
	out := new(LookupAddrFSResponse)
	err := grpc.Invoke(ctx, "/filesystem.FileSystemAPI/LookupAddrFS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileSystemAPI service

type FileSystemAPIServer interface {
	CreateFS(context.Context, *CreateFSRequest) (*CreateFSResponse, error)
	ListFS(context.Context, *ListFSRequest) (*ListFSResponse, error)
	ShowFS(context.Context, *ShowFSRequest) (*ShowFSResponse, error)
	DeleteFS(context.Context, *DeleteFSRequest) (*DeleteFSResponse, error)
	UpdateFS(context.Context, *UpdateFSRequest) (*UpdateFSResponse, error)
	GrantAddrFS(context.Context, *GrantAddrFSRequest) (*GrantAddrFSResponse, error)
	RevokeAddrFS(context.Context, *RevokeAddrFSRequest) (*RevokeAddrFSResponse, error)
	LookupAddrFS(context.Context, *LookupAddrFSRequest) (*LookupAddrFSResponse, error)
}

func RegisterFileSystemAPIServer(s *grpc.Server, srv FileSystemAPIServer) {
	s.RegisterService(&_FileSystemAPI_serviceDesc, srv)
}

func _FileSystemAPI_CreateFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemAPIServer).CreateFS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.FileSystemAPI/CreateFS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemAPIServer).CreateFS(ctx, req.(*CreateFSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystemAPI_ListFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemAPIServer).ListFS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.FileSystemAPI/ListFS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemAPIServer).ListFS(ctx, req.(*ListFSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystemAPI_ShowFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemAPIServer).ShowFS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.FileSystemAPI/ShowFS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemAPIServer).ShowFS(ctx, req.(*ShowFSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystemAPI_DeleteFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemAPIServer).DeleteFS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.FileSystemAPI/DeleteFS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemAPIServer).DeleteFS(ctx, req.(*DeleteFSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystemAPI_UpdateFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemAPIServer).UpdateFS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.FileSystemAPI/UpdateFS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemAPIServer).UpdateFS(ctx, req.(*UpdateFSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystemAPI_GrantAddrFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantAddrFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemAPIServer).GrantAddrFS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.FileSystemAPI/GrantAddrFS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemAPIServer).GrantAddrFS(ctx, req.(*GrantAddrFSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystemAPI_RevokeAddrFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeAddrFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemAPIServer).RevokeAddrFS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.FileSystemAPI/RevokeAddrFS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemAPIServer).RevokeAddrFS(ctx, req.(*RevokeAddrFSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystemAPI_LookupAddrFS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupAddrFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemAPIServer).LookupAddrFS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.FileSystemAPI/LookupAddrFS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemAPIServer).LookupAddrFS(ctx, req.(*LookupAddrFSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileSystemAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "filesystem.FileSystemAPI",
	HandlerType: (*FileSystemAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFS",
			Handler:    _FileSystemAPI_CreateFS_Handler,
		},
		{
			MethodName: "ListFS",
			Handler:    _FileSystemAPI_ListFS_Handler,
		},
		{
			MethodName: "ShowFS",
			Handler:    _FileSystemAPI_ShowFS_Handler,
		},
		{
			MethodName: "DeleteFS",
			Handler:    _FileSystemAPI_DeleteFS_Handler,
		},
		{
			MethodName: "UpdateFS",
			Handler:    _FileSystemAPI_UpdateFS_Handler,
		},
		{
			MethodName: "GrantAddrFS",
			Handler:    _FileSystemAPI_GrantAddrFS_Handler,
		},
		{
			MethodName: "RevokeAddrFS",
			Handler:    _FileSystemAPI_RevokeAddrFS_Handler,
		},
		{
			MethodName: "LookupAddrFS",
			Handler:    _FileSystemAPI_LookupAddrFS_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x95, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0x92, 0x36, 0xf4, 0x8c, 0xd2, 0xe1, 0x4e, 0x28, 0x84, 0x09, 0x26, 0x5f, 0x21,
	0x21, 0x7a, 0xb1, 0x5d, 0x23, 0x54, 0x98, 0x32, 0x4d, 0x1a, 0xa8, 0xaa, 0xd9, 0x03, 0x98, 0xc5,
	0x15, 0xd5, 0xda, 0x38, 0xd4, 0x2e, 0xd3, 0x6e, 0x79, 0x5e, 0x1e, 0x02, 0xc7, 0x89, 0x17, 0x3b,
	0x64, 0x50, 0x45, 0xd3, 0xee, 0xe2, 0x3f, 0xf9, 0x9d, 0xef, 0x34, 0xdf, 0x77, 0x0a, 0x7b, 0xf3,
	0xc5, 0x92, 0x89, 0x1b, 0x21, 0xd9, 0x6a, 0x9c, 0xad, 0xb9, 0xe4, 0x08, 0xaa, 0x1d, 0x7c, 0x0c,
	0xdd, 0xcf, 0x3c, 0x89, 0x09, 0x42, 0xe0, 0x7f, 0xa1, 0x2b, 0x16, 0x7a, 0x87, 0xde, 0x9b, 0xfe,
	0xcc, 0x4f, 0xd5, 0x33, 0x7a, 0x0e, 0x3d, 0x22, 0xa9, 0xdc, 0x88, 0xb0, 0xa3, 0x77, 0x7b, 0x42,
	0xaf, 0xf0, 0x35, 0x0c, 0x3f, 0xad, 0x19, 0x95, 0x2c, 0x26, 0x33, 0xf6, 0x63, 0xc3, 0x84, 0x44,
	0x21, 0x04, 0x93, 0xcb, 0x4b, 0x99, 0x6e, 0x56, 0x25, 0x21, 0xa0, 0xc5, 0x32, 0x87, 0xc4, 0x44,
	0xa3, 0x4b, 0xc8, 0x5c, 0xaf, 0xd0, 0x3e, 0x74, 0xbf, 0xf2, 0x2b, 0x96, 0x86, 0x8f, 0xf4, 0x76,
	0x57, 0xe6, 0x0b, 0x74, 0x00, 0xfd, 0xb3, 0xe9, 0x24, 0x49, 0xd6, 0x4c, 0x88, 0xd0, 0xd7, 0x27,
	0xfd, 0x85, 0xd9, 0xc0, 0x27, 0xb0, 0x57, 0x15, 0x16, 0x19, 0x4f, 0x85, 0x2d, 0xd2, 0xb3, 0x45,
	0xe6, 0x8a, 0xa6, 0xf4, 0x66, 0xc9, 0x69, 0x52, 0x16, 0x0e, 0xb2, 0x62, 0x89, 0x3f, 0xc0, 0xe0,
	0x7c, 0x21, 0xe4, 0x36, 0xe2, 0x6f, 0x45, 0x76, 0x2c, 0x91, 0xf8, 0x23, 0x3c, 0x35, 0x80, 0xd6,
	0x22, 0x08, 0x0c, 0xc8, 0x77, 0x7e, 0xbd, 0x8d, 0x08, 0xf5, 0x69, 0x62, 0xb2, 0x30, 0x04, 0x7f,
	0xae, 0x9e, 0x9b, 0x7f, 0xbd, 0x5c, 0x98, 0x81, 0xb6, 0x16, 0x76, 0x01, 0xc3, 0x13, 0xb6, 0x64,
	0xdb, 0x7d, 0xdc, 0xed, 0xa5, 0xa9, 0x4f, 0x57, 0x61, 0x5b, 0x8b, 0xfb, 0xe5, 0xc1, 0xf0, 0x22,
	0x4b, 0xe8, 0x3d, 0xab, 0x43, 0x6f, 0x21, 0x88, 0x8b, 0x50, 0x68, 0xd3, 0xed, 0x1e, 0x3d, 0x1b,
	0x5b, 0xb1, 0xd1, 0x09, 0x99, 0x05, 0xe5, 0x4e, 0xde, 0x4a, 0xa5, 0xa1, 0x75, 0x2b, 0x4b, 0x40,
	0xa7, 0x6b, 0x9a, 0xca, 0xdc, 0xdb, 0xf7, 0xdb, 0x8c, 0xba, 0x99, 0x43, 0xcb, 0xf8, 0xf8, 0x54,
	0x3d, 0xe3, 0x77, 0x30, 0x72, 0xaa, 0xfd, 0x5b, 0x36, 0x5e, 0xc1, 0x68, 0xc6, 0x7e, 0x2a, 0xda,
	0xc3, 0xa8, 0x1b, 0xc3, 0xbe, 0x5b, 0xee, 0x3f, 0xf2, 0xde, 0xc3, 0xe8, 0x9c, 0xf3, 0xab, 0x4d,
	0xe6, 0xca, 0x33, 0x22, 0x3c, 0x4b, 0x84, 0x29, 0xd7, 0x71, 0xcb, 0xb9, 0xaf, 0x57, 0xe5, 0x44,
	0x43, 0xb9, 0xa3, 0xdf, 0x3e, 0x0c, 0x72, 0x7b, 0x10, 0x6d, 0x87, 0xc9, 0xf4, 0x0c, 0x9d, 0xc2,
	0x63, 0x33, 0x88, 0xd0, 0x4b, 0xdb, 0x2a, 0xb5, 0xb9, 0x18, 0x1d, 0x34, 0x1f, 0x16, 0x05, 0xf1,
	0x0e, 0x9a, 0x40, 0xaf, 0x18, 0x25, 0xe8, 0x85, 0x7d, 0xd3, 0x99, 0x4f, 0x51, 0xd4, 0x74, 0x64,
	0x23, 0x8a, 0xd0, 0xbb, 0x08, 0x67, 0xba, 0xb8, 0x08, 0x77, 0x46, 0x28, 0x84, 0x6a, 0xc7, 0x84,
	0xd3, 0x6d, 0xa7, 0x36, 0x09, 0xdc, 0x76, 0xea, 0x79, 0x2e, 0x40, 0x26, 0x1a, 0x2e, 0xa8, 0x16,
	0x5a, 0x17, 0x54, 0x4f, 0x93, 0x02, 0x4d, 0x61, 0xd7, 0xf2, 0x2b, 0x7a, 0x65, 0x5f, 0xff, 0x3b,
	0x36, 0xd1, 0xeb, 0x3b, 0xcf, 0x6f, 0x89, 0x04, 0x9e, 0xd8, 0x1e, 0x43, 0xce, 0x2b, 0x0d, 0x66,
	0x8f, 0x0e, 0xef, 0xbe, 0x60, 0x43, 0x6d, 0x27, 0xb9, 0xd0, 0x06, 0x8b, 0xba, 0xd0, 0x26, 0x13,
	0xe2, 0x9d, 0x6f, 0x3d, 0xfd, 0x37, 0x7d, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x50, 0x44,
	0x8c, 0xba, 0x07, 0x00, 0x00,
}
