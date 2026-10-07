// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlestorageftpserver


type GoogleStorageFtpServerInternalConfig struct {
	// consumer_accept_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_storage_ftp_server#consumer_accept_list GoogleStorageFtpServer#consumer_accept_list}
	ConsumerAcceptList interface{} `field:"optional" json:"consumerAcceptList" yaml:"consumerAcceptList"`
	// consumer_reject_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_storage_ftp_server#consumer_reject_list GoogleStorageFtpServer#consumer_reject_list}
	ConsumerRejectList interface{} `field:"optional" json:"consumerRejectList" yaml:"consumerRejectList"`
}

