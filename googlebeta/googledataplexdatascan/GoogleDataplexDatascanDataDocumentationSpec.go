// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataplexdatascan


type GoogleDataplexDatascanDataDocumentationSpec struct {
	// If set, the latest DataScan job result will be published to Knowledge Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataplex_datascan#catalog_publishing_enabled GoogleDataplexDatascan#catalog_publishing_enabled}
	CatalogPublishingEnabled interface{} `field:"optional" json:"catalogPublishingEnabled" yaml:"catalogPublishingEnabled"`
	// The SQL dialect to use in the generated SQL queries.
	//
	// If not specified, the default dialect is Google SQL. Possible values: ["GOOGLE_SQL", "SPARK_SQL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataplex_datascan#sql_dialect GoogleDataplexDatascan#sql_dialect}
	SqlDialect *string `field:"optional" json:"sqlDialect" yaml:"sqlDialect"`
}

