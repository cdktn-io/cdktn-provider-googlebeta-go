// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataplexdataproduct


type GoogleDataplexDataProductAccessApprovalConfig struct {
	// Specifies the email addresses of users who are potential approvers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_dataplex_data_product#approver_emails GoogleDataplexDataProduct#approver_emails}
	ApproverEmails *[]*string `field:"optional" json:"approverEmails" yaml:"approverEmails"`
}

