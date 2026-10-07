// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googledatalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference interface {
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
	InternalValue() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabel
	SetInternalValue(val *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabel)
	LabelFieldsToMatch() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelLabelFieldsToMatchList
	LabelFieldsToMatchInput() interface{}
	LabelId() *string
	SetLabelId(val *string)
	LabelIdInput() *string
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
	PutLabelFieldsToMatch(value interface{})
	ResetLabelFieldsToMatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference
type jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) InternalValue() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabel {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) LabelFieldsToMatch() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelLabelFieldsToMatchList {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelLabelFieldsToMatchList
	_jsii_.Get(
		j,
		"labelFieldsToMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) LabelFieldsToMatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelFieldsToMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) LabelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) LabelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference_Override(g GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference)SetInternalValue(val *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference)SetLabelId(val *string) {
	if err := j.validateSetLabelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelId",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) PutLabelFieldsToMatch(value interface{}) {
	if err := g.validatePutLabelFieldsToMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLabelFieldsToMatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) ResetLabelFieldsToMatch() {
	_jsii_.InvokeVoid(
		g,
		"resetLabelFieldsToMatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

