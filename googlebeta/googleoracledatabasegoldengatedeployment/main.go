// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengatedeployment

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeployment",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createTime", GoGetter: "CreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "deletionPolicy", GoGetter: "DeletionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deletionPolicyInput", GoGetter: "DeletionPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveLabels", GoGetter: "EffectiveLabels"},
			_jsii_.MemberProperty{JsiiProperty: "entitlementId", GoGetter: "EntitlementId"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "gcpOracleZone", GoGetter: "GcpOracleZone"},
			_jsii_.MemberProperty{JsiiProperty: "gcpOracleZoneInput", GoGetter: "GcpOracleZoneInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "goldengateDeploymentId", GoGetter: "GoldengateDeploymentId"},
			_jsii_.MemberProperty{JsiiProperty: "goldengateDeploymentIdInput", GoGetter: "GoldengateDeploymentIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ociUrl", GoGetter: "OciUrl"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetwork", GoGetter: "OdbNetwork"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkInput", GoGetter: "OdbNetworkInput"},
			_jsii_.MemberProperty{JsiiProperty: "odbSubnet", GoGetter: "OdbSubnet"},
			_jsii_.MemberProperty{JsiiProperty: "odbSubnetInput", GoGetter: "OdbSubnetInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putProperties", GoMethod: "PutProperties"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeletionPolicy", GoMethod: "ResetDeletionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetGcpOracleZone", GoMethod: "ResetGcpOracleZone"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetOdbNetwork", GoMethod: "ResetOdbNetwork"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformLabels", GoGetter: "TerraformLabels"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentConfig",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentProperties",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentProperties)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupSchedule",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupSchedule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupScheduleOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupScheduleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backupScheduledTime", GoGetter: "BackupScheduledTime"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "compartmentId", GoGetter: "CompartmentId"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyBackupScheduled", GoGetter: "FrequencyBackupScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "metadataOnly", GoGetter: "MetadataOnly"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupScheduleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticDataOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticDataOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diagnosticEndTime", GoGetter: "DiagnosticEndTime"},
			_jsii_.MemberProperty{JsiiProperty: "diagnosticStartTime", GoGetter: "DiagnosticStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "diagnosticState", GoGetter: "DiagnosticState"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "object", GoGetter: "Object"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticDataOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIps",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIpsList",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIpsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIpsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIpsOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIpsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ingressIpAddress", GoGetter: "IngressIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIpsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesLocks",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesLocks)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesLocksList",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesLocksList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesLocksList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesLocksOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesLocksOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "compartmentId", GoGetter: "CompartmentId"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createTime", GoGetter: "CreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "message", GoGetter: "Message"},
			_jsii_.MemberProperty{JsiiProperty: "relatedResourceId", GoGetter: "RelatedResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesLocksOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bundleReleaseUpgradePeriodDays", GoGetter: "BundleReleaseUpgradePeriodDays"},
			_jsii_.MemberProperty{JsiiProperty: "bundleReleaseUpgradePeriodDaysInput", GoGetter: "BundleReleaseUpgradePeriodDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "interimReleaseUpgradePeriodDays", GoGetter: "InterimReleaseUpgradePeriodDays"},
			_jsii_.MemberProperty{JsiiProperty: "interimReleaseUpgradePeriodDaysInput", GoGetter: "InterimReleaseUpgradePeriodDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isInterimReleaseAutoUpgradeEnabled", GoGetter: "IsInterimReleaseAutoUpgradeEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isInterimReleaseAutoUpgradeEnabledInput", GoGetter: "IsInterimReleaseAutoUpgradeEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "majorReleaseUpgradePeriodDays", GoGetter: "MajorReleaseUpgradePeriodDays"},
			_jsii_.MemberProperty{JsiiProperty: "majorReleaseUpgradePeriodDaysInput", GoGetter: "MajorReleaseUpgradePeriodDaysInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBundleReleaseUpgradePeriodDays", GoMethod: "ResetBundleReleaseUpgradePeriodDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetInterimReleaseUpgradePeriodDays", GoMethod: "ResetInterimReleaseUpgradePeriodDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsInterimReleaseAutoUpgradeEnabled", GoMethod: "ResetIsInterimReleaseAutoUpgradeEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMajorReleaseUpgradePeriodDays", GoMethod: "ResetMajorReleaseUpgradePeriodDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityPatchUpgradePeriodDays", GoMethod: "ResetSecurityPatchUpgradePeriodDays"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "securityPatchUpgradePeriodDays", GoGetter: "SecurityPatchUpgradePeriodDays"},
			_jsii_.MemberProperty{JsiiProperty: "securityPatchUpgradePeriodDaysInput", GoGetter: "SecurityPatchUpgradePeriodDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindowOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindowOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "day", GoGetter: "Day"},
			_jsii_.MemberProperty{JsiiProperty: "dayInput", GoGetter: "DayInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startHour", GoGetter: "StartHour"},
			_jsii_.MemberProperty{JsiiProperty: "startHourInput", GoGetter: "StartHourInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindowOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMappingOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMappingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "administratorGroupId", GoGetter: "AdministratorGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "operatorGroupId", GoGetter: "OperatorGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupId", GoGetter: "SecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupId", GoGetter: "UserGroupId"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMappingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adminPassword", GoGetter: "AdminPassword"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordInput", GoGetter: "AdminPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordSecretVersion", GoGetter: "AdminPasswordSecretVersion"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordSecretVersionInput", GoGetter: "AdminPasswordSecretVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "adminUsername", GoGetter: "AdminUsername"},
			_jsii_.MemberProperty{JsiiProperty: "adminUsernameInput", GoGetter: "AdminUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "credentialStore", GoGetter: "CredentialStore"},
			_jsii_.MemberProperty{JsiiProperty: "deployment", GoGetter: "Deployment"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentInput", GoGetter: "DeploymentInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupRolesMapping", GoGetter: "GroupRolesMapping"},
			_jsii_.MemberProperty{JsiiProperty: "groupRolesMappingInput", GoGetter: "GroupRolesMappingInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityDomainId", GoGetter: "IdentityDomainId"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "oggVersion", GoGetter: "OggVersion"},
			_jsii_.MemberProperty{JsiiProperty: "oggVersionInput", GoGetter: "OggVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordSecretId", GoGetter: "PasswordSecretId"},
			_jsii_.MemberMethod{JsiiMethod: "putGroupRolesMapping", GoMethod: "PutGroupRolesMapping"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminPassword", GoMethod: "ResetAdminPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminPasswordSecretVersion", GoMethod: "ResetAdminPasswordSecretVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupRolesMapping", GoMethod: "ResetGroupRolesMapping"},
			_jsii_.MemberMethod{JsiiMethod: "resetOggVersion", GoMethod: "ResetOggVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backupSchedule", GoGetter: "BackupSchedule"},
			_jsii_.MemberProperty{JsiiProperty: "backupScheduleInput", GoGetter: "BackupScheduleInput"},
			_jsii_.MemberProperty{JsiiProperty: "category", GoGetter: "Category"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCount", GoGetter: "CpuCoreCount"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCountInput", GoGetter: "CpuCoreCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentBackupId", GoGetter: "DeploymentBackupId"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentDiagnosticData", GoGetter: "DeploymentDiagnosticData"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentDiagnosticDataInput", GoGetter: "DeploymentDiagnosticDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentRole", GoGetter: "DeploymentRole"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentType", GoGetter: "DeploymentType"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentTypeInput", GoGetter: "DeploymentTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentUrl", GoGetter: "DeploymentUrl"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "environmentType", GoGetter: "EnvironmentType"},
			_jsii_.MemberProperty{JsiiProperty: "environmentTypeInput", GoGetter: "EnvironmentTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqdn", GoGetter: "Fqdn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "healthy", GoGetter: "Healthy"},
			_jsii_.MemberProperty{JsiiProperty: "ingressIps", GoGetter: "IngressIps"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isAutoScalingEnabled", GoGetter: "IsAutoScalingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isAutoScalingEnabledInput", GoGetter: "IsAutoScalingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "isLatestVersion", GoGetter: "IsLatestVersion"},
			_jsii_.MemberProperty{JsiiProperty: "isPublic", GoGetter: "IsPublic"},
			_jsii_.MemberProperty{JsiiProperty: "isStorageUtilizationLimitExceeded", GoGetter: "IsStorageUtilizationLimitExceeded"},
			_jsii_.MemberProperty{JsiiProperty: "lastBackupScheduleTime", GoGetter: "LastBackupScheduleTime"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModel", GoGetter: "LicenseModel"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModelInput", GoGetter: "LicenseModelInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleDetails", GoGetter: "LifecycleDetails"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleState", GoGetter: "LifecycleState"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleSubState", GoGetter: "LifecycleSubState"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerId", GoGetter: "LoadBalancerId"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerSubnetId", GoGetter: "LoadBalancerSubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "locks", GoGetter: "Locks"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceConfig", GoGetter: "MaintenanceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceConfigInput", GoGetter: "MaintenanceConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindow", GoGetter: "MaintenanceWindow"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowInput", GoGetter: "MaintenanceWindowInput"},
			_jsii_.MemberProperty{JsiiProperty: "nextBackupScheduleTime", GoGetter: "NextBackupScheduleTime"},
			_jsii_.MemberProperty{JsiiProperty: "nextMaintenanceActionType", GoGetter: "NextMaintenanceActionType"},
			_jsii_.MemberProperty{JsiiProperty: "nextMaintenanceDescription", GoGetter: "NextMaintenanceDescription"},
			_jsii_.MemberProperty{JsiiProperty: "nextMaintenanceTime", GoGetter: "NextMaintenanceTime"},
			_jsii_.MemberProperty{JsiiProperty: "nsgIds", GoGetter: "NsgIds"},
			_jsii_.MemberProperty{JsiiProperty: "ocid", GoGetter: "Ocid"},
			_jsii_.MemberProperty{JsiiProperty: "oggData", GoGetter: "OggData"},
			_jsii_.MemberProperty{JsiiProperty: "oggDataInput", GoGetter: "OggDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "oggVersionSupportEndTime", GoGetter: "OggVersionSupportEndTime"},
			_jsii_.MemberProperty{JsiiProperty: "placements", GoGetter: "Placements"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddress", GoGetter: "PrivateIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddress", GoGetter: "PublicIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "putBackupSchedule", GoMethod: "PutBackupSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "putDeploymentDiagnosticData", GoMethod: "PutDeploymentDiagnosticData"},
			_jsii_.MemberMethod{JsiiMethod: "putMaintenanceConfig", GoMethod: "PutMaintenanceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putMaintenanceWindow", GoMethod: "PutMaintenanceWindow"},
			_jsii_.MemberMethod{JsiiMethod: "putOggData", GoMethod: "PutOggData"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackupSchedule", GoMethod: "ResetBackupSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuCoreCount", GoMethod: "ResetCpuCoreCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentDiagnosticData", GoMethod: "ResetDeploymentDiagnosticData"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironmentType", GoMethod: "ResetEnvironmentType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsAutoScalingEnabled", GoMethod: "ResetIsAutoScalingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseModel", GoMethod: "ResetLicenseModel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceConfig", GoMethod: "ResetMaintenanceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceWindow", GoMethod: "ResetMaintenanceWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleChangeTime", GoGetter: "RoleChangeTime"},
			_jsii_.MemberProperty{JsiiProperty: "storageUtilizationBytes", GoGetter: "StorageUtilizationBytes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updateTime", GoGetter: "UpdateTime"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeRequiredTime", GoGetter: "UpgradeRequiredTime"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacements",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacements)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacementsList",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacementsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacementsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacementsOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacementsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityDomain", GoGetter: "AvailabilityDomain"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "faultDomain", GoGetter: "FaultDomain"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacementsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentTimeouts",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentTimeoutsOutputReference",
		reflect.TypeOf((*GoogleOracleDatabaseGoldengateDeploymentTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
