// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlefirestorefield


type GoogleFirestoreFieldTtlConfig struct {
	// The offset, relative to the timestamp value from the field, used to determine the document's expiration time.
	//
	// Formatted as the number of seconds followed by 's'. For example, "60s" represents an offset of one minute. The number of seconds must be between 1 and 2147483647 inclusive. To configure no offset, omit this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_firestore_field#expiration_offset GoogleFirestoreField#expiration_offset}
	ExpirationOffset *string `field:"optional" json:"expirationOffset" yaml:"expirationOffset"`
}

