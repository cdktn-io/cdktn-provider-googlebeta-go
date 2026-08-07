// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryroutine


type GoogleBigqueryRoutineArgumentsTableTypeColumns struct {
	// The name of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_bigquery_routine#name GoogleBigqueryRoutine#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A JSON schema for the data type of the column.
	//
	// Required unless argumentKind = ANY_TYPE.
	// ~>**NOTE**: Because this field expects a JSON string, any changes to the string
	// will create a diff, even if the JSON itself hasn't changed. If the API returns
	// a different value for the same schema, e.g. it switched the order of values
	// or replaced STRUCT field type with RECORD field type, we currently cannot
	// suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_bigquery_routine#type GoogleBigqueryRoutine#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

