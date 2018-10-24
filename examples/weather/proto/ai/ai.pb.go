// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/weather/proto/ai/ai.proto

package cloudprovider_apis_ai

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Question struct {
	Question             string   `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ded88fef1e5bf3, []int{0}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

type Answer struct {
	Answer               string   `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ded88fef1e5bf3, []int{1}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func init() {
	proto.RegisterType((*Question)(nil), "cloudprovider.apis.ai.Question")
	proto.RegisterType((*Answer)(nil), "cloudprovider.apis.ai.Answer")
}

func init() { proto.RegisterFile("examples/weather/proto/ai/ai.proto", fileDescriptor_85ded88fef1e5bf3) }

var fileDescriptor_85ded88fef1e5bf3 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x4f, 0x4d, 0x2c, 0xc9, 0x48, 0x2d, 0xd2, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0xd7, 0x4f, 0xcc, 0xd4, 0x4f, 0xcc, 0xd4, 0x03, 0xb3, 0x85, 0x44, 0x93, 0x73, 0xf2,
	0x4b, 0x53, 0x0a, 0x8a, 0xf2, 0xcb, 0x32, 0x53, 0x52, 0x8b, 0xf4, 0x12, 0x0b, 0x32, 0x8b, 0xf5,
	0x12, 0x33, 0x95, 0xd4, 0xb8, 0x38, 0x02, 0x4b, 0x53, 0x8b, 0x4b, 0x32, 0xf3, 0xf3, 0x84, 0xa4,
	0xb8, 0x38, 0x0a, 0xa1, 0x6c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x49, 0x81,
	0x8b, 0xcd, 0x31, 0xaf, 0xb8, 0x3c, 0xb5, 0x48, 0x48, 0x8c, 0x8b, 0x2d, 0x11, 0xcc, 0x82, 0xaa,
	0x81, 0xf2, 0x8c, 0x7c, 0xb9, 0x98, 0x1c, 0x33, 0x85, 0xdc, 0xb9, 0x98, 0x1d, 0x8b, 0xb3, 0x85,
	0xe4, 0xf5, 0xb0, 0x5a, 0xa7, 0x07, 0xb3, 0x4b, 0x4a, 0x16, 0x87, 0x02, 0x88, 0x25, 0x4a, 0x0c,
	0x49, 0x6c, 0x60, 0x67, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x37, 0xe8, 0x16, 0x57, 0xdc,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AiClient is the client API for Ai service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AiClient interface {
	Ask(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Answer, error)
}

type aiClient struct {
	cc *grpc.ClientConn
}

func NewAiClient(cc *grpc.ClientConn) AiClient {
	return &aiClient{cc}
}

func (c *aiClient) Ask(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/cloudprovider.apis.ai.Ai/Ask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AiServer is the server API for Ai service.
type AiServer interface {
	Ask(context.Context, *Question) (*Answer, error)
}

func RegisterAiServer(s *grpc.Server, srv AiServer) {
	s.RegisterService(&_Ai_serviceDesc, srv)
}

func _Ai_Ask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Question)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiServer).Ask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprovider.apis.ai.Ai/Ask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiServer).Ask(ctx, req.(*Question))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ai_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudprovider.apis.ai.Ai",
	HandlerType: (*AiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ask",
			Handler:    _Ai_Ask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/weather/proto/ai/ai.proto",
}
