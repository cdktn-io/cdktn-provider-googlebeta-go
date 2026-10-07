// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagcorpus


type GoogleVertexAiRagCorpusVectorDbConfigRagManagedDbAnn struct {
	// Number of leaf nodes in the tree-based structure. Default value is 500.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#leaf_count GoogleVertexAiRagCorpus#leaf_count}
	LeafCount *float64 `field:"optional" json:"leafCount" yaml:"leafCount"`
	// The depth of the tree-based structure. Only depth values of 2 and 3 are supported. Default value is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_rag_corpus#tree_depth GoogleVertexAiRagCorpus#tree_depth}
	TreeDepth *float64 `field:"optional" json:"treeDepth" yaml:"treeDepth"`
}

