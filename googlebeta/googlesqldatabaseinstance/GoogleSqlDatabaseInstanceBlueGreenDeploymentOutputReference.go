// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlesqldatabaseinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference interface {
	cdktn.ComplexObject
	ActiveEnvironment() *string
	SetActiveEnvironment(val *string)
	ActiveEnvironmentInput() *string
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
	DeleteOldSourceOnDestroy() interface{}
	SetDeleteOldSourceOnDestroy(val interface{})
	DeleteOldSourceOnDestroyInput() interface{}
	DeploymentName() *string
	SetDeploymentName(val *string)
	DeploymentNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSqlDatabaseInstanceBlueGreenDeployment
	SetInternalValue(val *GoogleSqlDatabaseInstanceBlueGreenDeployment)
	OldSourceInstanceConnectionName() *string
	OldSourceInstanceName() *string
	TargetDatabaseVersion() *string
	SetTargetDatabaseVersion(val *string)
	TargetDatabaseVersionInput() *string
	TargetInstanceConnectionName() *string
	TargetInstanceIpAddresses() GoogleSqlDatabaseInstanceBlueGreenDeploymentTargetInstanceIpAddressesList
	TargetInstanceName() *string
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
	ResetActiveEnvironment()
	ResetDeleteOldSourceOnDestroy()
	ResetTargetDatabaseVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference
type jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) ActiveEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) ActiveEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) DeleteOldSourceOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOldSourceOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) DeleteOldSourceOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOldSourceOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) DeploymentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) DeploymentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) InternalValue() *GoogleSqlDatabaseInstanceBlueGreenDeployment {
	var returns *GoogleSqlDatabaseInstanceBlueGreenDeployment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) OldSourceInstanceConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldSourceInstanceConnectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) OldSourceInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldSourceInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetDatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDatabaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetDatabaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDatabaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetInstanceConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInstanceConnectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetInstanceIpAddresses() GoogleSqlDatabaseInstanceBlueGreenDeploymentTargetInstanceIpAddressesList {
	var returns GoogleSqlDatabaseInstanceBlueGreenDeploymentTargetInstanceIpAddressesList
	_jsii_.Get(
		j,
		"targetInstanceIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) TargetInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference_Override(g GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetActiveEnvironment(val *string) {
	if err := j.validateSetActiveEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeEnvironment",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetDeleteOldSourceOnDestroy(val interface{}) {
	if err := j.validateSetDeleteOldSourceOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOldSourceOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetDeploymentName(val *string) {
	if err := j.validateSetDeploymentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentName",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetInternalValue(val *GoogleSqlDatabaseInstanceBlueGreenDeployment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetTargetDatabaseVersion(val *string) {
	if err := j.validateSetTargetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDatabaseVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) ResetActiveEnvironment() {
	_jsii_.InvokeVoid(
		g,
		"resetActiveEnvironment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) ResetDeleteOldSourceOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteOldSourceOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) ResetTargetDatabaseVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetDatabaseVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceBlueGreenDeploymentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

