// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool


type GoogleCesToolFileSearchTool struct {
	// Required. The tool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_ces_tool#name GoogleCesTool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional. The type of the corpus. Default is FULLY_MANAGED. Possible values: CORPUS_TYPE_UNSPECIFIED USER_OWNED FULLY_MANAGED Possible values: ["CORPUS_TYPE_UNSPECIFIED", "USER_OWNED", "FULLY_MANAGED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_ces_tool#corpus_type GoogleCesTool#corpus_type}
	CorpusType *string `field:"optional" json:"corpusType" yaml:"corpusType"`
	// Optional. The tool description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_ces_tool#description GoogleCesTool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional. The corpus where files are stored. Format: projects/{project}/locations/{location}/ragCorpora/{rag_corpus}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_ces_tool#file_corpus GoogleCesTool#file_corpus}
	FileCorpus *string `field:"optional" json:"fileCorpus" yaml:"fileCorpus"`
}

