// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule struct {
	// likelihood_adjustment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#likelihood_adjustment GoogleDataLossPreventionContentPolicy#likelihood_adjustment}
	LikelihoodAdjustment *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment `field:"required" json:"likelihoodAdjustment" yaml:"likelihoodAdjustment"`
	// adjust_by_image_findings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#adjust_by_image_findings GoogleDataLossPreventionContentPolicy#adjust_by_image_findings}
	AdjustByImageFindings *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings `field:"optional" json:"adjustByImageFindings" yaml:"adjustByImageFindings"`
	// adjust_by_matching_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#adjust_by_matching_info_types GoogleDataLossPreventionContentPolicy#adjust_by_matching_info_types}
	AdjustByMatchingInfoTypes *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes `field:"optional" json:"adjustByMatchingInfoTypes" yaml:"adjustByMatchingInfoTypes"`
}

