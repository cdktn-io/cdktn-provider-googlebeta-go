// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesecuresourcemanagerrepository


type GoogleSecureSourceManagerRepositoryScanConfig struct {
	// secret_scan_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_secure_source_manager_repository#secret_scan_config GoogleSecureSourceManagerRepository#secret_scan_config}
	SecretScanConfig *GoogleSecureSourceManagerRepositoryScanConfigSecretScanConfig `field:"optional" json:"secretScanConfig" yaml:"secretScanConfig"`
}

