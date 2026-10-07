// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment struct {
	// Set the likelihood of a finding to a fixed value. Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#fixed_likelihood GoogleDataLossPreventionContentPolicy#fixed_likelihood}
	FixedLikelihood *string `field:"required" json:"fixedLikelihood" yaml:"fixedLikelihood"`
}

