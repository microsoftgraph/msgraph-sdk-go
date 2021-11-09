package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AgreementAcceptance struct {
    Entity
    // The identifier of the agreement file accepted by the user.
    agreementFileId *string;
    // The identifier of the agreement.
    agreementId *string;
    // The display name of the device used for accepting the agreement.
    deviceDisplayName *string;
    // The unique identifier of the device used for accepting the agreement.
    deviceId *string;
    // The operating system used to accept the agreement.
    deviceOSType *string;
    // The operating system version of the device used to accept the agreement.
    deviceOSVersion *string;
    // The expiration date time of the acceptance. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    recordedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The state of the agreement acceptance. Possible values are: accepted, declined.
    state *AgreementAcceptanceState;
    // Display name of the user when the acceptance was recorded.
    userDisplayName *string;
    // Email of the user when the acceptance was recorded.
    userEmail *string;
    // The identifier of the user who accepted the agreement.
    userId *string;
    // UPN of the user when the acceptance was recorded.
    userPrincipalName *string;
}
// Instantiates a new agreementAcceptance and sets the default values.
func NewAgreementAcceptance()(*AgreementAcceptance) {
    m := &AgreementAcceptance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the agreementFileId property value. The identifier of the agreement file accepted by the user.
func (m *AgreementAcceptance) GetAgreementFileId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.agreementFileId
    }
}
// Gets the agreementId property value. The identifier of the agreement.
func (m *AgreementAcceptance) GetAgreementId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.agreementId
    }
}
// Gets the deviceDisplayName property value. The display name of the device used for accepting the agreement.
func (m *AgreementAcceptance) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
// Gets the deviceId property value. The unique identifier of the device used for accepting the agreement.
func (m *AgreementAcceptance) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceOSType property value. The operating system used to accept the agreement.
func (m *AgreementAcceptance) GetDeviceOSType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOSType
    }
}
// Gets the deviceOSVersion property value. The operating system version of the device used to accept the agreement.
func (m *AgreementAcceptance) GetDeviceOSVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOSVersion
    }
}
// Gets the expirationDateTime property value. The expiration date time of the acceptance. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementAcceptance) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the recordedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementAcceptance) GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.recordedDateTime
    }
}
// Gets the state property value. The state of the agreement acceptance. Possible values are: accepted, declined.
func (m *AgreementAcceptance) GetState()(*AgreementAcceptanceState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the userDisplayName property value. Display name of the user when the acceptance was recorded.
func (m *AgreementAcceptance) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// Gets the userEmail property value. Email of the user when the acceptance was recorded.
func (m *AgreementAcceptance) GetUserEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userEmail
    }
}
// Gets the userId property value. The identifier of the user who accepted the agreement.
func (m *AgreementAcceptance) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userPrincipalName property value. UPN of the user when the acceptance was recorded.
func (m *AgreementAcceptance) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *AgreementAcceptance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agreementFileId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgreementFileId(val)
        }
        return nil
    }
    res["agreementId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgreementId(val)
        }
        return nil
    }
    res["deviceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
        }
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceOSType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOSType(val)
        }
        return nil
    }
    res["deviceOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOSVersion(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["recordedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordedDateTime(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAgreementAcceptanceState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AgreementAcceptanceState)
            m.SetState(&cast)
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
    res["userEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
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
    return res
}
func (m *AgreementAcceptance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AgreementAcceptance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("agreementFileId", m.GetAgreementFileId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("agreementId", m.GetAgreementId())
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
        err = writer.WriteStringValue("deviceOSType", m.GetDeviceOSType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceOSVersion", m.GetDeviceOSVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("recordedDateTime", m.GetRecordedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
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
        err = writer.WriteStringValue("userEmail", m.GetUserEmail())
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
    return nil
}
// Sets the agreementFileId property value. The identifier of the agreement file accepted by the user.
// Parameters:
//  - value : Value to set for the agreementFileId property.
func (m *AgreementAcceptance) SetAgreementFileId(value *string)() {
    m.agreementFileId = value
}
// Sets the agreementId property value. The identifier of the agreement.
// Parameters:
//  - value : Value to set for the agreementId property.
func (m *AgreementAcceptance) SetAgreementId(value *string)() {
    m.agreementId = value
}
// Sets the deviceDisplayName property value. The display name of the device used for accepting the agreement.
// Parameters:
//  - value : Value to set for the deviceDisplayName property.
func (m *AgreementAcceptance) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// Sets the deviceId property value. The unique identifier of the device used for accepting the agreement.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *AgreementAcceptance) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceOSType property value. The operating system used to accept the agreement.
// Parameters:
//  - value : Value to set for the deviceOSType property.
func (m *AgreementAcceptance) SetDeviceOSType(value *string)() {
    m.deviceOSType = value
}
// Sets the deviceOSVersion property value. The operating system version of the device used to accept the agreement.
// Parameters:
//  - value : Value to set for the deviceOSVersion property.
func (m *AgreementAcceptance) SetDeviceOSVersion(value *string)() {
    m.deviceOSVersion = value
}
// Sets the expirationDateTime property value. The expiration date time of the acceptance. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *AgreementAcceptance) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the recordedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the recordedDateTime property.
func (m *AgreementAcceptance) SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.recordedDateTime = value
}
// Sets the state property value. The state of the agreement acceptance. Possible values are: accepted, declined.
// Parameters:
//  - value : Value to set for the state property.
func (m *AgreementAcceptance) SetState(value *AgreementAcceptanceState)() {
    m.state = value
}
// Sets the userDisplayName property value. Display name of the user when the acceptance was recorded.
// Parameters:
//  - value : Value to set for the userDisplayName property.
func (m *AgreementAcceptance) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// Sets the userEmail property value. Email of the user when the acceptance was recorded.
// Parameters:
//  - value : Value to set for the userEmail property.
func (m *AgreementAcceptance) SetUserEmail(value *string)() {
    m.userEmail = value
}
// Sets the userId property value. The identifier of the user who accepted the agreement.
// Parameters:
//  - value : Value to set for the userId property.
func (m *AgreementAcceptance) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userPrincipalName property value. UPN of the user when the acceptance was recorded.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *AgreementAcceptance) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
