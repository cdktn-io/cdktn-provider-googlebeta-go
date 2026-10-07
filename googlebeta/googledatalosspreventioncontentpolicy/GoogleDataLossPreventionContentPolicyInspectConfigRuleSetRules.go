// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRules struct {
	// adjustment_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#adjustment_rule GoogleDataLossPreventionContentPolicy#adjustment_rule}
	AdjustmentRule *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule `field:"optional" json:"adjustmentRule" yaml:"adjustmentRule"`
	// exclusion_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#exclusion_rule GoogleDataLossPreventionContentPolicy#exclusion_rule}
	ExclusionRule *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule `field:"optional" json:"exclusionRule" yaml:"exclusionRule"`
	// hotword_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#hotword_rule GoogleDataLossPreventionContentPolicy#hotword_rule}
	HotwordRule *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesHotwordRule `field:"optional" json:"hotwordRule" yaml:"hotwordRule"`
}

