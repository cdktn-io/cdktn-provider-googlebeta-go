// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment struct {
	// Set the likelihood of a finding to a fixed value.
	//
	// Either this or relative_likelihood can be set. Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#fixed_likelihood GoogleDataLossPreventionContentPolicy#fixed_likelihood}
	FixedLikelihood *string `field:"optional" json:"fixedLikelihood" yaml:"fixedLikelihood"`
	// Increase or decrease the likelihood by the specified number of levels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#relative_likelihood GoogleDataLossPreventionContentPolicy#relative_likelihood}
	RelativeLikelihood *float64 `field:"optional" json:"relativeLikelihood" yaml:"relativeLikelihood"`
}

