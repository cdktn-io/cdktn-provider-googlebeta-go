// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagerserviceperimeters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleaccesscontextmanagerserviceperimeters/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Modifiers() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsModifiersList
	ModifiersInput() interface{}
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
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
	PutModifiers(value interface{})
	ResetModifiers()
	ResetPattern()
	ResetService()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference
type jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) Modifiers() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsModifiersList {
	var returns GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsModifiersList
	_jsii_.Get(
		j,
		"modifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) ModifiersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAccessContextManagerServicePerimeters.GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference_Override(g GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAccessContextManagerServicePerimeters.GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) PutModifiers(value interface{}) {
	if err := g.validatePutModifiersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putModifiers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) ResetModifiers() {
	_jsii_.InvokeVoid(
		g,
		"resetModifiers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) ResetPattern() {
	_jsii_.InvokeVoid(
		g,
		"resetPattern",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		g,
		"resetService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesAllowedServicePatternsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

