// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference interface {
	cdktn.ComplexObject
	AmazonKinesisConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference
	AmazonKinesisConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties
	AmazonRedshiftConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionPropertiesOutputReference
	AmazonRedshiftConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties
	AmazonS3ConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionPropertiesOutputReference
	AmazonS3ConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties
	AzureDataLakeStorageConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionPropertiesOutputReference
	AzureDataLakeStorageConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties
	AzureSynapseAnalyticsConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionPropertiesOutputReference
	AzureSynapseAnalyticsConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectionType() *string
	SetConnectionType(val *string)
	ConnectionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabricksConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionPropertiesOutputReference
	DatabricksConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties
	Db2ConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference
	Db2ConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ElasticsearchConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionPropertiesOutputReference
	ElasticsearchConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties
	// Experimental.
	Fqn() *string
	GenericConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGenericConnectionPropertiesOutputReference
	GenericConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties
	GoldengateConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionPropertiesOutputReference
	GoldengateConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties
	GoogleBigQueryConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionPropertiesOutputReference
	GoogleBigQueryConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties
	GoogleCloudStorageConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionPropertiesOutputReference
	GoogleCloudStorageConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties
	GooglePubsubConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionPropertiesOutputReference
	GooglePubsubConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties
	HdfsConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesHdfsConnectionPropertiesOutputReference
	HdfsConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties
	IcebergConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesOutputReference
	IcebergConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties
	IngressIpAddresses() *[]*string
	InternalValue() *GoogleOracleDatabaseGoldengateConnectionProperties
	SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionProperties)
	JavaMessageServiceConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference
	JavaMessageServiceConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties
	KafkaConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference
	KafkaConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties
	KafkaSchemaRegistryConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference
	KafkaSchemaRegistryConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties
	LifecycleDetails() *string
	LifecycleState() *string
	MicrosoftFabricConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionPropertiesOutputReference
	MicrosoftFabricConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties
	MicrosoftSqlserverConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference
	MicrosoftSqlserverConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties
	MongodbConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesMongodbConnectionPropertiesOutputReference
	MongodbConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties
	MysqlConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference
	MysqlConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties
	Ocid() *string
	OciObjectStorageConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionPropertiesOutputReference
	OciObjectStorageConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties
	OracleAiDataPlatformConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference
	OracleAiDataPlatformConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties
	OracleConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesOracleConnectionPropertiesOutputReference
	OracleConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties
	OracleNosqlConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference
	OracleNosqlConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties
	PostgresqlConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionPropertiesOutputReference
	PostgresqlConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties
	RedisConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesRedisConnectionPropertiesOutputReference
	RedisConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties
	RoutingMethod() *string
	SetRoutingMethod(val *string)
	RoutingMethodInput() *string
	SnowflakeConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionPropertiesOutputReference
	SnowflakeConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateTime() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAmazonKinesisConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties)
	PutAmazonRedshiftConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties)
	PutAmazonS3ConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties)
	PutAzureDataLakeStorageConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties)
	PutAzureSynapseAnalyticsConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties)
	PutDatabricksConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties)
	PutDb2ConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties)
	PutElasticsearchConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties)
	PutGenericConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties)
	PutGoldengateConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties)
	PutGoogleBigQueryConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties)
	PutGoogleCloudStorageConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties)
	PutGooglePubsubConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties)
	PutHdfsConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties)
	PutIcebergConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties)
	PutJavaMessageServiceConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties)
	PutKafkaConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties)
	PutKafkaSchemaRegistryConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties)
	PutMicrosoftFabricConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties)
	PutMicrosoftSqlserverConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties)
	PutMongodbConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties)
	PutMysqlConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties)
	PutOciObjectStorageConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties)
	PutOracleAiDataPlatformConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties)
	PutOracleConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties)
	PutOracleNosqlConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties)
	PutPostgresqlConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties)
	PutRedisConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties)
	PutSnowflakeConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties)
	ResetAmazonKinesisConnectionProperties()
	ResetAmazonRedshiftConnectionProperties()
	ResetAmazonS3ConnectionProperties()
	ResetAzureDataLakeStorageConnectionProperties()
	ResetAzureSynapseAnalyticsConnectionProperties()
	ResetDatabricksConnectionProperties()
	ResetDb2ConnectionProperties()
	ResetDescription()
	ResetElasticsearchConnectionProperties()
	ResetGenericConnectionProperties()
	ResetGoldengateConnectionProperties()
	ResetGoogleBigQueryConnectionProperties()
	ResetGoogleCloudStorageConnectionProperties()
	ResetGooglePubsubConnectionProperties()
	ResetHdfsConnectionProperties()
	ResetIcebergConnectionProperties()
	ResetJavaMessageServiceConnectionProperties()
	ResetKafkaConnectionProperties()
	ResetKafkaSchemaRegistryConnectionProperties()
	ResetMicrosoftFabricConnectionProperties()
	ResetMicrosoftSqlserverConnectionProperties()
	ResetMongodbConnectionProperties()
	ResetMysqlConnectionProperties()
	ResetOciObjectStorageConnectionProperties()
	ResetOracleAiDataPlatformConnectionProperties()
	ResetOracleConnectionProperties()
	ResetOracleNosqlConnectionProperties()
	ResetPostgresqlConnectionProperties()
	ResetRedisConnectionProperties()
	ResetRoutingMethod()
	ResetSnowflakeConnectionProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonKinesisConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"amazonKinesisConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonKinesisConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties
	_jsii_.Get(
		j,
		"amazonKinesisConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonRedshiftConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"amazonRedshiftConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonRedshiftConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties
	_jsii_.Get(
		j,
		"amazonRedshiftConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonS3ConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"amazonS3ConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AmazonS3ConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties
	_jsii_.Get(
		j,
		"amazonS3ConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AzureDataLakeStorageConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"azureDataLakeStorageConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AzureDataLakeStorageConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties
	_jsii_.Get(
		j,
		"azureDataLakeStorageConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AzureSynapseAnalyticsConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"azureSynapseAnalyticsConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) AzureSynapseAnalyticsConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties
	_jsii_.Get(
		j,
		"azureSynapseAnalyticsConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) DatabricksConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"databricksConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) DatabricksConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties
	_jsii_.Get(
		j,
		"databricksConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) Db2ConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"db2ConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) Db2ConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties
	_jsii_.Get(
		j,
		"db2ConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ElasticsearchConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ElasticsearchConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties
	_jsii_.Get(
		j,
		"elasticsearchConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GenericConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGenericConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesGenericConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"genericConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GenericConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties
	_jsii_.Get(
		j,
		"genericConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GoldengateConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"goldengateConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GoldengateConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties
	_jsii_.Get(
		j,
		"goldengateConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GoogleBigQueryConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"googleBigQueryConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GoogleBigQueryConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties
	_jsii_.Get(
		j,
		"googleBigQueryConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GoogleCloudStorageConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorageConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GoogleCloudStorageConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties
	_jsii_.Get(
		j,
		"googleCloudStorageConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GooglePubsubConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"googlePubsubConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GooglePubsubConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties
	_jsii_.Get(
		j,
		"googlePubsubConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) HdfsConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesHdfsConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesHdfsConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"hdfsConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) HdfsConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties
	_jsii_.Get(
		j,
		"hdfsConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) IcebergConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"icebergConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) IcebergConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties
	_jsii_.Get(
		j,
		"icebergConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) IngressIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) JavaMessageServiceConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"javaMessageServiceConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) JavaMessageServiceConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties
	_jsii_.Get(
		j,
		"javaMessageServiceConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) KafkaConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"kafkaConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) KafkaConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties
	_jsii_.Get(
		j,
		"kafkaConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) KafkaSchemaRegistryConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"kafkaSchemaRegistryConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) KafkaSchemaRegistryConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties
	_jsii_.Get(
		j,
		"kafkaSchemaRegistryConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) MicrosoftFabricConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"microsoftFabricConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) MicrosoftFabricConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties
	_jsii_.Get(
		j,
		"microsoftFabricConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) MicrosoftSqlserverConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"microsoftSqlserverConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) MicrosoftSqlserverConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties
	_jsii_.Get(
		j,
		"microsoftSqlserverConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) MongodbConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesMongodbConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesMongodbConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"mongodbConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) MongodbConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties
	_jsii_.Get(
		j,
		"mongodbConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) MysqlConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"mysqlConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) MysqlConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties
	_jsii_.Get(
		j,
		"mysqlConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) OciObjectStorageConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"ociObjectStorageConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) OciObjectStorageConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties
	_jsii_.Get(
		j,
		"ociObjectStorageConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleAiDataPlatformConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"oracleAiDataPlatformConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleAiDataPlatformConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties
	_jsii_.Get(
		j,
		"oracleAiDataPlatformConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesOracleConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesOracleConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"oracleConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties
	_jsii_.Get(
		j,
		"oracleConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleNosqlConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"oracleNosqlConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) OracleNosqlConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties
	_jsii_.Get(
		j,
		"oracleNosqlConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PostgresqlConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"postgresqlConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PostgresqlConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties
	_jsii_.Get(
		j,
		"postgresqlConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) RedisConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesRedisConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesRedisConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"redisConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) RedisConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties
	_jsii_.Get(
		j,
		"redisConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) RoutingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) RoutingMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) SnowflakeConnectionProperties() GoogleOracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionPropertiesOutputReference {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"snowflakeConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) SnowflakeConnectionPropertiesInput() *GoogleOracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties
	_jsii_.Get(
		j,
		"snowflakeConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference_Override(g GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference)SetConnectionType(val *string) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference)SetRoutingMethod(val *string) {
	if err := j.validateSetRoutingMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAmazonKinesisConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties) {
	if err := g.validatePutAmazonKinesisConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAmazonKinesisConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAmazonRedshiftConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonRedshiftConnectionProperties) {
	if err := g.validatePutAmazonRedshiftConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAmazonRedshiftConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAmazonS3ConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonS3ConnectionProperties) {
	if err := g.validatePutAmazonS3ConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAmazonS3ConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAzureDataLakeStorageConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureDataLakeStorageConnectionProperties) {
	if err := g.validatePutAzureDataLakeStorageConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureDataLakeStorageConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutAzureSynapseAnalyticsConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesAzureSynapseAnalyticsConnectionProperties) {
	if err := g.validatePutAzureSynapseAnalyticsConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureSynapseAnalyticsConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutDatabricksConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesDatabricksConnectionProperties) {
	if err := g.validatePutDatabricksConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatabricksConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutDb2ConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties) {
	if err := g.validatePutDb2ConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDb2ConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutElasticsearchConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties) {
	if err := g.validatePutElasticsearchConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putElasticsearchConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGenericConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGenericConnectionProperties) {
	if err := g.validatePutGenericConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGenericConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGoldengateConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGoldengateConnectionProperties) {
	if err := g.validatePutGoldengateConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoldengateConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGoogleBigQueryConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleBigQueryConnectionProperties) {
	if err := g.validatePutGoogleBigQueryConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleBigQueryConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGoogleCloudStorageConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGoogleCloudStorageConnectionProperties) {
	if err := g.validatePutGoogleCloudStorageConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleCloudStorageConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutGooglePubsubConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesGooglePubsubConnectionProperties) {
	if err := g.validatePutGooglePubsubConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGooglePubsubConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutHdfsConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties) {
	if err := g.validatePutHdfsConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHdfsConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutIcebergConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties) {
	if err := g.validatePutIcebergConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIcebergConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutJavaMessageServiceConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties) {
	if err := g.validatePutJavaMessageServiceConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJavaMessageServiceConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutKafkaConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties) {
	if err := g.validatePutKafkaConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKafkaConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutKafkaSchemaRegistryConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties) {
	if err := g.validatePutKafkaSchemaRegistryConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKafkaSchemaRegistryConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutMicrosoftFabricConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftFabricConnectionProperties) {
	if err := g.validatePutMicrosoftFabricConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMicrosoftFabricConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutMicrosoftSqlserverConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties) {
	if err := g.validatePutMicrosoftSqlserverConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMicrosoftSqlserverConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutMongodbConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties) {
	if err := g.validatePutMongodbConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMongodbConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutMysqlConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties) {
	if err := g.validatePutMysqlConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMysqlConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutOciObjectStorageConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesOciObjectStorageConnectionProperties) {
	if err := g.validatePutOciObjectStorageConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOciObjectStorageConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutOracleAiDataPlatformConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties) {
	if err := g.validatePutOracleAiDataPlatformConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOracleAiDataPlatformConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutOracleConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties) {
	if err := g.validatePutOracleConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOracleConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutOracleNosqlConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties) {
	if err := g.validatePutOracleNosqlConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOracleNosqlConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutPostgresqlConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties) {
	if err := g.validatePutPostgresqlConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostgresqlConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutRedisConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties) {
	if err := g.validatePutRedisConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRedisConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) PutSnowflakeConnectionProperties(value *GoogleOracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties) {
	if err := g.validatePutSnowflakeConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSnowflakeConnectionProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAmazonKinesisConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAmazonKinesisConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAmazonRedshiftConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAmazonRedshiftConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAmazonS3ConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAmazonS3ConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAzureDataLakeStorageConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureDataLakeStorageConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetAzureSynapseAnalyticsConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureSynapseAnalyticsConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetDatabricksConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabricksConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetDb2ConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetDb2ConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetElasticsearchConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetElasticsearchConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGenericConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetGenericConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGoldengateConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetGoldengateConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGoogleBigQueryConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleBigQueryConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGoogleCloudStorageConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleCloudStorageConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetGooglePubsubConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetGooglePubsubConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetHdfsConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetHdfsConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetIcebergConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetIcebergConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetJavaMessageServiceConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetJavaMessageServiceConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetKafkaConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetKafkaConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetKafkaSchemaRegistryConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetKafkaSchemaRegistryConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetMicrosoftFabricConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetMicrosoftFabricConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetMicrosoftSqlserverConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetMicrosoftSqlserverConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetMongodbConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetMongodbConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetMysqlConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetMysqlConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetOciObjectStorageConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetOciObjectStorageConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetOracleAiDataPlatformConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetOracleAiDataPlatformConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetOracleConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetOracleConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetOracleNosqlConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetOracleNosqlConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetPostgresqlConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetPostgresqlConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetRedisConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetRedisConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetRoutingMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetRoutingMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ResetSnowflakeConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetSnowflakeConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

