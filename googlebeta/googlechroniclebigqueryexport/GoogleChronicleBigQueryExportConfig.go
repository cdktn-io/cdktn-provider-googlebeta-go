// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclebigqueryexport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleBigQueryExportConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#instance GoogleChronicleBigQueryExport#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#location GoogleChronicleBigQueryExport#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The BigQueryExportPackage entitled for the Chronicle instance. Possible values: ["BIG_QUERY_EXPORT_PACKAGE_BYOBQ", "BIG_QUERY_EXPORT_PACKAGE_ADVANCED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#big_query_export_package GoogleChronicleBigQueryExport#big_query_export_package}
	BigQueryExportPackage *string `field:"optional" json:"bigQueryExportPackage" yaml:"bigQueryExportPackage"`
	// entity_graph_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#entity_graph_settings GoogleChronicleBigQueryExport#entity_graph_settings}
	EntityGraphSettings *GoogleChronicleBigQueryExportEntityGraphSettings `field:"optional" json:"entityGraphSettings" yaml:"entityGraphSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#id GoogleChronicleBigQueryExport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ioc_matches_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#ioc_matches_settings GoogleChronicleBigQueryExport#ioc_matches_settings}
	IocMatchesSettings *GoogleChronicleBigQueryExportIocMatchesSettings `field:"optional" json:"iocMatchesSettings" yaml:"iocMatchesSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#project GoogleChronicleBigQueryExport#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rule_detections_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#rule_detections_settings GoogleChronicleBigQueryExport#rule_detections_settings}
	RuleDetectionsSettings *GoogleChronicleBigQueryExportRuleDetectionsSettings `field:"optional" json:"ruleDetectionsSettings" yaml:"ruleDetectionsSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#timeouts GoogleChronicleBigQueryExport#timeouts}
	Timeouts *GoogleChronicleBigQueryExportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// udm_events_aggregates_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#udm_events_aggregates_settings GoogleChronicleBigQueryExport#udm_events_aggregates_settings}
	UdmEventsAggregatesSettings *GoogleChronicleBigQueryExportUdmEventsAggregatesSettings `field:"optional" json:"udmEventsAggregatesSettings" yaml:"udmEventsAggregatesSettings"`
	// udm_events_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_big_query_export#udm_events_settings GoogleChronicleBigQueryExport#udm_events_settings}
	UdmEventsSettings *GoogleChronicleBigQueryExportUdmEventsSettings `field:"optional" json:"udmEventsSettings" yaml:"udmEventsSettings"`
}

