// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatafusioninstance


type GoogleDataFusionInstanceMaintenancePolicyMaintenanceWindow struct {
	// recurring_time_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_fusion_instance#recurring_time_window GoogleDataFusionInstance#recurring_time_window}
	RecurringTimeWindow *GoogleDataFusionInstanceMaintenancePolicyMaintenanceWindowRecurringTimeWindow `field:"required" json:"recurringTimeWindow" yaml:"recurringTimeWindow"`
}

