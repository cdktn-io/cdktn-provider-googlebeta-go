// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevectorsearchindex

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVectorSearchIndexConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#collection_id GoogleVectorSearchIndex#collection_id}
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// The collection schema field to index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#index_field GoogleVectorSearchIndex#index_field}
	IndexField *string `field:"required" json:"indexField" yaml:"indexField"`
	// ID of the Index to create.
	//
	// The id must be 1-63 characters long, and comply with
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	// Specifically, it must be 1-63 characters long and match the regular
	// expression '[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#index_id GoogleVectorSearchIndex#index_id}
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#location GoogleVectorSearchIndex#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// dedicated_infrastructure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#dedicated_infrastructure GoogleVectorSearchIndex#dedicated_infrastructure}
	DedicatedInfrastructure *GoogleVectorSearchIndexDedicatedInfrastructure `field:"optional" json:"dedicatedInfrastructure" yaml:"dedicatedInfrastructure"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#deletion_policy GoogleVectorSearchIndex#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// dense_scann block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#dense_scann GoogleVectorSearchIndex#dense_scann}
	DenseScann *GoogleVectorSearchIndexDenseScann `field:"optional" json:"denseScann" yaml:"denseScann"`
	// User-specified description of the index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#description GoogleVectorSearchIndex#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-specified display name of the index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#display_name GoogleVectorSearchIndex#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Distance metric used for indexing. If not specified, will default to 'DOT_PRODUCT'. Possible values: ["DOT_PRODUCT", "COSINE_DISTANCE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#distance_metric GoogleVectorSearchIndex#distance_metric}
	DistanceMetric *string `field:"optional" json:"distanceMetric" yaml:"distanceMetric"`
	// The fields to push into the index to enable fast ANN inline filtering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#filter_fields GoogleVectorSearchIndex#filter_fields}
	FilterFields *[]*string `field:"optional" json:"filterFields" yaml:"filterFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#id GoogleVectorSearchIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#labels GoogleVectorSearchIndex#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#project GoogleVectorSearchIndex#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The fields to push into the index to enable inline data retrieval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#store_fields GoogleVectorSearchIndex#store_fields}
	StoreFields *[]*string `field:"optional" json:"storeFields" yaml:"storeFields"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vector_search_index#timeouts GoogleVectorSearchIndex#timeouts}
	Timeouts *GoogleVectorSearchIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

