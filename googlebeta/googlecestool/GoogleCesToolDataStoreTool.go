// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool


type GoogleCesToolDataStoreTool struct {
	// The data store tool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#name GoogleCesTool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// boost_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#boost_specs GoogleCesTool#boost_specs}
	BoostSpecs interface{} `field:"optional" json:"boostSpecs" yaml:"boostSpecs"`
	// data_store_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#data_store_source GoogleCesTool#data_store_source}
	DataStoreSource *GoogleCesToolDataStoreToolDataStoreSource `field:"optional" json:"dataStoreSource" yaml:"dataStoreSource"`
	// The tool description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#description GoogleCesTool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// engine_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#engine_source GoogleCesTool#engine_source}
	EngineSource *GoogleCesToolDataStoreToolEngineSource `field:"optional" json:"engineSource" yaml:"engineSource"`
	// Optional. The filter parameter behavior. Possible values: FILTER_PARAMETER_BEHAVIOR_UNSPECIFIED ALWAYS_INCLUDE NEVER_INCLUDE Possible values: ["FILTER_PARAMETER_BEHAVIOR_UNSPECIFIED", "ALWAYS_INCLUDE", "NEVER_INCLUDE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#filter_parameter_behavior GoogleCesTool#filter_parameter_behavior}
	FilterParameterBehavior *string `field:"optional" json:"filterParameterBehavior" yaml:"filterParameterBehavior"`
	// Number of search results to return per query. The default value is 10. The maximum allowed value is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#max_results GoogleCesTool#max_results}
	MaxResults *float64 `field:"optional" json:"maxResults" yaml:"maxResults"`
	// modality_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#modality_configs GoogleCesTool#modality_configs}
	ModalityConfigs interface{} `field:"optional" json:"modalityConfigs" yaml:"modalityConfigs"`
}

