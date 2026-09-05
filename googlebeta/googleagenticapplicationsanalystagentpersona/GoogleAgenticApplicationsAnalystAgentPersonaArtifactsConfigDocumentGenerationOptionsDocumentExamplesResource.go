// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResource struct {
	// bigquery_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#bigquery_resource GoogleAgenticApplicationsAnalystAgentPersona#bigquery_resource}
	BigqueryResource *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResource `field:"optional" json:"bigqueryResource" yaml:"bigqueryResource"`
	// A user-friendly name for this resource. This can be shown to the user and used by the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#display_label GoogleAgenticApplicationsAnalystAgentPersona#display_label}
	DisplayLabel *string `field:"optional" json:"displayLabel" yaml:"displayLabel"`
	// f1_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#f1_resource GoogleAgenticApplicationsAnalystAgentPersona#f1_resource}
	F1Resource *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1Resource `field:"optional" json:"f1Resource" yaml:"f1Resource"`
	// google_cloud_storage_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#google_cloud_storage_resource GoogleAgenticApplicationsAnalystAgentPersona#google_cloud_storage_resource}
	GoogleCloudStorageResource *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResource `field:"optional" json:"googleCloudStorageResource" yaml:"googleCloudStorageResource"`
	// google_drive_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#google_drive_resource GoogleAgenticApplicationsAnalystAgentPersona#google_drive_resource}
	GoogleDriveResource *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResource `field:"optional" json:"googleDriveResource" yaml:"googleDriveResource"`
	// A description of the resource. The model may use this, it will not be shown to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#model_description GoogleAgenticApplicationsAnalystAgentPersona#model_description}
	ModelDescription *string `field:"optional" json:"modelDescription" yaml:"modelDescription"`
	// raw_file_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#raw_file_resource GoogleAgenticApplicationsAnalystAgentPersona#raw_file_resource}
	RawFileResource *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResource `field:"optional" json:"rawFileResource" yaml:"rawFileResource"`
	// If true, use RAG to retrieve relevant information from the resources.
	//
	// Must only be set for file-based resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#use_rag GoogleAgenticApplicationsAnalystAgentPersona#use_rag}
	UseRag interface{} `field:"optional" json:"useRag" yaml:"useRag"`
}

