// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlecontainercluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference interface {
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
	DelayUntil() GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntilOutputReference
	DelayUntilInput() *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindow
	SetInternalValue(val *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindow)
	Recurrence() *string
	SetRecurrence(val *string)
	RecurrenceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WindowDuration() *string
	SetWindowDuration(val *string)
	WindowDurationInput() *string
	WindowStartTime() GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTimeOutputReference
	WindowStartTimeInput() *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime
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
	PutDelayUntil(value *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil)
	PutWindowStartTime(value *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime)
	ResetDelayUntil()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference
type jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) DelayUntil() GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntilOutputReference {
	var returns GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntilOutputReference
	_jsii_.Get(
		j,
		"delayUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) DelayUntilInput() *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil {
	var returns *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil
	_jsii_.Get(
		j,
		"delayUntilInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) InternalValue() *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindow {
	var returns *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) Recurrence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) RecurrenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) WindowDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) WindowDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) WindowStartTime() GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTimeOutputReference {
	var returns GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTimeOutputReference
	_jsii_.Get(
		j,
		"windowStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) WindowStartTimeInput() *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime {
	var returns *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime
	_jsii_.Get(
		j,
		"windowStartTimeInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference_Override(g GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetInternalValue(val *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetRecurrence(val *string) {
	if err := j.validateSetRecurrenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrence",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference)SetWindowDuration(val *string) {
	if err := j.validateSetWindowDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowDuration",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) PutDelayUntil(value *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil) {
	if err := g.validatePutDelayUntilParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDelayUntil",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) PutWindowStartTime(value *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime) {
	if err := g.validatePutWindowStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWindowStartTime",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ResetDelayUntil() {
	_jsii_.InvokeVoid(
		g,
		"resetDelayUntil",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

