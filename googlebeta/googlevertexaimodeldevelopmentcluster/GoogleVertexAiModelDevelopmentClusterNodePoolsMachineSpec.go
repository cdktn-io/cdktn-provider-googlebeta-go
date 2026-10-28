// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpec struct {
	// The type of the machine.
	//
	// See the [list of machine types supported for
	// prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	// See the [list of machine types supported for custom
	// training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	// For DeployedModel this field is optional, and the default
	// value is 'n1-standard-2'. For BatchPredictionJob or as part of
	// WorkerPoolSpec this field is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#machine_type GoogleVertexAiModelDevelopmentCluster#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#reservation_affinity GoogleVertexAiModelDevelopmentCluster#reservation_affinity}
	ReservationAffinity *GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpecReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
}

