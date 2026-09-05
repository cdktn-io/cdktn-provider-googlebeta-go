// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestoolset


type GoogleCesToolsetConnectorToolset struct {
	// The full resource name of the referenced Integration Connectors Connection. Format: 'projects/{project}/locations/{location}/connections/{connection}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#connection GoogleCesToolset#connection}
	Connection *string `field:"required" json:"connection" yaml:"connection"`
	// connector_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#connector_actions GoogleCesToolset#connector_actions}
	ConnectorActions interface{} `field:"required" json:"connectorActions" yaml:"connectorActions"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#auth_config GoogleCesToolset#auth_config}
	AuthConfig *GoogleCesToolsetConnectorToolsetAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
}

