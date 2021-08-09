// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: types/microvm.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HostPathVolumeSource_HostPathType int32

const (
	HostPathVolumeSource_UNKNOWN HostPathVolumeSource_HostPathType = 0
	// RAW_FILE represents a file on the host to use as a source for a volume. It should be a raw fs.
	HostPathVolumeSource_RAW_FILE HostPathVolumeSource_HostPathType = 1
)

// Enum value maps for HostPathVolumeSource_HostPathType.
var (
	HostPathVolumeSource_HostPathType_name = map[int32]string{
		0: "UNKNOWN",
		1: "RAW_FILE",
	}
	HostPathVolumeSource_HostPathType_value = map[string]int32{
		"UNKNOWN":  0,
		"RAW_FILE": 1,
	}
)

func (x HostPathVolumeSource_HostPathType) Enum() *HostPathVolumeSource_HostPathType {
	p := new(HostPathVolumeSource_HostPathType)
	*p = x
	return p
}

func (x HostPathVolumeSource_HostPathType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HostPathVolumeSource_HostPathType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_microvm_proto_enumTypes[0].Descriptor()
}

func (HostPathVolumeSource_HostPathType) Type() protoreflect.EnumType {
	return &file_types_microvm_proto_enumTypes[0]
}

func (x HostPathVolumeSource_HostPathType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HostPathVolumeSource_HostPathType.Descriptor instead.
func (HostPathVolumeSource_HostPathType) EnumDescriptor() ([]byte, []int) {
	return file_types_microvm_proto_rawDescGZIP(), []int{6, 0}
}

// MicroVMSpec represents the specification for a microvm.
type MicroVMSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the identifier of the microvm.
	// If this empty at creation time a ID will be automatically generated.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Namespace is the name of the namespace the microvm belongs to.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Labels allows you to include extra data for the microvms.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// VCPU specifies how many vcpu the machine will be allocated.
	Vcpu int32 `protobuf:"varint,4,opt,name=vcpu,proto3" json:"vcpu,omitempty"`
	// MemoryInMb is the amount of memory in megabytes that the machine will be allocated.
	MemoryInMb int32 `protobuf:"varint,5,opt,name=memory_in_mb,json=memoryInMb,proto3" json:"memory_in_mb,omitempty"`
	// Kernel is the details of the kernel to use .
	Kernel *Kernel `protobuf:"bytes,6,opt,name=kernel,proto3" json:"kernel,omitempty"`
	// Initrd_image is the container image to use for the initial ramdisk.
	InitrdImage *string `protobuf:"bytes,7,opt,name=initrd_image,json=initrdImage,proto3,oneof" json:"initrd_image,omitempty"`
	// Volumes specifies the volumes to be attached to the microvm. There must be
	// at lease one volume that will act as the root volume.
	Volumes []*Volume `protobuf:"bytes,8,rep,name=volumes,proto3" json:"volumes,omitempty"`
	// Interfaces specifies the network interfaces to be attached to the microvm.
	Interfaces []*NetworkInterface `protobuf:"bytes,9,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	// CreatedAt indicates the time the microvm was created at.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// UpdatedAt indicates the time the microvm was last updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *MicroVMSpec) Reset() {
	*x = MicroVMSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_microvm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MicroVMSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroVMSpec) ProtoMessage() {}

func (x *MicroVMSpec) ProtoReflect() protoreflect.Message {
	mi := &file_types_microvm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroVMSpec.ProtoReflect.Descriptor instead.
func (*MicroVMSpec) Descriptor() ([]byte, []int) {
	return file_types_microvm_proto_rawDescGZIP(), []int{0}
}

func (x *MicroVMSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MicroVMSpec) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *MicroVMSpec) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MicroVMSpec) GetVcpu() int32 {
	if x != nil {
		return x.Vcpu
	}
	return 0
}

func (x *MicroVMSpec) GetMemoryInMb() int32 {
	if x != nil {
		return x.MemoryInMb
	}
	return 0
}

func (x *MicroVMSpec) GetKernel() *Kernel {
	if x != nil {
		return x.Kernel
	}
	return nil
}

func (x *MicroVMSpec) GetInitrdImage() string {
	if x != nil && x.InitrdImage != nil {
		return *x.InitrdImage
	}
	return ""
}

func (x *MicroVMSpec) GetVolumes() []*Volume {
	if x != nil {
		return x.Volumes
	}
	return nil
}

func (x *MicroVMSpec) GetInterfaces() []*NetworkInterface {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

func (x *MicroVMSpec) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MicroVMSpec) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Kernel represents the configuration for a kernel.
type Kernel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Image is the container image to use.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Cmdline is the kernel command line args.
	Cmdline string `protobuf:"bytes,2,opt,name=cmdline,proto3" json:"cmdline,omitempty"`
	// Filename is used to specify the name of the kernel file
	// in the Image.
	Filename *string `protobuf:"bytes,3,opt,name=filename,proto3,oneof" json:"filename,omitempty"`
}

func (x *Kernel) Reset() {
	*x = Kernel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_microvm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kernel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kernel) ProtoMessage() {}

func (x *Kernel) ProtoReflect() protoreflect.Message {
	mi := &file_types_microvm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kernel.ProtoReflect.Descriptor instead.
func (*Kernel) Descriptor() ([]byte, []int) {
	return file_types_microvm_proto_rawDescGZIP(), []int{1}
}

func (x *Kernel) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Kernel) GetCmdline() string {
	if x != nil {
		return x.Cmdline
	}
	return ""
}

func (x *Kernel) GetFilename() string {
	if x != nil && x.Filename != nil {
		return *x.Filename
	}
	return ""
}

type NetworkInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AllowMetadataReq indicates if the network interface allows metadata requests.
	AllowMetadataReq bool `protobuf:"varint,1,opt,name=allow_metadata_req,json=allowMetadataReq,proto3" json:"allow_metadata_req,omitempty"`
	// GuestMAC allows the specifying of a specifi MAC address to use for the interface. If
	// not supplied a autogenerated MAC address will be used.y
	GuestMac *string `protobuf:"bytes,2,opt,name=guest_mac,json=guestMac,proto3,oneof" json:"guest_mac,omitempty"`
	// GuestDeviceName is the name of the network interface to create in the microvm. If this
	// is not supplied than a device name will be assigned automatically
	GuestDeviceName *string `protobuf:"bytes,3,opt,name=guest_device_name,json=guestDeviceName,proto3,oneof" json:"guest_device_name,omitempty"`
}

func (x *NetworkInterface) Reset() {
	*x = NetworkInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_microvm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterface) ProtoMessage() {}

func (x *NetworkInterface) ProtoReflect() protoreflect.Message {
	mi := &file_types_microvm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterface.ProtoReflect.Descriptor instead.
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return file_types_microvm_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkInterface) GetAllowMetadataReq() bool {
	if x != nil {
		return x.AllowMetadataReq
	}
	return false
}

func (x *NetworkInterface) GetGuestMac() string {
	if x != nil && x.GuestMac != nil {
		return *x.GuestMac
	}
	return ""
}

func (x *NetworkInterface) GetGuestDeviceName() string {
	if x != nil && x.GuestDeviceName != nil {
		return *x.GuestDeviceName
	}
	return ""
}

// Volume represents the configuration for a volume to be attached to a microvm.
type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the uinique identifier of the volume.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// IsRoot specifies that the volume is to be used as the root volume. A machine
	// must have a root volume.
	IsRoot bool `protobuf:"varint,2,opt,name=is_root,json=isRoot,proto3" json:"is_root,omitempty"`
	// IsReadOnly specifies that the volume is to be mounted readonly.
	IsReadOnly bool `protobuf:"varint,3,opt,name=is_read_only,json=isReadOnly,proto3" json:"is_read_only,omitempty"`
	// MountPoint is the mount point for the volume in the microvm.
	MountPoint string `protobuf:"bytes,4,opt,name=mount_point,json=mountPoint,proto3" json:"mount_point,omitempty"`
	// Source is where the volume will be sourced from.
	Source *VolumeSource `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	// PartitionID is the uuid of the boot partition.
	PartitionId *string `protobuf:"bytes,6,opt,name=partition_id,json=partitionId,proto3,oneof" json:"partition_id,omitempty"`
	// Size is the size to resize this volume to.
	SizeInMb *int32 `protobuf:"varint,7,opt,name=size_in_mb,json=sizeInMb,proto3,oneof" json:"size_in_mb,omitempty"` // TODO: add rate limiting
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_microvm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_types_microvm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_types_microvm_proto_rawDescGZIP(), []int{3}
}

