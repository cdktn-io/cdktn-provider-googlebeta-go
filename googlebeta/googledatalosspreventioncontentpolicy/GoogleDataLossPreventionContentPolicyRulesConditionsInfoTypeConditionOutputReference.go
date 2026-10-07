// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googledatalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference interface {
	cdktn.ComplexObject
	AnyInfoType() GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionAnyInfoTypeOutputReference
	AnyInfoTypeInput() *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionAnyInfoType
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
	InfoTypes() GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypesOutputReference
	InfoTypesInput() *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypes
	InternalValue() *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeCondition
	SetInternalValue(val *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeCondition)
	MinCount() *float64
	SetMinCount(val *float64)
	MinCountInput() *float64
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
	PutAnyInfoType(value *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionAnyInfoType)
	PutInfoTypes(value *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypes)
	ResetAnyInfoType()
	ResetInfoTypes()
	ResetMinCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference
type jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) AnyInfoType() GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionAnyInfoTypeOutputReference {
	var returns GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionAnyInfoTypeOutputReference
	_jsii_.Get(
		j,
		"anyInfoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) AnyInfoTypeInput() *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionAnyInfoType {
	var returns *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionAnyInfoType
	_jsii_.Get(
		j,
		"anyInfoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) InfoTypes() GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypesOutputReference {
	var returns GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypesOutputReference
	_jsii_.Get(
		j,
		"infoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) InfoTypesInput() *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypes {
	var returns *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypes
	_jsii_.Get(
		j,
		"infoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) InternalValue() *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeCondition {
	var returns *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) MinCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) MinCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference_Override(g GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference)SetInternalValue(val *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference)SetMinCount(val *float64) {
	if err := j.validateSetMinCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCount",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) PutAnyInfoType(value *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionAnyInfoType) {
	if err := g.validatePutAnyInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnyInfoType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) PutInfoTypes(value *GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionInfoTypes) {
	if err := g.validatePutInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInfoTypes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) ResetAnyInfoType() {
	_jsii_.InvokeVoid(
		g,
		"resetAnyInfoType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) ResetInfoTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetInfoTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) ResetMinCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyRulesConditionsInfoTypeConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

