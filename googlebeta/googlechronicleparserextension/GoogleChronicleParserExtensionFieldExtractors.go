// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparserextension


type GoogleChronicleParserExtensionFieldExtractors struct {
	// Whether to append repeated fields or not. When false, repeated fields will be replaced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_chronicle_parser_extension#append_repeated_fields GoogleChronicleParserExtension#append_repeated_fields}
	AppendRepeatedFields interface{} `field:"optional" json:"appendRepeatedFields" yaml:"appendRepeatedFields"`
	// extractors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_chronicle_parser_extension#extractors GoogleChronicleParserExtension#extractors}
	Extractors interface{} `field:"optional" json:"extractors" yaml:"extractors"`
	// Possible values: JSON CSV XML.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_chronicle_parser_extension#log_format GoogleChronicleParserExtension#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// preprocess_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_chronicle_parser_extension#preprocess_config GoogleChronicleParserExtension#preprocess_config}
	PreprocessConfig *GoogleChronicleParserExtensionFieldExtractorsPreprocessConfig `field:"optional" json:"preprocessConfig" yaml:"preprocessConfig"`
}

