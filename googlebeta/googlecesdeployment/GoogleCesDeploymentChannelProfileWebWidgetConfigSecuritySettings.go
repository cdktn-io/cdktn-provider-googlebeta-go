// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesdeployment


type GoogleCesDeploymentChannelProfileWebWidgetConfigSecuritySettings struct {
	// The origins that are allowed to host the web widget.
	//
	// An origin is defined by RFC 6454. If empty, all origins are allowed. A maximum of 100 origins is allowed. Example: "https://example.com"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_deployment#allowed_origins GoogleCesDeployment#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Indicates whether origin check for the web widget is enabled.
	//
	// If true, the web widget will check the origin of the website that loads the web widget and only allow it to be loaded in the same origin or any of the allowed origins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_deployment#enable_origin_check GoogleCesDeployment#enable_origin_check}
	EnableOriginCheck interface{} `field:"optional" json:"enableOriginCheck" yaml:"enableOriginCheck"`
	// Indicates whether public access to the web widget is enabled.
	//
	// If true, the web widget will be publicly accessible. If false, the web widget must be integrated with your own authentication and authorization system to return valid credentials for accessing the CES agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_deployment#enable_public_access GoogleCesDeployment#enable_public_access}
	EnablePublicAccess interface{} `field:"optional" json:"enablePublicAccess" yaml:"enablePublicAccess"`
	// Indicates whether reCAPTCHA verification for the web widget is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_deployment#enable_recaptcha GoogleCesDeployment#enable_recaptcha}
	EnableRecaptcha interface{} `field:"optional" json:"enableRecaptcha" yaml:"enableRecaptcha"`
}

