// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleapigeednszone

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleApigeeDnsZoneConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Description for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_apigee_dns_zone#description GoogleApigeeDnsZone#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// ID of the dns zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_apigee_dns_zone#dns_zone_id GoogleApigeeDnsZone#dns_zone_id}
	DnsZoneId *string `field:"required" json:"dnsZoneId" yaml:"dnsZoneId"`
	// Doamin for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_apigee_dns_zone#domain GoogleApigeeDnsZone#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_apigee_dns_zone#org_id GoogleApigeeDnsZone#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_apigee_dns_zone#peering_config GoogleApigeeDnsZone#peering_config}
	PeeringConfig *GoogleApigeeDnsZonePeeringConfig `field:"required" json:"peeringConfig" yaml:"peeringConfig"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_apigee_dns_zone#deletion_policy GoogleApigeeDnsZone#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_apigee_dns_zone#id GoogleApigeeDnsZone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_apigee_dns_zone#timeouts GoogleApigeeDnsZone#timeouts}
	Timeouts *GoogleApigeeDnsZoneTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

