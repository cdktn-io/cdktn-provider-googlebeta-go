// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesevaluation


type GoogleCesEvaluationGoldenTurnsStepsAgentTransfer struct {
	// The resource name of the target agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#target_agent GoogleCesEvaluation#target_agent}
	TargetAgent *string `field:"required" json:"targetAgent" yaml:"targetAgent"`
}

