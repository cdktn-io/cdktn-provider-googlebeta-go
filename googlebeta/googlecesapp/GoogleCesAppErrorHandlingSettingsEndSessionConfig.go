// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesapp


type GoogleCesAppErrorHandlingSettingsEndSessionConfig struct {
	// Whether to escalate the session in EndSession. If session is escalated, metadata in EndSession will contain session_escalated = true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_app#escalate_session GoogleCesApp#escalate_session}
	EscalateSession interface{} `field:"optional" json:"escalateSession" yaml:"escalateSession"`
}

