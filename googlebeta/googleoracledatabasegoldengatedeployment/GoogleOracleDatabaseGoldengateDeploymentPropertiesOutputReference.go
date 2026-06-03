// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengatedeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googleoracledatabasegoldengatedeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference interface {
	cdktn.ComplexObject
	BackupSchedule() GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupScheduleOutputReference
	BackupScheduleInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupSchedule
	Category() *string
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
	CpuCoreCount() *float64
	SetCpuCoreCount(val *float64)
	CpuCoreCountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentBackupId() *string
	DeploymentDiagnosticData() GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticDataOutputReference
	DeploymentDiagnosticDataInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData
	DeploymentRole() *string
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
	DeploymentUrl() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnvironmentType() *string
	SetEnvironmentType(val *string)
	EnvironmentTypeInput() *string
	Fqdn() *string
	// Experimental.
	Fqn() *string
	Healthy() cdktn.IResolvable
	IngressIps() GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIpsList
	InternalValue() *GoogleOracleDatabaseGoldengateDeploymentProperties
	SetInternalValue(val *GoogleOracleDatabaseGoldengateDeploymentProperties)
	IsAutoScalingEnabled() interface{}
	SetIsAutoScalingEnabled(val interface{})
	IsAutoScalingEnabledInput() interface{}
	IsLatestVersion() cdktn.IResolvable
	IsPublic() cdktn.IResolvable
	IsStorageUtilizationLimitExceeded() cdktn.IResolvable
	LastBackupScheduleTime() *string
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	LifecycleDetails() *string
	LifecycleState() *string
	LifecycleSubState() *string
	LoadBalancerId() *string
	LoadBalancerSubnetId() *string
	Locks() GoogleOracleDatabaseGoldengateDeploymentPropertiesLocksList
	MaintenanceConfig() GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference
	MaintenanceConfigInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig
	MaintenanceWindow() GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindowOutputReference
	MaintenanceWindowInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow
	NextBackupScheduleTime() *string
	NextMaintenanceActionType() *string
	NextMaintenanceDescription() *string
	NextMaintenanceTime() *string
	NsgIds() *[]*string
	Ocid() *string
	OggData() GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference
	OggDataInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData
	OggVersionSupportEndTime() *string
	Placements() GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacementsList
	PrivateIpAddress() *string
	PublicIpAddress() *string
	RoleChangeTime() *string
	StorageUtilizationBytes() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateTime() *string
	UpgradeRequiredTime() *string
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
	PutBackupSchedule(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupSchedule)
	PutDeploymentDiagnosticData(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData)
	PutMaintenanceConfig(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig)
	PutMaintenanceWindow(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow)
	PutOggData(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData)
	ResetBackupSchedule()
	ResetCpuCoreCount()
	ResetDeploymentDiagnosticData()
	ResetDescription()
	ResetEnvironmentType()
	ResetIsAutoScalingEnabled()
	ResetLicenseModel()
	ResetMaintenanceConfig()
	ResetMaintenanceWindow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) BackupSchedule() GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupScheduleOutputReference {
	var returns GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupScheduleOutputReference
	_jsii_.Get(
		j,
		"backupSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) BackupScheduleInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupSchedule {
	var returns *GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupSchedule
	_jsii_.Get(
		j,
		"backupScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) CpuCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentBackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentBackupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentDiagnosticData() GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticDataOutputReference {
	var returns GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticDataOutputReference
	_jsii_.Get(
		j,
		"deploymentDiagnosticData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentDiagnosticDataInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData {
	var returns *GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData
	_jsii_.Get(
		j,
		"deploymentDiagnosticDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) EnvironmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) EnvironmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) Healthy() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"healthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) IngressIps() GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIpsList {
	var returns GoogleOracleDatabaseGoldengateDeploymentPropertiesIngressIpsList
	_jsii_.Get(
		j,
		"ingressIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateDeploymentProperties {
	var returns *GoogleOracleDatabaseGoldengateDeploymentProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsAutoScalingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsAutoScalingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsLatestVersion() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isLatestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsPublic() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsStorageUtilizationLimitExceeded() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isStorageUtilizationLimitExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) LastBackupScheduleTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastBackupScheduleTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) LifecycleSubState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleSubState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) LoadBalancerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) LoadBalancerSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) Locks() GoogleOracleDatabaseGoldengateDeploymentPropertiesLocksList {
	var returns GoogleOracleDatabaseGoldengateDeploymentPropertiesLocksList
	_jsii_.Get(
		j,
		"locks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) MaintenanceConfig() GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference {
	var returns GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference
	_jsii_.Get(
		j,
		"maintenanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) MaintenanceConfigInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig {
	var returns *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig
	_jsii_.Get(
		j,
		"maintenanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) MaintenanceWindow() GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindowOutputReference {
	var returns GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) MaintenanceWindowInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow {
	var returns *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) NextBackupScheduleTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextBackupScheduleTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) NextMaintenanceActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceActionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) NextMaintenanceDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) NextMaintenanceTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) NsgIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nsgIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) OggData() GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference {
	var returns GoogleOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference
	_jsii_.Get(
		j,
		"oggData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) OggDataInput() *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData {
	var returns *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData
	_jsii_.Get(
		j,
		"oggDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) OggVersionSupportEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oggVersionSupportEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) Placements() GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacementsList {
	var returns GoogleOracleDatabaseGoldengateDeploymentPropertiesPlacementsList
	_jsii_.Get(
		j,
		"placements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) RoleChangeTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleChangeTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) StorageUtilizationBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageUtilizationBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) UpgradeRequiredTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeRequiredTime",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference_Override(g GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateDeployment.GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetCpuCoreCount(val *float64) {
	if err := j.validateSetCpuCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetEnvironmentType(val *string) {
	if err := j.validateSetEnvironmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateDeploymentProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetIsAutoScalingEnabled(val interface{}) {
	if err := j.validateSetIsAutoScalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoScalingEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutBackupSchedule(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesBackupSchedule) {
	if err := g.validatePutBackupScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBackupSchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutDeploymentDiagnosticData(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData) {
	if err := g.validatePutDeploymentDiagnosticDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeploymentDiagnosticData",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutMaintenanceConfig(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig) {
	if err := g.validatePutMaintenanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutMaintenanceWindow(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow) {
	if err := g.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutOggData(value *GoogleOracleDatabaseGoldengateDeploymentPropertiesOggData) {
	if err := g.validatePutOggDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOggData",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetBackupSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetCpuCoreCount() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuCoreCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetDeploymentDiagnosticData() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentDiagnosticData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetEnvironmentType() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironmentType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetIsAutoScalingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIsAutoScalingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		g,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetMaintenanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateDeploymentPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

