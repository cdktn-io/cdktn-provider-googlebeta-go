// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengatedeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengatedeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference interface {
	cdktn.ComplexObject
	BundleReleaseUpgradePeriodDays() *float64
	SetBundleReleaseUpgradePeriodDays(val *float64)
	BundleReleaseUpgradePeriodDaysInput() *float64
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
	InterimReleaseUpgradePeriodDays() *float64
	SetInterimReleaseUpgradePeriodDays(val *float64)
	InterimReleaseUpgradePeriodDaysInput() *float64
	InternalValue() *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig
	SetInternalValue(val *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig)
	IsInterimReleaseAutoUpgradeEnabled() interface{}
	SetIsInterimReleaseAutoUpgradeEnabled(val interface{})
	IsInterimReleaseAutoUpgradeEnabledInput() interface{}
	MajorReleaseUpgradePeriodDays() *float64
	SetMajorReleaseUpgradePeriodDays(val *float64)
	MajorReleaseUpgradePeriodDaysInput() *float64
	SecurityPatchUpgradePeriodDays() *float64
	SetSecurityPatchUpgradePeriodDays(val *float64)
	SecurityPatchUpgradePeriodDaysInput() *float64
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
	ResetBundleReleaseUpgradePeriodDays()
	ResetInterimReleaseUpgradePeriodDays()
	ResetIsInterimReleaseAutoUpgradeEnabled()
	ResetMajorReleaseUpgradePeriodDays()
	ResetSecurityPatchUpgradePeriodDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) BundleReleaseUpgradePeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bundleReleaseUpgradePeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) BundleReleaseUpgradePeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bundleReleaseUpgradePeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InterimReleaseUpgradePeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interimReleaseUpgradePeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InterimReleaseUpgradePeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interimReleaseUpgradePeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig {
	var returns *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) IsInterimReleaseAutoUpgradeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isInterimReleaseAutoUpgradeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) IsInterimReleaseAutoUpgradeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isInterimReleaseAutoUpgradeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) MajorReleaseUpgradePeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"majorReleaseUpgradePeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) MajorReleaseUpgradePeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"majorReleaseUpgradePeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) SecurityPatchUpgradePeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securityPatchUpgradePeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) SecurityPatchUpgradePeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securityPatchUpgradePeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference_Override(g GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetBundleReleaseUpgradePeriodDays(val *float64) {
	if err := j.validateSetBundleReleaseUpgradePeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bundleReleaseUpgradePeriodDays",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetInterimReleaseUpgradePeriodDays(val *float64) {
	if err := j.validateSetInterimReleaseUpgradePeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interimReleaseUpgradePeriodDays",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetIsInterimReleaseAutoUpgradeEnabled(val interface{}) {
	if err := j.validateSetIsInterimReleaseAutoUpgradeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isInterimReleaseAutoUpgradeEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetMajorReleaseUpgradePeriodDays(val *float64) {
	if err := j.validateSetMajorReleaseUpgradePeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"majorReleaseUpgradePeriodDays",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetSecurityPatchUpgradePeriodDays(val *float64) {
	if err := j.validateSetSecurityPatchUpgradePeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPatchUpgradePeriodDays",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetBundleReleaseUpgradePeriodDays() {
	_jsii_.InvokeVoid(
		g,
		"resetBundleReleaseUpgradePeriodDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetInterimReleaseUpgradePeriodDays() {
	_jsii_.InvokeVoid(
		g,
		"resetInterimReleaseUpgradePeriodDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetIsInterimReleaseAutoUpgradeEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIsInterimReleaseAutoUpgradeEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetMajorReleaseUpgradePeriodDays() {
	_jsii_.InvokeVoid(
		g,
		"resetMajorReleaseUpgradePeriodDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetSecurityPatchUpgradePeriodDays() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityPatchUpgradePeriodDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

