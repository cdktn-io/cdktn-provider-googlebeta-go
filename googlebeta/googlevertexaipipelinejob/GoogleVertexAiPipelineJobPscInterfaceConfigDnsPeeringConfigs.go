// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipipelinejob


type GoogleVertexAiPipelineJobPscInterfaceConfigDnsPeeringConfigs struct {
	// The DNS name suffix of the zone being peered to, e.g., "my-internal-domain.corp.". Must end with a dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_pipeline_job#domain GoogleVertexAiPipelineJob#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The VPC network name in the targetProject where the DNS zone specified by domain is visible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_pipeline_job#target_network GoogleVertexAiPipelineJob#target_network}
	TargetNetwork *string `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
	// The project ID hosting the Cloud DNS managed zone that contains the above domain.
	//
	// The Vertex AI Service Agent must be granted the dns.peer role on this project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_pipeline_job#target_project GoogleVertexAiPipelineJob#target_project}
	TargetProject *string `field:"required" json:"targetProject" yaml:"targetProject"`
}

