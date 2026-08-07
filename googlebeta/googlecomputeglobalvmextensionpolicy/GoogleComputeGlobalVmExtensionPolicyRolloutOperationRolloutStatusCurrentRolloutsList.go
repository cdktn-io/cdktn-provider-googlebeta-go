// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeglobalvmextensionpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlecomputeglobalvmextensionpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList
type jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList {
	_init_.Initialize()

	if err := validateNewGoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeGlobalVmExtensionPolicy.GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList_Override(g GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeGlobalVmExtensionPolicy.GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) Get(index *float64) GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutStatusCurrentRolloutsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

