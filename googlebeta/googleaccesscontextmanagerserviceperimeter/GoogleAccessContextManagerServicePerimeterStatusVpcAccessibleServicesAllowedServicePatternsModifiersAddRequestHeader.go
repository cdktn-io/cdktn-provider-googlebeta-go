// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagerserviceperimeter


type GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServicesAllowedServicePatternsModifiersAddRequestHeader struct {
	// HTTP header key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_access_context_manager_service_perimeter#key GoogleAccessContextManagerServicePerimeter#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// HTTP header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_access_context_manager_service_perimeter#value GoogleAccessContextManagerServicePerimeter#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

