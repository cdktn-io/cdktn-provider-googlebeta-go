// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshublistingsubscription


type GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionPushConfig struct {
	// Endpoint configuration attributes that can be used to control different aspects of the message delivery.
	//
	// The only currently supported attribute is 'x-goog-version', which you can use to change the format of the
	// pushed message. This attribute indicates the version of the data expected by the endpoint. This controls
	// the shape of the pushed message (i.e., its fields and metadata). If not present during the
	// 'CreateSubscription' call, it will default to the version of the Pub/Sub API used to make such call.
	// If not present in a 'ModifyPushConfig' call, its value will not be changed. 'GetSubscription' calls
	// will always return a valid version, even if the subscription was created without this attribute.
	// The only supported values for the 'x-goog-version' attribute are: 'v1beta1': uses the push format
	// defined in the v1beta1 Pub/Sub API. 'v1' or 'v1beta2': uses the push format defined in the v1 Pub/Sub API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#attributes GoogleBigqueryAnalyticsHubListingSubscription#attributes}
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// no_wrapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#no_wrapper GoogleBigqueryAnalyticsHubListingSubscription#no_wrapper}
	NoWrapper *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionPushConfigNoWrapper `field:"optional" json:"noWrapper" yaml:"noWrapper"`
	// oidc_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#oidc_token GoogleBigqueryAnalyticsHubListingSubscription#oidc_token}
	OidcToken *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionPushConfigOidcToken `field:"optional" json:"oidcToken" yaml:"oidcToken"`
	// A URL locating the endpoint to which messages should be pushed. For example, a Webhook endpoint might use 'https://example.com/push'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#push_endpoint GoogleBigqueryAnalyticsHubListingSubscription#push_endpoint}
	PushEndpoint *string `field:"optional" json:"pushEndpoint" yaml:"pushEndpoint"`
}

