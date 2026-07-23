// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivetable


type GoogleBiglakeHiveTableStorageDescriptor struct {
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#columns GoogleBiglakeHiveTable#columns}
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// Reducer grouping columns, clustering columns, and bucketing columns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#bucket_cols GoogleBiglakeHiveTable#bucket_cols}
	BucketCols *[]*string `field:"optional" json:"bucketCols" yaml:"bucketCols"`
	// Whether the table data is compressed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#compressed GoogleBiglakeHiveTable#compressed}
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// The fully qualified Java class name of the input format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#input_format GoogleBiglakeHiveTable#input_format}
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// The Cloud Storage URI where the table data is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#location_uri GoogleBiglakeHiveTable#location_uri}
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
	// The number of buckets in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#num_buckets GoogleBiglakeHiveTable#num_buckets}
	NumBuckets *float64 `field:"optional" json:"numBuckets" yaml:"numBuckets"`
	// The fully qualified Java class name of the output format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#output_format GoogleBiglakeHiveTable#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Key-value pairs for the storage descriptor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#parameters GoogleBiglakeHiveTable#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// serde_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#serde_info GoogleBiglakeHiveTable#serde_info}
	SerdeInfo *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo `field:"optional" json:"serdeInfo" yaml:"serdeInfo"`
	// skewed_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#skewed_info GoogleBiglakeHiveTable#skewed_info}
	SkewedInfo *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// sort_cols block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#sort_cols GoogleBiglakeHiveTable#sort_cols}
	SortCols interface{} `field:"optional" json:"sortCols" yaml:"sortCols"`
	// Whether the table is stored as sub directories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#stored_as_sub_dirs GoogleBiglakeHiveTable#stored_as_sub_dirs}
	StoredAsSubDirs interface{} `field:"optional" json:"storedAsSubDirs" yaml:"storedAsSubDirs"`
}

