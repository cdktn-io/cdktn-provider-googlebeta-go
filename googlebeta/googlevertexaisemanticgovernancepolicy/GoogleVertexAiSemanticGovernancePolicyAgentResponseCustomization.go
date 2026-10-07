// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaisemanticgovernancepolicy


type GoogleVertexAiSemanticGovernancePolicyAgentResponseCustomization struct {
	// Custom message shown to the end user when the policy check results in a denial.
	//
	// Use this
	// to explain the rationale to the user. Max 1000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_semantic_governance_policy#denial_message GoogleVertexAiSemanticGovernancePolicy#denial_message}
	DenialMessage *string `field:"optional" json:"denialMessage" yaml:"denialMessage"`
}

