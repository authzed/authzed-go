// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: authzed/api/v1/debug.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckDebugTrace_PermissionType int32

const (
	CheckDebugTrace_PERMISSION_TYPE_UNSPECIFIED CheckDebugTrace_PermissionType = 0
	CheckDebugTrace_PERMISSION_TYPE_RELATION    CheckDebugTrace_PermissionType = 1
	CheckDebugTrace_PERMISSION_TYPE_PERMISSION  CheckDebugTrace_PermissionType = 2
)

// Enum value maps for CheckDebugTrace_PermissionType.
var (
	CheckDebugTrace_PermissionType_name = map[int32]string{
		0: "PERMISSION_TYPE_UNSPECIFIED",
		1: "PERMISSION_TYPE_RELATION",
		2: "PERMISSION_TYPE_PERMISSION",
	}
	CheckDebugTrace_PermissionType_value = map[string]int32{
		"PERMISSION_TYPE_UNSPECIFIED": 0,
		"PERMISSION_TYPE_RELATION":    1,
		"PERMISSION_TYPE_PERMISSION":  2,
	}
)

func (x CheckDebugTrace_PermissionType) Enum() *CheckDebugTrace_PermissionType {
	p := new(CheckDebugTrace_PermissionType)
	*p = x
	return p
}

func (x CheckDebugTrace_PermissionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckDebugTrace_PermissionType) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1_debug_proto_enumTypes[0].Descriptor()
}

func (CheckDebugTrace_PermissionType) Type() protoreflect.EnumType {
	return &file_authzed_api_v1_debug_proto_enumTypes[0]
}

func (x CheckDebugTrace_PermissionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckDebugTrace_PermissionType.Descriptor instead.
func (CheckDebugTrace_PermissionType) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{1, 0}
}

type CheckDebugTrace_Permissionship int32

const (
	CheckDebugTrace_PERMISSIONSHIP_UNSPECIFIED            CheckDebugTrace_Permissionship = 0
	CheckDebugTrace_PERMISSIONSHIP_NO_PERMISSION          CheckDebugTrace_Permissionship = 1
	CheckDebugTrace_PERMISSIONSHIP_HAS_PERMISSION         CheckDebugTrace_Permissionship = 2
	CheckDebugTrace_PERMISSIONSHIP_CONDITIONAL_PERMISSION CheckDebugTrace_Permissionship = 3
)

// Enum value maps for CheckDebugTrace_Permissionship.
var (
	CheckDebugTrace_Permissionship_name = map[int32]string{
		0: "PERMISSIONSHIP_UNSPECIFIED",
		1: "PERMISSIONSHIP_NO_PERMISSION",
		2: "PERMISSIONSHIP_HAS_PERMISSION",
		3: "PERMISSIONSHIP_CONDITIONAL_PERMISSION",
	}
	CheckDebugTrace_Permissionship_value = map[string]int32{
		"PERMISSIONSHIP_UNSPECIFIED":            0,
		"PERMISSIONSHIP_NO_PERMISSION":          1,
		"PERMISSIONSHIP_HAS_PERMISSION":         2,
		"PERMISSIONSHIP_CONDITIONAL_PERMISSION": 3,
	}
)

func (x CheckDebugTrace_Permissionship) Enum() *CheckDebugTrace_Permissionship {
	p := new(CheckDebugTrace_Permissionship)
	*p = x
	return p
}

func (x CheckDebugTrace_Permissionship) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckDebugTrace_Permissionship) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1_debug_proto_enumTypes[1].Descriptor()
}

func (CheckDebugTrace_Permissionship) Type() protoreflect.EnumType {
	return &file_authzed_api_v1_debug_proto_enumTypes[1]
}

func (x CheckDebugTrace_Permissionship) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckDebugTrace_Permissionship.Descriptor instead.
func (CheckDebugTrace_Permissionship) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{1, 1}
}

type CaveatEvalInfo_Result int32

const (
	CaveatEvalInfo_RESULT_UNSPECIFIED          CaveatEvalInfo_Result = 0
	CaveatEvalInfo_RESULT_UNEVALUATED          CaveatEvalInfo_Result = 1
	CaveatEvalInfo_RESULT_FALSE                CaveatEvalInfo_Result = 2
	CaveatEvalInfo_RESULT_TRUE                 CaveatEvalInfo_Result = 3
	CaveatEvalInfo_RESULT_MISSING_SOME_CONTEXT CaveatEvalInfo_Result = 4
)

