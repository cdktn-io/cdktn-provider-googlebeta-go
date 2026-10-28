// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicestelemetrypolicy


type GoogleNetworkServicesTelemetryPolicyTracingConfiguration struct {
	// custom_span_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_services_telemetry_policy#custom_span_attributes GoogleNetworkServicesTelemetryPolicy#custom_span_attributes}
	CustomSpanAttributes interface{} `field:"optional" json:"customSpanAttributes" yaml:"customSpanAttributes"`
	// parent_based_sampling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_services_telemetry_policy#parent_based_sampling GoogleNetworkServicesTelemetryPolicy#parent_based_sampling}
	ParentBasedSampling *GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSampling `field:"optional" json:"parentBasedSampling" yaml:"parentBasedSampling"`
	// Specifies the tracing sampling rate.
	//
	// The sampling rate is a float number in the interval [0, 1]. The sampling rate is applied to all requests to the targeted resources. The default value is 0.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_services_telemetry_policy#sampling_rate GoogleNetworkServicesTelemetryPolicy#sampling_rate}
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
}

