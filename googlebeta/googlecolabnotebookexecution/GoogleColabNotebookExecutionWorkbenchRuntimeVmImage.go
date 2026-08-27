// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabnotebookexecution


type GoogleColabNotebookExecutionWorkbenchRuntimeVmImage struct {
	// Use this VM image family to find the image; the newest image in this family will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_colab_notebook_execution#family GoogleColabNotebookExecution#family}
	Family *string `field:"optional" json:"family" yaml:"family"`
	// Use VM image name to find the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_colab_notebook_execution#name GoogleColabNotebookExecution#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the Google Cloud project that this VM image belongs to. Format: {project_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_colab_notebook_execution#project GoogleColabNotebookExecution#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

