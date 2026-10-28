// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterLustreExportConfig struct {
	// The Cloud Storage bucket that backups are written to, without the "gs://" scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#destination_bucket GoogleVertexAiModelDevelopmentCluster#destination_bucket}
	DestinationBucket *string `field:"required" json:"destinationBucket" yaml:"destinationBucket"`
	// Days of the week on which a backup runs, each at most once and none of them DAY_OF_WEEK_UNSPECIFIED.
	//
	// If empty, a backup runs every day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#days_of_week GoogleVertexAiModelDevelopmentCluster#days_of_week}
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// The Managed Lustre instances to back up, each named as "projects/{project}/locations/{zone}/instances/{instance}".
	//
	// Every entry
	// must be attached to one of the cluster's node pools, and no entry may
	// repeat. If empty, every attached instance is backed up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#lustres GoogleVertexAiModelDevelopmentCluster#lustres}
	Lustres *[]*string `field:"optional" json:"lustres" yaml:"lustres"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#start_time GoogleVertexAiModelDevelopmentCluster#start_time}
	StartTime *GoogleVertexAiModelDevelopmentClusterLustreExportConfigStartTime `field:"optional" json:"startTime" yaml:"startTime"`
}

