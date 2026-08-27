// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptions struct {
	// document_examples block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#document_examples GoogleAgenticApplicationsAnalystAgentPersona#document_examples}
	DocumentExamples interface{} `field:"optional" json:"documentExamples" yaml:"documentExamples"`
	// Format for document export. Possible values: PDF DOCX GOOGLE_DOCS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#export_format GoogleAgenticApplicationsAnalystAgentPersona#export_format}
	ExportFormat *string `field:"optional" json:"exportFormat" yaml:"exportFormat"`
}

