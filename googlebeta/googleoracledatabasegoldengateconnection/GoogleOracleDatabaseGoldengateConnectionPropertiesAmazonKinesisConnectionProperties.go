// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties struct {
	// Access key ID to access the Amazon Kinesis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#access_key_id GoogleOracleDatabaseGoldengateConnection#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// The name of the AWS region. If not provided, Goldengate will default to 'us-west-1'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#aws_region GoogleOracleDatabaseGoldengateConnection#aws_region}
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The endpoint URL of the Amazon Kinesis service. e.g.: 'https://kinesis.us-east-1.amazonaws.com' If not provided, Goldengate will default to 'https://kinesis..amazonaws.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#endpoint GoogleOracleDatabaseGoldengateConnection#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Secret access key to access the Amazon Kinesis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#secret_access_key_secret GoogleOracleDatabaseGoldengateConnection#secret_access_key_secret}
	SecretAccessKeySecret *string `field:"optional" json:"secretAccessKeySecret" yaml:"secretAccessKeySecret"`
	// The technology type of AmazonKinesisConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
}

