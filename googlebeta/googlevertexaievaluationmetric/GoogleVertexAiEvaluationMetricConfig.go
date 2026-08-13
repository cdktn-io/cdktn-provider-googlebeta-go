// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaievaluationmetric

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiEvaluationMetricConfig struct {
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
	// The user-friendly display name for the EvaluationMetric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#display_name GoogleVertexAiEvaluationMetric#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The region of the EvaluationMetric. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#region GoogleVertexAiEvaluationMetric#region}
	Region *string `field:"required" json:"region" yaml:"region"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#deletion_policy GoogleVertexAiEvaluationMetric#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// A description of the EvaluationMetric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#description GoogleVertexAiEvaluationMetric#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#encryption_spec GoogleVertexAiEvaluationMetric#encryption_spec}
	EncryptionSpec *GoogleVertexAiEvaluationMetricEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// The ID to use for the EvaluationMetric, which will become the final component of the resource name.
	//
	// This value should be 1-63 characters,
	// and valid characters are /[a-z][0-9]-/. The first character must be
	// a lowercase letter, and the last character must be a lowercase letter
	// or number. If not provided, the server will generate a unique ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#evaluation_metric_id GoogleVertexAiEvaluationMetric#evaluation_metric_id}
	EvaluationMetricId *string `field:"optional" json:"evaluationMetricId" yaml:"evaluationMetricId"`
	// The Google Cloud Storage URI that stores the metric specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#gcs_uri GoogleVertexAiEvaluationMetric#gcs_uri}
	GcsUri *string `field:"optional" json:"gcsUri" yaml:"gcsUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#id GoogleVertexAiEvaluationMetric#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels for the EvaluationMetric.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#labels GoogleVertexAiEvaluationMetric#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The metric configuration as a JSON string.
	//
	// Uses camelCase field names
	// to match the API format. Supports LLM-based metrics and custom code
	// execution metrics.
	// See the [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/Metric)
	// for the full schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#metric GoogleVertexAiEvaluationMetric#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#project GoogleVertexAiEvaluationMetric#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_evaluation_metric#timeouts GoogleVertexAiEvaluationMetric#timeouts}
	Timeouts *GoogleVertexAiEvaluationMetricTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

