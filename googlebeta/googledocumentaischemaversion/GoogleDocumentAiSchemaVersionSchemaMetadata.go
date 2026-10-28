// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledocumentaischemaversion


type GoogleDocumentAiSchemaVersionSchemaMetadata struct {
	// If true, on a given page, there can be multiple 'document' annotations covering it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#document_allow_multiple_labels GoogleDocumentAiSchemaVersion#document_allow_multiple_labels}
	DocumentAllowMultipleLabels interface{} `field:"optional" json:"documentAllowMultipleLabels" yaml:"documentAllowMultipleLabels"`
	// If true, a 'document' entity type can be applied to subdocument (splitting).
	//
	// Otherwise, it can only be applied to the entire document (classification).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#document_splitter GoogleDocumentAiSchemaVersion#document_splitter}
	DocumentSplitter interface{} `field:"optional" json:"documentSplitter" yaml:"documentSplitter"`
	// If set, all the nested entities must be prefixed with the parents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#prefixed_naming_on_properties GoogleDocumentAiSchemaVersion#prefixed_naming_on_properties}
	PrefixedNamingOnProperties interface{} `field:"optional" json:"prefixedNamingOnProperties" yaml:"prefixedNamingOnProperties"`
	// If set, this will skip the naming format validation in the schema.
	//
	// So the string values in 'DocumentSchema.EntityType.name' and 'DocumentSchema.EntityType.Property.name' will not be checked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#skip_naming_validation GoogleDocumentAiSchemaVersion#skip_naming_validation}
	SkipNamingValidation interface{} `field:"optional" json:"skipNamingValidation" yaml:"skipNamingValidation"`
}

