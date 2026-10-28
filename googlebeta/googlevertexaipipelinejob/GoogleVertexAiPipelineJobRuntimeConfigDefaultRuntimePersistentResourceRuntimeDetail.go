// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipipelinejob


type GoogleVertexAiPipelineJobRuntimeConfigDefaultRuntimePersistentResourceRuntimeDetail struct {
	// PersistentResource name, format: 'projects/{project}/locations/{location}/persistentResources/{persistent_resource}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_pipeline_job#persistent_resource_name GoogleVertexAiPipelineJob#persistent_resource_name}
	PersistentResourceName *string `field:"optional" json:"persistentResourceName" yaml:"persistentResourceName"`
	// Sets the behavior when a task has reached the defined wait duration while waiting for task resources to become available in the specified persistent resource.
	//
	// Possible values are: 'TASK_RESOURCE_UNAVAILABLE_TIMEOUT_BEHAVIOR_UNSPECIFIED', 'FAIL', 'FALL_BACK_TO_ON_DEMAND'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_pipeline_job#task_resource_unavailable_timeout_behavior GoogleVertexAiPipelineJob#task_resource_unavailable_timeout_behavior}
	TaskResourceUnavailableTimeoutBehavior *string `field:"optional" json:"taskResourceUnavailableTimeoutBehavior" yaml:"taskResourceUnavailableTimeoutBehavior"`
	// The maxWaitTime for the task if the persistent resource is unavailable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_pipeline_job#task_resource_unavailable_wait_time_ms GoogleVertexAiPipelineJob#task_resource_unavailable_wait_time_ms}
	TaskResourceUnavailableWaitTimeMs *string `field:"optional" json:"taskResourceUnavailableWaitTimeMs" yaml:"taskResourceUnavailableWaitTimeMs"`
}

