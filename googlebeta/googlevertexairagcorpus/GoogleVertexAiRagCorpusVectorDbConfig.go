// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagcorpus


type GoogleVertexAiRagCorpusVectorDbConfig struct {
	// api_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#api_auth GoogleVertexAiRagCorpus#api_auth}
	ApiAuth *GoogleVertexAiRagCorpusVectorDbConfigApiAuth `field:"optional" json:"apiAuth" yaml:"apiAuth"`
	// pinecone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#pinecone GoogleVertexAiRagCorpus#pinecone}
	Pinecone *GoogleVertexAiRagCorpusVectorDbConfigPinecone `field:"optional" json:"pinecone" yaml:"pinecone"`
	// rag_embedding_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#rag_embedding_model_config GoogleVertexAiRagCorpus#rag_embedding_model_config}
	RagEmbeddingModelConfig *GoogleVertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig `field:"optional" json:"ragEmbeddingModelConfig" yaml:"ragEmbeddingModelConfig"`
	// rag_managed_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#rag_managed_db GoogleVertexAiRagCorpus#rag_managed_db}
	RagManagedDb *GoogleVertexAiRagCorpusVectorDbConfigRagManagedDb `field:"optional" json:"ragManagedDb" yaml:"ragManagedDb"`
	// vertex_vector_search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#vertex_vector_search GoogleVertexAiRagCorpus#vertex_vector_search}
	VertexVectorSearch *GoogleVertexAiRagCorpusVectorDbConfigVertexVectorSearch `field:"optional" json:"vertexVectorSearch" yaml:"vertexVectorSearch"`
}

