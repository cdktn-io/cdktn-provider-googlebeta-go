// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleconfigdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleconfigdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleConfigDeploymentTerraformBlueprintOutputReference interface {
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
	GcsSource() *string
	SetGcsSource(val *string)
	GcsSourceInput() *string
	GitSource() GoogleConfigDeploymentTerraformBlueprintGitSourceOutputReference
	GitSourceInput() *GoogleConfigDeploymentTerraformBlueprintGitSource
	InputValues() GoogleConfigDeploymentTerraformBlueprintInputValuesList
	InputValuesInput() interface{}
	InternalValue() *GoogleConfigDeploymentTerraformBlueprint
	SetInternalValue(val *GoogleConfigDeploymentTerraformBlueprint)
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
	PutGitSource(value *GoogleConfigDeploymentTerraformBlueprintGitSource)
	PutInputValues(value interface{})
	ResetGcsSource()
	ResetGitSource()
	ResetInputValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleConfigDeploymentTerraformBlueprintOutputReference
type jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GcsSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GcsSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GitSource() GoogleConfigDeploymentTerraformBlueprintGitSourceOutputReference {
	var returns GoogleConfigDeploymentTerraformBlueprintGitSourceOutputReference
	_jsii_.Get(
		j,
		"gitSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GitSourceInput() *GoogleConfigDeploymentTerraformBlueprintGitSource {
	var returns *GoogleConfigDeploymentTerraformBlueprintGitSource
	_jsii_.Get(
		j,
		"gitSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) InputValues() GoogleConfigDeploymentTerraformBlueprintInputValuesList {
	var returns GoogleConfigDeploymentTerraformBlueprintInputValuesList
	_jsii_.Get(
		j,
		"inputValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) InputValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) InternalValue() *GoogleConfigDeploymentTerraformBlueprint {
	var returns *GoogleConfigDeploymentTerraformBlueprint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleConfigDeploymentTerraformBlueprintOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleConfigDeploymentTerraformBlueprintOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleConfigDeploymentTerraformBlueprintOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleConfigDeployment.GoogleConfigDeploymentTerraformBlueprintOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleConfigDeploymentTerraformBlueprintOutputReference_Override(g GoogleConfigDeploymentTerraformBlueprintOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleConfigDeployment.GoogleConfigDeploymentTerraformBlueprintOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference)SetGcsSource(val *string) {
	if err := j.validateSetGcsSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsSource",
		val,
	)
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference)SetInternalValue(val *GoogleConfigDeploymentTerraformBlueprint) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) PutGitSource(value *GoogleConfigDeploymentTerraformBlueprintGitSource) {
	if err := g.validatePutGitSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGitSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) PutInputValues(value interface{}) {
	if err := g.validatePutInputValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInputValues",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) ResetGcsSource() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) ResetGitSource() {
	_jsii_.InvokeVoid(
		g,
		"resetGitSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) ResetInputValues() {
	_jsii_.InvokeVoid(
		g,
		"resetInputValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleConfigDeploymentTerraformBlueprintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

