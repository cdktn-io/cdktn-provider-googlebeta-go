// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googletagstagbindingcollection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleTagsTagBindingCollectionConfig struct {
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
	// The full resource name of the resource to which the tags are bound. E.g. //cloudresourcemanager.googleapis.com/projects/123.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_tags_tag_binding_collection#full_resource_name GoogleTagsTagBindingCollection#full_resource_name}
	FullResourceName *string `field:"required" json:"fullResourceName" yaml:"fullResourceName"`
	// A map of tag keys to values directly bound to this resource, specified in namespaced format.
	//
	// For example:
	//   "123/environment": "production"
	// Keys must be namespaced names of TagKeys, and values must be short names of TagValues.
	// This field is non-authoritative. Terraform will only manage the precise tags present in this map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_tags_tag_binding_collection#tags GoogleTagsTagBindingCollection#tags}
	Tags *map[string]*string `field:"required" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_tags_tag_binding_collection#id GoogleTagsTagBindingCollection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the TagBindingCollection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_tags_tag_binding_collection#location GoogleTagsTagBindingCollection#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_tags_tag_binding_collection#timeouts GoogleTagsTagBindingCollection#timeouts}
	Timeouts *GoogleTagsTagBindingCollectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

