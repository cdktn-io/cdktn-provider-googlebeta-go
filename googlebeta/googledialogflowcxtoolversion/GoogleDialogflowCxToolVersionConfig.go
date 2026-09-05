// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowcxtoolversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDialogflowCxToolVersionConfig struct {
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
	// The display name of the tool version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_cx_tool_version#display_name GoogleDialogflowCxToolVersion#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The tool to create a Version for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/tools/<Tool ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_cx_tool_version#parent GoogleDialogflowCxToolVersion#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_cx_tool_version#tool GoogleDialogflowCxToolVersion#tool}
	Tool *GoogleDialogflowCxToolVersionTool `field:"required" json:"tool" yaml:"tool"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_cx_tool_version#deletion_policy GoogleDialogflowCxToolVersion#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_cx_tool_version#id GoogleDialogflowCxToolVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_cx_tool_version#timeouts GoogleDialogflowCxToolVersion#timeouts}
	Timeouts *GoogleDialogflowCxToolVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

