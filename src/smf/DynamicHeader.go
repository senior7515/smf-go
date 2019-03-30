// automatically generated by the FlatBuffers compiler, do not modify

package smf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// \brief used for extra headers, ala HTTP
/// The use case for the core is to support
/// zipkin/google-Dapper style tracing
type DynamicHeader struct {
	_tab flatbuffers.Table
}

func GetRootAsDynamicHeader(buf []byte, offset flatbuffers.UOffsetT) *DynamicHeader {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DynamicHeader{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DynamicHeader) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DynamicHeader) Table() flatbuffers.Table {
	return rcv._tab
}

/// alows for binary search lookup
/// use with CreateVectorOfSortedTables<> instead of the CreateVector
func (rcv *DynamicHeader) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// alows for binary search lookup
/// use with CreateVectorOfSortedTables<> instead of the CreateVector
func (rcv *DynamicHeader) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DynamicHeaderStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DynamicHeaderAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func DynamicHeaderAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func DynamicHeaderEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}