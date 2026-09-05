// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesagentgateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleNetworkServicesAgentGatewayConfig struct {
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
	// The location of the agent gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#location GoogleNetworkServicesAgentGateway#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the AgentGateway resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#name GoogleNetworkServicesAgentGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#deletion_policy GoogleNetworkServicesAgentGateway#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#description GoogleNetworkServicesAgentGateway#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// google_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#google_managed GoogleNetworkServicesAgentGateway#google_managed}
	GoogleManaged *GoogleNetworkServicesAgentGatewayGoogleManaged `field:"optional" json:"googleManaged" yaml:"googleManaged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#id GoogleNetworkServicesAgentGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the AgentGateway resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#labels GoogleNetworkServicesAgentGateway#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#network_config GoogleNetworkServicesAgentGateway#network_config}
	NetworkConfig *GoogleNetworkServicesAgentGatewayNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#project GoogleNetworkServicesAgentGateway#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// List of protocols supported by an Agent Gateway. Possible values: ["MCP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#protocols GoogleNetworkServicesAgentGateway#protocols}
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// A list of Agent registries containing the agents, MCP servers and tools governed by the Agent Gateway.
	//
	// Note: Currently limited to project-scoped registries Must be of format
	// '//agentregistry.googleapis.com/{version}/projects/{{project}}/locations/{{location}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#registries GoogleNetworkServicesAgentGateway#registries}
	Registries *[]*string `field:"optional" json:"registries" yaml:"registries"`
	// self_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#self_managed GoogleNetworkServicesAgentGateway#self_managed}
	SelfManaged *GoogleNetworkServicesAgentGatewaySelfManaged `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_network_services_agent_gateway#timeouts GoogleNetworkServicesAgentGateway#timeouts}
	Timeouts *GoogleNetworkServicesAgentGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

