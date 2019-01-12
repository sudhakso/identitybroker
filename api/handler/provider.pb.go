// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provider.proto

package api

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

type Provider_State int32

const (
	Provider_ENABLED  Provider_State = 0
	Provider_DISABLED Provider_State = 1
	Provider_ERROR    Provider_State = 2
)

var Provider_State_name = map[int32]string{
	0: "ENABLED",
	1: "DISABLED",
	2: "ERROR",
}

var Provider_State_value = map[string]int32{
	"ENABLED":  0,
	"DISABLED": 1,
	"ERROR":    2,
}

func (x Provider_State) String() string {
	return proto.EnumName(Provider_State_name, int32(x))
}

func (Provider_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{3, 0}
}

type Credential struct {
	ApiKey               string   `protobuf:"bytes,1,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	AuthUrl              string   `protobuf:"bytes,3,opt,name=authUrl,proto3" json:"authUrl,omitempty"`
	Base64Cert           string   `protobuf:"bytes,4,opt,name=base64Cert,proto3" json:"base64Cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{0}
}

func (m *Credential) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credential.Unmarshal(m, b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return xxx_messageInfo_Credential.Size(m)
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

func (m *Credential) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *Credential) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Credential) GetAuthUrl() string {
	if m != nil {
		return m.AuthUrl
	}
	return ""
}

func (m *Credential) GetBase64Cert() string {
	if m != nil {
		return m.Base64Cert
	}
	return ""
}

type ProviderRegistrationOpts struct {
	Namespace            string      `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ProviderType         string      `protobuf:"bytes,2,opt,name=providerType,proto3" json:"providerType,omitempty"`
	Cred                 *Credential `protobuf:"bytes,3,opt,name=cred,proto3" json:"cred,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ProviderRegistrationOpts) Reset()         { *m = ProviderRegistrationOpts{} }
func (m *ProviderRegistrationOpts) String() string { return proto.CompactTextString(m) }
func (*ProviderRegistrationOpts) ProtoMessage()    {}
func (*ProviderRegistrationOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{1}
}

func (m *ProviderRegistrationOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderRegistrationOpts.Unmarshal(m, b)
}
func (m *ProviderRegistrationOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderRegistrationOpts.Marshal(b, m, deterministic)
}
func (m *ProviderRegistrationOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderRegistrationOpts.Merge(m, src)
}
func (m *ProviderRegistrationOpts) XXX_Size() int {
	return xxx_messageInfo_ProviderRegistrationOpts.Size(m)
}
func (m *ProviderRegistrationOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderRegistrationOpts.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderRegistrationOpts proto.InternalMessageInfo

func (m *ProviderRegistrationOpts) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ProviderRegistrationOpts) GetProviderType() string {
	if m != nil {
		return m.ProviderType
	}
	return ""
}

func (m *ProviderRegistrationOpts) GetCred() *Credential {
	if m != nil {
		return m.Cred
	}
	return nil
}

type RegistrationStatus struct {
	Error                bool     `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	OriginalError        string   `protobuf:"bytes,2,opt,name=originalError,proto3" json:"originalError,omitempty"`
	ProviderId           string   `protobuf:"bytes,3,opt,name=providerId,proto3" json:"providerId,omitempty"`
	ProviderNamespace    string   `protobuf:"bytes,4,opt,name=providerNamespace,proto3" json:"providerNamespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationStatus) Reset()         { *m = RegistrationStatus{} }
func (m *RegistrationStatus) String() string { return proto.CompactTextString(m) }
func (*RegistrationStatus) ProtoMessage()    {}
func (*RegistrationStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{2}
}

func (m *RegistrationStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationStatus.Unmarshal(m, b)
}
func (m *RegistrationStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationStatus.Marshal(b, m, deterministic)
}
func (m *RegistrationStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationStatus.Merge(m, src)
}
func (m *RegistrationStatus) XXX_Size() int {
	return xxx_messageInfo_RegistrationStatus.Size(m)
}
func (m *RegistrationStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationStatus proto.InternalMessageInfo

func (m *RegistrationStatus) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *RegistrationStatus) GetOriginalError() string {
	if m != nil {
		return m.OriginalError
	}
	return ""
}

func (m *RegistrationStatus) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *RegistrationStatus) GetProviderNamespace() string {
	if m != nil {
		return m.ProviderNamespace
	}
	return ""
}

type Provider struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Namespace            string          `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	State                Provider_State  `protobuf:"varint,3,opt,name=state,proto3,enum=api.Provider_State" json:"state,omitempty"`
	ProviderType         string          `protobuf:"bytes,4,opt,name=providerType,proto3" json:"providerType,omitempty"`
	ResourcePathprefix   string          `protobuf:"bytes,5,opt,name=resourcePathprefix,proto3" json:"resourcePathprefix,omitempty"`
	ResourceTypes        []*ResourceType `protobuf:"bytes,6,rep,name=resourceTypes,proto3" json:"resourceTypes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{3}
}

