// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob


type GoogleVertexAiCustomJobJobSpecWorkerPoolSpecs struct {
	// container_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#container_spec GoogleVertexAiCustomJob#container_spec}
	ContainerSpec *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec `field:"optional" json:"containerSpec" yaml:"containerSpec"`
	// disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#disk_spec GoogleVertexAiCustomJob#disk_spec}
	DiskSpec *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec `field:"optional" json:"diskSpec" yaml:"diskSpec"`
	// lustre_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#lustre_mounts GoogleVertexAiCustomJob#lustre_mounts}
	LustreMounts interface{} `field:"optional" json:"lustreMounts" yaml:"lustreMounts"`
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#machine_spec GoogleVertexAiCustomJob#machine_spec}
	MachineSpec *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec `field:"optional" json:"machineSpec" yaml:"machineSpec"`
	// nfs_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#nfs_mounts GoogleVertexAiCustomJob#nfs_mounts}
	NfsMounts interface{} `field:"optional" json:"nfsMounts" yaml:"nfsMounts"`
	// python_package_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#python_package_spec GoogleVertexAiCustomJob#python_package_spec}
	PythonPackageSpec *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec `field:"optional" json:"pythonPackageSpec" yaml:"pythonPackageSpec"`
	// The number of worker replicas to use for this worker pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#replica_count GoogleVertexAiCustomJob#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
}

