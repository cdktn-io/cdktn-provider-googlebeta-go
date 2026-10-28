// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputesnapshotrecyclebinpolicy


type GoogleComputeSnapshotRecycleBinPolicyRulesStandardSnapshots struct {
	// The retention duration for snapshots in the recycle bin after which the snapshots are automatically deleted from recycle bin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_compute_snapshot_recycle_bin_policy#retention_duration_days GoogleComputeSnapshotRecycleBinPolicy#retention_duration_days}
	RetentionDurationDays *float64 `field:"optional" json:"retentionDurationDays" yaml:"retentionDurationDays"`
}

