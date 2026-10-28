// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlepubsubsubscription


type GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey struct {
	// The key fields to construct the row key from.
	//
	// The fields must be present in the message
	// as top-level fields; JSON path expressions will not traverse into nested objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_pubsub_subscription#key_fields GooglePubsubSubscription#key_fields}
	KeyFields *[]*string `field:"required" json:"keyFields" yaml:"keyFields"`
	// The byte sequence used to delimit concatenated key fields, as a base64-encoded string (for example, 'Iw==' for '#').
	//
	// Must be specified if multiple key fields are used. The
	// delimiter must contain at least 1 character and at most 50 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_pubsub_subscription#delimiter GooglePubsubSubscription#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
}

