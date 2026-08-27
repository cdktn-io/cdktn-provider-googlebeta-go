// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagerserviceperimeter


type GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressFromSourcesPscEndpoint struct {
	// The full resource name of the global forwarding rule that identifies a Private Service Connect endpoint. Forwarding rule format: '//compute.googleapis.com/projects/{PROJECT_ID}/global/forwardingRules/{FORWARDING_RULE_ID}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_access_context_manager_service_perimeter#forwarding_rule GoogleAccessContextManagerServicePerimeter#forwarding_rule}
	ForwardingRule *string `field:"optional" json:"forwardingRule" yaml:"forwardingRule"`
}

