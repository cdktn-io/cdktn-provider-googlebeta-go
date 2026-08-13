// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googledataproccluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference interface {
	cdktn.ComplexObject
	AttachedDiskConfig() GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfigList
	AttachedDiskConfigInput() interface{}
	BootDiskProvisionedIops() *float64
	SetBootDiskProvisionedIops(val *float64)
	BootDiskProvisionedIopsInput() *float64
	BootDiskProvisionedThroughput() *float64
	SetBootDiskProvisionedThroughput(val *float64)
	BootDiskProvisionedThroughputInput() *float64
	BootDiskSizeGb() *float64
	SetBootDiskSizeGb(val *float64)
	BootDiskSizeGbInput() *float64
	BootDiskType() *string
	SetBootDiskType(val *string)
	BootDiskTypeInput() *string
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
	InternalValue() *GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig
	SetInternalValue(val *GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig)
	LocalSsdInterface() *string
	SetLocalSsdInterface(val *string)
	LocalSsdInterfaceInput() *string
	NumLocalSsds() *float64
	SetNumLocalSsds(val *float64)
	NumLocalSsdsInput() *float64
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
	PutAttachedDiskConfig(value interface{})
	ResetAttachedDiskConfig()
	ResetBootDiskProvisionedIops()
	ResetBootDiskProvisionedThroughput()
	ResetBootDiskSizeGb()
	ResetBootDiskType()
	ResetLocalSsdInterface()
	ResetNumLocalSsds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference
type jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) AttachedDiskConfig() GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfigList {
	var returns GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfigList
	_jsii_.Get(
		j,
		"attachedDiskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) AttachedDiskConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachedDiskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) InternalValue() *GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig {
	var returns *GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) LocalSsdInterface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) LocalSsdInterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) NumLocalSsds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) NumLocalSsdsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference_Override(g GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskProvisionedIops(val *float64) {
	if err := j.validateSetBootDiskProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskProvisionedIops",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskProvisionedThroughput(val *float64) {
	if err := j.validateSetBootDiskProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskProvisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskType(val *string) {
	if err := j.validateSetBootDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetInternalValue(val *GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetLocalSsdInterface(val *string) {
	if err := j.validateSetLocalSsdInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdInterface",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetNumLocalSsds(val *float64) {
	if err := j.validateSetNumLocalSsdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numLocalSsds",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) PutAttachedDiskConfig(value interface{}) {
	if err := g.validatePutAttachedDiskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAttachedDiskConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetAttachedDiskConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAttachedDiskConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskProvisionedIops() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskProvisionedIops",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskProvisionedThroughput() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskProvisionedThroughput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetLocalSsdInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalSsdInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetNumLocalSsds() {
	_jsii_.InvokeVoid(
		g,
		"resetNumLocalSsds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

