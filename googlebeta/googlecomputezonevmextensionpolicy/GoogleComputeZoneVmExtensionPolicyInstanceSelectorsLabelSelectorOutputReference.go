// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputezonevmextensionpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlecomputezonevmextensionpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InclusionLabels() *map[string]*string
	SetInclusionLabels(val *map[string]*string)
	InclusionLabelsInput() *map[string]*string
	InternalValue() *GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelector
	SetInternalValue(val *GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelector)
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
	ResetInclusionLabels()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference
type jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) InclusionLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inclusionLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) InclusionLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inclusionLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) InternalValue() *GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelector {
	var returns *GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelector
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeZoneVmExtensionPolicy.GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference_Override(g GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeZoneVmExtensionPolicy.GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference)SetInclusionLabels(val *map[string]*string) {
	if err := j.validateSetInclusionLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionLabels",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference)SetInternalValue(val *GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelector) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) ResetInclusionLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetInclusionLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelectorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