func (m *Provider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Provider.Unmarshal(m, b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return xxx_messageInfo_Provider.Size(m)
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Provider) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Provider) GetState() Provider_State {
	if m != nil {
		return m.State
	}
	return Provider_ENABLED
}

func (m *Provider) GetProviderType() string {
	if m != nil {
		return m.ProviderType
	}
	return ""
}

func (m *Provider) GetResourcePathprefix() string {
	if m != nil {
		return m.ResourcePathprefix
	}
	return ""
}

func (m *Provider) GetResourceTypes() []*ResourceType {
	if m != nil {
		return m.ResourceTypes
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.Provider_State", Provider_State_name, Provider_State_value)
	proto.RegisterType((*Credential)(nil), "api.Credential")
	proto.RegisterType((*ProviderRegistrationOpts)(nil), "api.ProviderRegistrationOpts")
	proto.RegisterType((*RegistrationStatus)(nil), "api.RegistrationStatus")
	proto.RegisterType((*Provider)(nil), "api.Provider")
}

func init() { proto.RegisterFile("provider.proto", fileDescriptor_c6a9f3c02af3d1c8) }

var fileDescriptor_c6a9f3c02af3d1c8 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0x38, 0x75, 0x26, 0x6d, 0x68, 0x87, 0x0a, 0xac, 0x08, 0x50, 0x65, 0x38, 0x14,
	0x09, 0x7c, 0x08, 0x08, 0xce, 0xd0, 0x04, 0x29, 0x02, 0xb5, 0xd5, 0x16, 0x3e, 0x60, 0x1b, 0x0f,
	0xed, 0x4a, 0xc1, 0x5e, 0xcd, 0x6e, 0x10, 0xe5, 0xc6, 0x8f, 0xc0, 0x85, 0x0f, 0x45, 0x59, 0x7b,
	0x5b, 0x3b, 0xcd, 0x8d, 0xe3, 0x7b, 0x6f, 0xbc, 0xf3, 0xde, 0xd3, 0x18, 0x86, 0x9a, 0xcb, 0xef,
	0x2a, 0x27, 0xce, 0x34, 0x97, 0xb6, 0xc4, 0x8e, 0xd4, 0x6a, 0x34, 0x64, 0x32, 0xe5, 0x92, 0xe7,
	0x54, 0x91, 0xe9, 0x4f, 0x80, 0x63, 0xa6, 0x9c, 0x0a, 0xab, 0xe4, 0x02, 0x1f, 0x40, 0x4f, 0x6a,
	0xf5, 0x91, 0xae, 0x93, 0xe0, 0x30, 0x38, 0xea, 0x8b, 0x1a, 0xe1, 0x08, 0xe2, 0xf9, 0x42, 0x51,
	0x61, 0x67, 0x79, 0x12, 0x3a, 0xe5, 0x06, 0x63, 0x02, 0xdb, 0x72, 0x69, 0xaf, 0xbe, 0xf0, 0x22,
	0xe9, 0x38, 0xc9, 0x43, 0x7c, 0x02, 0x70, 0x21, 0x0d, 0xbd, 0x79, 0x7d, 0x4c, 0x6c, 0x93, 0xae,
	0x13, 0x1b, 0x4c, 0xfa, 0x2b, 0x80, 0xe4, 0xac, 0xf6, 0x28, 0xe8, 0x52, 0x19, 0xcb, 0xd2, 0xaa,
	0xb2, 0x38, 0xd5, 0xd6, 0xe0, 0x23, 0xe8, 0x17, 0xf2, 0x1b, 0x19, 0x2d, 0xe7, 0x54, 0xbb, 0xb9,
	0x25, 0x30, 0x85, 0x1d, 0x9f, 0xee, 0xf3, 0xb5, 0xa6, 0xda, 0x54, 0x8b, 0xc3, 0xa7, 0xd0, 0x9d,
	0x33, 0xe5, 0xce, 0xd5, 0x60, 0x7c, 0x2f, 0x93, 0x5a, 0x65, 0xb7, 0x59, 0x85, 0x13, 0xd3, 0xdf,
	0x01, 0x60, 0x73, 0xf7, 0xb9, 0x95, 0x76, 0x69, 0xf0, 0x00, 0x22, 0x62, 0x2e, 0xd9, 0x6d, 0x8e,
	0x45, 0x05, 0xf0, 0x19, 0xec, 0x96, 0xac, 0x2e, 0x55, 0x21, 0x17, 0x53, 0xa7, 0x56, 0x6b, 0xdb,
	0xe4, 0x2a, 0xb6, 0xf7, 0x31, 0xcb, 0xeb, 0x4e, 0x1a, 0x0c, 0xbe, 0x80, 0x7d, 0x8f, 0x4e, 0x6e,
	0x12, 0x56, 0xed, 0xdc, 0x15, 0xd2, 0x3f, 0x21, 0xc4, 0xbe, 0x24, 0x1c, 0x42, 0xa8, 0xf2, 0xba,
	0x8d, 0x50, 0xe5, 0xed, 0x92, 0xc2, 0xf5, 0x92, 0x9e, 0x43, 0x64, 0xac, 0xb4, 0xe4, 0x3c, 0x0c,
	0xc7, 0xf7, 0x5d, 0x03, 0xfe, 0xad, 0x6c, 0x95, 0x94, 0x44, 0x35, 0x71, 0xa7, 0xcf, 0xee, 0x86,
	0x3e, 0x33, 0x40, 0x7f, 0x3c, 0x67, 0xd2, 0x5e, 0x69, 0xa6, 0xaf, 0xea, 0x47, 0x12, 0xb9, 0xc9,
	0x0d, 0x0a, 0xbe, 0x85, 0x5d, 0xcf, 0xae, 0xbe, 0x37, 0x49, 0xef, 0xb0, 0x73, 0x34, 0x18, 0xef,
	0x3b, 0x1b, 0xa2, 0xa1, 0x88, 0xf6, 0x5c, 0xfa, 0x12, 0x22, 0x67, 0x0e, 0x07, 0xb0, 0x3d, 0x3d,
	0x79, 0xf7, 0xfe, 0xd3, 0x74, 0xb2, 0xb7, 0x85, 0x3b, 0x10, 0x4f, 0x66, 0xe7, 0x15, 0x0a, 0xb0,
	0x0f, 0xd1, 0x54, 0x88, 0x53, 0xb1, 0x17, 0x8e, 0xff, 0x06, 0x70, 0xb0, 0xe9, 0x8c, 0x70, 0x02,
	0x71, 0x85, 0x89, 0xf1, 0x71, 0x2b, 0xfc, 0xfa, 0xb5, 0x8d, 0x1e, 0xd6, 0xa6, 0xd6, 0x0f, 0x21,
	0xdd, 0xc2, 0x0f, 0x00, 0x13, 0xfa, 0xff, 0x77, 0x2e, 0x7a, 0xee, 0x87, 0x7b, 0xf5, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0xa2, 0xdb, 0x3b, 0x65, 0x97, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProviderRegistrationClient is the client API for ProviderRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProviderRegistrationClient interface {
	Register(ctx context.Context, in *ProviderRegistrationOpts, opts ...grpc.CallOption) (*RegistrationStatus, error)
	DeRegister(ctx context.Context, in *ProviderRegistrationOpts, opts ...grpc.CallOption) (*RegistrationStatus, error)
}

type providerRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewProviderRegistrationClient(cc *grpc.ClientConn) ProviderRegistrationClient {
	return &providerRegistrationClient{cc}
}

func (c *providerRegistrationClient) Register(ctx context.Context, in *ProviderRegistrationOpts, opts ...grpc.CallOption) (*RegistrationStatus, error) {
	out := new(RegistrationStatus)
	err := c.cc.Invoke(ctx, "/api.ProviderRegistration/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerRegistrationClient) DeRegister(ctx context.Context, in *ProviderRegistrationOpts, opts ...grpc.CallOption) (*RegistrationStatus, error) {
	out := new(RegistrationStatus)
	err := c.cc.Invoke(ctx, "/api.ProviderRegistration/DeRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderRegistrationServer is the server API for ProviderRegistration service.
type ProviderRegistrationServer interface {
	Register(context.Context, *ProviderRegistrationOpts) (*RegistrationStatus, error)
	DeRegister(context.Context, *ProviderRegistrationOpts) (*RegistrationStatus, error)
}

func RegisterProviderRegistrationServer(s *grpc.Server, srv ProviderRegistrationServer) {
	s.RegisterService(&_ProviderRegistration_serviceDesc, srv)
}

func _ProviderRegistration_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderRegistrationOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderRegistrationServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ProviderRegistration/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderRegistrationServer).Register(ctx, req.(*ProviderRegistrationOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderRegistration_DeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderRegistrationOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderRegistrationServer).DeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ProviderRegistration/DeRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderRegistrationServer).DeRegister(ctx, req.(*ProviderRegistrationOpts))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ProviderRegistration",
	HandlerType: (*ProviderRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ProviderRegistration_Register_Handler,
		},
		{
			MethodName: "DeRegister",
			Handler:    _ProviderRegistration_DeRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}
