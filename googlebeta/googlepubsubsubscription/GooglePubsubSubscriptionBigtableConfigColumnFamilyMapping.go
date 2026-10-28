// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlepubsubsubscription


type GooglePubsubSubscriptionBigtableConfigColumnFamilyMapping struct {
	// delimited_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_pubsub_subscription#delimited_key GooglePubsubSubscription#delimited_key}
	DelimitedKey *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey `field:"optional" json:"delimitedKey" yaml:"delimitedKey"`
	// row_key_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_pubsub_subscription#row_key_schema GooglePubsubSubscription#row_key_schema}
	RowKeySchema *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchema `field:"optional" json:"rowKeySchema" yaml:"rowKeySchema"`
}

