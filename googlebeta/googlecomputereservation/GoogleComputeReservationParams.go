// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputereservation


type GoogleComputeReservationParams struct {
	// Resource manager tags to be bound to the reservation.
	//
	// Tag keys and values have the
	// same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id},
	// and values are in the format tagValues/456.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_compute_reservation#resource_manager_tags GoogleComputeReservation#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

