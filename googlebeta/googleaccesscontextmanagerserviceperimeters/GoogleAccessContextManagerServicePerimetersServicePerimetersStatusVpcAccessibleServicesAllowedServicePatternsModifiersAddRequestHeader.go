// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServicesAllowedServicePatternsModifiersAddRequestHeader struct {
	// HTTP header key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_access_context_manager_service_perimeters#key GoogleAccessContextManagerServicePerimeters#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// HTTP header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_access_context_manager_service_perimeters#value GoogleAccessContextManagerServicePerimeters#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

