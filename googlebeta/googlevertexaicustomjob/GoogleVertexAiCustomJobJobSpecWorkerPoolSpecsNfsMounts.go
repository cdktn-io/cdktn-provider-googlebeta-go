// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob


type GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsNfsMounts struct {
	// Destination mount path. The path must be absolute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#mount_point GoogleVertexAiCustomJob#mount_point}
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// Source path exported from NFS database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#path GoogleVertexAiCustomJob#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// IP address of the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#server GoogleVertexAiCustomJob#server}
	Server *string `field:"required" json:"server" yaml:"server"`
}

