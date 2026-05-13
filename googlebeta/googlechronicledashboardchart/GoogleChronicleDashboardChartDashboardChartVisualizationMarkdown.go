// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationMarkdown struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_chronicle_dashboard_chart#content GoogleChronicleDashboardChart#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_chronicle_dashboard_chart#properties GoogleChronicleDashboardChart#properties}
	Properties *GoogleChronicleDashboardChartDashboardChartVisualizationMarkdownProperties `field:"optional" json:"properties" yaml:"properties"`
}

