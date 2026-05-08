// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2service


type GoogleCloudRunV2ServiceMultiRegionSettings struct {
	// The list of regions to deploy the multi-region Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#regions GoogleCloudRunV2Service#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

