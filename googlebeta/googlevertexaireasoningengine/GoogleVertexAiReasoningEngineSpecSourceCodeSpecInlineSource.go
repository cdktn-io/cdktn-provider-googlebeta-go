// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSource struct {
	// Required. Input only. The application source code archive, provided as a compressed tarball (.tar.gz) file. A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#source_archive GoogleVertexAiReasoningEngine#source_archive}
	SourceArchive *string `field:"optional" json:"sourceArchive" yaml:"sourceArchive"`
}

