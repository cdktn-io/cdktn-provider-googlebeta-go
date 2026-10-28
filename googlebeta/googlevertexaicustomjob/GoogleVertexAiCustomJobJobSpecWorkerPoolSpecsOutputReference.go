// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexaicustomjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference interface {
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
	ContainerSpec() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpecOutputReference
	ContainerSpecInput() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskSpec() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpecOutputReference
	DiskSpecInput() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LustreMounts() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsList
	LustreMountsInput() interface{}
	MachineSpec() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpecOutputReference
	MachineSpecInput() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec
	NfsMounts() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsNfsMountsList
	NfsMountsInput() interface{}
	PythonPackageSpec() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference
	PythonPackageSpecInput() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec
	ReplicaCount() *float64
	SetReplicaCount(val *float64)
	ReplicaCountInput() *float64
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
	PutContainerSpec(value *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec)
	PutDiskSpec(value *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec)
	PutLustreMounts(value interface{})
	PutMachineSpec(value *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec)
	PutNfsMounts(value interface{})
	PutPythonPackageSpec(value *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec)
	ResetContainerSpec()
	ResetDiskSpec()
	ResetLustreMounts()
	ResetMachineSpec()
	ResetNfsMounts()
	ResetPythonPackageSpec()
	ResetReplicaCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference
type jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ContainerSpec() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpecOutputReference {
	var returns GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpecOutputReference
	_jsii_.Get(
		j,
		"containerSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ContainerSpecInput() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec {
	var returns *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec
	_jsii_.Get(
		j,
		"containerSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) DiskSpec() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpecOutputReference {
	var returns GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpecOutputReference
	_jsii_.Get(
		j,
		"diskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) DiskSpecInput() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec {
	var returns *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec
	_jsii_.Get(
		j,
		"diskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) LustreMounts() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsList {
	var returns GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsList
	_jsii_.Get(
		j,
		"lustreMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) LustreMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lustreMountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) MachineSpec() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpecOutputReference {
	var returns GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) MachineSpecInput() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec {
	var returns *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) NfsMounts() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsNfsMountsList {
	var returns GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsNfsMountsList
	_jsii_.Get(
		j,
		"nfsMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) NfsMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsMountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PythonPackageSpec() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference {
	var returns GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference
	_jsii_.Get(
		j,
		"pythonPackageSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PythonPackageSpecInput() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec {
	var returns *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec
	_jsii_.Get(
		j,
		"pythonPackageSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference_Override(g GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutContainerSpec(value *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec) {
	if err := g.validatePutContainerSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutDiskSpec(value *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec) {
	if err := g.validatePutDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiskSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutLustreMounts(value interface{}) {
	if err := g.validatePutLustreMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLustreMounts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutMachineSpec(value *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec) {
	if err := g.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutNfsMounts(value interface{}) {
	if err := g.validatePutNfsMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNfsMounts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutPythonPackageSpec(value *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec) {
	if err := g.validatePutPythonPackageSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPythonPackageSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetContainerSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetDiskSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetLustreMounts() {
	_jsii_.InvokeVoid(
		g,
		"resetLustreMounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetMachineSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetNfsMounts() {
	_jsii_.InvokeVoid(
		g,
		"resetNfsMounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetPythonPackageSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonPackageSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

