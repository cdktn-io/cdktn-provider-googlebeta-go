// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource


type GoogleVertexAiPersistentResourceResourcePoolsDiskSpec struct {
	// Size in GB of the boot disk (default is 100GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_persistent_resource#boot_disk_size_gb GoogleVertexAiPersistentResource#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// Type of the boot disk.
	//
	// For non-A3U machines, the default value is
	// "pd-ssd", for A3U machines, the default value is "hyperdisk-balanced".
	// Valid values: "pd-ssd" (Persistent Disk Solid State Drive),
	// "pd-standard" (Persistent Disk Hard Disk Drive) or "hyperdisk-balanced".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_persistent_resource#boot_disk_type GoogleVertexAiPersistentResource#boot_disk_type}
	BootDiskType *string `field:"optional" json:"bootDiskType" yaml:"bootDiskType"`
}

