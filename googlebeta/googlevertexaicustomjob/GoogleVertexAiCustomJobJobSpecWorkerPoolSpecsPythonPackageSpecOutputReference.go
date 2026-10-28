// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexaicustomjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference interface {
	cdktn.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
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
	Env() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecEnvList
	EnvInput() interface{}
	ExecutorImageUri() *string
	SetExecutorImageUri(val *string)
	ExecutorImageUriInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec
	SetInternalValue(val *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec)
	PackageUris() *[]*string
	SetPackageUris(val *[]*string)
	PackageUrisInput() *[]*string
	PythonModule() *string
	SetPythonModule(val *string)
	PythonModuleInput() *string
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
	PutEnv(value interface{})
	ResetArgs()
	ResetEnv()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference
type jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) Env() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecEnvList {
	var returns GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ExecutorImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executorImageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ExecutorImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executorImageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) InternalValue() *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec {
	var returns *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PackageUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PackageUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PythonModule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonModule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PythonModuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonModuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference_Override(g GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetExecutorImageUri(val *string) {
	if err := j.validateSetExecutorImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executorImageUri",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetInternalValue(val *GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetPackageUris(val *[]*string) {
	if err := j.validateSetPackageUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageUris",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetPythonModule(val *string) {
	if err := j.validateSetPythonModuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonModule",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PutEnv(value interface{}) {
	if err := g.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnv",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		g,
		"resetEnv",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

