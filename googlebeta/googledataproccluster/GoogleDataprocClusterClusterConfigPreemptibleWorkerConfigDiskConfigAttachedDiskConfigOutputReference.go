// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googledataproccluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference interface {
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
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProvisionedIops() *float64
	SetProvisionedIops(val *float64)
	ProvisionedIopsInput() *float64
	ProvisionedThroughput() *float64
	SetProvisionedThroughput(val *float64)
	ProvisionedThroughputInput() *float64
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
	ResetDiskSizeGb()
	ResetDiskType()
	ResetProvisionedIops()
	ResetProvisionedThroughput()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference
type jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference_Override(g GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference)SetProvisionedIops(val *float64) {
	if err := j.validateSetProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedIops",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference)SetProvisionedThroughput(val *float64) {
	if err := j.validateSetProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ResetProvisionedIops() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisionedIops",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigAttachedDiskConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

