// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicestelemetrypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlenetworkservicestelemetrypolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference interface {
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
	CustomSpanAttributes() GoogleNetworkServicesTelemetryPolicyTracingConfigurationCustomSpanAttributesList
	CustomSpanAttributesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetworkServicesTelemetryPolicyTracingConfiguration
	SetInternalValue(val *GoogleNetworkServicesTelemetryPolicyTracingConfiguration)
	ParentBasedSampling() GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSamplingOutputReference
	ParentBasedSamplingInput() *GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSampling
	SamplingRate() *float64
	SetSamplingRate(val *float64)
	SamplingRateInput() *float64
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
	PutCustomSpanAttributes(value interface{})
	PutParentBasedSampling(value *GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSampling)
	ResetCustomSpanAttributes()
	ResetParentBasedSampling()
	ResetSamplingRate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference
type jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) CustomSpanAttributes() GoogleNetworkServicesTelemetryPolicyTracingConfigurationCustomSpanAttributesList {
	var returns GoogleNetworkServicesTelemetryPolicyTracingConfigurationCustomSpanAttributesList
	_jsii_.Get(
		j,
		"customSpanAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) CustomSpanAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSpanAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) InternalValue() *GoogleNetworkServicesTelemetryPolicyTracingConfiguration {
	var returns *GoogleNetworkServicesTelemetryPolicyTracingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) ParentBasedSampling() GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSamplingOutputReference {
	var returns GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSamplingOutputReference
	_jsii_.Get(
		j,
		"parentBasedSampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) ParentBasedSamplingInput() *GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSampling {
	var returns *GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSampling
	_jsii_.Get(
		j,
		"parentBasedSamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) SamplingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) SamplingRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkServicesTelemetryPolicy.GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference_Override(g GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkServicesTelemetryPolicy.GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference)SetInternalValue(val *GoogleNetworkServicesTelemetryPolicyTracingConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference)SetSamplingRate(val *float64) {
	if err := j.validateSetSamplingRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplingRate",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) PutCustomSpanAttributes(value interface{}) {
	if err := g.validatePutCustomSpanAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomSpanAttributes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) PutParentBasedSampling(value *GoogleNetworkServicesTelemetryPolicyTracingConfigurationParentBasedSampling) {
	if err := g.validatePutParentBasedSamplingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParentBasedSampling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) ResetCustomSpanAttributes() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomSpanAttributes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) ResetParentBasedSampling() {
	_jsii_.InvokeVoid(
		g,
		"resetParentBasedSampling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) ResetSamplingRate() {
	_jsii_.InvokeVoid(
		g,
		"resetSamplingRate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesTelemetryPolicyTracingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

