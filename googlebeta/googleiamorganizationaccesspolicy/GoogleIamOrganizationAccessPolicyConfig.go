// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamorganizationaccesspolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleIamOrganizationAccessPolicyConfig struct {
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
	// The ID to use for the access policy, which will become the final component of the access policy's resource name.
	//
	// This value must start with a lowercase letter followed by up to 62
	// lowercase letters, numbers, hyphens, or dots. Pattern,
	// /a-z{2,62}/.
	// This value must be unique among all access policies with the same parent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_organization_access_policy#access_policy_id GoogleIamOrganizationAccessPolicy#access_policy_id}
	AccessPolicyId *string `field:"required" json:"accessPolicyId" yaml:"accessPolicyId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_organization_access_policy#location GoogleIamOrganizationAccessPolicy#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_organization_access_policy#organization GoogleIamOrganizationAccessPolicy#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size limitations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_organization_access_policy#annotations GoogleIamOrganizationAccessPolicy#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_organization_access_policy#deletion_policy GoogleIamOrganizationAccessPolicy#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_organization_access_policy#details GoogleIamOrganizationAccessPolicy#details}
	Details *GoogleIamOrganizationAccessPolicyDetails `field:"optional" json:"details" yaml:"details"`
	// The description of the access policy. Must be less than or equal to 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_organization_access_policy#display_name GoogleIamOrganizationAccessPolicy#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_organization_access_policy#id GoogleIamOrganizationAccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_organization_access_policy#timeouts GoogleIamOrganizationAccessPolicy#timeouts}
	Timeouts *GoogleIamOrganizationAccessPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

