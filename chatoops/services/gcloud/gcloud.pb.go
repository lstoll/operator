// Code generated by protoc-gen-go.
// source: gcloud/gcloud.proto
// DO NOT EDIT!

/*
Package gcloud is a generated protocol buffer package.

It is generated from these files:
	gcloud/gcloud.proto

It has these top-level messages:
	CreateDevInstanceRequest
	CreateDevInstanceResponse
	Instance
	ListInstancesRequest
	ListInstancesResponse
*/
package gcloud

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import operator "github.com/sr/operator"

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

type CreateDevInstanceRequest struct {
	Source *operator.Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
}

func (m *CreateDevInstanceRequest) Reset()                    { *m = CreateDevInstanceRequest{} }
func (m *CreateDevInstanceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDevInstanceRequest) ProtoMessage()               {}
func (*CreateDevInstanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateDevInstanceRequest) GetSource() *operator.Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type CreateDevInstanceResponse struct {
	Output *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *CreateDevInstanceResponse) Reset()                    { *m = CreateDevInstanceResponse{} }
func (m *CreateDevInstanceResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDevInstanceResponse) ProtoMessage()               {}
func (*CreateDevInstanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateDevInstanceResponse) GetOutput() *operator.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

type Instance struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Zone   string `protobuf:"bytes,4,opt,name=zone" json:"zone,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListInstancesRequest struct {
	Source    *operator.Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	ProjectId string           `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
}

