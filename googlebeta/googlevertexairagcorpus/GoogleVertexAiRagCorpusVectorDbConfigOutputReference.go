// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagcorpus

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlevertexairagcorpus/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiRagCorpusVectorDbConfigOutputReference interface {
	cdktn.ComplexObject
	ApiAuth() GoogleVertexAiRagCorpusVectorDbConfigApiAuthOutputReference
	ApiAuthInput() *GoogleVertexAiRagCorpusVectorDbConfigApiAuth
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
	InternalValue() *GoogleVertexAiRagCorpusVectorDbConfig
	SetInternalValue(val *GoogleVertexAiRagCorpusVectorDbConfig)
	Pinecone() GoogleVertexAiRagCorpusVectorDbConfigPineconeOutputReference
	PineconeInput() *GoogleVertexAiRagCorpusVectorDbConfigPinecone
	RagEmbeddingModelConfig() GoogleVertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfigOutputReference
	RagEmbeddingModelConfigInput() *GoogleVertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig
	RagManagedDb() GoogleVertexAiRagCorpusVectorDbConfigRagManagedDbOutputReference
	RagManagedDbInput() *GoogleVertexAiRagCorpusVectorDbConfigRagManagedDb
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VertexVectorSearch() GoogleVertexAiRagCorpusVectorDbConfigVertexVectorSearchOutputReference
	VertexVectorSearchInput() *GoogleVertexAiRagCorpusVectorDbConfigVertexVectorSearch
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
	PutApiAuth(value *GoogleVertexAiRagCorpusVectorDbConfigApiAuth)
	PutPinecone(value *GoogleVertexAiRagCorpusVectorDbConfigPinecone)
	PutRagEmbeddingModelConfig(value *GoogleVertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig)
	PutRagManagedDb(value *GoogleVertexAiRagCorpusVectorDbConfigRagManagedDb)
	PutVertexVectorSearch(value *GoogleVertexAiRagCorpusVectorDbConfigVertexVectorSearch)
	ResetApiAuth()
	ResetPinecone()
	ResetRagEmbeddingModelConfig()
	ResetRagManagedDb()
	ResetVertexVectorSearch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiRagCorpusVectorDbConfigOutputReference
type jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ApiAuth() GoogleVertexAiRagCorpusVectorDbConfigApiAuthOutputReference {
	var returns GoogleVertexAiRagCorpusVectorDbConfigApiAuthOutputReference
	_jsii_.Get(
		j,
		"apiAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ApiAuthInput() *GoogleVertexAiRagCorpusVectorDbConfigApiAuth {
	var returns *GoogleVertexAiRagCorpusVectorDbConfigApiAuth
	_jsii_.Get(
		j,
		"apiAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) InternalValue() *GoogleVertexAiRagCorpusVectorDbConfig {
	var returns *GoogleVertexAiRagCorpusVectorDbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) Pinecone() GoogleVertexAiRagCorpusVectorDbConfigPineconeOutputReference {
	var returns GoogleVertexAiRagCorpusVectorDbConfigPineconeOutputReference
	_jsii_.Get(
		j,
		"pinecone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) PineconeInput() *GoogleVertexAiRagCorpusVectorDbConfigPinecone {
	var returns *GoogleVertexAiRagCorpusVectorDbConfigPinecone
	_jsii_.Get(
		j,
		"pineconeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) RagEmbeddingModelConfig() GoogleVertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfigOutputReference {
	var returns GoogleVertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfigOutputReference
	_jsii_.Get(
		j,
		"ragEmbeddingModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) RagEmbeddingModelConfigInput() *GoogleVertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig {
	var returns *GoogleVertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig
	_jsii_.Get(
		j,
		"ragEmbeddingModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) RagManagedDb() GoogleVertexAiRagCorpusVectorDbConfigRagManagedDbOutputReference {
	var returns GoogleVertexAiRagCorpusVectorDbConfigRagManagedDbOutputReference
	_jsii_.Get(
		j,
		"ragManagedDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) RagManagedDbInput() *GoogleVertexAiRagCorpusVectorDbConfigRagManagedDb {
	var returns *GoogleVertexAiRagCorpusVectorDbConfigRagManagedDb
	_jsii_.Get(
		j,
		"ragManagedDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) VertexVectorSearch() GoogleVertexAiRagCorpusVectorDbConfigVertexVectorSearchOutputReference {
	var returns GoogleVertexAiRagCorpusVectorDbConfigVertexVectorSearchOutputReference
	_jsii_.Get(
		j,
		"vertexVectorSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) VertexVectorSearchInput() *GoogleVertexAiRagCorpusVectorDbConfigVertexVectorSearch {
	var returns *GoogleVertexAiRagCorpusVectorDbConfigVertexVectorSearch
	_jsii_.Get(
		j,
		"vertexVectorSearchInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiRagCorpusVectorDbConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiRagCorpusVectorDbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiRagCorpusVectorDbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiRagCorpus.GoogleVertexAiRagCorpusVectorDbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiRagCorpusVectorDbConfigOutputReference_Override(g GoogleVertexAiRagCorpusVectorDbConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiRagCorpus.GoogleVertexAiRagCorpusVectorDbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference)SetInternalValue(val *GoogleVertexAiRagCorpusVectorDbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) PutApiAuth(value *GoogleVertexAiRagCorpusVectorDbConfigApiAuth) {
	if err := g.validatePutApiAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiAuth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) PutPinecone(value *GoogleVertexAiRagCorpusVectorDbConfigPinecone) {
	if err := g.validatePutPineconeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPinecone",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) PutRagEmbeddingModelConfig(value *GoogleVertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig) {
	if err := g.validatePutRagEmbeddingModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRagEmbeddingModelConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) PutRagManagedDb(value *GoogleVertexAiRagCorpusVectorDbConfigRagManagedDb) {
	if err := g.validatePutRagManagedDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRagManagedDb",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) PutVertexVectorSearch(value *GoogleVertexAiRagCorpusVectorDbConfigVertexVectorSearch) {
	if err := g.validatePutVertexVectorSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVertexVectorSearch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ResetApiAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetApiAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ResetPinecone() {
	_jsii_.InvokeVoid(
		g,
		"resetPinecone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ResetRagEmbeddingModelConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRagEmbeddingModelConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ResetRagManagedDb() {
	_jsii_.InvokeVoid(
		g,
		"resetRagManagedDb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ResetVertexVectorSearch() {
	_jsii_.InvokeVoid(
		g,
		"resetVertexVectorSearch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiRagCorpusVectorDbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

