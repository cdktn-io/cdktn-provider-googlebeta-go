// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlebiglakehivetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleBiglakeHiveTableStorageDescriptorOutputReference interface {
	cdktn.ComplexObject
	BucketCols() *[]*string
	SetBucketCols(val *[]*string)
	BucketColsInput() *[]*string
	Columns() GoogleBiglakeHiveTableStorageDescriptorColumnsList
	ColumnsInput() interface{}
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
	Compressed() interface{}
	SetCompressed(val interface{})
	CompressedInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InputFormat() *string
	SetInputFormat(val *string)
	InputFormatInput() *string
	InternalValue() *GoogleBiglakeHiveTableStorageDescriptor
	SetInternalValue(val *GoogleBiglakeHiveTableStorageDescriptor)
	LocationUri() *string
	SetLocationUri(val *string)
	LocationUriInput() *string
	NumBuckets() *float64
	SetNumBuckets(val *float64)
	NumBucketsInput() *float64
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	SerdeInfo() GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference
	SerdeInfoInput() *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo
	SkewedInfo() GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference
	SkewedInfoInput() *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo
	SortCols() GoogleBiglakeHiveTableStorageDescriptorSortColsList
	SortColsInput() interface{}
	StoredAsSubDirs() interface{}
	SetStoredAsSubDirs(val interface{})
	StoredAsSubDirsInput() interface{}
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
	PutColumns(value interface{})
	PutSerdeInfo(value *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo)
	PutSkewedInfo(value *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo)
	PutSortCols(value interface{})
	ResetBucketCols()
	ResetCompressed()
	ResetInputFormat()
	ResetLocationUri()
	ResetNumBuckets()
	ResetOutputFormat()
	ResetParameters()
	ResetSerdeInfo()
	ResetSkewedInfo()
	ResetSortCols()
	ResetStoredAsSubDirs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBiglakeHiveTableStorageDescriptorOutputReference
type jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) BucketCols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketCols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) BucketColsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) Columns() GoogleBiglakeHiveTableStorageDescriptorColumnsList {
	var returns GoogleBiglakeHiveTableStorageDescriptorColumnsList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) InternalValue() *GoogleBiglakeHiveTableStorageDescriptor {
	var returns *GoogleBiglakeHiveTableStorageDescriptor
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) LocationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) NumBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) NumBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) SerdeInfo() GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference {
	var returns GoogleBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference
	_jsii_.Get(
		j,
		"serdeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) SerdeInfoInput() *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo {
	var returns *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo
	_jsii_.Get(
		j,
		"serdeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) SkewedInfo() GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference {
	var returns GoogleBiglakeHiveTableStorageDescriptorSkewedInfoOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) SkewedInfoInput() *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo {
	var returns *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) SortCols() GoogleBiglakeHiveTableStorageDescriptorSortColsList {
	var returns GoogleBiglakeHiveTableStorageDescriptorSortColsList
	_jsii_.Get(
		j,
		"sortCols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) SortColsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortColsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) StoredAsSubDirs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) StoredAsSubDirsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBiglakeHiveTableStorageDescriptorOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBiglakeHiveTableStorageDescriptorOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBiglakeHiveTableStorageDescriptorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBiglakeHiveTable.GoogleBiglakeHiveTableStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBiglakeHiveTableStorageDescriptorOutputReference_Override(g GoogleBiglakeHiveTableStorageDescriptorOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBiglakeHiveTable.GoogleBiglakeHiveTableStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetBucketCols(val *[]*string) {
	if err := j.validateSetBucketColsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketCols",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetCompressed(val interface{}) {
	if err := j.validateSetCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetInternalValue(val *GoogleBiglakeHiveTableStorageDescriptor) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetLocationUri(val *string) {
	if err := j.validateSetLocationUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationUri",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetNumBuckets(val *float64) {
	if err := j.validateSetNumBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numBuckets",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetStoredAsSubDirs(val interface{}) {
	if err := j.validateSetStoredAsSubDirsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedAsSubDirs",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) PutColumns(value interface{}) {
	if err := g.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putColumns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) PutSerdeInfo(value *GoogleBiglakeHiveTableStorageDescriptorSerdeInfo) {
	if err := g.validatePutSerdeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSerdeInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) PutSkewedInfo(value *GoogleBiglakeHiveTableStorageDescriptorSkewedInfo) {
	if err := g.validatePutSkewedInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) PutSortCols(value interface{}) {
	if err := g.validatePutSortColsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSortCols",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetBucketCols() {
	_jsii_.InvokeVoid(
		g,
		"resetBucketCols",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		g,
		"resetCompressed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetLocationUri() {
	_jsii_.InvokeVoid(
		g,
		"resetLocationUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetNumBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetNumBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetSerdeInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSerdeInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetSortCols() {
	_jsii_.InvokeVoid(
		g,
		"resetSortCols",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ResetStoredAsSubDirs() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredAsSubDirs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBiglakeHiveTableStorageDescriptorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

