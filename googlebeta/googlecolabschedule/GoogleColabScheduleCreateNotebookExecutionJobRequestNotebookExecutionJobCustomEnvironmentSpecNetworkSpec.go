// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabschedule


type GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec struct {
	// Whether to enable public internet access. Default false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_colab_schedule#enable_internet_access GoogleColabSchedule#enable_internet_access}
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_colab_schedule#network GoogleColabSchedule#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The name of the subnet that this instance is in. Format: 'projects/{project_id_or_number}/regions/{region}/subnetworks/{subnetwork_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_colab_schedule#subnetwork GoogleColabSchedule#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

