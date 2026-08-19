// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool


type GoogleCesToolWidgetTool struct {
	// Required. The display name of the widget tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#name GoogleCesTool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// data_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#data_mapping GoogleCesTool#data_mapping}
	DataMapping *GoogleCesToolWidgetToolDataMapping `field:"optional" json:"dataMapping" yaml:"dataMapping"`
	// Optional. The description of the widget tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#description GoogleCesTool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#parameters GoogleCesTool#parameters}
	Parameters *GoogleCesToolWidgetToolParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// text_response_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#text_response_config GoogleCesTool#text_response_config}
	TextResponseConfig *GoogleCesToolWidgetToolTextResponseConfig `field:"optional" json:"textResponseConfig" yaml:"textResponseConfig"`
	// Optional. Configuration for rendering the widget. Represents a JSON object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#ui_config GoogleCesTool#ui_config}
	UiConfig *string `field:"optional" json:"uiConfig" yaml:"uiConfig"`
	// Optional.
	//
	// The type of the widget tool. If not specified, the default type will be CUSTOMIZED.
	// Possible values:
	// WIDGET_TYPE_UNSPECIFIED
	// CUSTOM
	// PRODUCT_CAROUSEL
	// PRODUCT_DETAILS
	// QUICK_ACTIONS
	// PRODUCT_COMPARISON
	// ADVANCED_PRODUCT_DETAILS
	// SHORT_FORM
	// OVERALL_SATISFACTION
	// ORDER_SUMMARY
	// APPOINTMENT_DETAILS
	// APPOINTMENT_SCHEDULER
	// CONTACT_FORM Possible values: ["WIDGET_TYPE_UNSPECIFIED", "CUSTOM", "PRODUCT_CAROUSEL", "PRODUCT_DETAILS", "QUICK_ACTIONS", "PRODUCT_COMPARISON", "ADVANCED_PRODUCT_DETAILS", "SHORT_FORM", "OVERALL_SATISFACTION", "ORDER_SUMMARY", "APPOINTMENT_DETAILS", "APPOINTMENT_SCHEDULER", "CONTACT_FORM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool#widget_type GoogleCesTool#widget_type}
	WidgetType *string `field:"optional" json:"widgetType" yaml:"widgetType"`
}

