// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglestoragecontrolfolderintelligencefindingssummary

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleStorageControlFolderIntelligenceFindingsSummaryConfig struct {
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
	// The ID of the Google Cloud Folder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/data-sources/google_storage_control_folder_intelligence_findings_summary#folder DataGoogleStorageControlFolderIntelligenceFindingsSummary#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// The filter expression. Supports filtering by FindingType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/data-sources/google_storage_control_folder_intelligence_findings_summary#filter DataGoogleStorageControlFolderIntelligenceFindingsSummary#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/data-sources/google_storage_control_folder_intelligence_findings_summary#id DataGoogleStorageControlFolderIntelligenceFindingsSummary#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the intelligence findings summary.
	//
	// Currently default value is global and users cannot use for input for now.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/data-sources/google_storage_control_folder_intelligence_findings_summary#location DataGoogleStorageControlFolderIntelligenceFindingsSummary#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Determines the granularity of the findings when the parent is an organization or folder.
	//
	// Possible values are PARENT and PROJECT. Default value is PARENT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/data-sources/google_storage_control_folder_intelligence_findings_summary#resource_scope DataGoogleStorageControlFolderIntelligenceFindingsSummary#resource_scope}
	ResourceScope *string `field:"optional" json:"resourceScope" yaml:"resourceScope"`
}

