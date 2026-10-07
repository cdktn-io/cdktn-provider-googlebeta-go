// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatabasemigrationserviceprivateconnection


type GoogleDatabaseMigrationServicePrivateConnectionReservedPublicIpConfig struct {
	// Optional. Number of static public IP addresses to reserve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_database_migration_service_private_connection#nat_ips_count GoogleDatabaseMigrationServicePrivateConnection#nat_ips_count}
	NatIpsCount *float64 `field:"optional" json:"natIpsCount" yaml:"natIpsCount"`
}

