// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralgoogleserviceaccountkey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralGoogleServiceAccountKeyConfig struct {
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
	// Whether to fetch the public key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_key#fetch_key EphemeralGoogleServiceAccountKey#fetch_key}
	FetchKey interface{} `field:"optional" json:"fetchKey" yaml:"fetchKey"`
	// The algorithm used to generate the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_key#key_algorithm EphemeralGoogleServiceAccountKey#key_algorithm}
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// The name of the service account key.
	//
	// This must have format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{KEYID}`, where `{ACCOUNT}` is the email address or unique id of the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_key#name EphemeralGoogleServiceAccountKey#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The private key, base64 encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_key#private_key EphemeralGoogleServiceAccountKey#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The type of the private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_key#private_key_type EphemeralGoogleServiceAccountKey#private_key_type}
	PrivateKeyType *string `field:"optional" json:"privateKeyType" yaml:"privateKeyType"`
	// The public key, base64 encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_key#public_key_data EphemeralGoogleServiceAccountKey#public_key_data}
	PublicKeyData *string `field:"optional" json:"publicKeyData" yaml:"publicKeyData"`
	// The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_key#public_key_type EphemeralGoogleServiceAccountKey#public_key_type}
	PublicKeyType *string `field:"optional" json:"publicKeyType" yaml:"publicKeyType"`
	// The ID of the parent service account of the key.
	//
	// This can be a string in the format {ACCOUNT} or projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}, where {ACCOUNT} is the email address or unique id of the service account. If the {ACCOUNT} syntax is used, the project will be inferred from the provider's configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/ephemeral-resources/google_service_account_key#service_account_id EphemeralGoogleServiceAccountKey#service_account_id}
	ServiceAccountId *string `field:"optional" json:"serviceAccountId" yaml:"serviceAccountId"`
}

