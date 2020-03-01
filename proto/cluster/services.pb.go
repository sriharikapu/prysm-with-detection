// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/cluster/services.proto

package prysm_internal_cluster

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PrivateKeyRequest struct {
	PodName              string   `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	NumberOfKeys         uint64   `protobuf:"varint,2,opt,name=number_of_keys,json=numberOfKeys,proto3" json:"number_of_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivateKeyRequest) Reset()         { *m = PrivateKeyRequest{} }
func (m *PrivateKeyRequest) String() string { return proto.CompactTextString(m) }
func (*PrivateKeyRequest) ProtoMessage()    {}
func (*PrivateKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f118b0fcaa41cfbe, []int{0}
}

func (m *PrivateKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateKeyRequest.Unmarshal(m, b)
}
func (m *PrivateKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateKeyRequest.Marshal(b, m, deterministic)
}
func (m *PrivateKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateKeyRequest.Merge(m, src)
}
func (m *PrivateKeyRequest) XXX_Size() int {
	return xxx_messageInfo_PrivateKeyRequest.Size(m)
}
func (m *PrivateKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateKeyRequest proto.InternalMessageInfo

func (m *PrivateKeyRequest) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *PrivateKeyRequest) GetNumberOfKeys() uint64 {
	if m != nil {
		return m.NumberOfKeys
	}
	return 0
}

type PrivateKeyResponse struct {
	PrivateKey           []byte       `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"` // Deprecated: Do not use.
	PrivateKeys          *PrivateKeys `protobuf:"bytes,2,opt,name=private_keys,json=privateKeys,proto3" json:"private_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PrivateKeyResponse) Reset()         { *m = PrivateKeyResponse{} }
func (m *PrivateKeyResponse) String() string { return proto.CompactTextString(m) }
func (*PrivateKeyResponse) ProtoMessage()    {}
func (*PrivateKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f118b0fcaa41cfbe, []int{1}
}

func (m *PrivateKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateKeyResponse.Unmarshal(m, b)
}
func (m *PrivateKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateKeyResponse.Marshal(b, m, deterministic)
}
func (m *PrivateKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateKeyResponse.Merge(m, src)
}
func (m *PrivateKeyResponse) XXX_Size() int {
	return xxx_messageInfo_PrivateKeyResponse.Size(m)
}
func (m *PrivateKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateKeyResponse proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *PrivateKeyResponse) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *PrivateKeyResponse) GetPrivateKeys() *PrivateKeys {
	if m != nil {
		return m.PrivateKeys
	}
	return nil
}

type PrivateKeys struct {
	PrivateKeys          [][]byte `protobuf:"bytes,1,rep,name=private_keys,json=privateKeys,proto3" json:"private_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivateKeys) Reset()         { *m = PrivateKeys{} }
func (m *PrivateKeys) String() string { return proto.CompactTextString(m) }
func (*PrivateKeys) ProtoMessage()    {}
func (*PrivateKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_f118b0fcaa41cfbe, []int{2}
}

func (m *PrivateKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateKeys.Unmarshal(m, b)
}
func (m *PrivateKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateKeys.Marshal(b, m, deterministic)
}
func (m *PrivateKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateKeys.Merge(m, src)
}
func (m *PrivateKeys) XXX_Size() int {
	return xxx_messageInfo_PrivateKeys.Size(m)
}
func (m *PrivateKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateKeys.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateKeys proto.InternalMessageInfo

func (m *PrivateKeys) GetPrivateKeys() [][]byte {
	if m != nil {
		return m.PrivateKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*PrivateKeyRequest)(nil), "prysm.internal.cluster.PrivateKeyRequest")
	proto.RegisterType((*PrivateKeyResponse)(nil), "prysm.internal.cluster.PrivateKeyResponse")
	proto.RegisterType((*PrivateKeys)(nil), "prysm.internal.cluster.PrivateKeys")
}

func init() { proto.RegisterFile("proto/cluster/services.proto", fileDescriptor_f118b0fcaa41cfbe) }

var fileDescriptor_f118b0fcaa41cfbe = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x14, 0xa7, 0xaf, 0x45, 0x30, 0x07, 0xa9, 0xe2, 0xa1, 0x76, 0x1e, 0xaa, 0x87,
	0x4c, 0xe6, 0x37, 0xf0, 0xe0, 0x65, 0xa0, 0x12, 0xbd, 0xd7, 0x6e, 0x7b, 0x83, 0xe2, 0x9a, 0xc4,
	0xbc, 0x74, 0xd0, 0xa3, 0xdf, 0x5c, 0x6c, 0xa6, 0x8d, 0x22, 0xb8, 0xeb, 0x3f, 0xbf, 0xf7, 0x4b,
	0xf2, 0x7f, 0x70, 0x66, 0xac, 0x76, 0x7a, 0x3c, 0x5f, 0x35, 0xe4, 0xd0, 0x8e, 0x09, 0xed, 0xba,
	0x9a, 0x23, 0x89, 0x2e, 0xe6, 0xc7, 0xc6, 0xb6, 0x54, 0x8b, 0x4a, 0x39, 0xb4, 0xaa, 0x5c, 0x89,
	0x0d, 0x96, 0x3d, 0xc3, 0xd1, 0xa3, 0xad, 0xd6, 0xa5, 0xc3, 0x29, 0xb6, 0x12, 0xdf, 0x1a, 0x24,
	0xc7, 0x4f, 0x60, 0xdf, 0xe8, 0x45, 0xa1, 0xca, 0x1a, 0x13, 0x96, 0xb2, 0xfc, 0x40, 0x0e, 0x8d,
	0x5e, 0xdc, 0x97, 0x35, 0xf2, 0x0b, 0x38, 0x54, 0x4d, 0x3d, 0x43, 0x5b, 0xe8, 0x65, 0xf1, 0x8a,
	0x2d, 0x25, 0x83, 0x94, 0xe5, 0xbb, 0x32, 0xf6, 0xe9, 0xc3, 0x72, 0x8a, 0x2d, 0x65, 0xef, 0x0c,
	0x78, 0xa8, 0x25, 0xa3, 0x15, 0x21, 0x1f, 0x41, 0x64, 0x7c, 0xfa, 0x39, 0xda, 0xa9, 0xe3, 0xdb,
	0x41, 0xc2, 0x24, 0x98, 0x6f, 0x98, 0xdf, 0x41, 0x1c, 0x40, 0xde, 0x1f, 0x4d, 0x46, 0xe2, 0xef,
	0x0f, 0x88, 0xfe, 0x1a, 0x92, 0x51, 0xaf, 0xa1, 0xec, 0x1a, 0xa2, 0xe0, 0x8c, 0x9f, 0xff, 0xd2,
	0xb2, 0x74, 0x27, 0x8f, 0x7f, 0x4c, 0x4c, 0x9a, 0xb0, 0x8b, 0x27, 0xdf, 0x1f, 0x7f, 0x81, 0xe1,
	0x57, 0x2d, 0x97, 0xff, 0xbf, 0x61, 0x83, 0x9e, 0x5e, 0x6d, 0x83, 0xfa, 0x56, 0x66, 0x7b, 0xdd,
	0x86, 0x6e, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x28, 0x1a, 0xa8, 0xa6, 0xc1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrivateKeyServiceClient is the client API for PrivateKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrivateKeyServiceClient interface {
	Request(ctx context.Context, in *PrivateKeyRequest, opts ...grpc.CallOption) (*PrivateKeyResponse, error)
}

type privateKeyServiceClient struct {
	cc *grpc.ClientConn
}

func NewPrivateKeyServiceClient(cc *grpc.ClientConn) PrivateKeyServiceClient {
	return &privateKeyServiceClient{cc}
}

func (c *privateKeyServiceClient) Request(ctx context.Context, in *PrivateKeyRequest, opts ...grpc.CallOption) (*PrivateKeyResponse, error) {
	out := new(PrivateKeyResponse)
	err := c.cc.Invoke(ctx, "/prysm.internal.cluster.PrivateKeyService/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateKeyServiceServer is the server API for PrivateKeyService service.
type PrivateKeyServiceServer interface {
	Request(context.Context, *PrivateKeyRequest) (*PrivateKeyResponse, error)
}

// UnimplementedPrivateKeyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPrivateKeyServiceServer struct {
}

func (*UnimplementedPrivateKeyServiceServer) Request(ctx context.Context, req *PrivateKeyRequest) (*PrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}

func RegisterPrivateKeyServiceServer(s *grpc.Server, srv PrivateKeyServiceServer) {
	s.RegisterService(&_PrivateKeyService_serviceDesc, srv)
}

func _PrivateKeyService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateKeyServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prysm.internal.cluster.PrivateKeyService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateKeyServiceServer).Request(ctx, req.(*PrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivateKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prysm.internal.cluster.PrivateKeyService",
	HandlerType: (*PrivateKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _PrivateKeyService_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cluster/services.proto",
}