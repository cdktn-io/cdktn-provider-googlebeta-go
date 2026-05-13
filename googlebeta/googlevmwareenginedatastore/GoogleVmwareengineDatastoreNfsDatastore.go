// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevmwareenginedatastore


type GoogleVmwareengineDatastoreNfsDatastore struct {
	// google_file_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vmwareengine_datastore#google_file_service GoogleVmwareengineDatastore#google_file_service}
	GoogleFileService *GoogleVmwareengineDatastoreNfsDatastoreGoogleFileService `field:"optional" json:"googleFileService" yaml:"googleFileService"`
	// third_party_file_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vmwareengine_datastore#third_party_file_service GoogleVmwareengineDatastore#third_party_file_service}
	ThirdPartyFileService *GoogleVmwareengineDatastoreNfsDatastoreThirdPartyFileService `field:"optional" json:"thirdPartyFileService" yaml:"thirdPartyFileService"`
}

