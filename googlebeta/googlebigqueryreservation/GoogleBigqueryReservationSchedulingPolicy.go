// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryreservation


type GoogleBigqueryReservationSchedulingPolicy struct {
	// If present and > 0, the reservation will attempt to limit the concurrency of jobs running for any particular project within it to the given value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_bigquery_reservation#concurrency GoogleBigqueryReservation#concurrency}
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// If present and > 0, the reservation will attempt to limit the slot consumption of queries running for any particular project within it to the given value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_bigquery_reservation#max_slots GoogleBigqueryReservation#max_slots}
	MaxSlots *float64 `field:"optional" json:"maxSlots" yaml:"maxSlots"`
}

