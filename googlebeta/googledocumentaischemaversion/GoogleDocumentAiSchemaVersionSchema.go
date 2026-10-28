// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledocumentaischemaversion


type GoogleDocumentAiSchemaVersionSchema struct {
	// Description of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#description GoogleDocumentAiSchemaVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name to show users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#display_name GoogleDocumentAiSchemaVersion#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Document-level prompt provided by the user.
	//
	// This custom text is injected into the AI model's prompt to provide extra, document-wide guidance for processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#document_prompt GoogleDocumentAiSchemaVersion#document_prompt}
	DocumentPrompt *string `field:"optional" json:"documentPrompt" yaml:"documentPrompt"`
	// entity_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#entity_types GoogleDocumentAiSchemaVersion#entity_types}
	EntityTypes interface{} `field:"optional" json:"entityTypes" yaml:"entityTypes"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#metadata GoogleDocumentAiSchemaVersion#metadata}
	Metadata *GoogleDocumentAiSchemaVersionSchemaMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

