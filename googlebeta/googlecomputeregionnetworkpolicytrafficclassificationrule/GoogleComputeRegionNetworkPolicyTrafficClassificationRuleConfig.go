// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionnetworkpolicytrafficclassificationrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleComputeRegionNetworkPolicyTrafficClassificationRuleConfig struct {
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
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#match GoogleComputeRegionNetworkPolicyTrafficClassificationRule#match}
	Match *GoogleComputeRegionNetworkPolicyTrafficClassificationRuleMatch `field:"required" json:"match" yaml:"match"`
	// The firewall policy of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#network_policy GoogleComputeRegionNetworkPolicyTrafficClassificationRule#network_policy}
	NetworkPolicy *string `field:"required" json:"networkPolicy" yaml:"networkPolicy"`
	// An integer indicating the priority of a rule in the list.
	//
	// The priority must be a positive value between 1 and 2147482647.
	// The priority values from 2147482648 to 2147483647 (1000) are reserved for system default network policy rules.
	// Rules are evaluated from highest to lowest priority where 1 is the highest priority and 2147483647 is the lowest priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#priority GoogleComputeRegionNetworkPolicyTrafficClassificationRule#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#action GoogleComputeRegionNetworkPolicyTrafficClassificationRule#action}
	Action *GoogleComputeRegionNetworkPolicyTrafficClassificationRuleAction `field:"optional" json:"action" yaml:"action"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#deletion_policy GoogleComputeRegionNetworkPolicyTrafficClassificationRule#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// An optional description for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#description GoogleComputeRegionNetworkPolicyTrafficClassificationRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Denotes whether the network policy rule is disabled.
	//
	// When set to true, the network policy rule is not enforced and traffic behaves as if it did not exist.
	// If this is unspecified, the network policy rule will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#disabled GoogleComputeRegionNetworkPolicyTrafficClassificationRule#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#id GoogleComputeRegionNetworkPolicyTrafficClassificationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#project GoogleComputeRegionNetworkPolicyTrafficClassificationRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The location of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#region GoogleComputeRegionNetworkPolicyTrafficClassificationRule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// An optional name for the rule. This field is not a unique identifier and can be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#rule_name GoogleComputeRegionNetworkPolicyTrafficClassificationRule#rule_name}
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// target_secure_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#target_secure_tags GoogleComputeRegionNetworkPolicyTrafficClassificationRule#target_secure_tags}
	TargetSecureTags interface{} `field:"optional" json:"targetSecureTags" yaml:"targetSecureTags"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#target_service_accounts GoogleComputeRegionNetworkPolicyTrafficClassificationRule#target_service_accounts}
	TargetServiceAccounts *[]*string `field:"optional" json:"targetServiceAccounts" yaml:"targetServiceAccounts"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_region_network_policy_traffic_classification_rule#timeouts GoogleComputeRegionNetworkPolicyTrafficClassificationRule#timeouts}
	Timeouts *GoogleComputeRegionNetworkPolicyTrafficClassificationRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

