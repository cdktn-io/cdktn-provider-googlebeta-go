// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatabasemigrationserviceprivateconnection


type GoogleDatabaseMigrationServicePrivateConnectionPscInterfaceConfig struct {
	// Fully qualified name of the Network Attachment that DMS will connect to. Format: projects/{project}/regions/{region}/networkAttachments/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_database_migration_service_private_connection#network_attachment GoogleDatabaseMigrationServicePrivateConnection#network_attachment}
	NetworkAttachment *string `field:"required" json:"networkAttachment" yaml:"networkAttachment"`
}

