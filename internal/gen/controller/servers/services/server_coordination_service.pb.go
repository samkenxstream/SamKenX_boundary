// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: controller/servers/services/v1/server_coordination_service.proto

package services

import (
	servers "github.com/hashicorp/boundary/internal/servers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CONNECTIONSTATUS int32

const (
	CONNECTIONSTATUS_CONNECTIONSTATUS_UNSPECIFIED CONNECTIONSTATUS = 0
	CONNECTIONSTATUS_CONNECTIONSTATUS_AUTHORIZED  CONNECTIONSTATUS = 1
	CONNECTIONSTATUS_CONNECTIONSTATUS_CONNECTED   CONNECTIONSTATUS = 2
	CONNECTIONSTATUS_CONNECTIONSTATUS_CLOSED      CONNECTIONSTATUS = 3
)

// Enum value maps for CONNECTIONSTATUS.
var (
	CONNECTIONSTATUS_name = map[int32]string{
		0: "CONNECTIONSTATUS_UNSPECIFIED",
		1: "CONNECTIONSTATUS_AUTHORIZED",
		2: "CONNECTIONSTATUS_CONNECTED",
		3: "CONNECTIONSTATUS_CLOSED",
	}
	CONNECTIONSTATUS_value = map[string]int32{
		"CONNECTIONSTATUS_UNSPECIFIED": 0,
		"CONNECTIONSTATUS_AUTHORIZED":  1,
		"CONNECTIONSTATUS_CONNECTED":   2,
		"CONNECTIONSTATUS_CLOSED":      3,
	}
)

func (x CONNECTIONSTATUS) Enum() *CONNECTIONSTATUS {
	p := new(CONNECTIONSTATUS)
	*p = x
	return p
}

func (x CONNECTIONSTATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CONNECTIONSTATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_servers_services_v1_server_coordination_service_proto_enumTypes[0].Descriptor()
}

func (CONNECTIONSTATUS) Type() protoreflect.EnumType {
	return &file_controller_servers_services_v1_server_coordination_service_proto_enumTypes[0]
}

func (x CONNECTIONSTATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CONNECTIONSTATUS.Descriptor instead.
func (CONNECTIONSTATUS) EnumDescriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{0}
}

type SESSIONSTATUS int32

const (
	SESSIONSTATUS_SESSIONSTATUS_UNSPECIFIED SESSIONSTATUS = 0
	SESSIONSTATUS_SESSIONSTATUS_PENDING     SESSIONSTATUS = 1
	SESSIONSTATUS_SESSIONSTATUS_ACTIVE      SESSIONSTATUS = 2
	SESSIONSTATUS_SESSIONSTATUS_CANCELING   SESSIONSTATUS = 3
	SESSIONSTATUS_SESSIONSTATUS_TERMINATED  SESSIONSTATUS = 4
)

// Enum value maps for SESSIONSTATUS.
var (
	SESSIONSTATUS_name = map[int32]string{
		0: "SESSIONSTATUS_UNSPECIFIED",
		1: "SESSIONSTATUS_PENDING",
		2: "SESSIONSTATUS_ACTIVE",
		3: "SESSIONSTATUS_CANCELING",
		4: "SESSIONSTATUS_TERMINATED",
	}
	SESSIONSTATUS_value = map[string]int32{
		"SESSIONSTATUS_UNSPECIFIED": 0,
		"SESSIONSTATUS_PENDING":     1,
		"SESSIONSTATUS_ACTIVE":      2,
		"SESSIONSTATUS_CANCELING":   3,
		"SESSIONSTATUS_TERMINATED":  4,
	}
)

func (x SESSIONSTATUS) Enum() *SESSIONSTATUS {
	p := new(SESSIONSTATUS)
	*p = x
	return p
}

func (x SESSIONSTATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SESSIONSTATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_servers_services_v1_server_coordination_service_proto_enumTypes[1].Descriptor()
}

func (SESSIONSTATUS) Type() protoreflect.EnumType {
	return &file_controller_servers_services_v1_server_coordination_service_proto_enumTypes[1]
}

func (x SESSIONSTATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SESSIONSTATUS.Descriptor instead.
func (SESSIONSTATUS) EnumDescriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{1}
}

type JOBTYPE int32

const (
	JOBTYPE_JOBTYPE_UNSPECIFIED JOBTYPE = 0
	JOBTYPE_JOBTYPE_SESSION     JOBTYPE = 1
)

// Enum value maps for JOBTYPE.
var (
	JOBTYPE_name = map[int32]string{
		0: "JOBTYPE_UNSPECIFIED",
		1: "JOBTYPE_SESSION",
	}
	JOBTYPE_value = map[string]int32{
		"JOBTYPE_UNSPECIFIED": 0,
		"JOBTYPE_SESSION":     1,
	}
)

func (x JOBTYPE) Enum() *JOBTYPE {
	p := new(JOBTYPE)
	*p = x
	return p
}

func (x JOBTYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JOBTYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_servers_services_v1_server_coordination_service_proto_enumTypes[2].Descriptor()
}

func (JOBTYPE) Type() protoreflect.EnumType {
	return &file_controller_servers_services_v1_server_coordination_service_proto_enumTypes[2]
}

func (x JOBTYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JOBTYPE.Descriptor instead.
func (JOBTYPE) EnumDescriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{2}
}

type CHANGETYPE int32

const (
	CHANGETYPE_CHANGETYPE_UNSPECIFIED CHANGETYPE = 0
	// Indicates to the worker to update its knowledge of the state, which may
	// cause action to be taken.
	CHANGETYPE_CHANGETYPE_UPDATE_STATE CHANGETYPE = 1
)

// Enum value maps for CHANGETYPE.
var (
	CHANGETYPE_name = map[int32]string{
		0: "CHANGETYPE_UNSPECIFIED",
		1: "CHANGETYPE_UPDATE_STATE",
	}
	CHANGETYPE_value = map[string]int32{
		"CHANGETYPE_UNSPECIFIED":  0,
		"CHANGETYPE_UPDATE_STATE": 1,
	}
)

func (x CHANGETYPE) Enum() *CHANGETYPE {
	p := new(CHANGETYPE)
	*p = x
	return p
}

func (x CHANGETYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CHANGETYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_servers_services_v1_server_coordination_service_proto_enumTypes[3].Descriptor()
}

func (CHANGETYPE) Type() protoreflect.EnumType {
	return &file_controller_servers_services_v1_server_coordination_service_proto_enumTypes[3]
}

func (x CHANGETYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CHANGETYPE.Descriptor instead.
func (CHANGETYPE) EnumDescriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{3}
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionId string           `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Status       CONNECTIONSTATUS `protobuf:"varint,2,opt,name=status,proto3,enum=controller.servers.services.v1.CONNECTIONSTATUS" json:"status,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{0}
}

func (x *Connection) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *Connection) GetStatus() CONNECTIONSTATUS {
	if x != nil {
		return x.Status
	}
	return CONNECTIONSTATUS_CONNECTIONSTATUS_UNSPECIFIED
}

type SessionJobInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId   string        `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Status      SESSIONSTATUS `protobuf:"varint,2,opt,name=status,proto3,enum=controller.servers.services.v1.SESSIONSTATUS" json:"status,omitempty"`
	Connections []*Connection `protobuf:"bytes,3,rep,name=connections,proto3" json:"connections,omitempty"`
}

func (x *SessionJobInfo) Reset() {
	*x = SessionJobInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionJobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionJobInfo) ProtoMessage() {}

func (x *SessionJobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionJobInfo.ProtoReflect.Descriptor instead.
func (*SessionJobInfo) Descriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{1}
}

func (x *SessionJobInfo) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *SessionJobInfo) GetStatus() SESSIONSTATUS {
	if x != nil {
		return x.Status
	}
	return SESSIONSTATUS_SESSIONSTATUS_UNSPECIFIED
}

func (x *SessionJobInfo) GetConnections() []*Connection {
	if x != nil {
		return x.Connections
	}
	return nil
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type JOBTYPE `protobuf:"varint,1,opt,name=type,proto3,enum=controller.servers.services.v1.JOBTYPE" json:"type,omitempty"`
	// Types that are assignable to JobInfo:
	//	*Job_SessionInfo
	JobInfo isJob_JobInfo `protobuf_oneof:"job_info"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{2}
}

func (x *Job) GetType() JOBTYPE {
	if x != nil {
		return x.Type
	}
	return JOBTYPE_JOBTYPE_UNSPECIFIED
}

func (m *Job) GetJobInfo() isJob_JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

func (x *Job) GetSessionInfo() *SessionJobInfo {
	if x, ok := x.GetJobInfo().(*Job_SessionInfo); ok {
		return x.SessionInfo
	}
	return nil
}

type isJob_JobInfo interface {
	isJob_JobInfo()
}

type Job_SessionInfo struct {
	// This value is specified when type is JOBTYPE_SESSION.
	SessionInfo *SessionJobInfo `protobuf:"bytes,2,opt,name=session_info,json=sessionInfo,proto3,oneof"`
}

func (*Job_SessionInfo) isJob_JobInfo() {}

type JobStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *JobStatus) Reset() {
	*x = JobStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatus) ProtoMessage() {}

