// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevectorsearchdataobject


type GoogleVectorSearchDataObjectVectorsSparse struct {
	// The indices corresponding to the entries in 'values'. Must have the same length as 'values'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vector_search_data_object#indices GoogleVectorSearchDataObject#indices}
	Indices *[]*float64 `field:"required" json:"indices" yaml:"indices"`
	// The non-zero float values of the sparse vector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vector_search_data_object#values GoogleVectorSearchDataObject#values}
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

