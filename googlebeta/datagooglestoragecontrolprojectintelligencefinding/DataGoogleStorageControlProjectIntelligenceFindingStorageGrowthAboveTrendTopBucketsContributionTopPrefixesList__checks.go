// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package datagooglestoragecontrolprojectintelligencefinding

import (
	"fmt"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingStorageGrowthAboveTrendTopBucketsContributionTopPrefixesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingStorageGrowthAboveTrendTopBucketsContributionTopPrefixesList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingStorageGrowthAboveTrendTopBucketsContributionTopPrefixesList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingStorageGrowthAboveTrendTopBucketsContributionTopPrefixesList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingStorageGrowthAboveTrendTopBucketsContributionTopPrefixesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataGoogleStorageControlProjectIntelligenceFindingStorageGrowthAboveTrendTopBucketsContributionTopPrefixesList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataGoogleStorageControlProjectIntelligenceFindingStorageGrowthAboveTrendTopBucketsContributionTopPrefixesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

