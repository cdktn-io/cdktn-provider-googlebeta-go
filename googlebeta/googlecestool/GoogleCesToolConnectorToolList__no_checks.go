// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package googlecestool

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GoogleCesToolConnectorToolList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GoogleCesToolConnectorToolList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GoogleCesToolConnectorToolList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GoogleCesToolConnectorToolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GoogleCesToolConnectorToolList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GoogleCesToolConnectorToolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGoogleCesToolConnectorToolListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

