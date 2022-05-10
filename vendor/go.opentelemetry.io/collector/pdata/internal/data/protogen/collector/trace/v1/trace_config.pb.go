// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/trace/v1/trace_config.proto

package v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// How spans should be sampled:
// - Always off
// - Always on
// - Always follow the parent Span's decision (off if no parent).
type ConstantSampler_ConstantDecision int32

const (
	ConstantSampler_ALWAYS_OFF    ConstantSampler_ConstantDecision = 0
	ConstantSampler_ALWAYS_ON     ConstantSampler_ConstantDecision = 1
	ConstantSampler_ALWAYS_PARENT ConstantSampler_ConstantDecision = 2
)

var ConstantSampler_ConstantDecision_name = map[int32]string{
	0: "ALWAYS_OFF",
	1: "ALWAYS_ON",
	2: "ALWAYS_PARENT",
}

var ConstantSampler_ConstantDecision_value = map[string]int32{
	"ALWAYS_OFF":    0,
	"ALWAYS_ON":     1,
	"ALWAYS_PARENT": 2,
}

func (x ConstantSampler_ConstantDecision) String() string {
	return proto.EnumName(ConstantSampler_ConstantDecision_name, int32(x))
}

func (ConstantSampler_ConstantDecision) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{1, 0}
}

// Global configuration of the trace service. All fields must be specified, or
// the default (zero) values will be used for each type.
type TraceConfig struct {
	// The global default sampler used to make decisions on span sampling.
	//
	// Types that are valid to be assigned to Sampler:
	//	*TraceConfig_ConstantSampler
	//	*TraceConfig_TraceIdRatioBased
	//	*TraceConfig_RateLimitingSampler
	Sampler isTraceConfig_Sampler `protobuf_oneof:"sampler"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributes int64 `protobuf:"varint,4,opt,name=max_number_of_attributes,json=maxNumberOfAttributes,proto3" json:"max_number_of_attributes,omitempty"`
	// The global default max number of annotation events per span.
	MaxNumberOfTimedEvents int64 `protobuf:"varint,5,opt,name=max_number_of_timed_events,json=maxNumberOfTimedEvents,proto3" json:"max_number_of_timed_events,omitempty"`
	// The global default max number of attributes per timed event.
	MaxNumberOfAttributesPerTimedEvent int64 `protobuf:"varint,6,opt,name=max_number_of_attributes_per_timed_event,json=maxNumberOfAttributesPerTimedEvent,proto3" json:"max_number_of_attributes_per_timed_event,omitempty"`
	// The global default max number of link entries per span.
	MaxNumberOfLinks int64 `protobuf:"varint,7,opt,name=max_number_of_links,json=maxNumberOfLinks,proto3" json:"max_number_of_links,omitempty"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributesPerLink int64 `protobuf:"varint,8,opt,name=max_number_of_attributes_per_link,json=maxNumberOfAttributesPerLink,proto3" json:"max_number_of_attributes_per_link,omitempty"`
}

