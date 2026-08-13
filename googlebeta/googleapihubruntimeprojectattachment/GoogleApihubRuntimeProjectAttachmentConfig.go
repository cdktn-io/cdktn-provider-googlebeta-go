// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleapihubruntimeprojectattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleApihubRuntimeProjectAttachmentConfig struct {
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
	// Part of 'parent'. See documentation of 'projectsId'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apihub_runtime_project_attachment#location GoogleApihubRuntimeProjectAttachment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required.
	//
	// Immutable. Google cloud project name in the format: "projects/abc" or "projects/123".
	// As input, project name with either project id or number are accepted.
	// As output, this field will contain project number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apihub_runtime_project_attachment#runtime_project GoogleApihubRuntimeProjectAttachment#runtime_project}
	RuntimeProject *string `field:"required" json:"runtimeProject" yaml:"runtimeProject"`
	// The ID to use for the Runtime Project Attachment, which will become the final component of the Runtime Project Attachment's name.
	//
	// The ID must be the same
	// as the project ID of the Google cloud project specified in the
	// runtime_project_attachment.runtime_project field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apihub_runtime_project_attachment#runtime_project_attachment_id GoogleApihubRuntimeProjectAttachment#runtime_project_attachment_id}
	RuntimeProjectAttachmentId *string `field:"required" json:"runtimeProjectAttachmentId" yaml:"runtimeProjectAttachmentId"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apihub_runtime_project_attachment#deletion_policy GoogleApihubRuntimeProjectAttachment#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apihub_runtime_project_attachment#id GoogleApihubRuntimeProjectAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apihub_runtime_project_attachment#project GoogleApihubRuntimeProjectAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apihub_runtime_project_attachment#timeouts GoogleApihubRuntimeProjectAttachment#timeouts}
	Timeouts *GoogleApihubRuntimeProjectAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

