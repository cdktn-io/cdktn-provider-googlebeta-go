// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource


type GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigs struct {
	// The DNS name suffix of the zone being peered to, e.g., "my-internal-domain.corp.". Must end with a dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_persistent_resource#domain GoogleVertexAiPersistentResource#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The VPC network name in the target_project where the DNS zone specified by 'domain' is visible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_persistent_resource#target_network GoogleVertexAiPersistentResource#target_network}
	TargetNetwork *string `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
	// The project ID hosting the Cloud DNS managed zone that contains the 'domain'.
	//
	// The Vertex AI Service Agent requires the
	// dns.peer role on this project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_persistent_resource#target_project GoogleVertexAiPersistentResource#target_project}
	TargetProject *string `field:"required" json:"targetProject" yaml:"targetProject"`
}

