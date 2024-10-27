// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// VolumeSourceApplyConfiguration represents an declarative configuration of the VolumeSource type for use
// with apply.
type VolumeSourceApplyConfiguration struct {
	HostPath              *HostPathVolumeSourceApplyConfiguration              `json:"hostPath,omitempty"`
	EmptyDir              *EmptyDirVolumeSourceApplyConfiguration              `json:"emptyDir,omitempty"`
	GCEPersistentDisk     *GCEPersistentDiskVolumeSourceApplyConfiguration     `json:"gcePersistentDisk,omitempty"`
	AWSElasticBlockStore  *AWSElasticBlockStoreVolumeSourceApplyConfiguration  `json:"awsElasticBlockStore,omitempty"`
	GitRepo               *GitRepoVolumeSourceApplyConfiguration               `json:"gitRepo,omitempty"`
	Secret                *SecretVolumeSourceApplyConfiguration                `json:"secret,omitempty"`
	NFS                   *NFSVolumeSourceApplyConfiguration                   `json:"nfs,omitempty"`
	ISCSI                 *ISCSIVolumeSourceApplyConfiguration                 `json:"iscsi,omitempty"`
	Glusterfs             *GlusterfsVolumeSourceApplyConfiguration             `json:"glusterfs,omitempty"`
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSourceApplyConfiguration `json:"persistentVolumeClaim,omitempty"`
	RBD                   *RBDVolumeSourceApplyConfiguration                   `json:"rbd,omitempty"`
	FlexVolume            *FlexVolumeSourceApplyConfiguration                  `json:"flexVolume,omitempty"`
	Cinder                *CinderVolumeSourceApplyConfiguration                `json:"cinder,omitempty"`
	CephFS                *CephFSVolumeSourceApplyConfiguration                `json:"cephfs,omitempty"`
	Flocker               *FlockerVolumeSourceApplyConfiguration               `json:"flocker,omitempty"`
	DownwardAPI           *DownwardAPIVolumeSourceApplyConfiguration           `json:"downwardAPI,omitempty"`
	FC                    *FCVolumeSourceApplyConfiguration                    `json:"fc,omitempty"`
	AzureFile             *AzureFileVolumeSourceApplyConfiguration             `json:"azureFile,omitempty"`
	ConfigMap             *ConfigMapVolumeSourceApplyConfiguration             `json:"configMap,omitempty"`
	VsphereVolume         *VsphereVirtualDiskVolumeSourceApplyConfiguration    `json:"vsphereVolume,omitempty"`
	Quobyte               *QuobyteVolumeSourceApplyConfiguration               `json:"quobyte,omitempty"`
	AzureDisk             *AzureDiskVolumeSourceApplyConfiguration             `json:"azureDisk,omitempty"`
	PhotonPersistentDisk  *PhotonPersistentDiskVolumeSourceApplyConfiguration  `json:"photonPersistentDisk,omitempty"`
	Projected             *ProjectedVolumeSourceApplyConfiguration             `json:"projected,omitempty"`
	PortworxVolume        *PortworxVolumeSourceApplyConfiguration              `json:"portworxVolume,omitempty"`
	ScaleIO               *ScaleIOVolumeSourceApplyConfiguration               `json:"scaleIO,omitempty"`
	StorageOS             *StorageOSVolumeSourceApplyConfiguration             `json:"storageos,omitempty"`
	CSI                   *CSIVolumeSourceApplyConfiguration                   `json:"csi,omitempty"`
	Ephemeral             *EphemeralVolumeSourceApplyConfiguration             `json:"ephemeral,omitempty"`
}

// VolumeSourceApplyConfiguration constructs an declarative configuration of the VolumeSource type for use with
// apply.
func VolumeSource() *VolumeSourceApplyConfiguration {
	return &VolumeSourceApplyConfiguration{}
}

// WithHostPath sets the HostPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostPath field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithHostPath(value *HostPathVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.HostPath = value
	return b
}

// WithEmptyDir sets the EmptyDir field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EmptyDir field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithEmptyDir(value *EmptyDirVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.EmptyDir = value
	return b
}

