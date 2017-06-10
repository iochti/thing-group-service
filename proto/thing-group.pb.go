// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/iochti/thing-group-service/proto/thing-group.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	github.com/iochti/thing-group-service/proto/thing-group.proto

It has these top-level messages:
	ThingGroup
	GroupIDRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ThingGroup struct {
	ID          int32                       `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name        string                      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string                      `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	CreatedAt   *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt   *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *ThingGroup) Reset()                    { *m = ThingGroup{} }
func (m *ThingGroup) String() string            { return proto1.CompactTextString(m) }
func (*ThingGroup) ProtoMessage()               {}
func (*ThingGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ThingGroup) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ThingGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ThingGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ThingGroup) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ThingGroup) GetUpdatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type GroupIDRequest struct {
	ID int32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *GroupIDRequest) Reset()                    { *m = GroupIDRequest{} }
func (m *GroupIDRequest) String() string            { return proto1.CompactTextString(m) }
func (*GroupIDRequest) ProtoMessage()               {}
func (*GroupIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GroupIDRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto1.RegisterType((*ThingGroup)(nil), "proto.ThingGroup")
	proto1.RegisterType((*GroupIDRequest)(nil), "proto.GroupIDRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ThingGroupSvc service

type ThingGroupSvcClient interface {
	GetGroup(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*ThingGroup, error)
	UpdateGroup(ctx context.Context, in *ThingGroup, opts ...grpc.CallOption) (*ThingGroup, error)
	CreateGroup(ctx context.Context, in *ThingGroup, opts ...grpc.CallOption) (*ThingGroup, error)
	DeleteGroup(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type thingGroupSvcClient struct {
	cc *grpc.ClientConn
}

func NewThingGroupSvcClient(cc *grpc.ClientConn) ThingGroupSvcClient {
	return &thingGroupSvcClient{cc}
}

func (c *thingGroupSvcClient) GetGroup(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*ThingGroup, error) {
	out := new(ThingGroup)
	err := grpc.Invoke(ctx, "/proto.ThingGroupSvc/GetGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingGroupSvcClient) UpdateGroup(ctx context.Context, in *ThingGroup, opts ...grpc.CallOption) (*ThingGroup, error) {
	out := new(ThingGroup)
	err := grpc.Invoke(ctx, "/proto.ThingGroupSvc/UpdateGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingGroupSvcClient) CreateGroup(ctx context.Context, in *ThingGroup, opts ...grpc.CallOption) (*ThingGroup, error) {
	out := new(ThingGroup)
	err := grpc.Invoke(ctx, "/proto.ThingGroupSvc/CreateGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingGroupSvcClient) DeleteGroup(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/proto.ThingGroupSvc/DeleteGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ThingGroupSvc service

type ThingGroupSvcServer interface {
	GetGroup(context.Context, *GroupIDRequest) (*ThingGroup, error)
	UpdateGroup(context.Context, *ThingGroup) (*ThingGroup, error)
	CreateGroup(context.Context, *ThingGroup) (*ThingGroup, error)
	DeleteGroup(context.Context, *GroupIDRequest) (*google_protobuf.Empty, error)
}

func RegisterThingGroupSvcServer(s *grpc.Server, srv ThingGroupSvcServer) {
	s.RegisterService(&_ThingGroupSvc_serviceDesc, srv)
}

func _ThingGroupSvc_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingGroupSvcServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThingGroupSvc/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingGroupSvcServer).GetGroup(ctx, req.(*GroupIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingGroupSvc_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThingGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingGroupSvcServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThingGroupSvc/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingGroupSvcServer).UpdateGroup(ctx, req.(*ThingGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingGroupSvc_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThingGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingGroupSvcServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThingGroupSvc/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingGroupSvcServer).CreateGroup(ctx, req.(*ThingGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingGroupSvc_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingGroupSvcServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThingGroupSvc/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingGroupSvcServer).DeleteGroup(ctx, req.(*GroupIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ThingGroupSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ThingGroupSvc",
	HandlerType: (*ThingGroupSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroup",
			Handler:    _ThingGroupSvc_GetGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _ThingGroupSvc_UpdateGroup_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _ThingGroupSvc_CreateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _ThingGroupSvc_DeleteGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/iochti/thing-group-service/proto/thing-group.proto",
}

func init() {
	proto1.RegisterFile("github.com/iochti/thing-group-service/proto/thing-group.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbb, 0xfd, 0xb7, 0x7f, 0xec, 0x04, 0x0b, 0x2e, 0x28, 0x21, 0x1e, 0x0c, 0x39, 0xf5,
	0xd2, 0x0d, 0x54, 0x14, 0x3c, 0x28, 0x14, 0x23, 0xa5, 0xd7, 0x58, 0xcf, 0x92, 0x6e, 0xc7, 0x74,
	0xa1, 0xe9, 0xc6, 0x64, 0x52, 0xf0, 0x1b, 0xfa, 0x85, 0xbc, 0x4b, 0x36, 0x09, 0xad, 0x2d, 0x45,
	0x3c, 0x25, 0xcc, 0xbc, 0xdf, 0xdb, 0xf7, 0x96, 0x85, 0xfb, 0x58, 0xd1, 0xb2, 0x98, 0x0b, 0xa9,
	0x13, 0x5f, 0x69, 0xb9, 0x24, 0xe5, 0xd3, 0x52, 0xad, 0xe3, 0x61, 0x9c, 0xe9, 0x22, 0x1d, 0xe6,
	0x98, 0x6d, 0x94, 0x44, 0x3f, 0xcd, 0x34, 0xe9, 0xdd, 0x8d, 0x30, 0x13, 0xde, 0x35, 0x1f, 0xe7,
	0x32, 0xd6, 0x3a, 0x5e, 0xd5, 0xb2, 0x79, 0xf1, 0xe6, 0x63, 0x92, 0xd2, 0x47, 0xa5, 0x71, 0xae,
	0xf6, 0x97, 0xa4, 0x12, 0xcc, 0x29, 0x4a, 0x6a, 0x13, 0xef, 0x93, 0x01, 0xcc, 0x4a, 0xeb, 0x49,
	0xe9, 0xcc, 0xfb, 0xd0, 0x9e, 0x06, 0x36, 0x73, 0xd9, 0xa0, 0x1b, 0xb6, 0xa7, 0x01, 0xe7, 0xd0,
	0x59, 0x47, 0x09, 0xda, 0x6d, 0x97, 0x0d, 0x7a, 0xa1, 0xf9, 0xe7, 0x2e, 0x58, 0x0b, 0xcc, 0x65,
	0xa6, 0x52, 0x52, 0x7a, 0x6d, 0xff, 0x33, 0xab, 0xdd, 0x11, 0xbf, 0x03, 0x90, 0x19, 0x46, 0x84,
	0x8b, 0xd7, 0x88, 0xec, 0x8e, 0xcb, 0x06, 0xd6, 0xc8, 0x11, 0x55, 0x14, 0xd1, 0x44, 0x11, 0xb3,
	0x26, 0x4a, 0xd8, 0xab, 0xd5, 0x63, 0x2a, 0xd1, 0x22, 0x5d, 0x34, 0x68, 0xf7, 0x77, 0xb4, 0x56,
	0x8f, 0xc9, 0x73, 0xa1, 0x6f, 0x4a, 0x4c, 0x83, 0x10, 0xdf, 0x0b, 0xcc, 0x69, 0xbf, 0xcd, 0xe8,
	0x8b, 0xc1, 0xe9, 0xb6, 0xec, 0xf3, 0x46, 0xf2, 0x5b, 0x38, 0x99, 0x20, 0x55, 0xdd, 0xcf, 0x2b,
	0x7f, 0xf1, 0xd3, 0xc4, 0x39, 0xab, 0xc7, 0x5b, 0xd0, 0x6b, 0xf1, 0x1b, 0xb0, 0x5e, 0xcc, 0xc1,
	0x15, 0x7a, 0xa8, 0x39, 0x8a, 0x3d, 0x9a, 0xaa, 0x7f, 0xc3, 0x1e, 0xc0, 0x0a, 0x70, 0x85, 0x0d,
	0x76, 0x24, 0xe8, 0xc5, 0xc1, 0x35, 0x3d, 0x95, 0x2f, 0xc1, 0x6b, 0xcd, 0xff, 0x9b, 0xc9, 0xf5,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xca, 0x9a, 0x5c, 0x71, 0x02, 0x00, 0x00,
}
