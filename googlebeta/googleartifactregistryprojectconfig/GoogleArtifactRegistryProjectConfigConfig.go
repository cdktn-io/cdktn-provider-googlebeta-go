// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleartifactregistryprojectconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleArtifactRegistryProjectConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_artifact_registry_project_config#id GoogleArtifactRegistryProjectConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the location this config is located in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_artifact_registry_project_config#location GoogleArtifactRegistryProjectConfig#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// platform_logs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_artifact_registry_project_config#platform_logs_config GoogleArtifactRegistryProjectConfig#platform_logs_config}
	PlatformLogsConfig *GoogleArtifactRegistryProjectConfigPlatformLogsConfig `field:"optional" json:"platformLogsConfig" yaml:"platformLogsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_artifact_registry_project_config#project GoogleArtifactRegistryProjectConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_artifact_registry_project_config#timeouts GoogleArtifactRegistryProjectConfig#timeouts}
	Timeouts *GoogleArtifactRegistryProjectConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

