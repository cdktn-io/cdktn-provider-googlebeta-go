// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleappenginestandardappversion


type GoogleAppEngineStandardAppVersionVpcAccessNetworkInterfaces struct {
	// The name of the VPC network to which the version connects (e.g. 'projects/my-project/global/networks/default').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_app_engine_standard_app_version#network GoogleAppEngineStandardAppVersion#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The name of the subnetwork to which the version connects (e.g. 'projects/my-project/regions/us-central1/subnetworks/default').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_app_engine_standard_app_version#subnetwork GoogleAppEngineStandardAppVersion#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// Network tags applied to this App Engine version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_app_engine_standard_app_version#tags GoogleAppEngineStandardAppVersion#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

