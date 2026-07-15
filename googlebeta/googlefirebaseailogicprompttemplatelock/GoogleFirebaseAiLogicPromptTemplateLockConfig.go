// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlefirebaseailogicprompttemplatelock

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleFirebaseAiLogicPromptTemplateLockConfig struct {
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
	// The location of the prompt template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_firebase_ai_logic_prompt_template_lock#location GoogleFirebaseAiLogicPromptTemplateLock#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the prompt template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_firebase_ai_logic_prompt_template_lock#template_id GoogleFirebaseAiLogicPromptTemplateLock#template_id}
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_firebase_ai_logic_prompt_template_lock#deletion_policy GoogleFirebaseAiLogicPromptTemplateLock#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_firebase_ai_logic_prompt_template_lock#id GoogleFirebaseAiLogicPromptTemplateLock#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_firebase_ai_logic_prompt_template_lock#project GoogleFirebaseAiLogicPromptTemplateLock#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// For the 'global' location only.
	//
	// If true, the modifyLock operation will
	// apply to the global region only. Otherwise, the operation will also
	// propagate to all applicable regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_firebase_ai_logic_prompt_template_lock#regional_propagation_disabled GoogleFirebaseAiLogicPromptTemplateLock#regional_propagation_disabled}
	RegionalPropagationDisabled interface{} `field:"optional" json:"regionalPropagationDisabled" yaml:"regionalPropagationDisabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_firebase_ai_logic_prompt_template_lock#timeouts GoogleFirebaseAiLogicPromptTemplateLock#timeouts}
	Timeouts *GoogleFirebaseAiLogicPromptTemplateLockTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

