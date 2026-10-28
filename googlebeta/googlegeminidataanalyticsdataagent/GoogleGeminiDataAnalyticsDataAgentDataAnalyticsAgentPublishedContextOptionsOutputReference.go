// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlegeminidataanalyticsdataagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference interface {
	cdktn.ComplexObject
	Analysis() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsAnalysisOutputReference
	AnalysisInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsAnalysis
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
	Datasource() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsDatasourceOutputReference
	DatasourceInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsDatasource
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions
	SetInternalValue(val *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions)
	Model() *string
	SetModel(val *string)
	ModelInput() *string
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
	PutAnalysis(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsAnalysis)
	PutDatasource(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsDatasource)
	ResetAnalysis()
	ResetDatasource()
	ResetModel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference
type jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) Analysis() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsAnalysisOutputReference {
	var returns GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsAnalysisOutputReference
	_jsii_.Get(
		j,
		"analysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) AnalysisInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsAnalysis {
	var returns *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsAnalysis
	_jsii_.Get(
		j,
		"analysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) Datasource() GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsDatasourceOutputReference {
	var returns GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsDatasourceOutputReference
	_jsii_.Get(
		j,
		"datasource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) DatasourceInput() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsDatasource {
	var returns *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsDatasource
	_jsii_.Get(
		j,
		"datasourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) InternalValue() *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions {
	var returns *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) Model() *string {
	var returns *string
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) ModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGeminiDataAnalyticsDataAgent.GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference_Override(g GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGeminiDataAnalyticsDataAgent.GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference)SetInternalValue(val *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference)SetModel(val *string) {
	if err := j.validateSetModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"model",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) PutAnalysis(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsAnalysis) {
	if err := g.validatePutAnalysisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnalysis",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) PutDatasource(value *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsDatasource) {
	if err := g.validatePutDatasourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatasource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) ResetAnalysis() {
	_jsii_.InvokeVoid(
		g,
		"resetAnalysis",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) ResetDatasource() {
	_jsii_.InvokeVoid(
		g,
		"resetDatasource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) ResetModel() {
	_jsii_.InvokeVoid(
		g,
		"resetModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

