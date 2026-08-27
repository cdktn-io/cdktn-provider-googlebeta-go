// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlefirestorechangestream

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleFirestoreChangeStreamConfig struct {
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
	// The ID to use for the change stream, which will become the final component of the change stream's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_firestore_change_stream#name GoogleFirestoreChangeStream#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The duration for which change stream data is retained.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "86400s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_firestore_change_stream#retention_period GoogleFirestoreChangeStream#retention_period}
	RetentionPeriod *string `field:"required" json:"retentionPeriod" yaml:"retentionPeriod"`
	// collection_group_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_firestore_change_stream#collection_group_scope GoogleFirestoreChangeStream#collection_group_scope}
	CollectionGroupScope *GoogleFirestoreChangeStreamCollectionGroupScope `field:"optional" json:"collectionGroupScope" yaml:"collectionGroupScope"`
	// The Firestore database ID. Defaults to '"(default)"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_firestore_change_stream#database GoogleFirestoreChangeStream#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// database_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_firestore_change_stream#database_scope GoogleFirestoreChangeStream#database_scope}
	DatabaseScope *GoogleFirestoreChangeStreamDatabaseScope `field:"optional" json:"databaseScope" yaml:"databaseScope"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_firestore_change_stream#deletion_policy GoogleFirestoreChangeStream#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_firestore_change_stream#id GoogleFirestoreChangeStream#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_firestore_change_stream#project GoogleFirestoreChangeStream#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_firestore_change_stream#timeouts GoogleFirestoreChangeStream#timeouts}
	Timeouts *GoogleFirestoreChangeStreamTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

