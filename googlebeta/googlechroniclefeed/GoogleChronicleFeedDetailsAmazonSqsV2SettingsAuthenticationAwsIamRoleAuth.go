// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclefeed


type GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth struct {
	// AWS IAM Role for Identity Federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#aws_iam_role_arn GoogleChronicleFeed#aws_iam_role_arn}
	AwsIamRoleArn *string `field:"optional" json:"awsIamRoleArn" yaml:"awsIamRoleArn"`
	// Subject ID to use for SQS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#subject_id GoogleChronicleFeed#subject_id}
	SubjectId *string `field:"optional" json:"subjectId" yaml:"subjectId"`
}

