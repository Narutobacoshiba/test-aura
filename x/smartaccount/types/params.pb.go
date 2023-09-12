// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aura/smartaccount/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type CodeID struct {
	// whitelist code id
	CodeID uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	// status of code id
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *CodeID) Reset()         { *m = CodeID{} }
func (m *CodeID) String() string { return proto.CompactTextString(m) }
func (*CodeID) ProtoMessage()    {}
func (*CodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_74fa9ef7f6db6ea9, []int{0}
}
func (m *CodeID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodeID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeID.Merge(m, src)
}
func (m *CodeID) XXX_Size() int {
	return m.Size()
}
func (m *CodeID) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeID.DiscardUnknown(m)
}

var xxx_messageInfo_CodeID proto.InternalMessageInfo

func (m *CodeID) GetCodeID() uint64 {
	if m != nil {
		return m.CodeID
	}
	return 0
}

func (m *CodeID) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

// Params defines the parameters for the module.
type Params struct {
	// code_id whitelist indicates which contract can be initialized as smart account
	// using gov proposal for updates
	WhitelistCodeID []*CodeID `protobuf:"bytes,1,rep,name=whitelist_code_id,json=whitelistCodeId,proto3" json:"whitelist_code_id,omitempty"`
	// list of diable messages for smartaccount
	DisableMsgsList []string `protobuf:"bytes,2,rep,name=disable_msgs_list,json=disableMsgsList,proto3" json:"disable_msgs_list,omitempty"`
	// limit how much gas can be consumed by the `pre_execute` method
	MaxGasExecute uint64 `protobuf:"varint,3,opt,name=max_gas_execute,json=maxGasExecute,proto3" json:"max_gas_execute,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_74fa9ef7f6db6ea9, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetWhitelistCodeID() []*CodeID {
	if m != nil {
		return m.WhitelistCodeID
	}
	return nil
}

func (m *Params) GetDisableMsgsList() []string {
	if m != nil {
		return m.DisableMsgsList
	}
	return nil
}

func (m *Params) GetMaxGasExecute() uint64 {
	if m != nil {
		return m.MaxGasExecute
	}
	return 0
}

func init() {
	proto.RegisterType((*CodeID)(nil), "aura.smartaccount.CodeID")
	proto.RegisterType((*Params)(nil), "aura.smartaccount.Params")
}

func init() { proto.RegisterFile("aura/smartaccount/params.proto", fileDescriptor_74fa9ef7f6db6ea9) }

var fileDescriptor_74fa9ef7f6db6ea9 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x33, 0x6d, 0xc9, 0xbd, 0x77, 0x2e, 0x25, 0x24, 0x8a, 0x44, 0x17, 0x93, 0x52, 0x37,
	0x59, 0x68, 0x02, 0xba, 0xb2, 0x1b, 0xa1, 0x56, 0xa5, 0x68, 0x41, 0xb2, 0x11, 0xba, 0x09, 0xd3,
	0xcc, 0x90, 0x06, 0x9a, 0x4e, 0xc9, 0x99, 0xd0, 0xf8, 0x16, 0x2e, 0x5d, 0xf6, 0x71, 0x5c, 0x76,
	0xe9, 0xaa, 0x48, 0x0a, 0xe2, 0x63, 0x48, 0x92, 0x8a, 0x2d, 0xae, 0x66, 0xce, 0xf9, 0x66, 0xce,
	0xf9, 0x7f, 0x7e, 0x4c, 0x68, 0x9a, 0x50, 0x17, 0x62, 0x9a, 0x48, 0x1a, 0x04, 0x22, 0x9d, 0x4a,
	0x77, 0x46, 0x13, 0x1a, 0x83, 0x33, 0x4b, 0x84, 0x14, 0x86, 0x5e, 0x70, 0x67, 0x9b, 0x1f, 0xed,
	0x87, 0x22, 0x14, 0x25, 0x75, 0x8b, 0x5b, 0xf5, 0xb0, 0x7d, 0x87, 0xd5, 0x2b, 0xc1, 0x78, 0xbf,
	0x67, 0x1c, 0xe3, 0x3f, 0x81, 0x60, 0xdc, 0x8f, 0x98, 0x89, 0x5a, 0xc8, 0x6e, 0x74, 0x71, 0xbe,
	0xb2, 0x36, 0xd0, 0x53, 0x0b, 0xd4, 0x67, 0xc6, 0x01, 0x56, 0x41, 0x52, 0x99, 0x82, 0x59, 0x6b,
	0x21, 0xfb, 0xaf, 0xb7, 0xa9, 0x3a, 0x8d, 0xcf, 0x85, 0x85, 0xda, 0x1f, 0x08, 0xab, 0x0f, 0xa5,
	0x0c, 0x63, 0x88, 0xf5, 0xf9, 0x38, 0x92, 0x7c, 0x12, 0x81, 0xf4, 0x7f, 0xe6, 0xd6, 0xed, 0xff,
	0x67, 0x87, 0xce, 0x2f, 0x71, 0x4e, 0xb5, 0xa6, 0xbb, 0x97, 0xaf, 0x2c, 0xed, 0xf1, 0xfb, 0xdf,
	0x66, 0xb7, 0x36, 0xdf, 0x69, 0x30, 0xe3, 0x12, 0xeb, 0x2c, 0x02, 0x3a, 0x9a, 0x70, 0x3f, 0x86,
	0x10, 0xfc, 0x02, 0x99, 0xb5, 0x56, 0xdd, 0xfe, 0x57, 0x0d, 0xe8, 0x55, 0x70, 0x00, 0x21, 0xdc,
	0x47, 0x20, 0x3d, 0x8d, 0xed, 0x36, 0x8c, 0x0b, 0xac, 0xc5, 0x34, 0xf3, 0x43, 0x0a, 0x3e, 0xcf,
	0x78, 0x90, 0x4a, 0x6e, 0xd6, 0x4b, 0xcb, 0x7a, 0xbe, 0xb2, 0x9a, 0x03, 0x9a, 0xdd, 0x52, 0xb8,
	0xae, 0x80, 0xd7, 0x8c, 0xb7, 0xcb, 0x4e, 0xe3, 0x65, 0x61, 0x29, 0xdd, 0x9b, 0xd7, 0x9c, 0xa0,
	0x65, 0x4e, 0xd0, 0x7b, 0x4e, 0xd0, 0xf3, 0x9a, 0x28, 0xcb, 0x35, 0x51, 0xde, 0xd6, 0x44, 0x19,
	0x9e, 0x84, 0x91, 0x1c, 0xa7, 0x23, 0x27, 0x10, 0xb1, 0x5b, 0xd8, 0x3c, 0x9d, 0xce, 0xcb, 0xd3,
	0xcd, 0x76, 0xd3, 0x92, 0x4f, 0x33, 0x0e, 0x23, 0xb5, 0x0c, 0xe1, 0xfc, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0x83, 0xbb, 0xde, 0xcf, 0x01, 0x00, 0x00,
}

func (this *CodeID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodeID)
	if !ok {
		that2, ok := that.(CodeID)
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
	if this.CodeID != that1.CodeID {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (m *CodeID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodeID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status {
		i--
		if m.Status {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.CodeID != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxGasExecute != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxGasExecute))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DisableMsgsList) > 0 {
		for iNdEx := len(m.DisableMsgsList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DisableMsgsList[iNdEx])
			copy(dAtA[i:], m.DisableMsgsList[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.DisableMsgsList[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.WhitelistCodeID) > 0 {
		for iNdEx := len(m.WhitelistCodeID) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhitelistCodeID[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CodeID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeID != 0 {
		n += 1 + sovParams(uint64(m.CodeID))
	}
	if m.Status {
		n += 2
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.WhitelistCodeID) > 0 {
		for _, e := range m.WhitelistCodeID {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.DisableMsgsList) > 0 {
		for _, s := range m.DisableMsgsList {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.MaxGasExecute != 0 {
		n += 1 + sovParams(uint64(m.MaxGasExecute))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CodeID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: CodeID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.Status = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistCodeID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistCodeID = append(m.WhitelistCodeID, &CodeID{})
			if err := m.WhitelistCodeID[len(m.WhitelistCodeID)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableMsgsList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisableMsgsList = append(m.DisableMsgsList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGasExecute", wireType)
			}
			m.MaxGasExecute = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxGasExecute |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)