// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesdeployment


type GoogleCesDeploymentInstagramCredentials struct {
	// The Meta auth code provided by the embedded signup flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#auth_code GoogleCesDeployment#auth_code}
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// The Meta auth code provided by the embedded signup flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#auth_code_wo GoogleCesDeployment#auth_code_wo}
	AuthCodeWo *string `field:"optional" json:"authCodeWo" yaml:"authCodeWo"`
	// Triggers update of 'auth_code_wo' write-only.
	//
	// Increment this value when an update to 'auth_code_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#auth_code_wo_version GoogleCesDeployment#auth_code_wo_version}
	AuthCodeWoVersion *string `field:"optional" json:"authCodeWoVersion" yaml:"authCodeWoVersion"`
	// The Conversation Profile ID to use for the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#conversation_profile_id GoogleCesDeployment#conversation_profile_id}
	ConversationProfileId *string `field:"optional" json:"conversationProfileId" yaml:"conversationProfileId"`
}

