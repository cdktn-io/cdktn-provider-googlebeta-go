// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabschedule


type GoogleColabScheduleCreatePipelineJobRequest struct {
	// pipeline_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_colab_schedule#pipeline_job GoogleColabSchedule#pipeline_job}
	PipelineJob *GoogleColabScheduleCreatePipelineJobRequestPipelineJob `field:"required" json:"pipelineJob" yaml:"pipelineJob"`
	// The resource name of the Location to create the PipelineJob in. Format: 'projects/{project}/locations/{location}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_colab_schedule#parent GoogleColabSchedule#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
}

