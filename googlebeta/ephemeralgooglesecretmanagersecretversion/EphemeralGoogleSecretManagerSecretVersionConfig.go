// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralgooglesecretmanagersecretversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralGoogleSecretManagerSecretVersionConfig struct {
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
	// The secret to get the secret version for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/ephemeral-resources/google_secret_manager_secret_version#secret EphemeralGoogleSecretManagerSecretVersion#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// If true, the secret data returned will not get base64 decoded. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/ephemeral-resources/google_secret_manager_secret_version#is_secret_data_base64 EphemeralGoogleSecretManagerSecretVersion#is_secret_data_base64}
	IsSecretDataBase64 interface{} `field:"optional" json:"isSecretDataBase64" yaml:"isSecretDataBase64"`
	// The project to get the secret version for. If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/ephemeral-resources/google_secret_manager_secret_version#project EphemeralGoogleSecretManagerSecretVersion#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The version of the secret to get. If it is not provided, the latest version is retrieved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/ephemeral-resources/google_secret_manager_secret_version#version EphemeralGoogleSecretManagerSecretVersion#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

