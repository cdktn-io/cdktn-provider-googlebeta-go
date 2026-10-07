// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent


type GoogleCesAgentBeforeToolCallbacks struct {
	// The python code to execute for the callback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#python_code GoogleCesAgent#python_code}
	PythonCode *string `field:"required" json:"pythonCode" yaml:"pythonCode"`
	// Human-readable description of the callback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#description GoogleCesAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the callback is disabled. Disabled callbacks are ignored by the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#disabled GoogleCesAgent#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// If enabled, the callback will also be executed on intermediate model outputs.
	//
	// This setting only affects after model callback.
	// **ENABLE WITH CAUTION**. Typically after model callback only needs to be
	// executed after receiving all model responses. Enabling proactive execution
	// may have negative implication on the execution cost and latency, and
	// should only be enabled in rare situations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#proactive_execution_enabled GoogleCesAgent#proactive_execution_enabled}
	ProactiveExecutionEnabled interface{} `field:"optional" json:"proactiveExecutionEnabled" yaml:"proactiveExecutionEnabled"`
}

