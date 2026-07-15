// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpec struct {
	// agent_gateway_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#agent_gateway_config GoogleVertexAiReasoningEngine#agent_gateway_config}
	AgentGatewayConfig *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfig `field:"optional" json:"agentGatewayConfig" yaml:"agentGatewayConfig"`
	// Optional.
	//
	// The agent server mode specifies what features are used when deploy the agent to agent engine.
	// Possible values:
	// * 'STABLE': Stable agent server mode.
	// * 'EXPERIMENTAL': Experimental agent server mode. Possible values: ["STABLE", "EXPERIMENTAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#agent_server_mode GoogleVertexAiReasoningEngine#agent_server_mode}
	AgentServerMode *string `field:"optional" json:"agentServerMode" yaml:"agentServerMode"`
	// Optional. Concurrency for each container and agent server. Recommended value: 2 * cpu + 1. Defaults to 9.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#container_concurrency GoogleVertexAiReasoningEngine#container_concurrency}
	ContainerConcurrency *float64 `field:"optional" json:"containerConcurrency" yaml:"containerConcurrency"`
	// Optional.
	//
	// Whether to enable dedicated ingress endpoint for the deployment. If true, the deployment will be accessible via a dedicated endpoint. This is required to enable GKE V2 runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#dedicated_ingress_endpoint_enabled GoogleVertexAiReasoningEngine#dedicated_ingress_endpoint_enabled}
	DedicatedIngressEndpointEnabled interface{} `field:"optional" json:"dedicatedIngressEndpointEnabled" yaml:"dedicatedIngressEndpointEnabled"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#env GoogleVertexAiReasoningEngine#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// keep_alive_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#keep_alive_probe GoogleVertexAiReasoningEngine#keep_alive_probe}
	KeepAliveProbe *GoogleVertexAiReasoningEngineSpecDeploymentSpecKeepAliveProbe `field:"optional" json:"keepAliveProbe" yaml:"keepAliveProbe"`
	// Optional.
	//
	// The maximum number of application instances that can be
	// launched to handle increased traffic. Defaults to 100.
	// Range: [1, 1000]. If VPC-SC or PSC-I is enabled, the acceptable
	// range is [1, 100].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#max_instances GoogleVertexAiReasoningEngine#max_instances}
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// Optional.
	//
	// The minimum number of application instances that will be
	// kept running at all times. Defaults to 1. Range: [0, 10].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#min_instances GoogleVertexAiReasoningEngine#min_instances}
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// psc_interface_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#psc_interface_config GoogleVertexAiReasoningEngine#psc_interface_config}
	PscInterfaceConfig *GoogleVertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfig `field:"optional" json:"pscInterfaceConfig" yaml:"pscInterfaceConfig"`
	// Optional. Resource limits for each container. Only 'cpu' and 'memory' keys are supported.
	//
	// Defaults to {"cpu": "4", "memory": "4Gi"}.
	//
	// The only supported values for CPU are '1', '2', '4', '6' and '8'.
	// For more information, go to
	// https://cloud.google.com/run/docs/configuring/cpu.
	//
	// The only supported values for memory are '1Gi', '2Gi', ... '32 Gi'.
	// For more information, go to
	// https://cloud.google.com/run/docs/configuring/memory-limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#resource_limits GoogleVertexAiReasoningEngine#resource_limits}
	ResourceLimits *map[string]*string `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
	// secret_env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#secret_env GoogleVertexAiReasoningEngine#secret_env}
	SecretEnv interface{} `field:"optional" json:"secretEnv" yaml:"secretEnv"`
}

