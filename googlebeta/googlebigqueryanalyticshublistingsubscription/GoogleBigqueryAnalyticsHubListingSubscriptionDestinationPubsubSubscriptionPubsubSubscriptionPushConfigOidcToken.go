// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshublistingsubscription


type GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionPushConfigOidcToken struct {
	// Audience to be used when generating OIDC token.
	//
	// The audience claim identifies the recipients
	// that the JWT is intended for. The audience value is a single case-sensitive string. Having
	// multiple values (array) for the audience field is not supported. More info about the OIDC JWT
	// token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3 Note: if not specified,
	// the Push endpoint URL will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#audience GoogleBigqueryAnalyticsHubListingSubscription#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Service account email used for generating the OIDC token. For more information on setting up authentication, see Push subscriptions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#service_account_email GoogleBigqueryAnalyticsHubListingSubscription#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

