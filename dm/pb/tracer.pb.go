// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tracer.proto

package pb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetTSORequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetTSORequest) Reset()         { *m = GetTSORequest{} }
func (m *GetTSORequest) String() string { return proto.CompactTextString(m) }
func (*GetTSORequest) ProtoMessage()    {}
func (*GetTSORequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{0}
}
func (m *GetTSORequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTSORequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTSORequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTSORequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTSORequest.Merge(m, src)
}
func (m *GetTSORequest) XXX_Size() int {
	return m.Size()
}
func (m *GetTSORequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTSORequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTSORequest proto.InternalMessageInfo

func (m *GetTSORequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTSOResponse struct {
	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ts     int64  `protobuf:"varint,3,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (m *GetTSOResponse) Reset()         { *m = GetTSOResponse{} }
func (m *GetTSOResponse) String() string { return proto.CompactTextString(m) }
func (*GetTSOResponse) ProtoMessage()    {}
func (*GetTSOResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{1}
}
func (m *GetTSOResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTSOResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTSOResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTSOResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTSOResponse.Merge(m, src)
}
func (m *GetTSOResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTSOResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTSOResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTSOResponse proto.InternalMessageInfo

func (m *GetTSOResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *GetTSOResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetTSOResponse) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type CommonUploadResponse struct {
	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *CommonUploadResponse) Reset()         { *m = CommonUploadResponse{} }
func (m *CommonUploadResponse) String() string { return proto.CompactTextString(m) }
func (*CommonUploadResponse) ProtoMessage()    {}
func (*CommonUploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{2}
}
func (m *CommonUploadResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonUploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommonUploadResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommonUploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonUploadResponse.Merge(m, src)
}
func (m *CommonUploadResponse) XXX_Size() int {
	return m.Size()
}
func (m *CommonUploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonUploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonUploadResponse proto.InternalMessageInfo

func (m *CommonUploadResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *CommonUploadResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UploadSyncerBinlogEventRequest struct {
	Events []*SyncerBinlogEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *UploadSyncerBinlogEventRequest) Reset()         { *m = UploadSyncerBinlogEventRequest{} }
func (m *UploadSyncerBinlogEventRequest) String() string { return proto.CompactTextString(m) }
func (*UploadSyncerBinlogEventRequest) ProtoMessage()    {}
func (*UploadSyncerBinlogEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{3}
}
func (m *UploadSyncerBinlogEventRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploadSyncerBinlogEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploadSyncerBinlogEventRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploadSyncerBinlogEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadSyncerBinlogEventRequest.Merge(m, src)
}
func (m *UploadSyncerBinlogEventRequest) XXX_Size() int {
	return m.Size()
}
func (m *UploadSyncerBinlogEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadSyncerBinlogEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadSyncerBinlogEventRequest proto.InternalMessageInfo

func (m *UploadSyncerBinlogEventRequest) GetEvents() []*SyncerBinlogEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type UploadSyncerJobEventRequest struct {
	Events []*SyncerJobEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *UploadSyncerJobEventRequest) Reset()         { *m = UploadSyncerJobEventRequest{} }
func (m *UploadSyncerJobEventRequest) String() string { return proto.CompactTextString(m) }
func (*UploadSyncerJobEventRequest) ProtoMessage()    {}
func (*UploadSyncerJobEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{4}
}
func (m *UploadSyncerJobEventRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploadSyncerJobEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploadSyncerJobEventRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploadSyncerJobEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadSyncerJobEventRequest.Merge(m, src)
}
func (m *UploadSyncerJobEventRequest) XXX_Size() int {
	return m.Size()
}
func (m *UploadSyncerJobEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadSyncerJobEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadSyncerJobEventRequest proto.InternalMessageInfo

func (m *UploadSyncerJobEventRequest) GetEvents() []*SyncerJobEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTSORequest)(nil), "pb.GetTSORequest")
	proto.RegisterType((*GetTSOResponse)(nil), "pb.GetTSOResponse")
	proto.RegisterType((*CommonUploadResponse)(nil), "pb.CommonUploadResponse")
	proto.RegisterType((*UploadSyncerBinlogEventRequest)(nil), "pb.UploadSyncerBinlogEventRequest")
	proto.RegisterType((*UploadSyncerJobEventRequest)(nil), "pb.UploadSyncerJobEventRequest")
}

func init() { proto.RegisterFile("tracer.proto", fileDescriptor_6d422d7c66fbbd8f) }

var fileDescriptor_6d422d7c66fbbd8f = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x9b, 0x16, 0x8a, 0x7e, 0xea, 0xd0, 0xcf, 0xa9, 0x61, 0x42, 0x36, 0x72, 0x1a, 0x82,
	0x03, 0xe7, 0x1f, 0x90, 0x89, 0x88, 0xbb, 0x0c, 0xba, 0x79, 0xf0, 0x24, 0xab, 0x0b, 0x63, 0xb0,
	0x35, 0xb5, 0xc9, 0x04, 0xff, 0x85, 0x3f, 0xcb, 0xe3, 0x8e, 0x1e, 0xa5, 0x3d, 0xfa, 0x27, 0x24,
	0x6d, 0x37, 0xb7, 0x3a, 0x07, 0xde, 0xf2, 0x7d, 0x79, 0xdf, 0x87, 0xbc, 0x2f, 0x81, 0x5d, 0x1d,
	0xf5, 0x9f, 0x44, 0xd4, 0x08, 0x23, 0xa9, 0x25, 0xda, 0xa1, 0x5f, 0x39, 0xcc, 0x36, 0x8f, 0xea,
	0x35, 0x58, 0x5c, 0xf0, 0x2a, 0xec, 0xdd, 0x0a, 0xdd, 0xeb, 0x76, 0x3c, 0xf1, 0x3c, 0x15, 0x4a,
	0x63, 0x09, 0xec, 0xd1, 0x80, 0x92, 0x1a, 0xa9, 0x6f, 0x7b, 0xf6, 0x68, 0xc0, 0xdb, 0x50, 0x9a,
	0x0b, 0x54, 0x28, 0x03, 0x25, 0xf0, 0x18, 0xdc, 0x48, 0xa8, 0xe9, 0x58, 0xa7, 0xaa, 0x2d, 0x2f,
	0x9f, 0x70, 0x1f, 0x9c, 0x89, 0x1a, 0x52, 0x3b, 0xb5, 0x9a, 0xa3, 0x61, 0x69, 0x45, 0x9d, 0x1a,
	0xa9, 0x3b, 0x9e, 0xad, 0x15, 0xbf, 0x82, 0xf2, 0xb5, 0x9c, 0x4c, 0x64, 0x70, 0x1f, 0x8e, 0x65,
	0x7f, 0xf0, 0x7f, 0x22, 0xef, 0x00, 0xcb, 0xbc, 0xdd, 0x34, 0x44, 0x6b, 0x14, 0x8c, 0xe5, 0xf0,
	0xe6, 0x45, 0x04, 0x7a, 0xfe, 0xfe, 0x73, 0x70, 0x85, 0x99, 0x15, 0x25, 0x35, 0xa7, 0xbe, 0xd3,
	0x3c, 0x6a, 0x84, 0x7e, 0xe3, 0xb7, 0x3a, 0x17, 0xf1, 0x3b, 0x38, 0x5d, 0x06, 0xb6, 0xa5, 0xbf,
	0x42, 0x3b, 0x2b, 0xd0, 0xf0, 0x87, 0xb6, 0x90, 0xe6, 0x8a, 0xe6, 0x17, 0x01, 0xb7, 0x97, 0x56,
	0x8c, 0x17, 0xe0, 0x66, 0xa5, 0xe1, 0x81, 0x31, 0xac, 0x34, 0x5c, 0xc1, 0xe5, 0x55, 0xd6, 0x00,
	0xb7, 0xf0, 0x01, 0x4e, 0xfe, 0x48, 0x86, 0xdc, 0x18, 0x36, 0xc7, 0xae, 0x50, 0xa3, 0x59, 0x57,
	0x2e, 0xb7, 0xb0, 0x0b, 0xe5, 0x75, 0x19, 0xb1, 0x5a, 0xe4, 0x16, 0xd2, 0x6f, 0x82, 0xb6, 0xe8,
	0x7b, 0xcc, 0xc8, 0x2c, 0x66, 0xe4, 0x33, 0x66, 0xe4, 0x2d, 0x61, 0xd6, 0x2c, 0x61, 0xd6, 0x47,
	0xc2, 0x2c, 0xdf, 0x4d, 0x7f, 0xd6, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x04, 0x0e,
	0x15, 0x82, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TracerClient is the client API for Tracer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TracerClient interface {
	GetTSO(ctx context.Context, in *GetTSORequest, opts ...grpc.CallOption) (*GetTSOResponse, error)
	UploadSyncerBinlogEvent(ctx context.Context, in *UploadSyncerBinlogEventRequest, opts ...grpc.CallOption) (*CommonUploadResponse, error)
	UploadSyncerJobEvent(ctx context.Context, in *UploadSyncerJobEventRequest, opts ...grpc.CallOption) (*CommonUploadResponse, error)
}

type tracerClient struct {
	cc *grpc.ClientConn
}

func NewTracerClient(cc *grpc.ClientConn) TracerClient {
	return &tracerClient{cc}
}

func (c *tracerClient) GetTSO(ctx context.Context, in *GetTSORequest, opts ...grpc.CallOption) (*GetTSOResponse, error) {
	out := new(GetTSOResponse)
	err := c.cc.Invoke(ctx, "/pb.Tracer/GetTSO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracerClient) UploadSyncerBinlogEvent(ctx context.Context, in *UploadSyncerBinlogEventRequest, opts ...grpc.CallOption) (*CommonUploadResponse, error) {
	out := new(CommonUploadResponse)
	err := c.cc.Invoke(ctx, "/pb.Tracer/UploadSyncerBinlogEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracerClient) UploadSyncerJobEvent(ctx context.Context, in *UploadSyncerJobEventRequest, opts ...grpc.CallOption) (*CommonUploadResponse, error) {
	out := new(CommonUploadResponse)
	err := c.cc.Invoke(ctx, "/pb.Tracer/UploadSyncerJobEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TracerServer is the server API for Tracer service.
type TracerServer interface {
	GetTSO(context.Context, *GetTSORequest) (*GetTSOResponse, error)
	UploadSyncerBinlogEvent(context.Context, *UploadSyncerBinlogEventRequest) (*CommonUploadResponse, error)
	UploadSyncerJobEvent(context.Context, *UploadSyncerJobEventRequest) (*CommonUploadResponse, error)
}

// UnimplementedTracerServer can be embedded to have forward compatible implementations.
type UnimplementedTracerServer struct {
}

func (*UnimplementedTracerServer) GetTSO(ctx context.Context, req *GetTSORequest) (*GetTSOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTSO not implemented")
}
func (*UnimplementedTracerServer) UploadSyncerBinlogEvent(ctx context.Context, req *UploadSyncerBinlogEventRequest) (*CommonUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadSyncerBinlogEvent not implemented")
}
func (*UnimplementedTracerServer) UploadSyncerJobEvent(ctx context.Context, req *UploadSyncerJobEventRequest) (*CommonUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadSyncerJobEvent not implemented")
}

func RegisterTracerServer(s *grpc.Server, srv TracerServer) {
	s.RegisterService(&_Tracer_serviceDesc, srv)
}

func _Tracer_GetTSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTSORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServer).GetTSO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Tracer/GetTSO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServer).GetTSO(ctx, req.(*GetTSORequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracer_UploadSyncerBinlogEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSyncerBinlogEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServer).UploadSyncerBinlogEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Tracer/UploadSyncerBinlogEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServer).UploadSyncerBinlogEvent(ctx, req.(*UploadSyncerBinlogEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracer_UploadSyncerJobEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSyncerJobEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServer).UploadSyncerJobEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Tracer/UploadSyncerJobEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServer).UploadSyncerJobEvent(ctx, req.(*UploadSyncerJobEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Tracer",
	HandlerType: (*TracerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTSO",
			Handler:    _Tracer_GetTSO_Handler,
		},
		{
			MethodName: "UploadSyncerBinlogEvent",
			Handler:    _Tracer_UploadSyncerBinlogEvent_Handler,
		},
		{
			MethodName: "UploadSyncerJobEvent",
			Handler:    _Tracer_UploadSyncerJobEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracer.proto",
}

func (m *GetTSORequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTSORequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTSORequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTracer(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetTSOResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTSOResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTSOResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ts != 0 {
		i = encodeVarintTracer(dAtA, i, uint64(m.Ts))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintTracer(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Result {
		i--
		if m.Result {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CommonUploadResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonUploadResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonUploadResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintTracer(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Result {
		i--
		if m.Result {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UploadSyncerBinlogEventRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadSyncerBinlogEventRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploadSyncerBinlogEventRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTracer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UploadSyncerJobEventRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadSyncerJobEventRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploadSyncerJobEventRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTracer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTracer(dAtA []byte, offset int, v uint64) int {
	offset -= sovTracer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetTSORequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTracer(uint64(l))
	}
	return n
}

func (m *GetTSOResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result {
		n += 2
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovTracer(uint64(l))
	}
	if m.Ts != 0 {
		n += 1 + sovTracer(uint64(m.Ts))
	}
	return n
}

func (m *CommonUploadResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result {
		n += 2
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovTracer(uint64(l))
	}
	return n
}

func (m *UploadSyncerBinlogEventRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTracer(uint64(l))
		}
	}
	return n
}

func (m *UploadSyncerJobEventRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTracer(uint64(l))
		}
	}
	return n
}

func sovTracer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTracer(x uint64) (n int) {
	return sovTracer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetTSORequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
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
			return fmt.Errorf("proto: GetTSORequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTSORequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
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
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTracer
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
func (m *GetTSOResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
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
			return fmt.Errorf("proto: GetTSOResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTSOResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
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
			m.Result = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
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
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ts", wireType)
			}
			m.Ts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ts |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTracer
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
func (m *CommonUploadResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
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
			return fmt.Errorf("proto: CommonUploadResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonUploadResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
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
			m.Result = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
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
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTracer
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
func (m *UploadSyncerBinlogEventRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
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
			return fmt.Errorf("proto: UploadSyncerBinlogEventRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadSyncerBinlogEventRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
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
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &SyncerBinlogEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTracer
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
func (m *UploadSyncerJobEventRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
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
			return fmt.Errorf("proto: UploadSyncerJobEventRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadSyncerJobEventRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
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
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &SyncerJobEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTracer
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
func skipTracer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTracer
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
					return 0, ErrIntOverflowTracer
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
					return 0, ErrIntOverflowTracer
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
				return 0, ErrInvalidLengthTracer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTracer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTracer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTracer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTracer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTracer = fmt.Errorf("proto: unexpected end of group")
)
