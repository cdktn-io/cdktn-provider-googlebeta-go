// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googledatalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference interface {
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
	InternalValue() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity
	SetInternalValue(val *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WindowAfter() *float64
	SetWindowAfter(val *float64)
	WindowAfterInput() *float64
	WindowBefore() *float64
	SetWindowBefore(val *float64)
	WindowBeforeInput() *float64
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
	ResetWindowAfter()
	ResetWindowBefore()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference
type jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) InternalValue() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) WindowAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) WindowAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) WindowBefore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) WindowBeforeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowBeforeInput",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference_Override(g GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference)SetInternalValue(val *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference)SetWindowAfter(val *float64) {
	if err := j.validateSetWindowAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowAfter",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference)SetWindowBefore(val *float64) {
	if err := j.validateSetWindowBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowBefore",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) ResetWindowAfter() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowAfter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) ResetWindowBefore() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowBefore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

