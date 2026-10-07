// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlediscoveryenginedataconnector


type GoogleDiscoveryEngineDataConnectorMetadata struct {
	// The party that authored the connector, e.g. "Google" or a third-party provider name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_discovery_engine_data_connector#author GoogleDiscoveryEngineDataConnector#author}
	Author *string `field:"optional" json:"author" yaml:"author"`
	// Human-readable description of the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_discovery_engine_data_connector#description GoogleDiscoveryEngineDataConnector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Free-form, multi-line note about the connector's capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_discovery_engine_data_connector#note GoogleDiscoveryEngineDataConnector#note}
	Note *string `field:"optional" json:"note" yaml:"note"`
	// Short, subtitle-length description of the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_discovery_engine_data_connector#short_description GoogleDiscoveryEngineDataConnector#short_description}
	ShortDescription *string `field:"optional" json:"shortDescription" yaml:"shortDescription"`
	// Display title of the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_discovery_engine_data_connector#title GoogleDiscoveryEngineDataConnector#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

