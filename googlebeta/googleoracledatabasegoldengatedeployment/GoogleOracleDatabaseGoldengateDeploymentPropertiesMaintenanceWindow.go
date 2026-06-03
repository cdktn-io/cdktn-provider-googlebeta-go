// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengatedeployment


type GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow struct {
	// Possible values: MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_deployment#day GoogleOracleDatabaseGoldengateDeployment#day}
	Day *string `field:"required" json:"day" yaml:"day"`
	// Start hour for maintenance period. Hour is in UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_deployment#start_hour GoogleOracleDatabaseGoldengateDeployment#start_hour}
	StartHour *float64 `field:"required" json:"startHour" yaml:"startHour"`
}

