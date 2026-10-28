// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlepubsubsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlepubsubsubscription/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference interface {
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
	DelimitedKey() GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference
	DelimitedKeyInput() *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePubsubSubscriptionBigtableConfigColumnFamilyMapping
	SetInternalValue(val *GooglePubsubSubscriptionBigtableConfigColumnFamilyMapping)
	RowKeySchema() GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchemaOutputReference
	RowKeySchemaInput() *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchema
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
	PutDelimitedKey(value *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey)
	PutRowKeySchema(value *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchema)
	ResetDelimitedKey()
	ResetRowKeySchema()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference
type jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) DelimitedKey() GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference {
	var returns GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference
	_jsii_.Get(
		j,
		"delimitedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) DelimitedKeyInput() *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey {
	var returns *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey
	_jsii_.Get(
		j,
		"delimitedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) InternalValue() *GooglePubsubSubscriptionBigtableConfigColumnFamilyMapping {
	var returns *GooglePubsubSubscriptionBigtableConfigColumnFamilyMapping
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) RowKeySchema() GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchemaOutputReference {
	var returns GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchemaOutputReference
	_jsii_.Get(
		j,
		"rowKeySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) RowKeySchemaInput() *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchema {
	var returns *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchema
	_jsii_.Get(
		j,
		"rowKeySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference_Override(g GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference)SetInternalValue(val *GooglePubsubSubscriptionBigtableConfigColumnFamilyMapping) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) PutDelimitedKey(value *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey) {
	if err := g.validatePutDelimitedKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDelimitedKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) PutRowKeySchema(value *GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingRowKeySchema) {
	if err := g.validatePutRowKeySchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRowKeySchema",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) ResetDelimitedKey() {
	_jsii_.InvokeVoid(
		g,
		"resetDelimitedKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) ResetRowKeySchema() {
	_jsii_.InvokeVoid(
		g,
		"resetRowKeySchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigtableConfigColumnFamilyMappingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

