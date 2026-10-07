// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowtool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googledialogflowtool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDialogflowToolConnectorSpecActionsOutputReference interface {
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
	ConnectionActionId() *string
	SetConnectionActionId(val *string)
	ConnectionActionIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EntityOperation() GoogleDialogflowToolConnectorSpecActionsEntityOperationOutputReference
	EntityOperationInput() *GoogleDialogflowToolConnectorSpecActionsEntityOperation
	// Experimental.
	Fqn() *string
	InputFields() *[]*string
	SetInputFields(val *[]*string)
	InputFieldsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputFields() *[]*string
	SetOutputFields(val *[]*string)
	OutputFieldsInput() *[]*string
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
	PutEntityOperation(value *GoogleDialogflowToolConnectorSpecActionsEntityOperation)
	ResetConnectionActionId()
	ResetEntityOperation()
	ResetInputFields()
	ResetOutputFields()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowToolConnectorSpecActionsOutputReference
type jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ConnectionActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionActionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ConnectionActionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionActionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) EntityOperation() GoogleDialogflowToolConnectorSpecActionsEntityOperationOutputReference {
	var returns GoogleDialogflowToolConnectorSpecActionsEntityOperationOutputReference
	_jsii_.Get(
		j,
		"entityOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) EntityOperationInput() *GoogleDialogflowToolConnectorSpecActionsEntityOperation {
	var returns *GoogleDialogflowToolConnectorSpecActionsEntityOperation
	_jsii_.Get(
		j,
		"entityOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) InputFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) InputFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) OutputFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) OutputFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowToolConnectorSpecActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDialogflowToolConnectorSpecActionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowToolConnectorSpecActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowTool.GoogleDialogflowToolConnectorSpecActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowToolConnectorSpecActionsOutputReference_Override(g GoogleDialogflowToolConnectorSpecActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowTool.GoogleDialogflowToolConnectorSpecActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference)SetConnectionActionId(val *string) {
	if err := j.validateSetConnectionActionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionActionId",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference)SetInputFields(val *[]*string) {
	if err := j.validateSetInputFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFields",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference)SetOutputFields(val *[]*string) {
	if err := j.validateSetOutputFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFields",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) PutEntityOperation(value *GoogleDialogflowToolConnectorSpecActionsEntityOperation) {
	if err := g.validatePutEntityOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEntityOperation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ResetConnectionActionId() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionActionId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ResetEntityOperation() {
	_jsii_.InvokeVoid(
		g,
		"resetEntityOperation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ResetInputFields() {
	_jsii_.InvokeVoid(
		g,
		"resetInputFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ResetOutputFields() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowToolConnectorSpecActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

