// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlememorystoreinstance


type GoogleMemorystoreInstanceManagedBackupSource struct {
	// Example: 'projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backup}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_memorystore_instance#backup GoogleMemorystoreInstance#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

