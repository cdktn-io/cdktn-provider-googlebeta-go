// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengatedeployment


type GoogleOracleDatabaseGoldengateDeploymentProperties struct {
	// A valid Goldengate Deployment type. For a list of supported types, use the 'ListGoldengateDeploymentTypes' operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#deployment_type GoogleOracleDatabaseGoldengateDeployment#deployment_type}
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// ogg_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#ogg_data GoogleOracleDatabaseGoldengateDeployment#ogg_data}
	OggData *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData `field:"required" json:"oggData" yaml:"oggData"`
	// backup_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#backup_schedule GoogleOracleDatabaseGoldengateDeployment#backup_schedule}
	BackupSchedule *GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupSchedule `field:"optional" json:"backupSchedule" yaml:"backupSchedule"`
	// The Minimum number of OCPUs to be made available for this Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#cpu_core_count GoogleOracleDatabaseGoldengateDeployment#cpu_core_count}
	CpuCoreCount *float64 `field:"optional" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// deployment_diagnostic_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#deployment_diagnostic_data GoogleOracleDatabaseGoldengateDeployment#deployment_diagnostic_data}
	DeploymentDiagnosticData *GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData `field:"optional" json:"deploymentDiagnosticData" yaml:"deploymentDiagnosticData"`
	// The description of the GoldengateDeployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#description GoogleOracleDatabaseGoldengateDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The environment type of the GoldengateDeployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#environment_type GoogleOracleDatabaseGoldengateDeployment#environment_type}
	EnvironmentType *string `field:"optional" json:"environmentType" yaml:"environmentType"`
	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#is_auto_scaling_enabled GoogleOracleDatabaseGoldengateDeployment#is_auto_scaling_enabled}
	IsAutoScalingEnabled interface{} `field:"optional" json:"isAutoScalingEnabled" yaml:"isAutoScalingEnabled"`
	// The Oracle license model that applies to a Deployment. Possible values: LICENSE_INCLUDED BRING_YOUR_OWN_LICENSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#license_model GoogleOracleDatabaseGoldengateDeployment#license_model}
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// maintenance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#maintenance_config GoogleOracleDatabaseGoldengateDeployment#maintenance_config}
	MaintenanceConfig *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig `field:"optional" json:"maintenanceConfig" yaml:"maintenanceConfig"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_deployment#maintenance_window GoogleOracleDatabaseGoldengateDeployment#maintenance_window}
	MaintenanceWindow *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
}

