// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnectionassignment


type GoogleOracleDatabaseGoldengateConnectionAssignmentProperties struct {
	// The GoldengateConnection resource to be assigned. Format: projects/{project}/locations/{location}/goldengateConnections/{goldengate_connection}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection_assignment#goldengate_connection GoogleOracleDatabaseGoldengateConnectionAssignment#goldengate_connection}
	GoldengateConnection *string `field:"required" json:"goldengateConnection" yaml:"goldengateConnection"`
	// The GoldenGateDeployment to assign the connection to. Format: projects/{project}/locations/{location}/goldengateDeployments/{goldengate_deployment}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection_assignment#goldengate_deployment GoogleOracleDatabaseGoldengateConnectionAssignment#goldengate_deployment}
	GoldengateDeployment *string `field:"required" json:"goldengateDeployment" yaml:"goldengateDeployment"`
}

