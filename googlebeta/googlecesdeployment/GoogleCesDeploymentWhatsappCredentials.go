// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesdeployment


type GoogleCesDeploymentWhatsappCredentials struct {
	// The Business Account ID to use for the phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#business_account_id GoogleCesDeployment#business_account_id}
	BusinessAccountId *string `field:"required" json:"businessAccountId" yaml:"businessAccountId"`
	// The phone number to register with WhatsApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#phone_number GoogleCesDeployment#phone_number}
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// The WhatsApp Business Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#waba_id GoogleCesDeployment#waba_id}
	WabaId *string `field:"required" json:"wabaId" yaml:"wabaId"`
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
	// The 6-digit PIN created by the user for two-step verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#pin GoogleCesDeployment#pin}
	Pin *string `field:"optional" json:"pin" yaml:"pin"`
	// The 6-digit PIN created by the user for two-step verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#pin_wo GoogleCesDeployment#pin_wo}
	PinWo *string `field:"optional" json:"pinWo" yaml:"pinWo"`
	// Triggers update of 'pin_wo' write-only.
	//
	// Increment this value when an update to 'pin_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#pin_wo_version GoogleCesDeployment#pin_wo_version}
	PinWoVersion *string `field:"optional" json:"pinWoVersion" yaml:"pinWoVersion"`
}

