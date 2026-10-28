// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryreservationassignment


type GoogleBigqueryReservationAssignmentSchedulingPolicy struct {
	// If present and > 0, the reservation will attempt to limit the concurrency of jobs running for any particular project within it to the given value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_bigquery_reservation_assignment#concurrency GoogleBigqueryReservationAssignment#concurrency}
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// If present and > 0, the reservation will attempt to limit the slot consumption of queries running for any particular project within it to the given value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_bigquery_reservation_assignment#max_slots GoogleBigqueryReservationAssignment#max_slots}
	MaxSlots *float64 `field:"optional" json:"maxSlots" yaml:"maxSlots"`
}

