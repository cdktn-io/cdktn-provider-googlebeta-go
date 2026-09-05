// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule


type GoogleVertexAiScheduleCreateNotebookExecutionJobRequest struct {
	// notebook_execution_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_schedule#notebook_execution_job GoogleVertexAiSchedule#notebook_execution_job}
	NotebookExecutionJob *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob `field:"required" json:"notebookExecutionJob" yaml:"notebookExecutionJob"`
	// The resource name of the Location to create the NotebookExecutionJob. Format: 'projects/{project}/locations/{location}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_schedule#parent GoogleVertexAiSchedule#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// User specified ID for the NotebookExecutionJob.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_schedule#notebook_execution_job_id GoogleVertexAiSchedule#notebook_execution_job_id}
	NotebookExecutionJobId *string `field:"optional" json:"notebookExecutionJobId" yaml:"notebookExecutionJobId"`
}

