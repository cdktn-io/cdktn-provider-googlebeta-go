// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleconfigdeployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleConfigDeploymentConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#location GoogleConfigDeployment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The user-specified ID of the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#name GoogleConfigDeployment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required. User-specified Service Account (SA) credentials to be used when actuating resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#service_account GoogleConfigDeployment#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// terraform_blueprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#terraform_blueprint GoogleConfigDeployment#terraform_blueprint}
	TerraformBlueprint *GoogleConfigDeploymentTerraformBlueprint `field:"required" json:"terraformBlueprint" yaml:"terraformBlueprint"`
	// Optional. Arbitrary key-value metadata storage.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#annotations GoogleConfigDeployment#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Location for Cloud Build logs and artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#artifacts_gcs_bucket GoogleConfigDeployment#artifacts_gcs_bucket}
	ArtifactsGcsBucket *string `field:"optional" json:"artifactsGcsBucket" yaml:"artifactsGcsBucket"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#deletion_policy GoogleConfigDeployment#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// If true, deletes the deployment and its nested resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#force_destroy GoogleConfigDeployment#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#id GoogleConfigDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, attempts to automatically import resources on 409 conflict.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#import_existing_resources GoogleConfigDeployment#import_existing_resources}
	ImportExistingResources interface{} `field:"optional" json:"importExistingResources" yaml:"importExistingResources"`
	// Optional. User-defined metadata for the deployment.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#labels GoogleConfigDeployment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#project GoogleConfigDeployment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Controls quota checks. Possible values: ["ENABLED", "ENFORCED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#quota_validation GoogleConfigDeployment#quota_validation}
	QuotaValidation *string `field:"optional" json:"quotaValidation" yaml:"quotaValidation"`
	// Optional constraint on the Terraform version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#tf_version_constraint GoogleConfigDeployment#tf_version_constraint}
	TfVersionConstraint *string `field:"optional" json:"tfVersionConstraint" yaml:"tfVersionConstraint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#timeouts GoogleConfigDeployment#timeouts}
	Timeouts *GoogleConfigDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Custom Cloud Build worker pool resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_config_deployment#worker_pool GoogleConfigDeployment#worker_pool}
	WorkerPool *string `field:"optional" json:"workerPool" yaml:"workerPool"`
}

