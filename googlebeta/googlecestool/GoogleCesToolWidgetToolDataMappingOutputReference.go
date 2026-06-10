// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlecestool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesToolWidgetToolDataMappingOutputReference interface {
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
	FieldMappings() *map[string]*string
	SetFieldMappings(val *map[string]*string)
	FieldMappingsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesToolWidgetToolDataMapping
	SetInternalValue(val *GoogleCesToolWidgetToolDataMapping)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	PythonFunction() GoogleCesToolWidgetToolDataMappingPythonFunctionOutputReference
	PythonFunctionInput() *GoogleCesToolWidgetToolDataMappingPythonFunction
	SourceToolName() *string
	SetSourceToolName(val *string)
	SourceToolNameInput() *string
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
	PutPythonFunction(value *GoogleCesToolWidgetToolDataMappingPythonFunction)
	ResetFieldMappings()
	ResetMode()
	ResetPythonFunction()
	ResetSourceToolName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesToolWidgetToolDataMappingOutputReference
type jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) FieldMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fieldMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) FieldMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"fieldMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) InternalValue() *GoogleCesToolWidgetToolDataMapping {
	var returns *GoogleCesToolWidgetToolDataMapping
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) PythonFunction() GoogleCesToolWidgetToolDataMappingPythonFunctionOutputReference {
	var returns GoogleCesToolWidgetToolDataMappingPythonFunctionOutputReference
	_jsii_.Get(
		j,
		"pythonFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) PythonFunctionInput() *GoogleCesToolWidgetToolDataMappingPythonFunction {
	var returns *GoogleCesToolWidgetToolDataMappingPythonFunction
	_jsii_.Get(
		j,
		"pythonFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) SourceToolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceToolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) SourceToolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceToolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesToolWidgetToolDataMappingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesToolWidgetToolDataMappingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesToolWidgetToolDataMappingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolWidgetToolDataMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesToolWidgetToolDataMappingOutputReference_Override(g GoogleCesToolWidgetToolDataMappingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolWidgetToolDataMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference)SetFieldMappings(val *map[string]*string) {
	if err := j.validateSetFieldMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldMappings",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference)SetInternalValue(val *GoogleCesToolWidgetToolDataMapping) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference)SetSourceToolName(val *string) {
	if err := j.validateSetSourceToolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceToolName",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) PutPythonFunction(value *GoogleCesToolWidgetToolDataMappingPythonFunction) {
	if err := g.validatePutPythonFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPythonFunction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) ResetFieldMappings() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldMappings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) ResetPythonFunction() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonFunction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) ResetSourceToolName() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceToolName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolDataMappingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

