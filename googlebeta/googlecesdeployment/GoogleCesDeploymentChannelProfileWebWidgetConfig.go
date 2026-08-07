// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesdeployment


type GoogleCesDeploymentChannelProfileWebWidgetConfig struct {
	// The modality of the web widget. Possible values: MODALITY_UNSPECIFIED CHAT_AND_VOICE VOICE_ONLY CHAT_ONLY CHAT_VOICE_AND_VIDEO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_ces_deployment#modality GoogleCesDeployment#modality}
	Modality *string `field:"optional" json:"modality" yaml:"modality"`
	// security_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_ces_deployment#security_settings GoogleCesDeployment#security_settings}
	SecuritySettings *GoogleCesDeploymentChannelProfileWebWidgetConfigSecuritySettings `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// The theme of the web widget. Possible values: THEME_UNSPECIFIED LIGHT DARK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_ces_deployment#theme GoogleCesDeployment#theme}
	Theme *string `field:"optional" json:"theme" yaml:"theme"`
	// The title of the web widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_ces_deployment#web_widget_title GoogleCesDeployment#web_widget_title}
	WebWidgetTitle *string `field:"optional" json:"webWidgetTitle" yaml:"webWidgetTitle"`
}

