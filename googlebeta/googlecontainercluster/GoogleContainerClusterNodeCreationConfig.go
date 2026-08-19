// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodeCreationConfig struct {
	// NodeCreationMode defines the settings of node creation mode.
	//
	// Accepted values are:
	// * VIA_KUBELET: Kubelet registers itself.
	// * VIA_CONTROL_PLANE: gcp-controller-manager automatically creates the node object after CSR approval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_container_cluster#node_creation_mode GoogleContainerCluster#node_creation_mode}
	NodeCreationMode *string `field:"required" json:"nodeCreationMode" yaml:"nodeCreationMode"`
}

