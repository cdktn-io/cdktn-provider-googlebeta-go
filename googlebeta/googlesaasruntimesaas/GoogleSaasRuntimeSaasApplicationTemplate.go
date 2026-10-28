// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesaasruntimesaas


type GoogleSaasRuntimeSaasApplicationTemplate struct {
	// Reference to the ApplicationTemplate resource, in the format 'projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{applicationTemplate}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_saas_runtime_saas#application_template GoogleSaasRuntimeSaas#application_template}
	ApplicationTemplate *string `field:"required" json:"applicationTemplate" yaml:"applicationTemplate"`
	// Revision of the ApplicationTemplate to import, as a full resource name prefixed by 'application_template': 'projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{applicationTemplate}/revisions/{revision}'.
	//
	// Changes to the revision trigger a new import. If empty, the service
	// stores 'application_template' but doesn't import it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_saas_runtime_saas#revision GoogleSaasRuntimeSaas#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
}

