// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleserviceusagev2consumerpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleServiceUsageV2ConsumerPolicyConfig struct {
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
	// enable_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_service_usage_v2_consumer_policy#enable_rules GoogleServiceUsageV2ConsumerPolicy#enable_rules}
	EnableRules interface{} `field:"required" json:"enableRules" yaml:"enableRules"`
	// The name of the policy. Currently only the “default” policy name is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_service_usage_v2_consumer_policy#name GoogleServiceUsageV2ConsumerPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the parent.
	//
	// It can be a project, folder or organization in the format of  projects/<project_id or project_number>, folders/<folder_number> or organizations/<org_number> .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_service_usage_v2_consumer_policy#parent GoogleServiceUsageV2ConsumerPolicy#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// (Optional) Default value is false.
	//
	// If true, the usage of the service to be removed will be checked. If the service has been used within the past 30 days or was enabled in the last 3 days, an error will be thrown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_service_usage_v2_consumer_policy#check_usage_on_remove GoogleServiceUsageV2ConsumerPolicy#check_usage_on_remove}
	CheckUsageOnRemove interface{} `field:"optional" json:"checkUsageOnRemove" yaml:"checkUsageOnRemove"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_service_usage_v2_consumer_policy#deletion_policy GoogleServiceUsageV2ConsumerPolicy#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_service_usage_v2_consumer_policy#id GoogleServiceUsageV2ConsumerPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_service_usage_v2_consumer_policy#timeouts GoogleServiceUsageV2ConsumerPolicy#timeouts}
	Timeouts *GoogleServiceUsageV2ConsumerPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// (Optional) Default value is true.
	//
	// If true, this flag enforces dependency management within the consumer policy. When adding a new service, it verifies that all its dependencies are already present/added in the policy. Conversely, when removing a service, it ensures that no other services within the policy depend on the service to be removed. If the validation fails, a comprehensive message will be presented, outlining the missing dependencies and providing instructions on how to address the issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_service_usage_v2_consumer_policy#validate_dependencies GoogleServiceUsageV2ConsumerPolicy#validate_dependencies}
	ValidateDependencies interface{} `field:"optional" json:"validateDependencies" yaml:"validateDependencies"`
}

