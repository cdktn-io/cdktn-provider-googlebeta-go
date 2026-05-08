// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesservicelbpolicies


type GoogleNetworkServicesServiceLbPoliciesIsolationConfig struct {
	// The isolation granularity of the load balancer. Possible values: ["ISOLATION_GRANULARITY_UNSPECIFIED", "REGION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_service_lb_policies#isolation_granularity GoogleNetworkServicesServiceLbPolicies#isolation_granularity}
	IsolationGranularity *string `field:"optional" json:"isolationGranularity" yaml:"isolationGranularity"`
	// The isolation mode of the load balancer. Default value: "NEAREST" Possible values: ["ISOLATION_MODE_UNSPECIFIED", "NEAREST", "STRICT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_service_lb_policies#isolation_mode GoogleNetworkServicesServiceLbPolicies#isolation_mode}
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
}

