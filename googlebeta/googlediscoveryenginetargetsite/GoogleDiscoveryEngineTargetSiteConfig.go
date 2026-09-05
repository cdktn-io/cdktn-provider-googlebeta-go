// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlediscoveryenginetargetsite

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDiscoveryEngineTargetSiteConfig struct {
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
	// The unique id of the data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_target_site#data_store_id GoogleDiscoveryEngineTargetSite#data_store_id}
	DataStoreId *string `field:"required" json:"dataStoreId" yaml:"dataStoreId"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_target_site#location GoogleDiscoveryEngineTargetSite#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The user provided URI pattern from which the 'generated_uri_pattern' is generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_target_site#provided_uri_pattern GoogleDiscoveryEngineTargetSite#provided_uri_pattern}
	ProvidedUriPattern *string `field:"required" json:"providedUriPattern" yaml:"providedUriPattern"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_target_site#deletion_policy GoogleDiscoveryEngineTargetSite#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// If set to false, a uri_pattern is generated to include all pages whose address contains the provided_uri_pattern.
	//
	// If set to true, an uri_pattern
	// is generated to try to be an exact match of the provided_uri_pattern or
	// just the specific page if the provided_uri_pattern is a specific one.
	// provided_uri_pattern is always normalized to generate the URI pattern to
	// be used by the search engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_target_site#exact_match GoogleDiscoveryEngineTargetSite#exact_match}
	ExactMatch interface{} `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_target_site#id GoogleDiscoveryEngineTargetSite#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_target_site#project GoogleDiscoveryEngineTargetSite#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_target_site#timeouts GoogleDiscoveryEngineTargetSite#timeouts}
	Timeouts *GoogleDiscoveryEngineTargetSiteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The possible target site types. Possible values: ["INCLUDE", "EXCLUDE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_discovery_engine_target_site#type GoogleDiscoveryEngineTargetSite#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

