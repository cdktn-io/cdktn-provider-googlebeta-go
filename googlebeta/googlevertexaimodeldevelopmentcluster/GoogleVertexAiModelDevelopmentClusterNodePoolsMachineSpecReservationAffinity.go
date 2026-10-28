// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpecReservationAffinity struct {
	// Specifies the reservation affinity type. Possible values: NO_RESERVATION ANY_RESERVATION SPECIFIC_RESERVATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#reservation_affinity_type GoogleVertexAiModelDevelopmentCluster#reservation_affinity_type}
	ReservationAffinityType *string `field:"required" json:"reservationAffinityType" yaml:"reservationAffinityType"`
	// Corresponds to the label values of a reservation resource.
	//
	// This must be the
	// resource name of the reservation, reservation block, or reservation sub-
	// block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#values GoogleVertexAiModelDevelopmentCluster#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

