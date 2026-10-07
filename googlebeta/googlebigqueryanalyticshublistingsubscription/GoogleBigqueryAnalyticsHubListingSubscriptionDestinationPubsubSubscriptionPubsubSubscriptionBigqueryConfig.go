// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshublistingsubscription


type GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionBigqueryConfig struct {
	// When true and 'useTopicSchema' is true, any fields that are a part of the topic schema that are not part of the BigQuery table schema are dropped when writing to BigQuery.
	//
	// Otherwise, the schemas
	// must be kept in sync and any messages with extra fields are not written and remain in the
	// subscription's backlog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#drop_unknown_fields GoogleBigqueryAnalyticsHubListingSubscription#drop_unknown_fields}
	DropUnknownFields interface{} `field:"optional" json:"dropUnknownFields" yaml:"dropUnknownFields"`
	// The service account to use to write to BigQuery.
	//
	// The subscription creator or updater that
	// specifies this field must have 'iam.serviceAccounts.actAs' permission on the service account.
	// If not specified, the Pub/Sub service agent,
	// service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com, is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#service_account_email GoogleBigqueryAnalyticsHubListingSubscription#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The name of the table to which to write data, of the form {projectId}.{datasetId}.{tableId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#table GoogleBigqueryAnalyticsHubListingSubscription#table}
	Table *string `field:"optional" json:"table" yaml:"table"`
	// When true, use the BigQuery table's schema as the columns to write to in BigQuery.
	//
	// 'useTableSchema' and 'useTopicSchema' cannot be enabled at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#use_table_schema GoogleBigqueryAnalyticsHubListingSubscription#use_table_schema}
	UseTableSchema interface{} `field:"optional" json:"useTableSchema" yaml:"useTableSchema"`
	// When true, use the topic's schema as the columns to write to in BigQuery, if it exists.
	//
	// 'useTopicSchema' and 'useTableSchema' cannot be enabled at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#use_topic_schema GoogleBigqueryAnalyticsHubListingSubscription#use_topic_schema}
	UseTopicSchema interface{} `field:"optional" json:"useTopicSchema" yaml:"useTopicSchema"`
	// When true, write the subscription name, message_id, publish_time, attributes, and ordering_key to additional columns in the table.
	//
	// The subscription name, message_id, and publish_time fields
	// are put in their own columns while all other message properties (other than data) are written
	// to a JSON object in the attributes column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#write_metadata GoogleBigqueryAnalyticsHubListingSubscription#write_metadata}
	WriteMetadata interface{} `field:"optional" json:"writeMetadata" yaml:"writeMetadata"`
}

