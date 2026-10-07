// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlecesagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesAgentTransferRulesDeterministicTransferOutputReference interface {
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
	ExpressionCondition() GoogleCesAgentTransferRulesDeterministicTransferExpressionConditionOutputReference
	ExpressionConditionInput() *GoogleCesAgentTransferRulesDeterministicTransferExpressionCondition
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesAgentTransferRulesDeterministicTransfer
	SetInternalValue(val *GoogleCesAgentTransferRulesDeterministicTransfer)
	PythonCodeCondition() GoogleCesAgentTransferRulesDeterministicTransferPythonCodeConditionOutputReference
	PythonCodeConditionInput() *GoogleCesAgentTransferRulesDeterministicTransferPythonCodeCondition
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
	PutExpressionCondition(value *GoogleCesAgentTransferRulesDeterministicTransferExpressionCondition)
	PutPythonCodeCondition(value *GoogleCesAgentTransferRulesDeterministicTransferPythonCodeCondition)
	ResetExpressionCondition()
	ResetPythonCodeCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAgentTransferRulesDeterministicTransferOutputReference
type jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) ExpressionCondition() GoogleCesAgentTransferRulesDeterministicTransferExpressionConditionOutputReference {
	var returns GoogleCesAgentTransferRulesDeterministicTransferExpressionConditionOutputReference
	_jsii_.Get(
		j,
		"expressionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) ExpressionConditionInput() *GoogleCesAgentTransferRulesDeterministicTransferExpressionCondition {
	var returns *GoogleCesAgentTransferRulesDeterministicTransferExpressionCondition
	_jsii_.Get(
		j,
		"expressionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) InternalValue() *GoogleCesAgentTransferRulesDeterministicTransfer {
	var returns *GoogleCesAgentTransferRulesDeterministicTransfer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) PythonCodeCondition() GoogleCesAgentTransferRulesDeterministicTransferPythonCodeConditionOutputReference {
	var returns GoogleCesAgentTransferRulesDeterministicTransferPythonCodeConditionOutputReference
	_jsii_.Get(
		j,
		"pythonCodeCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) PythonCodeConditionInput() *GoogleCesAgentTransferRulesDeterministicTransferPythonCodeCondition {
	var returns *GoogleCesAgentTransferRulesDeterministicTransferPythonCodeCondition
	_jsii_.Get(
		j,
		"pythonCodeConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesAgentTransferRulesDeterministicTransferOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAgentTransferRulesDeterministicTransferOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAgentTransferRulesDeterministicTransferOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgentTransferRulesDeterministicTransferOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAgentTransferRulesDeterministicTransferOutputReference_Override(g GoogleCesAgentTransferRulesDeterministicTransferOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgentTransferRulesDeterministicTransferOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference)SetInternalValue(val *GoogleCesAgentTransferRulesDeterministicTransfer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) PutExpressionCondition(value *GoogleCesAgentTransferRulesDeterministicTransferExpressionCondition) {
	if err := g.validatePutExpressionConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExpressionCondition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) PutPythonCodeCondition(value *GoogleCesAgentTransferRulesDeterministicTransferPythonCodeCondition) {
	if err := g.validatePutPythonCodeConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPythonCodeCondition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) ResetExpressionCondition() {
	_jsii_.InvokeVoid(
		g,
		"resetExpressionCondition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) ResetPythonCodeCondition() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonCodeCondition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAgentTransferRulesDeterministicTransferOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

