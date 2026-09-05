// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatafusioninstance


type GoogleDataFusionInstanceMaintenancePolicy struct {
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_data_fusion_instance#maintenance_window GoogleDataFusionInstance#maintenance_window}
	MaintenanceWindow *GoogleDataFusionInstanceMaintenancePolicyMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
}

