// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionnetworkpolicytrafficclassificationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlecomputeregionnetworkpolicytrafficclassificationrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference interface {
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
	DscpMode() *string
	SetDscpMode(val *string)
	DscpModeInput() *string
	DscpValue() *float64
	SetDscpValue(val *float64)
	DscpValueInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRegionNetworkPolicyTrafficClassificationRuleAction
	SetInternalValue(val *GoogleComputeRegionNetworkPolicyTrafficClassificationRuleAction)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrafficClass() *string
	SetTrafficClass(val *string)
	TrafficClassInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetDscpMode()
	ResetDscpValue()
	ResetTrafficClass()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference
type jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) DscpMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dscpMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) DscpModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dscpModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) DscpValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dscpValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) DscpValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dscpValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) InternalValue() *GoogleComputeRegionNetworkPolicyTrafficClassificationRuleAction {
	var returns *GoogleComputeRegionNetworkPolicyTrafficClassificationRuleAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) TrafficClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) TrafficClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionNetworkPolicyTrafficClassificationRule.GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference_Override(g GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionNetworkPolicyTrafficClassificationRule.GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference)SetDscpMode(val *string) {
	if err := j.validateSetDscpModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dscpMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference)SetDscpValue(val *float64) {
	if err := j.validateSetDscpValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dscpValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference)SetInternalValue(val *GoogleComputeRegionNetworkPolicyTrafficClassificationRuleAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference)SetTrafficClass(val *string) {
	if err := j.validateSetTrafficClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficClass",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) ResetDscpMode() {
	_jsii_.InvokeVoid(
		g,
		"resetDscpMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) ResetDscpValue() {
	_jsii_.InvokeVoid(
		g,
		"resetDscpValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) ResetTrafficClass() {
	_jsii_.InvokeVoid(
		g,
		"resetTrafficClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkPolicyTrafficClassificationRuleActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

