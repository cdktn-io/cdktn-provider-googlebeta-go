// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryconnection


type GoogleBigqueryConnectionSalesforceDataCloud struct {
	// cross_cloud_cache_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_bigquery_connection#cross_cloud_cache_options GoogleBigqueryConnection#cross_cloud_cache_options}
	CrossCloudCacheOptions *GoogleBigqueryConnectionSalesforceDataCloudCrossCloudCacheOptions `field:"required" json:"crossCloudCacheOptions" yaml:"crossCloudCacheOptions"`
	// The URL to the user's Salesforce Data Cloud instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_bigquery_connection#instance_uri GoogleBigqueryConnection#instance_uri}
	InstanceUri *string `field:"required" json:"instanceUri" yaml:"instanceUri"`
	// The ID of the user's Salesforce tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_bigquery_connection#tenant_id GoogleBigqueryConnection#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

