// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyLoggingConfigsLogToBigQuery struct {
	// The dataset ID of the BigQuery table to log to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#dataset_id GoogleDataLossPreventionContentPolicy#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The project ID of the BigQuery table to log to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#project_id GoogleDataLossPreventionContentPolicy#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The table ID of the BigQuery table to log to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#table_id GoogleDataLossPreventionContentPolicy#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

