// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglestoragecontrolprojectintelligencefindingrevisions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/datagooglestoragecontrolprojectintelligencefindingrevisions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference interface {
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
	InternalValue() *DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixes
	SetInternalValue(val *DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixes)
	PercentageIncrease() *float64
	Prefix() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThrottledRequests() *string
	TotalEgressBytes() *string
	TotalOperationsCount() *string
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

// The jsii proxy struct for DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference
type jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) InternalValue() *DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixes {
	var returns *DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) PercentageIncrease() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentageIncrease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) ThrottledRequests() *string {
	var returns *string
	_jsii_.Get(
		j,
		"throttledRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) TotalEgressBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalEgressBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) TotalOperationsCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalOperationsCount",
		&returns,
	)
	return returns
}


func NewDataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleStorageControlProjectIntelligenceFindingRevisions.DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference_Override(d DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleStorageControlProjectIntelligenceFindingRevisions.DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference)SetInternalValue(val *DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingRevisionsRevisionsSnapshotThrottledRequestsSpikeTopBucketsContributionTopPrefixesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

