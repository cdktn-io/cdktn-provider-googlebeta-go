// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlecestool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesToolWidgetToolOutputReference interface {
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
	DataMapping() GoogleCesToolWidgetToolDataMappingOutputReference
	DataMappingInput() *GoogleCesToolWidgetToolDataMapping
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesToolWidgetTool
	SetInternalValue(val *GoogleCesToolWidgetTool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameters() GoogleCesToolWidgetToolParametersOutputReference
	ParametersInput() *GoogleCesToolWidgetToolParameters
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextResponseConfig() GoogleCesToolWidgetToolTextResponseConfigOutputReference
	TextResponseConfigInput() *GoogleCesToolWidgetToolTextResponseConfig
	UiConfig() *string
	SetUiConfig(val *string)
	UiConfigInput() *string
	WidgetType() *string
	SetWidgetType(val *string)
	WidgetTypeInput() *string
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
	PutDataMapping(value *GoogleCesToolWidgetToolDataMapping)
	PutParameters(value *GoogleCesToolWidgetToolParameters)
	PutTextResponseConfig(value *GoogleCesToolWidgetToolTextResponseConfig)
	ResetDataMapping()
	ResetDescription()
	ResetParameters()
	ResetTextResponseConfig()
	ResetUiConfig()
	ResetWidgetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesToolWidgetToolOutputReference
type jsiiProxy_GoogleCesToolWidgetToolOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) DataMapping() GoogleCesToolWidgetToolDataMappingOutputReference {
	var returns GoogleCesToolWidgetToolDataMappingOutputReference
	_jsii_.Get(
		j,
		"dataMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) DataMappingInput() *GoogleCesToolWidgetToolDataMapping {
	var returns *GoogleCesToolWidgetToolDataMapping
	_jsii_.Get(
		j,
		"dataMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) InternalValue() *GoogleCesToolWidgetTool {
	var returns *GoogleCesToolWidgetTool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) Parameters() GoogleCesToolWidgetToolParametersOutputReference {
	var returns GoogleCesToolWidgetToolParametersOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ParametersInput() *GoogleCesToolWidgetToolParameters {
	var returns *GoogleCesToolWidgetToolParameters
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) TextResponseConfig() GoogleCesToolWidgetToolTextResponseConfigOutputReference {
	var returns GoogleCesToolWidgetToolTextResponseConfigOutputReference
	_jsii_.Get(
		j,
		"textResponseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) TextResponseConfigInput() *GoogleCesToolWidgetToolTextResponseConfig {
	var returns *GoogleCesToolWidgetToolTextResponseConfig
	_jsii_.Get(
		j,
		"textResponseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) UiConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) UiConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) WidgetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference) WidgetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetTypeInput",
		&returns,
	)
	return returns
}


func NewGoogleCesToolWidgetToolOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesToolWidgetToolOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesToolWidgetToolOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesToolWidgetToolOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolWidgetToolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesToolWidgetToolOutputReference_Override(g GoogleCesToolWidgetToolOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolWidgetToolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference)SetInternalValue(val *GoogleCesToolWidgetTool) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference)SetUiConfig(val *string) {
	if err := j.validateSetUiConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uiConfig",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolOutputReference)SetWidgetType(val *string) {
	if err := j.validateSetWidgetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetType",
		val,
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) PutDataMapping(value *GoogleCesToolWidgetToolDataMapping) {
	if err := g.validatePutDataMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataMapping",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) PutParameters(value *GoogleCesToolWidgetToolParameters) {
	if err := g.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParameters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) PutTextResponseConfig(value *GoogleCesToolWidgetToolTextResponseConfig) {
	if err := g.validatePutTextResponseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTextResponseConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ResetDataMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetDataMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ResetTextResponseConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTextResponseConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ResetUiConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetUiConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ResetWidgetType() {
	_jsii_.InvokeVoid(
		g,
		"resetWidgetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

