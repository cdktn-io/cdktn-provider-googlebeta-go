// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googledatalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference interface {
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
	DetectionRules() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesList
	DetectionRulesInput() interface{}
	Dictionary() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionaryOutputReference
	DictionaryInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary
	ExclusionType() *string
	SetExclusionType(val *string)
	ExclusionTypeInput() *string
	FileLabelInfoType() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeOutputReference
	FileLabelInfoTypeInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType
	// Experimental.
	Fqn() *string
	InfoType() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoTypeOutputReference
	InfoTypeInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Likelihood() *string
	SetLikelihood(val *string)
	LikelihoodInput() *string
	MetadataKeyValueExpression() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpressionOutputReference
	MetadataKeyValueExpressionInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression
	Regex() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegexOutputReference
	RegexInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex
	SensitivityScore() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScoreOutputReference
	SensitivityScoreInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore
	StoredType() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredTypeOutputReference
	StoredTypeInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType
	SurrogateType() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateTypeOutputReference
	SurrogateTypeInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType
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
	PutDetectionRules(value interface{})
	PutDictionary(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary)
	PutFileLabelInfoType(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType)
	PutInfoType(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType)
	PutMetadataKeyValueExpression(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression)
	PutRegex(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex)
	PutSensitivityScore(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore)
	PutStoredType(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType)
	PutSurrogateType(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType)
	ResetDetectionRules()
	ResetDictionary()
	ResetExclusionType()
	ResetFileLabelInfoType()
	ResetLikelihood()
	ResetMetadataKeyValueExpression()
	ResetRegex()
	ResetSensitivityScore()
	ResetStoredType()
	ResetSurrogateType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference
type jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) DetectionRules() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesList {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDetectionRulesList
	_jsii_.Get(
		j,
		"detectionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) DetectionRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Dictionary() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionaryOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionaryOutputReference
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) DictionaryInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ExclusionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ExclusionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) FileLabelInfoType() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeOutputReference
	_jsii_.Get(
		j,
		"fileLabelInfoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) FileLabelInfoTypeInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType
	_jsii_.Get(
		j,
		"fileLabelInfoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InfoType() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoTypeOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoTypeOutputReference
	_jsii_.Get(
		j,
		"infoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InfoTypeInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType
	_jsii_.Get(
		j,
		"infoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Likelihood() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) LikelihoodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) MetadataKeyValueExpression() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpressionOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpressionOutputReference
	_jsii_.Get(
		j,
		"metadataKeyValueExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) MetadataKeyValueExpressionInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression
	_jsii_.Get(
		j,
		"metadataKeyValueExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Regex() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegexOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegexOutputReference
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) RegexInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) SensitivityScore() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScoreOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScoreOutputReference
	_jsii_.Get(
		j,
		"sensitivityScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) SensitivityScoreInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore
	_jsii_.Get(
		j,
		"sensitivityScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) StoredType() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredTypeOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredTypeOutputReference
	_jsii_.Get(
		j,
		"storedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) StoredTypeInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType
	_jsii_.Get(
		j,
		"storedTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) SurrogateType() GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateTypeOutputReference {
	var returns GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateTypeOutputReference
	_jsii_.Get(
		j,
		"surrogateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) SurrogateTypeInput() *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType {
	var returns *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType
	_jsii_.Get(
		j,
		"surrogateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference_Override(g GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionContentPolicy.GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetExclusionType(val *string) {
	if err := j.validateSetExclusionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetLikelihood(val *string) {
	if err := j.validateSetLikelihoodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likelihood",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutDetectionRules(value interface{}) {
	if err := g.validatePutDetectionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDetectionRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutDictionary(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary) {
	if err := g.validatePutDictionaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDictionary",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutFileLabelInfoType(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType) {
	if err := g.validatePutFileLabelInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFileLabelInfoType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutInfoType(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType) {
	if err := g.validatePutInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInfoType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutMetadataKeyValueExpression(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression) {
	if err := g.validatePutMetadataKeyValueExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetadataKeyValueExpression",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutRegex(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex) {
	if err := g.validatePutRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegex",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutSensitivityScore(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore) {
	if err := g.validatePutSensitivityScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSensitivityScore",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutStoredType(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType) {
	if err := g.validatePutStoredTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStoredType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) PutSurrogateType(value *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType) {
	if err := g.validatePutSurrogateTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSurrogateType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetDetectionRules() {
	_jsii_.InvokeVoid(
		g,
		"resetDetectionRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		g,
		"resetDictionary",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetExclusionType() {
	_jsii_.InvokeVoid(
		g,
		"resetExclusionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetFileLabelInfoType() {
	_jsii_.InvokeVoid(
		g,
		"resetFileLabelInfoType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetLikelihood() {
	_jsii_.InvokeVoid(
		g,
		"resetLikelihood",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetMetadataKeyValueExpression() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataKeyValueExpression",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetSensitivityScore() {
	_jsii_.InvokeVoid(
		g,
		"resetSensitivityScore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetStoredType() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ResetSurrogateType() {
	_jsii_.InvokeVoid(
		g,
		"resetSurrogateType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

