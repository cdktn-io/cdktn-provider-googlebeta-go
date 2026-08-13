// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabschedule


type GoogleColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig struct {
	// A path in a Cloud Storage bucket, which will be treated as the root output directory of the pipeline.
	//
	// It is used by the system to generate the paths of output artifacts. The artifact paths are generated with a sub-path pattern '{job_id}/{task_id}/{output_key}' under the specified output directory. The service account specified in this pipeline must have the 'storage.objects.get' and 'storage.objects.create' permissions for this bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#gcs_output_directory GoogleColabSchedule#gcs_output_directory}
	GcsOutputDirectory *string `field:"required" json:"gcsOutputDirectory" yaml:"gcsOutputDirectory"`
	// Possible values: PIPELINE_FAILURE_POLICY_FAIL_SLOW PIPELINE_FAILURE_POLICY_FAIL_FAST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#failure_policy GoogleColabSchedule#failure_policy}
	FailurePolicy *string `field:"optional" json:"failurePolicy" yaml:"failurePolicy"`
	// The runtime parameters of the PipelineJob.
	//
	// The parameters will be passed into PipelineJob.pipeline_spec to replace the placeholders at runtime. This field is used by pipelines built using 'PipelineJob.pipeline_spec.schema_version' 2.1.0, such as pipelines built using Kubeflow Pipelines SDK 1.9 or higher and the v2 DSL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#parameter_values GoogleColabSchedule#parameter_values}
	ParameterValues *map[string]*string `field:"optional" json:"parameterValues" yaml:"parameterValues"`
}