func (m *TraceConfig) Reset()         { *m = TraceConfig{} }
func (m *TraceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceConfig) ProtoMessage()    {}
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{0}
}
func (m *TraceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceConfig.Merge(m, src)
}
func (m *TraceConfig) XXX_Size() int {
	return m.Size()
}
func (m *TraceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceConfig proto.InternalMessageInfo

type isTraceConfig_Sampler interface {
	isTraceConfig_Sampler()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TraceConfig_ConstantSampler struct {
	ConstantSampler *ConstantSampler `protobuf:"bytes,1,opt,name=constant_sampler,json=constantSampler,proto3,oneof" json:"constant_sampler,omitempty"`
}
type TraceConfig_TraceIdRatioBased struct {
	TraceIdRatioBased *TraceIdRatioBased `protobuf:"bytes,2,opt,name=trace_id_ratio_based,json=traceIdRatioBased,proto3,oneof" json:"trace_id_ratio_based,omitempty"`
}
type TraceConfig_RateLimitingSampler struct {
	RateLimitingSampler *RateLimitingSampler `protobuf:"bytes,3,opt,name=rate_limiting_sampler,json=rateLimitingSampler,proto3,oneof" json:"rate_limiting_sampler,omitempty"`
}

func (*TraceConfig_ConstantSampler) isTraceConfig_Sampler()     {}
func (*TraceConfig_TraceIdRatioBased) isTraceConfig_Sampler()   {}
func (*TraceConfig_RateLimitingSampler) isTraceConfig_Sampler() {}

func (m *TraceConfig) GetSampler() isTraceConfig_Sampler {
	if m != nil {
		return m.Sampler
	}
	return nil
}

func (m *TraceConfig) GetConstantSampler() *ConstantSampler {
	if x, ok := m.GetSampler().(*TraceConfig_ConstantSampler); ok {
		return x.ConstantSampler
	}
	return nil
}

func (m *TraceConfig) GetTraceIdRatioBased() *TraceIdRatioBased {
	if x, ok := m.GetSampler().(*TraceConfig_TraceIdRatioBased); ok {
		return x.TraceIdRatioBased
	}
	return nil
}

func (m *TraceConfig) GetRateLimitingSampler() *RateLimitingSampler {
	if x, ok := m.GetSampler().(*TraceConfig_RateLimitingSampler); ok {
		return x.RateLimitingSampler
	}
	return nil
}

func (m *TraceConfig) GetMaxNumberOfAttributes() int64 {
	if m != nil {
		return m.MaxNumberOfAttributes
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfTimedEvents() int64 {
	if m != nil {
		return m.MaxNumberOfTimedEvents
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAttributesPerTimedEvent() int64 {
	if m != nil {
		return m.MaxNumberOfAttributesPerTimedEvent
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfLinks() int64 {
	if m != nil {
		return m.MaxNumberOfLinks
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAttributesPerLink() int64 {
	if m != nil {
		return m.MaxNumberOfAttributesPerLink
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TraceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TraceConfig_ConstantSampler)(nil),
		(*TraceConfig_TraceIdRatioBased)(nil),
		(*TraceConfig_RateLimitingSampler)(nil),
	}
}

// Sampler that always makes a constant decision on span sampling.
type ConstantSampler struct {
	Decision ConstantSampler_ConstantDecision `protobuf:"varint,1,opt,name=decision,proto3,enum=opentelemetry.proto.trace.v1.ConstantSampler_ConstantDecision" json:"decision,omitempty"`
}

func (m *ConstantSampler) Reset()         { *m = ConstantSampler{} }
func (m *ConstantSampler) String() string { return proto.CompactTextString(m) }
func (*ConstantSampler) ProtoMessage()    {}
func (*ConstantSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{1}
}
func (m *ConstantSampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConstantSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConstantSampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConstantSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConstantSampler.Merge(m, src)
}
func (m *ConstantSampler) XXX_Size() int {
	return m.Size()
}
func (m *ConstantSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ConstantSampler.DiscardUnknown(m)
}

var xxx_messageInfo_ConstantSampler proto.InternalMessageInfo

func (m *ConstantSampler) GetDecision() ConstantSampler_ConstantDecision {
	if m != nil {
		return m.Decision
	}
	return ConstantSampler_ALWAYS_OFF
}

// Sampler that tries to uniformly sample traces with a given ratio.
// The ratio of sampling a trace is equal to that of the specified ratio.
type TraceIdRatioBased struct {
	// The desired ratio of sampling. Must be within [0.0, 1.0].
	SamplingRatio float64 `protobuf:"fixed64,1,opt,name=samplingRatio,proto3" json:"samplingRatio,omitempty"`
}

func (m *TraceIdRatioBased) Reset()         { *m = TraceIdRatioBased{} }
func (m *TraceIdRatioBased) String() string { return proto.CompactTextString(m) }
func (*TraceIdRatioBased) ProtoMessage()    {}
func (*TraceIdRatioBased) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{2}
}
func (m *TraceIdRatioBased) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceIdRatioBased) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceIdRatioBased.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceIdRatioBased) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceIdRatioBased.Merge(m, src)
}
func (m *TraceIdRatioBased) XXX_Size() int {
	return m.Size()
}
func (m *TraceIdRatioBased) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceIdRatioBased.DiscardUnknown(m)
}

var xxx_messageInfo_TraceIdRatioBased proto.InternalMessageInfo

func (m *TraceIdRatioBased) GetSamplingRatio() float64 {
	if m != nil {
		return m.SamplingRatio
	}
	return 0
}

// Sampler that tries to sample with a rate per time window.
type RateLimitingSampler struct {
	// Rate per second.
	Qps int64 `protobuf:"varint,1,opt,name=qps,proto3" json:"qps,omitempty"`
}

func (m *RateLimitingSampler) Reset()         { *m = RateLimitingSampler{} }
func (m *RateLimitingSampler) String() string { return proto.CompactTextString(m) }
func (*RateLimitingSampler) ProtoMessage()    {}
func (*RateLimitingSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5936aa8fa6443e6f, []int{3}
}
func (m *RateLimitingSampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitingSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitingSampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimitingSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitingSampler.Merge(m, src)
}
func (m *RateLimitingSampler) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitingSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitingSampler.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitingSampler proto.InternalMessageInfo

func (m *RateLimitingSampler) GetQps() int64 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func init() {
	proto.RegisterEnum("opentelemetry.proto.trace.v1.ConstantSampler_ConstantDecision", ConstantSampler_ConstantDecision_name, ConstantSampler_ConstantDecision_value)
	proto.RegisterType((*TraceConfig)(nil), "opentelemetry.proto.trace.v1.TraceConfig")
	proto.RegisterType((*ConstantSampler)(nil), "opentelemetry.proto.trace.v1.ConstantSampler")
	proto.RegisterType((*TraceIdRatioBased)(nil), "opentelemetry.proto.trace.v1.TraceIdRatioBased")
	proto.RegisterType((*RateLimitingSampler)(nil), "opentelemetry.proto.trace.v1.RateLimitingSampler")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/trace/v1/trace_config.proto", fileDescriptor_5936aa8fa6443e6f)
}

var fileDescriptor_5936aa8fa6443e6f = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0xed, 0x86, 0xfe, 0xdd, 0xaa, 0xad, 0x3b, 0x6d, 0x91, 0x55, 0x55, 0xa6, 0x58, 0x48,
	0x64, 0xd3, 0x58, 0x29, 0x0b, 0x04, 0x0b, 0xa4, 0xa4, 0x3f, 0x14, 0x29, 0xa4, 0x91, 0x1b, 0x09,
	0x91, 0xcd, 0x68, 0x62, 0x4f, 0xac, 0x11, 0xf6, 0x8c, 0x19, 0x4f, 0xa3, 0xb2, 0xe7, 0x01, 0x78,
	0x00, 0xde, 0x82, 0x97, 0x60, 0xd9, 0x25, 0x4b, 0x94, 0xbc, 0x08, 0xf2, 0x38, 0x4d, 0x93, 0xb4,
	0x8d, 0xc4, 0x6e, 0xee, 0x39, 0x73, 0xbe, 0x7b, 0x27, 0xb9, 0x09, 0x78, 0x22, 0xa5, 0x5c, 0xd1,
	0x98, 0x26, 0x54, 0xc9, 0x6f, 0x5e, 0x2a, 0x85, 0x12, 0x9e, 0x92, 0x24, 0xa0, 0x5e, 0xbf, 0x5a,
	0x1c, 0x70, 0x20, 0x78, 0x8f, 0x45, 0x15, 0xed, 0xa1, 0xfd, 0xa9, 0x40, 0x21, 0x56, 0xf4, 0xbd,
	0x4a, 0xbf, 0xba, 0xb7, 0x13, 0x89, 0x48, 0x14, 0x90, 0xfc, 0x54, 0xd8, 0xee, 0xf7, 0x45, 0x58,
	0x6b, 0xe7, 0x57, 0x8e, 0x35, 0x09, 0x75, 0xc0, 0x0a, 0x04, 0xcf, 0x14, 0xe1, 0x0a, 0x67, 0x24,
	0x49, 0x63, 0x2a, 0x6d, 0xf3, 0xc0, 0x2c, 0xaf, 0x1d, 0x1d, 0x56, 0xe6, 0xe1, 0x2b, 0xc7, 0xa3,
	0xd4, 0x65, 0x11, 0x3a, 0x37, 0xfc, 0xcd, 0x60, 0x5a, 0x42, 0x5d, 0xd8, 0x29, 0xa6, 0x66, 0x21,
	0x96, 0x44, 0x31, 0x81, 0xbb, 0x24, 0xa3, 0xa1, 0xbd, 0xa0, 0xf9, 0xde, 0x7c, 0xbe, 0x1e, 0xf2,
	0x43, 0xe8, 0xe7, 0xb9, 0x7a, 0x1e, 0x3b, 0x37, 0xfc, 0x2d, 0x35, 0x2b, 0xa2, 0x08, 0x76, 0x25,
	0x51, 0x14, 0xc7, 0x2c, 0x61, 0x8a, 0xf1, 0x68, 0xfc, 0x88, 0x92, 0x6e, 0x52, 0x9d, 0xdf, 0xc4,
	0x27, 0x8a, 0x36, 0x46, 0xc9, 0xbb, 0x87, 0x6c, 0xcb, 0xfb, 0x32, 0x7a, 0x0d, 0x76, 0x42, 0xae,
	0x31, 0xbf, 0x4a, 0xba, 0x54, 0x62, 0xd1, 0xc3, 0x44, 0x29, 0xc9, 0xba, 0x57, 0x8a, 0x66, 0xf6,
	0x93, 0x03, 0xb3, 0x5c, 0xf2, 0x77, 0x13, 0x72, 0xdd, 0xd4, 0xf6, 0x45, 0xaf, 0x36, 0x36, 0xd1,
	0x5b, 0xd8, 0x9b, 0x0e, 0x2a, 0x96, 0xd0, 0x10, 0xd3, 0x3e, 0xe5, 0x2a, 0xb3, 0x17, 0x75, 0xf4,
	0xe9, 0x44, 0xb4, 0x9d, 0xdb, 0xa7, 0xda, 0x45, 0x6d, 0x28, 0x3f, 0xd6, 0x14, 0xa7, 0x54, 0x4e,
	0xa2, 0xec, 0x25, 0x4d, 0x72, 0x1f, 0x1c, 0xa2, 0x45, 0xe5, 0x1d, 0x16, 0x1d, 0xc2, 0xf6, 0x34,
	0x35, 0x66, 0xfc, 0x4b, 0x66, 0x2f, 0x6b, 0x80, 0x35, 0x01, 0x68, 0xe4, 0x3a, 0x7a, 0x0f, 0xcf,
	0xe7, 0x0e, 0x91, 0xa7, 0xed, 0x15, 0x1d, 0xde, 0x7f, 0xac, 0x7b, 0x4e, 0xaa, 0xaf, 0xc2, 0xf2,
	0xe8, 0xdb, 0x71, 0x7f, 0x99, 0xb0, 0x39, 0xb3, 0x41, 0xa8, 0x03, 0x2b, 0x21, 0x0d, 0x58, 0xc6,
	0x04, 0xd7, 0x2b, 0xb8, 0x71, 0xf4, 0xee, 0xbf, 0x56, 0x70, 0x5c, 0x9f, 0x8c, 0x28, 0xfe, 0x98,
	0xe7, 0x9e, 0x80, 0x35, 0xeb, 0xa2, 0x0d, 0x80, 0x5a, 0xe3, 0x53, 0xed, 0xf3, 0x25, 0xbe, 0x38,
	0x3b, 0xb3, 0x0c, 0xb4, 0x0e, 0xab, 0xb7, 0x75, 0xd3, 0x32, 0xd1, 0x16, 0xac, 0x8f, 0xca, 0x56,
	0xcd, 0x3f, 0x6d, 0xb6, 0xad, 0x05, 0xf7, 0x0d, 0x6c, 0xdd, 0x5b, 0x4b, 0xf4, 0x02, 0xd6, 0xf5,
	0xab, 0x18, 0x8f, 0xb4, 0xaa, 0x67, 0x37, 0xfd, 0x69, 0xd1, 0x7d, 0x09, 0xdb, 0x0f, 0x2c, 0x1b,
	0xb2, 0xa0, 0xf4, 0x35, 0xcd, 0x74, 0xa4, 0xe4, 0xe7, 0xc7, 0xfa, 0x4f, 0xf3, 0xf7, 0xc0, 0x31,
	0x6f, 0x06, 0x8e, 0xf9, 0x77, 0xe0, 0x98, 0x3f, 0x86, 0x8e, 0x71, 0x33, 0x74, 0x8c, 0x3f, 0x43,
	0xc7, 0x80, 0x67, 0x4c, 0xcc, 0xfd, 0x40, 0xea, 0xd6, 0xc4, 0x2f, 0xbb, 0x95, 0x5b, 0x2d, 0xb3,
	0xf3, 0x31, 0x9a, 0x0d, 0x31, 0xe1, 0x05, 0x22, 0x8e, 0x69, 0xa0, 0x84, 0xf4, 0xd2, 0x90, 0x28,
	0xe2, 0x31, 0xae, 0xa8, 0xe4, 0x24, 0xf6, 0x74, 0xa5, 0xa9, 0x11, 0xe5, 0x13, 0xd7, 0x6e, 0xff,
	0x86, 0xba, 0x4b, 0xda, 0x7c, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xd1, 0x4d, 0x12, 0xad,
	0x04, 0x00, 0x00,
}

func (m *TraceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxNumberOfAttributesPerLink != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAttributesPerLink))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxNumberOfLinks != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfLinks))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxNumberOfAttributesPerTimedEvent != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAttributesPerTimedEvent))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxNumberOfTimedEvents != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfTimedEvents))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxNumberOfAttributes != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAttributes))
		i--
		dAtA[i] = 0x20
	}
	if m.Sampler != nil {
		{
			size := m.Sampler.Size()
			i -= size
			if _, err := m.Sampler.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *TraceConfig_ConstantSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig_ConstantSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ConstantSampler != nil {
		{
			size, err := m.ConstantSampler.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTraceConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *TraceConfig_TraceIdRatioBased) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig_TraceIdRatioBased) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TraceIdRatioBased != nil {
		{
			size, err := m.TraceIdRatioBased.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTraceConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *TraceConfig_RateLimitingSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig_RateLimitingSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RateLimitingSampler != nil {
		{
			size, err := m.RateLimitingSampler.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTraceConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *ConstantSampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConstantSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConstantSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decision != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.Decision))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TraceIdRatioBased) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceIdRatioBased) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceIdRatioBased) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SamplingRatio != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.SamplingRatio))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitingSampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitingSampler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimitingSampler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Qps != 0 {
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.Qps))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTraceConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovTraceConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TraceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sampler != nil {
		n += m.Sampler.Size()
	}
	if m.MaxNumberOfAttributes != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAttributes))
	}
	if m.MaxNumberOfTimedEvents != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfTimedEvents))
	}
	if m.MaxNumberOfAttributesPerTimedEvent != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAttributesPerTimedEvent))
	}
	if m.MaxNumberOfLinks != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfLinks))
	}
	if m.MaxNumberOfAttributesPerLink != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAttributesPerLink))
	}
	return n
}

