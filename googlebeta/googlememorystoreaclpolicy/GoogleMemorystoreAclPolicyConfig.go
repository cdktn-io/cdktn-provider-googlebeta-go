// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlememorystoreaclpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleMemorystoreAclPolicyConfig struct {
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
	// The logical name of the ACL policy in the customer project with the following restrictions:.
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the customer project / location
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_memorystore_acl_policy#acl_policy_id GoogleMemorystoreAclPolicy#acl_policy_id}
	AclPolicyId *string `field:"required" json:"aclPolicyId" yaml:"aclPolicyId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_memorystore_acl_policy#location GoogleMemorystoreAclPolicy#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_memorystore_acl_policy#rules GoogleMemorystoreAclPolicy#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_memorystore_acl_policy#deletion_policy GoogleMemorystoreAclPolicy#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_memorystore_acl_policy#id GoogleMemorystoreAclPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_memorystore_acl_policy#project GoogleMemorystoreAclPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_memorystore_acl_policy#timeouts GoogleMemorystoreAclPolicy#timeouts}
	Timeouts *GoogleMemorystoreAclPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

