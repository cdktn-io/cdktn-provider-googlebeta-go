// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpec struct {
	// Optional. The A2A Agent Card for the agent (if available).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#agent_card GoogleVertexAiReasoningEngine#agent_card}
	AgentCard *string `field:"optional" json:"agentCard" yaml:"agentCard"`
	// Optional. The OSS agent framework used to develop the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#agent_framework GoogleVertexAiReasoningEngine#agent_framework}
	AgentFramework *string `field:"optional" json:"agentFramework" yaml:"agentFramework"`
	// build_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#build_spec GoogleVertexAiReasoningEngine#build_spec}
	BuildSpec *GoogleVertexAiReasoningEngineSpecBuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Optional. Declarations for object class methods in OpenAPI specification format.
	//
	// **Note**: When deploying via Terraform, this field must be populated manually.
	// Otherwise, client SDKs (like 'agent_engines.get()') will not be able to discover the methods, and calls to the engine (or A2A integrations) will fail.
	//
	// Depending on the template/framework used ('agent_framework'), the required class methods and their parameters differ:
	//
	// **Warning**: The configuration snippets below are illustrative, may not be exhaustive, and could stop working over time. For the most up-to-date method lists and schemas, please consult the respective SDK source code:
	// * For Google ADK: See [ADK Python SDK cli_deploy.py](https://github.com/google/adk-python/blob/68a780306e3bdd648a882ef34c0abf8e5148353e/src/google/adk/cli/cli_deploy.py#L109).
	// * For Langchain: See [Vertex AI Python SDK langchain.py](https://github.com/googleapis/python-aiplatform/blob/c8a38a085931b01f4d6071f0ab7a64cb42851829/agentplatform/agent_engines/templates/langchain.py#L642-L717).
	//
	// ### 1. Langchain Template
	// * 'query' (api_mode = "sync" or empty)
	// * 'stream_query' (api_mode = "stream")
	//
	// Example for Langchain:
	// ```hcl
	// class_methods = jsonencode([
	//   {
	//     name        = "query"
	//     api_mode    = "sync"
	//     description = "Queries the reasoning engine"
	//     parameters  = {
	//       type       = "object"
	//       required   = ["input"]
	//       properties = {
	//         input = {
	//           type        = "string"
	//           description = "The input prompt"
	//         }
	//       }
	//     }
	//   },
	//   {
	//     name        = "stream_query"
	//     api_mode    = "stream"
	//     description = "Streams queries from the reasoning engine"
	//     parameters  = {
	//       type       = "object"
	//       required   = ["input"]
	//       properties = {
	//         input = {
	//           type        = "string"
	//           description = "The input prompt"
	//         }
	//       }
	//     }
	//   }
	// ])
	// ```
	//
	// ### 2. Google ADK Template (Standard - No A2A)
	// For standard Google ADK (Agent Development Kit) deployments, you must define the following 11 methods:
	//
	// Example for Standard ADK:
	// ```hcl
	// class_methods = jsonencode([
	//   {
	//     name        = "get_session"
	//     api_mode    = ""
	//     description = "Retrieve session by ID"
	//     parameters  = {
	//       type     = "object"
	//       required = ["user_id", "session_id"]
	//       properties = {
	//         user_id    = { type = "string" }
	//         session_id = { type = "string" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "async_get_session"
	//     api_mode    = "async"
	//     description = "Retrieve session asynchronously by ID"
	//     parameters  = {
	//       type     = "object"
	//       required = ["user_id", "session_id"]
	//       properties = {
	//         user_id    = { type = "string" }
	//         session_id = { type = "string" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "list_sessions"
	//     api_mode    = ""
	//     description = "List all sessions for a user"
	//     parameters  = {
	//       type     = "object"
	//       required = ["user_id"]
	//       properties = {
	//         user_id = { type = "string" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "async_list_sessions"
	//     api_mode    = "async"
	//     description = "List all sessions for a user asynchronously"
	//     parameters  = {
	//       type     = "object"
	//       required = ["user_id"]
	//       properties = {
	//         user_id = { type = "string" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "create_session"
	//     api_mode    = ""
	//     description = "Create a new session"
	//     parameters  = {
	//       type     = "object"
	//       required = ["user_id"]
	//       properties = {
	//         user_id    = { type = "string" }
	//         session_id = { type = "string" }
	//         state      = { type = "object" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "async_create_session"
	//     api_mode    = "async"
	//     description = "Create a new session asynchronously"
	//     parameters  = {
	//       type     = "object"
	//       required = ["user_id"]
	//       properties = {
	//         user_id    = { type = "string" }
	//         session_id = { type = "string" }
	//         state      = { type = "object" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "delete_session"
	//     api_mode    = ""
	//     description = "Delete session by ID"
	//     parameters  = {
	//       type     = "object"
	//       required = ["user_id", "session_id"]
	//       properties = {
	//         user_id    = { type = "string" }
	//         session_id = { type = "string" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "async_delete_session"
	//     api_mode    = "async"
	//     description = "Delete session asynchronously by ID"
	//     parameters  = {
	//       type     = "object"
	//       required = ["user_id", "session_id"]
	//       properties = {
	//         user_id    = { type = "string" }
	//         session_id = { type = "string" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "stream_query"
	//     api_mode    = "stream"
	//     description = "Stream queries from the agent"
	//     parameters  = {
	//       type     = "object"
	//       required = ["message", "user_id"]
	//       properties = {
	//         message    = { description = "Message string or object" }
	//         user_id    = { type = "string" }
	//         session_id = { type = "string" }
	//         run_config = { type = "object" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "async_stream_query"
	//     api_mode    = "async_stream"
	//     description = "Stream queries asynchronously from the agent"
	//     parameters  = {
	//       type     = "object"
	//       required = ["message", "user_id"]
	//       properties = {
	//         message        = { description = "Message string or object" }
	//         user_id        = { type = "string" }
	//         session_id     = { type = "string" }
	//         session_events = { type = "array", items = { type = "object" } }
	//         run_config     = { type = "object" }
	//       }
	//     }
	//   },
	//   {
	//     name        = "streaming_agent_run_with_events"
	//     api_mode    = "async_stream"
	//     description = "Stream agent run with events asynchronously"
	//     parameters  = {
	//       type     = "object"
	//       required = ["request_json"]
	//       properties = {
	//         request_json = { type = "string" }
	//       }
	//     }
	//   }
	// ])
	// ```
	//
	// ### 3. Google ADK Template (A2A-Enabled)
	// If the agent integrates with the Gemini Enterprise Agent Registry (A2A), you must inject the 'a2a_agent_card' JSON metadata as a string **specifically inside the 'async_create_session' method definition**:
	//
	// Example for A2A-Enabled ADK:
	// ```hcl
	// locals {
	//   # Construct the A2A endpoint URL
	//   a2a_url = "https://us-central1-aiplatform.googleapis.com/v1/projects/my-project/locations/us-central1/reasoningEngines/my-agent/a2a"
	//
	//   agent_card = {
	//     name                 = "my-agent"
	//     description          = "A2A Agent"
	//     version              = "1.0.0"
	//     preferred_transport  = "HTTP_JSON"
	//     supported_interfaces = [{ url = local.a2a_url, protocol_binding = "HTTP_JSON" }]
	//     capabilities         = { streaming = true }
	//   }
	// }
	//
	// # In class_methods, append "a2a_agent_card" key ONLY to the "async_create_session" method:
	// class_methods = jsonencode([
	//   # ... other 10 standard methods (same as Standard ADK) ...
	//   {
	//     name        = "async_create_session"
	//     api_mode    = "async"
	//     description = "Create a new session asynchronously"
	//     parameters  = {
	//       type     = "object"
	//       required = ["user_id"]
	//       properties = {
	//         user_id    = { type = "string" }
	//         session_id = { type = "string" }
	//         state      = { type = "object" }
	//       }
	//     }
	//     # Inject the serialized Agent Card here
	//     a2a_agent_card = jsonencode(local.agent_card)
	//   }
	// ])
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#class_methods GoogleVertexAiReasoningEngine#class_methods}
	ClassMethods *string `field:"optional" json:"classMethods" yaml:"classMethods"`
	// container_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#container_spec GoogleVertexAiReasoningEngine#container_spec}
	ContainerSpec *GoogleVertexAiReasoningEngineSpecContainerSpec `field:"optional" json:"containerSpec" yaml:"containerSpec"`
	// deployment_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#deployment_spec GoogleVertexAiReasoningEngine#deployment_spec}
	DeploymentSpec *GoogleVertexAiReasoningEngineSpecDeploymentSpec `field:"optional" json:"deploymentSpec" yaml:"deploymentSpec"`
	// Optional. The resource name of the linked ExampleStore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#example_store GoogleVertexAiReasoningEngine#example_store}
	ExampleStore *string `field:"optional" json:"exampleStore" yaml:"exampleStore"`
	// Optional.
	//
	// The identity type to use for the Reasoning Engine.
	// If not specified, the 'service_account' field will be used if set,
	// otherwise the default Vertex AI Reasoning Engine Service Agent in the project will be used.
	// Possible values:
	// * 'SERVICE_ACCOUNT': Use a custom service account if the 'service_account' field is set, otherwise use the default Vertex AI Reasoning Engine Service Agent in the project.
	// * 'AGENT_IDENTITY': Use Agent Identity. The 'service_account' field must not be set. Possible values: ["SERVICE_ACCOUNT", "AGENT_IDENTITY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#identity_type GoogleVertexAiReasoningEngine#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// package_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#package_spec GoogleVertexAiReasoningEngine#package_spec}
	PackageSpec *GoogleVertexAiReasoningEngineSpecPackageSpec `field:"optional" json:"packageSpec" yaml:"packageSpec"`
	// Optional.
	//
	// The service account that the Reasoning Engine artifact runs
	// as. It should have "roles/storage.objectViewer" for reading the user
	// project's Cloud Storage and "roles/aiplatform.user" for using Vertex
	// extensions. If not specified, the Vertex AI Reasoning Engine service
	// Agent in the project will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#service_account GoogleVertexAiReasoningEngine#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// source_code_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#source_code_spec GoogleVertexAiReasoningEngine#source_code_spec}
	SourceCodeSpec *GoogleVertexAiReasoningEngineSpecSourceCodeSpec `field:"optional" json:"sourceCodeSpec" yaml:"sourceCodeSpec"`
}

