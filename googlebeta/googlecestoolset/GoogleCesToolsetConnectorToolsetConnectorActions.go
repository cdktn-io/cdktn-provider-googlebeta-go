// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestoolset


type GoogleCesToolsetConnectorToolsetConnectorActions struct {
	// ID of a Connection action for the tool to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#connection_action_id GoogleCesToolset#connection_action_id}
	ConnectionActionId *string `field:"optional" json:"connectionActionId" yaml:"connectionActionId"`
	// entity_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#entity_operation GoogleCesToolset#entity_operation}
	EntityOperation *GoogleCesToolsetConnectorToolsetConnectorActionsEntityOperation `field:"optional" json:"entityOperation" yaml:"entityOperation"`
	// Entity fields to use as inputs for the operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#input_fields GoogleCesToolset#input_fields}
	InputFields *[]*string `field:"optional" json:"inputFields" yaml:"inputFields"`
	// Entity fields to return from the operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#output_fields GoogleCesToolset#output_fields}
	OutputFields *[]*string `field:"optional" json:"outputFields" yaml:"outputFields"`
}

