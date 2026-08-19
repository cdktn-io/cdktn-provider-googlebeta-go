// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechronicleruledeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlechronicleruledeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference interface {
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
	EnsureEnrichmentCompleteness() interface{}
	SetEnsureEnrichmentCompleteness(val interface{})
	EnsureEnrichmentCompletenessInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleChronicleRuleDeploymentScheduleCustomizations
	SetInternalValue(val *GoogleChronicleRuleDeploymentScheduleCustomizations)
	LateArrivingDataAdjustment() *string
	SetLateArrivingDataAdjustment(val *string)
	LateArrivingDataAdjustmentInput() *string
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
	ResetEnsureEnrichmentCompleteness()
	ResetLateArrivingDataAdjustment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference
type jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) EnsureEnrichmentCompleteness() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ensureEnrichmentCompleteness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) EnsureEnrichmentCompletenessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ensureEnrichmentCompletenessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) InternalValue() *GoogleChronicleRuleDeploymentScheduleCustomizations {
	var returns *GoogleChronicleRuleDeploymentScheduleCustomizations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) LateArrivingDataAdjustment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lateArrivingDataAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) LateArrivingDataAdjustmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lateArrivingDataAdjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleRuleDeploymentScheduleCustomizationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleRuleDeployment.GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference_Override(g GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleRuleDeployment.GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference)SetEnsureEnrichmentCompleteness(val interface{}) {
	if err := j.validateSetEnsureEnrichmentCompletenessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ensureEnrichmentCompleteness",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference)SetInternalValue(val *GoogleChronicleRuleDeploymentScheduleCustomizations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference)SetLateArrivingDataAdjustment(val *string) {
	if err := j.validateSetLateArrivingDataAdjustmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lateArrivingDataAdjustment",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) ResetEnsureEnrichmentCompleteness() {
	_jsii_.InvokeVoid(
		g,
		"resetEnsureEnrichmentCompleteness",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) ResetLateArrivingDataAdjustment() {
	_jsii_.InvokeVoid(
		g,
		"resetLateArrivingDataAdjustment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleRuleDeploymentScheduleCustomizationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

