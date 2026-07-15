// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleconfigdeployment


type GoogleConfigDeploymentTerraformBlueprintGitSource struct {
	// Repository URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_config_deployment#repo GoogleConfigDeployment#repo}
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// Subdirectory within the repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_config_deployment#directory GoogleConfigDeployment#directory}
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Git reference (branch or tag).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_config_deployment#ref GoogleConfigDeployment#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

