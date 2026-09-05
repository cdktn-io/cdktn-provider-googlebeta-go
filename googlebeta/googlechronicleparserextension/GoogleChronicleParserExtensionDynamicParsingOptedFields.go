// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparserextension


type GoogleChronicleParserExtensionDynamicParsingOptedFields struct {
	// Path of the log field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_parser_extension#path GoogleChronicleParserExtension#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Sample value of the log field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_parser_extension#sample_value GoogleChronicleParserExtension#sample_value}
	SampleValue *string `field:"optional" json:"sampleValue" yaml:"sampleValue"`
}

