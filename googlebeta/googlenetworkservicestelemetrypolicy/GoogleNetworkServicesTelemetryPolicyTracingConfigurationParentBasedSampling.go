// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicestelemetrypolicy


type GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSampling struct {
	// Provides a master switch to respect a sampled bit in incoming traceparentID headers. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_services_telemetry_policy#enabled GoogleNetworkServicesTelemetryPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the sampling rate to use when the sampled bit is set in the incoming traceparentID header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_services_telemetry_policy#sampling_rate GoogleNetworkServicesTelemetryPolicy#sampling_rate}
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
}

