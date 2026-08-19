// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicledataexport


type GoogleChronicleDataExportIngestionLabels struct {
	// The key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_data_export#key GoogleChronicleDataExport#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_data_export#value GoogleChronicleDataExport#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

