// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataplexmetadatafeed


type GoogleDataplexMetadataFeedFilters struct {
	// The aspect types that you want to listen to.
	//
	// Depending on how the aspect is attached to the entry, in the format: projects/{project_id_or_number}/locations/{location}/aspectTypes/{aspect_type_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_dataplex_metadata_feed#aspect_types GoogleDataplexMetadataFeed#aspect_types}
	AspectTypes *[]*string `field:"optional" json:"aspectTypes" yaml:"aspectTypes"`
	// The type of change that you want to listen to. If not specified, all changes are published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_dataplex_metadata_feed#change_types GoogleDataplexMetadataFeed#change_types}
	ChangeTypes *[]*string `field:"optional" json:"changeTypes" yaml:"changeTypes"`
	// The entry types that you want to listen to, specified as relative resource names in the format projects/{project_id_or_number}/locations/{location}/entryTypes/{entry_type_id}.
	//
	// Only entries that belong to the specified entry types are published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_dataplex_metadata_feed#entry_types GoogleDataplexMetadataFeed#entry_types}
	EntryTypes *[]*string `field:"optional" json:"entryTypes" yaml:"entryTypes"`
}

