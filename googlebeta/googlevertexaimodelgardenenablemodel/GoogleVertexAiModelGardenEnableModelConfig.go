// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodelgardenenablemodel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiModelGardenEnableModelConfig struct {
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
	// The resource name of the Model Garden publisher model to enable. Format: 'publishers/{publisher}/models/{publisher_model}', optionally with a version suffix, for example 'publishers/google/models/paligemma@paligemma-224-float32'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_model_garden_enable_model#publisher_model_name GoogleVertexAiModelGardenEnableModel#publisher_model_name}
	PublisherModelName *string `field:"required" json:"publisherModelName" yaml:"publisherModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_model_garden_enable_model#id GoogleVertexAiModelGardenEnableModel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_model_garden_enable_model#project GoogleVertexAiModelGardenEnableModel#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_model_garden_enable_model#timeouts GoogleVertexAiModelGardenEnableModel#timeouts}
	Timeouts *GoogleVertexAiModelGardenEnableModelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

