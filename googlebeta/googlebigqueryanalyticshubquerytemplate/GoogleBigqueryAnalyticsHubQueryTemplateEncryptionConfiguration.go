// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshubquerytemplate


type GoogleBigqueryAnalyticsHubQueryTemplateEncryptionConfiguration struct {
	// The KMS key used to encrypt the query template. Format: 'projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_bigquery_analytics_hub_query_template#kms_key_name GoogleBigqueryAnalyticsHubQueryTemplate#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

