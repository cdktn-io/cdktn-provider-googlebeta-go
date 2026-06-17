// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties struct {
	// The name of the GoldengateDeployment associated with the GoldengateConnection. Format: projects/{project}/locations/{location}/goldengateDeployments/{goldengate_deployment}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#goldengate_deployment_id GoogleOracleDatabaseGoldengateConnection#goldengate_deployment_id}
	GoldengateDeploymentId *string `field:"optional" json:"goldengateDeploymentId" yaml:"goldengateDeploymentId"`
	// The host of the GoldengateConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#host GoogleOracleDatabaseGoldengateConnection#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Input only. The password used to connect to the Oracle Goldengate in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password used to connect to the Oracle Goldengate.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// The port of the GoldengateConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#port GoogleOracleDatabaseGoldengateConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The technology type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The username credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

