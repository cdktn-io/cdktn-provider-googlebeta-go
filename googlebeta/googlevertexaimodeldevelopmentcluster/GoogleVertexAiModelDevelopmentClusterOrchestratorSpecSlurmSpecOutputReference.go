// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexaimodeldevelopmentcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference interface {
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
	Config() *map[string]*string
	SetConfig(val *map[string]*string)
	ConfigInput() *map[string]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultPartitionId() *string
	SetDefaultPartitionId(val *string)
	DefaultPartitionIdInput() *string
	EpilogBashScripts() *[]*string
	SetEpilogBashScripts(val *[]*string)
	EpilogBashScriptsInput() *[]*string
	// Experimental.
	Fqn() *string
	HomeDirectoryStorage() *string
	SetHomeDirectoryStorage(val *string)
	HomeDirectoryStorageInput() *string
	InternalValue() *GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpec
	SetInternalValue(val *GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpec)
	LoginNodePoolId() *string
	SetLoginNodePoolId(val *string)
	LoginNodePoolIdInput() *string
	NodeSets() GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecNodeSetsList
	NodeSetsInput() interface{}
	Partitions() GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecPartitionsList
	PartitionsInput() interface{}
	PrologBashScripts() *[]*string
	SetPrologBashScripts(val *[]*string)
	PrologBashScriptsInput() *[]*string
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
	PutNodeSets(value interface{})
	PutPartitions(value interface{})
	ResetConfig()
	ResetDefaultPartitionId()
	ResetEpilogBashScripts()
	ResetNodeSets()
	ResetPartitions()
	ResetPrologBashScripts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference
type jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) Config() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) DefaultPartitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPartitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) DefaultPartitionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPartitionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) EpilogBashScripts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"epilogBashScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) EpilogBashScriptsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"epilogBashScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) HomeDirectoryStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) HomeDirectoryStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) InternalValue() *GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpec {
	var returns *GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) LoginNodePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginNodePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) LoginNodePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginNodePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) NodeSets() GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecNodeSetsList {
	var returns GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecNodeSetsList
	_jsii_.Get(
		j,
		"nodeSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) NodeSetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) Partitions() GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecPartitionsList {
	var returns GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecPartitionsList
	_jsii_.Get(
		j,
		"partitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) PartitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) PrologBashScripts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prologBashScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) PrologBashScriptsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prologBashScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiModelDevelopmentCluster.GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference_Override(g GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiModelDevelopmentCluster.GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetConfig(val *map[string]*string) {
	if err := j.validateSetConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"config",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetDefaultPartitionId(val *string) {
	if err := j.validateSetDefaultPartitionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPartitionId",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetEpilogBashScripts(val *[]*string) {
	if err := j.validateSetEpilogBashScriptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"epilogBashScripts",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetHomeDirectoryStorage(val *string) {
	if err := j.validateSetHomeDirectoryStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homeDirectoryStorage",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetInternalValue(val *GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetLoginNodePoolId(val *string) {
	if err := j.validateSetLoginNodePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginNodePoolId",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetPrologBashScripts(val *[]*string) {
	if err := j.validateSetPrologBashScriptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prologBashScripts",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) PutNodeSets(value interface{}) {
	if err := g.validatePutNodeSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeSets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) PutPartitions(value interface{}) {
	if err := g.validatePutPartitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPartitions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ResetConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ResetDefaultPartitionId() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultPartitionId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ResetEpilogBashScripts() {
	_jsii_.InvokeVoid(
		g,
		"resetEpilogBashScripts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ResetNodeSets() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeSets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ResetPartitions() {
	_jsii_.InvokeVoid(
		g,
		"resetPartitions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ResetPrologBashScripts() {
	_jsii_.InvokeVoid(
		g,
		"resetPrologBashScripts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

