// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclefindingsrefinementdeployment


type GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplication struct {
	// The CuratedRules this detection exclusion applies to. Format: projects/{project}/locations/{location}/instances/{instance}/curatedRules/{rule}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_chronicle_findings_refinement_deployment#curated_rules GoogleChronicleFindingsRefinementDeployment#curated_rules}
	CuratedRules *[]*string `field:"optional" json:"curatedRules" yaml:"curatedRules"`
	// The CuratedRuleSets this detection exclusion applies to. Format: projects/{project}/locations/{location}/instances/{instance}/curatedRuleSetCategories/{category}/curatedRuleSets/{rule_set}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_chronicle_findings_refinement_deployment#curated_rule_sets GoogleChronicleFindingsRefinementDeployment#curated_rule_sets}
	CuratedRuleSets *[]*string `field:"optional" json:"curatedRuleSets" yaml:"curatedRuleSets"`
	// The Rules this detection exclusion applies to. Format: projects/{project}/locations/{location}/instances/{instance}/rules/{rule}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_chronicle_findings_refinement_deployment#rules GoogleChronicleFindingsRefinementDeployment#rules}
	Rules *[]*string `field:"optional" json:"rules" yaml:"rules"`
}

