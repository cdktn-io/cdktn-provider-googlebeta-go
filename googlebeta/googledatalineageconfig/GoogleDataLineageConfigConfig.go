// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalineageconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataLineageConfigConfig struct {
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
	// ingestion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_data_lineage_config#ingestion GoogleDataLineageConfig#ingestion}
	Ingestion *GoogleDataLineageConfigIngestion `field:"required" json:"ingestion" yaml:"ingestion"`
	// The region of the data lineage configuration for integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_data_lineage_config#location GoogleDataLineageConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Parent scope for the config. Format: projects/{project-id|project-number} or folders/{folder-number} or organizations/{organization-number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_data_lineage_config#parent GoogleDataLineageConfig#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_data_lineage_config#deletion_policy GoogleDataLineageConfig#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_data_lineage_config#id GoogleDataLineageConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_data_lineage_config#timeouts GoogleDataLineageConfig#timeouts}
	Timeouts *GoogleDataLineageConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

