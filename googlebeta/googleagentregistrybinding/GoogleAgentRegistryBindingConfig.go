// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentregistrybinding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleAgentRegistryBindingConfig struct {
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
	// auth_provider_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#auth_provider_binding GoogleAgentRegistryBinding#auth_provider_binding}
	AuthProviderBinding *GoogleAgentRegistryBindingAuthProviderBinding `field:"required" json:"authProviderBinding" yaml:"authProviderBinding"`
	// The name of the Binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#binding_id GoogleAgentRegistryBinding#binding_id}
	BindingId *string `field:"required" json:"bindingId" yaml:"bindingId"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#location GoogleAgentRegistryBinding#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#source GoogleAgentRegistryBinding#source}
	Source *GoogleAgentRegistryBindingSource `field:"required" json:"source" yaml:"source"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#target GoogleAgentRegistryBinding#target}
	Target *GoogleAgentRegistryBindingTarget `field:"required" json:"target" yaml:"target"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#deletion_policy GoogleAgentRegistryBinding#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The description of the Binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#description GoogleAgentRegistryBinding#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-defined display name for the Binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#display_name GoogleAgentRegistryBinding#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#id GoogleAgentRegistryBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#project GoogleAgentRegistryBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_binding#timeouts GoogleAgentRegistryBinding#timeouts}
	Timeouts *GoogleAgentRegistryBindingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

