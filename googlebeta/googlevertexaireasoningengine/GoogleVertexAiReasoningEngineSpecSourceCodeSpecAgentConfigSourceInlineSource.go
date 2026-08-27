// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSource struct {
	// Required. Input only. The application source code archive, provided as a compressed tarball (.tar.gz) file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#source_archive GoogleVertexAiReasoningEngine#source_archive}
	SourceArchive *string `field:"required" json:"sourceArchive" yaml:"sourceArchive"`
}

