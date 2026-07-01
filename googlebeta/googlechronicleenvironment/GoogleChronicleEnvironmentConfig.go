// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleenvironment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleEnvironmentConfig struct {
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
	// MAX_NAME_LENGTH = 256 Name of the contact for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#contact GoogleChronicleEnvironment#contact}
	Contact *string `field:"required" json:"contact" yaml:"contact"`
	// MAX_NAME_LENGTH = 256 Email of the contact for the environment. Multiple emails can be sepereated with the ';' character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#contact_emails GoogleChronicleEnvironment#contact_emails}
	ContactEmails *string `field:"required" json:"contactEmails" yaml:"contactEmails"`
	// MAX_NAME_LENGTH = 256 Phone number of the contact for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#contact_phone GoogleChronicleEnvironment#contact_phone}
	ContactPhone *string `field:"required" json:"contactPhone" yaml:"contactPhone"`
	// MAX_NAME_LENGTH = 256 Description of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#description GoogleChronicleEnvironment#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name of the environment MAX_NAME_LENGTH = 256.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#display_name GoogleChronicleEnvironment#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#instance GoogleChronicleEnvironment#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#location GoogleChronicleEnvironment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Environment data retention in months.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#retention_duration GoogleChronicleEnvironment#retention_duration}
	RetentionDuration *float64 `field:"required" json:"retentionDuration" yaml:"retentionDuration"`
	// Environment nicknames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#aliases_json GoogleChronicleEnvironment#aliases_json}
	AliasesJson *string `field:"optional" json:"aliasesJson" yaml:"aliasesJson"`
	// data access scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#data_access_scopes_json GoogleChronicleEnvironment#data_access_scopes_json}
	DataAccessScopesJson *string `field:"optional" json:"dataAccessScopesJson" yaml:"dataAccessScopesJson"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#deletion_policy GoogleChronicleEnvironment#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Whether Terraform will be prevented from destroying the environment.
	//
	// Deleting an environment will remove all its data and all playbooks, environments, integrations instances, reports and agents related to the environment. Once you delete an environment, it cannot be reversed. Deleting environments via terraform destroy or terraform apply will only succeed if this field is false in the Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#deletion_protection GoogleChronicleEnvironment#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#id GoogleChronicleEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#project GoogleChronicleEnvironment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_chronicle_environment#timeouts GoogleChronicleEnvironment#timeouts}
	Timeouts *GoogleChronicleEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

