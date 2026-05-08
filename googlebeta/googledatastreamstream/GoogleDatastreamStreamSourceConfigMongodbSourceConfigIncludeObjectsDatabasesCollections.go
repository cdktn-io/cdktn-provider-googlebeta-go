// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatastreamstream


type GoogleDatastreamStreamSourceConfigMongodbSourceConfigIncludeObjectsDatabasesCollections struct {
	// Collection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#collection GoogleDatastreamStream#collection}
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#fields GoogleDatastreamStream#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
}

