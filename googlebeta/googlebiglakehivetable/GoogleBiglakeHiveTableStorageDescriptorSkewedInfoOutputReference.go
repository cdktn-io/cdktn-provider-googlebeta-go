// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlebiglakehivetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo
	SetInternalValue(val *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo)
	SkewedColNames() *[]*string
	SetSkewedColNames(val *[]*string)
	SkewedColNamesInput() *[]*string
	SkewedColValues() GoogleBiglakeHiveTableStorageDescriptorSkewedInfoSkewedColValuesList
	SkewedColValuesInput() interface{}
	SkewedKeyValuesLocations() GoogleBiglakeHiveTableStorageDescriptorSkewedInfoSkewedKeyValuesLocationsList
	SkewedKeyValuesLocationsInput() interface{}
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
	PutSkewedColValues(value interface{})
	PutSkewedKeyValuesLocations(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference
type jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) InternalValue() *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo {
	var returns *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) SkewedColNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) SkewedColNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) SkewedColValues() GoogleBiglakeHiveTableStorageDescriptorSkewedInfoSkewedColValuesList {
	var returns GoogleBiglakeHiveTableStorageDescriptorSkewedInfoSkewedColValuesList
	_jsii_.Get(
		j,
		"skewedColValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) SkewedColValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skewedColValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) SkewedKeyValuesLocations() GoogleBiglakeHiveTableStorageDescriptorSkewedInfoSkewedKeyValuesLocationsList {
	var returns GoogleBiglakeHiveTableStorageDescriptorSkewedInfoSkewedKeyValuesLocationsList
	_jsii_.Get(
		j,
		"skewedKeyValuesLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) SkewedKeyValuesLocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skewedKeyValuesLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBiglakeHiveTable.GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference_Override(g GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBiglakeHiveTable.GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference)SetInternalValue(val *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference)SetSkewedColNames(val *[]*string) {
	if err := j.validateSetSkewedColNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColNames",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) PutSkewedColValues(value interface{}) {
	if err := g.validatePutSkewedColValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSkewedColValues",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) PutSkewedKeyValuesLocations(value interface{}) {
	if err := g.validatePutSkewedKeyValuesLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSkewedKeyValuesLocations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

