// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevectorsearchindex


type GoogleVectorSearchIndexDedicatedInfrastructureAutoscalingSpec struct {
	// The maximum number of replicas.
	//
	// Must be >= 'min_replica_count'
	// and <= '1000'. If not set or set to '0', defaults to the greater
	// of 'min_replica_count' and '2' (or '5' for the v1beta version).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_index#max_replica_count GoogleVectorSearchIndex#max_replica_count}
	MaxReplicaCount *float64 `field:"optional" json:"maxReplicaCount" yaml:"maxReplicaCount"`
	// The minimum number of replicas.
	//
	// If not set or set to '0', defaults
	// to '2'. Must be >= '1' and <= '1000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_index#min_replica_count GoogleVectorSearchIndex#min_replica_count}
	MinReplicaCount *float64 `field:"optional" json:"minReplicaCount" yaml:"minReplicaCount"`
}

