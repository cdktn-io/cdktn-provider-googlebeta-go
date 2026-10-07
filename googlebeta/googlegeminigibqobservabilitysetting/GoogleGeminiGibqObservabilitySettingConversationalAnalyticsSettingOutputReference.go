// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminigibqobservabilitysetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlegeminigibqobservabilitysetting/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference interface {
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
	InternalValue() *GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSetting
	SetInternalValue(val *GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSetting)
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

// The jsii proxy struct for GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference
type jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) FeedbackEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"feedbackEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) FeedbackEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"feedbackEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) InternalValue() *GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSetting {
	var returns *GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) LoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) LoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) MetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) MetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) TracesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) TracesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracesEnabledInput",
		&returns,
	)
	return returns
}


func NewGoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGeminiGibqObservabilitySetting.GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference_Override(g GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGeminiGibqObservabilitySetting.GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference)SetFeedbackEnabled(val interface{}) {
	if err := j.validateSetFeedbackEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedbackEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference)SetInternalValue(val *GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSetting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference)SetLoggingEnabled(val interface{}) {
	if err := j.validateSetLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference)SetMetricsEnabled(val interface{}) {
	if err := j.validateSetMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference)SetTracesEnabled(val interface{}) {
	if err := j.validateSetTracesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracesEnabled",
		val,
	)
}

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) ResetFeedbackEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetFeedbackEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) ResetLoggingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) ResetMetricsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) ResetTracesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetTracesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

