// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshublistingsubscription


type GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionCloudStorageConfig struct {
	// avro_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#avro_config GoogleBigqueryAnalyticsHubListingSubscription#avro_config}
	AvroConfig *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionCloudStorageConfigAvroConfig `field:"optional" json:"avroConfig" yaml:"avroConfig"`
	// User-provided name for the Cloud Storage bucket.
	//
	// The bucket must be created by the user.
	// The bucket name must be without any prefix like "gs://". See the
	// [bucket naming requirements](https://cloud.google.com/storage/docs/buckets#naming).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#bucket GoogleBigqueryAnalyticsHubListingSubscription#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// User-provided format string specifying how to represent datetimes in Cloud Storage filenames. See the [datetime format guidance](https://cloud.google.com/pubsub/docs/create-cloudstorage-subscription#file_names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#filename_datetime_format GoogleBigqueryAnalyticsHubListingSubscription#filename_datetime_format}
	FilenameDatetimeFormat *string `field:"optional" json:"filenameDatetimeFormat" yaml:"filenameDatetimeFormat"`
	// User-provided prefix for Cloud Storage filename. See the [object naming requirements](https://cloud.google.com/storage/docs/objects#naming).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#filename_prefix GoogleBigqueryAnalyticsHubListingSubscription#filename_prefix}
	FilenamePrefix *string `field:"optional" json:"filenamePrefix" yaml:"filenamePrefix"`
	// User-provided suffix for Cloud Storage filename. See the [object naming requirements](https://cloud.google.com/storage/docs/objects#naming). Must not end in "/".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#filename_suffix GoogleBigqueryAnalyticsHubListingSubscription#filename_suffix}
	FilenameSuffix *string `field:"optional" json:"filenameSuffix" yaml:"filenameSuffix"`
	// The maximum bytes that can be written to a Cloud Storage file before a new file is created.
	//
	// Min 1 KB, max 10 GiB. The maxBytes limit may be exceeded in cases where messages are larger
	// than the limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#max_bytes GoogleBigqueryAnalyticsHubListingSubscription#max_bytes}
	MaxBytes *string `field:"optional" json:"maxBytes" yaml:"maxBytes"`
	// The maximum duration that can elapse before a new Cloud Storage file is created.
	//
	// Min 1 minute, max 10 minutes, default 5 minutes. May not exceed the subscription's
	// acknowledgement deadline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#max_duration GoogleBigqueryAnalyticsHubListingSubscription#max_duration}
	MaxDuration *string `field:"optional" json:"maxDuration" yaml:"maxDuration"`
	// The maximum number of messages that can be written to a Cloud Storage file before a new file is created.
	//
	// Min 1000 messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#max_messages GoogleBigqueryAnalyticsHubListingSubscription#max_messages}
	MaxMessages *string `field:"optional" json:"maxMessages" yaml:"maxMessages"`
	// The service account to use to write to Cloud Storage.
	//
	// The subscription creator or updater that
	// specifies this field must have 'iam.serviceAccounts.actAs' permission on the service account.
	// If not specified, the Pub/Sub service agent,
	// service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com, is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#service_account_email GoogleBigqueryAnalyticsHubListingSubscription#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

