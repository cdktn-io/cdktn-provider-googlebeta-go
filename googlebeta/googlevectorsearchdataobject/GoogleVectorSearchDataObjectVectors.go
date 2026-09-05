// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevectorsearchdataobject


type GoogleVectorSearchDataObjectVectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#field_name GoogleVectorSearchDataObject#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// dense block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#dense GoogleVectorSearchDataObject#dense}
	Dense *GoogleVectorSearchDataObjectVectorsDense `field:"optional" json:"dense" yaml:"dense"`
	// sparse block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#sparse GoogleVectorSearchDataObject#sparse}
	Sparse *GoogleVectorSearchDataObjectVectorsSparse `field:"optional" json:"sparse" yaml:"sparse"`
}

