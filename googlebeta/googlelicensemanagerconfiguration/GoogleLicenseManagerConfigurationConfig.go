// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlelicensemanagerconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleLicenseManagerConfigurationConfig struct {
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
	// Id of the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#configuration_id GoogleLicenseManagerConfiguration#configuration_id}
	ConfigurationId *string `field:"required" json:"configurationId" yaml:"configurationId"`
	// Number of units to bill for.
	//
	// When licensing a product that is billed per-user, this means number of users. When licensing a product that is billed per-pack (e.g. SQL Server), this means the number of packs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#license_count GoogleLicenseManagerConfiguration#license_count}
	LicenseCount *float64 `field:"required" json:"licenseCount" yaml:"licenseCount"`
	// The region where the configuration should be created.
	//
	// This region must be the same where the licensed software will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#location GoogleLicenseManagerConfiguration#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the product for which you are setting the license configuration.
	//
	// For supported products see https://docs.cloud.google.com/compute/docs/instances/windows/license-manager#supported-license-products. Available values include Office2021ProfessionalPlus
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#product GoogleLicenseManagerConfiguration#product}
	Product *string `field:"required" json:"product" yaml:"product"`
	// Whether the configuration is active.
	//
	// We suggest you deactivate a configuration instead of deleting it, and allow License Manager to manage deletion of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#active GoogleLicenseManagerConfiguration#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#deletion_policy GoogleLicenseManagerConfiguration#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#id GoogleLicenseManagerConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#labels GoogleLicenseManagerConfiguration#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#project GoogleLicenseManagerConfiguration#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_license_manager_configuration#timeouts GoogleLicenseManagerConfiguration#timeouts}
	Timeouts *GoogleLicenseManagerConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

