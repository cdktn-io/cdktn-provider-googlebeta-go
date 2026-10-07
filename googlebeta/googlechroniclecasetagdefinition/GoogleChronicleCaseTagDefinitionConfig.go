// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclecasetagdefinition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleCaseTagDefinitionConfig struct {
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
	// When checked, the tag will be assigned as the title of the case if it meets the conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#can_be_case_title GoogleChronicleCaseTagDefinition#can_be_case_title}
	CanBeCaseTitle interface{} `field:"required" json:"canBeCaseTitle" yaml:"canBeCaseTitle"`
	// The type of comparison to be used when comparing the value to the case.
	//
	// Possible values: ["EXACT", "START_WITH", "CONTAIN", "ENDS_WITH"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#comparison_type GoogleChronicleCaseTagDefinition#comparison_type}
	ComparisonType *string `field:"required" json:"comparisonType" yaml:"comparisonType"`
	// This is the name of the tag that will be applied to the case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#display_name GoogleChronicleCaseTagDefinition#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#instance GoogleChronicleCaseTagDefinition#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#location GoogleChronicleCaseTagDefinition#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The criteria to match the case against. Possible values: ["BY_VENDOR", "BY_PRODUCT", "BY_RULE_GENERATOR", "BY_ENTITY_PROPERTY_NAME", "DATA_DRIVEN", "SYSTEM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#match_criteria GoogleChronicleCaseTagDefinition#match_criteria}
	MatchCriteria *string `field:"required" json:"matchCriteria" yaml:"matchCriteria"`
	// Note that Google Security Operations merges priority with other alerts and entities and events so that the priority here is not absolute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#priority GoogleChronicleCaseTagDefinition#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Specific value to search in case - in addition to SearchIn property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#value GoogleChronicleCaseTagDefinition#value}
	Value *string `field:"required" json:"value" yaml:"value"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#deletion_policy GoogleChronicleCaseTagDefinition#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#id GoogleChronicleCaseTagDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#project GoogleChronicleCaseTagDefinition#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Specific Entity property name to search in case. This is relevant only when a SearchIn of type BY_ENTITY_PROPERTY_NAME was chosen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#property_name GoogleChronicleCaseTagDefinition#property_name}
	PropertyName *string `field:"optional" json:"propertyName" yaml:"propertyName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_chronicle_case_tag_definition#timeouts GoogleChronicleCaseTagDefinition#timeouts}
	Timeouts *GoogleChronicleCaseTagDefinitionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

