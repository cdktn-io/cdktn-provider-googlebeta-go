// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfig struct {
	// Enable or disable the Performance Capture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_sql_database_instance#enabled GoogleSqlDatabaseInstance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The minimum number of consecutive readings above threshold that triggers instance state capture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_sql_database_instance#probe_threshold GoogleSqlDatabaseInstance#probe_threshold}
	ProbeThreshold *float64 `field:"optional" json:"probeThreshold" yaml:"probeThreshold"`
	// The time interval in seconds between any two probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_sql_database_instance#probing_interval_seconds GoogleSqlDatabaseInstance#probing_interval_seconds}
	ProbingIntervalSeconds *float64 `field:"optional" json:"probingIntervalSeconds" yaml:"probingIntervalSeconds"`
	// The minimum number of server threads running to trigger the capture on primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_sql_database_instance#running_threads_threshold GoogleSqlDatabaseInstance#running_threads_threshold}
	RunningThreadsThreshold *float64 `field:"optional" json:"runningThreadsThreshold" yaml:"runningThreadsThreshold"`
	// The minimum number of seconds replica must be lagging behind primary to trigger capture on replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_sql_database_instance#seconds_behind_source_threshold GoogleSqlDatabaseInstance#seconds_behind_source_threshold}
	SecondsBehindSourceThreshold *float64 `field:"optional" json:"secondsBehindSourceThreshold" yaml:"secondsBehindSourceThreshold"`
	// The amount of time in seconds that a transaction needs to have been open before getting recorded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_sql_database_instance#transaction_duration_threshold GoogleSqlDatabaseInstance#transaction_duration_threshold}
	TransactionDurationThreshold *float64 `field:"optional" json:"transactionDurationThreshold" yaml:"transactionDurationThreshold"`
}

