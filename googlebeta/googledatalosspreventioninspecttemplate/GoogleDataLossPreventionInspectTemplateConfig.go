// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioninspecttemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataLossPreventionInspectTemplateConfig struct {
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
	// The parent of the inspect template in any of the following formats:.
	//
	// * 'projects/{{project}}'
	// * 'projects/{{project}}/locations/{{location}}'
	// * 'organizations/{{organization_id}}'
	// * 'organizations/{{organization_id}}/locations/{{location}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_loss_prevention_inspect_template#parent GoogleDataLossPreventionInspectTemplate#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Enables the use of [limited-availability built-in infoTypes](https://docs.cloud.google.com/sensitive-data-protection/docs/infotypes-reference#limited-availability-infotypes) in inspect_config. These infoTypes are supported only in specific regions and can cause scanning errors if used elsewhere.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_loss_prevention_inspect_template#allow_limited_availability_info_types GoogleDataLossPreventionInspectTemplate#allow_limited_availability_info_types}
	AllowLimitedAvailabilityInfoTypes interface{} `field:"optional" json:"allowLimitedAvailabilityInfoTypes" yaml:"allowLimitedAvailabilityInfoTypes"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_loss_prevention_inspect_template#deletion_policy GoogleDataLossPreventionInspectTemplate#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// A description of the inspect template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_loss_prevention_inspect_template#description GoogleDataLossPreventionInspectTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User set display name of the inspect template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_loss_prevention_inspect_template#display_name GoogleDataLossPreventionInspectTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_loss_prevention_inspect_template#id GoogleDataLossPreventionInspectTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inspect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_loss_prevention_inspect_template#inspect_config GoogleDataLossPreventionInspectTemplate#inspect_config}
	InspectConfig *GoogleDataLossPreventionInspectTemplateInspectConfig `field:"optional" json:"inspectConfig" yaml:"inspectConfig"`
	// The template id can contain uppercase and lowercase letters, numbers, and hyphens;
	//
	// that is, it must match the regular expression: [a-zA-Z\d-_]+. The maximum length is
	// 100 characters. Can be empty to allow the system to generate one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_loss_prevention_inspect_template#template_id GoogleDataLossPreventionInspectTemplate#template_id}
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_data_loss_prevention_inspect_template#timeouts GoogleDataLossPreventionInspectTemplate#timeouts}
	Timeouts *GoogleDataLossPreventionInspectTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

