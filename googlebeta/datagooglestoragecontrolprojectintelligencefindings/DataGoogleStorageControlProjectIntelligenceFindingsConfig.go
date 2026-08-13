// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglestoragecontrolprojectintelligencefindings

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleStorageControlProjectIntelligenceFindingsConfig struct {
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
	// The filter expression to apply. Supports filtering by type and associated_resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/data-sources/google_storage_control_project_intelligence_findings#filter DataGoogleStorageControlProjectIntelligenceFindings#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/data-sources/google_storage_control_project_intelligence_findings#id DataGoogleStorageControlProjectIntelligenceFindings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the intelligence findings. Currently default value is global and users cannot use for input for now.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/data-sources/google_storage_control_project_intelligence_findings#location DataGoogleStorageControlProjectIntelligenceFindings#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The maximum number of IntelligenceFinding resources to return.
	//
	// The maximum value is 100; values greater than 100 become 100. The default value is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/data-sources/google_storage_control_project_intelligence_findings#page_size DataGoogleStorageControlProjectIntelligenceFindings#page_size}
	PageSize *float64 `field:"optional" json:"pageSize" yaml:"pageSize"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/data-sources/google_storage_control_project_intelligence_findings#project DataGoogleStorageControlProjectIntelligenceFindings#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

