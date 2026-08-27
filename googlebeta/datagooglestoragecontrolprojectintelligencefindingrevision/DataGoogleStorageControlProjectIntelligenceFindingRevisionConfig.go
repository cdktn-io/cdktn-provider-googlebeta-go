// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglestoragecontrolprojectintelligencefindingrevision

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleStorageControlProjectIntelligenceFindingRevisionConfig struct {
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
	// The ID of the intelligence finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/data-sources/google_storage_control_project_intelligence_finding_revision#finding_id DataGoogleStorageControlProjectIntelligenceFindingRevision#finding_id}
	FindingId *string `field:"required" json:"findingId" yaml:"findingId"`
	// The ID of the finding revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/data-sources/google_storage_control_project_intelligence_finding_revision#revision_id DataGoogleStorageControlProjectIntelligenceFindingRevision#revision_id}
	RevisionId *string `field:"required" json:"revisionId" yaml:"revisionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/data-sources/google_storage_control_project_intelligence_finding_revision#id DataGoogleStorageControlProjectIntelligenceFindingRevision#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the intelligence finding. Currently default value is global and users cannot use for input for now.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/data-sources/google_storage_control_project_intelligence_finding_revision#location DataGoogleStorageControlProjectIntelligenceFindingRevision#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/data-sources/google_storage_control_project_intelligence_finding_revision#project DataGoogleStorageControlProjectIntelligenceFindingRevision#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

