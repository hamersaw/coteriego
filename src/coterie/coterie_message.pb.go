// Code generated by protoc-gen-go.
// source: coterie_message.proto
// DO NOT EDIT!

/*
Package coterie is a generated protocol buffer package.

It is generated from these files:
	coterie_message.proto

It has these top-level messages:
	CoterieMsg
	CloseConnectionMsg
	InsertEntryMsg
	InsertRecordMsg
	Record
	RecordBatchMsg
	ResultMsg
*/
package coterie

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CoterieMsg_Type int32

const (
	CoterieMsg_CLOSE_CONNECTION CoterieMsg_Type = 0
	CoterieMsg_INSERT_ENTRY     CoterieMsg_Type = 1
	CoterieMsg_INSERT_RECORD    CoterieMsg_Type = 2
	CoterieMsg_RECORD_BATCH     CoterieMsg_Type = 3
	CoterieMsg_RESULT           CoterieMsg_Type = 4
)

var CoterieMsg_Type_name = map[int32]string{
	0: "CLOSE_CONNECTION",
	1: "INSERT_ENTRY",
	2: "INSERT_RECORD",
	3: "RECORD_BATCH",
	4: "RESULT",
}
var CoterieMsg_Type_value = map[string]int32{
	"CLOSE_CONNECTION": 0,
	"INSERT_ENTRY":     1,
	"INSERT_RECORD":    2,
	"RECORD_BATCH":     3,
	"RESULT":           4,
}

func (x CoterieMsg_Type) String() string {
	return proto.EnumName(CoterieMsg_Type_name, int32(x))
}
func (CoterieMsg_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type CoterieMsg struct {
	Type               CoterieMsg_Type     `protobuf:"varint,1,opt,name=type,enum=coterie.CoterieMsg_Type" json:"type,omitempty"`
	CloseConnectionMsg *CloseConnectionMsg `protobuf:"bytes,2,opt,name=close_connection_msg" json:"close_connection_msg,omitempty"`
	InsertEntryMsg     *InsertEntryMsg     `protobuf:"bytes,3,opt,name=insert_entry_msg" json:"insert_entry_msg,omitempty"`
	InsertRecordMsg    *InsertRecordMsg    `protobuf:"bytes,4,opt,name=insert_record_msg" json:"insert_record_msg,omitempty"`
	RecordBatchMsg     *RecordBatchMsg     `protobuf:"bytes,5,opt,name=record_batch_msg" json:"record_batch_msg,omitempty"`
	ResultMsg          *ResultMsg          `protobuf:"bytes,6,opt,name=result_msg" json:"result_msg,omitempty"`
}

func (m *CoterieMsg) Reset()                    { *m = CoterieMsg{} }
func (m *CoterieMsg) String() string            { return proto.CompactTextString(m) }
func (*CoterieMsg) ProtoMessage()               {}
func (*CoterieMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CoterieMsg) GetCloseConnectionMsg() *CloseConnectionMsg {
	if m != nil {
		return m.CloseConnectionMsg
	}
	return nil
}

func (m *CoterieMsg) GetInsertEntryMsg() *InsertEntryMsg {
	if m != nil {
		return m.InsertEntryMsg
	}
	return nil
}

func (m *CoterieMsg) GetInsertRecordMsg() *InsertRecordMsg {
	if m != nil {
		return m.InsertRecordMsg
	}
	return nil
}

func (m *CoterieMsg) GetRecordBatchMsg() *RecordBatchMsg {
	if m != nil {
		return m.RecordBatchMsg
	}
	return nil
}

func (m *CoterieMsg) GetResultMsg() *ResultMsg {
	if m != nil {
		return m.ResultMsg
	}
	return nil
}

type CloseConnectionMsg struct {
	Reason string `protobuf:"bytes,1,opt,name=reason" json:"reason,omitempty"`
}

func (m *CloseConnectionMsg) Reset()                    { *m = CloseConnectionMsg{} }
func (m *CloseConnectionMsg) String() string            { return proto.CompactTextString(m) }
func (*CloseConnectionMsg) ProtoMessage()               {}
func (*CloseConnectionMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type InsertEntryMsg struct {
	Token uint64 `protobuf:"varint,1,opt,name=token" json:"token,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *InsertEntryMsg) Reset()                    { *m = InsertEntryMsg{} }
func (m *InsertEntryMsg) String() string            { return proto.CompactTextString(m) }
func (*InsertEntryMsg) ProtoMessage()               {}
func (*InsertEntryMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type InsertRecordMsg struct {
	Token  uint64  `protobuf:"varint,1,opt,name=token" json:"token,omitempty"`
	Record *Record `protobuf:"bytes,2,opt,name=record" json:"record,omitempty"`
}

func (m *InsertRecordMsg) Reset()                    { *m = InsertRecordMsg{} }
func (m *InsertRecordMsg) String() string            { return proto.CompactTextString(m) }
func (*InsertRecordMsg) ProtoMessage()               {}
func (*InsertRecordMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InsertRecordMsg) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type Record struct {
	Entries map[string]string `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Record) GetEntries() map[string]string {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RecordBatchMsg struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *RecordBatchMsg) Reset()                    { *m = RecordBatchMsg{} }
func (m *RecordBatchMsg) String() string            { return proto.CompactTextString(m) }
func (*RecordBatchMsg) ProtoMessage()               {}
func (*RecordBatchMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RecordBatchMsg) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type ResultMsg struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *ResultMsg) Reset()                    { *m = ResultMsg{} }
func (m *ResultMsg) String() string            { return proto.CompactTextString(m) }
func (*ResultMsg) ProtoMessage()               {}
func (*ResultMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*CoterieMsg)(nil), "coterie.CoterieMsg")
	proto.RegisterType((*CloseConnectionMsg)(nil), "coterie.CloseConnectionMsg")
	proto.RegisterType((*InsertEntryMsg)(nil), "coterie.InsertEntryMsg")
	proto.RegisterType((*InsertRecordMsg)(nil), "coterie.InsertRecordMsg")
	proto.RegisterType((*Record)(nil), "coterie.Record")
	proto.RegisterType((*RecordBatchMsg)(nil), "coterie.RecordBatchMsg")
	proto.RegisterType((*ResultMsg)(nil), "coterie.ResultMsg")
	proto.RegisterEnum("coterie.CoterieMsg_Type", CoterieMsg_Type_name, CoterieMsg_Type_value)
}

