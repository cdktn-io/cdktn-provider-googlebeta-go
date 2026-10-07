// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNetworkConfigAdditionalNodeNetworkConfigs struct {
	// Name of the VPC where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_container_node_pool#network GoogleContainerNodePool#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The IP stack type of the additional node interface.
	//
	// Possible values are IPV4, IPV4_IPV6 and IPV6. If unset, the value is inferred from the additional subnetwork.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_container_node_pool#stack_type GoogleContainerNodePool#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// Name of the subnetwork where the additional interface belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_container_node_pool#subnetwork GoogleContainerNodePool#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

