// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclefindingsrefinement


type GoogleChronicleFindingsRefinementOutcomeFilters struct {
	// The operator to be applied to the outcome variable. Possible values: EQUAL CONTAINS MATCHES_REGEX MATCHES_CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_chronicle_findings_refinement#outcome_filter_operator GoogleChronicleFindingsRefinement#outcome_filter_operator}
	OutcomeFilterOperator *string `field:"required" json:"outcomeFilterOperator" yaml:"outcomeFilterOperator"`
	// The value of the outcome variable to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_chronicle_findings_refinement#outcome_value GoogleChronicleFindingsRefinement#outcome_value}
	OutcomeValue *string `field:"required" json:"outcomeValue" yaml:"outcomeValue"`
	// The outcome variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_chronicle_findings_refinement#outcome_variable GoogleChronicleFindingsRefinement#outcome_variable}
	OutcomeVariable *string `field:"required" json:"outcomeVariable" yaml:"outcomeVariable"`
}

