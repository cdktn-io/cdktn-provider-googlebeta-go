// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglestoragecontrolprojectintelligencefindings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/datagooglestoragecontrolprojectintelligencefindings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference interface {
	cdktn.ComplexObject
	Bucket() *string
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
	Contribution() DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsContributionList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Error() DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsErrorList
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBuckets
	SetInternalValue(val *DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBuckets)
	PercentageIncrease() *float64
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
	TotalStorageGrowthBytes() *string
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

// The jsii proxy struct for DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference
type jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) Contribution() DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsContributionList {
	var returns DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsContributionList
	_jsii_.Get(
		j,
		"contribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) Error() DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsErrorList {
	var returns DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsErrorList
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) InternalValue() *DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBuckets {
	var returns *DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBuckets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) PercentageIncrease() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentageIncrease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) ThrottledRequests() *string {
	var returns *string
	_jsii_.Get(
		j,
		"throttledRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) TotalEgressBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalEgressBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) TotalOperationsCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalOperationsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) TotalStorageGrowthBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalStorageGrowthBytes",
		&returns,
	)
	return returns
}


func NewDataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleStorageControlProjectIntelligenceFindings.DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference_Override(d DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleStorageControlProjectIntelligenceFindings.DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference)SetInternalValue(val *DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBuckets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingsFindingsColdlineAndArchivalStorageOperationsSpikeTopBucketsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

