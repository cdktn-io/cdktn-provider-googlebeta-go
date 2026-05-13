// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegkehubrolloutsequence


type GoogleGkeHubRolloutSequenceIgnoredClustersSelector struct {
	// The label selector must be a valid CEL (Common Expression Language) expression which evaluates resource.labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_gke_hub_rollout_sequence#label_selector GoogleGkeHubRolloutSequence#label_selector}
	LabelSelector *string `field:"required" json:"labelSelector" yaml:"labelSelector"`
}

