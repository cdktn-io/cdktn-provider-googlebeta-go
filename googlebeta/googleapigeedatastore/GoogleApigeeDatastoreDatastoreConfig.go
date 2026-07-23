// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleapigeedatastore


type GoogleApigeeDatastoreDatastoreConfig struct {
	// The GCP project ID that the datastore target resides in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_apigee_datastore#project_id GoogleApigeeDatastore#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The name of the Cloud Storage bucket. Required for 'gcs' target type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_apigee_datastore#bucket_name GoogleApigeeDatastore#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The name of the BigQuery dataset. Required for 'bigquery' target type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_apigee_datastore#dataset_name GoogleApigeeDatastore#dataset_name}
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// The path within the Cloud Storage bucket. Used for 'gcs' target type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_apigee_datastore#path GoogleApigeeDatastore#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The prefix for BigQuery table names. Used for 'bigquery' target type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_apigee_datastore#table_prefix GoogleApigeeDatastore#table_prefix}
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
}

