// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataplexmetadatafeed


type GoogleDataplexMetadataFeedScope struct {
	// The entry groups whose entries you want to listen to. Must be in the format: projects/{project_id_or_number}/locations/{location_id}/entryGroups/{entry_group_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_dataplex_metadata_feed#entry_groups GoogleDataplexMetadataFeed#entry_groups}
	EntryGroups *[]*string `field:"optional" json:"entryGroups" yaml:"entryGroups"`
	// Whether the metadata feed is at the organization-level.
	//
	// If true, all changes happened to the entries in the same organization as the feed are published.
	// If false, you must specify a list of projects or a list of entry groups whose entries you want to listen to.The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_dataplex_metadata_feed#organization_level GoogleDataplexMetadataFeed#organization_level}
	OrganizationLevel interface{} `field:"optional" json:"organizationLevel" yaml:"organizationLevel"`
	// The projects whose entries you want to listen to.
	//
	// Must be in the same organization as the feed. Must be in the format: projects/{project_id_or_number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_dataplex_metadata_feed#projects GoogleDataplexMetadataFeed#projects}
	Projects *[]*string `field:"optional" json:"projects" yaml:"projects"`
}

