package netappvolume

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"azurerm.netappVolume.NetappVolume",
		reflect.TypeOf((*NetappVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountName", GoGetter: "AccountName"},
			_jsii_.MemberProperty{JsiiProperty: "accountNameInput", GoGetter: "AccountNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "azureVmwareDataStoreEnabled", GoGetter: "AzureVmwareDataStoreEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "azureVmwareDataStoreEnabledInput", GoGetter: "AzureVmwareDataStoreEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createFromSnapshotResourceId", GoGetter: "CreateFromSnapshotResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "createFromSnapshotResourceIdInput", GoGetter: "CreateFromSnapshotResourceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataProtectionReplication", GoGetter: "DataProtectionReplication"},
			_jsii_.MemberProperty{JsiiProperty: "dataProtectionReplicationInput", GoGetter: "DataProtectionReplicationInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataProtectionSnapshotPolicy", GoGetter: "DataProtectionSnapshotPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dataProtectionSnapshotPolicyInput", GoGetter: "DataProtectionSnapshotPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "exportPolicyRule", GoGetter: "ExportPolicyRule"},
			_jsii_.MemberProperty{JsiiProperty: "exportPolicyRuleInput", GoGetter: "ExportPolicyRuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "mountIpAddresses", GoGetter: "MountIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkFeatures", GoGetter: "NetworkFeatures"},
			_jsii_.MemberProperty{JsiiProperty: "networkFeaturesInput", GoGetter: "NetworkFeaturesInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "poolName", GoGetter: "PoolName"},
			_jsii_.MemberProperty{JsiiProperty: "poolNameInput", GoGetter: "PoolNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocols", GoGetter: "Protocols"},
			_jsii_.MemberProperty{JsiiProperty: "protocolsInput", GoGetter: "ProtocolsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDataProtectionReplication", GoMethod: "PutDataProtectionReplication"},
			_jsii_.MemberMethod{JsiiMethod: "putDataProtectionSnapshotPolicy", GoMethod: "PutDataProtectionSnapshotPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putExportPolicyRule", GoMethod: "PutExportPolicyRule"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureVmwareDataStoreEnabled", GoMethod: "ResetAzureVmwareDataStoreEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreateFromSnapshotResourceId", GoMethod: "ResetCreateFromSnapshotResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataProtectionReplication", GoMethod: "ResetDataProtectionReplication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataProtectionSnapshotPolicy", GoMethod: "ResetDataProtectionSnapshotPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetExportPolicyRule", GoMethod: "ResetExportPolicyRule"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkFeatures", GoMethod: "ResetNetworkFeatures"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtocols", GoMethod: "ResetProtocols"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityStyle", GoMethod: "ResetSecurityStyle"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnapshotDirectoryVisible", GoMethod: "ResetSnapshotDirectoryVisible"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetThroughputInMibps", GoMethod: "ResetThroughputInMibps"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetZone", GoMethod: "ResetZone"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityStyle", GoGetter: "SecurityStyle"},
			_jsii_.MemberProperty{JsiiProperty: "securityStyleInput", GoGetter: "SecurityStyleInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceLevel", GoGetter: "ServiceLevel"},
			_jsii_.MemberProperty{JsiiProperty: "serviceLevelInput", GoGetter: "ServiceLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotDirectoryVisible", GoGetter: "SnapshotDirectoryVisible"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotDirectoryVisibleInput", GoGetter: "SnapshotDirectoryVisibleInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageQuotaInGb", GoGetter: "StorageQuotaInGb"},
			_jsii_.MemberProperty{JsiiProperty: "storageQuotaInGbInput", GoGetter: "StorageQuotaInGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdInput", GoGetter: "SubnetIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "throughputInMibps", GoGetter: "ThroughputInMibps"},
			_jsii_.MemberProperty{JsiiProperty: "throughputInMibpsInput", GoGetter: "ThroughputInMibpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "volumePath", GoGetter: "VolumePath"},
			_jsii_.MemberProperty{JsiiProperty: "volumePathInput", GoGetter: "VolumePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "zone", GoGetter: "Zone"},
			_jsii_.MemberProperty{JsiiProperty: "zoneInput", GoGetter: "ZoneInput"},
		},
		func() interface{} {
			j := jsiiProxy_NetappVolume{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.netappVolume.NetappVolumeConfig",
		reflect.TypeOf((*NetappVolumeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"azurerm.netappVolume.NetappVolumeDataProtectionReplication",
		reflect.TypeOf((*NetappVolumeDataProtectionReplication)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.netappVolume.NetappVolumeDataProtectionReplicationOutputReference",
		reflect.TypeOf((*NetappVolumeDataProtectionReplicationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpointType", GoGetter: "EndpointType"},
			_jsii_.MemberProperty{JsiiProperty: "endpointTypeInput", GoGetter: "EndpointTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "remoteVolumeLocation", GoGetter: "RemoteVolumeLocation"},
			_jsii_.MemberProperty{JsiiProperty: "remoteVolumeLocationInput", GoGetter: "RemoteVolumeLocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "remoteVolumeResourceId", GoGetter: "RemoteVolumeResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "remoteVolumeResourceIdInput", GoGetter: "RemoteVolumeResourceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "replicationFrequency", GoGetter: "ReplicationFrequency"},
			_jsii_.MemberProperty{JsiiProperty: "replicationFrequencyInput", GoGetter: "ReplicationFrequencyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpointType", GoMethod: "ResetEndpointType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.netappVolume.NetappVolumeDataProtectionSnapshotPolicy",
		reflect.TypeOf((*NetappVolumeDataProtectionSnapshotPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.netappVolume.NetappVolumeDataProtectionSnapshotPolicyOutputReference",
		reflect.TypeOf((*NetappVolumeDataProtectionSnapshotPolicyOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotPolicyId", GoGetter: "SnapshotPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotPolicyIdInput", GoGetter: "SnapshotPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.netappVolume.NetappVolumeExportPolicyRule",
		reflect.TypeOf((*NetappVolumeExportPolicyRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.netappVolume.NetappVolumeExportPolicyRuleList",
		reflect.TypeOf((*NetappVolumeExportPolicyRuleList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_NetappVolumeExportPolicyRuleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.netappVolume.NetappVolumeExportPolicyRuleOutputReference",
		reflect.TypeOf((*NetappVolumeExportPolicyRuleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedClients", GoGetter: "AllowedClients"},
			_jsii_.MemberProperty{JsiiProperty: "allowedClientsInput", GoGetter: "AllowedClientsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "protocolsEnabled", GoGetter: "ProtocolsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "protocolsEnabledInput", GoGetter: "ProtocolsEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtocolsEnabled", GoMethod: "ResetProtocolsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRootAccessEnabled", GoMethod: "ResetRootAccessEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnixReadOnly", GoMethod: "ResetUnixReadOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnixReadWrite", GoMethod: "ResetUnixReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rootAccessEnabled", GoGetter: "RootAccessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "rootAccessEnabledInput", GoGetter: "RootAccessEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "ruleIndex", GoGetter: "RuleIndex"},
			_jsii_.MemberProperty{JsiiProperty: "ruleIndexInput", GoGetter: "RuleIndexInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unixReadOnly", GoGetter: "UnixReadOnly"},
			_jsii_.MemberProperty{JsiiProperty: "unixReadOnlyInput", GoGetter: "UnixReadOnlyInput"},
			_jsii_.MemberProperty{JsiiProperty: "unixReadWrite", GoGetter: "UnixReadWrite"},
			_jsii_.MemberProperty{JsiiProperty: "unixReadWriteInput", GoGetter: "UnixReadWriteInput"},
		},
		func() interface{} {
			j := jsiiProxy_NetappVolumeExportPolicyRuleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.netappVolume.NetappVolumeTimeouts",
		reflect.TypeOf((*NetappVolumeTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.netappVolume.NetappVolumeTimeoutsOutputReference",
		reflect.TypeOf((*NetappVolumeTimeoutsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_NetappVolumeTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
