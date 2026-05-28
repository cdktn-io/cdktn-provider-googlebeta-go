// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleapigeeapideployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleApigeeApiDeploymentConfig struct {
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
	// The Apigee Environment associated with the Apigee API deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_apigee_api_deployment#environment GoogleApigeeApiDeployment#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The Apigee Organization associated with the Apigee API deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_apigee_api_deployment#org_id GoogleApigeeApiDeployment#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// The Apigee API associated with the Apigee API deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_apigee_api_deployment#proxy_id GoogleApigeeApiDeployment#proxy_id}
	ProxyId *string `field:"required" json:"proxyId" yaml:"proxyId"`
	// The revision of the API proxy to be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_apigee_api_deployment#revision GoogleApigeeApiDeployment#revision}
	Revision *string `field:"required" json:"revision" yaml:"revision"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_apigee_api_deployment#deletion_policy GoogleApigeeApiDeployment#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_apigee_api_deployment#timeouts GoogleApigeeApiDeployment#timeouts}
	Timeouts *GoogleApigeeApiDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

