// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleapigeedatacollector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleApigeeDataCollectorConfig struct {
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
	// The ID for the data collector. Must begin with 'dc_'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_apigee_data_collector#data_collector_id GoogleApigeeDataCollector#data_collector_id}
	DataCollectorId *string `field:"required" json:"dataCollectorId" yaml:"dataCollectorId"`
	// The Apigee Organization associated with the Apigee data collector, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_apigee_data_collector#org_id GoogleApigeeDataCollector#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// The type of data this data collector will collect. Possible values: ["BOOLEAN", "DATETIME", "FLOAT", "INTEGER", "STRING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_apigee_data_collector#type GoogleApigeeDataCollector#type}
	Type *string `field:"required" json:"type" yaml:"type"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_apigee_data_collector#deletion_policy GoogleApigeeDataCollector#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// A description of the data collector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_apigee_data_collector#description GoogleApigeeDataCollector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_apigee_data_collector#id GoogleApigeeDataCollector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_apigee_data_collector#timeouts GoogleApigeeDataCollector#timeouts}
	Timeouts *GoogleApigeeDataCollectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

