// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlestorageftpuser


type GoogleStorageFtpUserUserCredentials struct {
	// The name of the credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_storage_ftp_user#credential_name GoogleStorageFtpUser#credential_name}
	CredentialName *string `field:"optional" json:"credentialName" yaml:"credentialName"`
	// The type of the credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_storage_ftp_user#credential_type GoogleStorageFtpUser#credential_type}
	CredentialType *string `field:"optional" json:"credentialType" yaml:"credentialType"`
	// The SSH public key body.
	//
	// A file either absolute or relative path should be provided which contains the ssh public key using file() interpolation in Terraform, not recommended to have key as a literal string in config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_storage_ftp_user#ssh_public_key_body GoogleStorageFtpUser#ssh_public_key_body}
	SshPublicKeyBody *string `field:"optional" json:"sshPublicKeyBody" yaml:"sshPublicKeyBody"`
}

