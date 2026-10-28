// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledocumentaischemaversion


type GoogleDocumentAiSchemaVersionSchemaEntityTypes struct {
	// Name of the type. It must be unique within the schema file and cannot be a "Common Type".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#name GoogleDocumentAiSchemaVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The entity type that this type is derived from. For now, one and only one should be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#base_types GoogleDocumentAiSchemaVersion#base_types}
	BaseTypes *[]*string `field:"optional" json:"baseTypes" yaml:"baseTypes"`
	// User defined name for the type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#display_name GoogleDocumentAiSchemaVersion#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// enum_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#enum_values GoogleDocumentAiSchemaVersion#enum_values}
	EnumValues *GoogleDocumentAiSchemaVersionSchemaEntityTypesEnumValues `field:"optional" json:"enumValues" yaml:"enumValues"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#properties GoogleDocumentAiSchemaVersion#properties}
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

