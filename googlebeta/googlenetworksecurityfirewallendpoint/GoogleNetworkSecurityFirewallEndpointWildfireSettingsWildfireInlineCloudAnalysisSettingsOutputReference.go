// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityfirewallendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlenetworksecurityfirewallendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference interface {
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
	InternalValue() *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings
	SetInternalValue(val *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings)
	MaxAnalysisDuration() *string
	SetMaxAnalysisDuration(val *string)
	MaxAnalysisDurationInput() *string
	SubmissionTimeoutLoggingDisabled() interface{}
	SetSubmissionTimeoutLoggingDisabled(val interface{})
	SubmissionTimeoutLoggingDisabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutAction() *string
	SetTimeoutAction(val *string)
	TimeoutActionInput() *string
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
	ResetMaxAnalysisDuration()
	ResetSubmissionTimeoutLoggingDisabled()
	ResetTimeoutAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference
type jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) InternalValue() *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings {
	var returns *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) MaxAnalysisDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxAnalysisDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) MaxAnalysisDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxAnalysisDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) SubmissionTimeoutLoggingDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"submissionTimeoutLoggingDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) SubmissionTimeoutLoggingDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"submissionTimeoutLoggingDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) TimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) TimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutActionInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecurityFirewallEndpoint.GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference_Override(g GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecurityFirewallEndpoint.GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference)SetInternalValue(val *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference)SetMaxAnalysisDuration(val *string) {
	if err := j.validateSetMaxAnalysisDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAnalysisDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference)SetSubmissionTimeoutLoggingDisabled(val interface{}) {
	if err := j.validateSetSubmissionTimeoutLoggingDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"submissionTimeoutLoggingDisabled",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference)SetTimeoutAction(val *string) {
	if err := j.validateSetTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutAction",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) ResetMaxAnalysisDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxAnalysisDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) ResetSubmissionTimeoutLoggingDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetSubmissionTimeoutLoggingDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) ResetTimeoutAction() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

