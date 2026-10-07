// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleenvironment


type GoogleChronicleEnvironmentDynamicParameters struct {
	// The ID of the dynamic parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_environment#dynamic_parameter_id GoogleChronicleEnvironment#dynamic_parameter_id}
	DynamicParameterId *float64 `field:"required" json:"dynamicParameterId" yaml:"dynamicParameterId"`
	// The value of the dynamic parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_environment#value GoogleChronicleEnvironment#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

