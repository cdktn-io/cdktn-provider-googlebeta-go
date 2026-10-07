// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig struct {
	// attached_disk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#attached_disk_config GoogleDataprocWorkflowTemplate#attached_disk_config}
	AttachedDiskConfig interface{} `field:"optional" json:"attachedDiskConfig" yaml:"attachedDiskConfig"`
	// Optional.
	//
	// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. This field is supported only if boot_disk_type is hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#boot_disk_provisioned_iops GoogleDataprocWorkflowTemplate#boot_disk_provisioned_iops}
	BootDiskProvisionedIops *float64 `field:"optional" json:"bootDiskProvisionedIops" yaml:"bootDiskProvisionedIops"`
	// Optional.
	//
	// Indicates how much throughput to provision for the disk. This sets the number of throughput mb per second that the disk can handle. Values must be greater than or equal to 1. This field is supported only if boot_disk_type is hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#boot_disk_provisioned_throughput GoogleDataprocWorkflowTemplate#boot_disk_provisioned_throughput}
	BootDiskProvisionedThroughput *float64 `field:"optional" json:"bootDiskProvisionedThroughput" yaml:"bootDiskProvisionedThroughput"`
	// Optional. Size in GB of the boot disk (default is 500GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#boot_disk_size_gb GoogleDataprocWorkflowTemplate#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// Optional.
	//
	// Type of the boot disk (default is "pd-standard"). Valid values: "pd-balanced" (Persistent Disk Balanced Solid State Drive), "pd-ssd" (Persistent Disk Solid State Drive), or "pd-standard" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#boot_disk_type GoogleDataprocWorkflowTemplate#boot_disk_type}
	BootDiskType *string `field:"optional" json:"bootDiskType" yaml:"bootDiskType"`
	// Optional.
	//
	// Interface type of local SSDs (default is "scsi"). Valid values: "scsi" (Small Computer System Interface), "nvme" (Non-Volatile Memory Express).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#local_ssd_interface GoogleDataprocWorkflowTemplate#local_ssd_interface}
	LocalSsdInterface *string `field:"optional" json:"localSsdInterface" yaml:"localSsdInterface"`
	// Optional.
	//
	// Number of attached SSDs, from 0 to 8 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#num_local_ssds GoogleDataprocWorkflowTemplate#num_local_ssds}
	NumLocalSsds *float64 `field:"optional" json:"numLocalSsds" yaml:"numLocalSsds"`
}

