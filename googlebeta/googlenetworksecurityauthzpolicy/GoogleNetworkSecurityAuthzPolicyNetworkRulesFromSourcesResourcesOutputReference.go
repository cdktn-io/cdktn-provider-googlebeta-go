// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityauthzpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlenetworksecurityauthzpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference interface {
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
	IamServiceAccount() GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesIamServiceAccountOutputReference
	IamServiceAccountInput() *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesIamServiceAccount
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TagValueIdSet() GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesTagValueIdSetOutputReference
	TagValueIdSetInput() *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesTagValueIdSet
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
	PutIamServiceAccount(value *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesIamServiceAccount)
	PutTagValueIdSet(value *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesTagValueIdSet)
	ResetIamServiceAccount()
	ResetTagValueIdSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference
type jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) IamServiceAccount() GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesIamServiceAccountOutputReference {
	var returns GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesIamServiceAccountOutputReference
	_jsii_.Get(
		j,
		"iamServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) IamServiceAccountInput() *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesIamServiceAccount {
	var returns *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesIamServiceAccount
	_jsii_.Get(
		j,
		"iamServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) TagValueIdSet() GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesTagValueIdSetOutputReference {
	var returns GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesTagValueIdSetOutputReference
	_jsii_.Get(
		j,
		"tagValueIdSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) TagValueIdSetInput() *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesTagValueIdSet {
	var returns *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesTagValueIdSet
	_jsii_.Get(
		j,
		"tagValueIdSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecurityAuthzPolicy.GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference_Override(g GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecurityAuthzPolicy.GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) PutIamServiceAccount(value *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesIamServiceAccount) {
	if err := g.validatePutIamServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIamServiceAccount",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) PutTagValueIdSet(value *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesTagValueIdSet) {
	if err := g.validatePutTagValueIdSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTagValueIdSet",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) ResetIamServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetIamServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) ResetTagValueIdSet() {
	_jsii_.InvokeVoid(
		g,
		"resetTagValueIdSet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

