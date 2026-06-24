// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparserextension


type GoogleChronicleParserExtensionFieldExtractorsExtractors struct {
	// Path in generated event which is to be populated.
	//
	// This is required if the
	// FieldExtractor is used to specify the parser extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_chronicle_parser_extension#destination_path GoogleChronicleParserExtension#destination_path}
	DestinationPath *string `field:"optional" json:"destinationPath" yaml:"destinationPath"`
	// Field path could be a json path, xml path or csv column name depending on log format.
	//
	// It refers to a section or substring in raw log.
	// This is required if the FieldExtractor is used to specify the parser
	// extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_chronicle_parser_extension#field_path GoogleChronicleParserExtension#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
	// Operator used for precondition. Possible values: EQUALS NOT_EQUALS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_chronicle_parser_extension#precondition_op GoogleChronicleParserExtension#precondition_op}
	PreconditionOp *string `field:"optional" json:"preconditionOp" yaml:"preconditionOp"`
	// Precondition path could be a json path, xml path or csv column name depending on log format.
	//
	// It refers to a section or substring in raw log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_chronicle_parser_extension#precondition_path GoogleChronicleParserExtension#precondition_path}
	PreconditionPath *string `field:"optional" json:"preconditionPath" yaml:"preconditionPath"`
	// Precondition value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_chronicle_parser_extension#precondition_value GoogleChronicleParserExtension#precondition_value}
	PreconditionValue *string `field:"optional" json:"preconditionValue" yaml:"preconditionValue"`
	// Value to be mapped to the destination path directly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_chronicle_parser_extension#value GoogleChronicleParserExtension#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

