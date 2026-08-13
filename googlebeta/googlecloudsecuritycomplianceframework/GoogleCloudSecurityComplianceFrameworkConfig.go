// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudsecuritycomplianceframework

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCloudSecurityComplianceFrameworkConfig struct {
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
	// ID of the framework.
	//
	// This is not the full name of the framework.
	// This is the last part of the full name of the framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#framework_id GoogleCloudSecurityComplianceFramework#framework_id}
	FrameworkId *string `field:"required" json:"frameworkId" yaml:"frameworkId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#location GoogleCloudSecurityComplianceFramework#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// cloud_control_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#cloud_control_details GoogleCloudSecurityComplianceFramework#cloud_control_details}
	CloudControlDetails interface{} `field:"optional" json:"cloudControlDetails" yaml:"cloudControlDetails"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#deletion_policy GoogleCloudSecurityComplianceFramework#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The description of the framework. The maximum length is 2000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#description GoogleCloudSecurityComplianceFramework#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name of the framework. The maximum length is 200 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#display_name GoogleCloudSecurityComplianceFramework#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#id GoogleCloudSecurityComplianceFramework#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#organization GoogleCloudSecurityComplianceFramework#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The parent resource in which to create the resource. Must be in one of the following formats: * 'projects/{{project}}' * 'organizations/{{organization}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#parent GoogleCloudSecurityComplianceFramework#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_cloud_security_compliance_framework#timeouts GoogleCloudSecurityComplianceFramework#timeouts}
	Timeouts *GoogleCloudSecurityComplianceFrameworkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

