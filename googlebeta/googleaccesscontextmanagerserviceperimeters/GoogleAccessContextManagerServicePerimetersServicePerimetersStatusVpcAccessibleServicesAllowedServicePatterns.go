// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatterns struct {
	// modifiers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_access_context_manager_service_perimeters#modifiers GoogleAccessContextManagerServicePerimeters#modifiers}
	Modifiers interface{} `field:"optional" json:"modifiers" yaml:"modifiers"`
	// URL pattern to allow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_access_context_manager_service_perimeters#pattern GoogleAccessContextManagerServicePerimeters#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Supported service to allow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_access_context_manager_service_perimeters#service GoogleAccessContextManagerServicePerimeters#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

