// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiScheduleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// User provided name of the Schedule.
	//
	// The name can be up to 128 characters long and can consist of any UTF-8 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#display_name GoogleVertexAiSchedule#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location of the Schedule. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#location GoogleVertexAiSchedule#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Maximum number of runs that can be started concurrently for this Schedule.
	//
	// This is the limit for starting the scheduled requests and not the execution of the operations/jobs created by the requests (if applicable).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#max_concurrent_run_count GoogleVertexAiSchedule#max_concurrent_run_count}
	MaxConcurrentRunCount *string `field:"required" json:"maxConcurrentRunCount" yaml:"maxConcurrentRunCount"`
	// Whether new scheduled runs can be queued when max_concurrent_runs limit is reached.
	//
	// If set to true, new runs will be queued instead of skipped. Default to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#allow_queueing GoogleVertexAiSchedule#allow_queueing}
	AllowQueueing interface{} `field:"optional" json:"allowQueueing" yaml:"allowQueueing"`
	// create_notebook_execution_job_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#create_notebook_execution_job_request GoogleVertexAiSchedule#create_notebook_execution_job_request}
	CreateNotebookExecutionJobRequest *GoogleVertexAiScheduleCreateNotebookExecutionJobRequest `field:"optional" json:"createNotebookExecutionJobRequest" yaml:"createNotebookExecutionJobRequest"`
	// create_pipeline_job_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#create_pipeline_job_request GoogleVertexAiSchedule#create_pipeline_job_request}
	CreatePipelineJobRequest *GoogleVertexAiScheduleCreatePipelineJobRequest `field:"optional" json:"createPipelineJobRequest" yaml:"createPipelineJobRequest"`
	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#cron GoogleVertexAiSchedule#cron}
	Cron *string `field:"optional" json:"cron" yaml:"cron"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#deletion_policy GoogleVertexAiSchedule#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Timestamp after which no new runs can be scheduled.
	//
	// If specified, The schedule will be completed when either end_time is reached or when scheduled_run_count >= max_run_count. If not specified, new runs will keep getting scheduled until this Schedule is paused or deleted. Already scheduled runs will be allowed to complete. Unset if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#end_time GoogleVertexAiSchedule#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#id GoogleVertexAiSchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the maximum number of active runs that can be executed concurrently for this Schedule.
	//
	// This limits the number of runs that can be in a non-terminal state at the same time. Currently, this field is only supported for requests of type CreatePipelineJobRequest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#max_concurrent_active_run_count GoogleVertexAiSchedule#max_concurrent_active_run_count}
	MaxConcurrentActiveRunCount *string `field:"optional" json:"maxConcurrentActiveRunCount" yaml:"maxConcurrentActiveRunCount"`
	// Maximum run count of the schedule.
	//
	// If specified, The schedule will be completed when either started_run_count >= max_run_count or when end_time is reached. If not specified, new runs will keep getting scheduled until this Schedule is paused or deleted. Already scheduled runs will be allowed to complete. Unset if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#max_run_count GoogleVertexAiSchedule#max_run_count}
	MaxRunCount *string `field:"optional" json:"maxRunCount" yaml:"maxRunCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#project GoogleVertexAiSchedule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Timestamp after which the first run can be scheduled. Default to Schedule create time if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#start_time GoogleVertexAiSchedule#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#timeouts GoogleVertexAiSchedule#timeouts}
	Timeouts *GoogleVertexAiScheduleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

