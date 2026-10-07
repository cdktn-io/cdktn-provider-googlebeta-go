// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListStruct struct {
	// disk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#disk_config GoogleDataprocWorkflowTemplate#disk_config}
	DiskConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig `field:"optional" json:"diskConfig" yaml:"diskConfig"`
	// Optional. Full machine-type names, e.g. "n1-standard-16".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#machine_types GoogleDataprocWorkflowTemplate#machine_types}
	MachineTypes *[]*string `field:"optional" json:"machineTypes" yaml:"machineTypes"`
	// Optional.
	//
	// Preference of this instance selection. Lower number means higher preference. Dataproc will first try to create a VM based on the machine-type with priority rank and fallback to next rank based on availability. Machine types and instance selections with the same priority have the same preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#rank GoogleDataprocWorkflowTemplate#rank}
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

