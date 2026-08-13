// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengatedeployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateDeploymentConfig struct {
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
	// The display name for the GoldengateDeployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#display_name GoogleOracleDatabaseGoldengateDeployment#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID of the GoldengateDeployment to create.
	//
	// This value is
	// restricted to (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of
	// 63 characters in length. The value must start with a letter and end with a
	// letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#goldengate_deployment_id GoogleOracleDatabaseGoldengateDeployment#goldengate_deployment_id}
	GoldengateDeploymentId *string `field:"required" json:"goldengateDeploymentId" yaml:"goldengateDeploymentId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#location GoogleOracleDatabaseGoldengateDeployment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the OdbSubnet associated with the GoldengateDeployment for IP allocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#odb_subnet GoogleOracleDatabaseGoldengateDeployment#odb_subnet}
	OdbSubnet *string `field:"required" json:"odbSubnet" yaml:"odbSubnet"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#properties GoogleOracleDatabaseGoldengateDeployment#properties}
	Properties *GoogleOracleDatabaseGoldengateDeploymentProperties `field:"required" json:"properties" yaml:"properties"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "PREVENT".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#deletion_policy GoogleOracleDatabaseGoldengateDeployment#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The GCP Oracle zone where Oracle GoldengateDeployment is hosted.
	//
	// Example: us-east4-b-r2.
	// If not specified, the system will pick a zone based on availability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#gcp_oracle_zone GoogleOracleDatabaseGoldengateDeployment#gcp_oracle_zone}
	GcpOracleZone *string `field:"optional" json:"gcpOracleZone" yaml:"gcpOracleZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#id GoogleOracleDatabaseGoldengateDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels or tags associated with the GoldengateDeployment.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#labels GoogleOracleDatabaseGoldengateDeployment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the OdbNetwork associated with the GoldengateDeployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#odb_network GoogleOracleDatabaseGoldengateDeployment#odb_network}
	OdbNetwork *string `field:"optional" json:"odbNetwork" yaml:"odbNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#project GoogleOracleDatabaseGoldengateDeployment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_deployment#timeouts GoogleOracleDatabaseGoldengateDeployment#timeouts}
	Timeouts *GoogleOracleDatabaseGoldengateDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

