// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontactcenterinsightsencryptionspec

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleContactCenterInsightsEncryptionSpecConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of customer-managed encryption key that is used to secure a resource and its sub-resources.
	//
	// If empty, the resource is secured by the default Google encryption key.
	// Only the key in the same location as this resource is allowed to be used for encryption.
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{key}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_contact_center_insights_encryption_spec#kms_key GoogleContactCenterInsightsEncryptionSpec#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
	// The location in which the encryptionSpec is to be initialized.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_contact_center_insights_encryption_spec#location GoogleContactCenterInsightsEncryptionSpec#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_contact_center_insights_encryption_spec#id GoogleContactCenterInsightsEncryptionSpec#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_contact_center_insights_encryption_spec#project GoogleContactCenterInsightsEncryptionSpec#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_contact_center_insights_encryption_spec#timeouts GoogleContactCenterInsightsEncryptionSpec#timeouts}
	Timeouts *GoogleContactCenterInsightsEncryptionSpecTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

