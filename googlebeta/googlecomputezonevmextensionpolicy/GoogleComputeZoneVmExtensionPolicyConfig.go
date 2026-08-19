// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputezonevmextensionpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleComputeZoneVmExtensionPolicyConfig struct {
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
	// extension_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#extension_policies GoogleComputeZoneVmExtensionPolicy#extension_policies}
	ExtensionPolicies interface{} `field:"required" json:"extensionPolicies" yaml:"extensionPolicies"`
	// Name of the resource. Provided by the client when the resource is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#name GoogleComputeZoneVmExtensionPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the zone for this request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#zone GoogleComputeZoneVmExtensionPolicy#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#deletion_policy GoogleComputeZoneVmExtensionPolicy#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#description GoogleComputeZoneVmExtensionPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// instance_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#instance_selectors GoogleComputeZoneVmExtensionPolicy#instance_selectors}
	InstanceSelectors interface{} `field:"optional" json:"instanceSelectors" yaml:"instanceSelectors"`
	// Priority of this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#priority GoogleComputeZoneVmExtensionPolicy#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#project GoogleComputeZoneVmExtensionPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#timeouts GoogleComputeZoneVmExtensionPolicy#timeouts}
	Timeouts *GoogleComputeZoneVmExtensionPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

