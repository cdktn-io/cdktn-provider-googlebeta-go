// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigKubeletConfigCrashLoopBackOff struct {
	// The maximum duration the backoff delay can accrue to for container restarts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_container_node_pool#max_container_restart_period GoogleContainerNodePool#max_container_restart_period}
	MaxContainerRestartPeriod *string `field:"optional" json:"maxContainerRestartPeriod" yaml:"maxContainerRestartPeriod"`
}

