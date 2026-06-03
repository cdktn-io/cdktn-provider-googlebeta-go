// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemigrationcenterreport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlemigrationcenterreport/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList
type jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList {
	_init_.Initialize()

	if err := validateNewGoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMigrationCenterReport.GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList_Override(g GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMigrationCenterReport.GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) Get(index *float64) GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

