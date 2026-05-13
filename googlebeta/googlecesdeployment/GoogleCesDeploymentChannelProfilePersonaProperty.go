// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesdeployment


type GoogleCesDeploymentChannelProfilePersonaProperty struct {
	// The persona of the channel. Possible values: UNKNOWN CONCISE CHATTY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_deployment#persona GoogleCesDeployment#persona}
	Persona *string `field:"optional" json:"persona" yaml:"persona"`
}

