// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/type/timeofday.proto

package _type

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Represents a time of day. The date and time zone are either not significant
// or are specified elsewhere. An API may choose to allow leap seconds. Related
// types are [google.type.Date][google.type.Date] and `google.protobuf.Timestamp`.
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	// to allow the value "24:00:00" for scenarios like business closing time.
	Hours int32 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	// Minutes of hour of day. Must be from 0 to 59.
	Minutes int32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	// allow the value 60 if it allows leap-seconds.
	Seconds int32 `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos                int32    `protobuf:"varint,4,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeOfDay) Reset()      { *m = TimeOfDay{} }
func (*TimeOfDay) ProtoMessage() {}
func (*TimeOfDay) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c354ab94588cebb, []int{0}
}
func (m *TimeOfDay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeOfDay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeOfDay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeOfDay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeOfDay.Merge(m, src)
}
func (m *TimeOfDay) XXX_Size() int {
	return m.Size()
}
func (m *TimeOfDay) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeOfDay.DiscardUnknown(m)
}

var xxx_messageInfo_TimeOfDay proto.InternalMessageInfo

func (m *TimeOfDay) GetHours() int32 {
	if m != nil {
		return m.Hours
	}
	return 0
}

func (m *TimeOfDay) GetMinutes() int32 {
	if m != nil {
		return m.Minutes
	}
	return 0
}

func (m *TimeOfDay) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TimeOfDay) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func (*TimeOfDay) XXX_MessageName() string {
	return "google.type.TimeOfDay"
}
func init() {
	proto.RegisterType((*TimeOfDay)(nil), "google.type.TimeOfDay")
}

func init() { proto.RegisterFile("google/type/timeofday.proto", fileDescriptor_6c354ab94588cebb) }

var fileDescriptor_6c354ab94588cebb = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0xc9, 0xcc, 0x4d, 0xcd, 0x4f, 0x4b, 0x49,
	0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x48, 0xea, 0x81, 0x24, 0x95, 0xb2,
	0xb9, 0x38, 0x43, 0x32, 0x73, 0x53, 0xfd, 0xd3, 0x5c, 0x12, 0x2b, 0x85, 0x44, 0xb8, 0x58, 0x33,
	0xf2, 0x4b, 0x8b, 0x8a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c, 0x21, 0x09, 0x2e,
	0xf6, 0xdc, 0xcc, 0xbc, 0xd2, 0x92, 0xd4, 0x62, 0x09, 0x26, 0xb0, 0x38, 0x8c, 0x0b, 0x92, 0x29,
	0x4e, 0x4d, 0xce, 0xcf, 0x4b, 0x29, 0x96, 0x60, 0x86, 0xc8, 0x40, 0xb9, 0x20, 0x93, 0xf2, 0x12,
	0xf3, 0xf2, 0x8b, 0x25, 0x58, 0x20, 0x26, 0x81, 0x39, 0x4e, 0x49, 0x37, 0x1e, 0xca, 0x31, 0x7c,
	0x78, 0x28, 0xc7, 0xf8, 0xe3, 0xa1, 0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18,
	0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x17, 0x8f, 0xe4,
	0x18, 0x3e, 0x80, 0xc4, 0x1f, 0xcb, 0x31, 0x9e, 0x78, 0x2c, 0xc7, 0xc8, 0xc5, 0x9f, 0x9c, 0x9f,
	0xab, 0x87, 0xe4, 0x5e, 0x27, 0x3e, 0xb8, 0x6b, 0x03, 0x40, 0x9e, 0x09, 0x60, 0x8c, 0x62, 0x01,
	0x89, 0xff, 0x60, 0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0x12, 0x90, 0xc4, 0x06, 0xf6, 0xa4, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x7e, 0xc2, 0x1f, 0x03, 0x01, 0x00, 0x00,
}

