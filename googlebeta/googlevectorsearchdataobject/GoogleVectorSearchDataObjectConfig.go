// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevectorsearchdataobject

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVectorSearchDataObjectConfig struct {
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
	// The ID of the parent Collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#collection_id GoogleVectorSearchDataObject#collection_id}
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// ID of the DataObject to create.
	//
	// The id must be 1-63 characters long, and comply with
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	// Specifically, it must be 1-63 characters long and match the regular
	// expression '[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#data_object_id GoogleVectorSearchDataObject#data_object_id}
	DataObjectId *string `field:"required" json:"dataObjectId" yaml:"dataObjectId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#location GoogleVectorSearchDataObject#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The JSON data of the DataObject.
	//
	// Must be a JSON object whose field
	// names match the fields defined in the parent Collection's
	// 'data_schema'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#data GoogleVectorSearchDataObject#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#deletion_policy GoogleVectorSearchDataObject#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The etag of the DataObject, used for optimistic concurrency control on updates and deletes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#etag GoogleVectorSearchDataObject#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#id GoogleVectorSearchDataObject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#project GoogleVectorSearchDataObject#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#timeouts GoogleVectorSearchDataObject#timeouts}
	Timeouts *GoogleVectorSearchDataObjectTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vector_search_data_object#vectors GoogleVectorSearchDataObject#vectors}
	Vectors interface{} `field:"optional" json:"vectors" yaml:"vectors"`
}

