// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagcorpus


type GoogleVertexAiRagCorpusVectorDbConfigRagManagedDb struct {
	// ann block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#ann GoogleVertexAiRagCorpus#ann}
	Ann *GoogleVertexAiRagCorpusVectorDbConfigRagManagedDbAnn `field:"optional" json:"ann" yaml:"ann"`
	// knn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#knn GoogleVertexAiRagCorpus#knn}
	Knn *GoogleVertexAiRagCorpusVectorDbConfigRagManagedDbKnn `field:"optional" json:"knn" yaml:"knn"`
}

