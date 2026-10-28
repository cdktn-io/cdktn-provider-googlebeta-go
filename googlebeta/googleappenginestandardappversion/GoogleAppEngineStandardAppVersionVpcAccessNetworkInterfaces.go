// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleappenginestandardappversion


type GoogleAppEngineStandardAppVersionVpcAccessNetworkInterfaces struct {
	// The VPC network that the App Engine resource will be able to send traffic to.
	//
	// At least one of network or subnetwork must be specified. If both network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If network is not specified, it will be looked up from the subnetwork. Could be either a short name or a full path (e.g. 'default' or 'projects/my-project/global/networks/default').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_app_engine_standard_app_version#network GoogleAppEngineStandardAppVersion#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The VPC subnetwork that the App Engine resource will get IPs from.
	//
	// At least one of network or subnetwork must be specified. If both network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If subnetwork is not specified, the subnetwork with the same name with the network will be used. Could be either a short name or a full path (e.g. 'default' or 'projects/my-project/regions/us-central1/subnetworks/default').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_app_engine_standard_app_version#subnetwork GoogleAppEngineStandardAppVersion#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// Network tags applied to this App Engine version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_app_engine_standard_app_version#tags GoogleAppEngineStandardAppVersion#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

