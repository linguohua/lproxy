// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bw_report.proto

package bwreport

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type BandwidthStatistics struct {
	Statistics           []*BandwidthUsage `protobuf:"bytes,1,rep,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BandwidthStatistics) Reset()         { *m = BandwidthStatistics{} }
func (m *BandwidthStatistics) String() string { return proto.CompactTextString(m) }
func (*BandwidthStatistics) ProtoMessage()    {}
func (*BandwidthStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0a53df4d7a253ee, []int{0}
}

func (m *BandwidthStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BandwidthStatistics.Unmarshal(m, b)
}
func (m *BandwidthStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BandwidthStatistics.Marshal(b, m, deterministic)
}
func (m *BandwidthStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BandwidthStatistics.Merge(m, src)
}
func (m *BandwidthStatistics) XXX_Size() int {
	return xxx_messageInfo_BandwidthStatistics.Size(m)
}
func (m *BandwidthStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_BandwidthStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_BandwidthStatistics proto.InternalMessageInfo

func (m *BandwidthStatistics) GetStatistics() []*BandwidthUsage {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type BandwidthUsage struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	SendBytes            uint64   `protobuf:"varint,2,opt,name=send_bytes,json=sendBytes,proto3" json:"send_bytes,omitempty"`
	RecvBytes            uint64   `protobuf:"varint,3,opt,name=recv_bytes,json=recvBytes,proto3" json:"recv_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BandwidthUsage) Reset()         { *m = BandwidthUsage{} }
func (m *BandwidthUsage) String() string { return proto.CompactTextString(m) }
func (*BandwidthUsage) ProtoMessage()    {}
func (*BandwidthUsage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0a53df4d7a253ee, []int{1}
}

func (m *BandwidthUsage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BandwidthUsage.Unmarshal(m, b)
}
func (m *BandwidthUsage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BandwidthUsage.Marshal(b, m, deterministic)
}
func (m *BandwidthUsage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BandwidthUsage.Merge(m, src)
}
func (m *BandwidthUsage) XXX_Size() int {
	return xxx_messageInfo_BandwidthUsage.Size(m)
}
func (m *BandwidthUsage) XXX_DiscardUnknown() {
	xxx_messageInfo_BandwidthUsage.DiscardUnknown(m)
}

var xxx_messageInfo_BandwidthUsage proto.InternalMessageInfo

func (m *BandwidthUsage) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *BandwidthUsage) GetSendBytes() uint64 {
	if m != nil {
		return m.SendBytes
	}
	return 0
}

func (m *BandwidthUsage) GetRecvBytes() uint64 {
	if m != nil {
		return m.RecvBytes
	}
	return 0
}

type ReportResult struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportResult) Reset()         { *m = ReportResult{} }
func (m *ReportResult) String() string { return proto.CompactTextString(m) }
func (*ReportResult) ProtoMessage()    {}
func (*ReportResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0a53df4d7a253ee, []int{2}
}

func (m *ReportResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportResult.Unmarshal(m, b)
}
func (m *ReportResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportResult.Marshal(b, m, deterministic)
}
func (m *ReportResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportResult.Merge(m, src)
}
func (m *ReportResult) XXX_Size() int {
	return xxx_messageInfo_ReportResult.Size(m)
}
func (m *ReportResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportResult.DiscardUnknown(m)
}

var xxx_messageInfo_ReportResult proto.InternalMessageInfo

func (m *ReportResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*BandwidthStatistics)(nil), "BandwidthStatistics")
	proto.RegisterType((*BandwidthUsage)(nil), "BandwidthUsage")
	proto.RegisterType((*ReportResult)(nil), "ReportResult")
}

func init() { proto.RegisterFile("bw_report.proto", fileDescriptor_e0a53df4d7a253ee) }