func (m *ListInstancesRequest) Reset()                    { *m = ListInstancesRequest{} }
func (m *ListInstancesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListInstancesRequest) ProtoMessage()               {}
func (*ListInstancesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListInstancesRequest) GetSource() *operator.Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type ListInstancesResponse struct {
	Output  *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
	Source  *operator.Source `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Objects []*Instance      `protobuf:"bytes,3,rep,name=objects" json:"objects,omitempty"`
}

func (m *ListInstancesResponse) Reset()                    { *m = ListInstancesResponse{} }
func (m *ListInstancesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListInstancesResponse) ProtoMessage()               {}
func (*ListInstancesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListInstancesResponse) GetOutput() *operator.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *ListInstancesResponse) GetSource() *operator.Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *ListInstancesResponse) GetObjects() []*Instance {
	if m != nil {
		return m.Objects
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDevInstanceRequest)(nil), "gcloud.CreateDevInstanceRequest")
	proto.RegisterType((*CreateDevInstanceResponse)(nil), "gcloud.CreateDevInstanceResponse")
	proto.RegisterType((*Instance)(nil), "gcloud.Instance")
	proto.RegisterType((*ListInstancesRequest)(nil), "gcloud.ListInstancesRequest")
	proto.RegisterType((*ListInstancesResponse)(nil), "gcloud.ListInstancesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for GcloudService service

type GcloudServiceClient interface {
	CreateDevInstance(ctx context.Context, in *CreateDevInstanceRequest, opts ...grpc.CallOption) (*CreateDevInstanceResponse, error)
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
}

type gcloudServiceClient struct {
	cc *grpc.ClientConn
}

func NewGcloudServiceClient(cc *grpc.ClientConn) GcloudServiceClient {
	return &gcloudServiceClient{cc}
}

func (c *gcloudServiceClient) CreateDevInstance(ctx context.Context, in *CreateDevInstanceRequest, opts ...grpc.CallOption) (*CreateDevInstanceResponse, error) {
	out := new(CreateDevInstanceResponse)
	err := grpc.Invoke(ctx, "/gcloud.GcloudService/CreateDevInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcloudServiceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := grpc.Invoke(ctx, "/gcloud.GcloudService/ListInstances", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GcloudService service

type GcloudServiceServer interface {
	CreateDevInstance(context.Context, *CreateDevInstanceRequest) (*CreateDevInstanceResponse, error)
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
}

func RegisterGcloudServiceServer(s *grpc.Server, srv GcloudServiceServer) {
	s.RegisterService(&_GcloudService_serviceDesc, srv)
}

func _GcloudService_CreateDevInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateDevInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GcloudServiceServer).CreateDevInstance(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GcloudService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GcloudServiceServer).ListInstances(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _GcloudService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gcloud.GcloudService",
	HandlerType: (*GcloudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDevInstance",
			Handler:    _GcloudService_CreateDevInstance_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _GcloudService_ListInstances_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0x6d, 0x89, 0x76, 0xa4, 0x45, 0xd7, 0x1f, 0xd2, 0x60, 0xa1, 0xe6, 0x54, 0x3c, 0x54,
	0xa8, 0x8f, 0x60, 0x45, 0x0a, 0x05, 0x21, 0xbd, 0x88, 0x97, 0x92, 0x26, 0x83, 0x44, 0x34, 0x1b,
	0x77, 0x67, 0x7b, 0xf0, 0xe8, 0x23, 0x78, 0xf1, 0x81, 0x7c, 0x31, 0x93, 0xdd, 0x4d, 0x51, 0xdb,
	0x2a, 0x3d, 0xed, 0xec, 0x37, 0x33, 0xdf, 0x7e, 0xf3, 0xed, 0xc0, 0xe1, 0x43, 0xfc, 0xc4, 0x55,
	0x72, 0x61, 0x8e, 0x41, 0x2e, 0x38, 0x71, 0xe6, 0x9a, 0x9b, 0xdf, 0xe6, 0x39, 0x8a, 0x88, 0xb8,
	0x30, 0x78, 0x30, 0x02, 0xef, 0x4a, 0x60, 0x44, 0x38, 0xc2, 0xc5, 0x38, 0x93, 0x14, 0x65, 0x31,
	0x86, 0xf8, 0xa2, 0x50, 0x12, 0xeb, 0x83, 0x2b, 0xb9, 0x12, 0x31, 0x7a, 0x4e, 0xcf, 0xe9, 0xef,
	0x0d, 0xf7, 0x07, 0xcb, 0xe6, 0xa9, 0xc6, 0x43, 0x9b, 0x0f, 0xae, 0xa1, 0xb3, 0x86, 0x45, 0xe6,
	0x3c, 0x93, 0x58, 0xd2, 0x70, 0x45, 0xb9, 0xa2, 0x55, 0x9a, 0x5b, 0x8d, 0x87, 0x36, 0x1f, 0xdc,
	0xc3, 0x6e, 0xd5, 0xcd, 0xda, 0x50, 0x4b, 0x13, 0xdd, 0xd1, 0x0c, 0x8b, 0x88, 0x31, 0x68, 0x64,
	0xd1, 0x33, 0x7a, 0x35, 0x8d, 0xe8, 0x98, 0x9d, 0x14, 0x02, 0x29, 0x22, 0x25, 0xbd, 0xba, 0x46,
	0xed, 0xad, 0xac, 0x7d, 0xe5, 0x19, 0x7a, 0x0d, 0x53, 0x5b, 0xc6, 0xc1, 0x0c, 0x8e, 0x26, 0xa9,
	0xa4, 0x8a, 0x5f, 0x6e, 0x3d, 0x24, 0xeb, 0x02, 0x14, 0x9e, 0x3d, 0x62, 0x4c, 0xb3, 0x42, 0x99,
	0xd1, 0xd1, 0xb4, 0xc8, 0x38, 0x09, 0x3e, 0x1c, 0x38, 0xfe, 0xf5, 0xc2, 0xb6, 0x06, 0x7c, 0x13,
	0x53, 0xfb, 0x47, 0xcc, 0x39, 0xec, 0xf0, 0x79, 0xf9, 0x72, 0x39, 0x7b, 0x5d, 0x97, 0xda, 0xff,
	0x5e, 0xfa, 0x5f, 0x15, 0x0c, 0x3f, 0x1d, 0x68, 0xdd, 0xe8, 0xe4, 0x14, 0xc5, 0x22, 0x2d, 0xba,
	0xef, 0xe0, 0x60, 0xe5, 0xbf, 0x58, 0xaf, 0x62, 0xd8, 0xb4, 0x10, 0xfe, 0xd9, 0x1f, 0x15, 0x76,
	0xd6, 0x09, 0xb4, 0x7e, 0x98, 0xc0, 0x4e, 0xab, 0x9e, 0x75, 0xee, 0xfb, 0xdd, 0x0d, 0x59, 0xc3,
	0xe6, 0xc3, 0xfb, 0x5b, 0xc7, 0x6e, 0xee, 0xdc, 0xd5, 0x0b, 0x7b, 0xf9, 0x15, 0x00, 0x00, 0xff,
	0xff, 0xfd, 0x8a, 0x9f, 0x24, 0xdf, 0x02, 0x00, 0x00,
}
