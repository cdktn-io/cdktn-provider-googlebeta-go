// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties struct {
	// Azure client ID of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#client_id GoogleOracleDatabaseGoldengateConnection#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client secret associated with the client id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#client_secret GoogleOracleDatabaseGoldengateConnection#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Optional Microsoft Fabric service endpoint. Default value: https://onelake.dfs.fabric.microsoft.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#endpoint GoogleOracleDatabaseGoldengateConnection#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The technology type of MicrosoftFabricConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// Azure tenant ID of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#tenant_id GoogleOracleDatabaseGoldengateConnection#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

