// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlediscoveryengineschema

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDiscoveryEngineSchemaConfig struct {
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
	// The unique id of the data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_discovery_engine_schema#data_store_id GoogleDiscoveryEngineSchema#data_store_id}
	DataStoreId *string `field:"required" json:"dataStoreId" yaml:"dataStoreId"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_discovery_engine_schema#location GoogleDiscoveryEngineSchema#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The unique id of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_discovery_engine_schema#schema_id GoogleDiscoveryEngineSchema#schema_id}
	SchemaId *string `field:"required" json:"schemaId" yaml:"schemaId"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_discovery_engine_schema#deletion_policy GoogleDiscoveryEngineSchema#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_discovery_engine_schema#id GoogleDiscoveryEngineSchema#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The JSON representation of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_discovery_engine_schema#json_schema GoogleDiscoveryEngineSchema#json_schema}
	JsonSchema *string `field:"optional" json:"jsonSchema" yaml:"jsonSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_discovery_engine_schema#project GoogleDiscoveryEngineSchema#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_discovery_engine_schema#timeouts GoogleDiscoveryEngineSchema#timeouts}
	Timeouts *GoogleDiscoveryEngineSchemaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