func (m *TraceConfig_ConstantSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConstantSampler != nil {
		l = m.ConstantSampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *TraceConfig_TraceIdRatioBased) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TraceIdRatioBased != nil {
		l = m.TraceIdRatioBased.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *TraceConfig_RateLimitingSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RateLimitingSampler != nil {
		l = m.RateLimitingSampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *ConstantSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Decision != 0 {
		n += 1 + sovTraceConfig(uint64(m.Decision))
	}
	return n
}

func (m *TraceIdRatioBased) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SamplingRatio != 0 {
		n += 9
	}
	return n
}

func (m *RateLimitingSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Qps != 0 {
		n += 1 + sovTraceConfig(uint64(m.Qps))
	}
	return n
}

func sovTraceConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTraceConfig(x uint64) (n int) {
	return sovTraceConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TraceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
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
			return fmt.Errorf("proto: TraceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstantSampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
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
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ConstantSampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_ConstantSampler{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceIdRatioBased", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
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
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TraceIdRatioBased{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_TraceIdRatioBased{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitingSampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
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
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RateLimitingSampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_RateLimitingSampler{v}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAttributes", wireType)
			}
			m.MaxNumberOfAttributes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAttributes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfTimedEvents", wireType)
			}
			m.MaxNumberOfTimedEvents = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfTimedEvents |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAttributesPerTimedEvent", wireType)
			}
			m.MaxNumberOfAttributesPerTimedEvent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAttributesPerTimedEvent |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfLinks", wireType)
			}
			m.MaxNumberOfLinks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfLinks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAttributesPerLink", wireType)
			}
			m.MaxNumberOfAttributesPerLink = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAttributesPerLink |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTraceConfig
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
func (m *ConstantSampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
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
			return fmt.Errorf("proto: ConstantSampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConstantSampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decision", wireType)
			}
			m.Decision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decision |= ConstantSampler_ConstantDecision(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTraceConfig
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
func (m *TraceIdRatioBased) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
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
			return fmt.Errorf("proto: TraceIdRatioBased: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceIdRatioBased: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SamplingRatio", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.SamplingRatio = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTraceConfig
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
func (m *RateLimitingSampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
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
			return fmt.Errorf("proto: RateLimitingSampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitingSampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Qps", wireType)
			}
			m.Qps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Qps |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTraceConfig
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
func skipTraceConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTraceConfig
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
					return 0, ErrIntOverflowTraceConfig
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
					return 0, ErrIntOverflowTraceConfig
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
				return 0, ErrInvalidLengthTraceConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTraceConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTraceConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTraceConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTraceConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTraceConfig = fmt.Errorf("proto: unexpected end of group")
)
