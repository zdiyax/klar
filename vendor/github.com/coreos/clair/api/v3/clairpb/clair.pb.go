// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v3/clairpb/clair.proto

/*
Package clairpb is a generated protocol buffer package.

It is generated from these files:
	api/v3/clairpb/clair.proto

It has these top-level messages:
	Vulnerability
	Feature
	Layer
	ClairStatus
	GetAncestryRequest
	GetAncestryResponse
	PostAncestryRequest
	PostAncestryResponse
	GetNotificationRequest
	GetNotificationResponse
	PagedVulnerableAncestries
	MarkNotificationAsReadRequest
	MarkNotificationAsReadResponse
*/
package clairpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vulnerability struct {
	// The name of the vulnerability.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name of the namespace in which the vulnerability was detected.
	NamespaceName string `protobuf:"bytes,2,opt,name=namespace_name,json=namespaceName" json:"namespace_name,omitempty"`
	// A description of the vulnerability according to the source for the namespace.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// A link to the vulnerability according to the source for the namespace.
	Link string `protobuf:"bytes,4,opt,name=link" json:"link,omitempty"`
	// How dangerous the vulnerability is.
	Severity string `protobuf:"bytes,5,opt,name=severity" json:"severity,omitempty"`
	// Namespace agnostic metadata about the vulnerability.
	Metadata string `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
	// The feature that fixes this vulnerability.
	// This field only exists when a vulnerability is a part of a Feature.
	FixedBy string `protobuf:"bytes,7,opt,name=fixed_by,json=fixedBy" json:"fixed_by,omitempty"`
	// The Features that are affected by the vulnerability.
	// This field only exists when a vulnerability is a part of a Notification.
	AffectedVersions []*Feature `protobuf:"bytes,8,rep,name=affected_versions,json=affectedVersions" json:"affected_versions,omitempty"`
}

func (m *Vulnerability) Reset()                    { *m = Vulnerability{} }
func (m *Vulnerability) String() string            { return proto.CompactTextString(m) }
func (*Vulnerability) ProtoMessage()               {}
func (*Vulnerability) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Vulnerability) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vulnerability) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *Vulnerability) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Vulnerability) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Vulnerability) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *Vulnerability) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Vulnerability) GetFixedBy() string {
	if m != nil {
		return m.FixedBy
	}
	return ""
}

func (m *Vulnerability) GetAffectedVersions() []*Feature {
	if m != nil {
		return m.AffectedVersions
	}
	return nil
}

type Feature struct {
	// The name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name of the namespace in which the feature is detected.
	NamespaceName string `protobuf:"bytes,2,opt,name=namespace_name,json=namespaceName" json:"namespace_name,omitempty"`
	// The specific version of this feature.
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	// The format used to parse version numbers for the feature.
	VersionFormat string `protobuf:"bytes,4,opt,name=version_format,json=versionFormat" json:"version_format,omitempty"`
	// The list of vulnerabilities that affect the feature.
	Vulnerabilities []*Vulnerability `protobuf:"bytes,5,rep,name=vulnerabilities" json:"vulnerabilities,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *Feature) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Feature) GetVersionFormat() string {
	if m != nil {
		return m.VersionFormat
	}
	return ""
}

func (m *Feature) GetVulnerabilities() []*Vulnerability {
	if m != nil {
		return m.Vulnerabilities
	}
	return nil
}

type Layer struct {
	// The sha256 tarsum for the layer.
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *Layer) Reset()                    { *m = Layer{} }
func (m *Layer) String() string            { return proto.CompactTextString(m) }
func (*Layer) ProtoMessage()               {}
func (*Layer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Layer) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type ClairStatus struct {
	// The configured list of feature listers used to scan an ancestry.
	Listers []string `protobuf:"bytes,1,rep,name=listers" json:"listers,omitempty"`
	// The configured list of namespace detectors used to scan an ancestry.
	Detectors []string `protobuf:"bytes,2,rep,name=detectors" json:"detectors,omitempty"`
	// The time at which the updater last ran.
	LastUpdateTime *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
}

func (m *ClairStatus) Reset()                    { *m = ClairStatus{} }
func (m *ClairStatus) String() string            { return proto.CompactTextString(m) }
func (*ClairStatus) ProtoMessage()               {}
func (*ClairStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClairStatus) GetListers() []string {
	if m != nil {
		return m.Listers
	}
	return nil
}

func (m *ClairStatus) GetDetectors() []string {
	if m != nil {
		return m.Detectors
	}
	return nil
}

func (m *ClairStatus) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

type GetAncestryRequest struct {
	// The name of the desired ancestry.
	AncestryName string `protobuf:"bytes,1,opt,name=ancestry_name,json=ancestryName" json:"ancestry_name,omitempty"`
	// Whether to include vulnerabilities or not in the response.
	WithVulnerabilities bool `protobuf:"varint,2,opt,name=with_vulnerabilities,json=withVulnerabilities" json:"with_vulnerabilities,omitempty"`
	// Whether to include features or not in the response.
	WithFeatures bool `protobuf:"varint,3,opt,name=with_features,json=withFeatures" json:"with_features,omitempty"`
}

