// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule


type GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec struct {
	// Size in GB of the disk (default is 100GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_schedule#disk_size_gb GoogleVertexAiSchedule#disk_size_gb}
	DiskSizeGb *string `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Type of the disk (default is "pd-standard").
	//
	// Valid values: "pd-ssd" (Persistent Disk Solid State Drive) "pd-standard" (Persistent Disk Hard Disk Drive) "pd-balanced" (Balanced Persistent Disk) "pd-extreme" (Extreme Persistent Disk)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_schedule#disk_type GoogleVertexAiSchedule#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
}

