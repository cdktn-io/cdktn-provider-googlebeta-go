// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlecontainercluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference interface {
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
	GcpSecretManagerSecretUri() *string
	SetGcpSecretManagerSecretUri(val *string)
	GcpSecretManagerSecretUriInput() *string
	GcsGeneration() *float64
	SetGcsGeneration(val *float64)
	GcsGenerationInput() *float64
	GcsUri() *string
	SetGcsUri(val *string)
	GcsUriInput() *string
	InternalValue() *GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScript
	SetInternalValue(val *GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScript)
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
	ResetGcpSecretManagerSecretUri()
	ResetGcsGeneration()
	ResetGcsUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference
type jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcpSecretManagerSecretUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpSecretManagerSecretUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcpSecretManagerSecretUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpSecretManagerSecretUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcsGeneration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gcsGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcsGenerationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gcsGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcsUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) InternalValue() *GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScript {
	var returns *GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScript
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference_Override(g GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetGcpSecretManagerSecretUri(val *string) {
	if err := j.validateSetGcpSecretManagerSecretUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpSecretManagerSecretUri",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetGcsGeneration(val *float64) {
	if err := j.validateSetGcsGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsGeneration",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetGcsUri(val *string) {
	if err := j.validateSetGcsUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsUri",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetInternalValue(val *GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScript) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ResetGcpSecretManagerSecretUri() {
	_jsii_.InvokeVoid(
		g,
		"resetGcpSecretManagerSecretUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ResetGcsGeneration() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsGeneration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ResetGcsUri() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

