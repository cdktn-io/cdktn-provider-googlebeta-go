// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleapphubserviceprojectattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleApphubServiceProjectAttachmentConfig struct {
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
	// Required.
	//
	// The service project attachment identifier must contain the project_id of the service project specified in the service_project_attachment.service_project field. Hint: "projects/{project_id}"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apphub_service_project_attachment#service_project_attachment_id GoogleApphubServiceProjectAttachment#service_project_attachment_id}
	ServiceProjectAttachmentId *string `field:"required" json:"serviceProjectAttachmentId" yaml:"serviceProjectAttachmentId"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apphub_service_project_attachment#deletion_policy GoogleApphubServiceProjectAttachment#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apphub_service_project_attachment#id GoogleApphubServiceProjectAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apphub_service_project_attachment#project GoogleApphubServiceProjectAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// "Immutable.
	//
	// Service project name in the format: \"projects/abc\"
	// or \"projects/123\". As input, project name with either project id or number
	// are accepted. As output, this field will contain project number."
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apphub_service_project_attachment#service_project GoogleApphubServiceProjectAttachment#service_project}
	ServiceProject *string `field:"optional" json:"serviceProject" yaml:"serviceProject"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_apphub_service_project_attachment#timeouts GoogleApphubServiceProjectAttachment#timeouts}
	Timeouts *GoogleApphubServiceProjectAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