func (m *GetAncestryRequest) Reset()                    { *m = GetAncestryRequest{} }
func (m *GetAncestryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAncestryRequest) ProtoMessage()               {}
func (*GetAncestryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetAncestryRequest) GetAncestryName() string {
	if m != nil {
		return m.AncestryName
	}
	return ""
}

func (m *GetAncestryRequest) GetWithVulnerabilities() bool {
	if m != nil {
		return m.WithVulnerabilities
	}
	return false
}

func (m *GetAncestryRequest) GetWithFeatures() bool {
	if m != nil {
		return m.WithFeatures
	}
	return false
}

type GetAncestryResponse struct {
	// The ancestry requested.
	Ancestry *GetAncestryResponse_Ancestry `protobuf:"bytes,1,opt,name=ancestry" json:"ancestry,omitempty"`
	// The status of Clair at the time of the request.
	Status *ClairStatus `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *GetAncestryResponse) Reset()                    { *m = GetAncestryResponse{} }
func (m *GetAncestryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAncestryResponse) ProtoMessage()               {}
func (*GetAncestryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetAncestryResponse) GetAncestry() *GetAncestryResponse_Ancestry {
	if m != nil {
		return m.Ancestry
	}
	return nil
}

func (m *GetAncestryResponse) GetStatus() *ClairStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetAncestryResponse_Ancestry struct {
	// The name of the desired ancestry.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The list of features present in the ancestry.
	// This will only be provided if requested.
	Features []*Feature `protobuf:"bytes,2,rep,name=features" json:"features,omitempty"`
	// The layers present in the ancestry.
	Layers []*Layer `protobuf:"bytes,3,rep,name=layers" json:"layers,omitempty"`
	// The configured list of feature listers used to scan this ancestry.
	ScannedListers []string `protobuf:"bytes,4,rep,name=scanned_listers,json=scannedListers" json:"scanned_listers,omitempty"`
	// The configured list of namespace detectors used to scan an ancestry.
	ScannedDetectors []string `protobuf:"bytes,5,rep,name=scanned_detectors,json=scannedDetectors" json:"scanned_detectors,omitempty"`
}

func (m *GetAncestryResponse_Ancestry) Reset()                    { *m = GetAncestryResponse_Ancestry{} }
func (m *GetAncestryResponse_Ancestry) String() string            { return proto.CompactTextString(m) }
func (*GetAncestryResponse_Ancestry) ProtoMessage()               {}
func (*GetAncestryResponse_Ancestry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *GetAncestryResponse_Ancestry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetAncestryResponse_Ancestry) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *GetAncestryResponse_Ancestry) GetLayers() []*Layer {
	if m != nil {
		return m.Layers
	}
	return nil
}

func (m *GetAncestryResponse_Ancestry) GetScannedListers() []string {
	if m != nil {
		return m.ScannedListers
	}
	return nil
}

func (m *GetAncestryResponse_Ancestry) GetScannedDetectors() []string {
	if m != nil {
		return m.ScannedDetectors
	}
	return nil
}

type PostAncestryRequest struct {
	// The name of the ancestry being scanned.
	// If scanning OCI images, this should be the hash of the manifest.
	AncestryName string `protobuf:"bytes,1,opt,name=ancestry_name,json=ancestryName" json:"ancestry_name,omitempty"`
	// The format of the image being uploaded.
	Format string `protobuf:"bytes,2,opt,name=format" json:"format,omitempty"`
	// The layers to be scanned for this particular ancestry.
	Layers []*PostAncestryRequest_PostLayer `protobuf:"bytes,3,rep,name=layers" json:"layers,omitempty"`
}

func (m *PostAncestryRequest) Reset()                    { *m = PostAncestryRequest{} }
func (m *PostAncestryRequest) String() string            { return proto.CompactTextString(m) }
func (*PostAncestryRequest) ProtoMessage()               {}
func (*PostAncestryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PostAncestryRequest) GetAncestryName() string {
	if m != nil {
		return m.AncestryName
	}
	return ""
}

func (m *PostAncestryRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *PostAncestryRequest) GetLayers() []*PostAncestryRequest_PostLayer {
	if m != nil {
		return m.Layers
	}
	return nil
}

type PostAncestryRequest_PostLayer struct {
	// The hash of the layer.
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	// The location of the layer (URL or filepath).
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// Any HTTP Headers that need to be used if requesting a layer over HTTP(S).
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PostAncestryRequest_PostLayer) Reset()         { *m = PostAncestryRequest_PostLayer{} }
func (m *PostAncestryRequest_PostLayer) String() string { return proto.CompactTextString(m) }
func (*PostAncestryRequest_PostLayer) ProtoMessage()    {}
func (*PostAncestryRequest_PostLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

func (m *PostAncestryRequest_PostLayer) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *PostAncestryRequest_PostLayer) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PostAncestryRequest_PostLayer) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type PostAncestryResponse struct {
	// The status of Clair at the time of the request.
	Status *ClairStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *PostAncestryResponse) Reset()                    { *m = PostAncestryResponse{} }
func (m *PostAncestryResponse) String() string            { return proto.CompactTextString(m) }
func (*PostAncestryResponse) ProtoMessage()               {}
func (*PostAncestryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PostAncestryResponse) GetStatus() *ClairStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetNotificationRequest struct {
	// The current page of previous vulnerabilities for the ancestry.
	// This will be empty when it is the first page.
	OldVulnerabilityPage string `protobuf:"bytes,1,opt,name=old_vulnerability_page,json=oldVulnerabilityPage" json:"old_vulnerability_page,omitempty"`
	// The current page of vulnerabilities for the ancestry.
	// This will be empty when it is the first page.
	NewVulnerabilityPage string `protobuf:"bytes,2,opt,name=new_vulnerability_page,json=newVulnerabilityPage" json:"new_vulnerability_page,omitempty"`
	// The requested maximum number of results per page.
	Limit int32 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	// The name of the notification being requested.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *GetNotificationRequest) Reset()                    { *m = GetNotificationRequest{} }
func (m *GetNotificationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNotificationRequest) ProtoMessage()               {}
func (*GetNotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetNotificationRequest) GetOldVulnerabilityPage() string {
	if m != nil {
		return m.OldVulnerabilityPage
	}
	return ""
}

func (m *GetNotificationRequest) GetNewVulnerabilityPage() string {
	if m != nil {
		return m.NewVulnerabilityPage
	}
	return ""
}

func (m *GetNotificationRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetNotificationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetNotificationResponse struct {
	// The notification as requested.
	Notification *GetNotificationResponse_Notification `protobuf:"bytes,1,opt,name=notification" json:"notification,omitempty"`
}

func (m *GetNotificationResponse) Reset()                    { *m = GetNotificationResponse{} }
func (m *GetNotificationResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNotificationResponse) ProtoMessage()               {}
func (*GetNotificationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetNotificationResponse) GetNotification() *GetNotificationResponse_Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type GetNotificationResponse_Notification struct {
	// The name of the requested notification.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The time at which the notification was created.
	Created string `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	// The time at which the notification was last sent out.
	Notified string `protobuf:"bytes,3,opt,name=notified" json:"notified,omitempty"`
	// The time at which a notification has been deleted.
	Deleted string `protobuf:"bytes,4,opt,name=deleted" json:"deleted,omitempty"`
	// The previous vulnerability and a paginated view of the ancestries it affects.
	Old *PagedVulnerableAncestries `protobuf:"bytes,5,opt,name=old" json:"old,omitempty"`
	// The newly updated vulnerability and a paginated view of the ancestries it affects.
	New *PagedVulnerableAncestries `protobuf:"bytes,6,opt,name=new" json:"new,omitempty"`
}

