// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGetHttpHeaders struct {
	// Required. The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_cloud_run_v2_worker_pool#port GoogleCloudRunV2WorkerPool#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Optional. The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_cloud_run_v2_worker_pool#value GoogleCloudRunV2WorkerPool#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