var fileDescriptor_e0a53df4d7a253ee = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xdf, 0x4a, 0x87, 0x30,
	0x1c, 0xc5, 0x5b, 0x9a, 0xe0, 0xb7, 0x3f, 0xc2, 0xea, 0x42, 0x82, 0x40, 0x76, 0xe5, 0xd5, 0x04,
	0x7b, 0x03, 0x2f, 0x7a, 0x80, 0x45, 0xd7, 0xe2, 0xdc, 0xa8, 0x41, 0x38, 0xd9, 0xbe, 0x26, 0xbd,
	0x7d, 0x6c, 0x8a, 0x18, 0xfc, 0xee, 0xce, 0xce, 0xe7, 0xb0, 0x73, 0x36, 0x28, 0xe4, 0xda, 0x3b,
	0x3d, 0x5b, 0x87, 0x7c, 0x76, 0x16, 0x2d, 0x7b, 0x83, 0xc7, 0x6e, 0x98, 0xd4, 0x6a, 0x14, 0x7e,
	0xbd, 0xe3, 0x80, 0xc6, 0xa3, 0x19, 0x3d, 0x6d, 0x00, 0xfc, 0x71, 0x2a, 0x49, 0x95, 0xd4, 0xb7,
	0x6d, 0xc1, 0x8f, 0xe4, 0x87, 0x1f, 0x3e, 0xb5, 0x38, 0x45, 0x98, 0x84, 0x87, 0xff, 0x94, 0x52,
	0x48, 0x97, 0xc5, 0xa8, 0x92, 0x54, 0xa4, 0xce, 0x45, 0xd4, 0xf4, 0x05, 0xc0, 0xeb, 0x49, 0xf5,
	0xf2, 0x17, 0xb5, 0x2f, 0xaf, 0x2b, 0x52, 0xa7, 0x22, 0x0f, 0x4e, 0x17, 0x8c, 0x80, 0x9d, 0x1e,
	0x7f, 0x76, 0x9c, 0x6c, 0x38, 0x38, 0x11, 0x33, 0x06, 0x77, 0x22, 0x6e, 0x17, 0xda, 0x2f, 0xdf,
	0x18, 0x1a, 0x46, 0xab, 0x74, 0x6c, 0xb8, 0x11, 0x51, 0xb7, 0x1d, 0x14, 0xc7, 0x8e, 0x2d, 0x4c,
	0x1b, 0xc8, 0x76, 0xf5, 0xc4, 0x2f, 0xbc, 0xf5, 0xf9, 0x9e, 0x9f, 0x6f, 0x65, 0x57, 0x32, 0x8b,
	0x5f, 0xf3, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xcf, 0xd7, 0x5f, 0x2d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BandwidthReportClient is the client API for BandwidthReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BandwidthReportClient interface {
	Report(ctx context.Context, in *BandwidthStatistics, opts ...grpc.CallOption) (*ReportResult, error)
}

type bandwidthReportClient struct {
	cc *grpc.ClientConn
}

func NewBandwidthReportClient(cc *grpc.ClientConn) BandwidthReportClient {
	return &bandwidthReportClient{cc}
}

func (c *bandwidthReportClient) Report(ctx context.Context, in *BandwidthStatistics, opts ...grpc.CallOption) (*ReportResult, error) {
	out := new(ReportResult)
	err := c.cc.Invoke(ctx, "/BandwidthReport/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BandwidthReportServer is the server API for BandwidthReport service.
type BandwidthReportServer interface {
	Report(context.Context, *BandwidthStatistics) (*ReportResult, error)
}

// UnimplementedBandwidthReportServer can be embedded to have forward compatible implementations.
type UnimplementedBandwidthReportServer struct {
}

func (*UnimplementedBandwidthReportServer) Report(ctx context.Context, req *BandwidthStatistics) (*ReportResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}

func RegisterBandwidthReportServer(s *grpc.Server, srv BandwidthReportServer) {
	s.RegisterService(&_BandwidthReport_serviceDesc, srv)
}

func _BandwidthReport_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BandwidthStatistics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BandwidthReportServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BandwidthReport/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BandwidthReportServer).Report(ctx, req.(*BandwidthStatistics))
	}
	return interceptor(ctx, in, info, handler)
}

var _BandwidthReport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BandwidthReport",
	HandlerType: (*BandwidthReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _BandwidthReport_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bw_report.proto",
}