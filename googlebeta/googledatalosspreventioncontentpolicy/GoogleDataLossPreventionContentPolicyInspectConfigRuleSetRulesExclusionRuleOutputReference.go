// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googledatalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference interface {
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
	Dictionary() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference
	DictionaryInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary
	ExcludeByHotword() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference
	ExcludeByHotwordInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword
	ExcludeByImageFindings() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsOutputReference
	ExcludeByImageFindingsInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings
	ExcludeInfoTypes() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference
	ExcludeInfoTypesInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule
	SetInternalValue(val *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule)
	MatchingType() *string
	SetMatchingType(val *string)
	MatchingTypeInput() *string
	Regex() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegexOutputReference
	RegexInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex
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
	PutDictionary(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary)
	PutExcludeByHotword(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword)
	PutExcludeByImageFindings(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings)
	PutExcludeInfoTypes(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	PutRegex(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex)
	ResetDictionary()
	ResetExcludeByHotword()
	ResetExcludeByImageFindings()
	ResetExcludeInfoTypes()
	ResetRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference
type jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) Dictionary() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) DictionaryInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByHotword() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference
	_jsii_.Get(
		j,
		"excludeByHotword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByHotwordInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword
	_jsii_.Get(
		j,
		"excludeByHotwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByImageFindings() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindingsOutputReference
	_jsii_.Get(
		j,
		"excludeByImageFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByImageFindingsInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings
	_jsii_.Get(
		j,
		"excludeByImageFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeInfoTypes() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference
	_jsii_.Get(
		j,
		"excludeInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeInfoTypesInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	_jsii_.Get(
		j,
		"excludeInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) InternalValue() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) MatchingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) MatchingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) Regex() GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegexOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegexOutputReference
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) RegexInput() *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference_Override(g GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetInternalValue(val *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetMatchingType(val *string) {
	if err := j.validateSetMatchingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutDictionary(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary) {
	if err := g.validatePutDictionaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDictionary",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeByHotword(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByHotword) {
	if err := g.validatePutExcludeByHotwordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeByHotword",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeByImageFindings(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeByImageFindings) {
	if err := g.validatePutExcludeByImageFindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeByImageFindings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeInfoTypes(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) {
	if err := g.validatePutExcludeInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeInfoTypes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) PutRegex(value *GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleRegex) {
	if err := g.validatePutRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegex",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		g,
		"resetDictionary",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeByHotword() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeByHotword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeByImageFindings() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeByImageFindings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeInfoTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeInfoTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

