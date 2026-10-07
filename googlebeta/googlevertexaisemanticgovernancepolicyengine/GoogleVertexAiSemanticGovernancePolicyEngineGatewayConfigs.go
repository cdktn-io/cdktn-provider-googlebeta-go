// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaisemanticgovernancepolicyengine


type GoogleVertexAiSemanticGovernancePolicyEngineGatewayConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_semantic_governance_policy_engine#name GoogleVertexAiSemanticGovernancePolicyEngine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Additional consumer projects permitted to attach their own PSC endpoint to this gateway's ServiceAttachment.
	//
	// This is the "decoupled" mode, where
	// the customer creates the PSC endpoint in a project other than this
	// gateway's network project. Each listed project is VPC-SC enforced: it
	// must be within the caller's service perimeter. The owning
	// SemanticGovernancePolicyEngine's own project is always permitted
	// implicitly and need not be listed. Format: projects/{project} (ID or number).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_semantic_governance_policy_engine#allowed_projects GoogleVertexAiSemanticGovernancePolicyEngine#allowed_projects}
	AllowedProjects *[]*string `field:"optional" json:"allowedProjects" yaml:"allowedProjects"`
	// The name of the private Cloud DNS managed zone in which the backend creates the DNS record set for this gateway's PSC endpoint.
	//
	// This is the
	// managed-zone resource name, not a fully-qualified domain name. The zone
	// must already exist and be attached to the gateway's VPC at provision
	// time. The name must match '^[a-z0-9.-]{1,63}$'. Must be set together
	// with 'network' and 'subnetwork' (all three or none).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_semantic_governance_policy_engine#dns_zone_name GoogleVertexAiSemanticGovernancePolicyEngine#dns_zone_name}
	DnsZoneName *string `field:"optional" json:"dnsZoneName" yaml:"dnsZoneName"`
	// The URI of the network resource where the gateway's PSC endpoint is provisioned.
	//
	// Format: projects/{project}/global/networks/{network}.
	// 'network', 'subnetwork', and 'dns_zone_name' must all be set together
	// or all omitted; setting only some is rejected by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_semantic_governance_policy_engine#network GoogleVertexAiSemanticGovernancePolicyEngine#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The URI of the subnetwork resource where the gateway's PSC endpoint is provisioned.
	//
	// Format:
	// projects/{project}/regions/{region}/subnetworks/{subnetwork}. Must be
	// set together with 'network' and 'dns_zone_name' (all three or none).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_semantic_governance_policy_engine#subnetwork GoogleVertexAiSemanticGovernancePolicyEngine#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

