// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglebigqueryroutineiampolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleBigqueryRoutineIamPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/data-sources/google_bigquery_routine_iam_policy#dataset_id DataGoogleBigqueryRoutineIamPolicy#dataset_id}.
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/data-sources/google_bigquery_routine_iam_policy#routine_id DataGoogleBigqueryRoutineIamPolicy#routine_id}.
	RoutineId *string `field:"required" json:"routineId" yaml:"routineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/data-sources/google_bigquery_routine_iam_policy#id DataGoogleBigqueryRoutineIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/data-sources/google_bigquery_routine_iam_policy#project DataGoogleBigqueryRoutineIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

