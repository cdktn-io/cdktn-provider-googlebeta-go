// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglestoragecontrolfolderintelligencefindingssummary

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/datagooglestoragecontrolfolderintelligencefindingssummary/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference interface {
	cdktn.ComplexObject
	Category() *string
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummaries
	SetInternalValue(val *DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummaries)
	Severity() *string
	SummaryDetails() DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesSummaryDetailsList
	TargetResource() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	UpdateTime() *string
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

// The jsii proxy struct for DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference
type jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) InternalValue() *DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummaries {
	var returns *DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummaries
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) SummaryDetails() DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesSummaryDetailsList {
	var returns DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesSummaryDetailsList
	_jsii_.Get(
		j,
		"summaryDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) TargetResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewDataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleStorageControlFolderIntelligenceFindingsSummary.DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference_Override(d DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleStorageControlFolderIntelligenceFindingsSummary.DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference)SetInternalValue(val *DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummaries) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlFolderIntelligenceFindingsSummaryFindingSummariesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

