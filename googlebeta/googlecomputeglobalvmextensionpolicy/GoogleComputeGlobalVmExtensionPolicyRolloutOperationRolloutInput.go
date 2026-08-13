// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeglobalvmextensionpolicy


type GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput struct {
	// Specifies the behavior of the rollout if a conflict is detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_compute_global_vm_extension_policy#conflict_behavior GoogleComputeGlobalVmExtensionPolicy#conflict_behavior}
	ConflictBehavior *string `field:"optional" json:"conflictBehavior" yaml:"conflictBehavior"`
	// The name of the rollout plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_compute_global_vm_extension_policy#name GoogleComputeGlobalVmExtensionPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the predefined rollout plan for the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_compute_global_vm_extension_policy#predefined_rollout_plan GoogleComputeGlobalVmExtensionPolicy#predefined_rollout_plan}
	PredefinedRolloutPlan *string `field:"optional" json:"predefinedRolloutPlan" yaml:"predefinedRolloutPlan"`
	// The UUID that identifies a policy rollout retry attempt.
	//
	// It should only be set when retrying an existing rollout. Updating this field along with other policy fields (description, extension_policies, instance_selectors, priority) in the same plan will return an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_compute_global_vm_extension_policy#retry_uuid GoogleComputeGlobalVmExtensionPolicy#retry_uuid}
	RetryUuid *string `field:"optional" json:"retryUuid" yaml:"retryUuid"`
}

