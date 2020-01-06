// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tracer_base.proto

package pb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

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

type TraceType int32

const (
	TraceType_DummyEvent  TraceType = 0
	TraceType_BinlogEvent TraceType = 1
	TraceType_JobEvent    TraceType = 2
)

var TraceType_name = map[int32]string{
	0: "DummyEvent",
	1: "BinlogEvent",
	2: "JobEvent",
}

var TraceType_value = map[string]int32{
	"DummyEvent":  0,
	"BinlogEvent": 1,
	"JobEvent":    2,
}

func (x TraceType) String() string {
	return proto.EnumName(TraceType_name, int32(x))
}

func (TraceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be09eaf025e7bb5c, []int{0}
}

type BaseEvent struct {
	Filename string    `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Line     int32     `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	Tso      int64     `protobuf:"varint,3,opt,name=tso,proto3" json:"tso,omitempty"`
	TraceID  string    `protobuf:"bytes,4,opt,name=traceID,proto3" json:"traceID,omitempty"`
	GroupID  string    `protobuf:"bytes,5,opt,name=groupID,proto3" json:"groupID,omitempty"`
	Type     TraceType `protobuf:"varint,6,opt,name=type,proto3,enum=pb.TraceType" json:"type,omitempty"`
}

func (m *BaseEvent) Reset()         { *m = BaseEvent{} }
func (m *BaseEvent) String() string { return proto.CompactTextString(m) }
func (*BaseEvent) ProtoMessage()    {}
func (*BaseEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_be09eaf025e7bb5c, []int{0}
}
func (m *BaseEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseEvent.Merge(m, src)
}
func (m *BaseEvent) XXX_Size() int {
	return m.Size()
}
func (m *BaseEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BaseEvent proto.InternalMessageInfo

func (m *BaseEvent) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *BaseEvent) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *BaseEvent) GetTso() int64 {
	if m != nil {
		return m.Tso
	}
	return 0
}

func (m *BaseEvent) GetTraceID() string {
	if m != nil {
		return m.TraceID
	}
	return ""
}

func (m *BaseEvent) GetGroupID() string {
	if m != nil {
		return m.GroupID
	}
	return ""
}

func (m *BaseEvent) GetType() TraceType {
	if m != nil {
		return m.Type
	}
	return TraceType_DummyEvent
}

func init() {
	proto.RegisterEnum("pb.TraceType", TraceType_name, TraceType_value)
	proto.RegisterType((*BaseEvent)(nil), "pb.BaseEvent")
}

func init() { proto.RegisterFile("tracer_base.proto", fileDescriptor_be09eaf025e7bb5c) }

