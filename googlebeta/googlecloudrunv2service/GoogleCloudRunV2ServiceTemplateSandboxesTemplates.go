// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateSandboxesTemplates struct {
	// Name of the container image in Dockerhub or Artifact Registry. If the host is not provided, Dockerhub is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_cloud_run_v2_service#image GoogleCloudRunV2Service#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// Name of the sandbox specified as a DNS_LABEL (RFC 1123).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_cloud_run_v2_service#name GoogleCloudRunV2Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_cloud_run_v2_service#args GoogleCloudRunV2Service#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_cloud_run_v2_service#command GoogleCloudRunV2Service#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_cloud_run_v2_service#env GoogleCloudRunV2Service#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// volume_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_cloud_run_v2_service#volume_mounts GoogleCloudRunV2Service#volume_mounts}
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	// Container's working directory.
	//
	// If not specified, the container runtime's default will be used, which might be configured in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_cloud_run_v2_service#working_dir GoogleCloudRunV2Service#working_dir}
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

