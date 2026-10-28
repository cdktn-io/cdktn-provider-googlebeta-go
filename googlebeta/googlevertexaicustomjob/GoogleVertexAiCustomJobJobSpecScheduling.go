// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob


type GoogleVertexAiCustomJobJobSpecScheduling struct {
	// Indicates if the job should retry for internal errors after the job starts running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#disable_retries GoogleVertexAiCustomJob#disable_retries}
	DisableRetries interface{} `field:"optional" json:"disableRetries" yaml:"disableRetries"`
	// The maximum duration that a job will wait for the requested resources to be provisioned if the scheduling strategy is set to FLEX_START.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#max_wait_duration GoogleVertexAiCustomJob#max_wait_duration}
	MaxWaitDuration *string `field:"optional" json:"maxWaitDuration" yaml:"maxWaitDuration"`
	// Restarts the entire CustomJob if a worker gets restarted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#restart_job_on_worker_restart GoogleVertexAiCustomJob#restart_job_on_worker_restart}
	RestartJobOnWorkerRestart interface{} `field:"optional" json:"restartJobOnWorkerRestart" yaml:"restartJobOnWorkerRestart"`
	// This determines which type of scheduling strategy to use. Possible values: ["STANDARD", "SPOT", "FLEX_START"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#strategy GoogleVertexAiCustomJob#strategy}
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// The maximum job running time.
	//
	// The default is 7 days. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#timeout GoogleVertexAiCustomJob#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

