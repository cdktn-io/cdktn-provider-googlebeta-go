// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceTargetEnvironment struct {
	// Required. Target project for the Compute Engine instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_backup_dr_restore_workload#project GoogleBackupDrRestoreWorkload#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Required. The zone of the Compute Engine instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_backup_dr_restore_workload#zone GoogleBackupDrRestoreWorkload#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// If true, use the BackupDR P4SA credentials for same-project restores. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_backup_dr_restore_workload#use_project_service_account GoogleBackupDrRestoreWorkload#use_project_service_account}
	UseProjectServiceAccount interface{} `field:"optional" json:"useProjectServiceAccount" yaml:"useProjectServiceAccount"`
}

