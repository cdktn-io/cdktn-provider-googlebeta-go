// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterNodePools struct {
	// The unique ID referring to this node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#id GoogleVertexAiModelDevelopmentCluster#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#machine_spec GoogleVertexAiModelDevelopmentCluster#machine_spec}
	MachineSpec *GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpec `field:"required" json:"machineSpec" yaml:"machineSpec"`
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#boot_disk GoogleVertexAiModelDevelopmentCluster#boot_disk}
	BootDisk *GoogleVertexAiModelDevelopmentClusterNodePoolsBootDisk `field:"optional" json:"bootDisk" yaml:"bootDisk"`
	// Whether to enable public IPs for the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#enable_public_ips GoogleVertexAiModelDevelopmentCluster#enable_public_ips}
	EnablePublicIps interface{} `field:"optional" json:"enablePublicIps" yaml:"enablePublicIps"`
	// List of the existing Filestore instances to be mounted.
	//
	// For example:
	// if "projects/my-project/locations/us-central1-a/instances/my-instance" is
	// specified, the instance will be mounted to the node pool under the path
	// "mnt/filestore/"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#filestores GoogleVertexAiModelDevelopmentCluster#filestores}
	Filestores *[]*string `field:"optional" json:"filestores" yaml:"filestores"`
	// The maximum duration for instances in this node pool. This field is only applicable when 'provisioning_model' is set to 'FLEX_START'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#flex_start_max_duration GoogleVertexAiModelDevelopmentCluster#flex_start_max_duration}
	FlexStartMaxDuration *string `field:"optional" json:"flexStartMaxDuration" yaml:"flexStartMaxDuration"`
	// The labels with user-defined metadata to organize the node pool.
	//
	// These
	// labels are applied to the underlying compute resources (for example, the
	// VM instances) of the node pool, so they appear in billing exports and
	// Cloud Monitoring filters.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#labels GoogleVertexAiModelDevelopmentCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// List of the existing Google Managed Lustre instances to be mounted.
	//
	// Similar to the filestores, the instance will be mounted to the node pool
	// under the path "/mnt/lustre/"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#lustres GoogleVertexAiModelDevelopmentCluster#lustres}
	Lustres *[]*string `field:"optional" json:"lustres" yaml:"lustres"`
	// The node image to be used for the VMs in the node pool. In format of "projects//global/images/" or "projects/test-project/global/images/family/".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#node_image GoogleVertexAiModelDevelopmentCluster#node_image}
	NodeImage *string `field:"optional" json:"nodeImage" yaml:"nodeImage"`
	// Specifies the provisioning model for the node pool.
	//
	// Behavior for 'PROVISIONING_MODEL_UNSPECIFIED':
	// - For CPU node pools: Defaults to 'ON_DEMAND'.
	// - For GPU node pools:
	// - If
	// 'machine_spec.reservation_affinity'
	// is set, it defaults to 'RESERVATION'.
	// - If
	// 'machine_spec.reservation_affinity'
	// is NOT set, this is an INVALID state and the request will be rejected.
	// An explicit 'provisioning_model' must be provided.
	//
	// ## Constraints:
	//
	// 'machine_spec.reservation_affinity'
	// must not be set if 'provisioning_model' is 'ON_DEMAND', 'SPOT' or
	// ## 'FLEX_START'.
	//
	// 'machine_spec.reservation_affinity'
	// must be set if 'provisioning_model' is 'RESERVATION'.
	// Possible values:
	// ON_DEMAND
	// RESERVATION
	// SPOT
	// FLEX_START
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#provisioning_model GoogleVertexAiModelDevelopmentCluster#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// scaling_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#scaling_spec GoogleVertexAiModelDevelopmentCluster#scaling_spec}
	ScalingSpec *GoogleVertexAiModelDevelopmentClusterNodePoolsScalingSpec `field:"optional" json:"scalingSpec" yaml:"scalingSpec"`
	// Startup script for the node pool instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#startup_script GoogleVertexAiModelDevelopmentCluster#startup_script}
	StartupScript *string `field:"optional" json:"startupScript" yaml:"startupScript"`
	// The zone of the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#zone GoogleVertexAiModelDevelopmentCluster#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

