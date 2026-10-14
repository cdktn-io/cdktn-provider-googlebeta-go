// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebackupdrautoprotectionpolicy


type GoogleBackupDrAutoProtectionPolicyBackupPlanDetails struct {
	// The resource name of the backup plan to apply to the matching resources, in the format 'projects/{project}/locations/{location}/backupPlans/{backup_plan_id}'.
	//
	// The backup plan
	// must be in the same location as the auto protection policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_backup_dr_auto_protection_policy#backup_plan GoogleBackupDrAutoProtectionPolicy#backup_plan}
	BackupPlan *string `field:"required" json:"backupPlan" yaml:"backupPlan"`
	// The type of resource that this policy automatically protects, for example 'compute.googleapis.com/Instance' or 'compute.googleapis.com/Disk'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_backup_dr_auto_protection_policy#resource_type GoogleBackupDrAutoProtectionPolicy#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

