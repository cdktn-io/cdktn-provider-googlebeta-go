// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemigrationcenterreport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlemigrationcenterreport/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference interface {
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
	ComputeEngineFinding() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsComputeEngineFindingList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	DisplayName() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindings
	SetInternalValue(val *GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindings)
	MachinePreferences() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesList
	MonthlyCostCompute() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostComputeList
	MonthlyCostNetworkEgress() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostNetworkEgressList
	MonthlyCostOsLicense() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOsLicenseList
	MonthlyCostOther() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOtherList
	MonthlyCostStorage() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostStorageList
	MonthlyCostTotal() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostTotalList
	SoleTenantFinding() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsSoleTenantFindingList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VmwareEngineFinding() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsVmwareEngineFindingList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference
type jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ComputeEngineFinding() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsComputeEngineFindingList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsComputeEngineFindingList
	_jsii_.Get(
		j,
		"computeEngineFinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) InternalValue() *GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindings {
	var returns *GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MachinePreferences() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesList
	_jsii_.Get(
		j,
		"machinePreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostCompute() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostComputeList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostComputeList
	_jsii_.Get(
		j,
		"monthlyCostCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostNetworkEgress() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostNetworkEgressList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostNetworkEgressList
	_jsii_.Get(
		j,
		"monthlyCostNetworkEgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostOsLicense() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOsLicenseList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOsLicenseList
	_jsii_.Get(
		j,
		"monthlyCostOsLicense",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostOther() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOtherList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOtherList
	_jsii_.Get(
		j,
		"monthlyCostOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostStorage() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostStorageList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostStorageList
	_jsii_.Get(
		j,
		"monthlyCostStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostTotal() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostTotalList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostTotalList
	_jsii_.Get(
		j,
		"monthlyCostTotal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) SoleTenantFinding() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsSoleTenantFindingList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsSoleTenantFindingList
	_jsii_.Get(
		j,
		"soleTenantFinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) VmwareEngineFinding() GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsVmwareEngineFindingList {
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsVmwareEngineFindingList
	_jsii_.Get(
		j,
		"vmwareEngineFinding",
		&returns,
	)
	return returns
}


func NewGoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMigrationCenterReport.GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference_Override(g GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMigrationCenterReport.GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetInternalValue(val *GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

