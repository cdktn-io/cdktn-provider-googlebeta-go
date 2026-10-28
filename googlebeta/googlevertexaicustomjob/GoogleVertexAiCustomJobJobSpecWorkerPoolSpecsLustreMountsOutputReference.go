// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexaicustomjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference interface {
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
	Filesystem() *string
	SetFilesystem(val *string)
	FilesystemInput() *string
	// Experimental.
	Fqn() *string
	InstanceIp() *string
	SetInstanceIp(val *string)
	InstanceIpInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountPoint() *string
	SetMountPoint(val *string)
	MountPointInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VolumeHandle() *string
	SetVolumeHandle(val *string)
	VolumeHandleInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference
type jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) Filesystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) FilesystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) InstanceIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) InstanceIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) MountPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) MountPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) VolumeHandle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeHandle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) VolumeHandleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeHandleInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference_Override(g GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference)SetFilesystem(val *string) {
	if err := j.validateSetFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filesystem",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference)SetInstanceIp(val *string) {
	if err := j.validateSetInstanceIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceIp",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference)SetMountPoint(val *string) {
	if err := j.validateSetMountPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountPoint",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference)SetVolumeHandle(val *string) {
	if err := j.validateSetVolumeHandleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeHandle",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

