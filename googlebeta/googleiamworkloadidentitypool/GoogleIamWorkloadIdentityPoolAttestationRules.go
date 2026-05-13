// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamworkloadidentitypool


type GoogleIamWorkloadIdentityPoolAttestationRules struct {
	// A single workload operating on Google Cloud. For example: '//run.googleapis.com/projects/123/type/Service/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_iam_workload_identity_pool#google_cloud_resource GoogleIamWorkloadIdentityPool#google_cloud_resource}
	GoogleCloudResource *string `field:"required" json:"googleCloudResource" yaml:"googleCloudResource"`
}

