// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent


type GoogleCesAgentTransferRulesDisablePlannerTransfer struct {
	// expression_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#expression_condition GoogleCesAgent#expression_condition}
	ExpressionCondition *GoogleCesAgentTransferRulesDisablePlannerTransferExpressionCondition `field:"required" json:"expressionCondition" yaml:"expressionCondition"`
}

