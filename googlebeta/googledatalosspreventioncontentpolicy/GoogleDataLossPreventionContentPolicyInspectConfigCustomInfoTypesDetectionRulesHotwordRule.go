// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRule struct {
	// hotword_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#hotword_regex GoogleDataLossPreventionContentPolicy#hotword_regex}
	HotwordRegex *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex `field:"required" json:"hotwordRegex" yaml:"hotwordRegex"`
	// likelihood_adjustment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#likelihood_adjustment GoogleDataLossPreventionContentPolicy#likelihood_adjustment}
	LikelihoodAdjustment *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment `field:"required" json:"likelihoodAdjustment" yaml:"likelihoodAdjustment"`
	// proximity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#proximity GoogleDataLossPreventionContentPolicy#proximity}
	Proximity *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity `field:"required" json:"proximity" yaml:"proximity"`
}

