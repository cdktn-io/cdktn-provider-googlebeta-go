// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleParserConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#instance GoogleChronicleParser#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#location GoogleChronicleParser#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#logtype GoogleChronicleParser#logtype}
	Logtype *string `field:"required" json:"logtype" yaml:"logtype"`
	// if the parser is built using config documentation: https://cloud.google.com/chronicle/docs/preview/parser-extensions/parsing-overview.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#cbn GoogleChronicleParser#cbn}
	Cbn *string `field:"optional" json:"cbn" yaml:"cbn"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#deletion_policy GoogleChronicleParser#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#id GoogleChronicleParser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// low_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#low_code GoogleChronicleParser#low_code}
	LowCode *GoogleChronicleParserLowCode `field:"optional" json:"lowCode" yaml:"lowCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#project GoogleChronicleParser#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#timeouts GoogleChronicleParser#timeouts}
	Timeouts *GoogleChronicleParserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Flag to bypass parser validation when no logs are found.
	//
	// If enabled, the parser won't be be rejected during the validation
	// phase when no logs are found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#validated_on_empty_logs GoogleChronicleParser#validated_on_empty_logs}
	ValidatedOnEmptyLogs interface{} `field:"optional" json:"validatedOnEmptyLogs" yaml:"validatedOnEmptyLogs"`
	// If true, bypasses parser validation. If enabled, the parser won't be rejected during the validation phase and validation will be skipped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#validation_skipped GoogleChronicleParser#validation_skipped}
	ValidationSkipped interface{} `field:"optional" json:"validationSkipped" yaml:"validationSkipped"`
	// version_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_parser#version_info GoogleChronicleParser#version_info}
	VersionInfo *GoogleChronicleParserVersionInfo `field:"optional" json:"versionInfo" yaml:"versionInfo"`
}

