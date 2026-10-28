// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob


type GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec struct {
	// The URI of a container image in Artifact Registry that will run the provided Python package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#executor_image_uri GoogleVertexAiCustomJob#executor_image_uri}
	ExecutorImageUri *string `field:"required" json:"executorImageUri" yaml:"executorImageUri"`
	// The Google Cloud Storage location of the Python package files which are the training program and its dependent packages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#package_uris GoogleVertexAiCustomJob#package_uris}
	PackageUris *[]*string `field:"required" json:"packageUris" yaml:"packageUris"`
	// The Python module name to run after installing the packages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#python_module GoogleVertexAiCustomJob#python_module}
	PythonModule *string `field:"required" json:"pythonModule" yaml:"pythonModule"`
	// Command line arguments to be passed to the Python task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#args GoogleVertexAiCustomJob#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_custom_job#env GoogleVertexAiCustomJob#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
}

