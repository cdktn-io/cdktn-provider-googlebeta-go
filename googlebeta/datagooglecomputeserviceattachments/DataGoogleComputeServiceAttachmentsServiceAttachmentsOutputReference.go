// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagooglecomputeserviceattachments

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/datagooglecomputeserviceattachments/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference interface {
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
	ConnectedEndpoints() DataGoogleComputeServiceAttachmentsServiceAttachmentsConnectedEndpointsList
	ConnectionPreference() *string
	ConsumerAcceptLists() DataGoogleComputeServiceAttachmentsServiceAttachmentsConsumerAcceptListsList
	ConsumerRejectLists() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeletionPolicy() *string
	Description() *string
	DomainNames() *[]*string
	EnableProxyProtocol() cdktn.IResolvable
	Fingerprint() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleComputeServiceAttachmentsServiceAttachments
	SetInternalValue(val *DataGoogleComputeServiceAttachmentsServiceAttachments)
	Name() *string
	NatIpsPerEndpoint() *float64
	NatSubnets() *[]*string
	Project() *string
	PropagatedConnectionLimit() *float64
	PscServiceAttachmentId() DataGoogleComputeServiceAttachmentsServiceAttachmentsPscServiceAttachmentIdList
	ReconcileConnections() cdktn.IResolvable
	Region() *string
	SelfLink() *string
	SendPropagatedConnectionLimitIfZero() cdktn.IResolvable
	ShowNatIps() cdktn.IResolvable
	TargetService() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TunnelingConfig() DataGoogleComputeServiceAttachmentsServiceAttachmentsTunnelingConfigList
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

// The jsii proxy struct for DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference
type jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ConnectedEndpoints() DataGoogleComputeServiceAttachmentsServiceAttachmentsConnectedEndpointsList {
	var returns DataGoogleComputeServiceAttachmentsServiceAttachmentsConnectedEndpointsList
	_jsii_.Get(
		j,
		"connectedEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ConnectionPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ConsumerAcceptLists() DataGoogleComputeServiceAttachmentsServiceAttachmentsConsumerAcceptListsList {
	var returns DataGoogleComputeServiceAttachmentsServiceAttachmentsConsumerAcceptListsList
	_jsii_.Get(
		j,
		"consumerAcceptLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ConsumerRejectLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerRejectLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) EnableProxyProtocol() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableProxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) InternalValue() *DataGoogleComputeServiceAttachmentsServiceAttachments {
	var returns *DataGoogleComputeServiceAttachmentsServiceAttachments
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) NatIpsPerEndpoint() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"natIpsPerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) NatSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) PropagatedConnectionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"propagatedConnectionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) PscServiceAttachmentId() DataGoogleComputeServiceAttachmentsServiceAttachmentsPscServiceAttachmentIdList {
	var returns DataGoogleComputeServiceAttachmentsServiceAttachmentsPscServiceAttachmentIdList
	_jsii_.Get(
		j,
		"pscServiceAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ReconcileConnections() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"reconcileConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) SendPropagatedConnectionLimitIfZero() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"sendPropagatedConnectionLimitIfZero",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ShowNatIps() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"showNatIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) TargetService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) TunnelingConfig() DataGoogleComputeServiceAttachmentsServiceAttachmentsTunnelingConfigList {
	var returns DataGoogleComputeServiceAttachmentsServiceAttachmentsTunnelingConfigList
	_jsii_.Get(
		j,
		"tunnelingConfig",
		&returns,
	)
	return returns
}


func NewDataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleComputeServiceAttachments.DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference_Override(d DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleComputeServiceAttachments.DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference)SetInternalValue(val *DataGoogleComputeServiceAttachmentsServiceAttachments) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeServiceAttachmentsServiceAttachmentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

