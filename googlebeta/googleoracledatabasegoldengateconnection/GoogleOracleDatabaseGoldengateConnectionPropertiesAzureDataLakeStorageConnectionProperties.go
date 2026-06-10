// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties struct {
	// Sets the Azure storage account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#account GoogleOracleDatabaseGoldengateConnection#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Azure storage account key. This property is required when 'authentication_type' is set to 'SHARED_KEY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#account_key_secret GoogleOracleDatabaseGoldengateConnection#account_key_secret}
	AccountKeySecret *string `field:"optional" json:"accountKeySecret" yaml:"accountKeySecret"`
	// Authentication mechanism to access Azure Data Lake Storage. Possible values: SHARED_KEY SHARED_ACCESS_SIGNATURE AZURE_ACTIVE_DIRECTORY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#authentication_type GoogleOracleDatabaseGoldengateConnection#authentication_type}
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The endpoint used for authentication with Microsoft Entra ID (formerly Azure Active Directory). Default value: https://login.microsoftonline.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#azure_authority_host GoogleOracleDatabaseGoldengateConnection#azure_authority_host}
	AzureAuthorityHost *string `field:"optional" json:"azureAuthorityHost" yaml:"azureAuthorityHost"`
	// Azure tenant ID of the application. This property is required when 'authentication_type' is set to 'AZURE_ACTIVE_DIRECTORY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#azure_tenant_id GoogleOracleDatabaseGoldengateConnection#azure_tenant_id}
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// Azure client ID of the application. This property is required when 'authentication_type' is set to 'AZURE_ACTIVE_DIRECTORY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#client_id GoogleOracleDatabaseGoldengateConnection#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Azure client secret (aka application password) for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#client_secret GoogleOracleDatabaseGoldengateConnection#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Azure Storage service endpoint. e.g: https://test.blob.core.windows.net.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#endpoint GoogleOracleDatabaseGoldengateConnection#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Credential that uses a shared access signature (SAS) to authenticate to an Azure Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#sas_token_secret GoogleOracleDatabaseGoldengateConnection#sas_token_secret}
	SasTokenSecret *string `field:"optional" json:"sasTokenSecret" yaml:"sasTokenSecret"`
	// The technology type of AzureDataLakeStorageConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
}

