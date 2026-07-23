// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule


type GoogleVertexAiScheduleCreatePipelineJobRequest struct {
	// The resource name of the Location to create the PipelineJob in. Format: 'projects/{project}/locations/{location}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_vertex_ai_schedule#parent GoogleVertexAiSchedule#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// pipeline_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_vertex_ai_schedule#pipeline_job GoogleVertexAiSchedule#pipeline_job}
	PipelineJob *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJob `field:"required" json:"pipelineJob" yaml:"pipelineJob"`
	// The ID to use for the PipelineJob, which will become the final component of the PipelineJob name.
	//
	// If not provided, an ID will be automatically generated. This value should be less than 128 characters, and valid characters are '/a-z-/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_vertex_ai_schedule#pipeline_job_id GoogleVertexAiSchedule#pipeline_job_id}
	PipelineJobId *string `field:"optional" json:"pipelineJobId" yaml:"pipelineJobId"`
}

