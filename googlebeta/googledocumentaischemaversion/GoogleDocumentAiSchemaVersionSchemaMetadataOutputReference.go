// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledocumentaischemaversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googledocumentaischemaversion/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference interface {
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
	DocumentAllowMultipleLabels() interface{}
	SetDocumentAllowMultipleLabels(val interface{})
	DocumentAllowMultipleLabelsInput() interface{}
	DocumentSplitter() interface{}
	SetDocumentSplitter(val interface{})
	DocumentSplitterInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDocumentAiSchemaVersionSchemaMetadata
	SetInternalValue(val *GoogleDocumentAiSchemaVersionSchemaMetadata)
	PrefixedNamingOnProperties() interface{}
	SetPrefixedNamingOnProperties(val interface{})
	PrefixedNamingOnPropertiesInput() interface{}
	SkipNamingValidation() interface{}
	SetSkipNamingValidation(val interface{})
	SkipNamingValidationInput() interface{}
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
	ResetDocumentAllowMultipleLabels()
	ResetDocumentSplitter()
	ResetPrefixedNamingOnProperties()
	ResetSkipNamingValidation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference
type jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) DocumentAllowMultipleLabels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentAllowMultipleLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) DocumentAllowMultipleLabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentAllowMultipleLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) DocumentSplitter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentSplitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) DocumentSplitterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentSplitterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) InternalValue() *GoogleDocumentAiSchemaVersionSchemaMetadata {
	var returns *GoogleDocumentAiSchemaVersionSchemaMetadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) PrefixedNamingOnProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prefixedNamingOnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) PrefixedNamingOnPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prefixedNamingOnPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) SkipNamingValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNamingValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) SkipNamingValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNamingValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDocumentAiSchemaVersionSchemaMetadataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDocumentAiSchemaVersionSchemaMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDocumentAiSchemaVersion.GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDocumentAiSchemaVersionSchemaMetadataOutputReference_Override(g GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDocumentAiSchemaVersion.GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference)SetDocumentAllowMultipleLabels(val interface{}) {
	if err := j.validateSetDocumentAllowMultipleLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentAllowMultipleLabels",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference)SetDocumentSplitter(val interface{}) {
	if err := j.validateSetDocumentSplitterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentSplitter",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference)SetInternalValue(val *GoogleDocumentAiSchemaVersionSchemaMetadata) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference)SetPrefixedNamingOnProperties(val interface{}) {
	if err := j.validateSetPrefixedNamingOnPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixedNamingOnProperties",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference)SetSkipNamingValidation(val interface{}) {
	if err := j.validateSetSkipNamingValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipNamingValidation",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) ResetDocumentAllowMultipleLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetDocumentAllowMultipleLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) ResetDocumentSplitter() {
	_jsii_.InvokeVoid(
		g,
		"resetDocumentSplitter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) ResetPrefixedNamingOnProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetPrefixedNamingOnProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) ResetSkipNamingValidation() {
	_jsii_.InvokeVoid(
		g,
		"resetSkipNamingValidation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDocumentAiSchemaVersionSchemaMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

