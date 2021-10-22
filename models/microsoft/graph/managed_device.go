package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedDevice struct {
    Entity
    activationLockBypassCode *string;
    androidSecurityPatchLevel *string;
    azureADDeviceId *string;
    azureADRegistered *bool;
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    complianceState *ComplianceState;
    configurationManagerClientEnabledFeatures *ConfigurationManagerClientEnabledFeatures;
    deviceActionResults []DeviceActionResult;
    deviceCategory *DeviceCategory;
    deviceCategoryDisplayName *string;
    deviceCompliancePolicyStates []DeviceCompliancePolicyState;
    deviceConfigurationStates []DeviceConfigurationState;
    deviceEnrollmentType *DeviceEnrollmentType;
    deviceHealthAttestationState *DeviceHealthAttestationState;
    deviceName *string;
    deviceRegistrationState *DeviceRegistrationState;
    easActivated *bool;
    easActivationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    easDeviceId *string;
    emailAddress *string;
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    ethernetMacAddress *string;
    exchangeAccessState *DeviceManagementExchangeAccessState;
    exchangeAccessStateReason *DeviceManagementExchangeAccessStateReason;
    exchangeLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    freeStorageSpaceInBytes *int64;
    iccid *string;
    imei *string;
    isEncrypted *bool;
    isSupervised *bool;
    jailBroken *string;
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    managedDeviceName *string;
    managedDeviceOwnerType *ManagedDeviceOwnerType;
    managementAgent *ManagementAgentType;
    manufacturer *string;
    meid *string;
    model *string;
    notes *string;
    operatingSystem *string;
    osVersion *string;
    partnerReportedThreatState *ManagedDevicePartnerReportedHealthState;
    phoneNumber *string;
    physicalMemoryInBytes *int64;
    remoteAssistanceSessionErrorDetails *string;
    remoteAssistanceSessionUrl *string;
    serialNumber *string;
    subscriberCarrier *string;
    totalStorageSpaceInBytes *int64;
    udid *string;
    userDisplayName *string;
    userId *string;
    userPrincipalName *string;
    wiFiMacAddress *string;
}
func NewManagedDevice()(*ManagedDevice) {
    m := &ManagedDevice{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ManagedDevice) GetActivationLockBypassCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activationLockBypassCode
    }
}
func (m *ManagedDevice) GetAndroidSecurityPatchLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.androidSecurityPatchLevel
    }
}
func (m *ManagedDevice) GetAzureADDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureADDeviceId
    }
}
func (m *ManagedDevice) GetAzureADRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureADRegistered
    }
}
func (m *ManagedDevice) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.complianceGracePeriodExpirationDateTime
    }
}
func (m *ManagedDevice) GetComplianceState()(*ComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.complianceState
    }
}
func (m *ManagedDevice) GetConfigurationManagerClientEnabledFeatures()(*ConfigurationManagerClientEnabledFeatures) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerClientEnabledFeatures
    }
}
func (m *ManagedDevice) GetDeviceActionResults()([]DeviceActionResult) {
    if m == nil {
        return nil
    } else {
        return m.deviceActionResults
    }
}
func (m *ManagedDevice) GetDeviceCategory()(*DeviceCategory) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategory
    }
}
func (m *ManagedDevice) GetDeviceCategoryDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategoryDisplayName
    }
}
func (m *ManagedDevice) GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyState) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicyStates
    }
}
func (m *ManagedDevice) GetDeviceConfigurationStates()([]DeviceConfigurationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationStates
    }
}
func (m *ManagedDevice) GetDeviceEnrollmentType()(*DeviceEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentType
    }
}
func (m *ManagedDevice) GetDeviceHealthAttestationState()(*DeviceHealthAttestationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceHealthAttestationState
    }
}
func (m *ManagedDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *ManagedDevice) GetDeviceRegistrationState()(*DeviceRegistrationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRegistrationState
    }
}
func (m *ManagedDevice) GetEasActivated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.easActivated
    }
}
func (m *ManagedDevice) GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.easActivationDateTime
    }
}
func (m *ManagedDevice) GetEasDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.easDeviceId
    }
}
func (m *ManagedDevice) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
func (m *ManagedDevice) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDateTime
    }
}
func (m *ManagedDevice) GetEthernetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ethernetMacAddress
    }
}
func (m *ManagedDevice) GetExchangeAccessState()(*DeviceManagementExchangeAccessState) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAccessState
    }
}
func (m *ManagedDevice) GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAccessStateReason
    }
}
func (m *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.exchangeLastSuccessfulSyncDateTime
    }
}
func (m *ManagedDevice) GetFreeStorageSpaceInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.freeStorageSpaceInBytes
    }
}
func (m *ManagedDevice) GetIccid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.iccid
    }
}
func (m *ManagedDevice) GetImei()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imei
    }
}
func (m *ManagedDevice) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
func (m *ManagedDevice) GetIsSupervised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSupervised
    }
}
func (m *ManagedDevice) GetJailBroken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jailBroken
    }
}
func (m *ManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
func (m *ManagedDevice) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
func (m *ManagedDevice) GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceOwnerType
    }
}
func (m *ManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    if m == nil {
        return nil
    } else {
        return m.managementAgent
    }
}
func (m *ManagedDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
func (m *ManagedDevice) GetMeid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.meid
    }
}
func (m *ManagedDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *ManagedDevice) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
func (m *ManagedDevice) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
func (m *ManagedDevice) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
func (m *ManagedDevice) GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState) {
    if m == nil {
        return nil
    } else {
        return m.partnerReportedThreatState
    }
}
func (m *ManagedDevice) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
func (m *ManagedDevice) GetPhysicalMemoryInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.physicalMemoryInBytes
    }
}
func (m *ManagedDevice) GetRemoteAssistanceSessionErrorDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSessionErrorDetails
    }
}
func (m *ManagedDevice) GetRemoteAssistanceSessionUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSessionUrl
    }
}
func (m *ManagedDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
func (m *ManagedDevice) GetSubscriberCarrier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriberCarrier
    }
}
func (m *ManagedDevice) GetTotalStorageSpaceInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalStorageSpaceInBytes
    }
}
func (m *ManagedDevice) GetUdid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.udid
    }
}
func (m *ManagedDevice) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
func (m *ManagedDevice) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *ManagedDevice) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *ManagedDevice) GetWiFiMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wiFiMacAddress
    }
}
func (m *ManagedDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationLockBypassCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivationLockBypassCode(val)
        return nil
    }
    res["androidSecurityPatchLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAndroidSecurityPatchLevel(val)
        return nil
    }
    res["azureADDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureADDeviceId(val)
        return nil
    }
    res["azureADRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAzureADRegistered(val)
        return nil
    }
    res["complianceGracePeriodExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetComplianceGracePeriodExpirationDateTime(val)
        return nil
    }
    res["complianceState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceState)
        if err != nil {
            return err
        }
        cast := val.(ComplianceState)
        m.SetComplianceState(&cast)
        return nil
    }
    res["configurationManagerClientEnabledFeatures"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConfigurationManagerClientEnabledFeatures() })
        if err != nil {
            return err
        }
        m.SetConfigurationManagerClientEnabledFeatures(val.(*ConfigurationManagerClientEnabledFeatures))
        return nil
    }
    res["deviceActionResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceActionResult() })
        if err != nil {
            return err
        }
        res := make([]DeviceActionResult, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceActionResult))
        }
        m.SetDeviceActionResults(res)
        return nil
    }
    res["deviceCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCategory() })
        if err != nil {
            return err
        }
        m.SetDeviceCategory(val.(*DeviceCategory))
        return nil
    }
    res["deviceCategoryDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceCategoryDisplayName(val)
        return nil
    }
    res["deviceCompliancePolicyStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicyState() })
        if err != nil {
            return err
        }
        res := make([]DeviceCompliancePolicyState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCompliancePolicyState))
        }
        m.SetDeviceCompliancePolicyStates(res)
        return nil
    }
    res["deviceConfigurationStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationState() })
        if err != nil {
            return err
        }
        res := make([]DeviceConfigurationState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceConfigurationState))
        }
        m.SetDeviceConfigurationStates(res)
        return nil
    }
    res["deviceEnrollmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentType)
        if err != nil {
            return err
        }
        cast := val.(DeviceEnrollmentType)
        m.SetDeviceEnrollmentType(&cast)
        return nil
    }
    res["deviceHealthAttestationState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthAttestationState() })
        if err != nil {
            return err
        }
        m.SetDeviceHealthAttestationState(val.(*DeviceHealthAttestationState))
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["deviceRegistrationState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        cast := val.(DeviceRegistrationState)
        m.SetDeviceRegistrationState(&cast)
        return nil
    }
    res["easActivated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEasActivated(val)
        return nil
    }
    res["easActivationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEasActivationDateTime(val)
        return nil
    }
    res["easDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEasDeviceId(val)
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailAddress(val)
        return nil
    }
    res["enrolledDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEnrolledDateTime(val)
        return nil
    }
    res["ethernetMacAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEthernetMacAddress(val)
        return nil
    }
    res["exchangeAccessState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessState)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExchangeAccessState)
        m.SetExchangeAccessState(&cast)
        return nil
    }
    res["exchangeAccessStateReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessStateReason)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExchangeAccessStateReason)
        m.SetExchangeAccessStateReason(&cast)
        return nil
    }
    res["exchangeLastSuccessfulSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExchangeLastSuccessfulSyncDateTime(val)
        return nil
    }
    res["freeStorageSpaceInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetFreeStorageSpaceInBytes(val)
        return nil
    }
    res["iccid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIccid(val)
        return nil
    }
    res["imei"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetImei(val)
        return nil
    }
    res["isEncrypted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEncrypted(val)
        return nil
    }
    res["isSupervised"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSupervised(val)
        return nil
    }
    res["jailBroken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJailBroken(val)
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncDateTime(val)
        return nil
    }
    res["managedDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceName(val)
        return nil
    }
    res["managedDeviceOwnerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        cast := val.(ManagedDeviceOwnerType)
        m.SetManagedDeviceOwnerType(&cast)
        return nil
    }
    res["managementAgent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        cast := val.(ManagementAgentType)
        m.SetManagementAgent(&cast)
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["meid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMeid(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotes(val)
        return nil
    }
    res["operatingSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystem(val)
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsVersion(val)
        return nil
    }
    res["partnerReportedThreatState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDevicePartnerReportedHealthState)
        if err != nil {
            return err
        }
        cast := val.(ManagedDevicePartnerReportedHealthState)
        m.SetPartnerReportedThreatState(&cast)
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPhoneNumber(val)
        return nil
    }
    res["physicalMemoryInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPhysicalMemoryInBytes(val)
        return nil
    }
    res["remoteAssistanceSessionErrorDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemoteAssistanceSessionErrorDetails(val)
        return nil
    }
    res["remoteAssistanceSessionUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemoteAssistanceSessionUrl(val)
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSerialNumber(val)
        return nil
    }
    res["subscriberCarrier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubscriberCarrier(val)
        return nil
    }
    res["totalStorageSpaceInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalStorageSpaceInBytes(val)
        return nil
    }
    res["udid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUdid(val)
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserDisplayName(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["wiFiMacAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWiFiMacAddress(val)
        return nil
    }
    return res
}
func (m *ManagedDevice) IsNil()(bool) {
    return m == nil
}
func (m *ManagedDevice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activationLockBypassCode", m.GetActivationLockBypassCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("androidSecurityPatchLevel", m.GetAndroidSecurityPatchLevel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureADDeviceId", m.GetAzureADDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("azureADRegistered", m.GetAzureADRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("complianceGracePeriodExpirationDateTime", m.GetComplianceGracePeriodExpirationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetComplianceState() != nil {
        cast := m.GetComplianceState().String()
        err = writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configurationManagerClientEnabledFeatures", m.GetConfigurationManagerClientEnabledFeatures())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceActionResults()))
        for i, v := range m.GetDeviceActionResults() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceActionResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCategory", m.GetDeviceCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceCategoryDisplayName", m.GetDeviceCategoryDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCompliancePolicyStates()))
        for i, v := range m.GetDeviceCompliancePolicyStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicyStates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceConfigurationStates()))
        for i, v := range m.GetDeviceConfigurationStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentType() != nil {
        cast := m.GetDeviceEnrollmentType().String()
        err = writer.WriteStringValue("deviceEnrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceHealthAttestationState", m.GetDeviceHealthAttestationState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceRegistrationState() != nil {
        cast := m.GetDeviceRegistrationState().String()
        err = writer.WriteStringValue("deviceRegistrationState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("easActivated", m.GetEasActivated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("easActivationDateTime", m.GetEasActivationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("easDeviceId", m.GetEasDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("enrolledDateTime", m.GetEnrolledDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ethernetMacAddress", m.GetEthernetMacAddress())
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessState() != nil {
        cast := m.GetExchangeAccessState().String()
        err = writer.WriteStringValue("exchangeAccessState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessStateReason() != nil {
        cast := m.GetExchangeAccessStateReason().String()
        err = writer.WriteStringValue("exchangeAccessStateReason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("exchangeLastSuccessfulSyncDateTime", m.GetExchangeLastSuccessfulSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("freeStorageSpaceInBytes", m.GetFreeStorageSpaceInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("iccid", m.GetIccid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imei", m.GetImei())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEncrypted", m.GetIsEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSupervised", m.GetIsSupervised())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jailBroken", m.GetJailBroken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceOwnerType() != nil {
        cast := m.GetManagedDeviceOwnerType().String()
        err = writer.WriteStringValue("managedDeviceOwnerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementAgent() != nil {
        cast := m.GetManagementAgent().String()
        err = writer.WriteStringValue("managementAgent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("meid", m.GetMeid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerReportedThreatState() != nil {
        cast := m.GetPartnerReportedThreatState().String()
        err = writer.WriteStringValue("partnerReportedThreatState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("physicalMemoryInBytes", m.GetPhysicalMemoryInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remoteAssistanceSessionErrorDetails", m.GetRemoteAssistanceSessionErrorDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remoteAssistanceSessionUrl", m.GetRemoteAssistanceSessionUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriberCarrier", m.GetSubscriberCarrier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalStorageSpaceInBytes", m.GetTotalStorageSpaceInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("udid", m.GetUdid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("wiFiMacAddress", m.GetWiFiMacAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ManagedDevice) SetActivationLockBypassCode(value *string)() {
    m.activationLockBypassCode = value
}
func (m *ManagedDevice) SetAndroidSecurityPatchLevel(value *string)() {
    m.androidSecurityPatchLevel = value
}
func (m *ManagedDevice) SetAzureADDeviceId(value *string)() {
    m.azureADDeviceId = value
}
func (m *ManagedDevice) SetAzureADRegistered(value *bool)() {
    m.azureADRegistered = value
}
func (m *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceGracePeriodExpirationDateTime = value
}
func (m *ManagedDevice) SetComplianceState(value *ComplianceState)() {
    m.complianceState = value
}
func (m *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(value *ConfigurationManagerClientEnabledFeatures)() {
    m.configurationManagerClientEnabledFeatures = value
}
func (m *ManagedDevice) SetDeviceActionResults(value []DeviceActionResult)() {
    m.deviceActionResults = value
}
func (m *ManagedDevice) SetDeviceCategory(value *DeviceCategory)() {
    m.deviceCategory = value
}
func (m *ManagedDevice) SetDeviceCategoryDisplayName(value *string)() {
    m.deviceCategoryDisplayName = value
}
func (m *ManagedDevice) SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyState)() {
    m.deviceCompliancePolicyStates = value
}
func (m *ManagedDevice) SetDeviceConfigurationStates(value []DeviceConfigurationState)() {
    m.deviceConfigurationStates = value
}
func (m *ManagedDevice) SetDeviceEnrollmentType(value *DeviceEnrollmentType)() {
    m.deviceEnrollmentType = value
}
func (m *ManagedDevice) SetDeviceHealthAttestationState(value *DeviceHealthAttestationState)() {
    m.deviceHealthAttestationState = value
}
func (m *ManagedDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *ManagedDevice) SetDeviceRegistrationState(value *DeviceRegistrationState)() {
    m.deviceRegistrationState = value
}
func (m *ManagedDevice) SetEasActivated(value *bool)() {
    m.easActivated = value
}
func (m *ManagedDevice) SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.easActivationDateTime = value
}
func (m *ManagedDevice) SetEasDeviceId(value *string)() {
    m.easDeviceId = value
}
func (m *ManagedDevice) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
func (m *ManagedDevice) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrolledDateTime = value
}
func (m *ManagedDevice) SetEthernetMacAddress(value *string)() {
    m.ethernetMacAddress = value
}
func (m *ManagedDevice) SetExchangeAccessState(value *DeviceManagementExchangeAccessState)() {
    m.exchangeAccessState = value
}
func (m *ManagedDevice) SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)() {
    m.exchangeAccessStateReason = value
}
func (m *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.exchangeLastSuccessfulSyncDateTime = value
}
func (m *ManagedDevice) SetFreeStorageSpaceInBytes(value *int64)() {
    m.freeStorageSpaceInBytes = value
}
func (m *ManagedDevice) SetIccid(value *string)() {
    m.iccid = value
}
func (m *ManagedDevice) SetImei(value *string)() {
    m.imei = value
}
func (m *ManagedDevice) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
func (m *ManagedDevice) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
func (m *ManagedDevice) SetJailBroken(value *string)() {
    m.jailBroken = value
}
func (m *ManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
func (m *ManagedDevice) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
func (m *ManagedDevice) SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)() {
    m.managedDeviceOwnerType = value
}
func (m *ManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    m.managementAgent = value
}
func (m *ManagedDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
func (m *ManagedDevice) SetMeid(value *string)() {
    m.meid = value
}
func (m *ManagedDevice) SetModel(value *string)() {
    m.model = value
}
func (m *ManagedDevice) SetNotes(value *string)() {
    m.notes = value
}
func (m *ManagedDevice) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
func (m *ManagedDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
func (m *ManagedDevice) SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)() {
    m.partnerReportedThreatState = value
}
func (m *ManagedDevice) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
func (m *ManagedDevice) SetPhysicalMemoryInBytes(value *int64)() {
    m.physicalMemoryInBytes = value
}
func (m *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(value *string)() {
    m.remoteAssistanceSessionErrorDetails = value
}
func (m *ManagedDevice) SetRemoteAssistanceSessionUrl(value *string)() {
    m.remoteAssistanceSessionUrl = value
}
func (m *ManagedDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
func (m *ManagedDevice) SetSubscriberCarrier(value *string)() {
    m.subscriberCarrier = value
}
func (m *ManagedDevice) SetTotalStorageSpaceInBytes(value *int64)() {
    m.totalStorageSpaceInBytes = value
}
func (m *ManagedDevice) SetUdid(value *string)() {
    m.udid = value
}
func (m *ManagedDevice) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
func (m *ManagedDevice) SetUserId(value *string)() {
    m.userId = value
}
func (m *ManagedDevice) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
func (m *ManagedDevice) SetWiFiMacAddress(value *string)() {
    m.wiFiMacAddress = value
}
