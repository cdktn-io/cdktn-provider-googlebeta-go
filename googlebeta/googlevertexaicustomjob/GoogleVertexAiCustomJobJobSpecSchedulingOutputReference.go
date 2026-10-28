// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexaicustomjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiCustomJobJobSpecSchedulingOutputReference interface {
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
	DisableRetries() interface{}
	SetDisableRetries(val interface{})
	DisableRetriesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiCustomJobJobSpecScheduling
	SetInternalValue(val *GoogleVertexAiCustomJobJobSpecScheduling)
	MaxWaitDuration() *string
	SetMaxWaitDuration(val *string)
	MaxWaitDurationInput() *string
	RestartJobOnWorkerRestart() interface{}
	SetRestartJobOnWorkerRestart(val interface{})
	RestartJobOnWorkerRestartInput() interface{}
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
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
	ResetDisableRetries()
	ResetMaxWaitDuration()
	ResetRestartJobOnWorkerRestart()
	ResetStrategy()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiCustomJobJobSpecSchedulingOutputReference
type jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) DisableRetries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) DisableRetriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) InternalValue() *GoogleVertexAiCustomJobJobSpecScheduling {
	var returns *GoogleVertexAiCustomJobJobSpecScheduling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) MaxWaitDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWaitDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) MaxWaitDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWaitDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) RestartJobOnWorkerRestart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartJobOnWorkerRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) RestartJobOnWorkerRestartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartJobOnWorkerRestartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiCustomJobJobSpecSchedulingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiCustomJobJobSpecSchedulingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiCustomJobJobSpecSchedulingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiCustomJobJobSpecSchedulingOutputReference_Override(g GoogleVertexAiCustomJobJobSpecSchedulingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetDisableRetries(val interface{}) {
	if err := j.validateSetDisableRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRetries",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetInternalValue(val *GoogleVertexAiCustomJobJobSpecScheduling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetMaxWaitDuration(val *string) {
	if err := j.validateSetMaxWaitDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWaitDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetRestartJobOnWorkerRestart(val interface{}) {
	if err := j.validateSetRestartJobOnWorkerRestartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartJobOnWorkerRestart",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) ResetDisableRetries() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableRetries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) ResetMaxWaitDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxWaitDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) ResetRestartJobOnWorkerRestart() {
	_jsii_.InvokeVoid(
		g,
		"resetRestartJobOnWorkerRestart",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecSchedulingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

