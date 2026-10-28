// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabnotebookexecution


type GoogleColabNotebookExecutionWorkbenchRuntime struct {
	// custom_container_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_colab_notebook_execution#custom_container_image GoogleColabNotebookExecution#custom_container_image}
	CustomContainerImage *GoogleColabNotebookExecutionWorkbenchRuntimeCustomContainerImage `field:"optional" json:"customContainerImage" yaml:"customContainerImage"`
	// vm_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_colab_notebook_execution#vm_image GoogleColabNotebookExecution#vm_image}
	VmImage *GoogleColabNotebookExecutionWorkbenchRuntimeVmImage `field:"optional" json:"vmImage" yaml:"vmImage"`
}