func (m *GetNotificationResponse_Notification) Reset()         { *m = GetNotificationResponse_Notification{} }
func (m *GetNotificationResponse_Notification) String() string { return proto.CompactTextString(m) }
func (*GetNotificationResponse_Notification) ProtoMessage()    {}
func (*GetNotificationResponse_Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{9, 0}
}

func (m *GetNotificationResponse_Notification) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetNotificationResponse_Notification) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *GetNotificationResponse_Notification) GetNotified() string {
	if m != nil {
		return m.Notified
	}
	return ""
}

func (m *GetNotificationResponse_Notification) GetDeleted() string {
	if m != nil {
		return m.Deleted
	}
	return ""
}

func (m *GetNotificationResponse_Notification) GetOld() *PagedVulnerableAncestries {
	if m != nil {
		return m.Old
	}
	return nil
}

func (m *GetNotificationResponse_Notification) GetNew() *PagedVulnerableAncestries {
	if m != nil {
		return m.New
	}
	return nil
}

type PagedVulnerableAncestries struct {
	// The identifier for the current page.
	CurrentPage string `protobuf:"bytes,1,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	// The token used to request the next page.
	// This will be empty when there are no more pages.
	NextPage string `protobuf:"bytes,2,opt,name=next_page,json=nextPage" json:"next_page,omitempty"`
	// The requested maximum number of results per page.
	Limit int32 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	// The vulnerability that affects a given set of ancestries.
	Vulnerability *Vulnerability `protobuf:"bytes,4,opt,name=vulnerability" json:"vulnerability,omitempty"`
	// The ancestries affected by a vulnerability.
	Ancestries []*PagedVulnerableAncestries_IndexedAncestryName `protobuf:"bytes,5,rep,name=ancestries" json:"ancestries,omitempty"`
}

func (m *PagedVulnerableAncestries) Reset()                    { *m = PagedVulnerableAncestries{} }
func (m *PagedVulnerableAncestries) String() string            { return proto.CompactTextString(m) }
func (*PagedVulnerableAncestries) ProtoMessage()               {}
func (*PagedVulnerableAncestries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PagedVulnerableAncestries) GetCurrentPage() string {
	if m != nil {
		return m.CurrentPage
	}
	return ""
}

func (m *PagedVulnerableAncestries) GetNextPage() string {
	if m != nil {
		return m.NextPage
	}
	return ""
}

func (m *PagedVulnerableAncestries) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PagedVulnerableAncestries) GetVulnerability() *Vulnerability {
	if m != nil {
		return m.Vulnerability
	}
	return nil
}

func (m *PagedVulnerableAncestries) GetAncestries() []*PagedVulnerableAncestries_IndexedAncestryName {
	if m != nil {
		return m.Ancestries
	}
	return nil
}

type PagedVulnerableAncestries_IndexedAncestryName struct {
	// The index is an ever increasing number associated with the particular ancestry.
	// This is useful if you're processing notifications, and need to keep track of the progress of paginating the results.
	Index int32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// The name of the ancestry.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PagedVulnerableAncestries_IndexedAncestryName) Reset() {
	*m = PagedVulnerableAncestries_IndexedAncestryName{}
}
func (m *PagedVulnerableAncestries_IndexedAncestryName) String() string {
	return proto.CompactTextString(m)
}
func (*PagedVulnerableAncestries_IndexedAncestryName) ProtoMessage() {}
func (*PagedVulnerableAncestries_IndexedAncestryName) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

func (m *PagedVulnerableAncestries_IndexedAncestryName) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PagedVulnerableAncestries_IndexedAncestryName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MarkNotificationAsReadRequest struct {
	// The name of the Notification that has been processed.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *MarkNotificationAsReadRequest) Reset()                    { *m = MarkNotificationAsReadRequest{} }
func (m *MarkNotificationAsReadRequest) String() string            { return proto.CompactTextString(m) }
func (*MarkNotificationAsReadRequest) ProtoMessage()               {}
func (*MarkNotificationAsReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MarkNotificationAsReadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MarkNotificationAsReadResponse struct {
}

func (m *MarkNotificationAsReadResponse) Reset()                    { *m = MarkNotificationAsReadResponse{} }
func (m *MarkNotificationAsReadResponse) String() string            { return proto.CompactTextString(m) }
func (*MarkNotificationAsReadResponse) ProtoMessage()               {}
func (*MarkNotificationAsReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*Vulnerability)(nil), "coreos.clair.Vulnerability")
	proto.RegisterType((*Feature)(nil), "coreos.clair.Feature")
	proto.RegisterType((*Layer)(nil), "coreos.clair.Layer")
	proto.RegisterType((*ClairStatus)(nil), "coreos.clair.ClairStatus")
	proto.RegisterType((*GetAncestryRequest)(nil), "coreos.clair.GetAncestryRequest")
	proto.RegisterType((*GetAncestryResponse)(nil), "coreos.clair.GetAncestryResponse")
	proto.RegisterType((*GetAncestryResponse_Ancestry)(nil), "coreos.clair.GetAncestryResponse.Ancestry")
	proto.RegisterType((*PostAncestryRequest)(nil), "coreos.clair.PostAncestryRequest")
	proto.RegisterType((*PostAncestryRequest_PostLayer)(nil), "coreos.clair.PostAncestryRequest.PostLayer")
	proto.RegisterType((*PostAncestryResponse)(nil), "coreos.clair.PostAncestryResponse")
	proto.RegisterType((*GetNotificationRequest)(nil), "coreos.clair.GetNotificationRequest")
	proto.RegisterType((*GetNotificationResponse)(nil), "coreos.clair.GetNotificationResponse")
	proto.RegisterType((*GetNotificationResponse_Notification)(nil), "coreos.clair.GetNotificationResponse.Notification")
	proto.RegisterType((*PagedVulnerableAncestries)(nil), "coreos.clair.PagedVulnerableAncestries")
	proto.RegisterType((*PagedVulnerableAncestries_IndexedAncestryName)(nil), "coreos.clair.PagedVulnerableAncestries.IndexedAncestryName")
	proto.RegisterType((*MarkNotificationAsReadRequest)(nil), "coreos.clair.MarkNotificationAsReadRequest")
	proto.RegisterType((*MarkNotificationAsReadResponse)(nil), "coreos.clair.MarkNotificationAsReadResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AncestryService service

type AncestryServiceClient interface {
	// The RPC used to read the results of scanning for a particular ancestry.
	GetAncestry(ctx context.Context, in *GetAncestryRequest, opts ...grpc.CallOption) (*GetAncestryResponse, error)
	// The RPC used to create a new scan of an ancestry.
	PostAncestry(ctx context.Context, in *PostAncestryRequest, opts ...grpc.CallOption) (*PostAncestryResponse, error)
}

type ancestryServiceClient struct {
	cc *grpc.ClientConn
}

func NewAncestryServiceClient(cc *grpc.ClientConn) AncestryServiceClient {
	return &ancestryServiceClient{cc}
}

func (c *ancestryServiceClient) GetAncestry(ctx context.Context, in *GetAncestryRequest, opts ...grpc.CallOption) (*GetAncestryResponse, error) {
	out := new(GetAncestryResponse)
	err := grpc.Invoke(ctx, "/coreos.clair.AncestryService/GetAncestry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ancestryServiceClient) PostAncestry(ctx context.Context, in *PostAncestryRequest, opts ...grpc.CallOption) (*PostAncestryResponse, error) {
	out := new(PostAncestryResponse)
	err := grpc.Invoke(ctx, "/coreos.clair.AncestryService/PostAncestry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AncestryService service

type AncestryServiceServer interface {
	// The RPC used to read the results of scanning for a particular ancestry.
	GetAncestry(context.Context, *GetAncestryRequest) (*GetAncestryResponse, error)
	// The RPC used to create a new scan of an ancestry.
	PostAncestry(context.Context, *PostAncestryRequest) (*PostAncestryResponse, error)
}

func RegisterAncestryServiceServer(s *grpc.Server, srv AncestryServiceServer) {
	s.RegisterService(&_AncestryService_serviceDesc, srv)
}

func _AncestryService_GetAncestry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAncestryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncestryServiceServer).GetAncestry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coreos.clair.AncestryService/GetAncestry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncestryServiceServer).GetAncestry(ctx, req.(*GetAncestryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AncestryService_PostAncestry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostAncestryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncestryServiceServer).PostAncestry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coreos.clair.AncestryService/PostAncestry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncestryServiceServer).PostAncestry(ctx, req.(*PostAncestryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AncestryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coreos.clair.AncestryService",
	HandlerType: (*AncestryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAncestry",
			Handler:    _AncestryService_GetAncestry_Handler,
		},
		{
			MethodName: "PostAncestry",
			Handler:    _AncestryService_PostAncestry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/clairpb/clair.proto",
}

// Client API for NotificationService service

type NotificationServiceClient interface {
	// The RPC used to get a particularly Notification.
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error)
	// The RPC used to mark a Notification as read after it has been processed.
	MarkNotificationAsRead(ctx context.Context, in *MarkNotificationAsReadRequest, opts ...grpc.CallOption) (*MarkNotificationAsReadResponse, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error) {
	out := new(GetNotificationResponse)
	err := grpc.Invoke(ctx, "/coreos.clair.NotificationService/GetNotification", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) MarkNotificationAsRead(ctx context.Context, in *MarkNotificationAsReadRequest, opts ...grpc.CallOption) (*MarkNotificationAsReadResponse, error) {
	out := new(MarkNotificationAsReadResponse)
	err := grpc.Invoke(ctx, "/coreos.clair.NotificationService/MarkNotificationAsRead", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NotificationService service

type NotificationServiceServer interface {
	// The RPC used to get a particularly Notification.
	GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationResponse, error)
	// The RPC used to mark a Notification as read after it has been processed.
	MarkNotificationAsRead(context.Context, *MarkNotificationAsReadRequest) (*MarkNotificationAsReadResponse, error)
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coreos.clair.NotificationService/GetNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotification(ctx, req.(*GetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_MarkNotificationAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkNotificationAsReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).MarkNotificationAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coreos.clair.NotificationService/MarkNotificationAsRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).MarkNotificationAsRead(ctx, req.(*MarkNotificationAsReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coreos.clair.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotification",
			Handler:    _NotificationService_GetNotification_Handler,
		},
		{
			MethodName: "MarkNotificationAsRead",
			Handler:    _NotificationService_MarkNotificationAsRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/clairpb/clair.proto",
}

func init() { proto.RegisterFile("api/v3/clairpb/clair.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0x4f, 0x73, 0xdb, 0x44,
	0x14, 0x1f, 0x29, 0x71, 0xec, 0x3c, 0x3b, 0x7f, 0xba, 0x76, 0x53, 0x45, 0x69, 0x8b, 0x23, 0xe8,
	0xb4, 0x93, 0x30, 0xf6, 0xc4, 0xe1, 0x50, 0xc2, 0x81, 0x49, 0xda, 0xa6, 0x74, 0xa6, 0x74, 0x3a,
	0x2a, 0xe4, 0x00, 0x07, 0xcf, 0x5a, 0x7a, 0x4e, 0x34, 0x91, 0x25, 0xa3, 0x5d, 0x3b, 0xf1, 0x74,
	0x7a, 0xe1, 0xca, 0x09, 0x38, 0xf0, 0x19, 0xb8, 0xf0, 0x25, 0xb8, 0x72, 0x02, 0x8e, 0x70, 0x63,
	0x06, 0xbe, 0x00, 0x77, 0x66, 0x57, 0x2b, 0x45, 0x72, 0xd4, 0x24, 0xa5, 0x27, 0xeb, 0xfd, 0xde,
	0x9f, 0x7d, 0x7f, 0x7e, 0x6f, 0x37, 0x01, 0x93, 0x0e, 0xbd, 0xf6, 0x78, 0xbb, 0xed, 0xf8, 0xd4,
	0x8b, 0x86, 0xbd, 0xf8, 0xb7, 0x35, 0x8c, 0x42, 0x1e, 0x92, 0x9a, 0x13, 0x46, 0x18, 0xb2, 0x96,
	0xc4, 0xcc, 0x77, 0x0e, 0xc3, 0xf0, 0xd0, 0xc7, 0xb6, 0xd4, 0xf5, 0x46, 0xfd, 0x36, 0xf7, 0x06,
	0xc8, 0x38, 0x1d, 0x0c, 0x63, 0x73, 0xf3, 0xa6, 0x32, 0x10, 0x11, 0x69, 0x10, 0x84, 0x9c, 0x72,
	0x2f, 0x0c, 0x58, 0xac, 0xb5, 0x7e, 0xd0, 0x61, 0xe1, 0x60, 0xe4, 0x07, 0x18, 0xd1, 0x9e, 0xe7,
	0x7b, 0x7c, 0x42, 0x08, 0xcc, 0x06, 0x74, 0x80, 0x86, 0xd6, 0xd4, 0xee, 0xcd, 0xdb, 0xf2, 0x9b,
	0xdc, 0x81, 0x45, 0xf1, 0xcb, 0x86, 0xd4, 0xc1, 0xae, 0xd4, 0xea, 0x52, 0xbb, 0x90, 0xa2, 0xcf,
	0x84, 0x59, 0x13, 0xaa, 0x2e, 0x32, 0x27, 0xf2, 0x86, 0xe2, 0x08, 0x63, 0x46, 0xda, 0x64, 0x21,
	0x11, 0xdc, 0xf7, 0x82, 0x63, 0x63, 0x36, 0x0e, 0x2e, 0xbe, 0x89, 0x09, 0x15, 0x86, 0x63, 0x8c,
	0x3c, 0x3e, 0x31, 0x4a, 0x12, 0x4f, 0x65, 0xa1, 0x1b, 0x20, 0xa7, 0x2e, 0xe5, 0xd4, 0x98, 0x8b,
	0x75, 0x89, 0x4c, 0x56, 0xa1, 0xd2, 0xf7, 0x4e, 0xd1, 0xed, 0xf6, 0x26, 0x46, 0x59, 0xea, 0xca,
	0x52, 0xde, 0x9b, 0x90, 0x3d, 0xb8, 0x46, 0xfb, 0x7d, 0x74, 0x38, 0xba, 0xdd, 0x31, 0x46, 0x4c,
	0x14, 0x6c, 0x54, 0x9a, 0x33, 0xf7, 0xaa, 0x9d, 0xeb, 0xad, 0x6c, 0xfb, 0x5a, 0xfb, 0x48, 0xf9,
	0x28, 0x42, 0x7b, 0x39, 0xb1, 0x3f, 0x50, 0xe6, 0xd6, 0x2f, 0x1a, 0x94, 0x95, 0xf6, 0x6d, 0x7a,
	0x62, 0x40, 0x59, 0x65, 0xa0, 0xfa, 0x91, 0x88, 0x22, 0x80, 0xfa, 0xec, 0xf6, 0xc3, 0x68, 0x40,
	0xb9, 0xea, 0xca, 0x82, 0x42, 0xf7, 0x25, 0x48, 0x1e, 0xc1, 0xd2, 0x38, 0x33, 0x20, 0x0f, 0x99,
	0x51, 0x92, 0x95, 0xac, 0xe5, 0x2b, 0xc9, 0x4d, 0xd1, 0x9e, 0xf6, 0xb1, 0xd6, 0xa0, 0xf4, 0x94,
	0x4e, 0x30, 0x12, 0xb5, 0x1c, 0x51, 0x76, 0x94, 0xd4, 0x22, 0xbe, 0xad, 0x6f, 0x34, 0xa8, 0x3e,
	0x10, 0x51, 0x5e, 0x70, 0xca, 0x47, 0x4c, 0x24, 0xed, 0x7b, 0x8c, 0x63, 0xc4, 0x0c, 0xad, 0x39,
	0x23, 0x92, 0x56, 0x22, 0xb9, 0x09, 0xf3, 0x2e, 0x72, 0x74, 0x78, 0x18, 0x31, 0x43, 0x97, 0xba,
	0x33, 0x80, 0x3c, 0x84, 0x65, 0x9f, 0x32, 0xde, 0x1d, 0x0d, 0x5d, 0xca, 0xb1, 0x2b, 0xa8, 0x28,
	0xab, 0xae, 0x76, 0xcc, 0x56, 0x4c, 0xc3, 0x56, 0xc2, 0xd3, 0xd6, 0x67, 0x09, 0x4f, 0xed, 0x45,
	0xe1, 0xf3, 0xb9, 0x74, 0x11, 0xa0, 0xf5, 0xad, 0x06, 0xe4, 0x31, 0xf2, 0xdd, 0xc0, 0x41, 0xc6,
	0xa3, 0x89, 0x8d, 0x5f, 0x8d, 0x90, 0x71, 0xf2, 0x2e, 0x2c, 0x50, 0x05, 0x75, 0x33, 0xd3, 0xa8,
	0x25, 0xa0, 0x6c, 0xf7, 0x16, 0x34, 0x4e, 0x3c, 0x7e, 0xd4, 0x9d, 0x6e, 0x99, 0x98, 0x4d, 0xc5,
	0xae, 0x0b, 0xdd, 0x41, 0x5e, 0x25, 0xe2, 0x4a, 0x97, 0x7e, 0x3c, 0x6c, 0x26, 0x33, 0xae, 0xd8,
	0x35, 0x01, 0x2a, 0x02, 0x30, 0xeb, 0x6f, 0x1d, 0xea, 0xb9, 0x9c, 0xd8, 0x30, 0x0c, 0x18, 0x92,
	0x7d, 0xa8, 0x24, 0xe7, 0xcb, 0x7c, 0xaa, 0x9d, 0x8d, 0xfc, 0x58, 0x0a, 0x9c, 0x5a, 0x29, 0x90,
	0xfa, 0x92, 0x2d, 0x98, 0x63, 0xb2, 0xf7, 0x32, 0xd3, 0x6a, 0x67, 0x35, 0x1f, 0x25, 0x33, 0x1c,
	0x5b, 0x19, 0x9a, 0xbf, 0x6b, 0x50, 0x49, 0x22, 0x15, 0x32, 0x74, 0x0b, 0x2a, 0x69, 0x4d, 0xfa,
	0x45, 0xe4, 0x4f, 0xcd, 0xc8, 0x26, 0xcc, 0xf9, 0x82, 0x25, 0xa2, 0x09, 0xc2, 0xa1, 0x9e, 0x77,
	0x90, 0x0c, 0xb2, 0x95, 0x09, 0xb9, 0x0b, 0x4b, 0xcc, 0xa1, 0x41, 0x80, 0x6e, 0x37, 0x61, 0xcb,
	0xac, 0x64, 0xc4, 0xa2, 0x82, 0x9f, 0x2a, 0xd2, 0x6c, 0xc2, 0xb5, 0xc4, 0xf0, 0x8c, 0x3c, 0x25,
	0x69, 0xba, 0xac, 0x14, 0x0f, 0x13, 0xdc, 0xfa, 0x53, 0x87, 0xfa, 0xf3, 0x90, 0xfd, 0xbf, 0xf1,
	0xaf, 0xc0, 0x9c, 0xda, 0xa5, 0x78, 0x19, 0x95, 0x44, 0x1e, 0x4c, 0xd5, 0xb5, 0x99, 0xaf, 0xab,
	0xe0, 0x3c, 0x89, 0xe5, 0xea, 0x35, 0x7f, 0xd6, 0x60, 0x3e, 0x45, 0x8b, 0xf6, 0x48, 0x60, 0x43,
	0xca, 0x8f, 0xd4, 0xe1, 0xf2, 0x9b, 0xd8, 0x50, 0x3e, 0x42, 0xea, 0x9e, 0x9d, 0x7d, 0xff, 0x0d,
	0xce, 0x6e, 0x7d, 0x12, 0xbb, 0x3e, 0x0a, 0x84, 0x36, 0x09, 0x64, 0xee, 0x40, 0x2d, 0xab, 0x20,
	0xcb, 0x30, 0x73, 0x8c, 0x13, 0x95, 0x8a, 0xf8, 0x24, 0x0d, 0x28, 0x8d, 0xa9, 0x3f, 0x4a, 0x2e,
	0xa5, 0x58, 0xd8, 0xd1, 0xef, 0x6b, 0xd6, 0x13, 0x68, 0xe4, 0x8f, 0x54, 0x4c, 0x3e, 0x63, 0xa0,
	0x76, 0x45, 0x06, 0x5a, 0x3f, 0x69, 0xb0, 0xf2, 0x18, 0xf9, 0xb3, 0x90, 0x7b, 0x7d, 0xcf, 0x91,
	0xef, 0x4a, 0x32, 0xad, 0x0f, 0x60, 0x25, 0xf4, 0xdd, 0xdc, 0x1a, 0x4e, 0xba, 0x43, 0x7a, 0x98,
	0x8c, 0xad, 0x11, 0xfa, 0x6e, 0xee, 0xc6, 0x7a, 0x4e, 0x0f, 0x51, 0x78, 0x05, 0x78, 0x52, 0xe4,
	0x15, 0x97, 0xd1, 0x08, 0xf0, 0xe4, 0xbc, 0x57, 0x03, 0x4a, 0xbe, 0x37, 0xf0, 0xb8, 0x5c, 0xdc,
	0x92, 0x1d, 0x0b, 0xe9, 0x46, 0xcc, 0x9e, 0x6d, 0x84, 0xf5, 0x87, 0x0e, 0x37, 0xce, 0x25, 0xac,
	0xea, 0x3f, 0x80, 0x5a, 0x90, 0xc1, 0x55, 0x17, 0x3a, 0xe7, 0xb6, 0xb9, 0xc8, 0xb9, 0x95, 0x03,
	0x73, 0x71, 0xcc, 0x7f, 0x34, 0xa8, 0x65, 0xd5, 0x85, 0xab, 0x6a, 0x40, 0xd9, 0x89, 0x90, 0x72,
	0x74, 0x55, 0xa5, 0x89, 0x28, 0x5e, 0xc0, 0x38, 0x1c, 0xba, 0xea, 0x01, 0x49, 0x65, 0xe1, 0xe5,
	0xa2, 0x8f, 0xc2, 0x2b, 0xae, 0x32, 0x11, 0xc9, 0x87, 0x30, 0x13, 0xfa, 0xae, 0x7c, 0x4e, 0xab,
	0x9d, 0xbb, 0x53, 0x84, 0xa3, 0x87, 0x98, 0xf6, 0xde, 0x47, 0x45, 0x04, 0x0f, 0x99, 0x2d, 0x7c,
	0x84, 0x6b, 0x80, 0x27, 0xf2, 0xb5, 0x7d, 0x13, 0xd7, 0x00, 0x4f, 0xac, 0x5f, 0x75, 0x58, 0x7d,
	0xad, 0x09, 0x59, 0x87, 0x9a, 0x33, 0x8a, 0x22, 0x0c, 0x78, 0x96, 0x08, 0x55, 0x85, 0xc9, 0x49,
	0xae, 0xc1, 0x7c, 0x80, 0xa7, 0x3c, 0x3b, 0xf2, 0x8a, 0x00, 0x2e, 0x18, 0xf3, 0x2e, 0x2c, 0xe4,
	0xe8, 0x22, 0x3b, 0x71, 0xc9, 0xe3, 0x98, 0xf7, 0x20, 0x5f, 0x02, 0xd0, 0x34, 0x4d, 0xf5, 0xb8,
	0x7e, 0x74, 0xc5, 0xc2, 0x5b, 0x4f, 0x02, 0x17, 0x4f, 0xd1, 0xdd, 0xcd, 0xdc, 0x42, 0x76, 0x26,
	0x9c, 0xf9, 0x31, 0xd4, 0x0b, 0x4c, 0x44, 0x31, 0x9e, 0x80, 0x65, 0x17, 0x4a, 0x76, 0x2c, 0xa4,
	0xd4, 0xd0, 0x33, 0x9c, 0xdd, 0x86, 0x5b, 0x9f, 0xd2, 0xe8, 0x38, 0x4b, 0xa1, 0x5d, 0x66, 0x23,
	0x75, 0x93, 0x55, 0x2b, 0xe0, 0x93, 0xd5, 0x84, 0xdb, 0xaf, 0x73, 0x8a, 0x19, 0xdb, 0xf9, 0x57,
	0x83, 0xa5, 0x24, 0xa3, 0x17, 0x18, 0x8d, 0x3d, 0x07, 0xc9, 0x08, 0xaa, 0x99, 0xe7, 0x8a, 0x34,
	0x2f, 0x78, 0xc9, 0xe4, 0xd1, 0xe6, 0xfa, 0xa5, 0x6f, 0x9d, 0xb5, 0xfe, 0xf5, 0x6f, 0x7f, 0x7d,
	0xaf, 0xaf, 0x91, 0xd5, 0x76, 0x72, 0x51, 0xb7, 0x5f, 0xe6, 0xee, 0xf1, 0x57, 0xe4, 0x18, 0x6a,
	0xd9, 0x1b, 0x89, 0xac, 0x5f, 0x7a, 0x41, 0x9a, 0xd6, 0x45, 0x26, 0xea, 0xe4, 0x86, 0x3c, 0x79,
	0xd1, 0x9a, 0x4f, 0x4f, 0xde, 0xd1, 0x36, 0x3a, 0x3f, 0xea, 0x50, 0xcf, 0xb6, 0x25, 0xa9, 0xfd,
	0x15, 0x2c, 0x4d, 0x2d, 0x37, 0x79, 0xef, 0x92, 0xdd, 0x8f, 0x53, 0xb9, 0x73, 0xa5, 0x1b, 0xc2,
	0xba, 0x25, 0xb3, 0xb9, 0x41, 0xae, 0xb7, 0xb3, 0xb7, 0x03, 0x6b, 0xbf, 0x8c, 0x7b, 0xf0, 0x9d,
	0x06, 0x2b, 0xc5, 0x13, 0x23, 0x53, 0x6f, 0xd5, 0x85, 0x64, 0x30, 0xdf, 0xbf, 0x9a, 0x71, 0x3e,
	0xa9, 0x8d, 0xe2, 0xa4, 0xf6, 0x6e, 0x43, 0xdd, 0x09, 0x07, 0xf9, 0x88, 0xc3, 0xde, 0x17, 0x65,
	0xf5, 0x5f, 0x49, 0x6f, 0x4e, 0xfe, 0x31, 0xb7, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x24, 0x28, 0x46, 0xae, 0x0c, 0x00, 0x00,
}
