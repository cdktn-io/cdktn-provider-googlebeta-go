// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleagenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference interface {
	cdktn.ComplexObject
	BigqueryDataset() *string
	SetBigqueryDataset(val *string)
	BigqueryDatasetInput() *string
	BigqueryTable() *string
	SetBigqueryTable(val *string)
	BigqueryTableInput() *string
	ColumnDescriptions() *map[string]*string
	SetColumnDescriptions(val *map[string]*string)
	ColumnDescriptionsInput() *map[string]*string
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
	InternalValue() *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource
	SetInternalValue(val *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource)
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
	ResetBigqueryDataset()
	ResetBigqueryTable()
	ResetColumnDescriptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference
type jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) BigqueryDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) BigqueryDatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) BigqueryTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) BigqueryTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) ColumnDescriptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"columnDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) ColumnDescriptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"columnDescriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) InternalValue() *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference_Override(g GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference)SetBigqueryDataset(val *string) {
	if err := j.validateSetBigqueryDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bigqueryDataset",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference)SetBigqueryTable(val *string) {
	if err := j.validateSetBigqueryTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bigqueryTable",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference)SetColumnDescriptions(val *map[string]*string) {
	if err := j.validateSetColumnDescriptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnDescriptions",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference)SetInternalValue(val *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) ResetBigqueryDataset() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDataset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) ResetBigqueryTable() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) ResetColumnDescriptions() {
	_jsii_.InvokeVoid(
		g,
		"resetColumnDescriptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

