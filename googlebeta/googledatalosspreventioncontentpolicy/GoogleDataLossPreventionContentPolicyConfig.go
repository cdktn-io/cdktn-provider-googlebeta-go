// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataLossPreventionContentPolicyConfig struct {
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
	// The parent of the content policy in any of the following formats:.
	//
	// * 'projects/{{project}}/locations/{{location}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#parent GoogleDataLossPreventionContentPolicy#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#rules GoogleDataLossPreventionContentPolicy#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// default_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#default_action GoogleDataLossPreventionContentPolicy#default_action}
	DefaultAction *GoogleDataLossPreventionContentPolicyDefaultAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#deletion_policy GoogleDataLossPreventionContentPolicy#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Display name (max 63 chars).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#display_name GoogleDataLossPreventionContentPolicy#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// failed_to_scan_supported_file_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#failed_to_scan_supported_file_type GoogleDataLossPreventionContentPolicy#failed_to_scan_supported_file_type}
	FailedToScanSupportedFileType *GoogleDataLossPreventionContentPolicyFailedToScanSupportedFileType `field:"optional" json:"failedToScanSupportedFileType" yaml:"failedToScanSupportedFileType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#id GoogleDataLossPreventionContentPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// input_too_large block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#input_too_large GoogleDataLossPreventionContentPolicy#input_too_large}
	InputTooLarge *GoogleDataLossPreventionContentPolicyInputTooLarge `field:"optional" json:"inputTooLarge" yaml:"inputTooLarge"`
	// inspect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#inspect_config GoogleDataLossPreventionContentPolicy#inspect_config}
	InspectConfig *GoogleDataLossPreventionContentPolicyInspectConfig `field:"optional" json:"inspectConfig" yaml:"inspectConfig"`
	// logging_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#logging_configs GoogleDataLossPreventionContentPolicy#logging_configs}
	LoggingConfigs interface{} `field:"optional" json:"loggingConfigs" yaml:"loggingConfigs"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#timeouts GoogleDataLossPreventionContentPolicy#timeouts}
	Timeouts *GoogleDataLossPreventionContentPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// unsupported_file_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#unsupported_file_type GoogleDataLossPreventionContentPolicy#unsupported_file_type}
	UnsupportedFileType *GoogleDataLossPreventionContentPolicyUnsupportedFileType `field:"optional" json:"unsupportedFileType" yaml:"unsupportedFileType"`
}

