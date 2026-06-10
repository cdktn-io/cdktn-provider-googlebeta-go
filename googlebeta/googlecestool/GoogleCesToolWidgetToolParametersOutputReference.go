// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlecestool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesToolWidgetToolParametersOutputReference interface {
	cdktn.ComplexObject
	AdditionalProperties() *string
	SetAdditionalProperties(val *string)
	AdditionalPropertiesInput() *string
	AnyOf() *string
	SetAnyOf(val *string)
	AnyOfInput() *string
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
	Default() *string
	SetDefault(val *string)
	DefaultInput() *string
	Defs() *string
	SetDefs(val *string)
	DefsInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enum() *[]*string
	SetEnum(val *[]*string)
	EnumInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesToolWidgetToolParameters
	SetInternalValue(val *GoogleCesToolWidgetToolParameters)
	Items() *string
	SetItems(val *string)
	ItemsInput() *string
	Maximum() *float64
	SetMaximum(val *float64)
	MaximumInput() *float64
	MaxItems() *float64
	SetMaxItems(val *float64)
	MaxItemsInput() *float64
	Minimum() *float64
	SetMinimum(val *float64)
	MinimumInput() *float64
	MinItems() *float64
	SetMinItems(val *float64)
	MinItemsInput() *float64
	Nullable() interface{}
	SetNullable(val interface{})
	NullableInput() interface{}
	PrefixItems() *string
	SetPrefixItems(val *string)
	PrefixItemsInput() *string
	Properties() *string
	SetProperties(val *string)
	PropertiesInput() *string
	Ref() *string
	SetRef(val *string)
	RefInput() *string
	Required() *[]*string
	SetRequired(val *[]*string)
	RequiredInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UniqueItems() interface{}
	SetUniqueItems(val interface{})
	UniqueItemsInput() interface{}
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
	ResetAdditionalProperties()
	ResetAnyOf()
	ResetDefault()
	ResetDefs()
	ResetDescription()
	ResetEnum()
	ResetItems()
	ResetMaximum()
	ResetMaxItems()
	ResetMinimum()
	ResetMinItems()
	ResetNullable()
	ResetPrefixItems()
	ResetProperties()
	ResetRef()
	ResetRequired()
	ResetTitle()
	ResetUniqueItems()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesToolWidgetToolParametersOutputReference
type jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) AdditionalProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) AdditionalPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) AnyOf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) AnyOfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Default() *string {
	var returns *string
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) DefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Defs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) DefsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Enum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) EnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) InternalValue() *GoogleCesToolWidgetToolParameters {
	var returns *GoogleCesToolWidgetToolParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Items() *string {
	var returns *string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ItemsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Maximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) MaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) MaxItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Minimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) MinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) MinItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) MinItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Nullable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) NullableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) PrefixItems() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) PrefixItemsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Properties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) PropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) RefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Required() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) RequiredInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) UniqueItems() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniqueItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) UniqueItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniqueItemsInput",
		&returns,
	)
	return returns
}


func NewGoogleCesToolWidgetToolParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesToolWidgetToolParametersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesToolWidgetToolParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolWidgetToolParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesToolWidgetToolParametersOutputReference_Override(g GoogleCesToolWidgetToolParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolWidgetToolParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetAdditionalProperties(val *string) {
	if err := j.validateSetAdditionalPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalProperties",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetAnyOf(val *string) {
	if err := j.validateSetAnyOfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyOf",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetDefault(val *string) {
	if err := j.validateSetDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetDefs(val *string) {
	if err := j.validateSetDefsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defs",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetEnum(val *[]*string) {
	if err := j.validateSetEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enum",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetInternalValue(val *GoogleCesToolWidgetToolParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetItems(val *string) {
	if err := j.validateSetItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetMaximum(val *float64) {
	if err := j.validateSetMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximum",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetMaxItems(val *float64) {
	if err := j.validateSetMaxItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxItems",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetMinimum(val *float64) {
	if err := j.validateSetMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimum",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetMinItems(val *float64) {
	if err := j.validateSetMinItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minItems",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetNullable(val interface{}) {
	if err := j.validateSetNullableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullable",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetPrefixItems(val *string) {
	if err := j.validateSetPrefixItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixItems",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetProperties(val *string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetRef(val *string) {
	if err := j.validateSetRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ref",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetRequired(val *[]*string) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference)SetUniqueItems(val interface{}) {
	if err := j.validateSetUniqueItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniqueItems",
		val,
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetAdditionalProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetAnyOf() {
	_jsii_.InvokeVoid(
		g,
		"resetAnyOf",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetDefault() {
	_jsii_.InvokeVoid(
		g,
		"resetDefault",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetDefs() {
	_jsii_.InvokeVoid(
		g,
		"resetDefs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetEnum() {
	_jsii_.InvokeVoid(
		g,
		"resetEnum",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		g,
		"resetItems",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetMaximum() {
	_jsii_.InvokeVoid(
		g,
		"resetMaximum",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetMaxItems() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxItems",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetMinimum() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimum",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetMinItems() {
	_jsii_.InvokeVoid(
		g,
		"resetMinItems",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetNullable() {
	_jsii_.InvokeVoid(
		g,
		"resetNullable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetPrefixItems() {
	_jsii_.InvokeVoid(
		g,
		"resetPrefixItems",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetRef() {
	_jsii_.InvokeVoid(
		g,
		"resetRef",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		g,
		"resetRequired",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		g,
		"resetTitle",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ResetUniqueItems() {
	_jsii_.InvokeVoid(
		g,
		"resetUniqueItems",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesToolWidgetToolParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

