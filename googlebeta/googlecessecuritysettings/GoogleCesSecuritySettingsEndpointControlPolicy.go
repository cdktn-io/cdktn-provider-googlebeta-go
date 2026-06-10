// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecessecuritysettings


type GoogleCesSecuritySettingsEndpointControlPolicy struct {
	// Optional. The allowed HTTP(s) origins that tools in the App are able to directly call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_ces_security_settings#allowed_origins GoogleCesSecuritySettings#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Optional. The scope in which this policy's allowedOrigins list is enforced. Possible values: ["ENFORCEMENT_SCOPE_UNSPECIFIED", "VPCSC_ONLY", "ALWAYS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_ces_security_settings#enforcement_scope GoogleCesSecuritySettings#enforcement_scope}
	EnforcementScope *string `field:"optional" json:"enforcementScope" yaml:"enforcementScope"`
}

