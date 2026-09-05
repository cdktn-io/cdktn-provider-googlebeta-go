// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivecatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleBiglakeHiveCatalogConfig struct {
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
	// Cloud Storage location path where the catalog data will be stored. Format: gs://bucket/path/to/catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_catalog#location_uri GoogleBiglakeHiveCatalog#location_uri}
	LocationUri *string `field:"required" json:"locationUri" yaml:"locationUri"`
	// Name of the Hive Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_catalog#name GoogleBiglakeHiveCatalog#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The primary location for mirroring the remote catalog metadata.
	//
	// It must be
	// a BigLake-supported location, and it should be proximate to the remote
	// catalog's location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_catalog#primary_location GoogleBiglakeHiveCatalog#primary_location}
	PrimaryLocation *string `field:"required" json:"primaryLocation" yaml:"primaryLocation"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_catalog#deletion_policy GoogleBiglakeHiveCatalog#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Description of the Hive catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_catalog#description GoogleBiglakeHiveCatalog#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_catalog#id GoogleBiglakeHiveCatalog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_catalog#project GoogleBiglakeHiveCatalog#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_catalog#timeouts GoogleBiglakeHiveCatalog#timeouts}
	Timeouts *GoogleBiglakeHiveCatalogTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

