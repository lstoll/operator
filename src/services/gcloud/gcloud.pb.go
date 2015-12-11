// Code generated by protoc-gen-go.
// source: services/gcloud/gcloud.proto
// DO NOT EDIT!

/*
Package gcloud is a generated protocol buffer package.

It is generated from these files:
	services/gcloud/gcloud.proto

It has these top-level messages:
	CreateContainerClusterRequest
	CreateContainerClusterResponse
	Instance
	ListInstancesRequest
	ListInstancesResponse
*/
package gcloud

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import operator "github.com/sr/operator/src/operator"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateContainerClusterRequest struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id" json:"project_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	NodeCount string `protobuf:"bytes,3,opt,name=node_count" json:"node_count,omitempty"`
	Zone      string `protobuf:"bytes,4,opt,name=zone" json:"zone,omitempty"`
}

func (m *CreateContainerClusterRequest) Reset()                    { *m = CreateContainerClusterRequest{} }
func (m *CreateContainerClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateContainerClusterRequest) ProtoMessage()               {}
func (*CreateContainerClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateContainerClusterResponse struct {
}

func (m *CreateContainerClusterResponse) Reset()                    { *m = CreateContainerClusterResponse{} }
func (m *CreateContainerClusterResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateContainerClusterResponse) ProtoMessage()               {}
func (*CreateContainerClusterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
	ProjectId string `protobuf:"bytes,1,opt,name=project_id" json:"project_id,omitempty"`
}

func (m *ListInstancesRequest) Reset()                    { *m = ListInstancesRequest{} }
func (m *ListInstancesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListInstancesRequest) ProtoMessage()               {}
func (*ListInstancesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListInstancesResponse struct {
	Output  *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
	Objects []*Instance      `protobuf:"bytes,2,rep,name=objects" json:"objects,omitempty"`
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

func (m *ListInstancesResponse) GetObjects() []*Instance {
	if m != nil {
		return m.Objects
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateContainerClusterRequest)(nil), "gcloud.CreateContainerClusterRequest")
	proto.RegisterType((*CreateContainerClusterResponse)(nil), "gcloud.CreateContainerClusterResponse")
	proto.RegisterType((*Instance)(nil), "gcloud.Instance")
	proto.RegisterType((*ListInstancesRequest)(nil), "gcloud.ListInstancesRequest")
	proto.RegisterType((*ListInstancesResponse)(nil), "gcloud.ListInstancesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for GCloudService service

type GCloudServiceClient interface {
	CreateContainerCluster(ctx context.Context, in *CreateContainerClusterRequest, opts ...grpc.CallOption) (*CreateContainerClusterResponse, error)
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
}

type gCloudServiceClient struct {
	cc *grpc.ClientConn
}

func NewGCloudServiceClient(cc *grpc.ClientConn) GCloudServiceClient {
	return &gCloudServiceClient{cc}
}

func (c *gCloudServiceClient) CreateContainerCluster(ctx context.Context, in *CreateContainerClusterRequest, opts ...grpc.CallOption) (*CreateContainerClusterResponse, error) {
	out := new(CreateContainerClusterResponse)
	err := grpc.Invoke(ctx, "/gcloud.GCloudService/CreateContainerCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gCloudServiceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := grpc.Invoke(ctx, "/gcloud.GCloudService/ListInstances", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GCloudService service

type GCloudServiceServer interface {
	CreateContainerCluster(context.Context, *CreateContainerClusterRequest) (*CreateContainerClusterResponse, error)
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
}

func RegisterGCloudServiceServer(s *grpc.Server, srv GCloudServiceServer) {
	s.RegisterService(&_GCloudService_serviceDesc, srv)
}

func _GCloudService_CreateContainerCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateContainerClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GCloudServiceServer).CreateContainerCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GCloudService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GCloudServiceServer).ListInstances(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _GCloudService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gcloud.GCloudService",
	HandlerType: (*GCloudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContainerCluster",
			Handler:    _GCloudService_CreateContainerCluster_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _GCloudService_ListInstances_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xdb, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xe9, 0x81, 0xa8, 0xa3, 0x15, 0x59, 0x3c, 0x84, 0xd0, 0x4a, 0x0c, 0x28, 0xe2, 0x45,
	0x0b, 0xf5, 0x0d, 0xcc, 0x85, 0x08, 0x05, 0x41, 0x6f, 0x85, 0x92, 0x6e, 0x86, 0x12, 0xa9, 0x3b,
	0x71, 0x77, 0xd6, 0x0b, 0x1f, 0xd0, 0xe7, 0x32, 0xa7, 0x15, 0x94, 0xa4, 0x5e, 0x25, 0x33, 0xf9,
	0xe6, 0x9f, 0x99, 0x7f, 0x02, 0x63, 0x83, 0xfa, 0x23, 0x93, 0x68, 0x66, 0x6b, 0xb9, 0x21, 0x9b,
	0x36, 0x8f, 0x69, 0xae, 0x89, 0x49, 0x78, 0x75, 0x14, 0x9c, 0x51, 0x8e, 0x3a, 0x61, 0xd2, 0x33,
	0xf7, 0x52, 0x03, 0x91, 0x84, 0x49, 0xac, 0x31, 0x61, 0x8c, 0x49, 0x71, 0x92, 0x29, 0xd4, 0xf1,
	0xc6, 0x1a, 0x46, 0xfd, 0x84, 0xef, 0x16, 0x0d, 0x0b, 0x01, 0x50, 0x90, 0xaf, 0x28, 0x79, 0x99,
	0xa5, 0x7e, 0x2f, 0xec, 0x5d, 0xef, 0x89, 0x03, 0x18, 0xaa, 0xe4, 0x0d, 0xfd, 0x7e, 0x15, 0x15,
	0x84, 0xa2, 0x14, 0x97, 0x92, 0xac, 0x62, 0x7f, 0xe0, 0x88, 0x4f, 0x52, 0xe8, 0x0f, 0xcb, 0x28,
	0x0a, 0xe1, 0xbc, 0xab, 0x89, 0xc9, 0x49, 0x19, 0x8c, 0xee, 0x60, 0xf7, 0x41, 0x19, 0x4e, 0x94,
	0xc4, 0x42, 0xae, 0xdf, 0xd1, 0xe9, 0x10, 0xbc, 0x82, 0x61, 0x6b, 0x5a, 0xbb, 0xdc, 0xc0, 0xf1,
	0x22, 0x33, 0xec, 0x74, 0xcc, 0x96, 0x0d, 0xa2, 0x17, 0x38, 0xf9, 0xc3, 0xd6, 0x83, 0x88, 0x10,
	0x3c, 0xb2, 0x9c, 0x5b, 0xae, 0xc0, 0xfd, 0xf9, 0xd1, 0xf4, 0xc7, 0xb0, 0xc7, 0x2a, 0x2f, 0x2e,
	0x60, 0x87, 0x56, 0xa5, 0x9a, 0x29, 0xa6, 0x1a, 0x54, 0x48, 0x63, 0xb9, 0x53, 0x9b, 0x7f, 0xf5,
	0x60, 0x74, 0x1f, 0x97, 0xb9, 0xe7, 0xfa, 0x38, 0x62, 0x0d, 0xa7, 0xed, 0x0e, 0x88, 0x4b, 0x57,
	0xbd, 0xf5, 0x0c, 0xc1, 0xd5, 0x7f, 0x58, 0x33, 0xff, 0x02, 0x46, 0xbf, 0x16, 0x13, 0x63, 0x57,
	0xd8, 0xe6, 0x4d, 0x30, 0xe9, 0xf8, 0x5a, 0xab, 0xad, 0xbc, 0xea, 0x27, 0xb9, 0xfd, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x9f, 0x72, 0xa7, 0xbb, 0x65, 0x02, 0x00, 0x00,
}
