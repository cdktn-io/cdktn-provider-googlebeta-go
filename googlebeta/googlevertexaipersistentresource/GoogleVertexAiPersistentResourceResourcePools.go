// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource


type GoogleVertexAiPersistentResourceResourcePools struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_persistent_resource#machine_spec GoogleVertexAiPersistentResource#machine_spec}
	MachineSpec *GoogleVertexAiPersistentResourceResourcePoolsMachineSpec `field:"required" json:"machineSpec" yaml:"machineSpec"`
	// autoscaling_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_persistent_resource#autoscaling_spec GoogleVertexAiPersistentResource#autoscaling_spec}
	AutoscalingSpec *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec `field:"optional" json:"autoscalingSpec" yaml:"autoscalingSpec"`
	// disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_persistent_resource#disk_spec GoogleVertexAiPersistentResource#disk_spec}
	DiskSpec *GoogleVertexAiPersistentResourceResourcePoolsDiskSpec `field:"optional" json:"diskSpec" yaml:"diskSpec"`
	// The unique ID in a PersistentResource for referring to this resource pool.
	//
	// User can specify it if necessary. Otherwise, it's generated
	// automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_persistent_resource#id GoogleVertexAiPersistentResource#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The total number of machines to use for this resource pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_persistent_resource#replica_count GoogleVertexAiPersistentResource#replica_count}
	ReplicaCount *string `field:"optional" json:"replicaCount" yaml:"replicaCount"`
}

