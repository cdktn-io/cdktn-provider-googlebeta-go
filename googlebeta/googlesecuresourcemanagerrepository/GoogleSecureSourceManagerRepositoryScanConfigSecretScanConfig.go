// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesecuresourcemanagerrepository


type GoogleSecureSourceManagerRepositoryScanConfigSecretScanConfig struct {
	// Enables secret scanning for the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_secure_source_manager_repository#enabled GoogleSecureSourceManagerRepository#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The DLP inspect template to use for secret scanning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_secure_source_manager_repository#inspect_template GoogleSecureSourceManagerRepository#inspect_template}
	InspectTemplate *string `field:"optional" json:"inspectTemplate" yaml:"inspectTemplate"`
}

