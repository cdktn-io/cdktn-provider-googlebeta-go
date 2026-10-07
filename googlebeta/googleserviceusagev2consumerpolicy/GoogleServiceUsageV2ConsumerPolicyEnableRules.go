// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleserviceusagev2consumerpolicy


type GoogleServiceUsageV2ConsumerPolicyEnableRules struct {
	// (Optional): List of service names to be enabled in the format of services/<service_name>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_service_usage_v2_consumer_policy#services GoogleServiceUsageV2ConsumerPolicy#services}
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
}

