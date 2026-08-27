// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool


type GoogleCesToolDataStoreToolDataStoreSource struct {
	// data_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_tool#data_store GoogleCesTool#data_store}
	DataStore *GoogleCesToolDataStoreToolDataStoreSourceDataStore `field:"optional" json:"dataStore" yaml:"dataStore"`
	// Optional. Filter specification for the DataStore. See: https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_tool#filter GoogleCesTool#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

