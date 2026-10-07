// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshublistingsubscription


type GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionExpirationPolicy struct {
	// Specifies the "time-to-live" duration for an associated resource.
	//
	// The resource expires if it
	// is not active for a period of 'ttl'. The definition of "activity" depends on the type of the
	// associated resource. The minimum and maximum allowed values for 'ttl' depend on the type of
	// the associated resource, as well. If 'ttl' is not set, the associated resource never expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#ttl GoogleBigqueryAnalyticsHubListingSubscription#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

