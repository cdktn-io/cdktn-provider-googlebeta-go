// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlevertexaionlineevaluator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference interface {
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
	InternalValue() *GoogleVertexAiOnlineEvaluatorCloudObservability
	SetInternalValue(val *GoogleVertexAiOnlineEvaluatorCloudObservability)
	LogView() *string
	SetLogView(val *string)
	LogViewInput() *string
	OpenTelemetry() GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetryOutputReference
	OpenTelemetryInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry
	SessionScope() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference
	SessionScopeInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScope
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TraceScope() GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeOutputReference
	TraceScopeInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScope
	TraceView() *string
	SetTraceView(val *string)
	TraceViewInput() *string
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
	PutOpenTelemetry(value *GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry)
	PutSessionScope(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScope)
	PutTraceScope(value *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScope)
	ResetLogView()
	ResetSessionScope()
	ResetTraceScope()
	ResetTraceView()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference
type jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) InternalValue() *GoogleVertexAiOnlineEvaluatorCloudObservability {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservability
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) LogView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) LogViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logViewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) OpenTelemetry() GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetryOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetryOutputReference
	_jsii_.Get(
		j,
		"openTelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) OpenTelemetryInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry
	_jsii_.Get(
		j,
		"openTelemetryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) SessionScope() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference
	_jsii_.Get(
		j,
		"sessionScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) SessionScopeInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScope {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScope
	_jsii_.Get(
		j,
		"sessionScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) TraceScope() GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeOutputReference
	_jsii_.Get(
		j,
		"traceScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) TraceScopeInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScope {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScope
	_jsii_.Get(
		j,
		"traceScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) TraceView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) TraceViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceViewInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiOnlineEvaluator.GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference_Override(g GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiOnlineEvaluator.GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetInternalValue(val *GoogleVertexAiOnlineEvaluatorCloudObservability) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetLogView(val *string) {
	if err := j.validateSetLogViewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logView",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetTraceView(val *string) {
	if err := j.validateSetTraceViewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceView",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) PutOpenTelemetry(value *GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry) {
	if err := g.validatePutOpenTelemetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOpenTelemetry",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) PutSessionScope(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScope) {
	if err := g.validatePutSessionScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSessionScope",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) PutTraceScope(value *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScope) {
	if err := g.validatePutTraceScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTraceScope",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) ResetLogView() {
	_jsii_.InvokeVoid(
		g,
		"resetLogView",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) ResetSessionScope() {
	_jsii_.InvokeVoid(
		g,
		"resetSessionScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) ResetTraceScope() {
	_jsii_.InvokeVoid(
		g,
		"resetTraceScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) ResetTraceView() {
	_jsii_.InvokeVoid(
		g,
		"resetTraceView",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

