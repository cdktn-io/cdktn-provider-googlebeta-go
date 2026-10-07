// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagcorpus


type GoogleVertexAiRagCorpusVectorDbConfigPinecone struct {
	// Pinecone index name. This value cannot be changed after it's set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#index_name GoogleVertexAiRagCorpus#index_name}
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
}

