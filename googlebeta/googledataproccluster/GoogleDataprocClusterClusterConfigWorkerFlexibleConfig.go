// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataproccluster


type GoogleDataprocClusterClusterConfigWorkerFlexibleConfig struct {
	// The number of vCPUs per instance.
	//
	// Must be a power of 2: 2, 4, 8, 16, or 32. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dataproc_cluster#core_count GoogleDataprocCluster#core_count}
	CoreCount *float64 `field:"required" json:"coreCount" yaml:"coreCount"`
	// Specifies the number of worker instances to create.
	//
	// Represents the desired capacity of the group and is updatable in place.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dataproc_cluster#num_instances GoogleDataprocCluster#num_instances}
	NumInstances *float64 `field:"required" json:"numInstances" yaml:"numInstances"`
	// The boot disk size in GB. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dataproc_cluster#boot_disk_size_gb GoogleDataprocCluster#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// The hardware class of instances.
	//
	// Allowed values are "STANDARD" and "PERFORMANCE". Defaults to "STANDARD". Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dataproc_cluster#hardware_class GoogleDataprocCluster#hardware_class}
	HardwareClass *string `field:"optional" json:"hardwareClass" yaml:"hardwareClass"`
}

