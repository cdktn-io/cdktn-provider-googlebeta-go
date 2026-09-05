// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevectorsearchindex


type GoogleVectorSearchIndexDedicatedInfrastructure struct {
	// autoscaling_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_index#autoscaling_spec GoogleVectorSearchIndex#autoscaling_spec}
	AutoscalingSpec *GoogleVectorSearchIndexDedicatedInfrastructureAutoscalingSpec `field:"optional" json:"autoscalingSpec" yaml:"autoscalingSpec"`
	// Mode of the dedicated infrastructure. Defaults to 'PERFORMANCE_OPTIMIZED'. Possible values: ["MODE_UNSPECIFIED", "STORAGE_OPTIMIZED", "PERFORMANCE_OPTIMIZED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_index#mode GoogleVectorSearchIndex#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

