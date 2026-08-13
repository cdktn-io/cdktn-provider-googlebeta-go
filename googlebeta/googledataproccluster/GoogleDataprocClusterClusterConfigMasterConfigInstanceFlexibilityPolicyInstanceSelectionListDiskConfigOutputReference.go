// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googledataproccluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference interface {
	cdktn.ComplexObject
	AttachedDiskConfig() GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfigList
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
	InternalValue() *GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig
	SetInternalValue(val *GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig)
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

// The jsii proxy struct for GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference
type jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) AttachedDiskConfig() GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfigList {
	var returns GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfigList
	_jsii_.Get(
		j,
		"attachedDiskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) AttachedDiskConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachedDiskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) InternalValue() *GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig {
	var returns *GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) LocalSsdInterface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) LocalSsdInterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) NumLocalSsds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) NumLocalSsdsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference_Override(g GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskProvisionedIops(val *float64) {
	if err := j.validateSetBootDiskProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskProvisionedIops",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskProvisionedThroughput(val *float64) {
	if err := j.validateSetBootDiskProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskProvisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskType(val *string) {
	if err := j.validateSetBootDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetInternalValue(val *GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetLocalSsdInterface(val *string) {
	if err := j.validateSetLocalSsdInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdInterface",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetNumLocalSsds(val *float64) {
	if err := j.validateSetNumLocalSsdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numLocalSsds",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) PutAttachedDiskConfig(value interface{}) {
	if err := g.validatePutAttachedDiskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAttachedDiskConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetAttachedDiskConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAttachedDiskConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskProvisionedIops() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskProvisionedIops",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskProvisionedThroughput() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskProvisionedThroughput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetLocalSsdInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalSsdInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetNumLocalSsds() {
	_jsii_.InvokeVoid(
		g,
		"resetNumLocalSsds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

