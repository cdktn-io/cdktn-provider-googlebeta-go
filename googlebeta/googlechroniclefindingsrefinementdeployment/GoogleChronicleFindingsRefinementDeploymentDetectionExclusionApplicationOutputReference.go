// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclefindingsrefinementdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlechroniclefindingsrefinementdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference interface {
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
	CuratedRules() *[]*string
	SetCuratedRules(val *[]*string)
	CuratedRuleSets() *[]*string
	SetCuratedRuleSets(val *[]*string)
	CuratedRuleSetsInput() *[]*string
	CuratedRulesInput() *[]*string
	DeletedCuratedRuleSets() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplication
	SetInternalValue(val *GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplication)
	Rules() *[]*string
	SetRules(val *[]*string)
	RulesInput() *[]*string
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
	ResetCuratedRules()
	ResetCuratedRuleSets()
	ResetRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference
type jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CuratedRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"curatedRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CuratedRuleSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"curatedRuleSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CuratedRuleSetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"curatedRuleSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CuratedRulesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"curatedRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) DeletedCuratedRuleSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedCuratedRuleSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) InternalValue() *GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplication {
	var returns *GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) Rules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) RulesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFindingsRefinementDeployment.GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference_Override(g GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFindingsRefinementDeployment.GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetCuratedRules(val *[]*string) {
	if err := j.validateSetCuratedRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"curatedRules",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetCuratedRuleSets(val *[]*string) {
	if err := j.validateSetCuratedRuleSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"curatedRuleSets",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetInternalValue(val *GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetRules(val *[]*string) {
	if err := j.validateSetRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ResetCuratedRules() {
	_jsii_.InvokeVoid(
		g,
		"resetCuratedRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ResetCuratedRuleSets() {
	_jsii_.InvokeVoid(
		g,
		"resetCuratedRuleSets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ResetRules() {
	_jsii_.InvokeVoid(
		g,
		"resetRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

