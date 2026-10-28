// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesaasruntimeunitkind


type GoogleSaasRuntimeUnitKindApplicationTemplateComponent struct {
	// Name of the component in the composite ApplicationTemplate (the root module ID, e.g. a composite component ID or 'native-components').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_saas_runtime_unit_kind#component GoogleSaasRuntimeUnitKind#component}
	Component *string `field:"optional" json:"component" yaml:"component"`
	// composite_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_saas_runtime_unit_kind#composite_ref GoogleSaasRuntimeUnitKind#composite_ref}
	CompositeRef *GoogleSaasRuntimeUnitKindApplicationTemplateComponentCompositeRef `field:"optional" json:"compositeRef" yaml:"compositeRef"`
	// Revision of the component. If the component does not have a revision, this is the revision of the composite ApplicationTemplate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_saas_runtime_unit_kind#revision GoogleSaasRuntimeUnitKind#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
}

