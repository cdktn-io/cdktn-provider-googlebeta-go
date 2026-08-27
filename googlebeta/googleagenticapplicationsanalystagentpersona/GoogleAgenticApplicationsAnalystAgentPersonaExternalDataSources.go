// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSources struct {
	// Whether this external data source is enabled for the current analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#enabled GoogleAgenticApplicationsAnalystAgentPersona#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// air_quality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#air_quality GoogleAgenticApplicationsAnalystAgentPersona#air_quality}
	AirQuality *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality `field:"optional" json:"airQuality" yaml:"airQuality"`
	// bureau_labor_statistics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#bureau_labor_statistics GoogleAgenticApplicationsAnalystAgentPersona#bureau_labor_statistics}
	BureauLaborStatistics *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics `field:"optional" json:"bureauLaborStatistics" yaml:"bureauLaborStatistics"`
	// coindesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#coindesk GoogleAgenticApplicationsAnalystAgentPersona#coindesk}
	Coindesk *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk `field:"optional" json:"coindesk" yaml:"coindesk"`
	// finnhub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#finnhub GoogleAgenticApplicationsAnalystAgentPersona#finnhub}
	Finnhub *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub `field:"optional" json:"finnhub" yaml:"finnhub"`
	// fred block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#fred GoogleAgenticApplicationsAnalystAgentPersona#fred}
	Fred *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred `field:"optional" json:"fred" yaml:"fred"`
	// sec_edgar block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#sec_edgar GoogleAgenticApplicationsAnalystAgentPersona#sec_edgar}
	SecEdgar *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar `field:"optional" json:"secEdgar" yaml:"secEdgar"`
	// treasury_securities_auctions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#treasury_securities_auctions GoogleAgenticApplicationsAnalystAgentPersona#treasury_securities_auctions}
	TreasurySecuritiesAuctions *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions `field:"optional" json:"treasurySecuritiesAuctions" yaml:"treasurySecuritiesAuctions"`
	// usda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#usda GoogleAgenticApplicationsAnalystAgentPersona#usda}
	Usda *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda `field:"optional" json:"usda" yaml:"usda"`
}

