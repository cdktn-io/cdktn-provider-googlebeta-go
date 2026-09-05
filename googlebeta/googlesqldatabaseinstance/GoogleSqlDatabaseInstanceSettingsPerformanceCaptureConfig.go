// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfig struct {
	// The minimum percentage of CPU utilization that triggers the performance capture.
	//
	// Valid range is 10 to 99. 0 disables the check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#cpu_utilization_threshold_percent GoogleSqlDatabaseInstance#cpu_utilization_threshold_percent}
	CpuUtilizationThresholdPercent *float64 `field:"optional" json:"cpuUtilizationThresholdPercent" yaml:"cpuUtilizationThresholdPercent"`
	// Enable or disable the Performance Capture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#enabled GoogleSqlDatabaseInstance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The minimum number of undo log entries in the history list length that triggers the performance capture.
	//
	// Valid range is 10000 to 10000000. 0 disables the check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#history_list_length_threshold_count GoogleSqlDatabaseInstance#history_list_length_threshold_count}
	HistoryListLengthThresholdCount *float64 `field:"optional" json:"historyListLengthThresholdCount" yaml:"historyListLengthThresholdCount"`
	// The minimum percentage of memory usage that triggers the performance capture.
	//
	// Valid range is 10 to 99. 0 disables the check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#memory_usage_threshold_percent GoogleSqlDatabaseInstance#memory_usage_threshold_percent}
	MemoryUsageThresholdPercent *float64 `field:"optional" json:"memoryUsageThresholdPercent" yaml:"memoryUsageThresholdPercent"`
	// The minimum number of consecutive readings above threshold that triggers instance state capture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#probe_threshold GoogleSqlDatabaseInstance#probe_threshold}
	ProbeThreshold *float64 `field:"optional" json:"probeThreshold" yaml:"probeThreshold"`
	// The time interval in seconds between any two probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#probing_interval_seconds GoogleSqlDatabaseInstance#probing_interval_seconds}
	ProbingIntervalSeconds *float64 `field:"optional" json:"probingIntervalSeconds" yaml:"probingIntervalSeconds"`
	// The minimum number of server threads running to trigger the capture on primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#running_threads_threshold GoogleSqlDatabaseInstance#running_threads_threshold}
	RunningThreadsThreshold *float64 `field:"optional" json:"runningThreadsThreshold" yaml:"runningThreadsThreshold"`
	// The minimum number of seconds replica must be lagging behind primary to trigger capture on replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#seconds_behind_source_threshold GoogleSqlDatabaseInstance#seconds_behind_source_threshold}
	SecondsBehindSourceThreshold *float64 `field:"optional" json:"secondsBehindSourceThreshold" yaml:"secondsBehindSourceThreshold"`
	// The minimum number of semaphore waits that triggers the performance capture.
	//
	// Valid range is 10 to 10000. 0 disables the check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#semaphore_wait_threshold_count GoogleSqlDatabaseInstance#semaphore_wait_threshold_count}
	SemaphoreWaitThresholdCount *float64 `field:"optional" json:"semaphoreWaitThresholdCount" yaml:"semaphoreWaitThresholdCount"`
	// The amount of time in seconds that a transaction needs to have been open before getting recorded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#transaction_duration_threshold GoogleSqlDatabaseInstance#transaction_duration_threshold}
	TransactionDurationThreshold *float64 `field:"optional" json:"transactionDurationThreshold" yaml:"transactionDurationThreshold"`
	// A list of users to exclude from transaction termination. Entries can be in the format 'user@host' or just 'user'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#transaction_kill_excluded_user_hosts GoogleSqlDatabaseInstance#transaction_kill_excluded_user_hosts}
	TransactionKillExcludedUserHosts *[]*string `field:"optional" json:"transactionKillExcludedUserHosts" yaml:"transactionKillExcludedUserHosts"`
	// The amount of time in seconds that a transaction needs to have been open before the watcher starts terminating it.
	//
	// Valid range is 60 to 604800. 0 disables termination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#transaction_kill_threshold_seconds GoogleSqlDatabaseInstance#transaction_kill_threshold_seconds}
	TransactionKillThresholdSeconds *float64 `field:"optional" json:"transactionKillThresholdSeconds" yaml:"transactionKillThresholdSeconds"`
	// Determines which transactions are allowed to be terminated when they exceed transaction_kill_threshold_seconds. Possible values are: "TRANSACTION_KILL_TYPE_UNSPECIFIED", "READ_ONLY_TRANSACTIONS", "ALL_TRANSACTIONS".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#transaction_kill_type GoogleSqlDatabaseInstance#transaction_kill_type}
	TransactionKillType *string `field:"optional" json:"transactionKillType" yaml:"transactionKillType"`
	// The minimum number of transactions in lock wait state that triggers the performance capture.
	//
	// Valid range is 10 to 10000. 0 disables the check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_sql_database_instance#transaction_lock_wait_threshold_count GoogleSqlDatabaseInstance#transaction_lock_wait_threshold_count}
	TransactionLockWaitThresholdCount *float64 `field:"optional" json:"transactionLockWaitThresholdCount" yaml:"transactionLockWaitThresholdCount"`
}

