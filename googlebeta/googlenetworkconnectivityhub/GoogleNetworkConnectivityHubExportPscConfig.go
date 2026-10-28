// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkconnectivityhub


type GoogleNetworkConnectivityHubExportPscConfig struct {
	// Whether Private Service Connect endpoints for global Google APIs are propagated. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_connectivity_hub#global_google_apis GoogleNetworkConnectivityHub#global_google_apis}
	GlobalGoogleApis interface{} `field:"optional" json:"globalGoogleApis" yaml:"globalGoogleApis"`
	// Whether Private Service Connect endpoints for published services (regional ILBs) and regional Google APIs are propagated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_connectivity_hub#published_services_and_regional_google_apis GoogleNetworkConnectivityHub#published_services_and_regional_google_apis}
	PublishedServicesAndRegionalGoogleApis interface{} `field:"optional" json:"publishedServicesAndRegionalGoogleApis" yaml:"publishedServicesAndRegionalGoogleApis"`
}

