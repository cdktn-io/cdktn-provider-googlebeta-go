// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool


type GoogleCesToolWidgetToolDataMapping struct {
	// Optional.
	//
	// A map of widget input parameter fields to the corresponding output fields of the source tool.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_tool#field_mappings GoogleCesTool#field_mappings}
	FieldMappings *map[string]*string `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// Optional. The mode of the data mapping. Possible values: MODE_UNSPECIFIED FIELD_MAPPING PYTHON_SCRIPT Possible values: ["MODE_UNSPECIFIED", "FIELD_MAPPING", "PYTHON_SCRIPT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_tool#mode GoogleCesTool#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// python_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_tool#python_function GoogleCesTool#python_function}
	PythonFunction *GoogleCesToolWidgetToolDataMappingPythonFunction `field:"optional" json:"pythonFunction" yaml:"pythonFunction"`
	// Optional.
	//
	// The resource name of the tool that provides the data for the widget (e.g., a search tool or a custom function).
	// Format: projects/{project}/locations/{location}/agents/{agent}/tools/{tool}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_tool#source_tool_name GoogleCesTool#source_tool_name}
	SourceToolName *string `field:"optional" json:"sourceToolName" yaml:"sourceToolName"`
}

