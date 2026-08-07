// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule


type GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource struct {
	// The version of the Cloud Storage object to read.
	//
	// If unset, the current version of the object is read. See https://cloud.google.com/storage/docs/metadata#generation-number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#generation GoogleVertexAiSchedule#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
	// The Cloud Storage uri pointing to the ipynb file. Format: 'gs://bucket/notebook_file.ipynb'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_schedule#uri GoogleVertexAiSchedule#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

