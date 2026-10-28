// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2instance


type GoogleCloudRunV2InstanceContainersLivenessProbeHttpGetHttpHeaders struct {
	// Required. The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_cloud_run_v2_instance#name GoogleCloudRunV2Instance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional. The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_cloud_run_v2_instance#value GoogleCloudRunV2Instance#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

