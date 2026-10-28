// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesaasruntimeunitkind


type GoogleSaasRuntimeUnitKindAppParams struct {
	// Grouping used to construct the name of the App Hub Application.
	//
	// Multiple UnitKinds can specify the same group to use the same
	// Application across their respective Units. Corresponds to the
	// app_boundary_id in the ADC composite ApplicationTemplate. If empty,
	// the service creates no App Hub application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_saas_runtime_unit_kind#group GoogleSaasRuntimeUnitKind#group}
	Group *string `field:"optional" json:"group" yaml:"group"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_saas_runtime_unit_kind#scope GoogleSaasRuntimeUnitKind#scope}
	Scope *GoogleSaasRuntimeUnitKindAppParamsScope `field:"optional" json:"scope" yaml:"scope"`
}

