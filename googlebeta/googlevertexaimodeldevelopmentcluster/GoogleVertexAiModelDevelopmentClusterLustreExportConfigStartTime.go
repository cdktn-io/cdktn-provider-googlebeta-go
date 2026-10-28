// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterLustreExportConfigStartTime struct {
	// Hours of a day in 24 hour format.
	//
	// Must be greater than or equal to 0 and
	// typically must be less than or equal to 23. An API may choose to allow the
	// value "24:00:00" for scenarios like business closing time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#hours GoogleVertexAiModelDevelopmentCluster#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Minutes of an hour. Must be greater than or equal to 0 and less than or equal to 59.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#minutes GoogleVertexAiModelDevelopmentCluster#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

