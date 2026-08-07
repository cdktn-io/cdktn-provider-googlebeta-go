// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleapigeedatastore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleApigeeDatastoreConfig struct {
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
	// datastore_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_apigee_datastore#datastore_config GoogleApigeeDatastore#datastore_config}
	DatastoreConfig *GoogleApigeeDatastoreDatastoreConfig `field:"required" json:"datastoreConfig" yaml:"datastoreConfig"`
	// The display name for the datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_apigee_datastore#display_name GoogleApigeeDatastore#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Apigee Organization associated with the Apigee datastore, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_apigee_datastore#org_id GoogleApigeeDatastore#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// The type of target for the datastore. Must be 'gcs' for Google Cloud Storage or 'bigquery' for BigQuery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_apigee_datastore#target_type GoogleApigeeDatastore#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_apigee_datastore#deletion_policy GoogleApigeeDatastore#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_apigee_datastore#id GoogleApigeeDatastore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_apigee_datastore#timeouts GoogleApigeeDatastore#timeouts}
	Timeouts *GoogleApigeeDatastoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