func (this *TimeOfDay) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*TimeOfDay)
	if !ok {
		that2, ok := that.(TimeOfDay)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Hours != that1.Hours {
		if this.Hours < that1.Hours {
			return -1
		}
		return 1
	}
	if this.Minutes != that1.Minutes {
		if this.Minutes < that1.Minutes {
			return -1
		}
		return 1
	}
	if this.Seconds != that1.Seconds {
		if this.Seconds < that1.Seconds {
			return -1
		}
		return 1
	}
	if this.Nanos != that1.Nanos {
		if this.Nanos < that1.Nanos {
			return -1
		}
		return 1
	}
	if c := bytes.Compare(this.XXX_unrecognized, that1.XXX_unrecognized); c != 0 {
		return c
	}
	return 0
}
func (this *TimeOfDay) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TimeOfDay)
	if !ok {
		that2, ok := that.(TimeOfDay)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Hours != that1.Hours {
		return false
	}
	if this.Minutes != that1.Minutes {
		return false
	}
	if this.Seconds != that1.Seconds {
		return false
	}
	if this.Nanos != that1.Nanos {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TimeOfDay) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&_type.TimeOfDay{")
	s = append(s, "Hours: "+fmt.Sprintf("%#v", this.Hours)+",\n")
	s = append(s, "Minutes: "+fmt.Sprintf("%#v", this.Minutes)+",\n")
	s = append(s, "Seconds: "+fmt.Sprintf("%#v", this.Seconds)+",\n")
	s = append(s, "Nanos: "+fmt.Sprintf("%#v", this.Nanos)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTimeofday(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TimeOfDay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeOfDay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeOfDay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Nanos != 0 {
		i = encodeVarintTimeofday(dAtA, i, uint64(m.Nanos))
		i--
		dAtA[i] = 0x20
	}
	if m.Seconds != 0 {
		i = encodeVarintTimeofday(dAtA, i, uint64(m.Seconds))
		i--
		dAtA[i] = 0x18
	}
	if m.Minutes != 0 {
		i = encodeVarintTimeofday(dAtA, i, uint64(m.Minutes))
		i--
		dAtA[i] = 0x10
	}
	if m.Hours != 0 {
		i = encodeVarintTimeofday(dAtA, i, uint64(m.Hours))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTimeofday(dAtA []byte, offset int, v uint64) int {
	offset -= sovTimeofday(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedTimeOfDay(r randyTimeofday, easy bool) *TimeOfDay {
	this := &TimeOfDay{}
	this.Hours = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Hours *= -1
	}
	this.Minutes = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Minutes *= -1
	}
	this.Seconds = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Seconds *= -1
	}
	this.Nanos = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Nanos *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTimeofday(r, 5)
	}
	return this
}

type randyTimeofday interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTimeofday(r randyTimeofday) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTimeofday(r randyTimeofday) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTimeofday(r)
	}
	return string(tmps)
}
func randUnrecognizedTimeofday(r randyTimeofday, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTimeofday(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTimeofday(dAtA []byte, r randyTimeofday, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTimeofday(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTimeofday(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTimeofday(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTimeofday(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTimeofday(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTimeofday(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTimeofday(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *TimeOfDay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Hours != 0 {
		n += 1 + sovTimeofday(uint64(m.Hours))
	}
	if m.Minutes != 0 {
		n += 1 + sovTimeofday(uint64(m.Minutes))
	}
	if m.Seconds != 0 {
		n += 1 + sovTimeofday(uint64(m.Seconds))
	}
	if m.Nanos != 0 {
		n += 1 + sovTimeofday(uint64(m.Nanos))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTimeofday(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTimeofday(x uint64) (n int) {
	return sovTimeofday(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TimeOfDay) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TimeOfDay{`,
		`Hours:` + fmt.Sprintf("%v", this.Hours) + `,`,
		`Minutes:` + fmt.Sprintf("%v", this.Minutes) + `,`,
		`Seconds:` + fmt.Sprintf("%v", this.Seconds) + `,`,
		`Nanos:` + fmt.Sprintf("%v", this.Nanos) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTimeofday(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TimeOfDay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeofday
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
			return fmt.Errorf("proto: TimeOfDay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeOfDay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hours |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minutes", wireType)
			}
			m.Minutes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Minutes |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seconds |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
			}
			m.Nanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeofday
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nanos |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimeofday(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimeofday
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
func skipTimeofday(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimeofday
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
					return 0, ErrIntOverflowTimeofday
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
					return 0, ErrIntOverflowTimeofday
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
				return 0, ErrInvalidLengthTimeofday
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTimeofday
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTimeofday
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTimeofday        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimeofday          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTimeofday = fmt.Errorf("proto: unexpected end of group")
)
