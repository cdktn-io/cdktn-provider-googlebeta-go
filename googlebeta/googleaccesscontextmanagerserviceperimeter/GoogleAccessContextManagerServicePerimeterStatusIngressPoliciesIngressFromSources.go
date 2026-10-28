// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagerserviceperimeter


type GoogleAccessContextManagerServicePerimeterStatusIngressPoliciesIngressFromSources struct {
	// An 'AccessLevel' resource name that allow resources within the 'ServicePerimeters' to be accessed from the internet.
	//
	// 'AccessLevels' listed
	// must be in the same policy as this 'ServicePerimeter'. Referencing a nonexistent
	// 'AccessLevel' will cause an error. If no 'AccessLevel' names are listed,
	// resources within the perimeter can only be accessed via Google Cloud calls
	// with request origins within the perimeter.
	// Example 'accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.'
	// If * is specified, then all IngressSources will be allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_access_context_manager_service_perimeter#access_level GoogleAccessContextManagerServicePerimeter#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// psc_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_access_context_manager_service_perimeter#psc_endpoint GoogleAccessContextManagerServicePerimeter#psc_endpoint}
	PscEndpoint *GoogleAccessContextManagerServicePerimeterStatusIngressPoliciesIngressFromSourcesPscEndpoint `field:"optional" json:"pscEndpoint" yaml:"pscEndpoint"`
	// A Google Cloud resource that is allowed to ingress the perimeter.
	//
	// Requests from these resources will be allowed to access perimeter data.
	// Projects, folders, organizations, and VPC networks are allowed.
	// Project format: 'projects/{project_number}'
	// Folder format: 'folders/{folder_number}'
	// Organization format: 'organizations/{organization_number}'
	// VPC network format: '//compute.googleapis.com/projects/{PROJECT_ID}/global/networks/{NAME}'.
	// The resource may be in any Google Cloud organization, not just the
	// organization that the perimeter is defined in. '*' is not allowed, the case
	// of allowing all Google Cloud resources only is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_access_context_manager_service_perimeter#resource GoogleAccessContextManagerServicePerimeter#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

