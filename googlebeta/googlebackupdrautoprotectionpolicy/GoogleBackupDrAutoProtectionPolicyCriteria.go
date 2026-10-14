// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebackupdrautoprotectionpolicy


type GoogleBackupDrAutoProtectionPolicyCriteria struct {
	// matching_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_backup_dr_auto_protection_policy#matching_conditions GoogleBackupDrAutoProtectionPolicy#matching_conditions}
	MatchingConditions interface{} `field:"required" json:"matchingConditions" yaml:"matchingConditions"`
}

