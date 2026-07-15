// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule


type GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpecReservationAffinity struct {
	// Specifies the reservation affinity type. Possible values: NO_RESERVATION ANY_RESERVATION SPECIFIC_RESERVATION SPECIFIC_THEN_ANY_RESERVATION SPECIFIC_THEN_NO_RESERVATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_schedule#reservation_affinity_type GoogleVertexAiSchedule#reservation_affinity_type}
	ReservationAffinityType *string `field:"required" json:"reservationAffinityType" yaml:"reservationAffinityType"`
	// Corresponds to the label key of a reservation resource.
	//
	// To target a SPECIFIC_RESERVATION by name, use 'compute.googleapis.com/reservation-name' as the key and specify the name of your reservation as its value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_schedule#key GoogleVertexAiSchedule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// When set to true, resources will be drawn from go/cloud-ai-gcp-pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_schedule#use_reservation_pool GoogleVertexAiSchedule#use_reservation_pool}
	UseReservationPool interface{} `field:"optional" json:"useReservationPool" yaml:"useReservationPool"`
	// Corresponds to the label values of a reservation resource.
	//
	// This must be the full resource name of the reservation or reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_schedule#values GoogleVertexAiSchedule#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

