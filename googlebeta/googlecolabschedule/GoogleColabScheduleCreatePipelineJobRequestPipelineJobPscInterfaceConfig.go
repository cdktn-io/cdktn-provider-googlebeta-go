// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabschedule


type GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig struct {
	// dns_peering_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_colab_schedule#dns_peering_configs GoogleColabSchedule#dns_peering_configs}
	DnsPeeringConfigs interface{} `field:"optional" json:"dnsPeeringConfigs" yaml:"dnsPeeringConfigs"`
	// The name of the Compute Engine [network attachment](https://cloud.google.com/vpc/docs/about-network-attachments) to attach to the resource within the region and user project. To specify this field, you must have already [created a network attachment] (https://cloud.google.com/vpc/docs/create-manage-network-attachments#create-network-attachments). This field is only used for resources using PSC-I.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_colab_schedule#network_attachment GoogleColabSchedule#network_attachment}
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
}

