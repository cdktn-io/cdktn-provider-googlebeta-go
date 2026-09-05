// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorage struct {
	// The type of Iceberg storage. Possible values: AMAZON_S3 GOOGLE_CLOUD_STORAGE AZURE_DATA_LAKE_STORAGE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#storage_type GoogleOracleDatabaseGoldengateConnection#storage_type}
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// amazon_s3_iceberg_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#amazon_s3_iceberg_storage GoogleOracleDatabaseGoldengateConnection#amazon_s3_iceberg_storage}
	AmazonS3IcebergStorage *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAmazonS3IcebergStorage `field:"optional" json:"amazonS3IcebergStorage" yaml:"amazonS3IcebergStorage"`
	// azure_data_lake_storage_iceberg_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#azure_data_lake_storage_iceberg_storage GoogleOracleDatabaseGoldengateConnection#azure_data_lake_storage_iceberg_storage}
	AzureDataLakeStorageIcebergStorage *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorage `field:"optional" json:"azureDataLakeStorageIcebergStorage" yaml:"azureDataLakeStorageIcebergStorage"`
	// google_cloud_storage_iceberg_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#google_cloud_storage_iceberg_storage GoogleOracleDatabaseGoldengateConnection#google_cloud_storage_iceberg_storage}
	GoogleCloudStorageIcebergStorage *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageGoogleCloudStorageIcebergStorage `field:"optional" json:"googleCloudStorageIcebergStorage" yaml:"googleCloudStorageIcebergStorage"`
}

