// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemigrationcenterassetsexportjob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleMigrationCenterAssetsExportJobConfig struct {
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
	// The ID to use for the asset export job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#assets_export_job_id GoogleMigrationCenterAssetsExportJob#assets_export_job_id}
	AssetsExportJobId *string `field:"required" json:"assetsExportJobId" yaml:"assetsExportJobId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#location GoogleMigrationCenterAssetsExportJob#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#condition GoogleMigrationCenterAssetsExportJob#condition}
	Condition *GoogleMigrationCenterAssetsExportJobCondition `field:"optional" json:"condition" yaml:"condition"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#deletion_policy GoogleMigrationCenterAssetsExportJob#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#id GoogleMigrationCenterAssetsExportJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs. Labels must meet the following constraints:.
	//
	// * Keys and values can contain only lowercase letters, numeric characters,
	// underscores, and dashes.
	// * All characters must use UTF-8 encoding, and international characters are
	// allowed.
	// * Keys must start with a lowercase letter or international character.
	// * Each resource is limited to a maximum of 64 labels.
	//
	// Both keys and values are additionally constrained to be <= 128 bytes.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#labels GoogleMigrationCenterAssetsExportJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// performance_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#performance_data GoogleMigrationCenterAssetsExportJob#performance_data}
	PerformanceData *GoogleMigrationCenterAssetsExportJobPerformanceData `field:"optional" json:"performanceData" yaml:"performanceData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#project GoogleMigrationCenterAssetsExportJob#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// When this value is set to 'true' the response will include all assets, including those that are hidden.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#show_hidden GoogleMigrationCenterAssetsExportJob#show_hidden}
	ShowHidden interface{} `field:"optional" json:"showHidden" yaml:"showHidden"`
	// signed_uri_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#signed_uri_destination GoogleMigrationCenterAssetsExportJob#signed_uri_destination}
	SignedUriDestination *GoogleMigrationCenterAssetsExportJobSignedUriDestination `field:"optional" json:"signedUriDestination" yaml:"signedUriDestination"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#timeouts GoogleMigrationCenterAssetsExportJob#timeouts}
	Timeouts *GoogleMigrationCenterAssetsExportJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

