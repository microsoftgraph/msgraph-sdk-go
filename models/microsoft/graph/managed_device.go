package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedDevice struct {
    Entity
    // Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
    activationLockBypassCode *string;
    // Android security patch level. This property is read-only.
    androidSecurityPatchLevel *string;
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureADDeviceId *string;
    // Whether the device is Azure Active Directory registered. This property is read-only.
    azureADRegistered *bool;
    // The DateTime when device compliance grace period expires. This property is read-only.
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
    complianceState *ComplianceState;
    // ConfigrMgr client enabled features. This property is read-only.
    configurationManagerClientEnabledFeatures *ConfigurationManagerClientEnabledFeatures;
    // List of ComplexType deviceActionResult objects. This property is read-only.
    deviceActionResults []DeviceActionResult;
    // Device category
    deviceCategory *DeviceCategory;
    // Device category display name. This property is read-only.
    deviceCategoryDisplayName *string;
    // Device compliance policy states for this device.
    deviceCompliancePolicyStates []DeviceCompliancePolicyState;
    // Device configuration states for this device.
    deviceConfigurationStates []DeviceConfigurationState;
    // Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount.
    deviceEnrollmentType *DeviceEnrollmentType;
    // The device health attestation state. This property is read-only.
    deviceHealthAttestationState *DeviceHealthAttestationState;
    // Name of the device. This property is read-only.
    deviceName *string;
    // Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
    deviceRegistrationState *DeviceRegistrationState;
    // Whether the device is Exchange ActiveSync activated. This property is read-only.
    easActivated *bool;
    // Exchange ActivationSync activation time of the device. This property is read-only.
    easActivationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Exchange ActiveSync Id of the device. This property is read-only.
    easDeviceId *string;
    // Email(s) for the user associated with the device. This property is read-only.
    emailAddress *string;
    // Enrollment time of the device. This property is read-only.
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Ethernet MAC. This property is read-only.
    ethernetMacAddress *string;
    // The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined.
    exchangeAccessState *DeviceManagementExchangeAccessState;
    // The reason for the device's access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
    exchangeAccessStateReason *DeviceManagementExchangeAccessStateReason;
    // Last time the device contacted Exchange. This property is read-only.
    exchangeLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Free Storage in Bytes. This property is read-only.
    freeStorageSpaceInBytes *int64;
    // Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
    iccid *string;
    // IMEI. This property is read-only.
    imei *string;
    // Device encryption status. This property is read-only.
    isEncrypted *bool;
    // Device supervised status. This property is read-only.
    isSupervised *bool;
    // whether the device is jail broken or rooted. This property is read-only.
    jailBroken *string;
    // The date and time that the device last completed a successful sync with Intune. This property is read-only.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Automatically generated name to identify a device. Can be overwritten to a user friendly name.
    managedDeviceName *string;
    // Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
    managedDeviceOwnerType *ManagedDeviceOwnerType;
    // Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
    managementAgent *ManagementAgentType;
    // Manufacturer of the device. This property is read-only.
    manufacturer *string;
    // MEID. This property is read-only.
    meid *string;
    // Model of the device. This property is read-only.
    model *string;
    // Notes on the device created by IT Admin
    notes *string;
    // Operating system of the device. Windows, iOS, etc. This property is read-only.
    operatingSystem *string;
    // Operating system version of the device. This property is read-only.
    osVersion *string;
    // Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
    partnerReportedThreatState *ManagedDevicePartnerReportedHealthState;
    // Phone number of the device. This property is read-only.
    phoneNumber *string;
    // Total Memory in Bytes. This property is read-only.
    physicalMemoryInBytes *int64;
    // An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
    remoteAssistanceSessionErrorDetails *string;
    // Url that allows a Remote Assistance session to be established with the device. This property is read-only.
    remoteAssistanceSessionUrl *string;
    // SerialNumber. This property is read-only.
    serialNumber *string;
    // Subscriber Carrier. This property is read-only.
    subscriberCarrier *string;
    // Total Storage in Bytes. This property is read-only.
    totalStorageSpaceInBytes *int64;
    // Unique Device Identifier for iOS and macOS devices. This property is read-only.
    udid *string;
    // User display name. This property is read-only.
    userDisplayName *string;
    // Unique Identifier for the user associated with the device. This property is read-only.
    userId *string;
    // Device user principal name. This property is read-only.
    userPrincipalName *string;
    // Wi-Fi MAC. This property is read-only.
    wiFiMacAddress *string;
}
// Instantiates a new managedDevice and sets the default values.
func NewManagedDevice()(*ManagedDevice) {
    m := &ManagedDevice{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
func (m *ManagedDevice) GetActivationLockBypassCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activationLockBypassCode
    }
}
// Gets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) GetAndroidSecurityPatchLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.androidSecurityPatchLevel
    }
}
// Gets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureADDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureADDeviceId
    }
}
// Gets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAzureADRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureADRegistered
    }
}
// Gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.complianceGracePeriodExpirationDateTime
    }
}
// Gets the complianceState property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
func (m *ManagedDevice) GetComplianceState()(*ComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.complianceState
    }
}
// Gets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) GetConfigurationManagerClientEnabledFeatures()(*ConfigurationManagerClientEnabledFeatures) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerClientEnabledFeatures
    }
}
// Gets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) GetDeviceActionResults()([]DeviceActionResult) {
    if m == nil {
        return nil
    } else {
        return m.deviceActionResults
    }
}
// Gets the deviceCategory property value. Device category
func (m *ManagedDevice) GetDeviceCategory()(*DeviceCategory) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategory
    }
}
// Gets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) GetDeviceCategoryDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategoryDisplayName
    }
}
// Gets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyState) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicyStates
    }
}
// Gets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) GetDeviceConfigurationStates()([]DeviceConfigurationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationStates
    }
}
// Gets the deviceEnrollmentType property value. Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount.
func (m *ManagedDevice) GetDeviceEnrollmentType()(*DeviceEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentType
    }
}
// Gets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) GetDeviceHealthAttestationState()(*DeviceHealthAttestationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceHealthAttestationState
    }
}
// Gets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the deviceRegistrationState property value. Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
func (m *ManagedDevice) GetDeviceRegistrationState()(*DeviceRegistrationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRegistrationState
    }
}
// Gets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) GetEasActivated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.easActivated
    }
}
// Gets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.easActivationDateTime
    }
}
// Gets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) GetEasDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.easDeviceId
    }
}
// Gets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// Gets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDateTime
    }
}
// Gets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
func (m *ManagedDevice) GetEthernetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ethernetMacAddress
    }
}
// Gets the exchangeAccessState property value. The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined.
func (m *ManagedDevice) GetExchangeAccessState()(*DeviceManagementExchangeAccessState) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAccessState
    }
}
// Gets the exchangeAccessStateReason property value. The reason for the device's access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
func (m *ManagedDevice) GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAccessStateReason
    }
}
// Gets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.exchangeLastSuccessfulSyncDateTime
    }
}
// Gets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetFreeStorageSpaceInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.freeStorageSpaceInBytes
    }
}
// Gets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
func (m *ManagedDevice) GetIccid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.iccid
    }
}
// Gets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) GetImei()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imei
    }
}
// Gets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
// Gets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) GetIsSupervised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSupervised
    }
}
// Gets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) GetJailBroken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jailBroken
    }
}
// Gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// Gets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
// Gets the managedDeviceOwnerType property value. Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
func (m *ManagedDevice) GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceOwnerType
    }
}
// Gets the managementAgent property value. Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *ManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    if m == nil {
        return nil
    } else {
        return m.managementAgent
    }
}
// Gets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) GetMeid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.meid
    }
}
// Gets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the notes property value. Notes on the device created by IT Admin
func (m *ManagedDevice) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Gets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// Gets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the partnerReportedThreatState property value. Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
func (m *ManagedDevice) GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState) {
    if m == nil {
        return nil
    } else {
        return m.partnerReportedThreatState
    }
}
// Gets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// Gets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
func (m *ManagedDevice) GetPhysicalMemoryInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.physicalMemoryInBytes
    }
}
// Gets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionErrorDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSessionErrorDetails
    }
}
// Gets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSessionUrl
    }
}
// Gets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// Gets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) GetSubscriberCarrier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriberCarrier
    }
}
// Gets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetTotalStorageSpaceInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalStorageSpaceInBytes
    }
}
// Gets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
func (m *ManagedDevice) GetUdid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.udid
    }
}
// Gets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// Gets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) GetWiFiMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wiFiMacAddress
    }
}
// The deserialization information for the current model
func (m *ManagedDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationLockBypassCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationLockBypassCode(val)
        }
        return nil
    }
    res["androidSecurityPatchLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidSecurityPatchLevel(val)
        }
        return nil
    }
    res["azureADDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADDeviceId(val)
        }
        return nil
    }
    res["azureADRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADRegistered(val)
        }
        return nil
    }
    res["complianceGracePeriodExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["complianceState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ComplianceState)
            m.SetComplianceState(&cast)
        }
        return nil
    }
    res["configurationManagerClientEnabledFeatures"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConfigurationManagerClientEnabledFeatures() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientEnabledFeatures(val.(*ConfigurationManagerClientEnabledFeatures))
        }
        return nil
    }
    res["deviceActionResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceActionResult() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceActionResult, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceActionResult))
            }
            m.SetDeviceActionResults(res)
        }
        return nil
    }
    res["deviceCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategory(val.(*DeviceCategory))
        }
        return nil
    }
    res["deviceCategoryDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategoryDisplayName(val)
        }
        return nil
    }
    res["deviceCompliancePolicyStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicyState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceCompliancePolicyState))
            }
            m.SetDeviceCompliancePolicyStates(res)
        }
        return nil
    }
    res["deviceConfigurationStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceConfigurationState))
            }
            m.SetDeviceConfigurationStates(res)
        }
        return nil
    }
    res["deviceEnrollmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceEnrollmentType)
            m.SetDeviceEnrollmentType(&cast)
        }
        return nil
    }
    res["deviceHealthAttestationState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthAttestationState() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceHealthAttestationState(val.(*DeviceHealthAttestationState))
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceRegistrationState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceRegistrationState)
            m.SetDeviceRegistrationState(&cast)
        }
        return nil
    }
    res["easActivated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivated(val)
        }
        return nil
    }
    res["easActivationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivationDateTime(val)
        }
        return nil
    }
    res["easDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasDeviceId(val)
        }
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["enrolledDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDateTime(val)
        }
        return nil
    }
    res["ethernetMacAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEthernetMacAddress(val)
        }
        return nil
    }
    res["exchangeAccessState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementExchangeAccessState)
            m.SetExchangeAccessState(&cast)
        }
        return nil
    }
    res["exchangeAccessStateReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessStateReason)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementExchangeAccessStateReason)
            m.SetExchangeAccessStateReason(&cast)
        }
        return nil
    }
    res["exchangeLastSuccessfulSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["freeStorageSpaceInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpaceInBytes(val)
        }
        return nil
    }
    res["iccid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIccid(val)
        }
        return nil
    }
    res["imei"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["isEncrypted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSupervised"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["jailBroken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJailBroken(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["managedDeviceOwnerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ManagedDeviceOwnerType)
            m.SetManagedDeviceOwnerType(&cast)
        }
        return nil
    }
    res["managementAgent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ManagementAgentType)
            m.SetManagementAgent(&cast)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["meid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["operatingSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["partnerReportedThreatState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDevicePartnerReportedHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ManagedDevicePartnerReportedHealthState)
            m.SetPartnerReportedThreatState(&cast)
        }
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["physicalMemoryInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhysicalMemoryInBytes(val)
        }
        return nil
    }
    res["remoteAssistanceSessionErrorDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionErrorDetails(val)
        }
        return nil
    }
    res["remoteAssistanceSessionUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionUrl(val)
        }
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["totalStorageSpaceInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpaceInBytes(val)
        }
        return nil
    }
    res["udid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUdid(val)
        }
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["wiFiMacAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiMacAddress(val)
        }
        return nil
    }
    return res
}
func (m *ManagedDevice) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
// Parameters:
//  - value : Value to set for the activationLockBypassCode property.
func (m *ManagedDevice) SetActivationLockBypassCode(value *string)() {
    m.activationLockBypassCode = value
}
// Sets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
// Parameters:
//  - value : Value to set for the androidSecurityPatchLevel property.
func (m *ManagedDevice) SetAndroidSecurityPatchLevel(value *string)() {
    m.androidSecurityPatchLevel = value
}
// Sets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
// Parameters:
//  - value : Value to set for the azureADDeviceId property.
func (m *ManagedDevice) SetAzureADDeviceId(value *string)() {
    m.azureADDeviceId = value
}
// Sets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
// Parameters:
//  - value : Value to set for the azureADRegistered property.
func (m *ManagedDevice) SetAzureADRegistered(value *bool)() {
    m.azureADRegistered = value
}
// Sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
// Parameters:
//  - value : Value to set for the complianceGracePeriodExpirationDateTime property.
func (m *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceGracePeriodExpirationDateTime = value
}
// Sets the complianceState property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
// Parameters:
//  - value : Value to set for the complianceState property.
func (m *ManagedDevice) SetComplianceState(value *ComplianceState)() {
    m.complianceState = value
}
// Sets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
// Parameters:
//  - value : Value to set for the configurationManagerClientEnabledFeatures property.
func (m *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(value *ConfigurationManagerClientEnabledFeatures)() {
    m.configurationManagerClientEnabledFeatures = value
}
// Sets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
// Parameters:
//  - value : Value to set for the deviceActionResults property.
func (m *ManagedDevice) SetDeviceActionResults(value []DeviceActionResult)() {
    m.deviceActionResults = value
}
// Sets the deviceCategory property value. Device category
// Parameters:
//  - value : Value to set for the deviceCategory property.
func (m *ManagedDevice) SetDeviceCategory(value *DeviceCategory)() {
    m.deviceCategory = value
}
// Sets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
// Parameters:
//  - value : Value to set for the deviceCategoryDisplayName property.
func (m *ManagedDevice) SetDeviceCategoryDisplayName(value *string)() {
    m.deviceCategoryDisplayName = value
}
// Sets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
// Parameters:
//  - value : Value to set for the deviceCompliancePolicyStates property.
func (m *ManagedDevice) SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyState)() {
    m.deviceCompliancePolicyStates = value
}
// Sets the deviceConfigurationStates property value. Device configuration states for this device.
// Parameters:
//  - value : Value to set for the deviceConfigurationStates property.
func (m *ManagedDevice) SetDeviceConfigurationStates(value []DeviceConfigurationState)() {
    m.deviceConfigurationStates = value
}
// Sets the deviceEnrollmentType property value. Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount.
// Parameters:
//  - value : Value to set for the deviceEnrollmentType property.
func (m *ManagedDevice) SetDeviceEnrollmentType(value *DeviceEnrollmentType)() {
    m.deviceEnrollmentType = value
}
// Sets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
// Parameters:
//  - value : Value to set for the deviceHealthAttestationState property.
func (m *ManagedDevice) SetDeviceHealthAttestationState(value *DeviceHealthAttestationState)() {
    m.deviceHealthAttestationState = value
}
// Sets the deviceName property value. Name of the device. This property is read-only.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *ManagedDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the deviceRegistrationState property value. Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
// Parameters:
//  - value : Value to set for the deviceRegistrationState property.
func (m *ManagedDevice) SetDeviceRegistrationState(value *DeviceRegistrationState)() {
    m.deviceRegistrationState = value
}
// Sets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
// Parameters:
//  - value : Value to set for the easActivated property.
func (m *ManagedDevice) SetEasActivated(value *bool)() {
    m.easActivated = value
}
// Sets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
// Parameters:
//  - value : Value to set for the easActivationDateTime property.
func (m *ManagedDevice) SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.easActivationDateTime = value
}
// Sets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
// Parameters:
//  - value : Value to set for the easDeviceId property.
func (m *ManagedDevice) SetEasDeviceId(value *string)() {
    m.easDeviceId = value
}
// Sets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
// Parameters:
//  - value : Value to set for the emailAddress property.
func (m *ManagedDevice) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// Sets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
// Parameters:
//  - value : Value to set for the enrolledDateTime property.
func (m *ManagedDevice) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrolledDateTime = value
}
// Sets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
// Parameters:
//  - value : Value to set for the ethernetMacAddress property.
func (m *ManagedDevice) SetEthernetMacAddress(value *string)() {
    m.ethernetMacAddress = value
}
// Sets the exchangeAccessState property value. The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined.
// Parameters:
//  - value : Value to set for the exchangeAccessState property.
func (m *ManagedDevice) SetExchangeAccessState(value *DeviceManagementExchangeAccessState)() {
    m.exchangeAccessState = value
}
// Sets the exchangeAccessStateReason property value. The reason for the device's access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
// Parameters:
//  - value : Value to set for the exchangeAccessStateReason property.
func (m *ManagedDevice) SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)() {
    m.exchangeAccessStateReason = value
}
// Sets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
// Parameters:
//  - value : Value to set for the exchangeLastSuccessfulSyncDateTime property.
func (m *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.exchangeLastSuccessfulSyncDateTime = value
}
// Sets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
// Parameters:
//  - value : Value to set for the freeStorageSpaceInBytes property.
func (m *ManagedDevice) SetFreeStorageSpaceInBytes(value *int64)() {
    m.freeStorageSpaceInBytes = value
}
// Sets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
// Parameters:
//  - value : Value to set for the iccid property.
func (m *ManagedDevice) SetIccid(value *string)() {
    m.iccid = value
}
// Sets the imei property value. IMEI. This property is read-only.
// Parameters:
//  - value : Value to set for the imei property.
func (m *ManagedDevice) SetImei(value *string)() {
    m.imei = value
}
// Sets the isEncrypted property value. Device encryption status. This property is read-only.
// Parameters:
//  - value : Value to set for the isEncrypted property.
func (m *ManagedDevice) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// Sets the isSupervised property value. Device supervised status. This property is read-only.
// Parameters:
//  - value : Value to set for the isSupervised property.
func (m *ManagedDevice) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
// Sets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
// Parameters:
//  - value : Value to set for the jailBroken property.
func (m *ManagedDevice) SetJailBroken(value *string)() {
    m.jailBroken = value
}
// Sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
// Parameters:
//  - value : Value to set for the lastSyncDateTime property.
func (m *ManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// Sets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
// Parameters:
//  - value : Value to set for the managedDeviceName property.
func (m *ManagedDevice) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// Sets the managedDeviceOwnerType property value. Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
// Parameters:
//  - value : Value to set for the managedDeviceOwnerType property.
func (m *ManagedDevice) SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)() {
    m.managedDeviceOwnerType = value
}
// Sets the managementAgent property value. Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
// Parameters:
//  - value : Value to set for the managementAgent property.
func (m *ManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    m.managementAgent = value
}
// Sets the manufacturer property value. Manufacturer of the device. This property is read-only.
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *ManagedDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the meid property value. MEID. This property is read-only.
// Parameters:
//  - value : Value to set for the meid property.
func (m *ManagedDevice) SetMeid(value *string)() {
    m.meid = value
}
// Sets the model property value. Model of the device. This property is read-only.
// Parameters:
//  - value : Value to set for the model property.
func (m *ManagedDevice) SetModel(value *string)() {
    m.model = value
}
// Sets the notes property value. Notes on the device created by IT Admin
// Parameters:
//  - value : Value to set for the notes property.
func (m *ManagedDevice) SetNotes(value *string)() {
    m.notes = value
}
// Sets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
// Parameters:
//  - value : Value to set for the operatingSystem property.
func (m *ManagedDevice) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// Sets the osVersion property value. Operating system version of the device. This property is read-only.
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *ManagedDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the partnerReportedThreatState property value. Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
// Parameters:
//  - value : Value to set for the partnerReportedThreatState property.
func (m *ManagedDevice) SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)() {
    m.partnerReportedThreatState = value
}
// Sets the phoneNumber property value. Phone number of the device. This property is read-only.
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *ManagedDevice) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// Sets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
// Parameters:
//  - value : Value to set for the physicalMemoryInBytes property.
func (m *ManagedDevice) SetPhysicalMemoryInBytes(value *int64)() {
    m.physicalMemoryInBytes = value
}
// Sets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
// Parameters:
//  - value : Value to set for the remoteAssistanceSessionErrorDetails property.
func (m *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(value *string)() {
    m.remoteAssistanceSessionErrorDetails = value
}
// Sets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
// Parameters:
//  - value : Value to set for the remoteAssistanceSessionUrl property.
func (m *ManagedDevice) SetRemoteAssistanceSessionUrl(value *string)() {
    m.remoteAssistanceSessionUrl = value
}
// Sets the serialNumber property value. SerialNumber. This property is read-only.
// Parameters:
//  - value : Value to set for the serialNumber property.
func (m *ManagedDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// Sets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
// Parameters:
//  - value : Value to set for the subscriberCarrier property.
func (m *ManagedDevice) SetSubscriberCarrier(value *string)() {
    m.subscriberCarrier = value
}
// Sets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
// Parameters:
//  - value : Value to set for the totalStorageSpaceInBytes property.
func (m *ManagedDevice) SetTotalStorageSpaceInBytes(value *int64)() {
    m.totalStorageSpaceInBytes = value
}
// Sets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
// Parameters:
//  - value : Value to set for the udid property.
func (m *ManagedDevice) SetUdid(value *string)() {
    m.udid = value
}
// Sets the userDisplayName property value. User display name. This property is read-only.
// Parameters:
//  - value : Value to set for the userDisplayName property.
func (m *ManagedDevice) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// Sets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
// Parameters:
//  - value : Value to set for the userId property.
func (m *ManagedDevice) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userPrincipalName property value. Device user principal name. This property is read-only.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *ManagedDevice) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
// Parameters:
//  - value : Value to set for the wiFiMacAddress property.
func (m *ManagedDevice) SetWiFiMacAddress(value *string)() {
    m.wiFiMacAddress = value
}
