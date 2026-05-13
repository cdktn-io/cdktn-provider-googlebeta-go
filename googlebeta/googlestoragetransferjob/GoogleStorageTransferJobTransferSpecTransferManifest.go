// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecTransferManifest struct {
	// Cloud Storage path to the manifest CSV.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_storage_transfer_job#location GoogleStorageTransferJob#location}
	Location *string `field:"required" json:"location" yaml:"location"`
}

