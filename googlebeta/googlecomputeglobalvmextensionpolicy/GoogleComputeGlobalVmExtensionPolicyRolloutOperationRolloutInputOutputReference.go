// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeglobalvmextensionpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlecomputeglobalvmextensionpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference interface {
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
	ConflictBehavior() *string
	SetConflictBehavior(val *string)
	ConflictBehaviorInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput
	SetInternalValue(val *GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput)
	Name() *string
	SetName(val *string)
	NameInput() *string
	PredefinedRolloutPlan() *string
	SetPredefinedRolloutPlan(val *string)
	PredefinedRolloutPlanInput() *string
	RetryUuid() *string
	SetRetryUuid(val *string)
	RetryUuidInput() *string
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
	ResetConflictBehavior()
	ResetName()
	ResetPredefinedRolloutPlan()
	ResetRetryUuid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference
type jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ConflictBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ConflictBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) InternalValue() *GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput {
	var returns *GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) PredefinedRolloutPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedRolloutPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) PredefinedRolloutPlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedRolloutPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) RetryUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) RetryUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeGlobalVmExtensionPolicy.GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference_Override(g GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeGlobalVmExtensionPolicy.GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetConflictBehavior(val *string) {
	if err := j.validateSetConflictBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conflictBehavior",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetInternalValue(val *GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetPredefinedRolloutPlan(val *string) {
	if err := j.validateSetPredefinedRolloutPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predefinedRolloutPlan",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetRetryUuid(val *string) {
	if err := j.validateSetRetryUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryUuid",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ResetConflictBehavior() {
	_jsii_.InvokeVoid(
		g,
		"resetConflictBehavior",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ResetPredefinedRolloutPlan() {
	_jsii_.InvokeVoid(
		g,
		"resetPredefinedRolloutPlan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ResetRetryUuid() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryUuid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

