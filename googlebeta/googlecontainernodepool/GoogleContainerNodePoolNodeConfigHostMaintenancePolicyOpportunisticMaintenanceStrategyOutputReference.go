// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlecontainernodepool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategy
	SetInternalValue(val *GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategy)
	MaintenanceAvailabilityWindow() *string
	SetMaintenanceAvailabilityWindow(val *string)
	MaintenanceAvailabilityWindowInput() *string
	MinNodesPerPool() *float64
	SetMinNodesPerPool(val *float64)
	MinNodesPerPoolInput() *float64
	NodeIdleTimeWindow() *string
	SetNodeIdleTimeWindow(val *string)
	NodeIdleTimeWindowInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference
type jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) InternalValue() *GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategy {
	var returns *GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) MaintenanceAvailabilityWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceAvailabilityWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) MaintenanceAvailabilityWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceAvailabilityWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) MinNodesPerPool() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodesPerPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) MinNodesPerPoolInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodesPerPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) NodeIdleTimeWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeIdleTimeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) NodeIdleTimeWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeIdleTimeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference_Override(g GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference)SetInternalValue(val *GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference)SetMaintenanceAvailabilityWindow(val *string) {
	if err := j.validateSetMaintenanceAvailabilityWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceAvailabilityWindow",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference)SetMinNodesPerPool(val *float64) {
	if err := j.validateSetMinNodesPerPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodesPerPool",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference)SetNodeIdleTimeWindow(val *string) {
	if err := j.validateSetNodeIdleTimeWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeIdleTimeWindow",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

