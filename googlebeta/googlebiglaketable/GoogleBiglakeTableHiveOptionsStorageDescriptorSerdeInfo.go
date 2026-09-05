// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglaketable


type GoogleBiglakeTableHiveOptionsStorageDescriptorSerdeInfo struct {
	// The fully qualified Java class name of the serialization library.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_table#serialization_lib GoogleBiglakeTable#serialization_lib}
	SerializationLib *string `field:"optional" json:"serializationLib" yaml:"serializationLib"`
}

