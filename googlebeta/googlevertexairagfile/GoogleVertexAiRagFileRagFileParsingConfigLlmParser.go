// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagfile


type GoogleVertexAiRagFileRagFileParsingConfigLlmParser struct {
	// The prompt to use for parsing. If not specified, a default prompt is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_rag_file#custom_parsing_prompt GoogleVertexAiRagFile#custom_parsing_prompt}
	CustomParsingPrompt *string `field:"optional" json:"customParsingPrompt" yaml:"customParsingPrompt"`
	// The maximum number of requests the job is allowed to make to the LLM model per minute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_rag_file#max_parsing_requests_per_min GoogleVertexAiRagFile#max_parsing_requests_per_min}
	MaxParsingRequestsPerMin *float64 `field:"optional" json:"maxParsingRequestsPerMin" yaml:"maxParsingRequestsPerMin"`
	// The name of a LLM model used for parsing. Format: 'projects/{project_id}/locations/{location}/publishers/{publisher}/models/{model}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_rag_file#model_name GoogleVertexAiRagFile#model_name}
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
}

