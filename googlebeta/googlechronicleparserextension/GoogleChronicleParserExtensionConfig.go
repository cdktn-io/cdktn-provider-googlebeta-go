// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparserextension

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleParserExtensionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#instance GoogleChronicleParserExtension#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#location GoogleChronicleParserExtension#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#log_type GoogleChronicleParserExtension#log_type}
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// Parser config could be a cbn snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#cbn_snippet GoogleChronicleParserExtension#cbn_snippet}
	CbnSnippet *string `field:"optional" json:"cbnSnippet" yaml:"cbnSnippet"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#deletion_policy GoogleChronicleParserExtension#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// dynamic_parsing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#dynamic_parsing GoogleChronicleParserExtension#dynamic_parsing}
	DynamicParsing *GoogleChronicleParserExtensionDynamicParsing `field:"optional" json:"dynamicParsing" yaml:"dynamicParsing"`
	// field_extractors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#field_extractors GoogleChronicleParserExtension#field_extractors}
	FieldExtractors *GoogleChronicleParserExtensionFieldExtractors `field:"optional" json:"fieldExtractors" yaml:"fieldExtractors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#id GoogleChronicleParserExtension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Raw log used to assist the user in creation of augmentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#log GoogleChronicleParserExtension#log}
	Log *string `field:"optional" json:"log" yaml:"log"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#project GoogleChronicleParserExtension#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#timeouts GoogleChronicleParserExtension#timeouts}
	Timeouts *GoogleChronicleParserExtensionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Flag to bypass parser extension validation.
	//
	// If enabled, the parser extension won't be rejected during the validation
	// phase and validation will be skipped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_chronicle_parser_extension#validation_skipped GoogleChronicleParserExtension#validation_skipped}
	ValidationSkipped interface{} `field:"optional" json:"validationSkipped" yaml:"validationSkipped"`
}

