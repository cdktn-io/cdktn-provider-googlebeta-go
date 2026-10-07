// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkmanagementnetworkmonitoringprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleNetworkManagementNetworkMonitoringProviderConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The location of the Network Monitoring Provider. Currently only 'global' is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_management_network_monitoring_provider#location GoogleNetworkManagementNetworkMonitoringProvider#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID to use for the Network Monitoring Provider. This will become the last component of the provider's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_management_network_monitoring_provider#network_monitoring_provider_id GoogleNetworkManagementNetworkMonitoringProvider#network_monitoring_provider_id}
	NetworkMonitoringProviderId *string `field:"required" json:"networkMonitoringProviderId" yaml:"networkMonitoringProviderId"`
	// The type of the Network Monitoring Provider. Currently only 'EXTERNAL' is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_management_network_monitoring_provider#provider_type GoogleNetworkManagementNetworkMonitoringProvider#provider_type}
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
	// The deletion policy for the Network Monitoring Provider.
	//
	// Setting 'deletion_policy = "FORCE"' forces the deletion of all nested resources
	// (MonitoringPoints, NetworkPaths, WebPaths) belonging to this provider on deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_management_network_monitoring_provider#deletion_policy GoogleNetworkManagementNetworkMonitoringProvider#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_management_network_monitoring_provider#id GoogleNetworkManagementNetworkMonitoringProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_management_network_monitoring_provider#project GoogleNetworkManagementNetworkMonitoringProvider#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_management_network_monitoring_provider#timeouts GoogleNetworkManagementNetworkMonitoringProvider#timeouts}
	Timeouts *GoogleNetworkManagementNetworkMonitoringProviderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

