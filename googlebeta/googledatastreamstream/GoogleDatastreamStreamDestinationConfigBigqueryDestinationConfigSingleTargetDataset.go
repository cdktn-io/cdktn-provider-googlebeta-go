// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatastreamstream


type GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset struct {
	// Dataset ID in the format projects/{project}/datasets/{dataset_id} or {project}:{dataset_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_datastream_stream#dataset_id GoogleDatastreamStream#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
}

