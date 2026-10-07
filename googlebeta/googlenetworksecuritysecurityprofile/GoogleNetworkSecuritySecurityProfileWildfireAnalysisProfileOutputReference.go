// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecuritysecurityprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlenetworksecuritysecurityprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference interface {
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
	InternalValue() *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfile
	SetInternalValue(val *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfile)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WildfireInlineCloudAnalysisRules() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList
	WildfireInlineCloudAnalysisRulesInput() interface{}
	WildfireInlineMlOverrides() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlOverridesList
	WildfireInlineMlOverridesInput() interface{}
	WildfireInlineMlSetting() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSettingOutputReference
	WildfireInlineMlSettingInput() *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting
	WildfireOverrides() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireOverridesList
	WildfireOverridesInput() interface{}
	WildfireRealtimeLookup() interface{}
	SetWildfireRealtimeLookup(val interface{})
	WildfireRealtimeLookupInput() interface{}
	WildfireSubmissionRules() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireSubmissionRulesList
	WildfireSubmissionRulesInput() interface{}
	WildfireThreatOverrides() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireThreatOverridesList
	WildfireThreatOverridesInput() interface{}
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
	PutWildfireInlineCloudAnalysisRules(value interface{})
	PutWildfireInlineMlOverrides(value interface{})
	PutWildfireInlineMlSetting(value *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting)
	PutWildfireOverrides(value interface{})
	PutWildfireSubmissionRules(value interface{})
	PutWildfireThreatOverrides(value interface{})
	ResetWildfireInlineCloudAnalysisRules()
	ResetWildfireInlineMlOverrides()
	ResetWildfireInlineMlSetting()
	ResetWildfireOverrides()
	ResetWildfireRealtimeLookup()
	ResetWildfireSubmissionRules()
	ResetWildfireThreatOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference
type jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) InternalValue() *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfile {
	var returns *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineCloudAnalysisRules() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList {
	var returns GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList
	_jsii_.Get(
		j,
		"wildfireInlineCloudAnalysisRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineCloudAnalysisRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireInlineCloudAnalysisRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineMlOverrides() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlOverridesList {
	var returns GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlOverridesList
	_jsii_.Get(
		j,
		"wildfireInlineMlOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineMlOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireInlineMlOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineMlSetting() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSettingOutputReference {
	var returns GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSettingOutputReference
	_jsii_.Get(
		j,
		"wildfireInlineMlSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineMlSettingInput() *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting {
	var returns *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting
	_jsii_.Get(
		j,
		"wildfireInlineMlSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireOverrides() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireOverridesList {
	var returns GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireOverridesList
	_jsii_.Get(
		j,
		"wildfireOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireRealtimeLookup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireRealtimeLookup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireRealtimeLookupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireSubmissionRules() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireSubmissionRulesList {
	var returns GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireSubmissionRulesList
	_jsii_.Get(
		j,
		"wildfireSubmissionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireSubmissionRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireSubmissionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireThreatOverrides() GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireThreatOverridesList {
	var returns GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireThreatOverridesList
	_jsii_.Get(
		j,
		"wildfireThreatOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireThreatOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireThreatOverridesInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecuritySecurityProfile.GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference_Override(g GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecuritySecurityProfile.GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetInternalValue(val *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetWildfireRealtimeLookup(val interface{}) {
	if err := j.validateSetWildfireRealtimeLookupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildfireRealtimeLookup",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireInlineCloudAnalysisRules(value interface{}) {
	if err := g.validatePutWildfireInlineCloudAnalysisRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWildfireInlineCloudAnalysisRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireInlineMlOverrides(value interface{}) {
	if err := g.validatePutWildfireInlineMlOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWildfireInlineMlOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireInlineMlSetting(value *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting) {
	if err := g.validatePutWildfireInlineMlSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWildfireInlineMlSetting",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireOverrides(value interface{}) {
	if err := g.validatePutWildfireOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWildfireOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireSubmissionRules(value interface{}) {
	if err := g.validatePutWildfireSubmissionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWildfireSubmissionRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireThreatOverrides(value interface{}) {
	if err := g.validatePutWildfireThreatOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWildfireThreatOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireInlineCloudAnalysisRules() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireInlineCloudAnalysisRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireInlineMlOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireInlineMlOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireInlineMlSetting() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireInlineMlSetting",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireRealtimeLookup() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireRealtimeLookup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireSubmissionRules() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireSubmissionRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireThreatOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireThreatOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