func (x *Volume) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Volume) GetIsRoot() bool {
	if x != nil {
		return x.IsRoot
	}
	return false
}

func (x *Volume) GetIsReadOnly() bool {
	if x != nil {
		return x.IsReadOnly
	}
	return false
}

func (x *Volume) GetMountPoint() string {
	if x != nil {
		return x.MountPoint
	}
	return ""
}

func (x *Volume) GetSource() *VolumeSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Volume) GetPartitionId() string {
	if x != nil && x.PartitionId != nil {
		return *x.PartitionId
	}
	return ""
}

func (x *Volume) GetSizeInMb() int32 {
	if x != nil && x.SizeInMb != nil {
		return *x.SizeInMb
	}
	return 0
}

// VolumeSource is the source of a volume. Based loosely on the volumes in Kubernetes Pod specs.
type VolumeSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Container is used to specify a source of a volume as a OCI container.
	ContainerSource *string `protobuf:"bytes,1,opt,name=container_source,json=containerSource,proto3,oneof" json:"container_source,omitempty"`
	// HostPath is used to specify a source of a volume as a file/device on the host.
	HostpathSource *HostPathVolumeSource `protobuf:"bytes,2,opt,name=hostpath_source,json=hostpathSource,proto3,oneof" json:"hostpath_source,omitempty"` //TODO: add CSI
}

