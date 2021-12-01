package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AgreementAcceptance 
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
// NewAgreementAcceptance instantiates a new agreementAcceptance and sets the default values.
func NewAgreementAcceptance()(*AgreementAcceptance) {
    m := &AgreementAcceptance{
        Entity: *NewEntity(),
    }
    return m
}
// GetAgreementFileId gets the agreementFileId property value. The identifier of the agreement file accepted by the user.
func (m *AgreementAcceptance) GetAgreementFileId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.agreementFileId
    }
}
// GetAgreementId gets the agreementId property value. The identifier of the agreement.
func (m *AgreementAcceptance) GetAgreementId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.agreementId
    }
}
// GetDeviceDisplayName gets the deviceDisplayName property value. The display name of the device used for accepting the agreement.
func (m *AgreementAcceptance) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
// GetDeviceId gets the deviceId property value. The unique identifier of the device used for accepting the agreement.
func (m *AgreementAcceptance) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDeviceOSType gets the deviceOSType property value. The operating system used to accept the agreement.
func (m *AgreementAcceptance) GetDeviceOSType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOSType
    }
}
// GetDeviceOSVersion gets the deviceOSVersion property value. The operating system version of the device used to accept the agreement.
func (m *AgreementAcceptance) GetDeviceOSVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOSVersion
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The expiration date time of the acceptance. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementAcceptance) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetRecordedDateTime gets the recordedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementAcceptance) GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.recordedDateTime
    }
}
// GetState gets the state property value. The state of the agreement acceptance. Possible values are: accepted, declined.
func (m *AgreementAcceptance) GetState()(*AgreementAcceptanceState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetUserDisplayName gets the userDisplayName property value. Display name of the user when the acceptance was recorded.
func (m *AgreementAcceptance) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserEmail gets the userEmail property value. Email of the user when the acceptance was recorded.
func (m *AgreementAcceptance) GetUserEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userEmail
    }
}
// GetUserId gets the userId property value. The identifier of the user who accepted the agreement.
func (m *AgreementAcceptance) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. UPN of the user when the acceptance was recorded.
func (m *AgreementAcceptance) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAgreementFileId sets the agreementFileId property value. The identifier of the agreement file accepted by the user.
func (m *AgreementAcceptance) SetAgreementFileId(value *string)() {
    if m != nil {
        m.agreementFileId = value
    }
}
// SetAgreementId sets the agreementId property value. The identifier of the agreement.
func (m *AgreementAcceptance) SetAgreementId(value *string)() {
    if m != nil {
        m.agreementId = value
    }
}
// SetDeviceDisplayName sets the deviceDisplayName property value. The display name of the device used for accepting the agreement.
func (m *AgreementAcceptance) SetDeviceDisplayName(value *string)() {
    if m != nil {
        m.deviceDisplayName = value
    }
}
// SetDeviceId sets the deviceId property value. The unique identifier of the device used for accepting the agreement.
func (m *AgreementAcceptance) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDeviceOSType sets the deviceOSType property value. The operating system used to accept the agreement.
func (m *AgreementAcceptance) SetDeviceOSType(value *string)() {
    if m != nil {
        m.deviceOSType = value
    }
}
// SetDeviceOSVersion sets the deviceOSVersion property value. The operating system version of the device used to accept the agreement.
func (m *AgreementAcceptance) SetDeviceOSVersion(value *string)() {
    if m != nil {
        m.deviceOSVersion = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expiration date time of the acceptance. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementAcceptance) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetRecordedDateTime sets the recordedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementAcceptance) SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.recordedDateTime = value
    }
}
// SetState sets the state property value. The state of the agreement acceptance. Possible values are: accepted, declined.
func (m *AgreementAcceptance) SetState(value *AgreementAcceptanceState)() {
    if m != nil {
        m.state = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. Display name of the user when the acceptance was recorded.
func (m *AgreementAcceptance) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserEmail sets the userEmail property value. Email of the user when the acceptance was recorded.
func (m *AgreementAcceptance) SetUserEmail(value *string)() {
    if m != nil {
        m.userEmail = value
    }
}
// SetUserId sets the userId property value. The identifier of the user who accepted the agreement.
func (m *AgreementAcceptance) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. UPN of the user when the acceptance was recorded.
func (m *AgreementAcceptance) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
