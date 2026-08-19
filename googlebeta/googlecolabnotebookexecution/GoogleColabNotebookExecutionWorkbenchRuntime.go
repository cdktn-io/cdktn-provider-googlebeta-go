// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabnotebookexecution


type GoogleColabNotebookExecutionWorkbenchRuntime struct {
	// vm_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_colab_notebook_execution#vm_image GoogleColabNotebookExecution#vm_image}
	VmImage *GoogleColabNotebookExecutionWorkbenchRuntimeVmImage `field:"required" json:"vmImage" yaml:"vmImage"`
}

