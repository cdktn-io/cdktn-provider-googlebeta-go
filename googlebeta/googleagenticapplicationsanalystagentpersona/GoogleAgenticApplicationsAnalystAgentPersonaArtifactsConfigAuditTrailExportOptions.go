// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigAuditTrailExportOptions struct {
	// If true, append the audit trail to the final response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_agentic_applications_analyst_agent_persona#append_audit_trail GoogleAgenticApplicationsAnalystAgentPersona#append_audit_trail}
	AppendAuditTrail interface{} `field:"optional" json:"appendAuditTrail" yaml:"appendAuditTrail"`
	// If true, export the audit trail as a separate artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_agentic_applications_analyst_agent_persona#export_audit_trail_artifact GoogleAgenticApplicationsAnalystAgentPersona#export_audit_trail_artifact}
	ExportAuditTrailArtifact interface{} `field:"optional" json:"exportAuditTrailArtifact" yaml:"exportAuditTrailArtifact"`
	// Format for audit trail export. Possible values: MARKDOWN HTML PDF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_agentic_applications_analyst_agent_persona#export_format GoogleAgenticApplicationsAnalystAgentPersona#export_format}
	ExportFormat *string `field:"optional" json:"exportFormat" yaml:"exportFormat"`
}

