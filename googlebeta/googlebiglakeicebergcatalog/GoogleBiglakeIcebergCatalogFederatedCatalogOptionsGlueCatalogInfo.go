// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakeicebergcatalog


type GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo struct {
	// The AWS region where the Glue catalog is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_biglake_iceberg_catalog#aws_region GoogleBiglakeIcebergCatalog#aws_region}
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The ARN of the AWS IAM role to assume for accessing the Glue catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_biglake_iceberg_catalog#aws_role_arn GoogleBiglakeIcebergCatalog#aws_role_arn}
	AwsRoleArn *string `field:"required" json:"awsRoleArn" yaml:"awsRoleArn"`
	// The AWS Glue warehouse identifier (account ID or S3 table bucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_biglake_iceberg_catalog#warehouse GoogleBiglakeIcebergCatalog#warehouse}
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
}

