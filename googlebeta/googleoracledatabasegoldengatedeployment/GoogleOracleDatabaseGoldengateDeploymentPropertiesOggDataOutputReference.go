// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengatedeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengatedeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference interface {
	cdktn.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminPasswordSecretVersion() *string
	SetAdminPasswordSecretVersion(val *string)
	AdminPasswordSecretVersionInput() *string
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	Certificate() *string
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
	CredentialStore() *string
	Deployment() *string
	SetDeployment(val *string)
	DeploymentInput() *string
	// Experimental.
	Fqn() *string
	GroupRolesMapping() GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMappingOutputReference
	GroupRolesMappingInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping
	IdentityDomainId() *string
	InternalValue() *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData
	SetInternalValue(val *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData)
	OggVersion() *string
	SetOggVersion(val *string)
	OggVersionInput() *string
	PasswordSecretId() *string
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
	PutGroupRolesMapping(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping)
	ResetAdminPassword()
	ResetAdminPasswordSecretVersion()
	ResetGroupRolesMapping()
	ResetOggVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminPasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminPasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) CredentialStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) Deployment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) DeploymentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GroupRolesMapping() GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMappingOutputReference {
	var returns GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMappingOutputReference
	_jsii_.Get(
		j,
		"groupRolesMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GroupRolesMappingInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping {
	var returns *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping
	_jsii_.Get(
		j,
		"groupRolesMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) IdentityDomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData {
	var returns *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) OggVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oggVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) OggVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oggVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) PasswordSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference_Override(g GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetAdminPasswordSecretVersion(val *string) {
	if err := j.validateSetAdminPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetDeployment(val *string) {
	if err := j.validateSetDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployment",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetOggVersion(val *string) {
	if err := j.validateSetOggVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oggVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) PutGroupRolesMapping(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping) {
	if err := g.validatePutGroupRolesMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGroupRolesMapping",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ResetAdminPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetAdminPasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ResetGroupRolesMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupRolesMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ResetOggVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetOggVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

