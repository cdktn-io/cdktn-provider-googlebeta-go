// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package providerfunctions

import (
	"fmt"
)

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateLocationFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateNameFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateProjectFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateRegionFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateRegionFromZoneParameters(zone *string) error {
	if zone == nil {
		return fmt.Errorf("parameter zone is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleBetaProviderFunctions) validateZoneFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func validateNewGoogleBetaProviderFunctionsParameters(providerLocalName *string) error {
	if providerLocalName == nil {
		return fmt.Errorf("parameter providerLocalName is required, but nil was provided")
	}

	return nil
}

