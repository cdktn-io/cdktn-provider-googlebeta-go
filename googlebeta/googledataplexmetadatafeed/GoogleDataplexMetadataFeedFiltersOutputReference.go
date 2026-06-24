// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataplexmetadatafeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googledataplexmetadatafeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataplexMetadataFeedFiltersOutputReference interface {
	cdktn.ComplexObject
	AspectTypes() *[]*string
	SetAspectTypes(val *[]*string)
	AspectTypesInput() *[]*string
	ChangeTypes() *[]*string
	SetChangeTypes(val *[]*string)
	ChangeTypesInput() *[]*string
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
	EntryTypes() *[]*string
	SetEntryTypes(val *[]*string)
	EntryTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataplexMetadataFeedFilters
	SetInternalValue(val *GoogleDataplexMetadataFeedFilters)
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
	ResetAspectTypes()
	ResetChangeTypes()
	ResetEntryTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexMetadataFeedFiltersOutputReference
type jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) AspectTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aspectTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) AspectTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aspectTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) ChangeTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"changeTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) ChangeTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"changeTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) EntryTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) EntryTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) InternalValue() *GoogleDataplexMetadataFeedFilters {
	var returns *GoogleDataplexMetadataFeedFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataplexMetadataFeedFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataplexMetadataFeedFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexMetadataFeedFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataplexMetadataFeed.GoogleDataplexMetadataFeedFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexMetadataFeedFiltersOutputReference_Override(g GoogleDataplexMetadataFeedFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataplexMetadataFeed.GoogleDataplexMetadataFeedFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference)SetAspectTypes(val *[]*string) {
	if err := j.validateSetAspectTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aspectTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference)SetChangeTypes(val *[]*string) {
	if err := j.validateSetChangeTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference)SetEntryTypes(val *[]*string) {
	if err := j.validateSetEntryTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference)SetInternalValue(val *GoogleDataplexMetadataFeedFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) ResetAspectTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetAspectTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) ResetChangeTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetChangeTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) ResetEntryTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetEntryTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

