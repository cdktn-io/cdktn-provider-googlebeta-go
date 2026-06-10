// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemigrationcenterimportjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlemigrationcenterimportjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference interface {
	cdktn.ComplexObject
	ArchiveError() GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsArchiveErrorList
	AssetTitle() *string
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
	CsvError() GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsCsvErrorList
	Errors() GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsErrorsList
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrors
	SetInternalValue(val *GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrors)
	RowNumber() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VmName() *string
	VmUuid() *string
	XlsxError() GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsXlsxErrorList
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

// The jsii proxy struct for GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference
type jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ArchiveError() GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsArchiveErrorList {
	var returns GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsArchiveErrorList
	_jsii_.Get(
		j,
		"archiveError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) AssetTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) CsvError() GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsCsvErrorList {
	var returns GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsCsvErrorList
	_jsii_.Get(
		j,
		"csvError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) Errors() GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsErrorsList {
	var returns GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsErrorsList
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) InternalValue() *GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrors {
	var returns *GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) RowNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) VmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) VmUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) XlsxError() GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsXlsxErrorList {
	var returns GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsXlsxErrorList
	_jsii_.Get(
		j,
		"xlsxError",
		&returns,
	)
	return returns
}


func NewGoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMigrationCenterImportJob.GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference_Override(g GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMigrationCenterImportJob.GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetInternalValue(val *GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrors) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

