// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference interface {
	cdktn.ComplexObject
	AccountKeySecret() *string
	SetAccountKeySecret(val *string)
	AccountKeySecretInput() *string
	AzureAccount() *string
	SetAzureAccount(val *string)
	AzureAccountInput() *string
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
	Container() *string
	SetContainer(val *string)
	ContainerInput() *string
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
	InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorage
	SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorage)
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
	ResetAccountKeySecret()
	ResetEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) AccountKeySecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKeySecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) AccountKeySecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKeySecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) AzureAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) AzureAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) Container() *string {
	var returns *string
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) ContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorage {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference_Override(g GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference)SetAccountKeySecret(val *string) {
	if err := j.validateSetAccountKeySecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountKeySecret",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference)SetAzureAccount(val *string) {
	if err := j.validateSetAzureAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference)SetContainer(val *string) {
	if err := j.validateSetContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) ResetAccountKeySecret() {
	_jsii_.InvokeVoid(
		g,
		"resetAccountKeySecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorageAzureDataLakeStorageIcebergStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

