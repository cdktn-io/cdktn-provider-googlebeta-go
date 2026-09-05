// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaitensorboardrun

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiTensorboardRunConfig struct {
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
	// User provided name of this TensorboardRun. This value must be unique among all TensorboardRuns belonging to the same parent TensorboardExperiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#display_name GoogleVertexAiTensorboardRun#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Tensorboard Experiment ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#experiment GoogleVertexAiTensorboardRun#experiment}
	Experiment *string `field:"required" json:"experiment" yaml:"experiment"`
	// The location of the Tensorboard Run. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#location GoogleVertexAiTensorboardRun#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The Tensorboard instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#tensorboard GoogleVertexAiTensorboardRun#tensorboard}
	Tensorboard *string `field:"required" json:"tensorboard" yaml:"tensorboard"`
	// The ID to use for the Tensorboard run, which becomes the final component of the Tensorboard run's resource name.
	//
	// This value should be 1-128 characters, and valid characters
	// are '/a-z-/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#tensorboard_run_id GoogleVertexAiTensorboardRun#tensorboard_run_id}
	TensorboardRunId *string `field:"required" json:"tensorboardRunId" yaml:"tensorboardRunId"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#deletion_policy GoogleVertexAiTensorboardRun#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Description of this TensorboardRun.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#description GoogleVertexAiTensorboardRun#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#id GoogleVertexAiTensorboardRun#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels with user-defined metadata to organize your TensorboardRuns.
	//
	// This field will be used to filter and visualize Runs in the Tensorboard UI.
	// For example, a Vertex AI training job can set a label
	// aiplatform.googleapis.com/training_job_id=xxxxx to all the runs created
	// within that job. An end user can set a label experiment_id=xxxxx for all
	// the runs produced in a Jupyter notebook. These runs can be grouped by a
	// label value and visualized together in the Tensorboard UI.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// No more than 64 user labels can be associated with one TensorboardRun
	// (System labels are excluded).
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	// and are immutable.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#labels GoogleVertexAiTensorboardRun#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#project GoogleVertexAiTensorboardRun#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_tensorboard_run#timeouts GoogleVertexAiTensorboardRun#timeouts}
	Timeouts *GoogleVertexAiTensorboardRunTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

