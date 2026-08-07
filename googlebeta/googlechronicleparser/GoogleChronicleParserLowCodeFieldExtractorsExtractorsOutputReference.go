// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlechronicleparser/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference interface {
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
	DestinationPath() *string
	SetDestinationPath(val *string)
	DestinationPathInput() *string
	FieldPath() *string
	SetFieldPath(val *string)
	FieldPathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PreconditionOp() *string
	SetPreconditionOp(val *string)
	PreconditionOpInput() *string
	PreconditionPath() *string
	SetPreconditionPath(val *string)
	PreconditionPathInput() *string
	PreconditionValue() *string
	SetPreconditionValue(val *string)
	PreconditionValueInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetDestinationPath()
	ResetFieldPath()
	ResetPreconditionOp()
	ResetPreconditionPath()
	ResetPreconditionValue()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference
type jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) DestinationPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) DestinationPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) FieldPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) FieldPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionOp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionOp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionOpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionOpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewGoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleParser.GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference_Override(g GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleParser.GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetDestinationPath(val *string) {
	if err := j.validateSetDestinationPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPath",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetFieldPath(val *string) {
	if err := j.validateSetFieldPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldPath",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetPreconditionOp(val *string) {
	if err := j.validateSetPreconditionOpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconditionOp",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetPreconditionPath(val *string) {
	if err := j.validateSetPreconditionPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconditionPath",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetPreconditionValue(val *string) {
	if err := j.validateSetPreconditionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconditionValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetDestinationPath() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetFieldPath() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetPreconditionOp() {
	_jsii_.InvokeVoid(
		g,
		"resetPreconditionOp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetPreconditionPath() {
	_jsii_.InvokeVoid(
		g,
		"resetPreconditionPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetPreconditionValue() {
	_jsii_.InvokeVoid(
		g,
		"resetPreconditionValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		g,
		"resetValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

