// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityfirewallendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlenetworksecurityfirewallendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetworkSecurityFirewallEndpointWildfireSettings
	SetInternalValue(val *GoogleNetworkSecurityFirewallEndpointWildfireSettings)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WildfireInlineCloudAnalysisSettings() GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference
	WildfireInlineCloudAnalysisSettingsInput() *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings
	WildfireRealtimeLookupDuration() *string
	SetWildfireRealtimeLookupDuration(val *string)
	WildfireRealtimeLookupDurationInput() *string
	WildfireRealtimeLookupTimeoutAction() *string
	SetWildfireRealtimeLookupTimeoutAction(val *string)
	WildfireRealtimeLookupTimeoutActionInput() *string
	WildfireRegion() *string
	SetWildfireRegion(val *string)
	WildfireRegionInput() *string
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
	PutWildfireInlineCloudAnalysisSettings(value *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings)
	ResetEnabled()
	ResetWildfireInlineCloudAnalysisSettings()
	ResetWildfireRealtimeLookupDuration()
	ResetWildfireRealtimeLookupTimeoutAction()
	ResetWildfireRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference
type jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) InternalValue() *GoogleNetworkSecurityFirewallEndpointWildfireSettings {
	var returns *GoogleNetworkSecurityFirewallEndpointWildfireSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireInlineCloudAnalysisSettings() GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference {
	var returns GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference
	_jsii_.Get(
		j,
		"wildfireInlineCloudAnalysisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireInlineCloudAnalysisSettingsInput() *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings {
	var returns *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings
	_jsii_.Get(
		j,
		"wildfireInlineCloudAnalysisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRealtimeLookupDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRealtimeLookupDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRealtimeLookupTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRealtimeLookupTimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupTimeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRegionInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecurityFirewallEndpoint.GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference_Override(g GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecurityFirewallEndpoint.GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetInternalValue(val *GoogleNetworkSecurityFirewallEndpointWildfireSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetWildfireRealtimeLookupDuration(val *string) {
	if err := j.validateSetWildfireRealtimeLookupDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildfireRealtimeLookupDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetWildfireRealtimeLookupTimeoutAction(val *string) {
	if err := j.validateSetWildfireRealtimeLookupTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildfireRealtimeLookupTimeoutAction",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetWildfireRegion(val *string) {
	if err := j.validateSetWildfireRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildfireRegion",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) PutWildfireInlineCloudAnalysisSettings(value *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings) {
	if err := g.validatePutWildfireInlineCloudAnalysisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWildfireInlineCloudAnalysisSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetWildfireInlineCloudAnalysisSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireInlineCloudAnalysisSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetWildfireRealtimeLookupDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireRealtimeLookupDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetWildfireRealtimeLookupTimeoutAction() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireRealtimeLookupTimeoutAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetWildfireRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetWildfireRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

