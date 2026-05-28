// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetappbackup


type GoogleNetappBackupOntapSource struct {
	// Name of the storage pool. This must be specified for creating backups for ONTAP mode volumes. Format: 'projects/{{project}}/locations/{{location}}/storagePools/{{storage_pool_id}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_netapp_backup#storage_pool GoogleNetappBackup#storage_pool}
	StoragePool *string `field:"required" json:"storagePool" yaml:"storagePool"`
	// The UUID of the ONTAP source volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_netapp_backup#volume_uuid GoogleNetappBackup#volume_uuid}
	VolumeUuid *string `field:"required" json:"volumeUuid" yaml:"volumeUuid"`
	// The UUID of the ONTAP source snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_netapp_backup#snapshot_uuid GoogleNetappBackup#snapshot_uuid}
	SnapshotUuid *string `field:"optional" json:"snapshotUuid" yaml:"snapshotUuid"`
}

