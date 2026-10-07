// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesdeployment


type GoogleCesDeploymentChannelProfileWhatsappConfig struct {
	// Required. The Meta phone number ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#phone_number_id GoogleCesDeployment#phone_number_id}
	PhoneNumberId *string `field:"required" json:"phoneNumberId" yaml:"phoneNumberId"`
	// Required. The WhatsApp Business Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#waba_id GoogleCesDeployment#waba_id}
	WabaId *string `field:"required" json:"wabaId" yaml:"wabaId"`
	// Optional. The phone number in E.164 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_deployment#phone_number GoogleCesDeployment#phone_number}
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

