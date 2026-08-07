// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlevertexaipersistentresource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiPersistentResourceResourcePoolsOutputReference interface {
	cdktn.ComplexObject
	AutoscalingSpec() GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference
	AutoscalingSpecInput() *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec
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
	DiskSpec() GoogleVertexAiPersistentResourceResourcePoolsDiskSpecOutputReference
	DiskSpecInput() *GoogleVertexAiPersistentResourceResourcePoolsDiskSpec
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MachineSpec() GoogleVertexAiPersistentResourceResourcePoolsMachineSpecOutputReference
	MachineSpecInput() *GoogleVertexAiPersistentResourceResourcePoolsMachineSpec
	ReplicaCount() *string
	SetReplicaCount(val *string)
	ReplicaCountInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UsedReplicaCount() *string
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
	PutAutoscalingSpec(value *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec)
	PutDiskSpec(value *GoogleVertexAiPersistentResourceResourcePoolsDiskSpec)
	PutMachineSpec(value *GoogleVertexAiPersistentResourceResourcePoolsMachineSpec)
	ResetAutoscalingSpec()
	ResetDiskSpec()
	ResetId()
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

// The jsii proxy struct for GoogleVertexAiPersistentResourceResourcePoolsOutputReference
type jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) AutoscalingSpec() GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference {
	var returns GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference
	_jsii_.Get(
		j,
		"autoscalingSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) AutoscalingSpecInput() *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec {
	var returns *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec
	_jsii_.Get(
		j,
		"autoscalingSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) DiskSpec() GoogleVertexAiPersistentResourceResourcePoolsDiskSpecOutputReference {
	var returns GoogleVertexAiPersistentResourceResourcePoolsDiskSpecOutputReference
	_jsii_.Get(
		j,
		"diskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) DiskSpecInput() *GoogleVertexAiPersistentResourceResourcePoolsDiskSpec {
	var returns *GoogleVertexAiPersistentResourceResourcePoolsDiskSpec
	_jsii_.Get(
		j,
		"diskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) MachineSpec() GoogleVertexAiPersistentResourceResourcePoolsMachineSpecOutputReference {
	var returns GoogleVertexAiPersistentResourceResourcePoolsMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) MachineSpecInput() *GoogleVertexAiPersistentResourceResourcePoolsMachineSpec {
	var returns *GoogleVertexAiPersistentResourceResourcePoolsMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ReplicaCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ReplicaCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) UsedReplicaCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usedReplicaCount",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiPersistentResourceResourcePoolsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleVertexAiPersistentResourceResourcePoolsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiPersistentResourceResourcePoolsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPersistentResource.GoogleVertexAiPersistentResourceResourcePoolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleVertexAiPersistentResourceResourcePoolsOutputReference_Override(g GoogleVertexAiPersistentResourceResourcePoolsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPersistentResource.GoogleVertexAiPersistentResourceResourcePoolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference)SetReplicaCount(val *string) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) PutAutoscalingSpec(value *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec) {
	if err := g.validatePutAutoscalingSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscalingSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) PutDiskSpec(value *GoogleVertexAiPersistentResourceResourcePoolsDiskSpec) {
	if err := g.validatePutDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiskSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) PutMachineSpec(value *GoogleVertexAiPersistentResourceResourcePoolsMachineSpec) {
	if err := g.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ResetAutoscalingSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscalingSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ResetDiskSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

