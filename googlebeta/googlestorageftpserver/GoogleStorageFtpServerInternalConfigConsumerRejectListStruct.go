// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlestorageftpserver


type GoogleStorageFtpServerInternalConfigConsumerRejectListStruct struct {
	// The project that is rejected from connecting, in the format 'projects/{project}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_storage_ftp_server#project GoogleStorageFtpServer#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

