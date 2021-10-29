package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ApplePushNotificationCertificate struct {
    Entity
    // Apple Id of the account used to create the MDM push certificate.
    appleIdentifier *string;
    // Not yet documented
    certificate *string;
    // Certificate serial number. This property is read-only.
    certificateSerialNumber *string;
    // The expiration date and time for Apple push notification certificate.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Last modified date and time for Apple push notification certificate.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Topic Id.
    topicIdentifier *string;
}
// Instantiates a new applePushNotificationCertificate and sets the default values.
func NewApplePushNotificationCertificate()(*ApplePushNotificationCertificate) {
    m := &ApplePushNotificationCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appleIdentifier property value. Apple Id of the account used to create the MDM push certificate.
func (m *ApplePushNotificationCertificate) GetAppleIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleIdentifier
    }
}
// Gets the certificate property value. Not yet documented
func (m *ApplePushNotificationCertificate) GetCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
// Gets the certificateSerialNumber property value. Certificate serial number. This property is read-only.
func (m *ApplePushNotificationCertificate) GetCertificateSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateSerialNumber
    }
}
// Gets the expirationDateTime property value. The expiration date and time for Apple push notification certificate.
func (m *ApplePushNotificationCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the lastModifiedDateTime property value. Last modified date and time for Apple push notification certificate.
func (m *ApplePushNotificationCertificate) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the topicIdentifier property value. Topic Id.
func (m *ApplePushNotificationCertificate) GetTopicIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.topicIdentifier
    }
}
// The deserialization information for the current model
func (m *ApplePushNotificationCertificate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appleIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppleIdentifier(val)
        return nil
    }
    res["certificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificate(val)
        return nil
    }
    res["certificateSerialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificateSerialNumber(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["topicIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTopicIdentifier(val)
        return nil
    }
    return res
}
func (m *ApplePushNotificationCertificate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ApplePushNotificationCertificate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appleIdentifier", m.GetAppleIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSerialNumber", m.GetCertificateSerialNumber())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("topicIdentifier", m.GetTopicIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appleIdentifier property value. Apple Id of the account used to create the MDM push certificate.
// Parameters:
//  - value : Value to set for the appleIdentifier property.
func (m *ApplePushNotificationCertificate) SetAppleIdentifier(value *string)() {
    m.appleIdentifier = value
}
// Sets the certificate property value. Not yet documented
// Parameters:
//  - value : Value to set for the certificate property.
func (m *ApplePushNotificationCertificate) SetCertificate(value *string)() {
    m.certificate = value
}
// Sets the certificateSerialNumber property value. Certificate serial number. This property is read-only.
// Parameters:
//  - value : Value to set for the certificateSerialNumber property.
func (m *ApplePushNotificationCertificate) SetCertificateSerialNumber(value *string)() {
    m.certificateSerialNumber = value
}
// Sets the expirationDateTime property value. The expiration date and time for Apple push notification certificate.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *ApplePushNotificationCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the lastModifiedDateTime property value. Last modified date and time for Apple push notification certificate.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *ApplePushNotificationCertificate) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the topicIdentifier property value. Topic Id.
// Parameters:
//  - value : Value to set for the topicIdentifier property.
func (m *ApplePushNotificationCertificate) SetTopicIdentifier(value *string)() {
    m.topicIdentifier = value
}
