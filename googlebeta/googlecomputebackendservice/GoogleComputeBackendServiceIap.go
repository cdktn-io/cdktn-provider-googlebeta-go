// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputebackendservice


type GoogleComputeBackendServiceIap struct {
	// Whether the serving infrastructure will authenticate and authorize all incoming requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_backend_service#enabled GoogleComputeBackendService#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// OAuth2 Client ID for IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_backend_service#oauth2_client_id GoogleComputeBackendService#oauth2_client_id}
	Oauth2ClientId *string `field:"optional" json:"oauth2ClientId" yaml:"oauth2ClientId"`
	// OAuth2 Client ID for IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_backend_service#oauth2_client_id_wo GoogleComputeBackendService#oauth2_client_id_wo}
	Oauth2ClientIdWo *string `field:"optional" json:"oauth2ClientIdWo" yaml:"oauth2ClientIdWo"`
	// Triggers update of 'oauth2_client_id_wo' write-only.
	//
	// Increment this value when an update to 'oauth2_client_id_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_backend_service#oauth2_client_id_wo_version GoogleComputeBackendService#oauth2_client_id_wo_version}
	Oauth2ClientIdWoVersion *string `field:"optional" json:"oauth2ClientIdWoVersion" yaml:"oauth2ClientIdWoVersion"`
	// OAuth2 Client Secret for IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_backend_service#oauth2_client_secret GoogleComputeBackendService#oauth2_client_secret}
	Oauth2ClientSecret *string `field:"optional" json:"oauth2ClientSecret" yaml:"oauth2ClientSecret"`
	// OAuth2 Client Secret for IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_backend_service#oauth2_client_secret_wo GoogleComputeBackendService#oauth2_client_secret_wo}
	Oauth2ClientSecretWo *string `field:"optional" json:"oauth2ClientSecretWo" yaml:"oauth2ClientSecretWo"`
	// Triggers update of 'oauth2_client_secret_wo' write-only.
	//
	// Increment this value when an update to 'oauth2_client_secret_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_backend_service#oauth2_client_secret_wo_version GoogleComputeBackendService#oauth2_client_secret_wo_version}
	Oauth2ClientSecretWoVersion *string `field:"optional" json:"oauth2ClientSecretWoVersion" yaml:"oauth2ClientSecretWoVersion"`
}

