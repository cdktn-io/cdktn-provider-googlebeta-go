// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataproccluster


type GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfig struct {
	// Size of the attached disk, specified in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_dataproc_cluster#disk_size_gb GoogleDataprocCluster#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// The disk type of the attached disk. Such as "pd-ssd" or "pd-standard".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_dataproc_cluster#disk_type GoogleDataprocCluster#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Indicates how many IOPS to provision for the disk.
	//
	// This sets the number of I/O operations per second that the disk can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_dataproc_cluster#provisioned_iops GoogleDataprocCluster#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Indicates how much throughput to provision for the disk.
	//
	// This sets the number of throughput mb per second that the disk can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_dataproc_cluster#provisioned_throughput GoogleDataprocCluster#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

