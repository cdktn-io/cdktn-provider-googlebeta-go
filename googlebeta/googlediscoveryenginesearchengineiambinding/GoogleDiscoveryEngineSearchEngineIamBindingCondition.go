// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlediscoveryenginesearchengineiambinding


type GoogleDiscoveryEngineSearchEngineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_search_engine_iam_binding#expression GoogleDiscoveryEngineSearchEngineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_search_engine_iam_binding#title GoogleDiscoveryEngineSearchEngineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_search_engine_iam_binding#description GoogleDiscoveryEngineSearchEngineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

