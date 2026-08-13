// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlediscoveryenginedatastore


type GoogleDiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesLayoutParsingConfig struct {
	// If true, the processed document will be made available for the GetProcessedDocument API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_discovery_engine_data_store#enable_get_processed_document GoogleDiscoveryEngineDataStore#enable_get_processed_document}
	EnableGetProcessedDocument interface{} `field:"optional" json:"enableGetProcessedDocument" yaml:"enableGetProcessedDocument"`
	// If true, the LLM based annotation is added to the image during parsing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_discovery_engine_data_store#enable_image_annotation GoogleDiscoveryEngineDataStore#enable_image_annotation}
	EnableImageAnnotation interface{} `field:"optional" json:"enableImageAnnotation" yaml:"enableImageAnnotation"`
	// If true, the pdf layout will be refined using an LLM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_discovery_engine_data_store#enable_llm_layout_parsing GoogleDiscoveryEngineDataStore#enable_llm_layout_parsing}
	EnableLlmLayoutParsing interface{} `field:"optional" json:"enableLlmLayoutParsing" yaml:"enableLlmLayoutParsing"`
	// If true, the LLM based annotation is added to the table during parsing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_discovery_engine_data_store#enable_table_annotation GoogleDiscoveryEngineDataStore#enable_table_annotation}
	EnableTableAnnotation interface{} `field:"optional" json:"enableTableAnnotation" yaml:"enableTableAnnotation"`
	// List of HTML classes to exclude from the parsed content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_discovery_engine_data_store#exclude_html_classes GoogleDiscoveryEngineDataStore#exclude_html_classes}
	ExcludeHtmlClasses *[]*string `field:"optional" json:"excludeHtmlClasses" yaml:"excludeHtmlClasses"`
	// List of HTML elements to exclude from the parsed content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_discovery_engine_data_store#exclude_html_elements GoogleDiscoveryEngineDataStore#exclude_html_elements}
	ExcludeHtmlElements *[]*string `field:"optional" json:"excludeHtmlElements" yaml:"excludeHtmlElements"`
	// List of HTML ids to exclude from the parsed content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_discovery_engine_data_store#exclude_html_ids GoogleDiscoveryEngineDataStore#exclude_html_ids}
	ExcludeHtmlIds *[]*string `field:"optional" json:"excludeHtmlIds" yaml:"excludeHtmlIds"`
	// Contains the required structure types to extract from the document. Supported values: 'shareholder-structure'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_discovery_engine_data_store#structured_content_types GoogleDiscoveryEngineDataStore#structured_content_types}
	StructuredContentTypes *[]*string `field:"optional" json:"structuredContentTypes" yaml:"structuredContentTypes"`
}

