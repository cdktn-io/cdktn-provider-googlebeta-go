// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlegeminidataanalyticsdataagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference interface {
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
	DatasourceReferences() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesOutputReference
	DatasourceReferencesInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferences
	ExampleQueries() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextExampleQueriesList
	ExampleQueriesInput() interface{}
	// Experimental.
	Fqn() *string
	GlossaryTerms() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextGlossaryTermsList
	GlossaryTermsInput() interface{}
	InternalValue() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContext
	SetInternalValue(val *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContext)
	LookerGoldenQueries() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextLookerGoldenQueriesList
	LookerGoldenQueriesInput() interface{}
	Options() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptionsOutputReference
	OptionsInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptions
	SchemaRelationships() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextSchemaRelationshipsList
	SchemaRelationshipsInput() interface{}
	SystemInstruction() *string
	SetSystemInstruction(val *string)
	SystemInstructionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserFunctions() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextUserFunctionsOutputReference
	UserFunctionsInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextUserFunctions
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
	PutDatasourceReferences(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferences)
	PutExampleQueries(value interface{})
	PutGlossaryTerms(value interface{})
	PutLookerGoldenQueries(value interface{})
	PutOptions(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptions)
	PutSchemaRelationships(value interface{})
	PutUserFunctions(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextUserFunctions)
	ResetExampleQueries()
	ResetGlossaryTerms()
	ResetLookerGoldenQueries()
	ResetOptions()
	ResetSchemaRelationships()
	ResetSystemInstruction()
	ResetUserFunctions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference
type jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) DatasourceReferences() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesOutputReference {
	var returns GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesOutputReference
	_jsii_.Get(
		j,
		"datasourceReferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) DatasourceReferencesInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferences {
	var returns *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferences
	_jsii_.Get(
		j,
		"datasourceReferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ExampleQueries() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextExampleQueriesList {
	var returns GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextExampleQueriesList
	_jsii_.Get(
		j,
		"exampleQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ExampleQueriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exampleQueriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GlossaryTerms() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextGlossaryTermsList {
	var returns GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextGlossaryTermsList
	_jsii_.Get(
		j,
		"glossaryTerms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GlossaryTermsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"glossaryTermsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) InternalValue() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContext {
	var returns *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) LookerGoldenQueries() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextLookerGoldenQueriesList {
	var returns GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextLookerGoldenQueriesList
	_jsii_.Get(
		j,
		"lookerGoldenQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) LookerGoldenQueriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lookerGoldenQueriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) Options() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptionsOutputReference {
	var returns GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptionsOutputReference
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) OptionsInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptions {
	var returns *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptions
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) SchemaRelationships() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextSchemaRelationshipsList {
	var returns GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextSchemaRelationshipsList
	_jsii_.Get(
		j,
		"schemaRelationships",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) SchemaRelationshipsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaRelationshipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) SystemInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) SystemInstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemInstructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) UserFunctions() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextUserFunctionsOutputReference {
	var returns GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextUserFunctionsOutputReference
	_jsii_.Get(
		j,
		"userFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) UserFunctionsInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextUserFunctions {
	var returns *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextUserFunctions
	_jsii_.Get(
		j,
		"userFunctionsInput",
		&returns,
	)
	return returns
}


func NewGoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGeminiDataAnalyticsDataAgent.GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference_Override(g GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGeminiDataAnalyticsDataAgent.GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference)SetInternalValue(val *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference)SetSystemInstruction(val *string) {
	if err := j.validateSetSystemInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemInstruction",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) PutDatasourceReferences(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferences) {
	if err := g.validatePutDatasourceReferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatasourceReferences",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) PutExampleQueries(value interface{}) {
	if err := g.validatePutExampleQueriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExampleQueries",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) PutGlossaryTerms(value interface{}) {
	if err := g.validatePutGlossaryTermsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGlossaryTerms",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) PutLookerGoldenQueries(value interface{}) {
	if err := g.validatePutLookerGoldenQueriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLookerGoldenQueries",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) PutOptions(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptions) {
	if err := g.validatePutOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) PutSchemaRelationships(value interface{}) {
	if err := g.validatePutSchemaRelationshipsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchemaRelationships",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) PutUserFunctions(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextUserFunctions) {
	if err := g.validatePutUserFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserFunctions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ResetExampleQueries() {
	_jsii_.InvokeVoid(
		g,
		"resetExampleQueries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ResetGlossaryTerms() {
	_jsii_.InvokeVoid(
		g,
		"resetGlossaryTerms",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ResetLookerGoldenQueries() {
	_jsii_.InvokeVoid(
		g,
		"resetLookerGoldenQueries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ResetOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ResetSchemaRelationships() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaRelationships",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ResetSystemInstruction() {
	_jsii_.InvokeVoid(
		g,
		"resetSystemInstruction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ResetUserFunctions() {
	_jsii_.InvokeVoid(
		g,
		"resetUserFunctions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

