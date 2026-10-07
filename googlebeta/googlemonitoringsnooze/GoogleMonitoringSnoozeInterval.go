// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemonitoringsnooze


type GoogleMonitoringSnoozeInterval struct {
	// The end of the time interval.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
	// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and
	// "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_monitoring_snooze#end_time GoogleMonitoringSnooze#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// The beginning of the time interval.
	//
	// The default value for the start time
	// is the end time. The start time must not be later than the end time.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
	// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and
	// "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_monitoring_snooze#start_time GoogleMonitoringSnooze#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

