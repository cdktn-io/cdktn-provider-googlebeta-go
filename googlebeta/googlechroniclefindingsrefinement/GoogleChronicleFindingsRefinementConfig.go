// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclefindingsrefinement

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleFindingsRefinementConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_findings_refinement#instance GoogleChronicleFindingsRefinement#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_findings_refinement#location GoogleChronicleFindingsRefinement#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Display name of the findings refinement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_findings_refinement#display_name GoogleChronicleFindingsRefinement#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_findings_refinement#id GoogleChronicleFindingsRefinement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// outcome_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_findings_refinement#outcome_filters GoogleChronicleFindingsRefinement#outcome_filters}
	OutcomeFilters interface{} `field:"optional" json:"outcomeFilters" yaml:"outcomeFilters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_findings_refinement#project GoogleChronicleFindingsRefinement#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The query for the findings refinement.
	//
	// Works in conjunction with the type
	// field to determine the findings refinement behavior. The syntax of this
	// query is the same as a UDM search string. See the following for more
	// information:
	// https://cloud.google.com/chronicle/docs/investigation/udm-search
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_findings_refinement#query GoogleChronicleFindingsRefinement#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_findings_refinement#timeouts GoogleChronicleFindingsRefinement#timeouts}
	Timeouts *GoogleChronicleFindingsRefinementTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// DETECTION_EXCLUSION is the only supported type of findings refinement. Possible values: DETECTION_EXCLUSION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_findings_refinement#type GoogleChronicleFindingsRefinement#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

