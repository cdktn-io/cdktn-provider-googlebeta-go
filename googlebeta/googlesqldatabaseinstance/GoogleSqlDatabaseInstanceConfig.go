// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesqldatabaseinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleSqlDatabaseInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The MySQL, PostgreSQL or SQL Server version to use.
	//
	// Supported values include MYSQL_5_6, MYSQL_5_7, MYSQL_8_0, MYSQL_8_4, POSTGRES_9_6, POSTGRES_10, POSTGRES_11, POSTGRES_12, POSTGRES_13, POSTGRES_14, POSTGRES_15, POSTGRES_16, POSTGRES_17, POSTGRES_18, SQLSERVER_2022_STANDARD, SQLSERVER_2022_ENTERPRISE, SQLSERVER_2022_EXPRESS, SQLSERVER_2022_WEB, SQLSERVER_2025_STANDARD, SQLSERVER_2025_ENTERPRISE, SQLSERVER_2025_EXPRESS, SQLSERVER_2025_WEB. Database Version Policies includes an up-to-date reference of supported versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#database_version GoogleSqlDatabaseInstance#database_version}
	DatabaseVersion *string `field:"required" json:"databaseVersion" yaml:"databaseVersion"`
	// The name of the BackupDR backup to restore from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#backupdr_backup GoogleSqlDatabaseInstance#backupdr_backup}
	BackupdrBackup *string `field:"optional" json:"backupdrBackup" yaml:"backupdrBackup"`
	// clone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#clone GoogleSqlDatabaseInstance#clone}
	Clone *GoogleSqlDatabaseInstanceClone `field:"optional" json:"clone" yaml:"clone"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#deletion_policy GoogleSqlDatabaseInstance#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Used to block Terraform from deleting a SQL Instance. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#deletion_protection GoogleSqlDatabaseInstance#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#encryption_key_name GoogleSqlDatabaseInstance#encryption_key_name}.
	EncryptionKeyName *string `field:"optional" json:"encryptionKeyName" yaml:"encryptionKeyName"`
	// Whether to enforce the new SQL network architecture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#enforce_new_sql_network_architecture GoogleSqlDatabaseInstance#enforce_new_sql_network_architecture}
	EnforceNewSqlNetworkArchitecture interface{} `field:"optional" json:"enforceNewSqlNetworkArchitecture" yaml:"enforceNewSqlNetworkArchitecture"`
	// The description of final backup if instance enable create final backup during instance deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#final_backup_description GoogleSqlDatabaseInstance#final_backup_description}
	FinalBackupDescription *string `field:"optional" json:"finalBackupDescription" yaml:"finalBackupDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#id GoogleSqlDatabaseInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When this parameter is set to true, Cloud SQL instances can perform in-place major version upgrades of read replicas along with the primary instance when 'database_version' is updated.
	//
	// This is an input-only field that is not persisted in the API and only takes effect during a major version upgrade.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#include_replicas_for_major_version_upgrade GoogleSqlDatabaseInstance#include_replicas_for_major_version_upgrade}
	IncludeReplicasForMajorVersionUpgrade interface{} `field:"optional" json:"includeReplicasForMajorVersionUpgrade" yaml:"includeReplicasForMajorVersionUpgrade"`
	// The type of the instance. See https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1/instances#SqlInstanceType for supported values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#instance_type GoogleSqlDatabaseInstance#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Maintenance version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#maintenance_version GoogleSqlDatabaseInstance#maintenance_version}
	MaintenanceVersion *string `field:"optional" json:"maintenanceVersion" yaml:"maintenanceVersion"`
	// The name of the instance that will act as the master in the replication setup.
	//
	// Note, this requires the master to have binary_log_enabled set, as well as existing backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#master_instance_name GoogleSqlDatabaseInstance#master_instance_name}
	MasterInstanceName *string `field:"optional" json:"masterInstanceName" yaml:"masterInstanceName"`
	// The name of the instance.
	//
	// If the name is left blank, Terraform will randomly generate one when the instance is first created. This is done because after a name is used, it cannot be reused for up to one week.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#name GoogleSqlDatabaseInstance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// For a read pool instance, the number of nodes in the read pool.
	//
	// For read pools with auto scaling enabled, this field is read only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#node_count GoogleSqlDatabaseInstance#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// point_in_time_restore_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#point_in_time_restore_context GoogleSqlDatabaseInstance#point_in_time_restore_context}
	PointInTimeRestoreContext *GoogleSqlDatabaseInstancePointInTimeRestoreContext `field:"optional" json:"pointInTimeRestoreContext" yaml:"pointInTimeRestoreContext"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#project GoogleSqlDatabaseInstance#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region the instance will sit in.
	//
	// Note, Cloud SQL is not available in all regions. A valid region must be provided to use this resource. If a region is not provided in the resource definition, the provider region will be used instead, but this will be an apply-time error for instances if the provider region is not supported with Cloud SQL. If you choose not to provide the region argument for this resource, make sure you understand this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#region GoogleSqlDatabaseInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// replica_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#replica_configuration GoogleSqlDatabaseInstance#replica_configuration}
	ReplicaConfiguration *GoogleSqlDatabaseInstanceReplicaConfiguration `field:"optional" json:"replicaConfiguration" yaml:"replicaConfiguration"`
	// The replicas of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#replica_names GoogleSqlDatabaseInstance#replica_names}
	ReplicaNames *[]*string `field:"optional" json:"replicaNames" yaml:"replicaNames"`
	// replication_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#replication_cluster GoogleSqlDatabaseInstance#replication_cluster}
	ReplicationCluster *GoogleSqlDatabaseInstanceReplicationCluster `field:"optional" json:"replicationCluster" yaml:"replicationCluster"`
	// restore_backup_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#restore_backup_context GoogleSqlDatabaseInstance#restore_backup_context}
	RestoreBackupContext *GoogleSqlDatabaseInstanceRestoreBackupContext `field:"optional" json:"restoreBackupContext" yaml:"restoreBackupContext"`
	// Initial root password. Required for MS SQL Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#root_password GoogleSqlDatabaseInstance#root_password}
	RootPassword *string `field:"optional" json:"rootPassword" yaml:"rootPassword"`
	// Initial root password.
	//
	// Required for MS SQL Server.
	// 				Note: This property is write-only and will not be read from the API. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#root_password_wo GoogleSqlDatabaseInstance#root_password_wo}
	RootPasswordWo *string `field:"optional" json:"rootPasswordWo" yaml:"rootPasswordWo"`
	// Triggers update of root_password_wo write-only. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#root_password_wo_version GoogleSqlDatabaseInstance#root_password_wo_version}
	RootPasswordWoVersion *string `field:"optional" json:"rootPasswordWoVersion" yaml:"rootPasswordWoVersion"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#settings GoogleSqlDatabaseInstance#settings}
	Settings *GoogleSqlDatabaseInstanceSettings `field:"optional" json:"settings" yaml:"settings"`
	// When set to true, Cloud SQL instances can switch storing point-in-time recovery transaction logs from a data disk to Cloud Storage, freeing up data disk space and enabling longer retention windows.
	//
	// This is an input-only field that is not persisted in the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#switch_transaction_logs_to_cloud_storage_enabled GoogleSqlDatabaseInstance#switch_transaction_logs_to_cloud_storage_enabled}
	SwitchTransactionLogsToCloudStorageEnabled interface{} `field:"optional" json:"switchTransactionLogsToCloudStorageEnabled" yaml:"switchTransactionLogsToCloudStorageEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_sql_database_instance#timeouts GoogleSqlDatabaseInstance#timeouts}
	Timeouts *GoogleSqlDatabaseInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

