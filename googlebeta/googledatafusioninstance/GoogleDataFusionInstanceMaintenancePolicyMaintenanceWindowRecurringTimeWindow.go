// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatafusioninstance


type GoogleDataFusionInstanceMaintenancePolicyMaintenanceWindowRecurringTimeWindow struct {
	// An RRULE with format RFC-5545 for how this window reccurs.
	//
	// They go on for the span of time between the start and end time. The only supported FREQ value is "WEEKLY". To have something repeat every weekday, use: "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_data_fusion_instance#recurrence GoogleDataFusionInstance#recurrence}
	Recurrence *string `field:"required" json:"recurrence" yaml:"recurrence"`
	// window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_data_fusion_instance#window GoogleDataFusionInstance#window}
	Window *GoogleDataFusionInstanceMaintenancePolicyMaintenanceWindowRecurringTimeWindowWindow `field:"required" json:"window" yaml:"window"`
}

