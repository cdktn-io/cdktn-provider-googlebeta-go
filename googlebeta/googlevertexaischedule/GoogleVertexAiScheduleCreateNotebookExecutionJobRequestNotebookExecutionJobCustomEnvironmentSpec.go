// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule


type GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_schedule#machine_spec GoogleVertexAiSchedule#machine_spec}
	MachineSpec *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec `field:"optional" json:"machineSpec" yaml:"machineSpec"`
	// network_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_schedule#network_spec GoogleVertexAiSchedule#network_spec}
	NetworkSpec *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec `field:"optional" json:"networkSpec" yaml:"networkSpec"`
	// persistent_disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_schedule#persistent_disk_spec GoogleVertexAiSchedule#persistent_disk_spec}
	PersistentDiskSpec *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec `field:"optional" json:"persistentDiskSpec" yaml:"persistentDiskSpec"`
}

