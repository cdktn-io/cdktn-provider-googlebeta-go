// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemigrationcenterreportconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleMigrationCenterReportConfigAConfig struct {
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
	// group_preferenceset_assignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_migration_center_report_config#group_preferenceset_assignments GoogleMigrationCenterReportConfigA#group_preferenceset_assignments}
	GroupPreferencesetAssignments interface{} `field:"required" json:"groupPreferencesetAssignments" yaml:"groupPreferencesetAssignments"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_migration_center_report_config#location GoogleMigrationCenterReportConfigA#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// User specified ID for the report config.
	//
	// It will become the last component
	// of the report config name. The ID must be unique within the project, must
	// conform with RFC-1034, is restricted to lower-cased letters, and has a
	// maximum length of 63 characters. The ID must match the regular expression:
	// [a-z]([a-z0-9-]{0,61}[a-z0-9])?.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_migration_center_report_config#report_config_id GoogleMigrationCenterReportConfigA#report_config_id}
	ReportConfigId *string `field:"required" json:"reportConfigId" yaml:"reportConfigId"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_migration_center_report_config#deletion_policy GoogleMigrationCenterReportConfigA#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Free-text description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_migration_center_report_config#description GoogleMigrationCenterReportConfigA#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-friendly display name. Maximum length is 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_migration_center_report_config#display_name GoogleMigrationCenterReportConfigA#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_migration_center_report_config#id GoogleMigrationCenterReportConfigA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_migration_center_report_config#project GoogleMigrationCenterReportConfigA#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_migration_center_report_config#timeouts GoogleMigrationCenterReportConfigA#timeouts}
	Timeouts *GoogleMigrationCenterReportConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

