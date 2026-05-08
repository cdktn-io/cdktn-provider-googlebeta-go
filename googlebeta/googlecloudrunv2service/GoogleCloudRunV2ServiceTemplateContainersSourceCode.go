// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersSourceCode struct {
	// cloud_storage_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#cloud_storage_source GoogleCloudRunV2Service#cloud_storage_source}
	CloudStorageSource *GoogleCloudRunV2ServiceTemplateContainersSourceCodeCloudStorageSource `field:"optional" json:"cloudStorageSource" yaml:"cloudStorageSource"`
}

