// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollection struct {
	// include_regexes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_data_loss_prevention_discovery_config#include_regexes GoogleDataLossPreventionDiscoveryConfig#include_regexes}
	IncludeRegexes *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionIncludeRegexes `field:"optional" json:"includeRegexes" yaml:"includeRegexes"`
	// include_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_data_loss_prevention_discovery_config#include_tags GoogleDataLossPreventionDiscoveryConfig#include_tags}
	IncludeTags *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionIncludeTags `field:"optional" json:"includeTags" yaml:"includeTags"`
}

