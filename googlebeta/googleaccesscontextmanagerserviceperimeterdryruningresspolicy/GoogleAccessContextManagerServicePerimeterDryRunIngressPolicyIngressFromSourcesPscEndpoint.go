// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagerserviceperimeterdryruningresspolicy


type GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromSourcesPscEndpoint struct {
	// The full resource name of the global forwarding rule that identifies a Private Service Connect endpoint. Forwarding rule format: '//compute.googleapis.com/projects/{PROJECT_ID}/global/forwardingRules/{FORWARDING_RULE_ID}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_access_context_manager_service_perimeter_dry_run_ingress_policy#forwarding_rule GoogleAccessContextManagerServicePerimeterDryRunIngressPolicy#forwarding_rule}
	ForwardingRule *string `field:"optional" json:"forwardingRule" yaml:"forwardingRule"`
}

