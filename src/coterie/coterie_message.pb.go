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
	Filter
	InsertEntryMsg
	InsertRecordMsg
	QueryMsg
	QueryResultMsg
	Record
	InsertRecordsMsg
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
	CoterieMsg_INSERT_RECORDS   CoterieMsg_Type = 3
	CoterieMsg_QUERY            CoterieMsg_Type = 4
	CoterieMsg_QUERY_RESULT     CoterieMsg_Type = 5
	CoterieMsg_RESULT           CoterieMsg_Type = 6
)

var CoterieMsg_Type_name = map[int32]string{
	0: "CLOSE_CONNECTION",
	1: "INSERT_ENTRY",
	2: "INSERT_RECORD",
	3: "INSERT_RECORDS",
	4: "QUERY",
	5: "QUERY_RESULT",
	6: "RESULT",
}
var CoterieMsg_Type_value = map[string]int32{
	"CLOSE_CONNECTION": 0,
	"INSERT_ENTRY":     1,
	"INSERT_RECORD":    2,
	"INSERT_RECORDS":   3,
	"QUERY":            4,
	"QUERY_RESULT":     5,
	"RESULT":           6,
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
	InsertRecordsMsg   *InsertRecordsMsg   `protobuf:"bytes,5,opt,name=insert_records_msg" json:"insert_records_msg,omitempty"`
	QueryMsg           *QueryMsg           `protobuf:"bytes,6,opt,name=query_msg" json:"query_msg,omitempty"`
	QueryResultMsg     *QueryResultMsg     `protobuf:"bytes,7,opt,name=query_result_msg" json:"query_result_msg,omitempty"`
	ResultMsg          *ResultMsg          `protobuf:"bytes,8,opt,name=result_msg" json:"result_msg,omitempty"`
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

func (m *CoterieMsg) GetInsertRecordsMsg() *InsertRecordsMsg {
	if m != nil {
		return m.InsertRecordsMsg
	}
	return nil
}

func (m *CoterieMsg) GetQueryMsg() *QueryMsg {
	if m != nil {
		return m.QueryMsg
	}
	return nil
}

func (m *CoterieMsg) GetQueryResultMsg() *QueryResultMsg {
	if m != nil {
		return m.QueryResultMsg
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

type Filter struct {
	FieldName string   `protobuf:"bytes,1,opt,name=field_name" json:"field_name,omitempty"`
	Type      string   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Arguments []string `protobuf:"bytes,3,rep,name=arguments" json:"arguments,omitempty"`
	Value     string   `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type InsertEntryMsg struct {
	Token uint64 `protobuf:"varint,1,opt,name=token" json:"token,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *InsertEntryMsg) Reset()                    { *m = InsertEntryMsg{} }
func (m *InsertEntryMsg) String() string            { return proto.CompactTextString(m) }
func (*InsertEntryMsg) ProtoMessage()               {}
func (*InsertEntryMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type InsertRecordMsg struct {
	Token  uint64  `protobuf:"varint,1,opt,name=token" json:"token,omitempty"`
	Record *Record `protobuf:"bytes,2,opt,name=record" json:"record,omitempty"`
}

func (m *InsertRecordMsg) Reset()                    { *m = InsertRecordMsg{} }
func (m *InsertRecordMsg) String() string            { return proto.CompactTextString(m) }
func (*InsertRecordMsg) ProtoMessage()               {}
func (*InsertRecordMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InsertRecordMsg) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type QueryMsg struct {
	FieldNames []string  `protobuf:"bytes,1,rep,name=field_names" json:"field_names,omitempty"`
	Filters    []*Filter `protobuf:"bytes,2,rep,name=filters" json:"filters,omitempty"`
}

func (m *QueryMsg) Reset()                    { *m = QueryMsg{} }
func (m *QueryMsg) String() string            { return proto.CompactTextString(m) }
func (*QueryMsg) ProtoMessage()               {}
func (*QueryMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueryMsg) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

type QueryResultMsg struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *QueryResultMsg) Reset()                    { *m = QueryResultMsg{} }
func (m *QueryResultMsg) String() string            { return proto.CompactTextString(m) }
func (*QueryResultMsg) ProtoMessage()               {}
func (*QueryResultMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueryResultMsg) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type Record struct {
	Entries map[string]string `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Record) GetEntries() map[string]string {
	if m != nil {
		return m.Entries
	}
	return nil
}

type InsertRecordsMsg struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *InsertRecordsMsg) Reset()                    { *m = InsertRecordsMsg{} }
func (m *InsertRecordsMsg) String() string            { return proto.CompactTextString(m) }
func (*InsertRecordsMsg) ProtoMessage()               {}
func (*InsertRecordsMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *InsertRecordsMsg) GetRecords() []*Record {
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
func (*ResultMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*CoterieMsg)(nil), "coterie.CoterieMsg")
	proto.RegisterType((*CloseConnectionMsg)(nil), "coterie.CloseConnectionMsg")
	proto.RegisterType((*Filter)(nil), "coterie.Filter")
	proto.RegisterType((*InsertEntryMsg)(nil), "coterie.InsertEntryMsg")
	proto.RegisterType((*InsertRecordMsg)(nil), "coterie.InsertRecordMsg")
	proto.RegisterType((*QueryMsg)(nil), "coterie.QueryMsg")
	proto.RegisterType((*QueryResultMsg)(nil), "coterie.QueryResultMsg")
	proto.RegisterType((*Record)(nil), "coterie.Record")
	proto.RegisterType((*InsertRecordsMsg)(nil), "coterie.InsertRecordsMsg")
	proto.RegisterType((*ResultMsg)(nil), "coterie.ResultMsg")
	proto.RegisterEnum("coterie.CoterieMsg_Type", CoterieMsg_Type_name, CoterieMsg_Type_value)
}

var fileDescriptor0 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0xf2, 0x3b, 0x9e, 0xf6, 0x4b, 0xdd, 0x51, 0x81, 0x00, 0x45, 0x20, 0x5f, 0x00, 0xa9,
	0xb2, 0x44, 0x91, 0x00, 0x15, 0x71, 0x00, 0xd7, 0x40, 0xa5, 0xd6, 0xa1, 0x9b, 0x14, 0x89, 0x93,
	0x15, 0xdc, 0x6d, 0x64, 0x35, 0xb1, 0x5b, 0xaf, 0x83, 0x54, 0x09, 0x21, 0xf1, 0x9f, 0xb3, 0x3b,
	0x5e, 0xdb, 0xb1, 0x1b, 0x0e, 0xdc, 0x66, 0x67, 0xde, 0x7b, 0x9e, 0x79, 0x9e, 0x5d, 0xb8, 0x13,
	0xc4, 0x29, 0x4f, 0x42, 0xee, 0x2f, 0xb8, 0x10, 0xd3, 0x19, 0xb7, 0xaf, 0x92, 0x38, 0x8d, 0xb1,
	0xa7, 0xd3, 0xd6, 0xef, 0x0e, 0x80, 0x93, 0xc5, 0x27, 0x62, 0x86, 0x7b, 0xd0, 0x4e, 0x6f, 0xae,
	0xf8, 0xb0, 0xf1, 0xa4, 0xf1, 0x6c, 0xb0, 0x3f, 0xb4, 0x35, 0xcc, 0x2e, 0x21, 0xf6, 0x44, 0xd6,
	0x19, 0xa1, 0xf0, 0x04, 0x76, 0x82, 0x79, 0x2c, 0xb8, 0x1f, 0xc4, 0x51, 0xc4, 0x83, 0x34, 0x8c,
	0x23, 0x7f, 0x21, 0x66, 0xc3, 0xa6, 0x64, 0x6f, 0xec, 0x3f, 0x2c, 0xd9, 0x0a, 0xe4, 0x14, 0x18,
	0xa9, 0xc2, 0x30, 0xb8, 0x95, 0xc3, 0xf7, 0x60, 0x86, 0x91, 0xe0, 0x49, 0xea, 0xf3, 0x28, 0x4d,
	0x6e, 0x48, 0xaa, 0x45, 0x52, 0xf7, 0x0a, 0xa9, 0x23, 0x02, 0xb8, 0xaa, 0xae, 0x64, 0x06, 0x61,
	0xe5, 0x8c, 0x87, 0xb0, 0xad, 0x25, 0x12, 0x1e, 0xc4, 0xc9, 0x39, 0x69, 0xb4, 0x49, 0x63, 0x58,
	0xd3, 0x60, 0x04, 0x50, 0x22, 0x5b, 0x61, 0x35, 0x81, 0x9f, 0x00, 0x2b, 0x2a, 0x82, 0x64, 0x3a,
	0x24, 0x73, 0x7f, 0xad, 0x8c, 0x50, 0x3a, 0x66, 0x58, 0xcb, 0xa0, 0x0d, 0xc6, 0xf5, 0x92, 0xeb,
	0x51, 0xba, 0xc4, 0xdf, 0x2e, 0xf8, 0xa7, 0xaa, 0xa2, 0x78, 0xfd, 0x6b, 0x1d, 0x29, 0x07, 0x32,
	0x7c, 0xc2, 0xc5, 0x72, 0x9e, 0x12, 0xad, 0x57, 0x73, 0x80, 0x68, 0x8c, 0xea, 0xe4, 0xc0, 0x75,
	0xe5, 0x8c, 0x2f, 0x00, 0x56, 0xc8, 0x7d, 0x22, 0x63, 0x41, 0x2e, 0x79, 0x46, 0x92, 0x87, 0xd6,
	0x2f, 0x68, 0xab, 0x9f, 0x8a, 0x3b, 0x60, 0x3a, 0xc7, 0xa3, 0xb1, 0xeb, 0x3b, 0x23, 0xcf, 0x73,
	0x9d, 0xc9, 0xd1, 0xc8, 0x33, 0xff, 0x43, 0x13, 0x36, 0x8f, 0xbc, 0xb1, 0xcb, 0x26, 0xbe, 0xeb,
	0x4d, 0xd8, 0x37, 0xb3, 0x81, 0xdb, 0xf0, 0xbf, 0xce, 0x30, 0xd7, 0x19, 0xb1, 0x43, 0xb3, 0x89,
	0x08, 0x83, 0x4a, 0x6a, 0x6c, 0xb6, 0xd0, 0x80, 0xce, 0xe9, 0x99, 0x2b, 0x19, 0x6d, 0xa5, 0x41,
	0xa1, 0xac, 0x8e, 0xcf, 0x8e, 0x27, 0x66, 0x07, 0x01, 0xba, 0x3a, 0xee, 0x5a, 0x7b, 0x80, 0xb7,
	0x37, 0x04, 0xef, 0x42, 0x37, 0xe1, 0x53, 0x11, 0x47, 0xb4, 0x8c, 0x06, 0xd3, 0x27, 0x2b, 0x86,
	0xee, 0xc7, 0x70, 0x2e, 0xc7, 0xc1, 0x47, 0x00, 0x17, 0x21, 0x9f, 0x9f, 0xfb, 0xd1, 0x74, 0xc1,
	0x35, 0xca, 0xa0, 0x8c, 0x27, 0x13, 0xb2, 0xa7, 0x6c, 0x97, 0x9b, 0x54, 0xc8, 0x36, 0x76, 0x17,
	0x8c, 0x69, 0x32, 0x5b, 0x2e, 0xe4, 0x82, 0x09, 0xb9, 0x5b, 0x2d, 0xc5, 0x28, 0x12, 0xd2, 0x80,
	0xce, 0x8f, 0xe9, 0x7c, 0xc9, 0x69, 0x63, 0x0c, 0x96, 0x1d, 0x2c, 0x4f, 0xce, 0x56, 0xdd, 0x32,
	0x89, 0x4b, 0xe3, 0x4b, 0x9e, 0x75, 0xd6, 0x66, 0xd9, 0x41, 0x0e, 0xd9, 0xba, 0xe4, 0x37, 0xfa,
	0x73, 0x2a, 0x2c, 0xf5, 0x5a, 0xab, 0x7a, 0x5f, 0x60, 0xab, 0xb6, 0x81, 0x7f, 0x11, 0x7c, 0xaa,
	0x1c, 0x50, 0x10, 0x7d, 0xa1, 0xb6, 0x56, 0x7e, 0xa3, 0x4a, 0x33, 0x5d, 0xb6, 0xbe, 0x42, 0x3f,
	0x5f, 0x26, 0x7c, 0x0c, 0x1b, 0xa5, 0x29, 0x42, 0x0a, 0xaa, 0x19, 0xa1, 0x70, 0x45, 0xe0, 0x73,
	0xe8, 0x5d, 0x90, 0x7f, 0x42, 0xca, 0xb6, 0x2a, 0xb2, 0x99, 0xaf, 0x2c, 0xaf, 0x5b, 0x6f, 0x61,
	0x50, 0xdd, 0x36, 0x45, 0xd6, 0x57, 0x82, 0x94, 0xd7, 0xf4, 0x94, 0xd7, 0xad, 0x9f, 0xf2, 0x0f,
	0x53, 0x88, 0xaf, 0xa0, 0xa7, 0x2e, 0x74, 0xc8, 0x73, 0xd2, 0x6e, 0x8d, 0x64, 0xbb, 0x59, 0x99,
	0x0c, 0x66, 0x39, 0xf8, 0xc1, 0x01, 0x6c, 0xae, 0x16, 0x72, 0x83, 0x1b, 0x6b, 0x0c, 0x6e, 0xae,
	0x18, 0x7c, 0xd0, 0x7c, 0xd3, 0xb0, 0xde, 0x81, 0x59, 0xbf, 0x9f, 0xff, 0xd2, 0xfc, 0x6b, 0x30,
	0xca, 0xa1, 0x87, 0xd0, 0x13, 0xcb, 0x20, 0x90, 0x0f, 0x28, 0x7d, 0xbb, 0xcf, 0xf2, 0xa3, 0xea,
	0x28, 0x7f, 0xef, 0x64, 0x47, 0x32, 0xfc, 0xd0, 0xfc, 0xdc, 0xf8, 0xde, 0xa5, 0x37, 0xf6, 0xe5,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xc2, 0x24, 0xab, 0x7c, 0x05, 0x00, 0x00,
}
