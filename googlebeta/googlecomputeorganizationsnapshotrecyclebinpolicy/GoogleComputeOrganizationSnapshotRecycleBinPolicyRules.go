// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeorganizationsnapshotrecyclebinpolicy


type GoogleComputeOrganizationSnapshotRecycleBinPolicyRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_organization_snapshot_recycle_bin_policy#name GoogleComputeOrganizationSnapshotRecycleBinPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// standard_snapshots block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_organization_snapshot_recycle_bin_policy#standard_snapshots GoogleComputeOrganizationSnapshotRecycleBinPolicy#standard_snapshots}
	StandardSnapshots *GoogleComputeOrganizationSnapshotRecycleBinPolicyRulesStandardSnapshots `field:"optional" json:"standardSnapshots" yaml:"standardSnapshots"`
}

