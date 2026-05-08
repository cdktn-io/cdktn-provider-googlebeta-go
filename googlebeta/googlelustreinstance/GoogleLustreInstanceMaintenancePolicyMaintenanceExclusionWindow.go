// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlelustreinstance


type GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindow struct {
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#end_date GoogleLustreInstance#end_date}
	EndDate *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDate `field:"required" json:"endDate" yaml:"endDate"`
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#start_date GoogleLustreInstance#start_date}
	StartDate *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#time GoogleLustreInstance#time}
	Time *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime `field:"required" json:"time" yaml:"time"`
}

