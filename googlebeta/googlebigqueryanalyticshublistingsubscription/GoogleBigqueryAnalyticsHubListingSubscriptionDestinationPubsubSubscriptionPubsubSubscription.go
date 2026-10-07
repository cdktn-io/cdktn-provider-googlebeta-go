// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshublistingsubscription


type GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscription struct {
	// Name of the subscription. Format is 'projects/{project}/subscriptions/{sub}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#name GoogleBigqueryAnalyticsHubListingSubscription#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The approximate amount of time (on a best-effort basis) Pub/Sub waits for the subscriber to acknowledge receipt before resending the message.
	//
	// In the interval after the message is delivered
	// and before it is acknowledged, it is considered to be outstanding. During that time period, the
	// message will not be redelivered (on a best-effort basis). For pull subscriptions, this value is
	// used as the initial value for the ack deadline. To override this value for a given message, call
	// 'ModifyAckDeadline' with the corresponding 'ack_id' if using non-streaming pull or send the
	// 'ack_id' in a 'StreamingModifyAckDeadlineRequest' if using streaming pull. The minimum custom
	// deadline you can specify is 10 seconds. The maximum custom deadline you can specify is 600
	// seconds (10 minutes). If this parameter is 0, a default value of 10 seconds is used. For push
	// delivery, this value is also used to set the request timeout for the call to the push endpoint.
	// If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver
	// the message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#ack_deadline_seconds GoogleBigqueryAnalyticsHubListingSubscription#ack_deadline_seconds}
	AckDeadlineSeconds *float64 `field:"optional" json:"ackDeadlineSeconds" yaml:"ackDeadlineSeconds"`
	// bigquery_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#bigquery_config GoogleBigqueryAnalyticsHubListingSubscription#bigquery_config}
	BigqueryConfig *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionBigqueryConfig `field:"optional" json:"bigqueryConfig" yaml:"bigqueryConfig"`
	// cloud_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#cloud_storage_config GoogleBigqueryAnalyticsHubListingSubscription#cloud_storage_config}
	CloudStorageConfig *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionCloudStorageConfig `field:"optional" json:"cloudStorageConfig" yaml:"cloudStorageConfig"`
	// dead_letter_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#dead_letter_policy GoogleBigqueryAnalyticsHubListingSubscription#dead_letter_policy}
	DeadLetterPolicy *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionDeadLetterPolicy `field:"optional" json:"deadLetterPolicy" yaml:"deadLetterPolicy"`
	// Indicates whether the subscription is detached from its topic.
	//
	// Detached subscriptions don't
	// receive messages from their topic and don't retain any backlog. 'Pull' and 'StreamingPull'
	// requests will return FAILED_PRECONDITION. If the subscription is a push subscription, pushes
	// to the endpoint will not be made.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#detached GoogleBigqueryAnalyticsHubListingSubscription#detached}
	Detached interface{} `field:"optional" json:"detached" yaml:"detached"`
	// If true, Pub/Sub provides the following guarantees for the delivery of a message with a given value of 'message_id' on this subscription: The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires.
	//
	// An acknowledged message will
	// not be resent to a subscriber. Note that subscribers may still receive multiple copies of a
	// message when 'enableExactlyOnceDelivery' is true if the message was published multiple times by
	// a publisher client. These copies are considered distinct by Pub/Sub and have distinct 'message_id'
	// values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#enable_exactly_once_delivery GoogleBigqueryAnalyticsHubListingSubscription#enable_exactly_once_delivery}
	EnableExactlyOnceDelivery interface{} `field:"optional" json:"enableExactlyOnceDelivery" yaml:"enableExactlyOnceDelivery"`
	// If true, messages published with the same 'ordering_key' in 'PubsubMessage' will be delivered to the subscribers in the order in which they are received by the Pub/Sub system.
	//
	// Otherwise, they may be delivered in any order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#enable_message_ordering GoogleBigqueryAnalyticsHubListingSubscription#enable_message_ordering}
	EnableMessageOrdering interface{} `field:"optional" json:"enableMessageOrdering" yaml:"enableMessageOrdering"`
	// expiration_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#expiration_policy GoogleBigqueryAnalyticsHubListingSubscription#expiration_policy}
	ExpirationPolicy *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionExpirationPolicy `field:"optional" json:"expirationPolicy" yaml:"expirationPolicy"`
	// An expression written in the Pub/Sub filter language.
	//
	// If non-empty, then only 'PubsubMessage's
	// whose 'attributes' field matches the filter are delivered on this subscription. If empty, then
	// no messages are filtered out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#filter GoogleBigqueryAnalyticsHubListingSubscription#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// See [Creating and managing labels](https://cloud.google.com/pubsub/docs/labels).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#labels GoogleBigqueryAnalyticsHubListingSubscription#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// How long to retain unacknowledged messages in the subscription's backlog, from the moment a message is published.
	//
	// If 'retainAckedMessages' is true, then this also configures the retention
	// of acknowledged messages, and thus configures how far back in time a Seek can be done. Defaults
	// to 7 days. Cannot be more than 31 days or less than 10 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#message_retention_duration GoogleBigqueryAnalyticsHubListingSubscription#message_retention_duration}
	MessageRetentionDuration *string `field:"optional" json:"messageRetentionDuration" yaml:"messageRetentionDuration"`
	// push_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#push_config GoogleBigqueryAnalyticsHubListingSubscription#push_config}
	PushConfig *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionPushConfig `field:"optional" json:"pushConfig" yaml:"pushConfig"`
	// Indicates whether to retain acknowledged messages.
	//
	// If true, then messages are not expunged from
	// the subscription's backlog, even if they are acknowledged, until they fall out of the
	// 'messageRetentionDuration' window. This must be true if you would like to Seek to a timestamp
	// in the past to replay previously-acknowledged messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#retain_acked_messages GoogleBigqueryAnalyticsHubListingSubscription#retain_acked_messages}
	RetainAckedMessages interface{} `field:"optional" json:"retainAckedMessages" yaml:"retainAckedMessages"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#retry_policy GoogleBigqueryAnalyticsHubListingSubscription#retry_policy}
	RetryPolicy *GoogleBigqueryAnalyticsHubListingSubscriptionDestinationPubsubSubscriptionPubsubSubscriptionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

