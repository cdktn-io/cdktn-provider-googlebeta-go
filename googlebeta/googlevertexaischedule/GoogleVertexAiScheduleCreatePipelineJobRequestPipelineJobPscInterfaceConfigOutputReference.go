// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlevertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference interface {
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
	DnsPeeringConfigs() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsList
	DnsPeeringConfigsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	SetInternalValue(val *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig)
	NetworkAttachment() *string
	SetNetworkAttachment(val *string)
	NetworkAttachmentInput() *string
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
	PutDnsPeeringConfigs(value interface{})
	ResetDnsPeeringConfigs()
	ResetNetworkAttachment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference
type jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) DnsPeeringConfigs() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsList {
	var returns GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsList
	_jsii_.Get(
		j,
		"dnsPeeringConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) DnsPeeringConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsPeeringConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) InternalValue() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig {
	var returns *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) NetworkAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) NetworkAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference_Override(g GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetInternalValue(val *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetNetworkAttachment(val *string) {
	if err := j.validateSetNetworkAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAttachment",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) PutDnsPeeringConfigs(value interface{}) {
	if err := g.validatePutDnsPeeringConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDnsPeeringConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ResetDnsPeeringConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsPeeringConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ResetNetworkAttachment() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkAttachment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

