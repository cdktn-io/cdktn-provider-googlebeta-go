// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleparserextension

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlechronicleparserextension/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleParserExtensionFieldExtractorsOutputReference interface {
	cdktn.ComplexObject
	AppendRepeatedFields() interface{}
	SetAppendRepeatedFields(val interface{})
	AppendRepeatedFieldsInput() interface{}
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
	Extractors() GoogleChronicleParserExtensionFieldExtractorsExtractorsList
	ExtractorsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleChronicleParserExtensionFieldExtractors
	SetInternalValue(val *GoogleChronicleParserExtensionFieldExtractors)
	LogFormat() *string
	SetLogFormat(val *string)
	LogFormatInput() *string
	PreprocessConfig() GoogleChronicleParserExtensionFieldExtractorsPreprocessConfigOutputReference
	PreprocessConfigInput() *GoogleChronicleParserExtensionFieldExtractorsPreprocessConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TransformedCbnSnippet() *string
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
	PutExtractors(value interface{})
	PutPreprocessConfig(value *GoogleChronicleParserExtensionFieldExtractorsPreprocessConfig)
	ResetAppendRepeatedFields()
	ResetExtractors()
	ResetLogFormat()
	ResetPreprocessConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleParserExtensionFieldExtractorsOutputReference
type jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) AppendRepeatedFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendRepeatedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) AppendRepeatedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendRepeatedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) Extractors() GoogleChronicleParserExtensionFieldExtractorsExtractorsList {
	var returns GoogleChronicleParserExtensionFieldExtractorsExtractorsList
	_jsii_.Get(
		j,
		"extractors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) ExtractorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) InternalValue() *GoogleChronicleParserExtensionFieldExtractors {
	var returns *GoogleChronicleParserExtensionFieldExtractors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) LogFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) LogFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) PreprocessConfig() GoogleChronicleParserExtensionFieldExtractorsPreprocessConfigOutputReference {
	var returns GoogleChronicleParserExtensionFieldExtractorsPreprocessConfigOutputReference
	_jsii_.Get(
		j,
		"preprocessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) PreprocessConfigInput() *GoogleChronicleParserExtensionFieldExtractorsPreprocessConfig {
	var returns *GoogleChronicleParserExtensionFieldExtractorsPreprocessConfig
	_jsii_.Get(
		j,
		"preprocessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) TransformedCbnSnippet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformedCbnSnippet",
		&returns,
	)
	return returns
}


func NewGoogleChronicleParserExtensionFieldExtractorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleParserExtensionFieldExtractorsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleParserExtensionFieldExtractorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleParserExtension.GoogleChronicleParserExtensionFieldExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleParserExtensionFieldExtractorsOutputReference_Override(g GoogleChronicleParserExtensionFieldExtractorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleParserExtension.GoogleChronicleParserExtensionFieldExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference)SetAppendRepeatedFields(val interface{}) {
	if err := j.validateSetAppendRepeatedFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appendRepeatedFields",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference)SetInternalValue(val *GoogleChronicleParserExtensionFieldExtractors) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference)SetLogFormat(val *string) {
	if err := j.validateSetLogFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) PutExtractors(value interface{}) {
	if err := g.validatePutExtractorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExtractors",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) PutPreprocessConfig(value *GoogleChronicleParserExtensionFieldExtractorsPreprocessConfig) {
	if err := g.validatePutPreprocessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPreprocessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) ResetAppendRepeatedFields() {
	_jsii_.InvokeVoid(
		g,
		"resetAppendRepeatedFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) ResetExtractors() {
	_jsii_.InvokeVoid(
		g,
		"resetExtractors",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) ResetLogFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetLogFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) ResetPreprocessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPreprocessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleParserExtensionFieldExtractorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

