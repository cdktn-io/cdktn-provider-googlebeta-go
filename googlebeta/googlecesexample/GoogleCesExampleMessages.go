// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesexample


type GoogleCesExampleMessages struct {
	// chunks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#chunks GoogleCesExample#chunks}
	Chunks interface{} `field:"optional" json:"chunks" yaml:"chunks"`
	// The role within the conversation, e.g., user, agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#role GoogleCesExample#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

