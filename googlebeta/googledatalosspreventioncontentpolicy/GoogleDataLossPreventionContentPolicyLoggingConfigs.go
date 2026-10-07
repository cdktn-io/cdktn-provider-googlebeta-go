// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyLoggingConfigs struct {
	// log_to_big_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#log_to_big_query GoogleDataLossPreventionContentPolicy#log_to_big_query}
	LogToBigQuery *GoogleDataLossPreventionContentPolicyLoggingConfigsLogToBigQuery `field:"optional" json:"logToBigQuery" yaml:"logToBigQuery"`
}

