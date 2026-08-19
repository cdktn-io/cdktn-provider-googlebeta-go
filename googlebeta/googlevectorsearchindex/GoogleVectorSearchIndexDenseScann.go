// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevectorsearchindex


type GoogleVectorSearchIndexDenseScann struct {
	// Feature norm type for the ScaNN index. Possible values: ["FEATURE_NORM_TYPE_UNSPECIFIED", "NONE", "UNIT_L2_NORM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#feature_norm_type GoogleVectorSearchIndex#feature_norm_type}
	FeatureNormType *string `field:"optional" json:"featureNormType" yaml:"featureNormType"`
}

