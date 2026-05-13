// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormat struct {
	// Specifies whether the client connects directly to the host[:port] in the connection URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_datastream_connection_profile#direct_connection GoogleDatastreamConnectionProfile#direct_connection}
	DirectConnection interface{} `field:"optional" json:"directConnection" yaml:"directConnection"`
}

