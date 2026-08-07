// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakeicebergcatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlebiglakeicebergcatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference interface {
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
	GlueCatalogInfo() GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference
	GlueCatalogInfoInput() *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo
	InternalValue() *GoogleBiglakeIcebergCatalogFederatedCatalogOptions
	SetInternalValue(val *GoogleBiglakeIcebergCatalogFederatedCatalogOptions)
	RefreshOptions() GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsOutputReference
	RefreshOptionsInput() *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions
	RefreshStatus() GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshStatusList
	SecretName() *string
	SetSecretName(val *string)
	SecretNameInput() *string
	ServiceDirectoryName() *string
	SetServiceDirectoryName(val *string)
	ServiceDirectoryNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnityCatalogInfo() GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference
	UnityCatalogInfoInput() *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo
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
	PutGlueCatalogInfo(value *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo)
	PutRefreshOptions(value *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions)
	PutUnityCatalogInfo(value *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo)
	ResetGlueCatalogInfo()
	ResetRefreshOptions()
	ResetSecretName()
	ResetServiceDirectoryName()
	ResetUnityCatalogInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference
type jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GlueCatalogInfo() GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference {
	var returns GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference
	_jsii_.Get(
		j,
		"glueCatalogInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GlueCatalogInfoInput() *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo {
	var returns *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo
	_jsii_.Get(
		j,
		"glueCatalogInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) InternalValue() *GoogleBiglakeIcebergCatalogFederatedCatalogOptions {
	var returns *GoogleBiglakeIcebergCatalogFederatedCatalogOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) RefreshOptions() GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsOutputReference {
	var returns GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsOutputReference
	_jsii_.Get(
		j,
		"refreshOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) RefreshOptionsInput() *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions {
	var returns *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions
	_jsii_.Get(
		j,
		"refreshOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) RefreshStatus() GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshStatusList {
	var returns GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshStatusList
	_jsii_.Get(
		j,
		"refreshStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) SecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ServiceDirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ServiceDirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) UnityCatalogInfo() GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference {
	var returns GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference
	_jsii_.Get(
		j,
		"unityCatalogInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) UnityCatalogInfoInput() *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo {
	var returns *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo
	_jsii_.Get(
		j,
		"unityCatalogInfoInput",
		&returns,
	)
	return returns
}


func NewGoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBiglakeIcebergCatalog.GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference_Override(g GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBiglakeIcebergCatalog.GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetInternalValue(val *GoogleBiglakeIcebergCatalogFederatedCatalogOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetSecretName(val *string) {
	if err := j.validateSetSecretNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretName",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetServiceDirectoryName(val *string) {
	if err := j.validateSetServiceDirectoryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceDirectoryName",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) PutGlueCatalogInfo(value *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo) {
	if err := g.validatePutGlueCatalogInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGlueCatalogInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) PutRefreshOptions(value *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions) {
	if err := g.validatePutRefreshOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRefreshOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) PutUnityCatalogInfo(value *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo) {
	if err := g.validatePutUnityCatalogInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUnityCatalogInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetGlueCatalogInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetGlueCatalogInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetRefreshOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetRefreshOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetSecretName() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetServiceDirectoryName() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetUnityCatalogInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetUnityCatalogInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

