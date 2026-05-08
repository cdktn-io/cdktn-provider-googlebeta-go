// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainerattachedcluster


type GoogleContainerAttachedClusterProxyConfigKubernetesSecret struct {
	// Name of the kubernetes secret containing the proxy config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_attached_cluster#name GoogleContainerAttachedCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace of the kubernetes secret containing the proxy config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_attached_cluster#namespace GoogleContainerAttachedCluster#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

