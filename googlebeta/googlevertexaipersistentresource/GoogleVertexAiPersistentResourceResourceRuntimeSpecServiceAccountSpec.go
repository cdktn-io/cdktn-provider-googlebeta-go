// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource


type GoogleVertexAiPersistentResourceResourceRuntimeSpecServiceAccountSpec struct {
	// If true, custom user-managed service account is enforced to run any workloads (for example, Vertex Jobs) on the resource.
	//
	// Otherwise, uses the [Vertex AI Custom Code Service
	// Agent](https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_persistent_resource#enable_custom_service_account GoogleVertexAiPersistentResource#enable_custom_service_account}
	EnableCustomServiceAccount interface{} `field:"required" json:"enableCustomServiceAccount" yaml:"enableCustomServiceAccount"`
}

