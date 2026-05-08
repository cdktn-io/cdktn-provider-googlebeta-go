// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamworkforcepoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googleiamworkforcepoolprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider google_iam_workforce_pool_provider}.
type GoogleIamWorkforcePoolProvider interface {
	cdktn.TerraformResource
	AttributeCondition() *string
	SetAttributeCondition(val *string)
	AttributeConditionInput() *string
	AttributeMapping() *map[string]*string
	SetAttributeMapping(val *map[string]*string)
	AttributeMappingInput() *map[string]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DetailedAuditLogging() interface{}
	SetDetailedAuditLogging(val interface{})
	DetailedAuditLoggingInput() interface{}
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExtendedAttributesOauth2Client() GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference
	ExtendedAttributesOauth2ClientInput() *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client
	ExtraAttributesOauth2Client() GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference
	ExtraAttributesOauth2ClientInput() *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Oidc() GoogleIamWorkforcePoolProviderOidcOutputReference
	OidcInput() *GoogleIamWorkforcePoolProviderOidc
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderId() *string
	SetProviderId(val *string)
	ProviderIdInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Saml() GoogleIamWorkforcePoolProviderSamlOutputReference
	SamlInput() *GoogleIamWorkforcePoolProviderSaml
	ScimUsage() *string
	SetScimUsage(val *string)
	ScimUsageInput() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleIamWorkforcePoolProviderTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkforcePoolId() *string
	SetWorkforcePoolId(val *string)
	WorkforcePoolIdInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutExtendedAttributesOauth2Client(value *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client)
	PutExtraAttributesOauth2Client(value *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client)
	PutOidc(value *GoogleIamWorkforcePoolProviderOidc)
	PutSaml(value *GoogleIamWorkforcePoolProviderSaml)
	PutTimeouts(value *GoogleIamWorkforcePoolProviderTimeouts)
	ResetAttributeCondition()
	ResetAttributeMapping()
	ResetDescription()
	ResetDetailedAuditLogging()
	ResetDisabled()
	ResetDisplayName()
	ResetExtendedAttributesOauth2Client()
	ResetExtraAttributesOauth2Client()
	ResetId()
	ResetOidc()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSaml()
	ResetScimUsage()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for GoogleIamWorkforcePoolProvider
type jsiiProxy_GoogleIamWorkforcePoolProvider struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) AttributeCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) AttributeConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) AttributeMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) AttributeMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) DetailedAuditLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedAuditLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) DetailedAuditLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedAuditLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ExtendedAttributesOauth2Client() GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference {
	var returns GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference
	_jsii_.Get(
		j,
		"extendedAttributesOauth2Client",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ExtendedAttributesOauth2ClientInput() *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client {
	var returns *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client
	_jsii_.Get(
		j,
		"extendedAttributesOauth2ClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ExtraAttributesOauth2Client() GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference {
	var returns GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference
	_jsii_.Get(
		j,
		"extraAttributesOauth2Client",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ExtraAttributesOauth2ClientInput() *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client {
	var returns *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client
	_jsii_.Get(
		j,
		"extraAttributesOauth2ClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Oidc() GoogleIamWorkforcePoolProviderOidcOutputReference {
	var returns GoogleIamWorkforcePoolProviderOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) OidcInput() *GoogleIamWorkforcePoolProviderOidc {
	var returns *GoogleIamWorkforcePoolProviderOidc
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ProviderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ProviderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Saml() GoogleIamWorkforcePoolProviderSamlOutputReference {
	var returns GoogleIamWorkforcePoolProviderSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) SamlInput() *GoogleIamWorkforcePoolProviderSaml {
	var returns *GoogleIamWorkforcePoolProviderSaml
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ScimUsage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scimUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) ScimUsageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scimUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) Timeouts() GoogleIamWorkforcePoolProviderTimeoutsOutputReference {
	var returns GoogleIamWorkforcePoolProviderTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) WorkforcePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workforcePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider) WorkforcePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workforcePoolIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider google_iam_workforce_pool_provider} Resource.
func NewGoogleIamWorkforcePoolProvider(scope constructs.Construct, id *string, config *GoogleIamWorkforcePoolProviderConfig) GoogleIamWorkforcePoolProvider {
	_init_.Initialize()

	if err := validateNewGoogleIamWorkforcePoolProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIamWorkforcePoolProvider{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider google_iam_workforce_pool_provider} Resource.
func NewGoogleIamWorkforcePoolProvider_Override(g GoogleIamWorkforcePoolProvider, scope constructs.Construct, id *string, config *GoogleIamWorkforcePoolProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProvider",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetAttributeCondition(val *string) {
	if err := j.validateSetAttributeConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeCondition",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetAttributeMapping(val *map[string]*string) {
	if err := j.validateSetAttributeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeMapping",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetDetailedAuditLogging(val interface{}) {
	if err := j.validateSetDetailedAuditLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detailedAuditLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetProviderId(val *string) {
	if err := j.validateSetProviderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerId",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetScimUsage(val *string) {
	if err := j.validateSetScimUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scimUsage",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProvider)SetWorkforcePoolId(val *string) {
	if err := j.validateSetWorkforcePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workforcePoolId",
		val,
	)
}

// Generates CDKTN code for importing a GoogleIamWorkforcePoolProvider resource upon running "cdktn plan <stack-name>".
func GoogleIamWorkforcePoolProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleIamWorkforcePoolProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GoogleIamWorkforcePoolProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIamWorkforcePoolProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleIamWorkforcePoolProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIamWorkforcePoolProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleIamWorkforcePoolProvider_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIamWorkforcePoolProvider_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProvider",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleIamWorkforcePoolProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) PutExtendedAttributesOauth2Client(value *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client) {
	if err := g.validatePutExtendedAttributesOauth2ClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExtendedAttributesOauth2Client",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) PutExtraAttributesOauth2Client(value *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client) {
	if err := g.validatePutExtraAttributesOauth2ClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExtraAttributesOauth2Client",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) PutOidc(value *GoogleIamWorkforcePoolProviderOidc) {
	if err := g.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOidc",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) PutSaml(value *GoogleIamWorkforcePoolProviderSaml) {
	if err := g.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSaml",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) PutTimeouts(value *GoogleIamWorkforcePoolProviderTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetAttributeCondition() {
	_jsii_.InvokeVoid(
		g,
		"resetAttributeCondition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetAttributeMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetAttributeMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetDetailedAuditLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetDetailedAuditLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetExtendedAttributesOauth2Client() {
	_jsii_.InvokeVoid(
		g,
		"resetExtendedAttributesOauth2Client",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetExtraAttributesOauth2Client() {
	_jsii_.InvokeVoid(
		g,
		"resetExtraAttributesOauth2Client",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetOidc() {
	_jsii_.InvokeVoid(
		g,
		"resetOidc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetSaml() {
	_jsii_.InvokeVoid(
		g,
		"resetSaml",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetScimUsage() {
	_jsii_.InvokeVoid(
		g,
		"resetScimUsage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

