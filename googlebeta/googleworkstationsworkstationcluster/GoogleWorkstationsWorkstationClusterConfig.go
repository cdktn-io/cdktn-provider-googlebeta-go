// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleworkstationsworkstationcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleWorkstationsWorkstationClusterConfig struct {
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
	// The relative resource name of the VPC network on which the instance can be accessed.
	//
	// It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#network GoogleWorkstationsWorkstationCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
	//
	// Must be part of the subnetwork specified for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#subnetwork GoogleWorkstationsWorkstationCluster#subnetwork}
	Subnetwork *string `field:"required" json:"subnetwork" yaml:"subnetwork"`
	// ID to use for the workstation cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#workstation_cluster_id GoogleWorkstationsWorkstationCluster#workstation_cluster_id}
	WorkstationClusterId *string `field:"required" json:"workstationClusterId" yaml:"workstationClusterId"`
	// Client-specified annotations. This is distinct from labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#annotations GoogleWorkstationsWorkstationCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#deletion_policy GoogleWorkstationsWorkstationCluster#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Human-readable name for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#display_name GoogleWorkstationsWorkstationCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// domain_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#domain_config GoogleWorkstationsWorkstationCluster#domain_config}
	DomainConfig *GoogleWorkstationsWorkstationClusterDomainConfig `field:"optional" json:"domainConfig" yaml:"domainConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#id GoogleWorkstationsWorkstationCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#labels GoogleWorkstationsWorkstationCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location where the workstation cluster should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#location GoogleWorkstationsWorkstationCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// private_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#private_cluster_config GoogleWorkstationsWorkstationCluster#private_cluster_config}
	PrivateClusterConfig *GoogleWorkstationsWorkstationClusterPrivateClusterConfig `field:"optional" json:"privateClusterConfig" yaml:"privateClusterConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#project GoogleWorkstationsWorkstationCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Resource manager tags bound to this resource. For example: "123/environment": "production", "123/costCenter": "marketing".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#tags GoogleWorkstationsWorkstationCluster#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#timeouts GoogleWorkstationsWorkstationCluster#timeouts}
	Timeouts *GoogleWorkstationsWorkstationClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Specifies the redirect URL for unauthorized requests received by workstation VMs in this cluster.
	//
	// Redirects to this endpoint will send a base64 encoded 'state' query param containing the target workstation name and original request hostname. The endpoint is responsible for retrieving a token using 'GenerateAccessToken' and redirecting back to the original hostname with the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#workstation_authorization_url GoogleWorkstationsWorkstationCluster#workstation_authorization_url}
	WorkstationAuthorizationUrl *string `field:"optional" json:"workstationAuthorizationUrl" yaml:"workstationAuthorizationUrl"`
	// Specifies the launch URL for workstations in this cluster.
	//
	// Requests sent to unstarted workstations will be redirected to this URL.
	// Requests redirected to the launch endpoint will be sent with a 'workstation' query parameter containing the full workstation resource. The launch endpoint is responsible for starting the workstation, polling it until it reaches 'STATE_RUNNING', and then issuing a redirect to the workstation's host URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_workstations_workstation_cluster#workstation_launch_url GoogleWorkstationsWorkstationCluster#workstation_launch_url}
	WorkstationLaunchUrl *string `field:"optional" json:"workstationLaunchUrl" yaml:"workstationLaunchUrl"`
}

