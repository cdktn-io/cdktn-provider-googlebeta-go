// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties struct {
	// Access key ID to access the Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#access_key_id GoogleOracleDatabaseGoldengateConnection#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// The Amazon Endpoint for S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#endpoint GoogleOracleDatabaseGoldengateConnection#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The name of the AWS region where the bucket is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#region GoogleOracleDatabaseGoldengateConnection#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Secret access key to access the Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#secret_access_key_secret GoogleOracleDatabaseGoldengateConnection#secret_access_key_secret}
	SecretAccessKeySecret *string `field:"optional" json:"secretAccessKeySecret" yaml:"secretAccessKeySecret"`
	// The technology type of AmazonS3Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
}

