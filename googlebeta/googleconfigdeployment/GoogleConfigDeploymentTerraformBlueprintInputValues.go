// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleconfigdeployment


type GoogleConfigDeploymentTerraformBlueprintInputValues struct {
	// The value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_config_deployment#input_value GoogleConfigDeployment#input_value}
	InputValue *string `field:"required" json:"inputValue" yaml:"inputValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_config_deployment#variable_name GoogleConfigDeployment#variable_name}.
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
}