func (x *VolumeSource) Reset() {
	*x = VolumeSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_microvm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeSource) ProtoMessage() {}

func (x *VolumeSource) ProtoReflect() protoreflect.Message {
	mi := &file_types_microvm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeSource.ProtoReflect.Descriptor instead.
func (*VolumeSource) Descriptor() ([]byte, []int) {
	return file_types_microvm_proto_rawDescGZIP(), []int{4}
}

func (x *VolumeSource) GetContainerSource() string {
	if x != nil && x.ContainerSource != nil {
		return *x.ContainerSource
	}
	return ""
}

func (x *VolumeSource) GetHostpathSource() *HostPathVolumeSource {
	if x != nil {
		return x.HostpathSource
	}
	return nil
}

// ContainerVolumeSource represents the details of a volume coming from a OCI image.
type ContainerVolumeSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Image specifies teh conatiner image to use for the volume.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ContainerVolumeSource) Reset() {
	*x = ContainerVolumeSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_microvm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerVolumeSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerVolumeSource) ProtoMessage() {}

func (x *ContainerVolumeSource) ProtoReflect() protoreflect.Message {
	mi := &file_types_microvm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerVolumeSource.ProtoReflect.Descriptor instead.
func (*ContainerVolumeSource) Descriptor() ([]byte, []int) {
	return file_types_microvm_proto_rawDescGZIP(), []int{5}
}

func (x *ContainerVolumeSource) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// HostPathVolumeSource represents the details of a volume coming from a file/device on the host.
type HostPathVolumeSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path on the host of the file/device to use as a source for a volume.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Type is the type of file/device on the host.
	Type HostPathVolumeSource_HostPathType `protobuf:"varint,2,opt,name=type,proto3,enum=reignite.types.HostPathVolumeSource_HostPathType" json:"type,omitempty"`
}

func (x *HostPathVolumeSource) Reset() {
	*x = HostPathVolumeSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_microvm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostPathVolumeSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostPathVolumeSource) ProtoMessage() {}

func (x *HostPathVolumeSource) ProtoReflect() protoreflect.Message {
	mi := &file_types_microvm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostPathVolumeSource.ProtoReflect.Descriptor instead.
func (*HostPathVolumeSource) Descriptor() ([]byte, []int) {
	return file_types_microvm_proto_rawDescGZIP(), []int{6}
}

func (x *HostPathVolumeSource) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HostPathVolumeSource) GetType() HostPathVolumeSource_HostPathType {
	if x != nil {
		return x.Type
	}
	return HostPathVolumeSource_UNKNOWN
}

var File_types_microvm_proto protoreflect.FileDescriptor

