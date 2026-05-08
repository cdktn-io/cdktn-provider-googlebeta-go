// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleprivatecacertificate


type GooglePrivatecaCertificateConfigSubjectKeyId struct {
	// The value of the KeyId in lowercase hexadecimal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privateca_certificate#key_id GooglePrivatecaCertificate#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

