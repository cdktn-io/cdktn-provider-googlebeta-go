// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent


type GoogleCesAgentTransferRulesDisablePlannerTransferExpressionCondition struct {
	// The string representation of cloud.api.Expression condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#expression GoogleCesAgent#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

