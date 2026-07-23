// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleartifactregistryprojectconfig


type GoogleArtifactRegistryProjectConfigPlatformLogsConfig struct {
	// The state of the platform logs: enabled or disabled. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_artifact_registry_project_config#logging_state GoogleArtifactRegistryProjectConfig#logging_state}
	LoggingState *string `field:"optional" json:"loggingState" yaml:"loggingState"`
	// The severity level for the logs.
	//
	// Logs will be generated if their
	// severity level is >= than the value of the severity level mentioned here. Possible values: ["DEBUG", "INFO", "NOTICE", "WARNING", "ERROR", "CRITICAL", "ALERT", "EMERGENCY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_artifact_registry_project_config#severity_level GoogleArtifactRegistryProjectConfig#severity_level}
	SeverityLevel *string `field:"optional" json:"severityLevel" yaml:"severityLevel"`
}

