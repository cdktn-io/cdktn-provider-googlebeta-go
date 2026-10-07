// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix struct {
	// Optional. The base capacity that will always use Standard VMs to avoid risk of premature allocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#standard_capacity_base GoogleDataprocWorkflowTemplate#standard_capacity_base}
	StandardCapacityBase *float64 `field:"optional" json:"standardCapacityBase" yaml:"standardCapacityBase"`
	// Optional. The percentage of target capacity that will use Standard VMs above standardCapacityBase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dataproc_workflow_template#standard_capacity_percent_above_base GoogleDataprocWorkflowTemplate#standard_capacity_percent_above_base}
	StandardCapacityPercentAboveBase *float64 `field:"optional" json:"standardCapacityPercentAboveBase" yaml:"standardCapacityPercentAboveBase"`
}

