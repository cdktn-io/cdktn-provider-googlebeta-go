// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabschedule


type GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_colab_schedule#machine_spec GoogleColabSchedule#machine_spec}
	MachineSpec *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec `field:"optional" json:"machineSpec" yaml:"machineSpec"`
	// network_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_colab_schedule#network_spec GoogleColabSchedule#network_spec}
	NetworkSpec *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec `field:"optional" json:"networkSpec" yaml:"networkSpec"`
	// persistent_disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_colab_schedule#persistent_disk_spec GoogleColabSchedule#persistent_disk_spec}
	PersistentDiskSpec *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec `field:"optional" json:"persistentDiskSpec" yaml:"persistentDiskSpec"`
}

