// automatically generated by the FlatBuffers compiler, do not modify

package rpc

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PayloadHeaders struct {
	_tab flatbuffers.Table
}

func GetRootAsPayloadHeaders(buf []byte, offset flatbuffers.UOffsetT) *PayloadHeaders {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PayloadHeaders{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PayloadHeaders) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PayloadHeaders) Table() flatbuffers.Table {
	return rcv._tab
}

/// Headers for forward compat.
func (rcv *PayloadHeaders) DynamicHeaders(obj *DynamicHeader, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *PayloadHeaders) DynamicHeadersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Headers for forward compat.
/// We need to chain the actual payload
func (rcv *PayloadHeaders) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// We need to chain the actual payload
func (rcv *PayloadHeaders) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *PayloadHeaders) Checksum() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PayloadHeaders) MutateChecksum(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *PayloadHeaders) Compression() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PayloadHeaders) MutateCompression(n int8) bool {
	return rcv._tab.MutateInt8Slot(10, n)
}

func PayloadHeadersStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func PayloadHeadersAddDynamicHeaders(builder *flatbuffers.Builder, dynamicHeaders flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(dynamicHeaders), 0)
}
func PayloadHeadersStartDynamicHeadersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PayloadHeadersAddSize(builder *flatbuffers.Builder, size uint32) {
	builder.PrependUint32Slot(1, size, 0)
}
func PayloadHeadersAddChecksum(builder *flatbuffers.Builder, checksum uint32) {
	builder.PrependUint32Slot(2, checksum, 0)
}
func PayloadHeadersAddCompression(builder *flatbuffers.Builder, compression int8) {
	builder.PrependInt8Slot(3, compression, 0)
}
func PayloadHeadersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
