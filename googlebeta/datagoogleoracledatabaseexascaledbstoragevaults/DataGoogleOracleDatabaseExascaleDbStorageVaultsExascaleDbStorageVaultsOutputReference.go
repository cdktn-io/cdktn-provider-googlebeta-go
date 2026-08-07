// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagoogleoracledatabaseexascaledbstoragevaults

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/datagoogleoracledatabaseexascaledbstoragevaults/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference interface {
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeletionPolicy() *string
	DeletionProtection() cdktn.IResolvable
	DisplayName() *string
	EffectiveLabels() cdktn.StringMap
	EntitlementId() *string
	ExadataInfrastructure() *string
	ExascaleDbStorageVaultId() *string
	// Experimental.
	Fqn() *string
	GcpOracleZone() *string
	InternalValue() *DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaults
	SetInternalValue(val *DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaults)
	Labels() cdktn.StringMap
	Location() *string
	Name() *string
	Project() *string
	Properties() DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsPropertiesList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	TerraformLabels() cdktn.StringMap
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

// The jsii proxy struct for DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference
type jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) DeletionProtection() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) EntitlementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) ExadataInfrastructure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exadataInfrastructure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) ExascaleDbStorageVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exascaleDbStorageVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GcpOracleZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpOracleZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) InternalValue() *DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaults {
	var returns *DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaults
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) Labels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) Properties() DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsPropertiesList {
	var returns DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsPropertiesList
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleOracleDatabaseExascaleDbStorageVaults.DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference_Override(d DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleOracleDatabaseExascaleDbStorageVaults.DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference)SetInternalValue(val *DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaults) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseExascaleDbStorageVaultsExascaleDbStorageVaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

