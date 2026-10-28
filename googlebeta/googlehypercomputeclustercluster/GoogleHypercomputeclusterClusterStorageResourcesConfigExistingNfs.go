// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterStorageResourcesConfigExistingNfs struct {
	// Remote mount path of the NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_hypercomputecluster_cluster#remote_mount GoogleHypercomputeclusterCluster#remote_mount}
	RemoteMount *string `field:"required" json:"remoteMount" yaml:"remoteMount"`
	// Server IP address of the NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_hypercomputecluster_cluster#server_ip_address GoogleHypercomputeclusterCluster#server_ip_address}
	ServerIpAddress *string `field:"required" json:"serverIpAddress" yaml:"serverIpAddress"`
	// Options to mount the NFS share, as expected by fs_mntops field in fstab(5), for example: "rw,vers=4.1,nconnect=2,noatime".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_hypercomputecluster_cluster#mount_options GoogleHypercomputeclusterCluster#mount_options}
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

