// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaitensorboardexperiment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiTensorboardExperimentConfig struct {
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
	// The location of the Tensorboard Experiment. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#location GoogleVertexAiTensorboardExperiment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The Tensorboard instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#tensorboard GoogleVertexAiTensorboardExperiment#tensorboard}
	Tensorboard *string `field:"required" json:"tensorboard" yaml:"tensorboard"`
	// The ID to use for the Tensorboard experiment, which becomes the final component of the Tensorboard experiment's resource name.
	//
	// This value should be 1-128 characters, and valid characters
	// are '/a-z-/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#tensorboard_experiment_id GoogleVertexAiTensorboardExperiment#tensorboard_experiment_id}
	TensorboardExperimentId *string `field:"required" json:"tensorboardExperimentId" yaml:"tensorboardExperimentId"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#deletion_policy GoogleVertexAiTensorboardExperiment#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Description of this TensorboardExperiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#description GoogleVertexAiTensorboardExperiment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User provided name of this TensorboardExperiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#display_name GoogleVertexAiTensorboardExperiment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#id GoogleVertexAiTensorboardExperiment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels with user-defined metadata to organize your TensorboardExperiment.
	//
	// Label keys and values cannot be longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// No more than 64 user labels can be associated with one Dataset (System
	// labels are excluded).
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// System reserved label keys are prefixed with 'aiplatform.googleapis.com/'
	// and are immutable. The following system labels exist for each Dataset:
	//
	// * 'aiplatform.googleapis.com/dataset_metadata_schema': output only. Its
	// value is the metadata_schema's title.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#labels GoogleVertexAiTensorboardExperiment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#project GoogleVertexAiTensorboardExperiment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Source of the TensorboardExperiment. Example: a custom training job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#source GoogleVertexAiTensorboardExperiment#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_tensorboard_experiment#timeouts GoogleVertexAiTensorboardExperiment#timeouts}
	Timeouts *GoogleVertexAiTensorboardExperimentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

