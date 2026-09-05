// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralgoogleserviceaccountjwt

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralGoogleServiceAccountJwtConfig struct {
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
	// A JSON-encoded JWT claims set that will be included in the signed JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/ephemeral-resources/google_service_account_jwt#payload EphemeralGoogleServiceAccountJwt#payload}
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// The email of the service account that will sign the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/ephemeral-resources/google_service_account_jwt#target_service_account EphemeralGoogleServiceAccountJwt#target_service_account}
	TargetServiceAccount *string `field:"required" json:"targetServiceAccount" yaml:"targetServiceAccount"`
	// Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/ephemeral-resources/google_service_account_jwt#delegates EphemeralGoogleServiceAccountJwt#delegates}
	Delegates *[]*string `field:"optional" json:"delegates" yaml:"delegates"`
	// Number of seconds until the JWT expires.
	//
	// If set and non-zero an `exp` claim will be added to the payload derived from the current timestamp plus expires_in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/ephemeral-resources/google_service_account_jwt#expires_in EphemeralGoogleServiceAccountJwt#expires_in}
	ExpiresIn *float64 `field:"optional" json:"expiresIn" yaml:"expiresIn"`
}

