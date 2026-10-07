// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputedisk


type GoogleComputeDiskDiskEncryptionKey struct {
	// The self link of the encryption key used to encrypt the disk.
	//
	// Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// ('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com') must have
	// 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_disk#kms_key_self_link GoogleComputeDisk#kms_key_self_link}
	KmsKeySelfLink *string `field:"optional" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
	// The service account used for the encryption request for the given KMS key.
	//
	// If absent, the Compute Engine Service Agent service account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_disk#kms_key_service_account GoogleComputeDisk#kms_key_service_account}
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_disk#raw_key GoogleComputeDisk#raw_key}
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_disk#raw_key_wo GoogleComputeDisk#raw_key_wo}
	RawKeyWo *string `field:"optional" json:"rawKeyWo" yaml:"rawKeyWo"`
	// Triggers update of 'raw_key_wo' write-only.
	//
	// Increment this value when an update to 'raw_key_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_disk#raw_key_wo_version GoogleComputeDisk#raw_key_wo_version}
	RawKeyWoVersion *string `field:"optional" json:"rawKeyWoVersion" yaml:"rawKeyWoVersion"`
	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.
	//
	// You can provide either the rawKey or the rsaEncryptedKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_disk#rsa_encrypted_key GoogleComputeDisk#rsa_encrypted_key}
	RsaEncryptedKey *string `field:"optional" json:"rsaEncryptedKey" yaml:"rsaEncryptedKey"`
	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.
	//
	// You can provide either the rawKey or the rsaEncryptedKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_disk#rsa_encrypted_key_wo GoogleComputeDisk#rsa_encrypted_key_wo}
	RsaEncryptedKeyWo *string `field:"optional" json:"rsaEncryptedKeyWo" yaml:"rsaEncryptedKeyWo"`
	// Triggers update of 'rsa_encrypted_key_wo' write-only.
	//
	// Increment this value when an update to 'rsa_encrypted_key_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_disk#rsa_encrypted_key_wo_version GoogleComputeDisk#rsa_encrypted_key_wo_version}
	RsaEncryptedKeyWoVersion *string `field:"optional" json:"rsaEncryptedKeyWoVersion" yaml:"rsaEncryptedKeyWoVersion"`
}