// Enum value maps for CaveatEvalInfo_Result.
var (
	CaveatEvalInfo_Result_name = map[int32]string{
		0: "RESULT_UNSPECIFIED",
		1: "RESULT_UNEVALUATED",
		2: "RESULT_FALSE",
		3: "RESULT_TRUE",
		4: "RESULT_MISSING_SOME_CONTEXT",
	}
	CaveatEvalInfo_Result_value = map[string]int32{
		"RESULT_UNSPECIFIED":          0,
		"RESULT_UNEVALUATED":          1,
		"RESULT_FALSE":                2,
		"RESULT_TRUE":                 3,
		"RESULT_MISSING_SOME_CONTEXT": 4,
	}
)

func (x CaveatEvalInfo_Result) Enum() *CaveatEvalInfo_Result {
	p := new(CaveatEvalInfo_Result)
	*p = x
	return p
}

func (x CaveatEvalInfo_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CaveatEvalInfo_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1_debug_proto_enumTypes[2].Descriptor()
}

func (CaveatEvalInfo_Result) Type() protoreflect.EnumType {
	return &file_authzed_api_v1_debug_proto_enumTypes[2]
}

func (x CaveatEvalInfo_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CaveatEvalInfo_Result.Descriptor instead.
func (CaveatEvalInfo_Result) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{2, 0}
}

