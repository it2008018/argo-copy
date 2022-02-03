// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/certificate/certificate.proto

// Certificate Service
//
// Certificate Service API performs CRUD actions against repository certificate
// resources.

package certificate

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
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

// Message to query the server for configured repository certificates
type RepositoryCertificateQuery struct {
	// A file-glob pattern (not regular expression) the host name has to match
	HostNamePattern string `protobuf:"bytes,1,opt,name=hostNamePattern,proto3" json:"hostNamePattern,omitempty"`
	// The type of the certificate to match (ssh or https)
	CertType string `protobuf:"bytes,2,opt,name=certType,proto3" json:"certType,omitempty"`
	// The sub type of the certificate to match (protocol dependent, usually only used for ssh certs)
	CertSubType          string   `protobuf:"bytes,3,opt,name=certSubType,proto3" json:"certSubType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepositoryCertificateQuery) Reset()         { *m = RepositoryCertificateQuery{} }
func (m *RepositoryCertificateQuery) String() string { return proto.CompactTextString(m) }
func (*RepositoryCertificateQuery) ProtoMessage()    {}
func (*RepositoryCertificateQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_387c41efc0710f00, []int{0}
}
func (m *RepositoryCertificateQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RepositoryCertificateQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RepositoryCertificateQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RepositoryCertificateQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepositoryCertificateQuery.Merge(m, src)
}
func (m *RepositoryCertificateQuery) XXX_Size() int {
	return m.Size()
}
func (m *RepositoryCertificateQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RepositoryCertificateQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RepositoryCertificateQuery proto.InternalMessageInfo

func (m *RepositoryCertificateQuery) GetHostNamePattern() string {
	if m != nil {
		return m.HostNamePattern
	}
	return ""
}

func (m *RepositoryCertificateQuery) GetCertType() string {
	if m != nil {
		return m.CertType
	}
	return ""
}

func (m *RepositoryCertificateQuery) GetCertSubType() string {
	if m != nil {
		return m.CertSubType
	}
	return ""
}

// Request to create a set of certificates
type RepositoryCertificateCreateRequest struct {
	// List of certificates to be created
	Certificates *v1alpha1.RepositoryCertificateList `protobuf:"bytes,1,opt,name=certificates,proto3" json:"certificates,omitempty"`
	// Whether to upsert already existing certificates
	Upsert               bool     `protobuf:"varint,2,opt,name=upsert,proto3" json:"upsert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepositoryCertificateCreateRequest) Reset()         { *m = RepositoryCertificateCreateRequest{} }
func (m *RepositoryCertificateCreateRequest) String() string { return proto.CompactTextString(m) }
func (*RepositoryCertificateCreateRequest) ProtoMessage()    {}
func (*RepositoryCertificateCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_387c41efc0710f00, []int{1}
}
func (m *RepositoryCertificateCreateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RepositoryCertificateCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RepositoryCertificateCreateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RepositoryCertificateCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepositoryCertificateCreateRequest.Merge(m, src)
}
func (m *RepositoryCertificateCreateRequest) XXX_Size() int {
	return m.Size()
}
func (m *RepositoryCertificateCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RepositoryCertificateCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RepositoryCertificateCreateRequest proto.InternalMessageInfo

func (m *RepositoryCertificateCreateRequest) GetCertificates() *v1alpha1.RepositoryCertificateList {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *RepositoryCertificateCreateRequest) GetUpsert() bool {
	if m != nil {
		return m.Upsert
	}
	return false
}

type RepositoryCertificateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepositoryCertificateResponse) Reset()         { *m = RepositoryCertificateResponse{} }
func (m *RepositoryCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*RepositoryCertificateResponse) ProtoMessage()    {}
func (*RepositoryCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_387c41efc0710f00, []int{2}
}
func (m *RepositoryCertificateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RepositoryCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RepositoryCertificateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RepositoryCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepositoryCertificateResponse.Merge(m, src)
}
func (m *RepositoryCertificateResponse) XXX_Size() int {
	return m.Size()
}
func (m *RepositoryCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RepositoryCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RepositoryCertificateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RepositoryCertificateQuery)(nil), "certificate.RepositoryCertificateQuery")
	proto.RegisterType((*RepositoryCertificateCreateRequest)(nil), "certificate.RepositoryCertificateCreateRequest")
	proto.RegisterType((*RepositoryCertificateResponse)(nil), "certificate.RepositoryCertificateResponse")
}