var fileDescriptor_be09eaf025e7bb5c = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x86, 0x19, 0xe0, 0xce, 0x63, 0xd4, 0xf3, 0x9c, 0x6a, 0x63, 0xb1, 0x41, 0x2b, 0x62, 0x41,
	0xa1, 0x9d, 0x25, 0xc1, 0xe2, 0x2c, 0xc9, 0xf5, 0x86, 0x35, 0xeb, 0x85, 0x04, 0xd8, 0x0d, 0xec,
	0x99, 0xf0, 0x16, 0x3e, 0x84, 0x0f, 0x63, 0x79, 0xa5, 0xa5, 0x81, 0x17, 0x31, 0x2c, 0x1e, 0xdd,
	0xff, 0xfd, 0x5f, 0x76, 0x67, 0x32, 0x78, 0x6d, 0x9a, 0xfc, 0x4d, 0x36, 0xaf, 0x22, 0x6f, 0x65,
	0xac, 0x1b, 0x65, 0x14, 0xb9, 0x5a, 0xdc, 0x7d, 0x01, 0x06, 0x49, 0xde, 0xca, 0xe7, 0x0f, 0x59,
	0x1b, 0xba, 0xc1, 0xd5, 0x7b, 0x51, 0xca, 0x3a, 0xaf, 0x24, 0x83, 0x10, 0xa2, 0x20, 0x9b, 0x99,
	0x08, 0xfd, 0xb2, 0xa8, 0x25, 0x73, 0x43, 0x88, 0x16, 0x99, 0xcd, 0xb4, 0x41, 0xcf, 0xb4, 0x8a,
	0x79, 0x21, 0x44, 0x5e, 0x36, 0x46, 0x62, 0x78, 0x66, 0x07, 0x6d, 0x53, 0xe6, 0xdb, 0x0f, 0x4e,
	0x38, 0x9a, 0x7d, 0xa3, 0x0e, 0x7a, 0x9b, 0xb2, 0xc5, 0x64, 0xfe, 0x91, 0x6e, 0xd1, 0x37, 0x9d,
	0x96, 0x6c, 0x19, 0x42, 0xb4, 0x7e, 0xb8, 0x8c, 0xb5, 0x88, 0x77, 0xe3, 0xa3, 0x5d, 0xa7, 0x65,
	0x66, 0xd5, 0xfd, 0x13, 0x06, 0x73, 0x45, 0x6b, 0xc4, 0xf4, 0x50, 0x55, 0x9d, 0xdd, 0x79, 0xe3,
	0xd0, 0x15, 0x9e, 0x27, 0x45, 0x5d, 0xaa, 0xfd, 0x54, 0x00, 0x5d, 0xe0, 0xea, 0x45, 0x89, 0x89,
	0xdc, 0x84, 0x7d, 0xf7, 0x1c, 0x8e, 0x3d, 0x87, 0xdf, 0x9e, 0xc3, 0xe7, 0xc0, 0x9d, 0xe3, 0xc0,
	0x9d, 0x9f, 0x81, 0x3b, 0x62, 0x69, 0xef, 0xf0, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x25, 0x73,
	0x7a, 0x49, 0x1c, 0x01, 0x00, 0x00,
}

func (m *BaseEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintTracerBase(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x30
	}
	if len(m.GroupID) > 0 {
		i -= len(m.GroupID)
		copy(dAtA[i:], m.GroupID)
		i = encodeVarintTracerBase(dAtA, i, uint64(len(m.GroupID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TraceID) > 0 {
		i -= len(m.TraceID)
		copy(dAtA[i:], m.TraceID)
		i = encodeVarintTracerBase(dAtA, i, uint64(len(m.TraceID)))
		i--
		dAtA[i] = 0x22
	}
	if m.Tso != 0 {
		i = encodeVarintTracerBase(dAtA, i, uint64(m.Tso))
		i--
		dAtA[i] = 0x18
	}
	if m.Line != 0 {
		i = encodeVarintTracerBase(dAtA, i, uint64(m.Line))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Filename) > 0 {
		i -= len(m.Filename)
		copy(dAtA[i:], m.Filename)
		i = encodeVarintTracerBase(dAtA, i, uint64(len(m.Filename)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTracerBase(dAtA []byte, offset int, v uint64) int {
	offset -= sovTracerBase(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Filename)
	if l > 0 {
		n += 1 + l + sovTracerBase(uint64(l))
	}
	if m.Line != 0 {
		n += 1 + sovTracerBase(uint64(m.Line))
	}
	if m.Tso != 0 {
		n += 1 + sovTracerBase(uint64(m.Tso))
	}
	l = len(m.TraceID)
	if l > 0 {
		n += 1 + l + sovTracerBase(uint64(l))
	}
	l = len(m.GroupID)
	if l > 0 {
		n += 1 + l + sovTracerBase(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTracerBase(uint64(m.Type))
	}
	return n
}

func sovTracerBase(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTracerBase(x uint64) (n int) {
	return sovTracerBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracerBase
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
			return fmt.Errorf("proto: BaseEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filename", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerBase
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
				return ErrInvalidLengthTracerBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filename = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			m.Line = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Line |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tso", wireType)
			}
			m.Tso = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tso |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerBase
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
				return ErrInvalidLengthTracerBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerBase
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
				return ErrInvalidLengthTracerBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= TraceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTracerBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracerBase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTracerBase
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
func skipTracerBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTracerBase
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
					return 0, ErrIntOverflowTracerBase
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
					return 0, ErrIntOverflowTracerBase
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
				return 0, ErrInvalidLengthTracerBase
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTracerBase
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTracerBase
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTracerBase        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTracerBase          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTracerBase = fmt.Errorf("proto: unexpected end of group")
)
