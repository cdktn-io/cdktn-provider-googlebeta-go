// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule


type GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource struct {
	// The commit SHA to read repository with. If unset, the file will be read at HEAD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_vertex_ai_schedule#commit_sha GoogleVertexAiSchedule#commit_sha}
	CommitSha *string `field:"optional" json:"commitSha" yaml:"commitSha"`
	// The resource name of the Dataform Repository. Format: 'projects/{project_id}/locations/{location}/repositories/{repository_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_vertex_ai_schedule#dataform_repository_resource_name GoogleVertexAiSchedule#dataform_repository_resource_name}
	DataformRepositoryResourceName *string `field:"optional" json:"dataformRepositoryResourceName" yaml:"dataformRepositoryResourceName"`
}

