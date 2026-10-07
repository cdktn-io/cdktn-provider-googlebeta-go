// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesexample


type GoogleCesExampleMessagesChunksBlob struct {
	// Raw bytes of the blob.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_example#data GoogleCesExample#data}
	Data *string `field:"required" json:"data" yaml:"data"`
	// The IANA standard MIME type of the source data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_example#mime_type GoogleCesExample#mime_type}
	MimeType *string `field:"required" json:"mimeType" yaml:"mimeType"`
}

