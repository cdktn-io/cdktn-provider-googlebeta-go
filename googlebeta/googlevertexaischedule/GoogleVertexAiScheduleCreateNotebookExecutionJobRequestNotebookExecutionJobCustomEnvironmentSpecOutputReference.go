// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlevertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference interface {
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
	InternalValue() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec
	SetInternalValue(val *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec)
	MachineSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpecOutputReference
	MachineSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec
	NetworkSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpecOutputReference
	NetworkSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec
	PersistentDiskSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpecOutputReference
	PersistentDiskSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec
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
	PutMachineSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec)
	PutNetworkSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec)
	PutPersistentDiskSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec)
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

// The jsii proxy struct for GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference
type jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) InternalValue() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) MachineSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpecOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) MachineSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) NetworkSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpecOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpecOutputReference
	_jsii_.Get(
		j,
		"networkSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) NetworkSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec
	_jsii_.Get(
		j,
		"networkSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PersistentDiskSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpecOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpecOutputReference
	_jsii_.Get(
		j,
		"persistentDiskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PersistentDiskSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec
	_jsii_.Get(
		j,
		"persistentDiskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference_Override(g GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetInternalValue(val *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PutMachineSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec) {
	if err := g.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PutNetworkSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec) {
	if err := g.validatePutNetworkSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) PutPersistentDiskSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec) {
	if err := g.validatePutPersistentDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPersistentDiskSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ResetMachineSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ResetNetworkSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ResetPersistentDiskSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetPersistentDiskSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

