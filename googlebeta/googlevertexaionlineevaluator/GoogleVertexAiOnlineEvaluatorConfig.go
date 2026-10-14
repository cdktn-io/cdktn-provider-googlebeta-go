// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiOnlineEvaluatorConfig struct {
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
	// The name of the agent that the OnlineEvaluator evaluates periodically.
	//
	// This value is used to filter the traces with a matching cloud.resource_id
	// and link the evaluation results with relevant dashboards/UIs. This field
	// is immutable; once set, it cannot be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#agent_resource GoogleVertexAiOnlineEvaluator#agent_resource}
	AgentResource *string `field:"required" json:"agentResource" yaml:"agentResource"`
	// cloud_observability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#cloud_observability GoogleVertexAiOnlineEvaluator#cloud_observability}
	CloudObservability *GoogleVertexAiOnlineEvaluatorCloudObservability `field:"required" json:"cloudObservability" yaml:"cloudObservability"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#config GoogleVertexAiOnlineEvaluator#config}
	Config *GoogleVertexAiOnlineEvaluatorConfigA `field:"required" json:"config" yaml:"config"`
	// metric_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#metric_sources GoogleVertexAiOnlineEvaluator#metric_sources}
	MetricSources interface{} `field:"required" json:"metricSources" yaml:"metricSources"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#deletion_policy GoogleVertexAiOnlineEvaluator#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Human-readable name for the OnlineEvaluator.
	//
	// The name doesn't have to be
	// unique. The name can consist of any UTF-8 characters. The maximum length
	// is 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#display_name GoogleVertexAiOnlineEvaluator#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#id GoogleVertexAiOnlineEvaluator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#project GoogleVertexAiOnlineEvaluator#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the OnlineEvaluator (e.g. us-central1). If unset, the provider's region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#region GoogleVertexAiOnlineEvaluator#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#timeouts GoogleVertexAiOnlineEvaluator#timeouts}
	Timeouts *GoogleVertexAiOnlineEvaluatorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

