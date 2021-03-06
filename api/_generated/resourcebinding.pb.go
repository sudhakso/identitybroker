// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resourcebinding.proto

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

type ResourceIdentifier struct {
	Type                 *ResourceType `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PathPrefix           string        `protobuf:"bytes,3,opt,name=pathPrefix,proto3" json:"pathPrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ResourceIdentifier) Reset()         { *m = ResourceIdentifier{} }
func (m *ResourceIdentifier) String() string { return proto.CompactTextString(m) }
func (*ResourceIdentifier) ProtoMessage()    {}
func (*ResourceIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_586d4f0847511622, []int{0}
}

func (m *ResourceIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceIdentifier.Unmarshal(m, b)
}
func (m *ResourceIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceIdentifier.Marshal(b, m, deterministic)
}
func (m *ResourceIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceIdentifier.Merge(m, src)
}
func (m *ResourceIdentifier) XXX_Size() int {
	return xxx_messageInfo_ResourceIdentifier.Size(m)
}
func (m *ResourceIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceIdentifier proto.InternalMessageInfo

func (m *ResourceIdentifier) GetType() *ResourceType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ResourceIdentifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceIdentifier) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

type ResourceLink struct {
	FromResource         *ResourceIdentifier `protobuf:"bytes,1,opt,name=fromResource,proto3" json:"fromResource,omitempty"`
	ToResource           *ResourceIdentifier `protobuf:"bytes,2,opt,name=toResource,proto3" json:"toResource,omitempty"`
	Annotations          []string            `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ResourceLink) Reset()         { *m = ResourceLink{} }
func (m *ResourceLink) String() string { return proto.CompactTextString(m) }
func (*ResourceLink) ProtoMessage()    {}
func (*ResourceLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_586d4f0847511622, []int{1}
}

func (m *ResourceLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceLink.Unmarshal(m, b)
}
func (m *ResourceLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceLink.Marshal(b, m, deterministic)
}
func (m *ResourceLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceLink.Merge(m, src)
}
func (m *ResourceLink) XXX_Size() int {
	return xxx_messageInfo_ResourceLink.Size(m)
}
func (m *ResourceLink) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceLink.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceLink proto.InternalMessageInfo

func (m *ResourceLink) GetFromResource() *ResourceIdentifier {
	if m != nil {
		return m.FromResource
	}
	return nil
}

func (m *ResourceLink) GetToResource() *ResourceIdentifier {
	if m != nil {
		return m.ToResource
	}
	return nil
}

func (m *ResourceLink) GetAnnotations() []string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type ResourceBindings struct {
	Bindings             []*ResourceLink `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ResourceBindings) Reset()         { *m = ResourceBindings{} }
func (m *ResourceBindings) String() string { return proto.CompactTextString(m) }
func (*ResourceBindings) ProtoMessage()    {}
func (*ResourceBindings) Descriptor() ([]byte, []int) {
	return fileDescriptor_586d4f0847511622, []int{2}
}

func (m *ResourceBindings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceBindings.Unmarshal(m, b)
}
func (m *ResourceBindings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceBindings.Marshal(b, m, deterministic)
}
func (m *ResourceBindings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceBindings.Merge(m, src)
}
func (m *ResourceBindings) XXX_Size() int {
	return xxx_messageInfo_ResourceBindings.Size(m)
}
func (m *ResourceBindings) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceBindings.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceBindings proto.InternalMessageInfo

func (m *ResourceBindings) GetBindings() []*ResourceLink {
	if m != nil {
		return m.Bindings
	}
	return nil
}

type ResourceBindingOpts struct {
	FromResourceId       *ResourceIdentifier `protobuf:"bytes,1,opt,name=fromResourceId,proto3" json:"fromResourceId,omitempty"`
	ToResourceId         *ResourceIdentifier `protobuf:"bytes,2,opt,name=toResourceId,proto3" json:"toResourceId,omitempty"`
	Provider             *ProviderOpts       `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ResourceBindingOpts) Reset()         { *m = ResourceBindingOpts{} }
func (m *ResourceBindingOpts) String() string { return proto.CompactTextString(m) }
func (*ResourceBindingOpts) ProtoMessage()    {}
func (*ResourceBindingOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_586d4f0847511622, []int{3}
}

func (m *ResourceBindingOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceBindingOpts.Unmarshal(m, b)
}
func (m *ResourceBindingOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceBindingOpts.Marshal(b, m, deterministic)
}
func (m *ResourceBindingOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceBindingOpts.Merge(m, src)
}
func (m *ResourceBindingOpts) XXX_Size() int {
	return xxx_messageInfo_ResourceBindingOpts.Size(m)
}
func (m *ResourceBindingOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceBindingOpts.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceBindingOpts proto.InternalMessageInfo

func (m *ResourceBindingOpts) GetFromResourceId() *ResourceIdentifier {
	if m != nil {
		return m.FromResourceId
	}
	return nil
}

func (m *ResourceBindingOpts) GetToResourceId() *ResourceIdentifier {
	if m != nil {
		return m.ToResourceId
	}
	return nil
}

func (m *ResourceBindingOpts) GetProvider() *ProviderOpts {
	if m != nil {
		return m.Provider
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceIdentifier)(nil), "api.ResourceIdentifier")
	proto.RegisterType((*ResourceLink)(nil), "api.ResourceLink")
	proto.RegisterType((*ResourceBindings)(nil), "api.ResourceBindings")
	proto.RegisterType((*ResourceBindingOpts)(nil), "api.ResourceBindingOpts")
}

func init() { proto.RegisterFile("resourcebinding.proto", fileDescriptor_586d4f0847511622) }

var fileDescriptor_586d4f0847511622 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0x33, 0x31,
	0x10, 0xc7, 0xbf, 0x74, 0xcb, 0x87, 0x9d, 0x96, 0xaa, 0x23, 0xc5, 0xd0, 0x83, 0x2c, 0x0b, 0x42,
	0x2f, 0xf6, 0x50, 0x0f, 0x1e, 0x3c, 0x48, 0xbd, 0x48, 0x41, 0xb0, 0x04, 0x5f, 0x20, 0x75, 0xd3,
	0x1a, 0xa4, 0x49, 0x48, 0x62, 0xb1, 0xcf, 0x23, 0xf8, 0x20, 0x3e, 0x99, 0xec, 0x76, 0xd3, 0xee,
	0xae, 0x62, 0x0f, 0xde, 0x96, 0xff, 0xfc, 0x32, 0x93, 0xdf, 0x64, 0xa1, 0x67, 0x85, 0xd3, 0xaf,
	0xf6, 0x49, 0xcc, 0xa4, 0x4a, 0xa5, 0x5a, 0x0c, 0x8d, 0xd5, 0x5e, 0x63, 0xc4, 0x8d, 0xec, 0x77,
	0x43, 0x6d, 0x13, 0x26, 0x1a, 0x90, 0x15, 0xc9, 0x24, 0x15, 0xca, 0xcb, 0xb9, 0x14, 0x16, 0xcf,
	0xa1, 0xe9, 0xd7, 0x46, 0x50, 0x12, 0x93, 0x41, 0x7b, 0x74, 0x3c, 0xe4, 0x46, 0x0e, 0x03, 0xf6,
	0xb8, 0x36, 0x82, 0xe5, 0x65, 0x44, 0x68, 0x2a, 0xbe, 0x14, 0xb4, 0x11, 0x93, 0x41, 0x8b, 0xe5,
	0xdf, 0x78, 0x06, 0x60, 0xb8, 0x7f, 0x9e, 0x5a, 0x31, 0x97, 0x6f, 0x34, 0xca, 0x2b, 0xa5, 0x24,
	0xf9, 0x20, 0xd0, 0x09, 0xad, 0xee, 0xa5, 0x7a, 0xc1, 0x6b, 0xe8, 0xcc, 0xad, 0x5e, 0x86, 0xac,
	0x98, 0x79, 0x5a, 0x99, 0xb9, 0xbb, 0x1a, 0xab, 0xc0, 0x78, 0x05, 0xe0, 0xf5, 0xf6, 0x68, 0xe3,
	0xf7, 0xa3, 0x25, 0x14, 0x63, 0x68, 0x73, 0xa5, 0xb4, 0xe7, 0x5e, 0x6a, 0xe5, 0x68, 0x14, 0x47,
	0x83, 0x16, 0x2b, 0x47, 0xc9, 0x18, 0x8e, 0x02, 0x7d, 0xbb, 0xd9, 0xa3, 0xc3, 0x0b, 0x38, 0x28,
	0x76, 0xea, 0x28, 0x89, 0xa3, 0x6f, 0xbb, 0xc9, 0x84, 0xd8, 0x16, 0x49, 0x3e, 0x09, 0x9c, 0xd4,
	0x7a, 0x3c, 0x18, 0xef, 0xf0, 0x06, 0xba, 0x65, 0x8b, 0x49, 0xba, 0x4f, 0xba, 0x86, 0x67, 0x3b,
	0xdb, 0xb9, 0x4c, 0xd2, 0x7d, 0xe2, 0x15, 0x38, 0x93, 0x30, 0x56, 0xaf, 0x64, 0x2a, 0x6c, 0xfe,
	0x3e, 0x41, 0x62, 0x5a, 0x84, 0xd9, 0x15, 0xd9, 0x16, 0x19, 0xbd, 0x13, 0x38, 0xac, 0x49, 0xe0,
	0x1d, 0xe0, 0xd8, 0x39, 0xb9, 0x50, 0x95, 0x97, 0xa4, 0x95, 0xf9, 0x25, 0xe1, 0x7e, 0xef, 0xa7,
	0x8a, 0x4b, 0xfe, 0x65, 0x8d, 0x98, 0x58, 0xea, 0x95, 0xf8, 0x63, 0xa3, 0xd9, 0xff, 0xfc, 0x77,
	0xbe, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x81, 0x73, 0x52, 0xfc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceBindingClient is the client API for ResourceBinding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceBindingClient interface {
	AssignResourceLink(ctx context.Context, in *ResourceBindingOpts, opts ...grpc.CallOption) (*ResourceBindings, error)
	RemoveResourceLink(ctx context.Context, in *ResourceBindingOpts, opts ...grpc.CallOption) (*ResourceBindings, error)
}

type resourceBindingClient struct {
	cc *grpc.ClientConn
}

func NewResourceBindingClient(cc *grpc.ClientConn) ResourceBindingClient {
	return &resourceBindingClient{cc}
}

func (c *resourceBindingClient) AssignResourceLink(ctx context.Context, in *ResourceBindingOpts, opts ...grpc.CallOption) (*ResourceBindings, error) {
	out := new(ResourceBindings)
	err := c.cc.Invoke(ctx, "/api.ResourceBinding/AssignResourceLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceBindingClient) RemoveResourceLink(ctx context.Context, in *ResourceBindingOpts, opts ...grpc.CallOption) (*ResourceBindings, error) {
	out := new(ResourceBindings)
	err := c.cc.Invoke(ctx, "/api.ResourceBinding/RemoveResourceLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceBindingServer is the server API for ResourceBinding service.
type ResourceBindingServer interface {
	AssignResourceLink(context.Context, *ResourceBindingOpts) (*ResourceBindings, error)
	RemoveResourceLink(context.Context, *ResourceBindingOpts) (*ResourceBindings, error)
}

func RegisterResourceBindingServer(s *grpc.Server, srv ResourceBindingServer) {
	s.RegisterService(&_ResourceBinding_serviceDesc, srv)
}

func _ResourceBinding_AssignResourceLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceBindingOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceBindingServer).AssignResourceLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ResourceBinding/AssignResourceLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceBindingServer).AssignResourceLink(ctx, req.(*ResourceBindingOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceBinding_RemoveResourceLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceBindingOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceBindingServer).RemoveResourceLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ResourceBinding/RemoveResourceLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceBindingServer).RemoveResourceLink(ctx, req.(*ResourceBindingOpts))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceBinding_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ResourceBinding",
	HandlerType: (*ResourceBindingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignResourceLink",
			Handler:    _ResourceBinding_AssignResourceLink_Handler,
		},
		{
			MethodName: "RemoveResourceLink",
			Handler:    _ResourceBinding_RemoveResourceLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resourcebinding.proto",
}
