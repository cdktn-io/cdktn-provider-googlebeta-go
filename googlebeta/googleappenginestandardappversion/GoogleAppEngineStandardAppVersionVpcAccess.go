// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleappenginestandardappversion


type GoogleAppEngineStandardAppVersionVpcAccess struct {
	// The egress setting for the VPC Access, controlling what traffic is diverted through it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_app_engine_standard_app_version#egress_setting GoogleAppEngineStandardAppVersion#egress_setting}
	EgressSetting *string `field:"optional" json:"egressSetting" yaml:"egressSetting"`
	// network_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_app_engine_standard_app_version#network_interfaces GoogleAppEngineStandardAppVersion#network_interfaces}
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The traffic egress setting for the VPC network interface, controlling what traffic is diverted through it.
	//
	// Possible values: ["ALL_TRAFFIC", "PRIVATE_IP_RANGES"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_app_engine_standard_app_version#vpc_egress GoogleAppEngineStandardAppVersion#vpc_egress}
	VpcEgress *string `field:"optional" json:"vpcEgress" yaml:"vpcEgress"`
}

