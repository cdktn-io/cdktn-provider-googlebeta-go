// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource


type GoogleVertexAiPersistentResourceResourcePoolsMachineSpec struct {
	// The number of accelerators to attach to the machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_persistent_resource#accelerator_count GoogleVertexAiPersistentResource#accelerator_count}
	AcceleratorCount *float64 `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// The type of accelerator(s) that may be attached to the machine. Possible values: NVIDIA_TESLA_K80 NVIDIA_TESLA_P100 NVIDIA_TESLA_V100 NVIDIA_TESLA_P4 NVIDIA_TESLA_T4 NVIDIA_TESLA_A100 NVIDIA_A100_80GB NVIDIA_L4 NVIDIA_H100_80GB NVIDIA_H100_MEGA_80GB NVIDIA_H200_141GB NVIDIA_B200 NVIDIA_GB200 NVIDIA_RTX_PRO_6000 TPU_V2 TPU_V3 TPU_V4_POD TPU_V5_LITEPOD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_persistent_resource#accelerator_type GoogleVertexAiPersistentResource#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// The type of the machine.
	//
	// See the [list of machine types supported for
	// prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	// See the [list of machine types supported for custom
	// training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_persistent_resource#machine_type GoogleVertexAiPersistentResource#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
}

