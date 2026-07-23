// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiapagentregistrymcpserveriambinding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleIapAgentRegistryMcpServerIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#mcp_server_id GoogleIapAgentRegistryMcpServerIamBinding#mcp_server_id}.
	McpServerId *string `field:"required" json:"mcpServerId" yaml:"mcpServerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#members GoogleIapAgentRegistryMcpServerIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#role GoogleIapAgentRegistryMcpServerIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#condition GoogleIapAgentRegistryMcpServerIamBinding#condition}
	Condition *GoogleIapAgentRegistryMcpServerIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#id GoogleIapAgentRegistryMcpServerIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#location GoogleIapAgentRegistryMcpServerIamBinding#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#project GoogleIapAgentRegistryMcpServerIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

