// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebackupdrautoprotectionpolicy


type GoogleBackupDrAutoProtectionPolicyCriteriaMatchingConditions struct {
	// label_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_backup_dr_auto_protection_policy#label_condition GoogleBackupDrAutoProtectionPolicy#label_condition}
	LabelCondition *GoogleBackupDrAutoProtectionPolicyCriteriaMatchingConditionsLabelCondition `field:"optional" json:"labelCondition" yaml:"labelCondition"`
}

