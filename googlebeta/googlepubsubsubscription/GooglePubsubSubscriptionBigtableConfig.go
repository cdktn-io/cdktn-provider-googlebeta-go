// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlepubsubsubscription


type GooglePubsubSubscriptionBigtableConfig struct {
	// The unique name of the table to write messages to, of the form 'projects/{project}/instances/{instance}/tables/{table}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_pubsub_subscription#table GooglePubsubSubscription#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// The app profile to use for the Bigtable writes.
	//
	// If not specified, the "default" application
	// profile will be used. The app profile must use single-cluster routing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_pubsub_subscription#app_profile_id GooglePubsubSubscription#app_profile_id}
	AppProfileId *string `field:"optional" json:"appProfileId" yaml:"appProfileId"`
	// column_family_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_pubsub_subscription#column_family_mapping GooglePubsubSubscription#column_family_mapping}
	ColumnFamilyMapping *GooglePubsubSubscriptionBigtableConfigColumnFamilyMapping `field:"optional" json:"columnFamilyMapping" yaml:"columnFamilyMapping"`
	// The service account to use to write to Bigtable.
	//
	// The subscription creator or updater that
	// specifies this field must have 'iam.serviceAccounts.actAs' permission on the service account.
	// If not specified, the Pub/Sub [service agent](https://cloud.google.com/iam/docs/service-agents),
	// service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com, is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_pubsub_subscription#service_account_email GooglePubsubSubscription#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// When true, write the subscription name, messageId, publishTime, attributes, and orderingKey to additional columns in the table under the 'pubsub_metadata' column family.
	//
	// The subscription name,
	// messageId, and publishTime fields are put in their own columns while all other message properties
	// (other than data) are written to a JSON object in the attributes column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_pubsub_subscription#write_metadata GooglePubsubSubscription#write_metadata}
	WriteMetadata interface{} `field:"optional" json:"writeMetadata" yaml:"writeMetadata"`
}

