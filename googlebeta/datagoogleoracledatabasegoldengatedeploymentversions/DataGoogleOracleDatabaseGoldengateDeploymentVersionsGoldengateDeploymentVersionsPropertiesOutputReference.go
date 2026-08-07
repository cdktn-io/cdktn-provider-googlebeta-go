// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagoogleoracledatabasegoldengatedeploymentversions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/datagoogleoracledatabasegoldengatedeploymentversions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference interface {
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
	DeploymentType() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsProperties
	SetInternalValue(val *DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsProperties)
	OggVersion() *string
	ReleaseTime() *string
	ReleaseType() *string
	SecurityFix() cdktn.IResolvable
	SupportEndTime() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference
type jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) InternalValue() *DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsProperties {
	var returns *DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) OggVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oggVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) ReleaseTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) ReleaseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) SecurityFix() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"securityFix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) SupportEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleOracleDatabaseGoldengateDeploymentVersions.DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference_Override(d DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleOracleDatabaseGoldengateDeploymentVersions.DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference)SetInternalValue(val *DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseGoldengateDeploymentVersionsGoldengateDeploymentVersionsPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