// DebugInformation defines debug information returned by an API call in a footer when
// requested with a specific debugging header.
//
// The specific debug information returned will depend on the type of the API call made.
//
// See the github.com/authzed/authzed-go project for the specific header and footer names.
type DebugInformation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// check holds debug information about a check request.
	Check *CheckDebugTrace `protobuf:"bytes,1,opt,name=check,proto3" json:"check,omitempty"`
	// schema_used holds the schema used for the request.
	SchemaUsed    string `protobuf:"bytes,2,opt,name=schema_used,json=schemaUsed,proto3" json:"schema_used,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DebugInformation) Reset() {
	*x = DebugInformation{}
	mi := &file_authzed_api_v1_debug_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebugInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugInformation) ProtoMessage() {}

func (x *DebugInformation) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_debug_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugInformation.ProtoReflect.Descriptor instead.
func (*DebugInformation) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{0}
}

func (x *DebugInformation) GetCheck() *CheckDebugTrace {
	if x != nil {
		return x.Check
	}
	return nil
}

func (x *DebugInformation) GetSchemaUsed() string {
	if x != nil {
		return x.SchemaUsed
	}
	return ""
}

// CheckDebugTrace is a recursive trace of the requests made for resolving a CheckPermission
// API call.
type CheckDebugTrace struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// resource holds the resource on which the Check was performed.
	// for batched calls, the object_id field contains a comma-separated list of object IDs
	// for all the resources checked in the batch.
	Resource *ObjectReference `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// permission holds the name of the permission or relation on which the Check was performed.
	Permission string `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	// permission_type holds information indicating whether it was a permission or relation.
	PermissionType CheckDebugTrace_PermissionType `protobuf:"varint,3,opt,name=permission_type,json=permissionType,proto3,enum=authzed.api.v1.CheckDebugTrace_PermissionType" json:"permission_type,omitempty"`
	// subject holds the subject on which the Check was performed. This will be static across all calls within
	// the same Check tree.
	Subject *SubjectReference `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	// result holds the result of the Check call.
	Result CheckDebugTrace_Permissionship `protobuf:"varint,5,opt,name=result,proto3,enum=authzed.api.v1.CheckDebugTrace_Permissionship" json:"result,omitempty"`
	// caveat_evaluation_info holds information about the caveat evaluated for this step of the trace.
	CaveatEvaluationInfo *CaveatEvalInfo `protobuf:"bytes,8,opt,name=caveat_evaluation_info,json=caveatEvaluationInfo,proto3" json:"caveat_evaluation_info,omitempty"`
	// duration holds the time spent executing this Check operation.
	Duration *durationpb.Duration `protobuf:"bytes,9,opt,name=duration,proto3" json:"duration,omitempty"`
	// resolution holds information about how the problem was resolved.
	//
	// Types that are valid to be assigned to Resolution:
	//
	//	*CheckDebugTrace_WasCachedResult
	//	*CheckDebugTrace_SubProblems_
	Resolution isCheckDebugTrace_Resolution `protobuf_oneof:"resolution"`
	// optional_expires_at is the time at which at least one of the relationships used to
	// compute this result, expires (if any). This is *not* related to the caching window.
	OptionalExpiresAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=optional_expires_at,json=optionalExpiresAt,proto3" json:"optional_expires_at,omitempty"`
	// trace_operation_id is a unique identifier for this trace's operation, that will
	// be shared for all traces created for the same check operation in SpiceDB.
	//
	// In cases where SpiceDB performs automatic batching of subproblems, this ID can be used
	// to correlate work that was shared across multiple traces.
	//
	// This identifier is generated by SpiceDB, is to be considered opaque to the caller
	// and only guaranteed to be unique within the same overall Check or CheckBulk operation.
	TraceOperationId string `protobuf:"bytes,11,opt,name=trace_operation_id,json=traceOperationId,proto3" json:"trace_operation_id,omitempty"`
	// source holds the source of the result. It is of the form:
	// `<sourcetype>:<sourceid>`, where sourcetype can be, among others:
	// `spicedb`, `materialize`, etc.
	Source        string `protobuf:"bytes,12,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckDebugTrace) Reset() {
	*x = CheckDebugTrace{}
	mi := &file_authzed_api_v1_debug_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckDebugTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckDebugTrace) ProtoMessage() {}

func (x *CheckDebugTrace) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_debug_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckDebugTrace.ProtoReflect.Descriptor instead.
func (*CheckDebugTrace) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{1}
}

func (x *CheckDebugTrace) GetResource() *ObjectReference {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *CheckDebugTrace) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *CheckDebugTrace) GetPermissionType() CheckDebugTrace_PermissionType {
	if x != nil {
		return x.PermissionType
	}
	return CheckDebugTrace_PERMISSION_TYPE_UNSPECIFIED
}

func (x *CheckDebugTrace) GetSubject() *SubjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *CheckDebugTrace) GetResult() CheckDebugTrace_Permissionship {
	if x != nil {
		return x.Result
	}
	return CheckDebugTrace_PERMISSIONSHIP_UNSPECIFIED
}

func (x *CheckDebugTrace) GetCaveatEvaluationInfo() *CaveatEvalInfo {
	if x != nil {
		return x.CaveatEvaluationInfo
	}
	return nil
}

func (x *CheckDebugTrace) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *CheckDebugTrace) GetResolution() isCheckDebugTrace_Resolution {
	if x != nil {
		return x.Resolution
	}
	return nil
}

func (x *CheckDebugTrace) GetWasCachedResult() bool {
	if x != nil {
		if x, ok := x.Resolution.(*CheckDebugTrace_WasCachedResult); ok {
			return x.WasCachedResult
		}
	}
	return false
}

func (x *CheckDebugTrace) GetSubProblems() *CheckDebugTrace_SubProblems {
	if x != nil {
		if x, ok := x.Resolution.(*CheckDebugTrace_SubProblems_); ok {
			return x.SubProblems
		}
	}
	return nil
}

func (x *CheckDebugTrace) GetOptionalExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OptionalExpiresAt
	}
	return nil
}

func (x *CheckDebugTrace) GetTraceOperationId() string {
	if x != nil {
		return x.TraceOperationId
	}
	return ""
}

func (x *CheckDebugTrace) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type isCheckDebugTrace_Resolution interface {
	isCheckDebugTrace_Resolution()
}

type CheckDebugTrace_WasCachedResult struct {
	// was_cached_result, if true, indicates that the result was found in the cache and returned directly.
	WasCachedResult bool `protobuf:"varint,6,opt,name=was_cached_result,json=wasCachedResult,proto3,oneof"`
}

type CheckDebugTrace_SubProblems_ struct {
	// sub_problems holds the sub problems that were executed to resolve the answer to this Check. An empty list
	// and a permissionship of PERMISSIONSHIP_HAS_PERMISSION indicates the subject was found within this relation.
	SubProblems *CheckDebugTrace_SubProblems `protobuf:"bytes,7,opt,name=sub_problems,json=subProblems,proto3,oneof"`
}

func (*CheckDebugTrace_WasCachedResult) isCheckDebugTrace_Resolution() {}

func (*CheckDebugTrace_SubProblems_) isCheckDebugTrace_Resolution() {}

// CaveatEvalInfo holds information about a caveat expression that was evaluated.
type CaveatEvalInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// expression is the expression that was evaluated.
	Expression string `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
	// result is the result of the evaluation.
	Result CaveatEvalInfo_Result `protobuf:"varint,2,opt,name=result,proto3,enum=authzed.api.v1.CaveatEvalInfo_Result" json:"result,omitempty"`
	// context consists of any named values that were used for evaluating the caveat expression.
	Context *structpb.Struct `protobuf:"bytes,3,opt,name=context,proto3" json:"context,omitempty"`
	// partial_caveat_info holds information of a partially-evaluated caveated response, if applicable.
	PartialCaveatInfo *PartialCaveatInfo `protobuf:"bytes,4,opt,name=partial_caveat_info,json=partialCaveatInfo,proto3" json:"partial_caveat_info,omitempty"`
	// caveat_name is the name of the caveat that was executed, if applicable.
	CaveatName    string `protobuf:"bytes,5,opt,name=caveat_name,json=caveatName,proto3" json:"caveat_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CaveatEvalInfo) Reset() {
	*x = CaveatEvalInfo{}
	mi := &file_authzed_api_v1_debug_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CaveatEvalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaveatEvalInfo) ProtoMessage() {}

func (x *CaveatEvalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_debug_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaveatEvalInfo.ProtoReflect.Descriptor instead.
func (*CaveatEvalInfo) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{2}
}

func (x *CaveatEvalInfo) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *CaveatEvalInfo) GetResult() CaveatEvalInfo_Result {
	if x != nil {
		return x.Result
	}
	return CaveatEvalInfo_RESULT_UNSPECIFIED
}

func (x *CaveatEvalInfo) GetContext() *structpb.Struct {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *CaveatEvalInfo) GetPartialCaveatInfo() *PartialCaveatInfo {
	if x != nil {
		return x.PartialCaveatInfo
	}
	return nil
}

func (x *CaveatEvalInfo) GetCaveatName() string {
	if x != nil {
		return x.CaveatName
	}
	return ""
}

type CheckDebugTrace_SubProblems struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Traces        []*CheckDebugTrace     `protobuf:"bytes,1,rep,name=traces,proto3" json:"traces,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckDebugTrace_SubProblems) Reset() {
	*x = CheckDebugTrace_SubProblems{}
	mi := &file_authzed_api_v1_debug_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckDebugTrace_SubProblems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckDebugTrace_SubProblems) ProtoMessage() {}

