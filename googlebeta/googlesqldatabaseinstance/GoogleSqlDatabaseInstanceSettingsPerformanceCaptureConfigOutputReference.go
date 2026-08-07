// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlesqldatabaseinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfig
	SetInternalValue(val *GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfig)
	ProbeThreshold() *float64
	SetProbeThreshold(val *float64)
	ProbeThresholdInput() *float64
	ProbingIntervalSeconds() *float64
	SetProbingIntervalSeconds(val *float64)
	ProbingIntervalSecondsInput() *float64
	RunningThreadsThreshold() *float64
	SetRunningThreadsThreshold(val *float64)
	RunningThreadsThresholdInput() *float64
	SecondsBehindSourceThreshold() *float64
	SetSecondsBehindSourceThreshold(val *float64)
	SecondsBehindSourceThresholdInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TransactionDurationThreshold() *float64
	SetTransactionDurationThreshold(val *float64)
	TransactionDurationThresholdInput() *float64
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
	ResetEnabled()
	ResetProbeThreshold()
	ResetProbingIntervalSeconds()
	ResetRunningThreadsThreshold()
	ResetSecondsBehindSourceThreshold()
	ResetTransactionDurationThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference
type jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) InternalValue() *GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfig {
	var returns *GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ProbeThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probeThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ProbeThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probeThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ProbingIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probingIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ProbingIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probingIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) RunningThreadsThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningThreadsThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) RunningThreadsThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningThreadsThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) SecondsBehindSourceThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsBehindSourceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) SecondsBehindSourceThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsBehindSourceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) TransactionDurationThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionDurationThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) TransactionDurationThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionDurationThresholdInput",
		&returns,
	)
	return returns
}


func NewGoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference_Override(g GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetInternalValue(val *GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetProbeThreshold(val *float64) {
	if err := j.validateSetProbeThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probeThreshold",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetProbingIntervalSeconds(val *float64) {
	if err := j.validateSetProbingIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probingIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetRunningThreadsThreshold(val *float64) {
	if err := j.validateSetRunningThreadsThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runningThreadsThreshold",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetSecondsBehindSourceThreshold(val *float64) {
	if err := j.validateSetSecondsBehindSourceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondsBehindSourceThreshold",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference)SetTransactionDurationThreshold(val *float64) {
	if err := j.validateSetTransactionDurationThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionDurationThreshold",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ResetProbeThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetProbeThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ResetProbingIntervalSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetProbingIntervalSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ResetRunningThreadsThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetRunningThreadsThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ResetSecondsBehindSourceThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondsBehindSourceThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ResetTransactionDurationThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetTransactionDurationThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsPerformanceCaptureConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

