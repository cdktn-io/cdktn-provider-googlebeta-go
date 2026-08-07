// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlebiglakehivetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeserializerClass() *string
	SetDeserializerClass(val *string)
	DeserializerClassInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo
	SetInternalValue(val *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	SerdeType() *string
	SetSerdeType(val *string)
	SerdeTypeInput() *string
	SerializationLib() *string
	SetSerializationLib(val *string)
	SerializationLibInput() *string
	SerializerClass() *string
	SetSerializerClass(val *string)
	SerializerClassInput() *string
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
	ResetDescription()
	ResetDeserializerClass()
	ResetParameters()
	ResetSerdeType()
	ResetSerializerClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference
type jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) DeserializerClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deserializerClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) DeserializerClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deserializerClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) InternalValue() *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo {
	var returns *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerdeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serdeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerdeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serdeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerializationLib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerializationLibInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerializerClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializerClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerializerClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializerClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBiglakeHiveTable.GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference_Override(g GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBiglakeHiveTable.GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetDeserializerClass(val *string) {
	if err := j.validateSetDeserializerClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deserializerClass",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetInternalValue(val *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetSerdeType(val *string) {
	if err := j.validateSetSerdeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serdeType",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetSerializationLib(val *string) {
	if err := j.validateSetSerializationLibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serializationLib",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetSerializerClass(val *string) {
	if err := j.validateSetSerializerClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serializerClass",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetDeserializerClass() {
	_jsii_.InvokeVoid(
		g,
		"resetDeserializerClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetSerdeType() {
	_jsii_.InvokeVoid(
		g,
		"resetSerdeType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetSerializerClass() {
	_jsii_.InvokeVoid(
		g,
		"resetSerializerClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

