// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralgoogleserviceaccountaccesstoken

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralGoogleServiceAccountAccessTokenConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// The scopes the new credential should have (e.g. `['cloud-platform']`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_access_token#scopes EphemeralGoogleServiceAccountAccessToken#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The service account to impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_access_token#target_service_account EphemeralGoogleServiceAccountAccessToken#target_service_account}
	TargetServiceAccount *string `field:"required" json:"targetServiceAccount" yaml:"targetServiceAccount"`
	// Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `['projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com']`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_access_token#delegates EphemeralGoogleServiceAccountAccessToken#delegates}
	Delegates *[]*string `field:"optional" json:"delegates" yaml:"delegates"`
	// Lifetime of the impersonated token (defaults to its max: `3600s`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_access_token#lifetime EphemeralGoogleServiceAccountAccessToken#lifetime}
	Lifetime *string `field:"optional" json:"lifetime" yaml:"lifetime"`
}

