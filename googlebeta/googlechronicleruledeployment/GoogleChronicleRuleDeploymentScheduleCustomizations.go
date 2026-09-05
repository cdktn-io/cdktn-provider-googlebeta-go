// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleruledeployment


type GoogleChronicleRuleDeploymentScheduleCustomizations struct {
	// Indicates whether to add additional delays and runs to rules to ensure enrichment completeness, with the trade-off of more late-arriving detections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_rule_deployment#ensure_enrichment_completeness GoogleChronicleRuleDeployment#ensure_enrichment_completeness}
	EnsureEnrichmentCompleteness interface{} `field:"optional" json:"ensureEnrichmentCompleteness" yaml:"ensureEnrichmentCompleteness"`
	// Delay the first rule execution run to account for late-arriving data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_rule_deployment#late_arriving_data_adjustment GoogleChronicleRuleDeployment#late_arriving_data_adjustment}
	LateArrivingDataAdjustment *string `field:"optional" json:"lateArrivingDataAdjustment" yaml:"lateArrivingDataAdjustment"`
}

