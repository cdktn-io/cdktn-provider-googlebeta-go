// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterManagedOpentelemetryConfig struct {
	// The scope of the Managed OpenTelemetry pipeline.
	//
	// Available options include NONE and COLLECTION_AND_INSTRUMENTATION_COMPONENTS. See https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#managedopentelemetryconfig for more information. To disable the feature, explicitly set this to NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_container_cluster#scope GoogleContainerCluster#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

