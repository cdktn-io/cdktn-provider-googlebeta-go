// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterSecurityConfigAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_gkeonprem_bare_metal_cluster#admin_users GoogleGkeonpremBareMetalCluster#admin_users}
	AdminUsers interface{} `field:"required" json:"adminUsers" yaml:"adminUsers"`
}

