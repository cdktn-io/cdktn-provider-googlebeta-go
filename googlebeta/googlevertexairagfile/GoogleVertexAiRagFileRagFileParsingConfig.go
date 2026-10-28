// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagfile


type GoogleVertexAiRagFileRagFileParsingConfig struct {
	// layout_parser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_rag_file#layout_parser GoogleVertexAiRagFile#layout_parser}
	LayoutParser *GoogleVertexAiRagFileRagFileParsingConfigLayoutParser `field:"optional" json:"layoutParser" yaml:"layoutParser"`
	// llm_parser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_rag_file#llm_parser GoogleVertexAiRagFile#llm_parser}
	LlmParser *GoogleVertexAiRagFileRagFileParsingConfigLlmParser `field:"optional" json:"llmParser" yaml:"llmParser"`
}

