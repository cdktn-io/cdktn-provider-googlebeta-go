// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleartifactregistryprojectconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googleartifactregistryprojectconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference interface {
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
	InternalValue() *GoogleArtifactRegistryProjectConfigPlatformLogsConfig
	SetInternalValue(val *GoogleArtifactRegistryProjectConfigPlatformLogsConfig)
	LoggingState() *string
	SetLoggingState(val *string)
	LoggingStateInput() *string
	SeverityLevel() *string
	SetSeverityLevel(val *string)
	SeverityLevelInput() *string
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
	ResetLoggingState()
	ResetSeverityLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference
type jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) InternalValue() *GoogleArtifactRegistryProjectConfigPlatformLogsConfig {
	var returns *GoogleArtifactRegistryProjectConfigPlatformLogsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) LoggingState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) LoggingStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) SeverityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) SeverityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleArtifactRegistryProjectConfig.GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference_Override(g GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleArtifactRegistryProjectConfig.GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference)SetInternalValue(val *GoogleArtifactRegistryProjectConfigPlatformLogsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference)SetLoggingState(val *string) {
	if err := j.validateSetLoggingStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingState",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference)SetSeverityLevel(val *string) {
	if err := j.validateSetSeverityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severityLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) ResetLoggingState() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) ResetSeverityLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetSeverityLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryProjectConfigPlatformLogsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

