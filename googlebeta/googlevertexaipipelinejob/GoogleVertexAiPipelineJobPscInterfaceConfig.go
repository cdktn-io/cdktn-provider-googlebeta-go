// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipipelinejob


type GoogleVertexAiPipelineJobPscInterfaceConfig struct {
	// dns_peering_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_pipeline_job#dns_peering_configs GoogleVertexAiPipelineJob#dns_peering_configs}
	DnsPeeringConfigs interface{} `field:"optional" json:"dnsPeeringConfigs" yaml:"dnsPeeringConfigs"`
	// The name of the Compute Engine network attachment to attach to the resource within the region and user project.
	//
	// To specify this field, you must have already created a network attachment. This field is only used for resources using PSC-I.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_pipeline_job#network_attachment GoogleVertexAiPipelineJob#network_attachment}
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
}

