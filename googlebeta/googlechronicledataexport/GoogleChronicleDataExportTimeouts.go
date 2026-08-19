// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicledataexport


type GoogleChronicleDataExportTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_data_export#create GoogleChronicleDataExport#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_data_export#delete GoogleChronicleDataExport#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

