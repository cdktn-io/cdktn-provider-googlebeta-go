// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileSpannerProfile struct {
	// The full project and resource path for Spanner database. Format: projects/{project}/instances/{instance}/databases/{database}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_datastream_connection_profile#database GoogleDatastreamConnectionProfile#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The regional Spanner endpoint. Format: https://spanner.{region}.rep.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_datastream_connection_profile#host GoogleDatastreamConnectionProfile#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
}

