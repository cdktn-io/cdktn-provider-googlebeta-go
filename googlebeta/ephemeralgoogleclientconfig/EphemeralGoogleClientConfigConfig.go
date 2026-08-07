// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralgoogleclientconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralGoogleClientConfigConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
}

