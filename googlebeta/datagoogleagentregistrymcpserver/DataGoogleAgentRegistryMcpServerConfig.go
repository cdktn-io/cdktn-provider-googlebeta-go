// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagoogleagentregistrymcpserver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleAgentRegistryMcpServerConfig struct {
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
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/data-sources/google_agent_registry_mcp_server#location DataGoogleAgentRegistryMcpServer#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A filter string that identifies a unique MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/data-sources/google_agent_registry_mcp_server#filter DataGoogleAgentRegistryMcpServer#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/data-sources/google_agent_registry_mcp_server#id DataGoogleAgentRegistryMcpServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The unique identifier for the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/data-sources/google_agent_registry_mcp_server#mcp_server_id DataGoogleAgentRegistryMcpServer#mcp_server_id}
	McpServerId *string `field:"optional" json:"mcpServerId" yaml:"mcpServerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/data-sources/google_agent_registry_mcp_server#project DataGoogleAgentRegistryMcpServer#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

