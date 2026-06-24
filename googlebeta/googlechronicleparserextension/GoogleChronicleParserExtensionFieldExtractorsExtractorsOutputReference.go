// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparserextension

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlechronicleparserextension/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference interface {
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

// The jsii proxy struct for GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference
type jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) DestinationPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) DestinationPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) FieldPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) FieldPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) PreconditionOp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionOp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) PreconditionOpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionOpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) PreconditionPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) PreconditionPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) PreconditionValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) PreconditionValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewGoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleParserExtension.GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference_Override(g GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleParserExtension.GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetDestinationPath(val *string) {
	if err := j.validateSetDestinationPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPath",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetFieldPath(val *string) {
	if err := j.validateSetFieldPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldPath",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetPreconditionOp(val *string) {
	if err := j.validateSetPreconditionOpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconditionOp",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetPreconditionPath(val *string) {
	if err := j.validateSetPreconditionPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconditionPath",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetPreconditionValue(val *string) {
	if err := j.validateSetPreconditionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconditionValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ResetDestinationPath() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ResetFieldPath() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ResetPreconditionOp() {
	_jsii_.InvokeVoid(
		g,
		"resetPreconditionOp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ResetPreconditionPath() {
	_jsii_.InvokeVoid(
		g,
		"resetPreconditionPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ResetPreconditionValue() {
	_jsii_.InvokeVoid(
		g,
		"resetPreconditionValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		g,
		"resetValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsExtractorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

