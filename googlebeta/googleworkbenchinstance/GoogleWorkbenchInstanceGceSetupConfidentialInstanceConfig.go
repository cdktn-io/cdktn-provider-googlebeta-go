// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleworkbenchinstance


type GoogleWorkbenchInstanceGceSetupConfidentialInstanceConfig struct {
	// Defines the type of technology used by the confidential instance. Possible values: ["SEV"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_workbench_instance#confidential_instance_type GoogleWorkbenchInstance#confidential_instance_type}
	ConfidentialInstanceType *string `field:"optional" json:"confidentialInstanceType" yaml:"confidentialInstanceType"`
}

