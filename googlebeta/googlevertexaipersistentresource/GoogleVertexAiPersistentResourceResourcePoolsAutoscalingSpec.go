// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource


type GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec struct {
	// max replicas in the node pool, must be ≥ replica_count and > min_replica_count or will throw error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#max_replica_count GoogleVertexAiPersistentResource#max_replica_count}
	MaxReplicaCount *string `field:"optional" json:"maxReplicaCount" yaml:"maxReplicaCount"`
	// min replicas in the node pool, must be ≤ replica_count and < max_replica_count or will throw error.
	//
	// For autoscaling enabled Ray-on-Vertex, we allow min_replica_count of a
	// resource_pool to be 0 to match the OSS Ray
	// behavior(https://docs.ray.io/en/latest/cluster/vms/user-guides/configuring-autoscaling.html#cluster-config-parameters).
	// As for Persistent Resource, the min_replica_count must be > 0, we added
	// a corresponding validation inside
	// CreatePersistentResourceRequestValidator.java.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#min_replica_count GoogleVertexAiPersistentResource#min_replica_count}
	MinReplicaCount *string `field:"optional" json:"minReplicaCount" yaml:"minReplicaCount"`
}

