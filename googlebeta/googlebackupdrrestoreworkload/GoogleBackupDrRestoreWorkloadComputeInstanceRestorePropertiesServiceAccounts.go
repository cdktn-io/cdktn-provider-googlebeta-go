// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccounts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_backup_dr_restore_workload#email GoogleBackupDrRestoreWorkload#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_backup_dr_restore_workload#scopes GoogleBackupDrRestoreWorkload#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

