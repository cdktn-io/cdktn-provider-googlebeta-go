// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlecesapp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesAppErrorHandlingSettingsOutputReference interface {
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
	EndSessionConfig() GoogleCesAppErrorHandlingSettingsEndSessionConfigOutputReference
	EndSessionConfigInput() *GoogleCesAppErrorHandlingSettingsEndSessionConfig
	ErrorHandlingStrategy() *string
	SetErrorHandlingStrategy(val *string)
	ErrorHandlingStrategyInput() *string
	FallbackResponseConfig() GoogleCesAppErrorHandlingSettingsFallbackResponseConfigOutputReference
	FallbackResponseConfigInput() *GoogleCesAppErrorHandlingSettingsFallbackResponseConfig
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesAppErrorHandlingSettings
	SetInternalValue(val *GoogleCesAppErrorHandlingSettings)
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
	PutEndSessionConfig(value *GoogleCesAppErrorHandlingSettingsEndSessionConfig)
	PutFallbackResponseConfig(value *GoogleCesAppErrorHandlingSettingsFallbackResponseConfig)
	ResetEndSessionConfig()
	ResetErrorHandlingStrategy()
	ResetFallbackResponseConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppErrorHandlingSettingsOutputReference
type jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) EndSessionConfig() GoogleCesAppErrorHandlingSettingsEndSessionConfigOutputReference {
	var returns GoogleCesAppErrorHandlingSettingsEndSessionConfigOutputReference
	_jsii_.Get(
		j,
		"endSessionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) EndSessionConfigInput() *GoogleCesAppErrorHandlingSettingsEndSessionConfig {
	var returns *GoogleCesAppErrorHandlingSettingsEndSessionConfig
	_jsii_.Get(
		j,
		"endSessionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) ErrorHandlingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorHandlingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) ErrorHandlingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorHandlingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) FallbackResponseConfig() GoogleCesAppErrorHandlingSettingsFallbackResponseConfigOutputReference {
	var returns GoogleCesAppErrorHandlingSettingsFallbackResponseConfigOutputReference
	_jsii_.Get(
		j,
		"fallbackResponseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) FallbackResponseConfigInput() *GoogleCesAppErrorHandlingSettingsFallbackResponseConfig {
	var returns *GoogleCesAppErrorHandlingSettingsFallbackResponseConfig
	_jsii_.Get(
		j,
		"fallbackResponseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) InternalValue() *GoogleCesAppErrorHandlingSettings {
	var returns *GoogleCesAppErrorHandlingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesAppErrorHandlingSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAppErrorHandlingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppErrorHandlingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppErrorHandlingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAppErrorHandlingSettingsOutputReference_Override(g GoogleCesAppErrorHandlingSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppErrorHandlingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference)SetErrorHandlingStrategy(val *string) {
	if err := j.validateSetErrorHandlingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorHandlingStrategy",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference)SetInternalValue(val *GoogleCesAppErrorHandlingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) PutEndSessionConfig(value *GoogleCesAppErrorHandlingSettingsEndSessionConfig) {
	if err := g.validatePutEndSessionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndSessionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) PutFallbackResponseConfig(value *GoogleCesAppErrorHandlingSettingsFallbackResponseConfig) {
	if err := g.validatePutFallbackResponseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFallbackResponseConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) ResetEndSessionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEndSessionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) ResetErrorHandlingStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetErrorHandlingStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) ResetFallbackResponseConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetFallbackResponseConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppErrorHandlingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