func (x *JobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatus.ProtoReflect.Descriptor instead.
func (*JobStatus) Descriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{3}
}

func (x *JobStatus) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The worker info. We could use information from the TLS connection but this
	// is easier and going the other route doesn't provide much benefit -- if you
	// get access to the key and spoof the connection, you're already compromised.
	Worker *servers.Server `protobuf:"bytes,10,opt,name=worker,proto3" json:"worker,omitempty"`
	// Jobs which this worker wants to report the status.
	Jobs []*JobStatus `protobuf:"bytes,20,rep,name=jobs,proto3" json:"jobs,omitempty"`
	// Whether to update tags from the Server block on this RPC. We only need to
	// do this at startup or (at some point) SIGHUP, so specifying when it's
	// changed allows us to avoid constant database operations for something that
	// won't change very often, if ever.
	UpdateTags bool `protobuf:"varint,30,opt,name=update_tags,json=updateTags,proto3" json:"update_tags,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{4}
}

func (x *StatusRequest) GetWorker() *servers.Server {
	if x != nil {
		return x.Worker
	}
	return nil
}

func (x *StatusRequest) GetJobs() []*JobStatus {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *StatusRequest) GetUpdateTags() bool {
	if x != nil {
		return x.UpdateTags
	}
	return false
}

type JobChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job         *Job       `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	RequestType CHANGETYPE `protobuf:"varint,2,opt,name=request_type,json=requestType,proto3,enum=controller.servers.services.v1.CHANGETYPE" json:"request_type,omitempty"`
}

func (x *JobChangeRequest) Reset() {
	*x = JobChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobChangeRequest) ProtoMessage() {}

func (x *JobChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobChangeRequest.ProtoReflect.Descriptor instead.
func (*JobChangeRequest) Descriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{5}
}

