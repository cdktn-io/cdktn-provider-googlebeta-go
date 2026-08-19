// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleagenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference interface {
	cdktn.ComplexObject
	AirQuality() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQualityOutputReference
	AirQualityInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality
	BureauLaborStatistics() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatisticsOutputReference
	BureauLaborStatisticsInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics
	Coindesk() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindeskOutputReference
	CoindeskInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Finnhub() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhubOutputReference
	FinnhubInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub
	// Experimental.
	Fqn() *string
	Fred() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFredOutputReference
	FredInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SecEdgar() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgarOutputReference
	SecEdgarInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar
	SelectionName() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TreasurySecuritiesAuctions() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctionsOutputReference
	TreasurySecuritiesAuctionsInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions
	Usda() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsdaOutputReference
	UsdaInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda
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
	PutAirQuality(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality)
	PutBureauLaborStatistics(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics)
	PutCoindesk(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk)
	PutFinnhub(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub)
	PutFred(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred)
	PutSecEdgar(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar)
	PutTreasurySecuritiesAuctions(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions)
	PutUsda(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda)
	ResetAirQuality()
	ResetBureauLaborStatistics()
	ResetCoindesk()
	ResetFinnhub()
	ResetFred()
	ResetSecEdgar()
	ResetTreasurySecuritiesAuctions()
	ResetUsda()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference
type jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) AirQuality() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQualityOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQualityOutputReference
	_jsii_.Get(
		j,
		"airQuality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) AirQualityInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality
	_jsii_.Get(
		j,
		"airQualityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) BureauLaborStatistics() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatisticsOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatisticsOutputReference
	_jsii_.Get(
		j,
		"bureauLaborStatistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) BureauLaborStatisticsInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics
	_jsii_.Get(
		j,
		"bureauLaborStatisticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Coindesk() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindeskOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindeskOutputReference
	_jsii_.Get(
		j,
		"coindesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) CoindeskInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk
	_jsii_.Get(
		j,
		"coindeskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Finnhub() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhubOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhubOutputReference
	_jsii_.Get(
		j,
		"finnhub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) FinnhubInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub
	_jsii_.Get(
		j,
		"finnhubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Fred() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFredOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFredOutputReference
	_jsii_.Get(
		j,
		"fred",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) FredInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred
	_jsii_.Get(
		j,
		"fredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) SecEdgar() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgarOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgarOutputReference
	_jsii_.Get(
		j,
		"secEdgar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) SecEdgarInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar
	_jsii_.Get(
		j,
		"secEdgarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) SelectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) TreasurySecuritiesAuctions() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctionsOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctionsOutputReference
	_jsii_.Get(
		j,
		"treasurySecuritiesAuctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) TreasurySecuritiesAuctionsInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions
	_jsii_.Get(
		j,
		"treasurySecuritiesAuctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Usda() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsdaOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsdaOutputReference
	_jsii_.Get(
		j,
		"usda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) UsdaInput() *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda
	_jsii_.Get(
		j,
		"usdaInput",
		&returns,
	)
	return returns
}


func NewGoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference_Override(g GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutAirQuality(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality) {
	if err := g.validatePutAirQualityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAirQuality",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutBureauLaborStatistics(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics) {
	if err := g.validatePutBureauLaborStatisticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBureauLaborStatistics",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutCoindesk(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk) {
	if err := g.validatePutCoindeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCoindesk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutFinnhub(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub) {
	if err := g.validatePutFinnhubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFinnhub",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutFred(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred) {
	if err := g.validatePutFredParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFred",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutSecEdgar(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar) {
	if err := g.validatePutSecEdgarParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecEdgar",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutTreasurySecuritiesAuctions(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions) {
	if err := g.validatePutTreasurySecuritiesAuctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTreasurySecuritiesAuctions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutUsda(value *GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda) {
	if err := g.validatePutUsdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUsda",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetAirQuality() {
	_jsii_.InvokeVoid(
		g,
		"resetAirQuality",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetBureauLaborStatistics() {
	_jsii_.InvokeVoid(
		g,
		"resetBureauLaborStatistics",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetCoindesk() {
	_jsii_.InvokeVoid(
		g,
		"resetCoindesk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetFinnhub() {
	_jsii_.InvokeVoid(
		g,
		"resetFinnhub",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetFred() {
	_jsii_.InvokeVoid(
		g,
		"resetFred",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetSecEdgar() {
	_jsii_.InvokeVoid(
		g,
		"resetSecEdgar",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetTreasurySecuritiesAuctions() {
	_jsii_.InvokeVoid(
		g,
		"resetTreasurySecuritiesAuctions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetUsda() {
	_jsii_.InvokeVoid(
		g,
		"resetUsda",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

