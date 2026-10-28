// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlecloudrunv2instance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCloudRunV2InstanceVolumesOutputReference interface {
	cdktn.ComplexObject
	CloudSqlInstance() GoogleCloudRunV2InstanceVolumesCloudSqlInstanceOutputReference
	CloudSqlInstanceInput() *GoogleCloudRunV2InstanceVolumesCloudSqlInstance
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
	EmptyDir() GoogleCloudRunV2InstanceVolumesEmptyDirOutputReference
	EmptyDirInput() *GoogleCloudRunV2InstanceVolumesEmptyDir
	// Experimental.
	Fqn() *string
	Gcs() GoogleCloudRunV2InstanceVolumesGcsOutputReference
	GcsInput() *GoogleCloudRunV2InstanceVolumesGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() GoogleCloudRunV2InstanceVolumesNfsOutputReference
	NfsInput() *GoogleCloudRunV2InstanceVolumesNfs
	Secret() GoogleCloudRunV2InstanceVolumesSecretOutputReference
	SecretInput() *GoogleCloudRunV2InstanceVolumesSecret
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
	PutCloudSqlInstance(value *GoogleCloudRunV2InstanceVolumesCloudSqlInstance)
	PutEmptyDir(value *GoogleCloudRunV2InstanceVolumesEmptyDir)
	PutGcs(value *GoogleCloudRunV2InstanceVolumesGcs)
	PutNfs(value *GoogleCloudRunV2InstanceVolumesNfs)
	PutSecret(value *GoogleCloudRunV2InstanceVolumesSecret)
	ResetCloudSqlInstance()
	ResetEmptyDir()
	ResetGcs()
	ResetNfs()
	ResetSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudRunV2InstanceVolumesOutputReference
type jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) CloudSqlInstance() GoogleCloudRunV2InstanceVolumesCloudSqlInstanceOutputReference {
	var returns GoogleCloudRunV2InstanceVolumesCloudSqlInstanceOutputReference
	_jsii_.Get(
		j,
		"cloudSqlInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) CloudSqlInstanceInput() *GoogleCloudRunV2InstanceVolumesCloudSqlInstance {
	var returns *GoogleCloudRunV2InstanceVolumesCloudSqlInstance
	_jsii_.Get(
		j,
		"cloudSqlInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) EmptyDir() GoogleCloudRunV2InstanceVolumesEmptyDirOutputReference {
	var returns GoogleCloudRunV2InstanceVolumesEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) EmptyDirInput() *GoogleCloudRunV2InstanceVolumesEmptyDir {
	var returns *GoogleCloudRunV2InstanceVolumesEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) Gcs() GoogleCloudRunV2InstanceVolumesGcsOutputReference {
	var returns GoogleCloudRunV2InstanceVolumesGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GcsInput() *GoogleCloudRunV2InstanceVolumesGcs {
	var returns *GoogleCloudRunV2InstanceVolumesGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) Nfs() GoogleCloudRunV2InstanceVolumesNfsOutputReference {
	var returns GoogleCloudRunV2InstanceVolumesNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) NfsInput() *GoogleCloudRunV2InstanceVolumesNfs {
	var returns *GoogleCloudRunV2InstanceVolumesNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) Secret() GoogleCloudRunV2InstanceVolumesSecretOutputReference {
	var returns GoogleCloudRunV2InstanceVolumesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) SecretInput() *GoogleCloudRunV2InstanceVolumesSecret {
	var returns *GoogleCloudRunV2InstanceVolumesSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunV2InstanceVolumesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCloudRunV2InstanceVolumesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2InstanceVolumesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCloudRunV2Instance.GoogleCloudRunV2InstanceVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2InstanceVolumesOutputReference_Override(g GoogleCloudRunV2InstanceVolumesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCloudRunV2Instance.GoogleCloudRunV2InstanceVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) PutCloudSqlInstance(value *GoogleCloudRunV2InstanceVolumesCloudSqlInstance) {
	if err := g.validatePutCloudSqlInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudSqlInstance",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) PutEmptyDir(value *GoogleCloudRunV2InstanceVolumesEmptyDir) {
	if err := g.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) PutGcs(value *GoogleCloudRunV2InstanceVolumesGcs) {
	if err := g.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) PutNfs(value *GoogleCloudRunV2InstanceVolumesNfs) {
	if err := g.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNfs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) PutSecret(value *GoogleCloudRunV2InstanceVolumesSecret) {
	if err := g.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecret",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) ResetCloudSqlInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudSqlInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		g,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		g,
		"resetGcs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		g,
		"resetNfs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2InstanceVolumesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

