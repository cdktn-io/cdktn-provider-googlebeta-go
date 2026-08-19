// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleAgenticApplicationsAnalystAgentPersonaConfig struct {
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
	// Id of the requesting object If auto-generating Id server-side, remove this field and analyst_agent_persona_id from the method_signature of Create RPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#analyst_agent_persona_id GoogleAgenticApplicationsAnalystAgentPersona#analyst_agent_persona_id}
	AnalystAgentPersonaId *string `field:"required" json:"analystAgentPersonaId" yaml:"analystAgentPersonaId"`
	// The display name of the persona, shown to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#display_name GoogleAgenticApplicationsAnalystAgentPersona#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#location GoogleAgenticApplicationsAnalystAgentPersona#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// artifact_examples block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#artifact_examples GoogleAgenticApplicationsAnalystAgentPersona#artifact_examples}
	ArtifactExamples interface{} `field:"optional" json:"artifactExamples" yaml:"artifactExamples"`
	// artifacts_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#artifacts_config GoogleAgenticApplicationsAnalystAgentPersona#artifacts_config}
	ArtifactsConfig *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfig `field:"optional" json:"artifactsConfig" yaml:"artifactsConfig"`
	// The customer-specific context to be used by the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#customer_context GoogleAgenticApplicationsAnalystAgentPersona#customer_context}
	CustomerContext *[]*string `field:"optional" json:"customerContext" yaml:"customerContext"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#deletion_policy GoogleAgenticApplicationsAnalystAgentPersona#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The description of the persona, shown to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#display_description GoogleAgenticApplicationsAnalystAgentPersona#display_description}
	DisplayDescription *string `field:"optional" json:"displayDescription" yaml:"displayDescription"`
	// external_data_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#external_data_sources GoogleAgenticApplicationsAnalystAgentPersona#external_data_sources}
	ExternalDataSources interface{} `field:"optional" json:"externalDataSources" yaml:"externalDataSources"`
	// The Gemini Enterprise Engine ID associated with this persona.
	//
	// If set, any requests coming from this GE Engine will be routed to this
	// persona.
	// If not set, requests from GE will only be routed to this persona if its
	// name ends in "/default".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#gemini_enterprise_engine GoogleAgenticApplicationsAnalystAgentPersona#gemini_enterprise_engine}
	GeminiEnterpriseEngine *string `field:"optional" json:"geminiEnterpriseEngine" yaml:"geminiEnterpriseEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#id GoogleAgenticApplicationsAnalystAgentPersona#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mcp_data_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#mcp_data_sources GoogleAgenticApplicationsAnalystAgentPersona#mcp_data_sources}
	McpDataSources interface{} `field:"optional" json:"mcpDataSources" yaml:"mcpDataSources"`
	// The description of the persona review, used by the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#model_description GoogleAgenticApplicationsAnalystAgentPersona#model_description}
	ModelDescription *string `field:"optional" json:"modelDescription" yaml:"modelDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#project GoogleAgenticApplicationsAnalystAgentPersona#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#resources GoogleAgenticApplicationsAnalystAgentPersona#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// Possible values: ANALYST_ROLE_GENERIC_FINANCE_ANALYST ANALYST_ROLE_CORPORATE_FINANCE_ANALYST ANALYST_ROLE_CROSS_ASSET_DERIVATIVES_STRATEGIST ANALYST_ROLE_KYC_ANALYST ANALYST_ROLE_SALES_TRADER ANALYST_ROLE_QUANT_ANALYST ANALYST_ROLE_EXCHANGE_MANAGER ANALYST_ROLE_PORTFOLIO_MANAGER ANALYST_ROLE_WEALTH_MANAGER ANALYST_ROLE_INSTITUTIONAL_PORTFOLIO_STRATEGIST ANALYST_ROLE_MNA_EXECUTION_ANALYST ANALYST_ROLE_ECM_ORIGINATION_STRATEGIST ANALYST_ROLE_LEVERAGED_FINANCE_SPECIALIST ANALYST_ROLE_INVESTMENT_RESEARCH_ANALYST ANALYST_ROLE_CORPORATE_BANKING_ANALYST ANALYST_ROLE_CREDIT_RISK_STRATEGIST ANALYST_ROLE_BEHAVIORAL_FINANCIAL_STRATEGIST ANALYST_ROLE_FUND_ACCOUNTANT ANALYST_ROLE_MODEL_VALIDATION_AUDITOR ANALYST_ROLE_PRIVATE_EQUITY_SPECIALIST ANALYST_ROLE_TREASURY_ANALYST ANALYST_ROLE_VENTURE_CAPITAL_ANALYST ANALYST_ROLE_AML_INVESTIGATOR ANALYST_ROLE_DUE_DILIGENCE_ANALYST ANALYST_ROLE_INSURANCE_CLAIMS_ANALYST ANALYST_ROLE_SPECIALTY_LIABILITY_UNDERWRITER ANALYST_ROLE_CATASTROPHE_EXPOSURE_MODELER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#role GoogleAgenticApplicationsAnalystAgentPersona#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// skills block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#skills GoogleAgenticApplicationsAnalystAgentPersona#skills}
	Skills interface{} `field:"optional" json:"skills" yaml:"skills"`
	// tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#tables GoogleAgenticApplicationsAnalystAgentPersona#tables}
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#timeouts GoogleAgenticApplicationsAnalystAgentPersona#timeouts}
	Timeouts *GoogleAgenticApplicationsAnalystAgentPersonaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

