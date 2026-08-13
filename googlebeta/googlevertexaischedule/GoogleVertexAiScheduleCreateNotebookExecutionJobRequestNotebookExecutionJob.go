// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule


type GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob struct {
	// custom_environment_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#custom_environment_spec GoogleVertexAiSchedule#custom_environment_spec}
	CustomEnvironmentSpec *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec `field:"optional" json:"customEnvironmentSpec" yaml:"customEnvironmentSpec"`
	// dataform_repository_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#dataform_repository_source GoogleVertexAiSchedule#dataform_repository_source}
	DataformRepositorySource *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource `field:"optional" json:"dataformRepositorySource" yaml:"dataformRepositorySource"`
	// direct_notebook_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#direct_notebook_source GoogleVertexAiSchedule#direct_notebook_source}
	DirectNotebookSource *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource `field:"optional" json:"directNotebookSource" yaml:"directNotebookSource"`
	// The display name of the NotebookExecutionJob.
	//
	// The name can be up to 128 characters long and can consist of any UTF-8 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#display_name GoogleVertexAiSchedule#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#encryption_spec GoogleVertexAiSchedule#encryption_spec}
	EncryptionSpec *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Max running time of the execution job in seconds (default 86400s / 24 hrs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#execution_timeout GoogleVertexAiSchedule#execution_timeout}
	ExecutionTimeout *string `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// The user email to run the execution as. Only supported by Colab runtimes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#execution_user GoogleVertexAiSchedule#execution_user}
	ExecutionUser *string `field:"optional" json:"executionUser" yaml:"executionUser"`
	// gcs_notebook_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#gcs_notebook_source GoogleVertexAiSchedule#gcs_notebook_source}
	GcsNotebookSource *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource `field:"optional" json:"gcsNotebookSource" yaml:"gcsNotebookSource"`
	// The Cloud Storage location to upload the result to. Format: 'gs://bucket-name'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#gcs_output_uri GoogleVertexAiSchedule#gcs_output_uri}
	GcsOutputUri *string `field:"optional" json:"gcsOutputUri" yaml:"gcsOutputUri"`
	// The name of the kernel to use during notebook execution. If unset, the default kernel is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#kernel_name GoogleVertexAiSchedule#kernel_name}
	KernelName *string `field:"optional" json:"kernelName" yaml:"kernelName"`
	// The labels with user-defined metadata to organize NotebookExecutionJobs.
	//
	// Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "aiplatform.googleapis.com/" and are immutable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#labels GoogleVertexAiSchedule#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The NotebookRuntimeTemplate to source compute configuration from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#notebook_runtime_template_resource_name GoogleVertexAiSchedule#notebook_runtime_template_resource_name}
	NotebookRuntimeTemplateResourceName *string `field:"optional" json:"notebookRuntimeTemplateResourceName" yaml:"notebookRuntimeTemplateResourceName"`
	// The user-defined parameters to use during notebook execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#parameters GoogleVertexAiSchedule#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The service account to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#service_account GoogleVertexAiSchedule#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// workbench_runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_schedule#workbench_runtime GoogleVertexAiSchedule#workbench_runtime}
	WorkbenchRuntime *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime `field:"optional" json:"workbenchRuntime" yaml:"workbenchRuntime"`
}

