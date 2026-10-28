// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledocumentaischemaversion


type GoogleDocumentAiSchemaVersionSchemaEntityTypesProperties struct {
	// The name of the property. Follows the same guidelines as the EntityType name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#name GoogleDocumentAiSchemaVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A reference to the value type of the property.
	//
	// This type is subject to the same conventions as the 'Entity.base_types' field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#value_type GoogleDocumentAiSchemaVersion#value_type}
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// User defined name for the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#display_name GoogleDocumentAiSchemaVersion#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Specifies how the entity's value is obtained. Possible values: ["EXTRACT", "DERIVE", "RELAXED_EXTRACT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#method GoogleDocumentAiSchemaVersion#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Occurrence type limits the number of instances an entity type appears in the document.
	//
	// Possible values: ["OPTIONAL_ONCE", "OPTIONAL_MULTIPLE", "REQUIRED_ONCE", "REQUIRED_MULTIPLE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_document_ai_schema_version#occurrence_type GoogleDocumentAiSchemaVersion#occurrence_type}
	OccurrenceType *string `field:"optional" json:"occurrenceType" yaml:"occurrenceType"`
}

