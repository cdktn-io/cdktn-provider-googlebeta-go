// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResource struct {
	// amazon_s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_data_loss_prevention_discovery_config#amazon_s3_bucket GoogleDataLossPreventionDiscoveryConfig#amazon_s3_bucket}
	AmazonS3Bucket *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3Bucket `field:"optional" json:"amazonS3Bucket" yaml:"amazonS3Bucket"`
}

