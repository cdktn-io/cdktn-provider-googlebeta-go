// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionnetworkpolicytrafficclassificationrule


type GoogleComputeRegionNetworkPolicyTrafficClassificationRuleAction struct {
	// DSCP mode.
	//
	// When set to AUTO, the DSCP value will be picked automatically based on selected trafficClass. Otherwise,dscpValue needs to be explicitly specified. Possible values: ["AUTO", "CUSTOM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#dscp_mode GoogleComputeRegionNetworkPolicyTrafficClassificationRule#dscp_mode}
	DscpMode *string `field:"optional" json:"dscpMode" yaml:"dscpMode"`
	// Custom DSCP value from 0-63 range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#dscp_value GoogleComputeRegionNetworkPolicyTrafficClassificationRule#dscp_value}
	DscpValue *float64 `field:"optional" json:"dscpValue" yaml:"dscpValue"`
	// The traffic class that should be applied to the matching packet. Possible values: ["TC1", "TC2", "TC3", "TC4", "TC5", "TC6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#traffic_class GoogleComputeRegionNetworkPolicyTrafficClassificationRule#traffic_class}
	TrafficClass *string `field:"optional" json:"trafficClass" yaml:"trafficClass"`
	// Always apply_traffic_classification for Traffic Classification Rules. Default value: "apply_traffic_classification" Possible values: ["apply_traffic_classification"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#type GoogleComputeRegionNetworkPolicyTrafficClassificationRule#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

