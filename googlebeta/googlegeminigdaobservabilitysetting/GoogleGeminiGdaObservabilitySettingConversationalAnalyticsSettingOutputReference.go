// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminigdaobservabilitysetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlegeminigdaobservabilitysetting/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference interface {
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
	FeedbackEnabled() interface{}
	SetFeedbackEnabled(val interface{})
	FeedbackEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSetting
	SetInternalValue(val *GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSetting)
	LoggingEnabled() interface{}
	SetLoggingEnabled(val interface{})
	LoggingEnabledInput() interface{}
	MetricsEnabled() interface{}
	SetMetricsEnabled(val interface{})
	MetricsEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TracesEnabled() interface{}
	SetTracesEnabled(val interface{})
	TracesEnabledInput() interface{}
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
	ResetFeedbackEnabled()
	ResetLoggingEnabled()
	ResetMetricsEnabled()
	ResetTracesEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference
type jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) FeedbackEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"feedbackEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) FeedbackEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"feedbackEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) InternalValue() *GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSetting {
	var returns *GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) LoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) LoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) MetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) MetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) TracesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) TracesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracesEnabledInput",
		&returns,
	)
	return returns
}


func NewGoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGeminiGdaObservabilitySetting.GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference_Override(g GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGeminiGdaObservabilitySetting.GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference)SetFeedbackEnabled(val interface{}) {
	if err := j.validateSetFeedbackEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedbackEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference)SetInternalValue(val *GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSetting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference)SetLoggingEnabled(val interface{}) {
	if err := j.validateSetLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference)SetMetricsEnabled(val interface{}) {
	if err := j.validateSetMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference)SetTracesEnabled(val interface{}) {
	if err := j.validateSetTracesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracesEnabled",
		val,
	)
}

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) ResetFeedbackEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetFeedbackEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) ResetLoggingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) ResetMetricsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) ResetTracesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetTracesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGeminiGdaObservabilitySettingConversationalAnalyticsSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

