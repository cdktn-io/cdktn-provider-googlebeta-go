// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googledatalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference interface {
	cdktn.ComplexObject
	AdjustByImageFindings() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsOutputReference
	AdjustByImageFindingsInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings
	AdjustByMatchingInfoTypes() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypesOutputReference
	AdjustByMatchingInfoTypesInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes
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
	InternalValue() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule
	SetInternalValue(val *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule)
	LikelihoodAdjustment() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustmentOutputReference
	LikelihoodAdjustmentInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment
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
	PutAdjustByImageFindings(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings)
	PutAdjustByMatchingInfoTypes(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes)
	PutLikelihoodAdjustment(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment)
	ResetAdjustByImageFindings()
	ResetAdjustByMatchingInfoTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference
type jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) AdjustByImageFindings() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindingsOutputReference
	_jsii_.Get(
		j,
		"adjustByImageFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) AdjustByImageFindingsInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings
	_jsii_.Get(
		j,
		"adjustByImageFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) AdjustByMatchingInfoTypes() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypesOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypesOutputReference
	_jsii_.Get(
		j,
		"adjustByMatchingInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) AdjustByMatchingInfoTypesInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes
	_jsii_.Get(
		j,
		"adjustByMatchingInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) InternalValue() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) LikelihoodAdjustment() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustmentOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustmentOutputReference
	_jsii_.Get(
		j,
		"likelihoodAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) LikelihoodAdjustmentInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment
	_jsii_.Get(
		j,
		"likelihoodAdjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference_Override(g GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetInternalValue(val *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) PutAdjustByImageFindings(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByImageFindings) {
	if err := g.validatePutAdjustByImageFindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdjustByImageFindings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) PutAdjustByMatchingInfoTypes(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleAdjustByMatchingInfoTypes) {
	if err := g.validatePutAdjustByMatchingInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdjustByMatchingInfoTypes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) PutLikelihoodAdjustment(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleLikelihoodAdjustment) {
	if err := g.validatePutLikelihoodAdjustmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLikelihoodAdjustment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ResetAdjustByImageFindings() {
	_jsii_.InvokeVoid(
		g,
		"resetAdjustByImageFindings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ResetAdjustByMatchingInfoTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetAdjustByMatchingInfoTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesAdjustmentRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

