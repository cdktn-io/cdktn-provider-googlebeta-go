// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentType struct {
	// encloses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#encloses GoogleDataLossPreventionContentPolicy#encloses}
	Encloses *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentTypeEncloses `field:"optional" json:"encloses" yaml:"encloses"`
	// fully_inside block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#fully_inside GoogleDataLossPreventionContentPolicy#fully_inside}
	FullyInside *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentTypeFullyInside `field:"optional" json:"fullyInside" yaml:"fullyInside"`
	// overlaps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#overlaps GoogleDataLossPreventionContentPolicy#overlaps}
	Overlaps *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsImageContainmentTypeOverlaps `field:"optional" json:"overlaps" yaml:"overlaps"`
}

