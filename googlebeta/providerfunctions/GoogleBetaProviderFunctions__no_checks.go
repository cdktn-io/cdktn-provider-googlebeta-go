// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package providerfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateLocationFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateNameFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateProjectFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateRegionFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateRegionFromZoneParameters(zone *string) error {
	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateZoneFromIdParameters(id *string) error {
	return nil
}

func validateNewGoogleBetaProviderFunctionsParameters(providerLocalName *string) error {
	return nil
}

