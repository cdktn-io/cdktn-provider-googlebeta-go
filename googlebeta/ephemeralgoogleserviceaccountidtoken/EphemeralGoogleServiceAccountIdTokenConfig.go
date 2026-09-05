// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralgoogleserviceaccountidtoken

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralGoogleServiceAccountIdTokenConfig struct {
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
	// The audience claim for the `id_token`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/ephemeral-resources/google_service_account_id_token#target_audience EphemeralGoogleServiceAccountIdToken#target_audience}
	TargetAudience *string `field:"required" json:"targetAudience" yaml:"targetAudience"`
	// Delegate chain of approvals needed to perform full impersonation.
	//
	// Specify the fully qualified service account name.  Used only when using impersonation mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/ephemeral-resources/google_service_account_id_token#delegates EphemeralGoogleServiceAccountIdToken#delegates}
	Delegates *[]*string `field:"optional" json:"delegates" yaml:"delegates"`
	// Include the verified email in the claim. Used only when using impersonation mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/ephemeral-resources/google_service_account_id_token#include_email EphemeralGoogleServiceAccountIdToken#include_email}
	IncludeEmail interface{} `field:"optional" json:"includeEmail" yaml:"includeEmail"`
	// The email of the service account being impersonated.  Used only when using impersonation mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/ephemeral-resources/google_service_account_id_token#target_service_account EphemeralGoogleServiceAccountIdToken#target_service_account}
	TargetServiceAccount *string `field:"optional" json:"targetServiceAccount" yaml:"targetServiceAccount"`
}

