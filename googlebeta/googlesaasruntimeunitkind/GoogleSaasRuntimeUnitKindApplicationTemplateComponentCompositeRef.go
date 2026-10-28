// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesaasruntimeunitkind


type GoogleSaasRuntimeUnitKindApplicationTemplateComponentCompositeRef struct {
	// Reference to the ApplicationTemplate resource, in the format 'projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{applicationTemplate}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_saas_runtime_unit_kind#application_template GoogleSaasRuntimeUnitKind#application_template}
	ApplicationTemplate *string `field:"required" json:"applicationTemplate" yaml:"applicationTemplate"`
	// Revision of the ApplicationTemplate, as a full resource name prefixed by 'application_template'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_saas_runtime_unit_kind#revision GoogleSaasRuntimeUnitKind#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
}

