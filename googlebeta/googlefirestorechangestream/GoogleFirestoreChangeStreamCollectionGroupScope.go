// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlefirestorechangestream


type GoogleFirestoreChangeStreamCollectionGroupScope struct {
	// The ID of the collection group to track.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_firestore_change_stream#collection_group_id GoogleFirestoreChangeStream#collection_group_id}
	CollectionGroupId *string `field:"required" json:"collectionGroupId" yaml:"collectionGroupId"`
}

