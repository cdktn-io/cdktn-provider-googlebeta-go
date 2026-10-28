// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlestoragecontrolprojectfeatureconfig


type GoogleStorageControlProjectFeatureConfigAutoAnnotateConfig struct {
	// models block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_storage_control_project_feature_config#models GoogleStorageControlProjectFeatureConfig#models}
	Models interface{} `field:"required" json:"models" yaml:"models"`
	// The geographic location (case-sensitive) where data is processed. Currently, only multi-region locations such as 'us', 'eu', or 'asia' are supported.
	//
	// To avoid duplicate billing, configure exactly one processing location for each bucket
	// and model combination. If you configure multiple processing locations for the same
	// bucket and model across different feature configurations, a separate annotation run is
	// triggered in each location, and you're billed for each resulting annotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_storage_control_project_feature_config#processing_location GoogleStorageControlProjectFeatureConfig#processing_location}
	ProcessingLocation *string `field:"required" json:"processingLocation" yaml:"processingLocation"`
}

