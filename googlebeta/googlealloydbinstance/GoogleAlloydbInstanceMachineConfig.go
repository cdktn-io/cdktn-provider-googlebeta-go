// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlealloydbinstance


type GoogleAlloydbInstanceMachineConfig struct {
	// The number of CPUs in the VM instance.
	//
	// For read pool instances, this
	// value is applied to the instances in the pool and is not replaced by
	// a fixed default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_alloydb_instance#cpu_count GoogleAlloydbInstance#cpu_count}
	CpuCount *float64 `field:"optional" json:"cpuCount" yaml:"cpuCount"`
	// Machine type of the VM instance. E.g. "n2-highmem-4", "n2-highmem-8", "c4a-highmem-4-lssd". 'cpu_count' must match the number of vCPUs in the machine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_alloydb_instance#machine_type GoogleAlloydbInstance#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
}

