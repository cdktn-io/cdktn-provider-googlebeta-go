// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsModifiersAddRequestHeader struct {
	// HTTP header key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_access_context_manager_service_perimeters#key GoogleAccessContextManagerServicePerimeters#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// HTTP header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_access_context_manager_service_perimeters#value GoogleAccessContextManagerServicePerimeters#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

