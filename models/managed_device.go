package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevice 
type ManagedDevice struct {
    Entity
    // Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
    activationLockBypassCode *string
    // Android security patch level. This property is read-only.
    androidSecurityPatchLevel *string
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureADDeviceId *string
    // Whether the device is Azure Active Directory registered. This property is read-only.
    azureADRegistered *bool
    // The DateTime when device compliance grace period expires. This property is read-only.
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
    complianceState *ComplianceState
    // ConfigrMgr client enabled features. This property is read-only.
    configurationManagerClientEnabledFeatures ConfigurationManagerClientEnabledFeaturesable
    // List of ComplexType deviceActionResult objects. This property is read-only.
    deviceActionResults []DeviceActionResultable
    // Device category
    deviceCategory DeviceCategoryable
    // Device category display name. This property is read-only.
    deviceCategoryDisplayName *string
    // Device compliance policy states for this device.
    deviceCompliancePolicyStates []DeviceCompliancePolicyStateable
    // Device configuration states for this device.
    deviceConfigurationStates []DeviceConfigurationStateable
    // Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount.
    deviceEnrollmentType *DeviceEnrollmentType
    // The device health attestation state. This property is read-only.
    deviceHealthAttestationState DeviceHealthAttestationStateable
    // Name of the device. This property is read-only.
    deviceName *string
    // Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
    deviceRegistrationState *DeviceRegistrationState
    // Whether the device is Exchange ActiveSync activated. This property is read-only.
    easActivated *bool
    // Exchange ActivationSync activation time of the device. This property is read-only.
    easActivationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Exchange ActiveSync Id of the device. This property is read-only.
    easDeviceId *string
    // Email(s) for the user associated with the device. This property is read-only.
    emailAddress *string
    // Enrollment time of the device. This property is read-only.
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Ethernet MAC. This property is read-only.
    ethernetMacAddress *string
    // The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined.
    exchangeAccessState *DeviceManagementExchangeAccessState
    // The reason for the device's access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
    exchangeAccessStateReason *DeviceManagementExchangeAccessStateReason
    // Last time the device contacted Exchange. This property is read-only.
    exchangeLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Free Storage in Bytes. This property is read-only.
    freeStorageSpaceInBytes *int64
    // Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
    iccid *string
    // IMEI. This property is read-only.
    imei *string
    // Device encryption status. This property is read-only.
    isEncrypted *bool
    // Device supervised status. This property is read-only.
    isSupervised *bool
    // whether the device is jail broken or rooted. This property is read-only.
    jailBroken *string
    // The date and time that the device last completed a successful sync with Intune. This property is read-only.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Automatically generated name to identify a device. Can be overwritten to a user friendly name.
    managedDeviceName *string
    // Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
    managedDeviceOwnerType *ManagedDeviceOwnerType
    // Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
    managementAgent *ManagementAgentType
    // Manufacturer of the device. This property is read-only.
    manufacturer *string
    // MEID. This property is read-only.
    meid *string
    // Model of the device. This property is read-only.
    model *string
    // Notes on the device created by IT Admin
    notes *string
    // Operating system of the device. Windows, iOS, etc. This property is read-only.
    operatingSystem *string
    // Operating system version of the device. This property is read-only.
    osVersion *string
    // Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
    partnerReportedThreatState *ManagedDevicePartnerReportedHealthState
    // Phone number of the device. This property is read-only.
    phoneNumber *string
    // Total Memory in Bytes. This property is read-only.
    physicalMemoryInBytes *int64
    // An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
    remoteAssistanceSessionErrorDetails *string
    // Url that allows a Remote Assistance session to be established with the device. This property is read-only.
    remoteAssistanceSessionUrl *string
    // SerialNumber. This property is read-only.
    serialNumber *string
    // Subscriber Carrier. This property is read-only.
    subscriberCarrier *string
    // Total Storage in Bytes. This property is read-only.
    totalStorageSpaceInBytes *int64
    // Unique Device Identifier for iOS and macOS devices. This property is read-only.
    udid *string
    // User display name. This property is read-only.
    userDisplayName *string
    // Unique Identifier for the user associated with the device. This property is read-only.
    userId *string
    // Device user principal name. This property is read-only.
    userPrincipalName *string
    // Wi-Fi MAC. This property is read-only.
    wiFiMacAddress *string
}
// NewManagedDevice instantiates a new managedDevice and sets the default values.
func NewManagedDevice()(*ManagedDevice) {
    m := &ManagedDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevice(), nil
}
// GetActivationLockBypassCode gets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
func (m *ManagedDevice) GetActivationLockBypassCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activationLockBypassCode
    }
}
// GetAndroidSecurityPatchLevel gets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) GetAndroidSecurityPatchLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.androidSecurityPatchLevel
    }
}
// GetAzureADDeviceId gets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureADDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureADDeviceId
    }
}
// GetAzureADRegistered gets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAzureADRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureADRegistered
    }
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.complianceGracePeriodExpirationDateTime
    }
}
// GetComplianceState gets the complianceState property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
func (m *ManagedDevice) GetComplianceState()(*ComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.complianceState
    }
}
// GetConfigurationManagerClientEnabledFeatures gets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerClientEnabledFeatures
    }
}
// GetDeviceActionResults gets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) GetDeviceActionResults()([]DeviceActionResultable) {
    if m == nil {
        return nil
    } else {
        return m.deviceActionResults
    }
}
// GetDeviceCategory gets the deviceCategory property value. Device category
func (m *ManagedDevice) GetDeviceCategory()(DeviceCategoryable) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategory
    }
}
// GetDeviceCategoryDisplayName gets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) GetDeviceCategoryDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategoryDisplayName
    }
}
// GetDeviceCompliancePolicyStates gets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicyStates
    }
}
// GetDeviceConfigurationStates gets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) GetDeviceConfigurationStates()([]DeviceConfigurationStateable) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationStates
    }
}
// GetDeviceEnrollmentType gets the deviceEnrollmentType property value. Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount.
func (m *ManagedDevice) GetDeviceEnrollmentType()(*DeviceEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentType
    }
}
// GetDeviceHealthAttestationState gets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable) {
    if m == nil {
        return nil
    } else {
        return m.deviceHealthAttestationState
    }
}
// GetDeviceName gets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetDeviceRegistrationState gets the deviceRegistrationState property value. Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
func (m *ManagedDevice) GetDeviceRegistrationState()(*DeviceRegistrationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRegistrationState
    }
}
// GetEasActivated gets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) GetEasActivated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.easActivated
    }
}
// GetEasActivationDateTime gets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.easActivationDateTime
    }
}
// GetEasDeviceId gets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) GetEasDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.easDeviceId
    }
}
// GetEmailAddress gets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// GetEnrolledDateTime gets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDateTime
    }
}
// GetEthernetMacAddress gets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
func (m *ManagedDevice) GetEthernetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ethernetMacAddress
    }
}
// GetExchangeAccessState gets the exchangeAccessState property value. The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined.
func (m *ManagedDevice) GetExchangeAccessState()(*DeviceManagementExchangeAccessState) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAccessState
    }
}
// GetExchangeAccessStateReason gets the exchangeAccessStateReason property value. The reason for the device's access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
func (m *ManagedDevice) GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAccessStateReason
    }
}
// GetExchangeLastSuccessfulSyncDateTime gets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.exchangeLastSuccessfulSyncDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationLockBypassCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationLockBypassCode(val)
        }
        return nil
    }
    res["androidSecurityPatchLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidSecurityPatchLevel(val)
        }
        return nil
    }
    res["azureADDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADDeviceId(val)
        }
        return nil
    }
    res["azureADRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADRegistered(val)
        }
        return nil
    }
    res["complianceGracePeriodExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["complianceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceState(val.(*ComplianceState))
        }
        return nil
    }
    res["configurationManagerClientEnabledFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientEnabledFeatures(val.(ConfigurationManagerClientEnabledFeaturesable))
        }
        return nil
    }
    res["deviceActionResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceActionResultable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceActionResultable)
            }
            m.SetDeviceActionResults(res)
        }
        return nil
    }
    res["deviceCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategory(val.(DeviceCategoryable))
        }
        return nil
    }
    res["deviceCategoryDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategoryDisplayName(val)
        }
        return nil
    }
    res["deviceCompliancePolicyStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicyStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCompliancePolicyStateable)
            }
            m.SetDeviceCompliancePolicyStates(res)
        }
        return nil
    }
    res["deviceConfigurationStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationStateable)
            }
            m.SetDeviceConfigurationStates(res)
        }
        return nil
    }
    res["deviceEnrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceEnrollmentType(val.(*DeviceEnrollmentType))
        }
        return nil
    }
    res["deviceHealthAttestationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceHealthAttestationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceHealthAttestationState(val.(DeviceHealthAttestationStateable))
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceRegistrationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegistrationState(val.(*DeviceRegistrationState))
        }
        return nil
    }
    res["easActivated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivated(val)
        }
        return nil
    }
    res["easActivationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivationDateTime(val)
        }
        return nil
    }
    res["easDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasDeviceId(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["enrolledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDateTime(val)
        }
        return nil
    }
    res["ethernetMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEthernetMacAddress(val)
        }
        return nil
    }
    res["exchangeAccessState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessState(val.(*DeviceManagementExchangeAccessState))
        }
        return nil
    }
    res["exchangeAccessStateReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessStateReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessStateReason(val.(*DeviceManagementExchangeAccessStateReason))
        }
        return nil
    }
    res["exchangeLastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["freeStorageSpaceInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpaceInBytes(val)
        }
        return nil
    }
    res["iccid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIccid(val)
        }
        return nil
    }
    res["imei"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["isEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["jailBroken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJailBroken(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["managedDeviceOwnerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    res["managementAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementAgent(val.(*ManagementAgentType))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["meid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["partnerReportedThreatState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDevicePartnerReportedHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerReportedThreatState(val.(*ManagedDevicePartnerReportedHealthState))
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["physicalMemoryInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhysicalMemoryInBytes(val)
        }
        return nil
    }
    res["remoteAssistanceSessionErrorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionErrorDetails(val)
        }
        return nil
    }
    res["remoteAssistanceSessionUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionUrl(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["totalStorageSpaceInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpaceInBytes(val)
        }
        return nil
    }
    res["udid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUdid(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["wiFiMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFreeStorageSpaceInBytes gets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetFreeStorageSpaceInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.freeStorageSpaceInBytes
    }
}
// GetIccid gets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
func (m *ManagedDevice) GetIccid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.iccid
    }
}
// GetImei gets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) GetImei()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imei
    }
}
// GetIsEncrypted gets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
// GetIsSupervised gets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) GetIsSupervised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSupervised
    }
}
// GetJailBroken gets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) GetJailBroken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jailBroken
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetManagedDeviceName gets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
// GetManagedDeviceOwnerType gets the managedDeviceOwnerType property value. Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
func (m *ManagedDevice) GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceOwnerType
    }
}
// GetManagementAgent gets the managementAgent property value. Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *ManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    if m == nil {
        return nil
    } else {
        return m.managementAgent
    }
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetMeid gets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) GetMeid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.meid
    }
}
// GetModel gets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetNotes gets the notes property value. Notes on the device created by IT Admin
func (m *ManagedDevice) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetOperatingSystem gets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// GetOsVersion gets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetPartnerReportedThreatState gets the partnerReportedThreatState property value. Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
func (m *ManagedDevice) GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState) {
    if m == nil {
        return nil
    } else {
        return m.partnerReportedThreatState
    }
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// GetPhysicalMemoryInBytes gets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
func (m *ManagedDevice) GetPhysicalMemoryInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.physicalMemoryInBytes
    }
}
// GetRemoteAssistanceSessionErrorDetails gets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionErrorDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSessionErrorDetails
    }
}
// GetRemoteAssistanceSessionUrl gets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSessionUrl
    }
}
// GetSerialNumber gets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) GetSubscriberCarrier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriberCarrier
    }
}
// GetTotalStorageSpaceInBytes gets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetTotalStorageSpaceInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalStorageSpaceInBytes
    }
}
// GetUdid gets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
func (m *ManagedDevice) GetUdid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.udid
    }
}
// GetUserDisplayName gets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserId gets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetWiFiMacAddress gets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) GetWiFiMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wiFiMacAddress
    }
}
// Serialize serializes information the current object
func (m *ManagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetComplianceState()).String()
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
    if m.GetDeviceActionResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceActionResults()))
        for i, v := range m.GetDeviceActionResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetDeviceCompliancePolicyStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicyStates()))
        for i, v := range m.GetDeviceCompliancePolicyStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicyStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationStates()))
        for i, v := range m.GetDeviceConfigurationStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentType() != nil {
        cast := (*m.GetDeviceEnrollmentType()).String()
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
        cast := (*m.GetDeviceRegistrationState()).String()
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
        cast := (*m.GetExchangeAccessState()).String()
        err = writer.WriteStringValue("exchangeAccessState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessStateReason() != nil {
        cast := (*m.GetExchangeAccessStateReason()).String()
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
        cast := (*m.GetManagedDeviceOwnerType()).String()
        err = writer.WriteStringValue("managedDeviceOwnerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementAgent() != nil {
        cast := (*m.GetManagementAgent()).String()
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
        cast := (*m.GetPartnerReportedThreatState()).String()
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
// SetActivationLockBypassCode sets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
func (m *ManagedDevice) SetActivationLockBypassCode(value *string)() {
    if m != nil {
        m.activationLockBypassCode = value
    }
}
// SetAndroidSecurityPatchLevel sets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) SetAndroidSecurityPatchLevel(value *string)() {
    if m != nil {
        m.androidSecurityPatchLevel = value
    }
}
// SetAzureADDeviceId sets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureADDeviceId(value *string)() {
    if m != nil {
        m.azureADDeviceId = value
    }
}
// SetAzureADRegistered sets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAzureADRegistered(value *bool)() {
    if m != nil {
        m.azureADRegistered = value
    }
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.complianceGracePeriodExpirationDateTime = value
    }
}
// SetComplianceState sets the complianceState property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
func (m *ManagedDevice) SetComplianceState(value *ComplianceState)() {
    if m != nil {
        m.complianceState = value
    }
}
// SetConfigurationManagerClientEnabledFeatures sets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)() {
    if m != nil {
        m.configurationManagerClientEnabledFeatures = value
    }
}
// SetDeviceActionResults sets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) SetDeviceActionResults(value []DeviceActionResultable)() {
    if m != nil {
        m.deviceActionResults = value
    }
}
// SetDeviceCategory sets the deviceCategory property value. Device category
func (m *ManagedDevice) SetDeviceCategory(value DeviceCategoryable)() {
    if m != nil {
        m.deviceCategory = value
    }
}
// SetDeviceCategoryDisplayName sets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) SetDeviceCategoryDisplayName(value *string)() {
    if m != nil {
        m.deviceCategoryDisplayName = value
    }
}
// SetDeviceCompliancePolicyStates sets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)() {
    if m != nil {
        m.deviceCompliancePolicyStates = value
    }
}
// SetDeviceConfigurationStates sets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) SetDeviceConfigurationStates(value []DeviceConfigurationStateable)() {
    if m != nil {
        m.deviceConfigurationStates = value
    }
}
// SetDeviceEnrollmentType sets the deviceEnrollmentType property value. Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount.
func (m *ManagedDevice) SetDeviceEnrollmentType(value *DeviceEnrollmentType)() {
    if m != nil {
        m.deviceEnrollmentType = value
    }
}
// SetDeviceHealthAttestationState sets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)() {
    if m != nil {
        m.deviceHealthAttestationState = value
    }
}
// SetDeviceName sets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetDeviceRegistrationState sets the deviceRegistrationState property value. Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
func (m *ManagedDevice) SetDeviceRegistrationState(value *DeviceRegistrationState)() {
    if m != nil {
        m.deviceRegistrationState = value
    }
}
// SetEasActivated sets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) SetEasActivated(value *bool)() {
    if m != nil {
        m.easActivated = value
    }
}
// SetEasActivationDateTime sets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.easActivationDateTime = value
    }
}
// SetEasDeviceId sets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) SetEasDeviceId(value *string)() {
    if m != nil {
        m.easDeviceId = value
    }
}
// SetEmailAddress sets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetEmailAddress(value *string)() {
    if m != nil {
        m.emailAddress = value
    }
}
// SetEnrolledDateTime sets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.enrolledDateTime = value
    }
}
// SetEthernetMacAddress sets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
func (m *ManagedDevice) SetEthernetMacAddress(value *string)() {
    if m != nil {
        m.ethernetMacAddress = value
    }
}
// SetExchangeAccessState sets the exchangeAccessState property value. The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined.
func (m *ManagedDevice) SetExchangeAccessState(value *DeviceManagementExchangeAccessState)() {
    if m != nil {
        m.exchangeAccessState = value
    }
}
// SetExchangeAccessStateReason sets the exchangeAccessStateReason property value. The reason for the device's access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
func (m *ManagedDevice) SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)() {
    if m != nil {
        m.exchangeAccessStateReason = value
    }
}
// SetExchangeLastSuccessfulSyncDateTime sets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.exchangeLastSuccessfulSyncDateTime = value
    }
}
// SetFreeStorageSpaceInBytes sets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetFreeStorageSpaceInBytes(value *int64)() {
    if m != nil {
        m.freeStorageSpaceInBytes = value
    }
}
// SetIccid sets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
func (m *ManagedDevice) SetIccid(value *string)() {
    if m != nil {
        m.iccid = value
    }
}
// SetImei sets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) SetImei(value *string)() {
    if m != nil {
        m.imei = value
    }
}
// SetIsEncrypted sets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) SetIsEncrypted(value *bool)() {
    if m != nil {
        m.isEncrypted = value
    }
}
// SetIsSupervised sets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) SetIsSupervised(value *bool)() {
    if m != nil {
        m.isSupervised = value
    }
}
// SetJailBroken sets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) SetJailBroken(value *string)() {
    if m != nil {
        m.jailBroken = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetManagedDeviceName sets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) SetManagedDeviceName(value *string)() {
    if m != nil {
        m.managedDeviceName = value
    }
}
// SetManagedDeviceOwnerType sets the managedDeviceOwnerType property value. Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
func (m *ManagedDevice) SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)() {
    if m != nil {
        m.managedDeviceOwnerType = value
    }
}
// SetManagementAgent sets the managementAgent property value. Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *ManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    if m != nil {
        m.managementAgent = value
    }
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetMeid sets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) SetMeid(value *string)() {
    if m != nil {
        m.meid = value
    }
}
// SetModel sets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetNotes sets the notes property value. Notes on the device created by IT Admin
func (m *ManagedDevice) SetNotes(value *string)() {
    if m != nil {
        m.notes = value
    }
}
// SetOperatingSystem sets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) SetOperatingSystem(value *string)() {
    if m != nil {
        m.operatingSystem = value
    }
}
// SetOsVersion sets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetPartnerReportedThreatState sets the partnerReportedThreatState property value. Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
func (m *ManagedDevice) SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)() {
    if m != nil {
        m.partnerReportedThreatState = value
    }
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) SetPhoneNumber(value *string)() {
    if m != nil {
        m.phoneNumber = value
    }
}
// SetPhysicalMemoryInBytes sets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
func (m *ManagedDevice) SetPhysicalMemoryInBytes(value *int64)() {
    if m != nil {
        m.physicalMemoryInBytes = value
    }
}
// SetRemoteAssistanceSessionErrorDetails sets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(value *string)() {
    if m != nil {
        m.remoteAssistanceSessionErrorDetails = value
    }
}
// SetRemoteAssistanceSessionUrl sets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionUrl(value *string)() {
    if m != nil {
        m.remoteAssistanceSessionUrl = value
    }
}
// SetSerialNumber sets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) SetSerialNumber(value *string)() {
    if m != nil {
        m.serialNumber = value
    }
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) SetSubscriberCarrier(value *string)() {
    if m != nil {
        m.subscriberCarrier = value
    }
}
// SetTotalStorageSpaceInBytes sets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetTotalStorageSpaceInBytes(value *int64)() {
    if m != nil {
        m.totalStorageSpaceInBytes = value
    }
}
// SetUdid sets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
func (m *ManagedDevice) SetUdid(value *string)() {
    if m != nil {
        m.udid = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserId sets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetWiFiMacAddress sets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) SetWiFiMacAddress(value *string)() {
    if m != nil {
        m.wiFiMacAddress = value
    }
}