func (x *CheckDebugTrace_SubProblems) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_debug_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckDebugTrace_SubProblems.ProtoReflect.Descriptor instead.
func (*CheckDebugTrace_SubProblems) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CheckDebugTrace_SubProblems) GetTraces() []*CheckDebugTrace {
	if x != nil {
		return x.Traces
	}
	return nil
}

var File_authzed_api_v1_debug_proto protoreflect.FileDescriptor

const file_authzed_api_v1_debug_proto_rawDesc = "" +
	"\n" +
	"\x1aauthzed/api/v1/debug.proto\x12\x0eauthzed.api.v1\x1a\x19authzed/api/v1/core.proto\x1a\x17validate/validate.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"j\n" +
	"\x10DebugInformation\x125\n" +
	"\x05check\x18\x01 \x01(\v2\x1f.authzed.api.v1.CheckDebugTraceR\x05check\x12\x1f\n" +
	"\vschema_used\x18\x02 \x01(\tR\n" +
	"schemaUsed\"\x85\t\n" +
	"\x0fCheckDebugTrace\x12E\n" +
	"\bresource\x18\x01 \x01(\v2\x1f.authzed.api.v1.ObjectReferenceB\b\xfaB\x05\x8a\x01\x02\x10\x01R\bresource\x12\x1e\n" +
	"\n" +
	"permission\x18\x02 \x01(\tR\n" +
	"permission\x12c\n" +
	"\x0fpermission_type\x18\x03 \x01(\x0e2..authzed.api.v1.CheckDebugTrace.PermissionTypeB\n" +
	"\xfaB\a\x82\x01\x04\x10\x01 \x00R\x0epermissionType\x12D\n" +
	"\asubject\x18\x04 \x01(\v2 .authzed.api.v1.SubjectReferenceB\b\xfaB\x05\x8a\x01\x02\x10\x01R\asubject\x12R\n" +
	"\x06result\x18\x05 \x01(\x0e2..authzed.api.v1.CheckDebugTrace.PermissionshipB\n" +
	"\xfaB\a\x82\x01\x04\x10\x01 \x00R\x06result\x12T\n" +
	"\x16caveat_evaluation_info\x18\b \x01(\v2\x1e.authzed.api.v1.CaveatEvalInfoR\x14caveatEvaluationInfo\x125\n" +
	"\bduration\x18\t \x01(\v2\x19.google.protobuf.DurationR\bduration\x12,\n" +
	"\x11was_cached_result\x18\x06 \x01(\bH\x00R\x0fwasCachedResult\x12P\n" +
	"\fsub_problems\x18\a \x01(\v2+.authzed.api.v1.CheckDebugTrace.SubProblemsH\x00R\vsubProblems\x12J\n" +
	"\x13optional_expires_at\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\x11optionalExpiresAt\x12,\n" +
	"\x12trace_operation_id\x18\v \x01(\tR\x10traceOperationId\x12\x16\n" +
	"\x06source\x18\f \x01(\tR\x06source\x1aF\n" +
	"\vSubProblems\x127\n" +
	"\x06traces\x18\x01 \x03(\v2\x1f.authzed.api.v1.CheckDebugTraceR\x06traces\"o\n" +
	"\x0ePermissionType\x12\x1f\n" +
	"\x1bPERMISSION_TYPE_UNSPECIFIED\x10\x00\x12\x1c\n" +
	"\x18PERMISSION_TYPE_RELATION\x10\x01\x12\x1e\n" +
	"\x1aPERMISSION_TYPE_PERMISSION\x10\x02\"\xa0\x01\n" +
	"\x0ePermissionship\x12\x1e\n" +
	"\x1aPERMISSIONSHIP_UNSPECIFIED\x10\x00\x12 \n" +
	"\x1cPERMISSIONSHIP_NO_PERMISSION\x10\x01\x12!\n" +
	"\x1dPERMISSIONSHIP_HAS_PERMISSION\x10\x02\x12)\n" +
	"%PERMISSIONSHIP_CONDITIONAL_PERMISSION\x10\x03B\x11\n" +
	"\n" +
	"resolution\x12\x03\xf8B\x01\"\x94\x03\n" +
	"\x0eCaveatEvalInfo\x12\x1e\n" +
	"\n" +
	"expression\x18\x01 \x01(\tR\n" +
	"expression\x12=\n" +
	"\x06result\x18\x02 \x01(\x0e2%.authzed.api.v1.CaveatEvalInfo.ResultR\x06result\x121\n" +
	"\acontext\x18\x03 \x01(\v2\x17.google.protobuf.StructR\acontext\x12Q\n" +
	"\x13partial_caveat_info\x18\x04 \x01(\v2!.authzed.api.v1.PartialCaveatInfoR\x11partialCaveatInfo\x12\x1f\n" +
	"\vcaveat_name\x18\x05 \x01(\tR\n" +
	"caveatName\"|\n" +
	"\x06Result\x12\x16\n" +
	"\x12RESULT_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12RESULT_UNEVALUATED\x10\x01\x12\x10\n" +
	"\fRESULT_FALSE\x10\x02\x12\x0f\n" +
	"\vRESULT_TRUE\x10\x03\x12\x1f\n" +
	"\x1bRESULT_MISSING_SOME_CONTEXT\x10\x04BJ\n" +
	"\x12com.authzed.api.v1P\x01Z2github.com/authzed/authzed-go/proto/authzed/api/v1b\x06proto3"

