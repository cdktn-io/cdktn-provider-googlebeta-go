// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiPersistentResourceConfig struct {
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
	// The ID to use for the PersistentResource, which become the final component of the PersistentResource's resource name.
	//
	// The maximum length is 63 characters, and valid characters
	// are '/^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#name GoogleVertexAiPersistentResource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// resource_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#resource_pools GoogleVertexAiPersistentResource#resource_pools}
	ResourcePools interface{} `field:"required" json:"resourcePools" yaml:"resourcePools"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#deletion_policy GoogleVertexAiPersistentResource#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The display name of the PersistentResource.
	//
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#display_name GoogleVertexAiPersistentResource#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#encryption_spec GoogleVertexAiPersistentResource#encryption_spec}
	EncryptionSpec *GoogleVertexAiPersistentResourceEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#id GoogleVertexAiPersistentResource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels with user-defined metadata to organize PersistentResource.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#labels GoogleVertexAiPersistentResource#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of the PersistentResource. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#location GoogleVertexAiPersistentResource#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The full name of the Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to peered with Vertex AI to host the persistent resources.
	//
	// For example, 'projects/12345/global/networks/myVPC'.
	// [Format](/compute/docs/reference/rest/v1/networks/insert)
	// is of the form 'projects/{project}/global/networks/{network}'.
	// Where {project} is a project number, as in '12345', and {network} is a
	// network name.
	//
	// To specify this field, you must have already [configured VPC Network
	// Peering for Vertex
	// AI](https://cloud.google.com/vertex-ai/docs/general/vpc-peering).
	//
	// If this field is left unspecified, the resources aren't peered with any
	// network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#network GoogleVertexAiPersistentResource#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#project GoogleVertexAiPersistentResource#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// psc_interface_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#psc_interface_config GoogleVertexAiPersistentResource#psc_interface_config}
	PscInterfaceConfig *GoogleVertexAiPersistentResourcePscInterfaceConfig `field:"optional" json:"pscInterfaceConfig" yaml:"pscInterfaceConfig"`
	// A list of names for the reserved IP ranges under the VPC network that can be used for this persistent resource.
	//
	// If set, we will deploy the persistent resource within the provided IP
	// ranges. Otherwise, the persistent resource is deployed to any IP
	// ranges under the provided VPC network.
	//
	// Example: ['vertex-ai-ip-range'].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#reserved_ip_ranges GoogleVertexAiPersistentResource#reserved_ip_ranges}
	ReservedIpRanges *[]*string `field:"optional" json:"reservedIpRanges" yaml:"reservedIpRanges"`
	// resource_runtime_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#resource_runtime_spec GoogleVertexAiPersistentResource#resource_runtime_spec}
	ResourceRuntimeSpec *GoogleVertexAiPersistentResourceResourceRuntimeSpec `field:"optional" json:"resourceRuntimeSpec" yaml:"resourceRuntimeSpec"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_persistent_resource#timeouts GoogleVertexAiPersistentResource#timeouts}
	Timeouts *GoogleVertexAiPersistentResourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

