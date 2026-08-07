// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecolabschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlecolabschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference interface {
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
	InternalValue() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec
	SetInternalValue(val *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec)
	MachineSpec() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpecOutputReference
	MachineSpecInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec
	NetworkSpec() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpecOutputReference
	NetworkSpecInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec
	PersistentDiskSpec() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpecOutputReference
	PersistentDiskSpecInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec
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
	PutMachineSpec(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec)
	PutNetworkSpec(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec)
	PutPersistentDiskSpec(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec)
	ResetMachineSpec()
	ResetNetworkSpec()
	ResetPersistentDiskSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference
type jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) InternalValue() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec {
	var returns *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) MachineSpec() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpecOutputReference {
	var returns GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) MachineSpecInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec {
	var returns *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) NetworkSpec() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpecOutputReference {
	var returns GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpecOutputReference
	_jsii_.Get(
		j,
		"networkSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) NetworkSpecInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec {
	var returns *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec
	_jsii_.Get(
		j,
		"networkSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PersistentDiskSpec() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpecOutputReference {
	var returns GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpecOutputReference
	_jsii_.Get(
		j,
		"persistentDiskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PersistentDiskSpecInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec {
	var returns *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec
	_jsii_.Get(
		j,
		"persistentDiskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleColabSchedule.GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference_Override(g GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleColabSchedule.GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetInternalValue(val *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PutMachineSpec(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec) {
	if err := g.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PutNetworkSpec(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec) {
	if err := g.validatePutNetworkSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PutPersistentDiskSpec(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec) {
	if err := g.validatePutPersistentDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPersistentDiskSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ResetMachineSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ResetNetworkSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ResetPersistentDiskSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetPersistentDiskSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

