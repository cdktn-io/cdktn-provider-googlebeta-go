// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTest struct {
	// The comparison to apply between the time series (indicated by filter and aggregation) and the threshold (indicated by threshold_value).
	//
	// The
	// comparison is applied on each time series, with
	// the time series on the left-hand side and the
	// threshold on the right-hand side.
	//
	// The Cloud Monitoring API only supports
	// 'COMPARISON_LT' and 'COMPARISON_GT' for SQL
	// row-count thresholds; the other values are kept
	// in the schema for backward compatibility with
	// imported state but will be rejected by the API.
	// See
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies#MetricThreshold. Possible values: ["COMPARISON_GT", "COMPARISON_GE", "COMPARISON_LT", "COMPARISON_LE", "COMPARISON_EQ", "COMPARISON_NE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_monitoring_alert_policy#comparison GoogleMonitoringAlertPolicy#comparison}
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The value against which to compare the row count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_monitoring_alert_policy#threshold GoogleMonitoringAlertPolicy#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

