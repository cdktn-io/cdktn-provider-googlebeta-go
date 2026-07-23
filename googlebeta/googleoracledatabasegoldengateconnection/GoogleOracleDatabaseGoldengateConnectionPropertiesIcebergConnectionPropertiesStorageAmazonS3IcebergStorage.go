// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAmazonS3IcebergStorage struct {
	// The access key ID of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#access_key_id GoogleOracleDatabaseGoldengateConnection#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// The bucket of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#bucket GoogleOracleDatabaseGoldengateConnection#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The region of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#region GoogleOracleDatabaseGoldengateConnection#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The scheme type of Amazon S3. Possible values: S3 S3A.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#scheme_type GoogleOracleDatabaseGoldengateConnection#scheme_type}
	SchemeType *string `field:"required" json:"schemeType" yaml:"schemeType"`
	// The endpoint of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#endpoint GoogleOracleDatabaseGoldengateConnection#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The secret access key of Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#secret_access_key_secret GoogleOracleDatabaseGoldengateConnection#secret_access_key_secret}
	SecretAccessKeySecret *string `field:"optional" json:"secretAccessKeySecret" yaml:"secretAccessKeySecret"`
}

