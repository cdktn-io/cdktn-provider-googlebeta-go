// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesguardrail


type GoogleCesGuardrailModelSafety struct {
	// safety_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_guardrail#safety_settings GoogleCesGuardrail#safety_settings}
	SafetySettings interface{} `field:"required" json:"safetySettings" yaml:"safetySettings"`
}

