// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlestoragecontrolprojectintelligenceconfig


type GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations struct {
	// List of locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_control_project_intelligence_config#locations GoogleStorageControlProjectIntelligenceConfig#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
}

