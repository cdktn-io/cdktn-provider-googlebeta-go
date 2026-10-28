// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagfile


type GoogleVertexAiRagFileGoogleDriveSourceResourceIds struct {
	// The ID of the Google Drive file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_rag_file#resource_id GoogleVertexAiRagFile#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The type of the Google Drive resource.
	//
	// Only individual files are
	// supported because each resource manages one RagFile. Default value: "RESOURCE_TYPE_FILE" Possible values: ["RESOURCE_TYPE_FILE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_rag_file#resource_type GoogleVertexAiRagFile#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

