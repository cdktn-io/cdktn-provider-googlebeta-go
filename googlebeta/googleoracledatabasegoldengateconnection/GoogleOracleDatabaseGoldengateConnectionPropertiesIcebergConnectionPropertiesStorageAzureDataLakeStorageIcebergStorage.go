// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorage struct {
	// The account of Azure Data Lake Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#azure_account GoogleOracleDatabaseGoldengateConnection#azure_account}
	AzureAccount *string `field:"required" json:"azureAccount" yaml:"azureAccount"`
	// The container of Azure Data Lake Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#container GoogleOracleDatabaseGoldengateConnection#container}
	Container *string `field:"required" json:"container" yaml:"container"`
	// The account key of Azure Data Lake Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#account_key_secret GoogleOracleDatabaseGoldengateConnection#account_key_secret}
	AccountKeySecret *string `field:"optional" json:"accountKeySecret" yaml:"accountKeySecret"`
	// The endpoint of Azure Data Lake Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#endpoint GoogleOracleDatabaseGoldengateConnection#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

