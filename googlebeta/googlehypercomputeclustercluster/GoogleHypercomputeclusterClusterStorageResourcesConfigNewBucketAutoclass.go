// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass struct {
	// Enables Auto-class feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_hypercomputecluster_cluster#enabled GoogleHypercomputeclusterCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Terminal storage class of the autoclass bucket Possible values: NEARLINE ARCHIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_hypercomputecluster_cluster#terminal_storage_class GoogleHypercomputeclusterCluster#terminal_storage_class}
	TerminalStorageClass *string `field:"optional" json:"terminalStorageClass" yaml:"terminalStorageClass"`
}

