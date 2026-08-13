// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabschedule


type GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob struct {
	// Required. The display name of the Notebook Execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#display_name GoogleColabSchedule#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Cloud Storage location to upload the result to. Format:'gs://bucket-name'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#gcs_output_uri GoogleColabSchedule#gcs_output_uri}
	GcsOutputUri *string `field:"required" json:"gcsOutputUri" yaml:"gcsOutputUri"`
	// custom_environment_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#custom_environment_spec GoogleColabSchedule#custom_environment_spec}
	CustomEnvironmentSpec *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec `field:"optional" json:"customEnvironmentSpec" yaml:"customEnvironmentSpec"`
	// dataform_repository_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#dataform_repository_source GoogleColabSchedule#dataform_repository_source}
	DataformRepositorySource *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource `field:"optional" json:"dataformRepositorySource" yaml:"dataformRepositorySource"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#encryption_spec GoogleColabSchedule#encryption_spec}
	EncryptionSpec *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Max running time of the execution job in seconds (default 86400s / 24 hrs).
	//
	// A duration in seconds with up to nine fractional digits, ending with "s". Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#execution_timeout GoogleColabSchedule#execution_timeout}
	ExecutionTimeout *string `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// The user email to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#execution_user GoogleColabSchedule#execution_user}
	ExecutionUser *string `field:"optional" json:"executionUser" yaml:"executionUser"`
	// gcs_notebook_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#gcs_notebook_source GoogleColabSchedule#gcs_notebook_source}
	GcsNotebookSource *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource `field:"optional" json:"gcsNotebookSource" yaml:"gcsNotebookSource"`
	// The name of the kernel to use during notebook execution. If unset, the default kernel is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#kernel_name GoogleColabSchedule#kernel_name}
	KernelName *string `field:"optional" json:"kernelName" yaml:"kernelName"`
	// The labels with user-defined metadata to organize NotebookExecutionJobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#labels GoogleColabSchedule#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The NotebookRuntimeTemplate to source compute configuration from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#notebook_runtime_template_resource_name GoogleColabSchedule#notebook_runtime_template_resource_name}
	NotebookRuntimeTemplateResourceName *string `field:"optional" json:"notebookRuntimeTemplateResourceName" yaml:"notebookRuntimeTemplateResourceName"`
	// The service account to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#service_account GoogleColabSchedule#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// workbench_runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_colab_schedule#workbench_runtime GoogleColabSchedule#workbench_runtime}
	WorkbenchRuntime *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime `field:"optional" json:"workbenchRuntime" yaml:"workbenchRuntime"`
}