func (x *JobChangeRequest) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *JobChangeRequest) GetRequestType() CHANGETYPE {
	if x != nil {
		return x.RequestType
	}
	return CHANGETYPE_CHANGETYPE_UNSPECIFIED
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Active controllers. This can be used for connection management.
	Controllers []*servers.Server `protobuf:"bytes,10,rep,name=controllers,proto3" json:"controllers,omitempty"`
	// List of jobs and the expected state changes.  For example, this will
	// include jobs witch change type of canceled for jobs which are active on a
	// worker but should be canceled. This could also contain a request to start a
	// job such as a worker -> worker proxy for establishing a session through an
	// enclave.
	JobsRequests []*JobChangeRequest `protobuf:"bytes,20,rep,name=jobs_requests,json=jobsRequests,proto3" json:"jobs_requests,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP(), []int{6}
}

func (x *StatusResponse) GetControllers() []*servers.Server {
	if x != nil {
		return x.Controllers
	}
	return nil
}

func (x *StatusResponse) GetJobsRequests() []*JobChangeRequest {
	if x != nil {
		return x.JobsRequests
	}
	return nil
}

var File_controller_servers_services_v1_server_coordination_service_proto protoreflect.FileDescriptor

var file_controller_servers_services_v1_server_coordination_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x23, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4c, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x03,
	0x4a, 0x6f, 0x62, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4a, 0x4f, 0x42, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x53, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x42, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35,
	0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62,
	0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0xa6, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x3d,
	0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x73, 0x22, 0x98,
	0x01, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x4d, 0x0a, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x54, 0x59, 0x50, 0x45, 0x52, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x0e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x55, 0x0a,
	0x0d, 0x6a, 0x6f, 0x62, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x6a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2a, 0x92, 0x01, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x9e, 0x01, 0x0a, 0x0d, 0x53, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x1d, 0x0a, 0x19, 0x53,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12,
	0x1b, 0x0a, 0x17, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18,
	0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x45,
	0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x37, 0x0a, 0x07, 0x4a, 0x4f,
	0x42, 0x54, 0x59, 0x50, 0x45, 0x12, 0x17, 0x0a, 0x13, 0x4a, 0x4f, 0x42, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x4a, 0x4f, 0x42, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x10, 0x01, 0x2a, 0x45, 0x0a, 0x0a, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x54, 0x59, 0x50,
	0x45, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a,
	0x17, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x01, 0x32, 0x86, 0x01, 0x0a, 0x19, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_servers_services_v1_server_coordination_service_proto_rawDescOnce sync.Once
	file_controller_servers_services_v1_server_coordination_service_proto_rawDescData = file_controller_servers_services_v1_server_coordination_service_proto_rawDesc
)

func file_controller_servers_services_v1_server_coordination_service_proto_rawDescGZIP() []byte {
	file_controller_servers_services_v1_server_coordination_service_proto_rawDescOnce.Do(func() {
		file_controller_servers_services_v1_server_coordination_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_servers_services_v1_server_coordination_service_proto_rawDescData)
	})
	return file_controller_servers_services_v1_server_coordination_service_proto_rawDescData
}

var file_controller_servers_services_v1_server_coordination_service_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_controller_servers_services_v1_server_coordination_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_controller_servers_services_v1_server_coordination_service_proto_goTypes = []interface{}{
	(CONNECTIONSTATUS)(0),    // 0: controller.servers.services.v1.CONNECTIONSTATUS
	(SESSIONSTATUS)(0),       // 1: controller.servers.services.v1.SESSIONSTATUS
	(JOBTYPE)(0),             // 2: controller.servers.services.v1.JOBTYPE
	(CHANGETYPE)(0),          // 3: controller.servers.services.v1.CHANGETYPE
	(*Connection)(nil),       // 4: controller.servers.services.v1.Connection
	(*SessionJobInfo)(nil),   // 5: controller.servers.services.v1.SessionJobInfo
	(*Job)(nil),              // 6: controller.servers.services.v1.Job
	(*JobStatus)(nil),        // 7: controller.servers.services.v1.JobStatus
	(*StatusRequest)(nil),    // 8: controller.servers.services.v1.StatusRequest
	(*JobChangeRequest)(nil), // 9: controller.servers.services.v1.JobChangeRequest
	(*StatusResponse)(nil),   // 10: controller.servers.services.v1.StatusResponse
	(*servers.Server)(nil),   // 11: controller.servers.v1.Server
}
var file_controller_servers_services_v1_server_coordination_service_proto_depIdxs = []int32{
	0,  // 0: controller.servers.services.v1.Connection.status:type_name -> controller.servers.services.v1.CONNECTIONSTATUS
	1,  // 1: controller.servers.services.v1.SessionJobInfo.status:type_name -> controller.servers.services.v1.SESSIONSTATUS
	4,  // 2: controller.servers.services.v1.SessionJobInfo.connections:type_name -> controller.servers.services.v1.Connection
	2,  // 3: controller.servers.services.v1.Job.type:type_name -> controller.servers.services.v1.JOBTYPE
	5,  // 4: controller.servers.services.v1.Job.session_info:type_name -> controller.servers.services.v1.SessionJobInfo
	6,  // 5: controller.servers.services.v1.JobStatus.job:type_name -> controller.servers.services.v1.Job
	11, // 6: controller.servers.services.v1.StatusRequest.worker:type_name -> controller.servers.v1.Server
	7,  // 7: controller.servers.services.v1.StatusRequest.jobs:type_name -> controller.servers.services.v1.JobStatus
	6,  // 8: controller.servers.services.v1.JobChangeRequest.job:type_name -> controller.servers.services.v1.Job
	3,  // 9: controller.servers.services.v1.JobChangeRequest.request_type:type_name -> controller.servers.services.v1.CHANGETYPE
	11, // 10: controller.servers.services.v1.StatusResponse.controllers:type_name -> controller.servers.v1.Server
	9,  // 11: controller.servers.services.v1.StatusResponse.jobs_requests:type_name -> controller.servers.services.v1.JobChangeRequest
	8,  // 12: controller.servers.services.v1.ServerCoordinationService.Status:input_type -> controller.servers.services.v1.StatusRequest
	10, // 13: controller.servers.services.v1.ServerCoordinationService.Status:output_type -> controller.servers.services.v1.StatusResponse
	13, // [13:14] is the sub-list for method output_type
	12, // [12:13] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_controller_servers_services_v1_server_coordination_service_proto_init() }
func file_controller_servers_services_v1_server_coordination_service_proto_init() {
	if File_controller_servers_services_v1_server_coordination_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionJobInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobChangeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_controller_servers_services_v1_server_coordination_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Job_SessionInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_servers_services_v1_server_coordination_service_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_servers_services_v1_server_coordination_service_proto_goTypes,
		DependencyIndexes: file_controller_servers_services_v1_server_coordination_service_proto_depIdxs,
		EnumInfos:         file_controller_servers_services_v1_server_coordination_service_proto_enumTypes,
		MessageInfos:      file_controller_servers_services_v1_server_coordination_service_proto_msgTypes,
	}.Build()
	File_controller_servers_services_v1_server_coordination_service_proto = out.File
	file_controller_servers_services_v1_server_coordination_service_proto_rawDesc = nil
	file_controller_servers_services_v1_server_coordination_service_proto_goTypes = nil
	file_controller_servers_services_v1_server_coordination_service_proto_depIdxs = nil
}
