// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package googlecolabschedule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateSetDomainParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigs:
		val := val.(*GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigs)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigs:
		val_ := val.(GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigs)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigs; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateSetTargetNetworkParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateSetTargetProjectParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGoogleColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