// WithGCEPersistentDisk sets the GCEPersistentDisk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GCEPersistentDisk field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithGCEPersistentDisk(value *GCEPersistentDiskVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.GCEPersistentDisk = value
	return b
}

// WithAWSElasticBlockStore sets the AWSElasticBlockStore field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AWSElasticBlockStore field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithAWSElasticBlockStore(value *AWSElasticBlockStoreVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.AWSElasticBlockStore = value
	return b
}

// WithGitRepo sets the GitRepo field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GitRepo field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithGitRepo(value *GitRepoVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.GitRepo = value
	return b
}

// WithSecret sets the Secret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Secret field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithSecret(value *SecretVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Secret = value
	return b
}

// WithNFS sets the NFS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NFS field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithNFS(value *NFSVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.NFS = value
	return b
}

// WithISCSI sets the ISCSI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ISCSI field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithISCSI(value *ISCSIVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.ISCSI = value
	return b
}

// WithGlusterfs sets the Glusterfs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Glusterfs field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithGlusterfs(value *GlusterfsVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Glusterfs = value
	return b
}

// WithPersistentVolumeClaim sets the PersistentVolumeClaim field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PersistentVolumeClaim field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithPersistentVolumeClaim(value *PersistentVolumeClaimVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.PersistentVolumeClaim = value
	return b
}

// WithRBD sets the RBD field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RBD field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithRBD(value *RBDVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.RBD = value
	return b
}

// WithFlexVolume sets the FlexVolume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FlexVolume field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithFlexVolume(value *FlexVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.FlexVolume = value
	return b
}

// WithCinder sets the Cinder field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cinder field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithCinder(value *CinderVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Cinder = value
	return b
}

// WithCephFS sets the CephFS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CephFS field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithCephFS(value *CephFSVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.CephFS = value
	return b
}

// WithFlocker sets the Flocker field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Flocker field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithFlocker(value *FlockerVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Flocker = value
	return b
}

// WithDownwardAPI sets the DownwardAPI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DownwardAPI field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithDownwardAPI(value *DownwardAPIVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.DownwardAPI = value
	return b
}

// WithFC sets the FC field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FC field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithFC(value *FCVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.FC = value
	return b
}

// WithAzureFile sets the AzureFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AzureFile field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithAzureFile(value *AzureFileVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.AzureFile = value
	return b
}

// WithConfigMap sets the ConfigMap field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigMap field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithConfigMap(value *ConfigMapVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.ConfigMap = value
	return b
}

// WithVsphereVolume sets the VsphereVolume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VsphereVolume field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithVsphereVolume(value *VsphereVirtualDiskVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.VsphereVolume = value
	return b
}

// WithQuobyte sets the Quobyte field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Quobyte field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithQuobyte(value *QuobyteVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Quobyte = value
	return b
}

// WithAzureDisk sets the AzureDisk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AzureDisk field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithAzureDisk(value *AzureDiskVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.AzureDisk = value
	return b
}

// WithPhotonPersistentDisk sets the PhotonPersistentDisk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PhotonPersistentDisk field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithPhotonPersistentDisk(value *PhotonPersistentDiskVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.PhotonPersistentDisk = value
	return b
}

// WithProjected sets the Projected field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Projected field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithProjected(value *ProjectedVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Projected = value
	return b
}

// WithPortworxVolume sets the PortworxVolume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortworxVolume field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithPortworxVolume(value *PortworxVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.PortworxVolume = value
	return b
}

// WithScaleIO sets the ScaleIO field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScaleIO field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithScaleIO(value *ScaleIOVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.ScaleIO = value
	return b
}

// WithStorageOS sets the StorageOS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageOS field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithStorageOS(value *StorageOSVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.StorageOS = value
	return b
}

// WithCSI sets the CSI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CSI field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithCSI(value *CSIVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.CSI = value
	return b
}

// WithEphemeral sets the Ephemeral field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ephemeral field is set to the value of the last call.
func (b *VolumeSourceApplyConfiguration) WithEphemeral(value *EphemeralVolumeSourceApplyConfiguration) *VolumeSourceApplyConfiguration {
	b.Ephemeral = value
	return b
}
