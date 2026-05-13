// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecRulesTemplateReference struct {
	// The resource name of the template entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dataplex_datascan#name GoogleDataplexDatascan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dataplex_datascan#values GoogleDataplexDatascan#values}
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

