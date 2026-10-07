// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent


type GoogleCesAgentTransferRulesDeterministicTransfer struct {
	// expression_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#expression_condition GoogleCesAgent#expression_condition}
	ExpressionCondition *GoogleCesAgentTransferRulesDeterministicTransferExpressionCondition `field:"optional" json:"expressionCondition" yaml:"expressionCondition"`
	// python_code_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#python_code_condition GoogleCesAgent#python_code_condition}
	PythonCodeCondition *GoogleCesAgentTransferRulesDeterministicTransferPythonCodeCondition `field:"optional" json:"pythonCodeCondition" yaml:"pythonCodeCondition"`
}

