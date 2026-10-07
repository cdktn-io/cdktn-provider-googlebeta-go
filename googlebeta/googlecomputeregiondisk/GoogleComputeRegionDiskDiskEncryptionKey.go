// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregiondisk


type GoogleComputeRegionDiskDiskEncryptionKey struct {
	// The name of the encryption key that is stored in Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_region_disk#kms_key_name GoogleComputeRegionDisk#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_region_disk#raw_key GoogleComputeRegionDisk#raw_key}
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_region_disk#raw_key_wo GoogleComputeRegionDisk#raw_key_wo}
	RawKeyWo *string `field:"optional" json:"rawKeyWo" yaml:"rawKeyWo"`
	// Triggers update of 'raw_key_wo' write-only.
	//
	// Increment this value when an update to 'raw_key_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_region_disk#raw_key_wo_version GoogleComputeRegionDisk#raw_key_wo_version}
	RawKeyWoVersion *string `field:"optional" json:"rawKeyWoVersion" yaml:"rawKeyWoVersion"`
	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.
	//
	// You can provide either the rawKey or the rsaEncryptedKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_region_disk#rsa_encrypted_key GoogleComputeRegionDisk#rsa_encrypted_key}
	RsaEncryptedKey *string `field:"optional" json:"rsaEncryptedKey" yaml:"rsaEncryptedKey"`
	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.
	//
	// You can provide either the rawKey or the rsaEncryptedKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_region_disk#rsa_encrypted_key_wo GoogleComputeRegionDisk#rsa_encrypted_key_wo}
	RsaEncryptedKeyWo *string `field:"optional" json:"rsaEncryptedKeyWo" yaml:"rsaEncryptedKeyWo"`
	// Triggers update of 'rsa_encrypted_key_wo' write-only.
	//
	// Increment this value when an update to 'rsa_encrypted_key_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_region_disk#rsa_encrypted_key_wo_version GoogleComputeRegionDisk#rsa_encrypted_key_wo_version}
	RsaEncryptedKeyWoVersion *string `field:"optional" json:"rsaEncryptedKeyWoVersion" yaml:"rsaEncryptedKeyWoVersion"`
}