func init() {
	proto.RegisterFile("server/certificate/certificate.proto", fileDescriptor_387c41efc0710f00)
}

var fileDescriptor_387c41efc0710f00 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4f, 0x6b, 0x14, 0x31,
	0x14, 0x27, 0x55, 0x4a, 0x4d, 0x05, 0x6d, 0x28, 0xa5, 0x0c, 0x75, 0x2d, 0x43, 0xc1, 0x52, 0x30,
	0x61, 0xc7, 0x8b, 0x78, 0x74, 0xc5, 0x53, 0x11, 0x9d, 0x0a, 0x82, 0x17, 0xc9, 0xce, 0x3e, 0x67,
	0x63, 0xa7, 0x49, 0x4c, 0xde, 0x0c, 0x2c, 0xde, 0xc4, 0x6f, 0xe0, 0x37, 0xd1, 0x93, 0xdf, 0x40,
	0xf1, 0x22, 0xf8, 0x05, 0x64, 0xf1, 0x83, 0xc8, 0x64, 0x5b, 0x9b, 0x91, 0x11, 0xbd, 0x2c, 0xf4,
	0xf6, 0xf2, 0x5e, 0xde, 0x7b, 0xbf, 0x3f, 0x21, 0x74, 0xcf, 0x83, 0x6b, 0xc0, 0x89, 0x02, 0x1c,
	0xaa, 0x97, 0xaa, 0x90, 0x08, 0x71, 0xcc, 0xad, 0x33, 0x68, 0xd8, 0x7a, 0x94, 0x4a, 0x36, 0x4b,
	0x53, 0x9a, 0x90, 0x17, 0x6d, 0xb4, 0xb8, 0x92, 0xec, 0x94, 0xc6, 0x94, 0x15, 0x08, 0x69, 0x95,
	0x90, 0x5a, 0x1b, 0x94, 0xa8, 0x8c, 0xf6, 0xa7, 0xd5, 0xc3, 0x52, 0xe1, 0xb4, 0x1e, 0xf3, 0xc2,
	0x9c, 0x08, 0xe9, 0x42, 0xfb, 0xab, 0x10, 0xdc, 0x2e, 0x26, 0xa2, 0xc9, 0x84, 0x3d, 0x2e, 0xdb,
	0x4e, 0x2f, 0xa4, 0xb5, 0x55, 0xbb, 0x46, 0x19, 0x2d, 0x9a, 0xa1, 0xac, 0xec, 0x54, 0x0e, 0x45,
	0x09, 0x1a, 0x9c, 0x44, 0x98, 0x2c, 0xa6, 0xa5, 0xef, 0x08, 0x4d, 0x72, 0xb0, 0xc6, 0x2b, 0x34,
	0x6e, 0x36, 0x3a, 0xc7, 0xf6, 0xa4, 0x06, 0x37, 0x63, 0xfb, 0xf4, 0xda, 0xd4, 0x78, 0x7c, 0x24,
	0x4f, 0xe0, 0xb1, 0x44, 0x04, 0xa7, 0xb7, 0xc9, 0x2e, 0xd9, 0xbf, 0x92, 0xff, 0x99, 0x66, 0x09,
	0x5d, 0x6b, 0x99, 0x3d, 0x9d, 0x59, 0xd8, 0x5e, 0x09, 0x57, 0x7e, 0x9f, 0xd9, 0x2e, 0x0d, 0xac,
	0x8f, 0xea, 0x71, 0x28, 0x5f, 0x0a, 0xe5, 0x38, 0x95, 0x7e, 0x22, 0x34, 0xed, 0x85, 0x31, 0x72,
	0x20, 0x11, 0x72, 0x78, 0x5d, 0x83, 0x47, 0xf6, 0x86, 0x5e, 0x8d, 0xe4, 0xf3, 0x01, 0xcb, 0x7a,
	0xf6, 0x8c, 0x9f, 0x4b, 0xc2, 0xcf, 0x24, 0x09, 0xc1, 0x8b, 0x62, 0xc2, 0x9b, 0x8c, 0xdb, 0xe3,
	0x92, 0xb7, 0x92, 0xf0, 0x48, 0x12, 0x7e, 0x26, 0x09, 0xef, 0xdd, 0x7b, 0xa8, 0x3c, 0xe6, 0x9d,
	0x65, 0x6c, 0x8b, 0xae, 0xd6, 0xd6, 0x83, 0xc3, 0xc0, 0x6f, 0x2d, 0x3f, 0x3d, 0xa5, 0x37, 0xe9,
	0x8d, 0xde, 0x11, 0x39, 0x78, 0x6b, 0xb4, 0x87, 0xec, 0xeb, 0x65, 0xca, 0xa2, 0xfc, 0x11, 0xb8,
	0x46, 0x15, 0xc0, 0x3e, 0x10, 0x7a, 0xbd, 0x5d, 0x33, 0x8a, 0x97, 0xdc, 0xe2, 0xf1, 0x93, 0xf9,
	0xbb, 0x33, 0xc9, 0xb2, 0x48, 0xa7, 0x3b, 0x6f, 0xbf, 0xff, 0x7c, 0xbf, 0xb2, 0xc5, 0x36, 0xc3,
	0xfb, 0x6b, 0x86, 0xa2, 0x23, 0xc2, 0x17, 0x42, 0x37, 0x16, 0x9e, 0x44, 0x7d, 0x4c, 0xfc, 0x1b,
	0x75, 0xc7, 0xc8, 0xe5, 0xa1, 0x3f, 0x08, 0xe8, 0xf7, 0xd2, 0x5e, 0xf4, 0xf7, 0xba, 0x86, 0x7e,
	0x24, 0x74, 0xe3, 0x01, 0x54, 0xd0, 0xe5, 0x72, 0x61, 0x1c, 0x38, 0xe8, 0xe5, 0x70, 0xff, 0xe1,
	0xe7, 0xf9, 0x80, 0x7c, 0x9b, 0x0f, 0xc8, 0x8f, 0xf9, 0x80, 0x3c, 0xbf, 0xfb, 0x7f, 0xbf, 0x41,
	0x51, 0x29, 0xd0, 0x18, 0x0f, 0x1a, 0xaf, 0x86, 0x0f, 0xe0, 0xce, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x35, 0xdf, 0x5b, 0xaf, 0xb7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateServiceClient interface {
	// List all available repository certificates
	ListCertificates(ctx context.Context, in *RepositoryCertificateQuery, opts ...grpc.CallOption) (*v1alpha1.RepositoryCertificateList, error)
	// Creates repository certificates on the server
	CreateCertificate(ctx context.Context, in *RepositoryCertificateCreateRequest, opts ...grpc.CallOption) (*v1alpha1.RepositoryCertificateList, error)
	// Delete the certificates that match the RepositoryCertificateQuery
	DeleteCertificate(ctx context.Context, in *RepositoryCertificateQuery, opts ...grpc.CallOption) (*v1alpha1.RepositoryCertificateList, error)
}

type certificateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCertificateServiceClient(cc *grpc.ClientConn) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) ListCertificates(ctx context.Context, in *RepositoryCertificateQuery, opts ...grpc.CallOption) (*v1alpha1.RepositoryCertificateList, error) {
	out := new(v1alpha1.RepositoryCertificateList)
	err := c.cc.Invoke(ctx, "/certificate.CertificateService/ListCertificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) CreateCertificate(ctx context.Context, in *RepositoryCertificateCreateRequest, opts ...grpc.CallOption) (*v1alpha1.RepositoryCertificateList, error) {
	out := new(v1alpha1.RepositoryCertificateList)
	err := c.cc.Invoke(ctx, "/certificate.CertificateService/CreateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) DeleteCertificate(ctx context.Context, in *RepositoryCertificateQuery, opts ...grpc.CallOption) (*v1alpha1.RepositoryCertificateList, error) {
	out := new(v1alpha1.RepositoryCertificateList)
	err := c.cc.Invoke(ctx, "/certificate.CertificateService/DeleteCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
type CertificateServiceServer interface {
	// List all available repository certificates
	ListCertificates(context.Context, *RepositoryCertificateQuery) (*v1alpha1.RepositoryCertificateList, error)
	// Creates repository certificates on the server
	CreateCertificate(context.Context, *RepositoryCertificateCreateRequest) (*v1alpha1.RepositoryCertificateList, error)
	// Delete the certificates that match the RepositoryCertificateQuery
	DeleteCertificate(context.Context, *RepositoryCertificateQuery) (*v1alpha1.RepositoryCertificateList, error)
}

// UnimplementedCertificateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateServiceServer struct {
}

func (*UnimplementedCertificateServiceServer) ListCertificates(ctx context.Context, req *RepositoryCertificateQuery) (*v1alpha1.RepositoryCertificateList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCertificates not implemented")
}
func (*UnimplementedCertificateServiceServer) CreateCertificate(ctx context.Context, req *RepositoryCertificateCreateRequest) (*v1alpha1.RepositoryCertificateList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}
func (*UnimplementedCertificateServiceServer) DeleteCertificate(ctx context.Context, req *RepositoryCertificateQuery) (*v1alpha1.RepositoryCertificateList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCertificate not implemented")
}

func RegisterCertificateServiceServer(s *grpc.Server, srv CertificateServiceServer) {
	s.RegisterService(&_CertificateService_serviceDesc, srv)
}

func _CertificateService_ListCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryCertificateQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).ListCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.CertificateService/ListCertificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).ListCertificates(ctx, req.(*RepositoryCertificateQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryCertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.CertificateService/CreateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, req.(*RepositoryCertificateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_DeleteCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryCertificateQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).DeleteCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.CertificateService/DeleteCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).DeleteCertificate(ctx, req.(*RepositoryCertificateQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCertificates",
			Handler:    _CertificateService_ListCertificates_Handler,
		},
		{
			MethodName: "CreateCertificate",
			Handler:    _CertificateService_CreateCertificate_Handler,
		},
		{
			MethodName: "DeleteCertificate",
			Handler:    _CertificateService_DeleteCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/certificate/certificate.proto",
}

func (m *RepositoryCertificateQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepositoryCertificateQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RepositoryCertificateQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CertSubType) > 0 {
		i -= len(m.CertSubType)
		copy(dAtA[i:], m.CertSubType)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.CertSubType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CertType) > 0 {
		i -= len(m.CertType)
		copy(dAtA[i:], m.CertType)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.CertType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HostNamePattern) > 0 {
		i -= len(m.HostNamePattern)
		copy(dAtA[i:], m.HostNamePattern)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.HostNamePattern)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RepositoryCertificateCreateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepositoryCertificateCreateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RepositoryCertificateCreateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Upsert {
		i--
		if m.Upsert {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Certificates != nil {
		{
			size, err := m.Certificates.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCertificate(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RepositoryCertificateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepositoryCertificateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RepositoryCertificateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintCertificate(dAtA []byte, offset int, v uint64) int {
	offset -= sovCertificate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RepositoryCertificateQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HostNamePattern)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.CertType)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.CertSubType)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RepositoryCertificateCreateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Certificates != nil {
		l = m.Certificates.Size()
		n += 1 + l + sovCertificate(uint64(l))
	}
	if m.Upsert {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RepositoryCertificateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCertificate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCertificate(x uint64) (n int) {
	return sovCertificate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RepositoryCertificateQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertificate
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
			return fmt.Errorf("proto: RepositoryCertificateQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepositoryCertificateQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostNamePattern", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostNamePattern = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertSubType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertSubType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertificate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RepositoryCertificateCreateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertificate
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
			return fmt.Errorf("proto: RepositoryCertificateCreateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepositoryCertificateCreateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Certificates == nil {
				m.Certificates = &v1alpha1.RepositoryCertificateList{}
			}
			if err := m.Certificates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upsert", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Upsert = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertificate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RepositoryCertificateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertificate
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
			return fmt.Errorf("proto: RepositoryCertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepositoryCertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertificate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCertificate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCertificate
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
					return 0, ErrIntOverflowCertificate
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
					return 0, ErrIntOverflowCertificate
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
				return 0, ErrInvalidLengthCertificate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCertificate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCertificate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCertificate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCertificate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCertificate = fmt.Errorf("proto: unexpected end of group")
)
