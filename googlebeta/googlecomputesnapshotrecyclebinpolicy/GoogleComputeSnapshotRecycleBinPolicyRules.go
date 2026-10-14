// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputesnapshotrecyclebinpolicy


type GoogleComputeSnapshotRecycleBinPolicyRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_snapshot_recycle_bin_policy#name GoogleComputeSnapshotRecycleBinPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// standard_snapshots block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_snapshot_recycle_bin_policy#standard_snapshots GoogleComputeSnapshotRecycleBinPolicy#standard_snapshots}
	StandardSnapshots *GoogleComputeSnapshotRecycleBinPolicyRulesStandardSnapshots `field:"optional" json:"standardSnapshots" yaml:"standardSnapshots"`
}

