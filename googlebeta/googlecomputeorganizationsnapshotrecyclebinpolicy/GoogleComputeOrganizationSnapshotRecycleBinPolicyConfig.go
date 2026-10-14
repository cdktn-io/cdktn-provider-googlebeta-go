// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeorganizationsnapshotrecyclebinpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleComputeOrganizationSnapshotRecycleBinPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The organization ID for this request (e.g., "organizations/1234").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_organization_snapshot_recycle_bin_policy#organization GoogleComputeOrganizationSnapshotRecycleBinPolicy#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_organization_snapshot_recycle_bin_policy#id GoogleComputeOrganizationSnapshotRecycleBinPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_organization_snapshot_recycle_bin_policy#rules GoogleComputeOrganizationSnapshotRecycleBinPolicy#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_organization_snapshot_recycle_bin_policy#timeouts GoogleComputeOrganizationSnapshotRecycleBinPolicy#timeouts}
	Timeouts *GoogleComputeOrganizationSnapshotRecycleBinPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

