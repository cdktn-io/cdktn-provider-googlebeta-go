// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestoolset


type GoogleCesToolsetMcpToolsetTlsConfig struct {
	// ca_certs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_toolset#ca_certs GoogleCesToolset#ca_certs}
	CaCerts interface{} `field:"required" json:"caCerts" yaml:"caCerts"`
}

