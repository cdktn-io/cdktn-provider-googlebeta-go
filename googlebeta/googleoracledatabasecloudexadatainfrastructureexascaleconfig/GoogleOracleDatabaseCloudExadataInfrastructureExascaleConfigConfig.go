// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasecloudexadatainfrastructureexascaleconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseCloudExadataInfrastructureExascaleConfigConfig struct {
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
	// A reference to CloudExadataInfrastructure resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure_exascale_config#cloud_exadata_infrastructure GoogleOracleDatabaseCloudExadataInfrastructureExascaleConfig#cloud_exadata_infrastructure}
	CloudExadataInfrastructure *string `field:"required" json:"cloudExadataInfrastructure" yaml:"cloudExadataInfrastructure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure_exascale_config#location GoogleOracleDatabaseCloudExadataInfrastructureExascaleConfig#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// The total storage to be allocated to Exascale in GBs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure_exascale_config#total_storage_size_gb GoogleOracleDatabaseCloudExadataInfrastructureExascaleConfig#total_storage_size_gb}
	TotalStorageSizeGb *float64 `field:"required" json:"totalStorageSizeGb" yaml:"totalStorageSizeGb"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure_exascale_config#deletion_policy GoogleOracleDatabaseCloudExadataInfrastructureExascaleConfig#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure_exascale_config#id GoogleOracleDatabaseCloudExadataInfrastructureExascaleConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure_exascale_config#project GoogleOracleDatabaseCloudExadataInfrastructureExascaleConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure_exascale_config#timeouts GoogleOracleDatabaseCloudExadataInfrastructureExascaleConfig#timeouts}
	Timeouts *GoogleOracleDatabaseCloudExadataInfrastructureExascaleConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

