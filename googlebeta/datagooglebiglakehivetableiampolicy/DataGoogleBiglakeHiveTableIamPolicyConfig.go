// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglebiglakehivetableiampolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleBiglakeHiveTableIamPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/data-sources/google_biglake_hive_table_iam_policy#catalog DataGoogleBiglakeHiveTableIamPolicy#catalog}.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/data-sources/google_biglake_hive_table_iam_policy#database DataGoogleBiglakeHiveTableIamPolicy#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/data-sources/google_biglake_hive_table_iam_policy#name DataGoogleBiglakeHiveTableIamPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/data-sources/google_biglake_hive_table_iam_policy#id DataGoogleBiglakeHiveTableIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/data-sources/google_biglake_hive_table_iam_policy#project DataGoogleBiglakeHiveTableIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

