// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeglobalvmextensionpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlecomputeglobalvmextensionpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExtensionName() *string
	SetExtensionName(val *string)
	ExtensionNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PinnedVersion() *string
	SetPinnedVersion(val *string)
	PinnedVersionInput() *string
	StringConfig() *string
	SetStringConfig(val *string)
	StringConfigInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetPinnedVersion()
	ResetStringConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference
type jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) ExtensionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) ExtensionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) PinnedVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinnedVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) PinnedVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinnedVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) StringConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) StringConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeGlobalVmExtensionPolicy.GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference_Override(g GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeGlobalVmExtensionPolicy.GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference)SetExtensionName(val *string) {
	if err := j.validateSetExtensionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference)SetPinnedVersion(val *string) {
	if err := j.validateSetPinnedVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinnedVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference)SetStringConfig(val *string) {
	if err := j.validateSetStringConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringConfig",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) ResetPinnedVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPinnedVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) ResetStringConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetStringConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyExtensionPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