var (
	file_authzed_api_v1_debug_proto_rawDescOnce sync.Once
	file_authzed_api_v1_debug_proto_rawDescData []byte
)

func file_authzed_api_v1_debug_proto_rawDescGZIP() []byte {
	file_authzed_api_v1_debug_proto_rawDescOnce.Do(func() {
		file_authzed_api_v1_debug_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_authzed_api_v1_debug_proto_rawDesc), len(file_authzed_api_v1_debug_proto_rawDesc)))
	})
	return file_authzed_api_v1_debug_proto_rawDescData
}

var file_authzed_api_v1_debug_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_authzed_api_v1_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_authzed_api_v1_debug_proto_goTypes = []any{
	(CheckDebugTrace_PermissionType)(0), // 0: authzed.api.v1.CheckDebugTrace.PermissionType
	(CheckDebugTrace_Permissionship)(0), // 1: authzed.api.v1.CheckDebugTrace.Permissionship
	(CaveatEvalInfo_Result)(0),          // 2: authzed.api.v1.CaveatEvalInfo.Result
	(*DebugInformation)(nil),            // 3: authzed.api.v1.DebugInformation
	(*CheckDebugTrace)(nil),             // 4: authzed.api.v1.CheckDebugTrace
	(*CaveatEvalInfo)(nil),              // 5: authzed.api.v1.CaveatEvalInfo
	(*CheckDebugTrace_SubProblems)(nil), // 6: authzed.api.v1.CheckDebugTrace.SubProblems
	(*ObjectReference)(nil),             // 7: authzed.api.v1.ObjectReference
	(*SubjectReference)(nil),            // 8: authzed.api.v1.SubjectReference
	(*durationpb.Duration)(nil),         // 9: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),       // 10: google.protobuf.Timestamp
	(*structpb.Struct)(nil),             // 11: google.protobuf.Struct
	(*PartialCaveatInfo)(nil),           // 12: authzed.api.v1.PartialCaveatInfo
}
var file_authzed_api_v1_debug_proto_depIdxs = []int32{
	4,  // 0: authzed.api.v1.DebugInformation.check:type_name -> authzed.api.v1.CheckDebugTrace
	7,  // 1: authzed.api.v1.CheckDebugTrace.resource:type_name -> authzed.api.v1.ObjectReference
	0,  // 2: authzed.api.v1.CheckDebugTrace.permission_type:type_name -> authzed.api.v1.CheckDebugTrace.PermissionType
	8,  // 3: authzed.api.v1.CheckDebugTrace.subject:type_name -> authzed.api.v1.SubjectReference
	1,  // 4: authzed.api.v1.CheckDebugTrace.result:type_name -> authzed.api.v1.CheckDebugTrace.Permissionship
	5,  // 5: authzed.api.v1.CheckDebugTrace.caveat_evaluation_info:type_name -> authzed.api.v1.CaveatEvalInfo
	9,  // 6: authzed.api.v1.CheckDebugTrace.duration:type_name -> google.protobuf.Duration
	6,  // 7: authzed.api.v1.CheckDebugTrace.sub_problems:type_name -> authzed.api.v1.CheckDebugTrace.SubProblems
	10, // 8: authzed.api.v1.CheckDebugTrace.optional_expires_at:type_name -> google.protobuf.Timestamp
	2,  // 9: authzed.api.v1.CaveatEvalInfo.result:type_name -> authzed.api.v1.CaveatEvalInfo.Result
	11, // 10: authzed.api.v1.CaveatEvalInfo.context:type_name -> google.protobuf.Struct
	12, // 11: authzed.api.v1.CaveatEvalInfo.partial_caveat_info:type_name -> authzed.api.v1.PartialCaveatInfo
	4,  // 12: authzed.api.v1.CheckDebugTrace.SubProblems.traces:type_name -> authzed.api.v1.CheckDebugTrace
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_authzed_api_v1_debug_proto_init() }
func file_authzed_api_v1_debug_proto_init() {
	if File_authzed_api_v1_debug_proto != nil {
		return
	}
	file_authzed_api_v1_core_proto_init()
	file_authzed_api_v1_debug_proto_msgTypes[1].OneofWrappers = []any{
		(*CheckDebugTrace_WasCachedResult)(nil),
		(*CheckDebugTrace_SubProblems_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_authzed_api_v1_debug_proto_rawDesc), len(file_authzed_api_v1_debug_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authzed_api_v1_debug_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1_debug_proto_depIdxs,
		EnumInfos:         file_authzed_api_v1_debug_proto_enumTypes,
		MessageInfos:      file_authzed_api_v1_debug_proto_msgTypes,
	}.Build()
	File_authzed_api_v1_debug_proto = out.File
	file_authzed_api_v1_debug_proto_goTypes = nil
	file_authzed_api_v1_debug_proto_depIdxs = nil
}
