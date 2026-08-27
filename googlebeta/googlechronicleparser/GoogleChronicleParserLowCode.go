// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparser


type GoogleChronicleParserLowCode struct {
	// field_extractors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser#field_extractors GoogleChronicleParser#field_extractors}
	FieldExtractors *GoogleChronicleParserLowCodeFieldExtractors `field:"optional" json:"fieldExtractors" yaml:"fieldExtractors"`
	// The log used to create this low code parser in the UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser#log GoogleChronicleParser#log}
	Log *string `field:"optional" json:"log" yaml:"log"`
}

