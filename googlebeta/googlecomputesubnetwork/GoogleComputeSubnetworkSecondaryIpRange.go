// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputesubnetwork


type GoogleComputeSubnetworkSecondaryIpRange struct {
	// The name associated with this subnetwork secondary range, used when adding an alias IP range to a VM instance.
	//
	// The name must
	// be 1-63 characters long, and comply with RFC1035. The name
	// must be unique within the subnetwork.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_compute_subnetwork#range_name GoogleComputeSubnetwork#range_name}
	RangeName *string `field:"required" json:"rangeName" yaml:"rangeName"`
	// The range of IP addresses belonging to this subnetwork secondary range.
	//
	// Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	// Field is optional when 'reserved_internal_range' is defined, otherwise required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_compute_subnetwork#ip_cidr_range GoogleComputeSubnetwork#ip_cidr_range}
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Reference to a Public Delegated Prefix (PDP) for BYOIP.
	//
	// This field should be specified for configuring BYOGUA internal IPv6 secondary range.
	// When specified along with the ip_cidr_range, the ip_cidr_range must lie within the PDP referenced by the 'ipCollection' field.
	// When specified without the ip_cidr_range, the range is auto-allocated from the PDP referenced by the 'ipCollection' field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_compute_subnetwork#ip_collection GoogleComputeSubnetwork#ip_collection}
	IpCollection *string `field:"optional" json:"ipCollection" yaml:"ipCollection"`
	// The IP version of the secondary range. If not specified, IPV4 is used. Possible values: ["IPV4", "IPV6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_compute_subnetwork#ip_version GoogleComputeSubnetwork#ip_version}
	IpVersion *string `field:"optional" json:"ipVersion" yaml:"ipVersion"`
	// The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com' E.g. 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_compute_subnetwork#reserved_internal_range GoogleComputeSubnetwork#reserved_internal_range}
	ReservedInternalRange *string `field:"optional" json:"reservedInternalRange" yaml:"reservedInternalRange"`
}

