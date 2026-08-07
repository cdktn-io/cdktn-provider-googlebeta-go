// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparserextension


type GoogleChronicleParserExtensionFieldExtractorsPreprocessConfig struct {
	// GROK Regex to extract the structured part of the log. syntax documentation: www.elastic.co/guide/en/logstash/current/plugins-filters-grok.html.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_chronicle_parser_extension#grok_regex GoogleChronicleParserExtension#grok_regex}
	GrokRegex *string `field:"optional" json:"grokRegex" yaml:"grokRegex"`
	// Target field name for the structured part of the log. This should match a SEMANTIC identifier from the grok expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_chronicle_parser_extension#target GoogleChronicleParserExtension#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

