// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlepubsubtopic


type GooglePubsubTopicSchemaSettings struct {
	// The name of the schema that messages published should be validated against.
	//
	// Format is projects/{project}/schemas/{schema}.
	// The value of this field will be _deleted-schema_
	// if the schema has been deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_pubsub_topic#schema GooglePubsubTopic#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The encoding of messages validated against schema. Default value: "ENCODING_UNSPECIFIED" Possible values: ["ENCODING_UNSPECIFIED", "JSON", "BINARY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_pubsub_topic#encoding GooglePubsubTopic#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The minimum (inclusive) revision allowed for validating messages.
	//
	// If empty or not present, allow any revision to be validated against last_revision or any revision created before.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_pubsub_topic#first_revision_id GooglePubsubTopic#first_revision_id}
	FirstRevisionId *string `field:"optional" json:"firstRevisionId" yaml:"firstRevisionId"`
	// The maximum (inclusive) revision allowed for validating messages.
	//
	// If empty or not present, allow any revision to be validated against first_revision or any revision created after.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_pubsub_topic#last_revision_id GooglePubsubTopic#last_revision_id}
	LastRevisionId *string `field:"optional" json:"lastRevisionId" yaml:"lastRevisionId"`
}

