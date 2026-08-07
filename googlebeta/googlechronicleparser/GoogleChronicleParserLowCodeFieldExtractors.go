// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparser


type GoogleChronicleParserLowCodeFieldExtractors struct {
	// Whether to append repeated fields or not. When false, repeated fields will be replaced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_chronicle_parser#append_repeated_fields GoogleChronicleParser#append_repeated_fields}
	AppendRepeatedFields interface{} `field:"optional" json:"appendRepeatedFields" yaml:"appendRepeatedFields"`
	// extractors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_chronicle_parser#extractors GoogleChronicleParser#extractors}
	Extractors interface{} `field:"optional" json:"extractors" yaml:"extractors"`
	// Possible values: JSON CSV XML.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_chronicle_parser#log_format GoogleChronicleParser#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// preprocess_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_chronicle_parser#preprocess_config GoogleChronicleParser#preprocess_config}
	PreprocessConfig *GoogleChronicleParserLowCodeFieldExtractorsPreprocessConfig `field:"optional" json:"preprocessConfig" yaml:"preprocessConfig"`
}

