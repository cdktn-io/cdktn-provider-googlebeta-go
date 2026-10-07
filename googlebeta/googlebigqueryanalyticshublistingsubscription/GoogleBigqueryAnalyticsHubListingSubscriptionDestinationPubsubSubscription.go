// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshublistingsubscription


type GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscription struct {
	// pubsub_subscription block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#pubsub_subscription GoogleBigqueryAnalyticsHubListingSubscription#pubsub_subscription}
	PubsubSubscription *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscription `field:"required" json:"pubsubSubscription" yaml:"pubsubSubscription"`
}

