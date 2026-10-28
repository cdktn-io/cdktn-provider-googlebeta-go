// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlestoragecontrolprojectfeatureconfig


type GoogleStorageControlProjectFeatureConfigFilterIncludedCloudStorageBuckets struct {
	// The names of the buckets to include.
	//
	// For Auto Annotate, regex wildcards aren't
	// supported, so specify exact bucket names (for example, 'my-bucket'). Each bucket
	// must already exist in the feature configuration's project. Up to 1,000 buckets can
	// be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_storage_control_project_feature_config#bucket_id_regexes GoogleStorageControlProjectFeatureConfig#bucket_id_regexes}
	BucketIdRegexes *[]*string `field:"required" json:"bucketIdRegexes" yaml:"bucketIdRegexes"`
}

