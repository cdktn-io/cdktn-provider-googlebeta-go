// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengatedeployment


type GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig struct {
	// Defines auto upgrade period for bundle releases.
	//
	// Manually configured period
	// cannot be longer than service defined period for bundle releases. This
	// period must be shorter or equal to major release upgrade period. Not
	// passing this field during create will equate to using the service default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_oracle_database_goldengate_deployment#bundle_release_upgrade_period_days GoogleOracleDatabaseGoldengateDeployment#bundle_release_upgrade_period_days}
	BundleReleaseUpgradePeriodDays *float64 `field:"optional" json:"bundleReleaseUpgradePeriodDays" yaml:"bundleReleaseUpgradePeriodDays"`
	// Defines auto upgrade period for interim releases. This period must be shorter or equal to bundle release upgrade period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_oracle_database_goldengate_deployment#interim_release_upgrade_period_days GoogleOracleDatabaseGoldengateDeployment#interim_release_upgrade_period_days}
	InterimReleaseUpgradePeriodDays *float64 `field:"optional" json:"interimReleaseUpgradePeriodDays" yaml:"interimReleaseUpgradePeriodDays"`
	// By default auto upgrade for interim releases are not enabled.
	//
	// If
	// auto-upgrade is enabled for interim release,  you have to specify
	// interim_release_upgrade_period_days too.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_oracle_database_goldengate_deployment#is_interim_release_auto_upgrade_enabled GoogleOracleDatabaseGoldengateDeployment#is_interim_release_auto_upgrade_enabled}
	IsInterimReleaseAutoUpgradeEnabled interface{} `field:"optional" json:"isInterimReleaseAutoUpgradeEnabled" yaml:"isInterimReleaseAutoUpgradeEnabled"`
	// Defines auto upgrade period for major releases.
	//
	// Manually configured period
	// cannot be longer than service defined period for major releases. Not
	// passing this field during create will equate to using the service default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_oracle_database_goldengate_deployment#major_release_upgrade_period_days GoogleOracleDatabaseGoldengateDeployment#major_release_upgrade_period_days}
	MajorReleaseUpgradePeriodDays *float64 `field:"optional" json:"majorReleaseUpgradePeriodDays" yaml:"majorReleaseUpgradePeriodDays"`
	// Defines auto upgrade period for releases with security fix.
	//
	// Manually
	// configured period cannot be longer than service defined period for security
	// releases. Not passing this field during create will equate to using the
	// service default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_oracle_database_goldengate_deployment#security_patch_upgrade_period_days GoogleOracleDatabaseGoldengateDeployment#security_patch_upgrade_period_days}
	SecurityPatchUpgradePeriodDays *float64 `field:"optional" json:"securityPatchUpgradePeriodDays" yaml:"securityPatchUpgradePeriodDays"`
}

