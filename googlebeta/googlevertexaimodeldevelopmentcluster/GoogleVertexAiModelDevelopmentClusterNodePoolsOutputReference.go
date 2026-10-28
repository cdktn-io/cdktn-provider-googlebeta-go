// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexaimodeldevelopmentcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference interface {
	cdktn.ComplexObject
	BootDisk() GoogleVertexAiModelDevelopmentClusterNodePoolsBootDiskOutputReference
	BootDiskInput() *GoogleVertexAiModelDevelopmentClusterNodePoolsBootDisk
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
	EffectiveFlexStartMaxDuration() *string
	EnablePublicIps() interface{}
	SetEnablePublicIps(val interface{})
	EnablePublicIpsInput() interface{}
	Filestores() *[]*string
	SetFilestores(val *[]*string)
	FilestoresInput() *[]*string
	FlexStartMaxDuration() *string
	SetFlexStartMaxDuration(val *string)
	FlexStartMaxDurationInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Lustres() *[]*string
	SetLustres(val *[]*string)
	LustresInput() *[]*string
	MachineSpec() GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpecOutputReference
	MachineSpecInput() *GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpec
	NodeImage() *string
	SetNodeImage(val *string)
	NodeImageInput() *string
	ProvisioningModel() *string
	SetProvisioningModel(val *string)
	ProvisioningModelInput() *string
	ScalingSpec() GoogleVertexAiModelDevelopmentClusterNodePoolsScalingSpecOutputReference
	ScalingSpecInput() *GoogleVertexAiModelDevelopmentClusterNodePoolsScalingSpec
	StartupScript() *string
	SetStartupScript(val *string)
	StartupScriptInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutBootDisk(value *GoogleVertexAiModelDevelopmentClusterNodePoolsBootDisk)
	PutMachineSpec(value *GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpec)
	PutScalingSpec(value *GoogleVertexAiModelDevelopmentClusterNodePoolsScalingSpec)
	ResetBootDisk()
	ResetEnablePublicIps()
	ResetFilestores()
	ResetFlexStartMaxDuration()
	ResetLabels()
	ResetLustres()
	ResetNodeImage()
	ResetProvisioningModel()
	ResetScalingSpec()
	ResetStartupScript()
	ResetZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference
type jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) BootDisk() GoogleVertexAiModelDevelopmentClusterNodePoolsBootDiskOutputReference {
	var returns GoogleVertexAiModelDevelopmentClusterNodePoolsBootDiskOutputReference
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) BootDiskInput() *GoogleVertexAiModelDevelopmentClusterNodePoolsBootDisk {
	var returns *GoogleVertexAiModelDevelopmentClusterNodePoolsBootDisk
	_jsii_.Get(
		j,
		"bootDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) EffectiveFlexStartMaxDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveFlexStartMaxDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) EnablePublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) EnablePublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) Filestores() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filestores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) FilestoresInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filestoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) FlexStartMaxDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flexStartMaxDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) FlexStartMaxDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flexStartMaxDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) Lustres() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lustres",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) LustresInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lustresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) MachineSpec() GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpecOutputReference {
	var returns GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) MachineSpecInput() *GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpec {
	var returns *GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) NodeImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) NodeImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ProvisioningModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ProvisioningModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ScalingSpec() GoogleVertexAiModelDevelopmentClusterNodePoolsScalingSpecOutputReference {
	var returns GoogleVertexAiModelDevelopmentClusterNodePoolsScalingSpecOutputReference
	_jsii_.Get(
		j,
		"scalingSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ScalingSpecInput() *GoogleVertexAiModelDevelopmentClusterNodePoolsScalingSpec {
	var returns *GoogleVertexAiModelDevelopmentClusterNodePoolsScalingSpec
	_jsii_.Get(
		j,
		"scalingSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) StartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) StartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiModelDevelopmentClusterNodePoolsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiModelDevelopmentCluster.GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference_Override(g GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiModelDevelopmentCluster.GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetEnablePublicIps(val interface{}) {
	if err := j.validateSetEnablePublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePublicIps",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetFilestores(val *[]*string) {
	if err := j.validateSetFilestoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filestores",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetFlexStartMaxDuration(val *string) {
	if err := j.validateSetFlexStartMaxDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flexStartMaxDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetLustres(val *[]*string) {
	if err := j.validateSetLustresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lustres",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetNodeImage(val *string) {
	if err := j.validateSetNodeImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeImage",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetProvisioningModel(val *string) {
	if err := j.validateSetProvisioningModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningModel",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetStartupScript(val *string) {
	if err := j.validateSetStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupScript",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) PutBootDisk(value *GoogleVertexAiModelDevelopmentClusterNodePoolsBootDisk) {
	if err := g.validatePutBootDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBootDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) PutMachineSpec(value *GoogleVertexAiModelDevelopmentClusterNodePoolsMachineSpec) {
	if err := g.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) PutScalingSpec(value *GoogleVertexAiModelDevelopmentClusterNodePoolsScalingSpec) {
	if err := g.validatePutScalingSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScalingSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetBootDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetEnablePublicIps() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePublicIps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetFilestores() {
	_jsii_.InvokeVoid(
		g,
		"resetFilestores",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetFlexStartMaxDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetFlexStartMaxDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetLustres() {
	_jsii_.InvokeVoid(
		g,
		"resetLustres",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetNodeImage() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetProvisioningModel() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisioningModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetScalingSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetScalingSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetStartupScript() {
	_jsii_.InvokeVoid(
		g,
		"resetStartupScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterNodePoolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

