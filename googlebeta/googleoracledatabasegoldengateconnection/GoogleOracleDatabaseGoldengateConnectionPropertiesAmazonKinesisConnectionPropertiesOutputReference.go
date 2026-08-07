// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference interface {
	cdktn.ComplexObject
	AccessKeyId() *string
	SetAccessKeyId(val *string)
	AccessKeyIdInput() *string
	AwsRegion() *string
	SetAwsRegion(val *string)
	AwsRegionInput() *string
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
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties
	SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties)
	SecretAccessKeySecret() *string
	SetSecretAccessKeySecret(val *string)
	SecretAccessKeySecretInput() *string
	TechnologyType() *string
	SetTechnologyType(val *string)
	TechnologyTypeInput() *string
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
	ResetAccessKeyId()
	ResetAwsRegion()
	ResetEndpoint()
	ResetSecretAccessKeySecret()
	ResetTechnologyType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) AccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) AccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) SecretAccessKeySecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeySecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) SecretAccessKeySecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeySecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference_Override(g GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetAccessKeyId(val *string) {
	if err := j.validateSetAccessKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKeyId",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetAwsRegion(val *string) {
	if err := j.validateSetAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetSecretAccessKeySecret(val *string) {
	if err := j.validateSetSecretAccessKeySecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretAccessKeySecret",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) ResetAccessKeyId() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessKeyId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) ResetSecretAccessKeySecret() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretAccessKeySecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		g,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesAmazonKinesisConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

