// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesToolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#app GoogleCesTool#app}
	App *string `field:"required" json:"app" yaml:"app"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#location GoogleCesTool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID to use for the tool, which will become the final component of the tool's resource name.
	//
	// If not provided, a unique ID will be
	// automatically assigned for the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#tool_id GoogleCesTool#tool_id}
	ToolId *string `field:"required" json:"toolId" yaml:"toolId"`
	// agent_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#agent_tool GoogleCesTool#agent_tool}
	AgentTool *GoogleCesToolAgentTool `field:"optional" json:"agentTool" yaml:"agentTool"`
	// client_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#client_function GoogleCesTool#client_function}
	ClientFunction *GoogleCesToolClientFunction `field:"optional" json:"clientFunction" yaml:"clientFunction"`
	// data_store_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#data_store_tool GoogleCesTool#data_store_tool}
	DataStoreTool *GoogleCesToolDataStoreTool `field:"optional" json:"dataStoreTool" yaml:"dataStoreTool"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#deletion_policy GoogleCesTool#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Possible values: SYNCHRONOUS ASYNCHRONOUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#execution_type GoogleCesTool#execution_type}
	ExecutionType *string `field:"optional" json:"executionType" yaml:"executionType"`
	// file_search_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#file_search_tool GoogleCesTool#file_search_tool}
	FileSearchTool *GoogleCesToolFileSearchTool `field:"optional" json:"fileSearchTool" yaml:"fileSearchTool"`
	// google_search_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#google_search_tool GoogleCesTool#google_search_tool}
	GoogleSearchTool *GoogleCesToolGoogleSearchTool `field:"optional" json:"googleSearchTool" yaml:"googleSearchTool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#id GoogleCesTool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#project GoogleCesTool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// python_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#python_function GoogleCesTool#python_function}
	PythonFunction *GoogleCesToolPythonFunction `field:"optional" json:"pythonFunction" yaml:"pythonFunction"`
	// The timeout for the tool execution.
	//
	// If not set, the default timeout is 30
	// seconds for SYNCHRONOUS tools and 60 seconds for ASYNCHRONOUS tools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#timeout GoogleCesTool#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#timeouts GoogleCesTool#timeouts}
	Timeouts *GoogleCesToolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// tool_fake_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#tool_fake_config GoogleCesTool#tool_fake_config}
	ToolFakeConfig *GoogleCesToolToolFakeConfig `field:"optional" json:"toolFakeConfig" yaml:"toolFakeConfig"`
	// widget_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#widget_tool GoogleCesTool#widget_tool}
	WidgetTool *GoogleCesToolWidgetTool `field:"optional" json:"widgetTool" yaml:"widgetTool"`
}

