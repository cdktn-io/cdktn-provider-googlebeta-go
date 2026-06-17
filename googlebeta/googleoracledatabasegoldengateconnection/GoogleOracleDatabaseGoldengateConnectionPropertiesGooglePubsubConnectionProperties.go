// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties struct {
	// The content of the service account key file containing the credentials required to use Google Pub/Sub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#service_account_key_file GoogleOracleDatabaseGoldengateConnection#service_account_key_file}
	ServiceAccountKeyFile *string `field:"optional" json:"serviceAccountKeyFile" yaml:"serviceAccountKeyFile"`
	// The technology type of GooglePubsubConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
}

