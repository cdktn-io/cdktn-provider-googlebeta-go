// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob


type GoogleVertexAiCustomJobJobSpec struct {
	// worker_pool_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#worker_pool_specs GoogleVertexAiCustomJob#worker_pool_specs}
	WorkerPoolSpecs interface{} `field:"required" json:"workerPoolSpecs" yaml:"workerPoolSpecs"`
	// base_output_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#base_output_directory GoogleVertexAiCustomJob#base_output_directory}
	BaseOutputDirectory *GoogleVertexAiCustomJobJobSpecBaseOutputDirectory `field:"optional" json:"baseOutputDirectory" yaml:"baseOutputDirectory"`
	// Whether you want Vertex AI to enable access to the customized dashboard in training chief container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#enable_dashboard_access GoogleVertexAiCustomJob#enable_dashboard_access}
	EnableDashboardAccess interface{} `field:"optional" json:"enableDashboardAccess" yaml:"enableDashboardAccess"`
	// Whether you want Vertex AI to enable interactive shell access to training containers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#enable_web_access GoogleVertexAiCustomJob#enable_web_access}
	EnableWebAccess interface{} `field:"optional" json:"enableWebAccess" yaml:"enableWebAccess"`
	// The full name of the Compute Engine network to which the Job should be peered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#network GoogleVertexAiCustomJob#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// psc_interface_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#psc_interface_config GoogleVertexAiCustomJob#psc_interface_config}
	PscInterfaceConfig *GoogleVertexAiCustomJobJobSpecPscInterfaceConfig `field:"optional" json:"pscInterfaceConfig" yaml:"pscInterfaceConfig"`
	// A list of names for the reserved ip ranges under the VPC network that can be used for this job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#reserved_ip_ranges GoogleVertexAiCustomJob#reserved_ip_ranges}
	ReservedIpRanges *[]*string `field:"optional" json:"reservedIpRanges" yaml:"reservedIpRanges"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#scheduling GoogleVertexAiCustomJob#scheduling}
	Scheduling *GoogleVertexAiCustomJobJobSpecScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// Specifies the service account for workload run-as account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#service_account GoogleVertexAiCustomJob#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// The name of a Vertex AI Tensorboard resource to which this CustomJob will upload Tensorboard logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#tensorboard GoogleVertexAiCustomJob#tensorboard}
	Tensorboard *string `field:"optional" json:"tensorboard" yaml:"tensorboard"`
}

