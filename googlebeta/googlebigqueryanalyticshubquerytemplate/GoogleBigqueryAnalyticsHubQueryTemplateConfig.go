// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshubquerytemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleBigqueryAnalyticsHubQueryTemplateConfig struct {
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
	// The ID of the data exchange.
	//
	// Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#data_exchange_id GoogleBigqueryAnalyticsHubQueryTemplate#data_exchange_id}
	DataExchangeId *string `field:"required" json:"dataExchangeId" yaml:"dataExchangeId"`
	// Human-readable display name of the QueryTemplate.
	//
	// The display name must
	// contain only Unicode letters, numbers (0-9), underscores (_), dashes (-),
	// spaces ( ), ampersands (&) and can't start or end with spaces. Default
	// value is an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#display_name GoogleBigqueryAnalyticsHubQueryTemplate#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the location this data exchange query template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#location GoogleBigqueryAnalyticsHubQueryTemplate#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Unique QueryTemplate ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#query_template_id GoogleBigqueryAnalyticsHubQueryTemplate#query_template_id}
	QueryTemplateId *string `field:"required" json:"queryTemplateId" yaml:"queryTemplateId"`
	// This field uses a custom implementation please refer to documentation under /hashicorp/terraform-provider-google-beta/website/docs/r/bigquery_analytics_hub_query_template.html.markdown for specifics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#deletion_policy GoogleBigqueryAnalyticsHubQueryTemplate#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Short description of the QueryTemplate.
	//
	// The description must not contain
	// Unicode non-characters and C0 and C1 control codes except tabs,
	// new lines, carriage returns, and page breaks.
	// Default value is an empty string. Max length: 2000 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#description GoogleBigqueryAnalyticsHubQueryTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Documentation describing the QueryTemplate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#documentation GoogleBigqueryAnalyticsHubQueryTemplate#documentation}
	Documentation *string `field:"optional" json:"documentation" yaml:"documentation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#id GoogleBigqueryAnalyticsHubQueryTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Email or URL of the primary point of contact of the QueryTemplate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#primary_contact GoogleBigqueryAnalyticsHubQueryTemplate#primary_contact}
	PrimaryContact *string `field:"optional" json:"primaryContact" yaml:"primaryContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#project GoogleBigqueryAnalyticsHubQueryTemplate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// routine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#routine GoogleBigqueryAnalyticsHubQueryTemplate#routine}
	Routine *GoogleBigqueryAnalyticsHubQueryTemplateRoutine `field:"optional" json:"routine" yaml:"routine"`
	// If set to 'true', the QueryTemplate will be submitted for approval and cannot be updated afterwards.
	//
	// This is a one-time action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#submit GoogleBigqueryAnalyticsHubQueryTemplate#submit}
	Submit interface{} `field:"optional" json:"submit" yaml:"submit"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_bigquery_analytics_hub_query_template#timeouts GoogleBigqueryAnalyticsHubQueryTemplate#timeouts}
	Timeouts *GoogleBigqueryAnalyticsHubQueryTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

