// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlecesapp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference interface {
	cdktn.ComplexObject
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	EnableOriginCheck() interface{}
	SetEnableOriginCheck(val interface{})
	EnableOriginCheckInput() interface{}
	EnablePublicAccess() interface{}
	SetEnablePublicAccess(val interface{})
	EnablePublicAccessInput() interface{}
	EnableRecaptcha() interface{}
	SetEnableRecaptcha(val interface{})
	EnableRecaptchaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettings
	SetInternalValue(val *GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettings)
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
	ResetAllowedOrigins()
	ResetEnableOriginCheck()
	ResetEnablePublicAccess()
	ResetEnableRecaptcha()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference
type jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnableOriginCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOriginCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnableOriginCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOriginCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnablePublicAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnablePublicAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnableRecaptcha() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRecaptcha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnableRecaptchaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRecaptchaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) InternalValue() *GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettings {
	var returns *GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference_Override(g GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetAllowedOrigins(val *[]*string) {
	if err := j.validateSetAllowedOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetEnableOriginCheck(val interface{}) {
	if err := j.validateSetEnableOriginCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOriginCheck",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetEnablePublicAccess(val interface{}) {
	if err := j.validateSetEnablePublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePublicAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetEnableRecaptcha(val interface{}) {
	if err := j.validateSetEnableRecaptchaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRecaptcha",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetInternalValue(val *GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ResetAllowedOrigins() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedOrigins",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ResetEnableOriginCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableOriginCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ResetEnablePublicAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePublicAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ResetEnableRecaptcha() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableRecaptcha",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

