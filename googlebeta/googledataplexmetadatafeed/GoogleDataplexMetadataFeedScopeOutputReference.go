// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataplexmetadatafeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googledataplexmetadatafeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDataplexMetadataFeedScopeOutputReference interface {
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
	EntryGroups() *[]*string
	SetEntryGroups(val *[]*string)
	EntryGroupsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataplexMetadataFeedScope
	SetInternalValue(val *GoogleDataplexMetadataFeedScope)
	OrganizationLevel() interface{}
	SetOrganizationLevel(val interface{})
	OrganizationLevelInput() interface{}
	Projects() *[]*string
	SetProjects(val *[]*string)
	ProjectsInput() *[]*string
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
	ResetEntryGroups()
	ResetOrganizationLevel()
	ResetProjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexMetadataFeedScopeOutputReference
type jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) EntryGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) EntryGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) InternalValue() *GoogleDataplexMetadataFeedScope {
	var returns *GoogleDataplexMetadataFeedScope
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) OrganizationLevel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) OrganizationLevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) Projects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) ProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataplexMetadataFeedScopeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataplexMetadataFeedScopeOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexMetadataFeedScopeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataplexMetadataFeed.GoogleDataplexMetadataFeedScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexMetadataFeedScopeOutputReference_Override(g GoogleDataplexMetadataFeedScopeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataplexMetadataFeed.GoogleDataplexMetadataFeedScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference)SetEntryGroups(val *[]*string) {
	if err := j.validateSetEntryGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryGroups",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference)SetInternalValue(val *GoogleDataplexMetadataFeedScope) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference)SetOrganizationLevel(val interface{}) {
	if err := j.validateSetOrganizationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference)SetProjects(val *[]*string) {
	if err := j.validateSetProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projects",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) ResetEntryGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetEntryGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) ResetOrganizationLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetOrganizationLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) ResetProjects() {
	_jsii_.InvokeVoid(
		g,
		"resetProjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataplexMetadataFeedScopeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