var fileDescriptor0 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0x7e, 0xf3, 0xb1, 0xf6, 0xcd, 0x61, 0xb4, 0x99, 0x35, 0x20, 0x02, 0x2e, 0x50, 0x6e, 0x00,
	0x69, 0x8a, 0xc4, 0x90, 0x00, 0x8d, 0xab, 0x35, 0x8b, 0xb4, 0x4a, 0x5b, 0x8a, 0xdc, 0x70, 0xc1,
	0x55, 0xd4, 0x05, 0xab, 0x44, 0xeb, 0xe2, 0xc9, 0x76, 0x91, 0x2a, 0xf1, 0x87, 0xf9, 0x17, 0xd8,
	0x27, 0x49, 0xd7, 0xa4, 0x70, 0x77, 0x8e, 0x9f, 0x0f, 0x3f, 0xc7, 0x1f, 0xf0, 0xa4, 0xe0, 0x8a,
	0x89, 0x92, 0xe5, 0x77, 0x4c, 0xca, 0xc5, 0x92, 0x45, 0xf7, 0x82, 0x2b, 0x4e, 0x86, 0xcd, 0x72,
	0xf8, 0xdb, 0x01, 0x88, 0xeb, 0xfa, 0x5a, 0x2e, 0xc9, 0x09, 0xb8, 0x6a, 0x73, 0xcf, 0x02, 0xeb,
	0x95, 0xf5, 0x66, 0x74, 0x1a, 0x44, 0x0d, 0x2d, 0x7a, 0xa0, 0x44, 0x99, 0xc6, 0x29, 0xb2, 0xc8,
	0x35, 0x1c, 0x17, 0x2b, 0x2e, 0x59, 0x5e, 0xf0, 0xaa, 0x62, 0x85, 0x2a, 0x79, 0x95, 0xdf, 0xc9,
	0x65, 0x60, 0x6b, 0xf5, 0xa3, 0xd3, 0x17, 0x0f, 0x6a, 0x43, 0x8a, 0xb7, 0x1c, 0xed, 0x42, 0x49,
	0xb1, 0xb7, 0x46, 0xce, 0xc1, 0x2f, 0x2b, 0xc9, 0x84, 0xca, 0x59, 0xa5, 0xc4, 0x06, 0xad, 0x1c,
	0xb4, 0x7a, 0xb6, 0xb5, 0x9a, 0x22, 0x21, 0x31, 0xb8, 0xb1, 0x19, 0x95, 0x9d, 0x9e, 0x5c, 0xc0,
	0x51, 0x63, 0x21, 0x58, 0xc1, 0xc5, 0x77, 0xf4, 0x70, 0xd1, 0x23, 0xe8, 0x79, 0x50, 0x24, 0x18,
	0x93, 0x71, 0xd9, 0x5d, 0x30, 0x41, 0x1a, 0xf9, 0xcd, 0x42, 0x15, 0x3f, 0xd0, 0xe4, 0xa0, 0x17,
	0xa4, 0x66, 0x4f, 0x0c, 0x8e, 0x41, 0x44, 0xa7, 0x27, 0xef, 0x00, 0x04, 0x93, 0xeb, 0x95, 0x42,
	0xf1, 0x00, 0xc5, 0x64, 0x47, 0x6c, 0x20, 0xa3, 0xf3, 0x44, 0x5b, 0x86, 0x39, 0xb8, 0xe6, 0x6c,
	0xc9, 0x31, 0xf8, 0xf1, 0xd5, 0x6c, 0x9e, 0xe4, 0xf1, 0x2c, 0x4d, 0x93, 0x38, 0x9b, 0xce, 0x52,
	0xff, 0x3f, 0xe2, 0xc3, 0xe1, 0x34, 0x9d, 0x27, 0x34, 0xcb, 0x93, 0x34, 0xa3, 0xdf, 0x7c, 0x8b,
	0x1c, 0xc1, 0xe3, 0x66, 0x85, 0x26, 0xf1, 0x8c, 0x5e, 0xf8, 0xb6, 0x21, 0xd5, 0x75, 0x3e, 0x39,
	0xcf, 0xe2, 0x4b, 0xdf, 0x21, 0x00, 0x03, 0x9a, 0xcc, 0xbf, 0x5e, 0x65, 0xbe, 0x1b, 0x9e, 0x00,
	0xd9, 0xbf, 0x09, 0xf2, 0x14, 0x06, 0x82, 0x2d, 0x24, 0xaf, 0xf0, 0xd2, 0x3d, 0xda, 0x74, 0x61,
	0x0a, 0xa3, 0xee, 0x61, 0xeb, 0x60, 0x07, 0x8a, 0xdf, 0xb2, 0x9a, 0xe8, 0xd2, 0xba, 0xd1, 0x7b,
	0x3a, 0xb7, 0x6c, 0x83, 0x77, 0xee, 0x51, 0x53, 0x1a, 0xde, 0xcf, 0xc5, 0x6a, 0xcd, 0xf0, 0xf2,
	0x3c, 0x5a, 0x37, 0xe1, 0x17, 0x18, 0xf7, 0x0e, 0xfe, 0x1f, 0x86, 0xaf, 0x4d, 0x20, 0x43, 0x69,
	0xde, 0xd1, 0xb8, 0x77, 0xe6, 0xb4, 0x81, 0xc3, 0x5f, 0x7a, 0x36, 0xac, 0xc8, 0x07, 0x18, 0x9a,
	0x27, 0x53, 0x32, 0xa9, 0xad, 0x1c, 0xad, 0x79, 0xd9, 0xd3, 0x44, 0x49, 0x0d, 0xe3, 0x2c, 0xb4,
	0x25, 0x3f, 0x3f, 0x83, 0xc3, 0x5d, 0xa0, 0x9d, 0xc5, 0xfa, 0xcb, 0x2c, 0xf6, 0xce, 0x2c, 0x67,
	0xf6, 0x27, 0x2b, 0xfc, 0x0c, 0xa3, 0xee, 0x1b, 0x20, 0x6f, 0x61, 0x58, 0x27, 0x6b, 0x53, 0xec,
	0x25, 0x6f, 0xf1, 0xf0, 0x23, 0x78, 0xdb, 0x37, 0x40, 0x02, 0x18, 0xca, 0x75, 0x51, 0xe8, 0x0f,
	0x8a, 0x3b, 0xff, 0x4f, 0xdb, 0xd6, 0xe4, 0x69, 0xff, 0x93, 0xce, 0xa3, 0xcb, 0x89, 0x7d, 0x69,
	0xdd, 0x0c, 0xf0, 0x0f, 0xbf, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xce, 0x0c, 0xe6, 0x9f, 0xdc,
	0x03, 0x00, 0x00,
}
