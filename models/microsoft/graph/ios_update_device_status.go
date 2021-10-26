package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IosUpdateDeviceStatus struct {
    Entity
    // The DateTime when device compliance grace period expires
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Device name of the DevicePolicyStatus.
    deviceDisplayName *string;
    // The device id that is being reported.
    deviceId *string;
    // The device model that is being reported
    deviceModel *string;
    // The installation status of the policy report. Possible values are: success, available, idle, unknown, downloading, downloadFailed, downloadRequiresComputer, downloadInsufficientSpace, downloadInsufficientPower, downloadInsufficientNetwork, installing, installInsufficientSpace, installInsufficientPower, installPhoneCallInProgress, installFailed, notSupportedOperation, sharedDeviceUserLoggedInError, deviceOsHigherThanDesiredOsVersion.
    installStatus *IosUpdatesInstallStatus;
    // Last modified date time of the policy report.
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The device version that is being reported.
    osVersion *string;
    // Compliance status of the policy report. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
    status *ComplianceStatus;
    // The User id that is being reported.
    userId *string;
    // The User Name that is being reported
    userName *string;
    // UserPrincipalName.
    userPrincipalName *string;
}
// Instantiates a new iosUpdateDeviceStatus and sets the default values.
func NewIosUpdateDeviceStatus()(*IosUpdateDeviceStatus) {
    m := &IosUpdateDeviceStatus{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
func (m *IosUpdateDeviceStatus) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.complianceGracePeriodExpirationDateTime
    }
}
// Gets the deviceDisplayName property value. Device name of the DevicePolicyStatus.
func (m *IosUpdateDeviceStatus) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
// Gets the deviceId property value. The device id that is being reported.
func (m *IosUpdateDeviceStatus) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceModel property value. The device model that is being reported
func (m *IosUpdateDeviceStatus) GetDeviceModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceModel
    }
}
// Gets the installStatus property value. The installation status of the policy report. Possible values are: success, available, idle, unknown, downloading, downloadFailed, downloadRequiresComputer, downloadInsufficientSpace, downloadInsufficientPower, downloadInsufficientNetwork, installing, installInsufficientSpace, installInsufficientPower, installPhoneCallInProgress, installFailed, notSupportedOperation, sharedDeviceUserLoggedInError, deviceOsHigherThanDesiredOsVersion.
func (m *IosUpdateDeviceStatus) GetInstallStatus()(*IosUpdatesInstallStatus) {
    if m == nil {
        return nil
    } else {
        return m.installStatus
    }
}
// Gets the lastReportedDateTime property value. Last modified date time of the policy report.
func (m *IosUpdateDeviceStatus) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
// Gets the osVersion property value. The device version that is being reported.
func (m *IosUpdateDeviceStatus) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the status property value. Compliance status of the policy report. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *IosUpdateDeviceStatus) GetStatus()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the userId property value. The User id that is being reported.
func (m *IosUpdateDeviceStatus) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userName property value. The User Name that is being reported
func (m *IosUpdateDeviceStatus) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// Gets the userPrincipalName property value. UserPrincipalName.
func (m *IosUpdateDeviceStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *IosUpdateDeviceStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceGracePeriodExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetComplianceGracePeriodExpirationDateTime(val)
        return nil
    }
    res["deviceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceDisplayName(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["deviceModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceModel(val)
        return nil
    }
    res["installStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIosUpdatesInstallStatus)
        if err != nil {
            return err
        }
        cast := val.(IosUpdatesInstallStatus)
        m.SetInstallStatus(&cast)
        return nil
    }
    res["lastReportedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastReportedDateTime(val)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        cast := val.(ComplianceStatus)
        m.SetStatus(&cast)
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
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserName(val)
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
    return res
}
func (m *IosUpdateDeviceStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IosUpdateDeviceStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("complianceGracePeriodExpirationDateTime", m.GetComplianceGracePeriodExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    if m.GetInstallStatus() != nil {
        cast := m.GetInstallStatus().String()
        err = writer.WriteStringValue("installStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
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
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
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
        err = writer.WriteStringValue("userName", m.GetUserName())
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
    return nil
}
// Sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
// Parameters:
//  - value : Value to set for the complianceGracePeriodExpirationDateTime property.
func (m *IosUpdateDeviceStatus) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceGracePeriodExpirationDateTime = value
}
// Sets the deviceDisplayName property value. Device name of the DevicePolicyStatus.
// Parameters:
//  - value : Value to set for the deviceDisplayName property.
func (m *IosUpdateDeviceStatus) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// Sets the deviceId property value. The device id that is being reported.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *IosUpdateDeviceStatus) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceModel property value. The device model that is being reported
// Parameters:
//  - value : Value to set for the deviceModel property.
func (m *IosUpdateDeviceStatus) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// Sets the installStatus property value. The installation status of the policy report. Possible values are: success, available, idle, unknown, downloading, downloadFailed, downloadRequiresComputer, downloadInsufficientSpace, downloadInsufficientPower, downloadInsufficientNetwork, installing, installInsufficientSpace, installInsufficientPower, installPhoneCallInProgress, installFailed, notSupportedOperation, sharedDeviceUserLoggedInError, deviceOsHigherThanDesiredOsVersion.
// Parameters:
//  - value : Value to set for the installStatus property.
func (m *IosUpdateDeviceStatus) SetInstallStatus(value *IosUpdatesInstallStatus)() {
    m.installStatus = value
}
// Sets the lastReportedDateTime property value. Last modified date time of the policy report.
// Parameters:
//  - value : Value to set for the lastReportedDateTime property.
func (m *IosUpdateDeviceStatus) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// Sets the osVersion property value. The device version that is being reported.
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *IosUpdateDeviceStatus) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the status property value. Compliance status of the policy report. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
// Parameters:
//  - value : Value to set for the status property.
func (m *IosUpdateDeviceStatus) SetStatus(value *ComplianceStatus)() {
    m.status = value
}
// Sets the userId property value. The User id that is being reported.
// Parameters:
//  - value : Value to set for the userId property.
func (m *IosUpdateDeviceStatus) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userName property value. The User Name that is being reported
// Parameters:
//  - value : Value to set for the userName property.
func (m *IosUpdateDeviceStatus) SetUserName(value *string)() {
    m.userName = value
}
// Sets the userPrincipalName property value. UserPrincipalName.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *IosUpdateDeviceStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
