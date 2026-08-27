// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleContainerNodePoolConfig struct {
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
	// The cluster to create the node pool for. Cluster must be present in location provided for zonal clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#cluster GoogleContainerNodePool#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#autoscaling GoogleContainerNodePool#autoscaling}
	Autoscaling *GoogleContainerNodePoolAutoscaling `field:"optional" json:"autoscaling" yaml:"autoscaling"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#deletion_policy GoogleContainerNodePool#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#id GoogleContainerNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When true, the provider ignores external changes (drift) to the node count by skipping GCE API queries to the Instance Group Managers.
	//
	// This is a performance optimization for large clusters that saves API quota. Setting this to true will result in missing managed_instance_group_urls in the state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#ignore_node_count_changes GoogleContainerNodePool#ignore_node_count_changes}
	IgnoreNodeCountChanges interface{} `field:"optional" json:"ignoreNodeCountChanges" yaml:"ignoreNodeCountChanges"`
	// The initial number of nodes for the pool.
	//
	// In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#initial_node_count GoogleContainerNodePool#initial_node_count}
	InitialNodeCount *float64 `field:"optional" json:"initialNodeCount" yaml:"initialNodeCount"`
	// The location (region or zone) of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#location GoogleContainerNodePool#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#maintenance_policy GoogleContainerNodePool#maintenance_policy}
	MaintenancePolicy interface{} `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#management GoogleContainerNodePool#management}
	Management *GoogleContainerNodePoolManagement `field:"optional" json:"management" yaml:"management"`
	// The maximum number of pods per node in this node pool.
	//
	// Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#max_pods_per_node GoogleContainerNodePool#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
	// The name of the node pool. If left blank, Terraform will auto-generate a unique name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#name GoogleContainerNodePool#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#name_prefix GoogleContainerNodePool#name_prefix}
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#network_config GoogleContainerNodePool#network_config}
	NetworkConfig *GoogleContainerNodePoolNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#node_config GoogleContainerNodePool#node_config}
	NodeConfig *GoogleContainerNodePoolNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// The number of nodes per instance group.
	//
	// This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#node_count GoogleContainerNodePool#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// node_drain_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#node_drain_config GoogleContainerNodePool#node_drain_config}
	NodeDrainConfig interface{} `field:"optional" json:"nodeDrainConfig" yaml:"nodeDrainConfig"`
	// The list of zones in which the node pool's nodes should be located.
	//
	// Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#node_locations GoogleContainerNodePool#node_locations}
	NodeLocations *[]*string `field:"optional" json:"nodeLocations" yaml:"nodeLocations"`
	// placement_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#placement_policy GoogleContainerNodePool#placement_policy}
	PlacementPolicy *GoogleContainerNodePoolPlacementPolicy `field:"optional" json:"placementPolicy" yaml:"placementPolicy"`
	// The ID of the project in which to create the node pool.
	//
	// If blank, the provider-configured project will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#project GoogleContainerNodePool#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// queued_provisioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#queued_provisioning GoogleContainerNodePool#queued_provisioning}
	QueuedProvisioning *GoogleContainerNodePoolQueuedProvisioning `field:"optional" json:"queuedProvisioning" yaml:"queuedProvisioning"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#timeouts GoogleContainerNodePool#timeouts}
	Timeouts *GoogleContainerNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#upgrade_settings GoogleContainerNodePool#upgrade_settings}
	UpgradeSettings *GoogleContainerNodePoolUpgradeSettings `field:"optional" json:"upgradeSettings" yaml:"upgradeSettings"`
	// The Kubernetes version for the nodes in this pool.
	//
	// Note that if this field and auto_upgrade are both specified, they will fight each other for what the node version should be, so setting both is highly discouraged. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the google_container_engine_versions data source's version_prefix field to approximate fuzzy versions in a Terraform-compatible way.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#version GoogleContainerNodePool#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

