// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemanagedkafkacluster


type GoogleManagedKafkaClusterGcpConfigAccessConfig struct {
	// network_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_managed_kafka_cluster#network_configs GoogleManagedKafkaCluster#network_configs}
	NetworkConfigs interface{} `field:"required" json:"networkConfigs" yaml:"networkConfigs"`
	// public_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_managed_kafka_cluster#public_cluster_config GoogleManagedKafkaCluster#public_cluster_config}
	PublicClusterConfig *GoogleManagedKafkaClusterGcpConfigAccessConfigPublicClusterConfig `field:"optional" json:"publicClusterConfig" yaml:"publicClusterConfig"`
}

