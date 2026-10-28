// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool


type GoogleCesToolDataStoreToolModalityConfigsSnippetsConfig struct {
	// Whether snippets are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_ces_tool#enable_snippets GoogleCesTool#enable_snippets}
	EnableSnippets interface{} `field:"optional" json:"enableSnippets" yaml:"enableSnippets"`
	// Number of snippets to return per query. If unset, returns all snippets from the service by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_ces_tool#max_snippets GoogleCesTool#max_snippets}
	MaxSnippets *float64 `field:"optional" json:"maxSnippets" yaml:"maxSnippets"`
}

