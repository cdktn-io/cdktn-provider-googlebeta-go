// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengatedeployment


type GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData struct {
	// The Goldengate deployment console username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_deployment#admin_username GoogleOracleDatabaseGoldengateDeployment#admin_username}
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// The name given to the Goldengate service deployment.
	//
	// The name must be 1 to
	// 32 characters long, must contain only alphanumeric characters and must
	// start with a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_deployment#deployment GoogleOracleDatabaseGoldengateDeployment#deployment}
	Deployment *string `field:"required" json:"deployment" yaml:"deployment"`
	// The Goldengate deployment console password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_deployment#admin_password GoogleOracleDatabaseGoldengateDeployment#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Input only. The Goldengate deployment console password secret version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_deployment#admin_password_secret_version GoogleOracleDatabaseGoldengateDeployment#admin_password_secret_version}
	AdminPasswordSecretVersion *string `field:"optional" json:"adminPasswordSecretVersion" yaml:"adminPasswordSecretVersion"`
	// group_roles_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_deployment#group_roles_mapping GoogleOracleDatabaseGoldengateDeployment#group_roles_mapping}
	GroupRolesMapping *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping `field:"optional" json:"groupRolesMapping" yaml:"groupRolesMapping"`
	// Version of OGG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_deployment#ogg_version GoogleOracleDatabaseGoldengateDeployment#ogg_version}
	OggVersion *string `field:"optional" json:"oggVersion" yaml:"oggVersion"`
}

