// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleapigeeenvironmentdebugmask

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleApigeeEnvironmentDebugmaskConfig struct {
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
	// The Apigee environment group associated with the Apigee environment, in the format organizations/{{org_name}}/environments/{{env_name}}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#env_id GoogleApigeeEnvironmentDebugmask#env_id}
	EnvId *string `field:"required" json:"envId" yaml:"envId"`
	// List of XPath expressions that specify the XML elements or attributes that the debug mask applies to for fault messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#fault_x_paths GoogleApigeeEnvironmentDebugmask#fault_x_paths}
	FaultXPaths *[]*string `field:"optional" json:"faultXPaths" yaml:"faultXPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#id GoogleApigeeEnvironmentDebugmask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Map of namespaces to URIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#namespaces GoogleApigeeEnvironmentDebugmask#namespaces}
	Namespaces *map[string]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// List of JSONPath expressions that specify the JSON elements or attributes that the debug mask applies to for request messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#request_json_paths GoogleApigeeEnvironmentDebugmask#request_json_paths}
	RequestJsonPaths *[]*string `field:"optional" json:"requestJsonPaths" yaml:"requestJsonPaths"`
	// List of XPath expressions that specify the XML elements or attributes that the debug mask applies to for request messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#request_x_paths GoogleApigeeEnvironmentDebugmask#request_x_paths}
	RequestXPaths *[]*string `field:"optional" json:"requestXPaths" yaml:"requestXPaths"`
	// List of JSONPath expressions that specify the JSON elements or attributes that the debug mask applies to for response messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#response_json_paths GoogleApigeeEnvironmentDebugmask#response_json_paths}
	ResponseJsonPaths *[]*string `field:"optional" json:"responseJsonPaths" yaml:"responseJsonPaths"`
	// List of XPath expressions that specify the XML elements or attributes that the debug mask applies to for response messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#response_x_paths GoogleApigeeEnvironmentDebugmask#response_x_paths}
	ResponseXPaths *[]*string `field:"optional" json:"responseXPaths" yaml:"responseXPaths"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#timeouts GoogleApigeeEnvironmentDebugmask#timeouts}
	Timeouts *GoogleApigeeEnvironmentDebugmaskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// List of variables that the debug mask applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_apigee_environment_debugmask#variables GoogleApigeeEnvironmentDebugmask#variables}
	Variables *[]*string `field:"optional" json:"variables" yaml:"variables"`
}

