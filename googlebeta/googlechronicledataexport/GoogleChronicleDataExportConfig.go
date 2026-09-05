// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicledataexport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleDataExportConfig struct {
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
	// Last, exclusive time from the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#end_time GoogleChronicleDataExport#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Link to the destination Cloud Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#gcs_bucket GoogleChronicleDataExport#gcs_bucket}
	GcsBucket *string `field:"required" json:"gcsBucket" yaml:"gcsBucket"`
	// The unique identifier for the Chronicle instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#instance GoogleChronicleDataExport#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#location GoogleChronicleDataExport#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Start, inclusive time from the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#start_time GoogleChronicleDataExport#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#id GoogleChronicleDataExport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The specific log types to include in the Data Export request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#include_log_types GoogleChronicleDataExport#include_log_types}
	IncludeLogTypes *[]*string `field:"optional" json:"includeLogTypes" yaml:"includeLogTypes"`
	// ingestion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#ingestion_labels GoogleChronicleDataExport#ingestion_labels}
	IngestionLabels interface{} `field:"optional" json:"ingestionLabels" yaml:"ingestionLabels"`
	// The namespaces used to filter the export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#namespaces GoogleChronicleDataExport#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#project GoogleChronicleDataExport#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_chronicle_data_export#timeouts GoogleChronicleDataExport#timeouts}
	Timeouts *GoogleChronicleDataExportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

