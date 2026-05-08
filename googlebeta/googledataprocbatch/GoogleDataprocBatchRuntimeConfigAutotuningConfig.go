// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataprocbatch


type GoogleDataprocBatchRuntimeConfigAutotuningConfig struct {
	// Optional. Scenarios for which tunings are applied. Possible values: ["AUTO", "SCALING", "BROADCAST_HASH_JOIN", "MEMORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_batch#scenarios GoogleDataprocBatch#scenarios}
	Scenarios *[]*string `field:"optional" json:"scenarios" yaml:"scenarios"`
}

