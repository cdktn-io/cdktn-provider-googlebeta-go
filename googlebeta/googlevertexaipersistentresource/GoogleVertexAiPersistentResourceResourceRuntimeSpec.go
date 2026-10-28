// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource


type GoogleVertexAiPersistentResourceResourceRuntimeSpec struct {
	// ray_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_persistent_resource#ray_spec GoogleVertexAiPersistentResource#ray_spec}
	RaySpec *GoogleVertexAiPersistentResourceResourceRuntimeSpecRaySpec `field:"optional" json:"raySpec" yaml:"raySpec"`
	// service_account_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_persistent_resource#service_account_spec GoogleVertexAiPersistentResource#service_account_spec}
	ServiceAccountSpec *GoogleVertexAiPersistentResourceResourceRuntimeSpecServiceAccountSpec `field:"optional" json:"serviceAccountSpec" yaml:"serviceAccountSpec"`
}

