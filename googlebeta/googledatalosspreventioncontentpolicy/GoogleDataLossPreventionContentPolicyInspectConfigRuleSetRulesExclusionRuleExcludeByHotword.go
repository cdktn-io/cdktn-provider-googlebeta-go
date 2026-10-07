// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword struct {
	// hotword_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#hotword_regex GoogleDataLossPreventionContentPolicy#hotword_regex}
	HotwordRegex *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordHotwordRegex `field:"required" json:"hotwordRegex" yaml:"hotwordRegex"`
	// proximity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#proximity GoogleDataLossPreventionContentPolicy#proximity}
	Proximity *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordProximity `field:"required" json:"proximity" yaml:"proximity"`
}

