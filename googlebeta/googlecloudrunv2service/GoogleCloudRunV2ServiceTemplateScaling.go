// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateScaling struct {
	// Determines a threshold for concurrency utilization before scaling begins.
	//
	// Accepted values are between 0.1 and 0.95 (inclusive) or 0.0 to disable concurrency utilization as threshold for scaling. CPU and concurrency scaling cannot both be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_cloud_run_v2_service#concurrency_utilization GoogleCloudRunV2Service#concurrency_utilization}
	ConcurrencyUtilization *float64 `field:"optional" json:"concurrencyUtilization" yaml:"concurrencyUtilization"`
	// Determines a threshold for CPU utilization before scaling begins.
	//
	// Accepted values are between 0.1 and 0.95 (inclusive) or 0.0 to disable CPU utilization as threshold for scaling. CPU and concurrency scaling cannot both be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_cloud_run_v2_service#cpu_utilization GoogleCloudRunV2Service#cpu_utilization}
	CpuUtilization *float64 `field:"optional" json:"cpuUtilization" yaml:"cpuUtilization"`
	// Maximum number of serving instances that this resource should have.
	//
	// Must not be less than minimum instance count. If absent, Cloud Run will calculate
	// a default value based on the project's available container instances quota in the region and specified instance size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_cloud_run_v2_service#max_instance_count GoogleCloudRunV2Service#max_instance_count}
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Minimum number of serving instances that this resource should have.
	//
	// Defaults to 0. Must not be greater than maximum instance count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_cloud_run_v2_service#min_instance_count GoogleCloudRunV2Service#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
}

