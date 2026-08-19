// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentidentityauthprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleAgentIdentityAuthProviderConfig struct {
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
	// The ID to use for the AuthProvider, which will become the final segment of the AuthProvider's resource name.
	//
	// This value should be 1-63 characters, and valid characters
	// are /a-z-/. The first character must be a lowercase letter, and the
	// last character must be a lowercase letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#auth_provider_id GoogleAgentIdentityAuthProvider#auth_provider_id}
	AuthProviderId *string `field:"required" json:"authProviderId" yaml:"authProviderId"`
	// auth_provider_type_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#auth_provider_type_params GoogleAgentIdentityAuthProvider#auth_provider_type_params}
	AuthProviderTypeParams *GoogleAgentIdentityAuthProviderAuthProviderTypeParams `field:"required" json:"authProviderTypeParams" yaml:"authProviderTypeParams"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#location GoogleAgentIdentityAuthProvider#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// List of scopes that are allowed to be requested for this auth_provider.
	//
	// If this list is non-empty, only scopes within this list may be requested.
	// If this list is empty, all scopes may be requested.
	// Scopes appearing in 'blocked_scopes' are disallowed even if they appear in
	// 'allowed_scopes'.
	// The number of allowed scopes is limited to 200.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#allowed_scopes GoogleAgentIdentityAuthProvider#allowed_scopes}
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// List of scopes that are blocked from being requested for this auth_provider.
	//
	// If a scope appears in this list, it will not be requested,
	// even if it also appears in 'allowed_scopes'. 'blocked_scopes' takes
	// precedence over 'allowed_scopes'. The number of blocked scopes is limited
	// to 200.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#blocked_scopes GoogleAgentIdentityAuthProvider#blocked_scopes}
	BlockedScopes *[]*string `field:"optional" json:"blockedScopes" yaml:"blockedScopes"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#deletion_policy GoogleAgentIdentityAuthProvider#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Description of the resource. Must be less than 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#description GoogleAgentIdentityAuthProvider#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#id GoogleAgentIdentityAuthProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#labels GoogleAgentIdentityAuthProvider#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#project GoogleAgentIdentityAuthProvider#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#timeouts GoogleAgentIdentityAuthProvider#timeouts}
	Timeouts *GoogleAgentIdentityAuthProviderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Input only. Represents the workload identity in IAM 'principal://' format of the agent(s) that will use this AuthProvider. Example: 'principal://agents.global.org-${ORG_ID}.system.id.goog/resources/aiplatform/projects/{PROJECT_ID}/locations/{LOCATIONS}/reasoningEngines/{ID}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agent_identity_auth_provider#workload_ids GoogleAgentIdentityAuthProvider#workload_ids}
	WorkloadIds *[]*string `field:"optional" json:"workloadIds" yaml:"workloadIds"`
}

