// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleconfigdeployment


type GoogleConfigDeploymentTerraformBlueprint struct {
	// URI of a GCS object containing the zipped Terraform blueprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_config_deployment#gcs_source GoogleConfigDeployment#gcs_source}
	GcsSource *string `field:"optional" json:"gcsSource" yaml:"gcsSource"`
	// git_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_config_deployment#git_source GoogleConfigDeployment#git_source}
	GitSource *GoogleConfigDeploymentTerraformBlueprintGitSource `field:"optional" json:"gitSource" yaml:"gitSource"`
	// input_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_config_deployment#input_values GoogleConfigDeployment#input_values}
	InputValues interface{} `field:"optional" json:"inputValues" yaml:"inputValues"`
}

