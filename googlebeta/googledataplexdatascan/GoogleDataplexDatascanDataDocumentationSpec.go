// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataplexdatascan


type GoogleDataplexDatascanDataDocumentationSpec struct {
	// If set, the latest DataScan job result will be published to Knowledge Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_dataplex_datascan#catalog_publishing_enabled GoogleDataplexDatascan#catalog_publishing_enabled}
	CatalogPublishingEnabled interface{} `field:"optional" json:"catalogPublishingEnabled" yaml:"catalogPublishingEnabled"`
}

