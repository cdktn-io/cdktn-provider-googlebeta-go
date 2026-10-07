// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression struct {
	// The regular expression for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#key_regex GoogleDataLossPreventionContentPolicy#key_regex}
	KeyRegex *string `field:"required" json:"keyRegex" yaml:"keyRegex"`
	// The regular expression for the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#value_regex GoogleDataLossPreventionContentPolicy#value_regex}
	ValueRegex *string `field:"required" json:"valueRegex" yaml:"valueRegex"`
}

