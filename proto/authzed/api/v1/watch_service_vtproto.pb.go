// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.1-0.20240319094008-0393e58bdf10
// source: authzed/api/v1/watch_service.proto

package v1

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	structpb1 "github.com/planetscale/vtprotobuf/types/known/structpb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *WatchRequest) CloneVT() *WatchRequest {
	if m == nil {
		return (*WatchRequest)(nil)
	}
	r := new(WatchRequest)
	r.OptionalStartCursor = m.OptionalStartCursor.CloneVT()
	if rhs := m.OptionalObjectTypes; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.OptionalObjectTypes = tmpContainer
	}
	if rhs := m.OptionalRelationshipFilters; rhs != nil {
		tmpContainer := make([]*RelationshipFilter, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.OptionalRelationshipFilters = tmpContainer
	}
	if rhs := m.OptionalUpdateKinds; rhs != nil {
		tmpContainer := make([]WatchKind, len(rhs))
		copy(tmpContainer, rhs)
		r.OptionalUpdateKinds = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *WatchRequest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *WatchResponse) CloneVT() *WatchResponse {
	if m == nil {
		return (*WatchResponse)(nil)
	}
	r := new(WatchResponse)
	r.ChangesThrough = m.ChangesThrough.CloneVT()
	r.OptionalTransactionMetadata = (*structpb.Struct)((*structpb1.Struct)(m.OptionalTransactionMetadata).CloneVT())
	r.SchemaUpdated = m.SchemaUpdated
	r.IsCheckpoint = m.IsCheckpoint
	if rhs := m.Updates; rhs != nil {
		tmpContainer := make([]*RelationshipUpdate, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Updates = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *WatchResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *WatchRequest) EqualVT(that *WatchRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.OptionalObjectTypes) != len(that.OptionalObjectTypes) {
		return false
	}
	for i, vx := range this.OptionalObjectTypes {
		vy := that.OptionalObjectTypes[i]
		if vx != vy {
			return false
		}
	}
	if !this.OptionalStartCursor.EqualVT(that.OptionalStartCursor) {
		return false
	}
	if len(this.OptionalRelationshipFilters) != len(that.OptionalRelationshipFilters) {
		return false
	}
	for i, vx := range this.OptionalRelationshipFilters {
		vy := that.OptionalRelationshipFilters[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &RelationshipFilter{}
			}
			if q == nil {
				q = &RelationshipFilter{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.OptionalUpdateKinds) != len(that.OptionalUpdateKinds) {
		return false
	}
	for i, vx := range this.OptionalUpdateKinds {
		vy := that.OptionalUpdateKinds[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *WatchRequest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*WatchRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *WatchResponse) EqualVT(that *WatchResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Updates) != len(that.Updates) {
		return false
	}
	for i, vx := range this.Updates {
		vy := that.Updates[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &RelationshipUpdate{}
			}
			if q == nil {
				q = &RelationshipUpdate{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if !this.ChangesThrough.EqualVT(that.ChangesThrough) {
		return false
	}
	if !(*structpb1.Struct)(this.OptionalTransactionMetadata).EqualVT((*structpb1.Struct)(that.OptionalTransactionMetadata)) {
		return false
	}
	if this.SchemaUpdated != that.SchemaUpdated {
		return false
	}
	if this.IsCheckpoint != that.IsCheckpoint {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *WatchResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*WatchResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *WatchRequest) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *WatchRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.OptionalUpdateKinds) > 0 {
		var pksize2 int
		for _, num := range m.OptionalUpdateKinds {
			pksize2 += protohelpers.SizeOfVarint(uint64(num))
		}
		i -= pksize2
		j1 := i
		for _, num1 := range m.OptionalUpdateKinds {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA[j1] = uint8(num)
			j1++
		}
		i = protohelpers.EncodeVarint(dAtA, i, uint64(pksize2))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OptionalRelationshipFilters) > 0 {
		for iNdEx := len(m.OptionalRelationshipFilters) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.OptionalRelationshipFilters[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.OptionalStartCursor != nil {
		size, err := m.OptionalStartCursor.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OptionalObjectTypes) > 0 {
		for iNdEx := len(m.OptionalObjectTypes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.OptionalObjectTypes[iNdEx])
			copy(dAtA[i:], m.OptionalObjectTypes[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.OptionalObjectTypes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *WatchResponse) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *WatchResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.IsCheckpoint {
		i--
		if m.IsCheckpoint {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.SchemaUpdated {
		i--
		if m.SchemaUpdated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.OptionalTransactionMetadata != nil {
		size, err := (*structpb1.Struct)(m.OptionalTransactionMetadata).MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.ChangesThrough != nil {
		size, err := m.ChangesThrough.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Updates) > 0 {
		for iNdEx := len(m.Updates) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Updates[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *WatchRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.OptionalObjectTypes) > 0 {
		for _, s := range m.OptionalObjectTypes {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.OptionalStartCursor != nil {
		l = m.OptionalStartCursor.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.OptionalRelationshipFilters) > 0 {
		for _, e := range m.OptionalRelationshipFilters {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.OptionalUpdateKinds) > 0 {
		l = 0
		for _, e := range m.OptionalUpdateKinds {
			l += protohelpers.SizeOfVarint(uint64(e))
		}
		n += 1 + protohelpers.SizeOfVarint(uint64(l)) + l
	}
	n += len(m.unknownFields)
	return n
}

func (m *WatchResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Updates) > 0 {
		for _, e := range m.Updates {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.ChangesThrough != nil {
		l = m.ChangesThrough.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.OptionalTransactionMetadata != nil {
		l = (*structpb1.Struct)(m.OptionalTransactionMetadata).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.SchemaUpdated {
		n += 2
	}
	if m.IsCheckpoint {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *WatchRequest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: WatchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalObjectTypes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OptionalObjectTypes = append(m.OptionalObjectTypes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalStartCursor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OptionalStartCursor == nil {
				m.OptionalStartCursor = &ZedToken{}
			}
			if err := m.OptionalStartCursor.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalRelationshipFilters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OptionalRelationshipFilters = append(m.OptionalRelationshipFilters, &RelationshipFilter{})
			if err := m.OptionalRelationshipFilters[len(m.OptionalRelationshipFilters)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v WatchKind
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= WatchKind(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.OptionalUpdateKinds = append(m.OptionalUpdateKinds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return protohelpers.ErrInvalidLength
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return protohelpers.ErrInvalidLength
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.OptionalUpdateKinds) == 0 {
					m.OptionalUpdateKinds = make([]WatchKind, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v WatchKind
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protohelpers.ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= WatchKind(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.OptionalUpdateKinds = append(m.OptionalUpdateKinds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalUpdateKinds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: WatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Updates = append(m.Updates, &RelationshipUpdate{})
			if err := m.Updates[len(m.Updates)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangesThrough", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChangesThrough == nil {
				m.ChangesThrough = &ZedToken{}
			}
			if err := m.ChangesThrough.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalTransactionMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OptionalTransactionMetadata == nil {
				m.OptionalTransactionMetadata = &structpb.Struct{}
			}
			if err := (*structpb1.Struct)(m.OptionalTransactionMetadata).UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaUpdated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
			m.SchemaUpdated = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsCheckpoint", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
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
			m.IsCheckpoint = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
