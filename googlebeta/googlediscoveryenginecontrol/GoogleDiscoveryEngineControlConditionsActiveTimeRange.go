// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlConditionsActiveTimeRange struct {
	// The end time of the active time range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#end_time GoogleDiscoveryEngineControl#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The start time of the active time range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#start_time GoogleDiscoveryEngineControl#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

