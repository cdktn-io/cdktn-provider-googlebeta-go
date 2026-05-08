// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesaasruntimeunitoperation


type GoogleSaasRuntimeUnitOperationUpgrade struct {
	// input_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_operation#input_variables GoogleSaasRuntimeUnitOperation#input_variables}
	InputVariables interface{} `field:"optional" json:"inputVariables" yaml:"inputVariables"`
	// Reference to the Release object to use for the Unit. (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_operation#release GoogleSaasRuntimeUnitOperation#release}
	Release *string `field:"optional" json:"release" yaml:"release"`
}