var file_types_microvm_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x04, 0x0a, 0x0b, 0x4d, 0x69, 0x63, 0x72, 0x6f,
	0x56, 0x4d, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x56, 0x4d, 0x53, 0x70, 0x65,
	0x63, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x63, 0x70, 0x75, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x63, 0x70, 0x75, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x4d, 0x62, 0x12, 0x2e, 0x0a, 0x06, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65,
	0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x52, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0c, 0x69,
	0x6e, 0x69, 0x74, 0x72, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x72, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x07, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x69, 0x67,
	0x6e, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x69, 0x6e, 0x69,
	0x74, 0x72, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x66, 0x0a, 0x06, 0x4b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6d, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6d, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xb7, 0x01, 0x0a, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x09, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x67, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x61, 0x63, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x0f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x6d, 0x61, 0x63, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x95, 0x02, 0x0a, 0x06,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x62, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x4d, 0x62,
	0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6e,
	0x5f, 0x6d, 0x62, 0x22, 0xbb, 0x01, 0x0a, 0x0c, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x52, 0x0a, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x72, 0x65, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48,
	0x6f, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x48, 0x01, 0x52, 0x0e, 0x68, 0x6f, 0x73, 0x74, 0x70, 0x61, 0x74, 0x68, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x2d, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x9c, 0x01, 0x0a, 0x14, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x45, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x72, 0x65,
	0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x0c, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x41, 0x57, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65,
	0x61, 0x76, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x69, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_microvm_proto_rawDescOnce sync.Once
	file_types_microvm_proto_rawDescData = file_types_microvm_proto_rawDesc
)

func file_types_microvm_proto_rawDescGZIP() []byte {
	file_types_microvm_proto_rawDescOnce.Do(func() {
		file_types_microvm_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_microvm_proto_rawDescData)
	})
	return file_types_microvm_proto_rawDescData
}

var file_types_microvm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_microvm_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_types_microvm_proto_goTypes = []interface{}{
	(HostPathVolumeSource_HostPathType)(0), // 0: reignite.types.HostPathVolumeSource.HostPathType
	(*MicroVMSpec)(nil),                    // 1: reignite.types.MicroVMSpec
	(*Kernel)(nil),                         // 2: reignite.types.Kernel
	(*NetworkInterface)(nil),               // 3: reignite.types.NetworkInterface
	(*Volume)(nil),                         // 4: reignite.types.Volume
	(*VolumeSource)(nil),                   // 5: reignite.types.VolumeSource
	(*ContainerVolumeSource)(nil),          // 6: reignite.types.ContainerVolumeSource
	(*HostPathVolumeSource)(nil),           // 7: reignite.types.HostPathVolumeSource
	nil,                                    // 8: reignite.types.MicroVMSpec.LabelsEntry
	(*timestamppb.Timestamp)(nil),          // 9: google.protobuf.Timestamp
}
var file_types_microvm_proto_depIdxs = []int32{
	8, // 0: reignite.types.MicroVMSpec.labels:type_name -> reignite.types.MicroVMSpec.LabelsEntry
	2, // 1: reignite.types.MicroVMSpec.kernel:type_name -> reignite.types.Kernel
	4, // 2: reignite.types.MicroVMSpec.volumes:type_name -> reignite.types.Volume
	3, // 3: reignite.types.MicroVMSpec.interfaces:type_name -> reignite.types.NetworkInterface
	9, // 4: reignite.types.MicroVMSpec.created_at:type_name -> google.protobuf.Timestamp
	9, // 5: reignite.types.MicroVMSpec.updated_at:type_name -> google.protobuf.Timestamp
	5, // 6: reignite.types.Volume.source:type_name -> reignite.types.VolumeSource
	7, // 7: reignite.types.VolumeSource.hostpath_source:type_name -> reignite.types.HostPathVolumeSource
	0, // 8: reignite.types.HostPathVolumeSource.type:type_name -> reignite.types.HostPathVolumeSource.HostPathType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_types_microvm_proto_init() }
func file_types_microvm_proto_init() {
	if File_types_microvm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_microvm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MicroVMSpec); i {
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
		file_types_microvm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kernel); i {
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
		file_types_microvm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInterface); i {
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
		file_types_microvm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
		file_types_microvm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeSource); i {
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
		file_types_microvm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerVolumeSource); i {
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
		file_types_microvm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostPathVolumeSource); i {
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
	file_types_microvm_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_types_microvm_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_types_microvm_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_types_microvm_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_types_microvm_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_microvm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_microvm_proto_goTypes,
		DependencyIndexes: file_types_microvm_proto_depIdxs,
		EnumInfos:         file_types_microvm_proto_enumTypes,
		MessageInfos:      file_types_microvm_proto_msgTypes,
	}.Build()
	File_types_microvm_proto = out.File
	file_types_microvm_proto_rawDesc = nil
	file_types_microvm_proto_goTypes = nil
	file_types_microvm_proto_depIdxs = nil
}
