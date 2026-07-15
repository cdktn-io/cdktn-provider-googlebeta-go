// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatafusioninstance


type GoogleDataFusionInstanceMaintenancePolicyMaintenanceWindowRecurringTimeWindowWindow struct {
	// The end time of the time window provided in RFC 3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_data_fusion_instance#end_time GoogleDataFusionInstance#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// The start time of the time window provided in RFC 3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_data_fusion_instance#start_time GoogleDataFusionInstance#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

