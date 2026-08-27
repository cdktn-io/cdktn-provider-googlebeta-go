// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomposerenvironment


type GoogleComposerEnvironmentConfigNodeConfigTrafficRoutingConfig struct {
	// Traffic routing mode for Cloud Run functions. Possible values: ["DIRECT", "VIA_NETWORK_ATTACHMENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_composer_environment#cloud_run_functions_routing GoogleComposerEnvironment#cloud_run_functions_routing}
	CloudRunFunctionsRouting *string `field:"optional" json:"cloudRunFunctionsRouting" yaml:"cloudRunFunctionsRouting"`
}

