// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebackupdrautoprotectionpolicy


type GoogleBackupDrAutoProtectionPolicyCriteriaMatchingConditionsLabelCondition struct {
	// The key of the label to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_backup_dr_auto_protection_policy#key GoogleBackupDrAutoProtectionPolicy#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The values of the label to match.
	//
	// A resource matches the condition if its label
	// value is equal to any of the listed values. Currently, only one value is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_backup_dr_auto_protection_policy#values GoogleBackupDrAutoProtectionPolicy#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

