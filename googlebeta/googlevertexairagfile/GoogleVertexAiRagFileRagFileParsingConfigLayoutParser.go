// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagfile


type GoogleVertexAiRagFileRagFileParsingConfigLayoutParser struct {
	// The maximum number of requests the job is allowed to make to the Document AI processor per minute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_rag_file#max_parsing_requests_per_min GoogleVertexAiRagFile#max_parsing_requests_per_min}
	MaxParsingRequestsPerMin *float64 `field:"optional" json:"maxParsingRequestsPerMin" yaml:"maxParsingRequestsPerMin"`
	// The full resource name of a Document AI processor or processor version of type 'LAYOUT_PARSER_PROCESSOR'. Format: 'projects/{project_id}/locations/{location}/processors/{processor_id}' or 'projects/{project_id}/locations/{location}/processors/{processor_id}/processorVersions/{processor_version_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_rag_file#processor_name GoogleVertexAiRagFile#processor_name}
	ProcessorName *string `field:"optional" json:"processorName" yaml:"processorName"`
}

