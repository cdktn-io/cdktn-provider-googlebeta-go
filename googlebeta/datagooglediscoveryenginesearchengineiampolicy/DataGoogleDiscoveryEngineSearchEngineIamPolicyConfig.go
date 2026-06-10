// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglediscoveryenginesearchengineiampolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleDiscoveryEngineSearchEngineIamPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/data-sources/google_discovery_engine_search_engine_iam_policy#collection_id DataGoogleDiscoveryEngineSearchEngineIamPolicy#collection_id}.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/data-sources/google_discovery_engine_search_engine_iam_policy#engine_id DataGoogleDiscoveryEngineSearchEngineIamPolicy#engine_id}.
	EngineId *string `field:"required" json:"engineId" yaml:"engineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/data-sources/google_discovery_engine_search_engine_iam_policy#id DataGoogleDiscoveryEngineSearchEngineIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/data-sources/google_discovery_engine_search_engine_iam_policy#location DataGoogleDiscoveryEngineSearchEngineIamPolicy#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/data-sources/google_discovery_engine_search_engine_iam_policy#project DataGoogleDiscoveryEngineSearchEngineIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

