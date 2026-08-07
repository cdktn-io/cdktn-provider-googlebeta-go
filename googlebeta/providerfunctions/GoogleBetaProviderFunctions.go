// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"
)

// Provider-defined functions of the google-beta provider.
type GoogleBetaProviderFunctions interface {
	// Takes a single string argument, which should be a resource id, self link, or OP style resource name.
	//
	// This function will either return the location name from the input string or raise an error due to no location being present in the string. The function uses the presence of "locations/{{location}}/" in the input string to identify the location name, e.g. when the function is passed the id "projects/my-project/locations/us-central1/services/my-service" as an argument it will return "us-central1".
	LocationFromId(id *string) *string
	// Takes a single string argument, which should be a resource's id, resource URI, self link, or full resource name.
	//
	// This function will return the short-form name of a resource from the input string, or raise an error due to a problem with the input string. The function returns the final element in the input string as the resource's name, e.g. when the function is passed the id "projects/my-project/zones/us-central1-c/instances/my-instance" as an argument it will return "my-instance".
	NameFromId(id *string) *string
	// Takes a single string argument, which should be a resource's id, resource URI, self link, or full resource name.
	//
	// This function will either return the project name from the input string or raise an error due to no project being present in the string. The function uses the presence of "projects/{{project}}/" in the input string to identify the project name, e.g. when the function is passed the id "projects/my-project/zones/us-central1-c/instances/my-instance" as an argument it will return "my-project".
	ProjectFromId(id *string) *string
	// Takes a single string argument, which should be a resource id, self link, or OP style resource name.
	//
	// This function will either return the region name from the input string or raise an error due to no region being present in the string. The function uses the presence of "regions/{{region}}/" in the input string to identify the region name, e.g. when the function is passed the id "projects/my-project/regions/us-central1/subnetworks/my-subnetwork" as an argument it will return "us-central1".
	RegionFromId(id *string) *string
	// Takes a single string argument, which should be a resource's zone.
	RegionFromZone(zone *string) *string
	// Takes a single string argument, which should be an id or self link of a resource.
	//
	// This function will either return the zone name from the input string or raise an error due to no zone being present in the string. The function uses the presence of "zones/{{zone}}/" in the input string to identify the zone name, e.g. when the function is passed the id "projects/my-project/zones/us-central1-c/instances/my-instance" as an argument it will return "us-central1-c".
	ZoneFromId(id *string) *string
}

// The jsii proxy struct for GoogleBetaProviderFunctions
type jsiiProxy_GoogleBetaProviderFunctions struct {
	_ byte // padding
}

func NewGoogleBetaProviderFunctions(providerLocalName *string) GoogleBetaProviderFunctions {
	_init_.Initialize()

	if err := validateNewGoogleBetaProviderFunctionsParameters(providerLocalName); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBetaProviderFunctions{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.providerFunctions.GoogleBetaProviderFunctions",
		[]interface{}{providerLocalName},
		&j,
	)

	return &j
}

func NewGoogleBetaProviderFunctions_Override(g GoogleBetaProviderFunctions, providerLocalName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.providerFunctions.GoogleBetaProviderFunctions",
		[]interface{}{providerLocalName},
		g,
	)
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) LocationFromId(id *string) *string {
	if err := g.validateLocationFromIdParameters(id); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"locationFromId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) NameFromId(id *string) *string {
	if err := g.validateNameFromIdParameters(id); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"nameFromId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) ProjectFromId(id *string) *string {
	if err := g.validateProjectFromIdParameters(id); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"projectFromId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) RegionFromId(id *string) *string {
	if err := g.validateRegionFromIdParameters(id); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"regionFromId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) RegionFromZone(zone *string) *string {
	if err := g.validateRegionFromZoneParameters(zone); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"regionFromZone",
		[]interface{}{zone},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) ZoneFromId(id *string) *string {
	if err := g.validateZoneFromIdParameters(id); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"zoneFromId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

