// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclenativedashboard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleNativeDashboardConfig struct {
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
	// The display name/title of the dashboard visible to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#display_name GoogleChronicleNativeDashboard#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID of the Chronicle instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#instance GoogleChronicleNativeDashboard#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the Chronicle instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#location GoogleChronicleNativeDashboard#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The access level of the dashboard. Possible values: DASHBOARD_PRIVATE DASHBOARD_PUBLIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#access GoogleChronicleNativeDashboard#access}
	Access *string `field:"optional" json:"access" yaml:"access"`
	// charts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#charts GoogleChronicleNativeDashboard#charts}
	Charts interface{} `field:"optional" json:"charts" yaml:"charts"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#deletion_policy GoogleChronicleNativeDashboard#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// A description of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#description GoogleChronicleNativeDashboard#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#filters GoogleChronicleNativeDashboard#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#id GoogleChronicleNativeDashboard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the dashboard is pinned by the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#is_pinned GoogleChronicleNativeDashboard#is_pinned}
	IsPinned interface{} `field:"optional" json:"isPinned" yaml:"isPinned"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#project GoogleChronicleNativeDashboard#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#timeouts GoogleChronicleNativeDashboard#timeouts}
	Timeouts *GoogleChronicleNativeDashboardTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of dashboard. Possible values: CURATED, PRIVATE, PUBLIC, CUSTOM, MARKETPLACE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_chronicle_native_dashboard#type GoogleChronicleNativeDashboard#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

