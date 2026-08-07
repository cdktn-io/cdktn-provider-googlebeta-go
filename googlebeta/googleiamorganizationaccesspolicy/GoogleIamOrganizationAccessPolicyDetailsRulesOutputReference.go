// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamorganizationaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleiamorganizationaccesspolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference interface {
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
	Conditions() GoogleIamOrganizationAccessPolicyDetailsRulesConditionsList
	ConditionsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Effect() *string
	SetEffect(val *string)
	EffectInput() *string
	ExcludedPrincipals() *[]*string
	SetExcludedPrincipals(val *[]*string)
	ExcludedPrincipalsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Operation() GoogleIamOrganizationAccessPolicyDetailsRulesOperationOutputReference
	OperationInput() *GoogleIamOrganizationAccessPolicyDetailsRulesOperation
	Principals() *[]*string
	SetPrincipals(val *[]*string)
	PrincipalsInput() *[]*string
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
	PutConditions(value interface{})
	PutOperation(value *GoogleIamOrganizationAccessPolicyDetailsRulesOperation)
	ResetConditions()
	ResetDescription()
	ResetExcludedPrincipals()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference
type jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) Conditions() GoogleIamOrganizationAccessPolicyDetailsRulesConditionsList {
	var returns GoogleIamOrganizationAccessPolicyDetailsRulesConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) Effect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) EffectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ExcludedPrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ExcludedPrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) Operation() GoogleIamOrganizationAccessPolicyDetailsRulesOperationOutputReference {
	var returns GoogleIamOrganizationAccessPolicyDetailsRulesOperationOutputReference
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) OperationInput() *GoogleIamOrganizationAccessPolicyDetailsRulesOperation {
	var returns *GoogleIamOrganizationAccessPolicyDetailsRulesOperation
	_jsii_.Get(
		j,
		"operationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) Principals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"principals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) PrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"principalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIamOrganizationAccessPolicyDetailsRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIamOrganizationAccessPolicyDetailsRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamOrganizationAccessPolicy.GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleIamOrganizationAccessPolicyDetailsRulesOutputReference_Override(g GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamOrganizationAccessPolicy.GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference)SetEffect(val *string) {
	if err := j.validateSetEffectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effect",
		val,
	)
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference)SetExcludedPrincipals(val *[]*string) {
	if err := j.validateSetExcludedPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedPrincipals",
		val,
	)
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference)SetPrincipals(val *[]*string) {
	if err := j.validateSetPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principals",
		val,
	)
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) PutConditions(value interface{}) {
	if err := g.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConditions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) PutOperation(value *GoogleIamOrganizationAccessPolicyDetailsRulesOperation) {
	if err := g.validatePutOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOperation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ResetConditions() {
	_jsii_.InvokeVoid(
		g,
		"resetConditions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ResetExcludedPrincipals() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedPrincipals",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIamOrganizationAccessPolicyDetailsRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

