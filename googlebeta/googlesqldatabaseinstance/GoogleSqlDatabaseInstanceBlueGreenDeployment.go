// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceBlueGreenDeployment struct {
	// A unique identifier for the blue-green deployment within the project and region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_sql_database_instance#deployment_name GoogleSqlDatabaseInstance#deployment_name}
	DeploymentName *string `field:"required" json:"deploymentName" yaml:"deploymentName"`
	// Permissible values: "SOURCE" (default) or "TARGET". Changing to "TARGET" triggers the atomic switchover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_sql_database_instance#active_environment GoogleSqlDatabaseInstance#active_environment}
	ActiveEnvironment *string `field:"optional" json:"activeEnvironment" yaml:"activeEnvironment"`
	// Defaults to false.
	//
	// If true, removing the blue_green_deployment block after switchover automatically deletes the former Blue (Source) instance from GCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_sql_database_instance#delete_old_source_on_destroy GoogleSqlDatabaseInstance#delete_old_source_on_destroy}
	DeleteOldSourceOnDestroy interface{} `field:"optional" json:"deleteOldSourceOnDestroy" yaml:"deleteOldSourceOnDestroy"`
	// The target database version for Major Version Upgrades (e.g., MYSQL_8_4 or POSTGRES_15).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_sql_database_instance#target_database_version GoogleSqlDatabaseInstance#target_database_version}
	TargetDatabaseVersion *string `field:"optional" json:"targetDatabaseVersion" yaml:"targetDatabaseVersion"`
}

