package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SelfSignedCertificate 
type SelfSignedCertificate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Custom key identifier.
    customKeyIdentifier []byte;
    // The friendly name for the key.
    displayName *string;
    // The date and time at which the credential expires. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The value for the key credential. Should be a base-64 encoded value.
    key []byte;
    // The unique identifier (GUID) for the key.
    keyId *string;
    // The date and time at which the credential becomes valid. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The thumbprint value for the key.
    thumbprint *string;
    // The type of key credential. 'AsymmetricX509Cert'.
    type_escaped *string;
    // A string that describes the purpose for which the key can be used. For example, 'Verify'.
    usage *string;
}
// NewSelfSignedCertificate instantiates a new SelfSignedCertificate and sets the default values.
func NewSelfSignedCertificate()(*SelfSignedCertificate) {
    m := &SelfSignedCertificate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSelfSignedCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSelfSignedCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSelfSignedCertificate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SelfSignedCertificate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCustomKeyIdentifier gets the customKeyIdentifier property value. Custom key identifier.
func (m *SelfSignedCertificate) GetCustomKeyIdentifier()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.customKeyIdentifier
    }
}
// GetDisplayName gets the displayName property value. The friendly name for the key.
func (m *SelfSignedCertificate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEndDateTime gets the endDateTime property value. The date and time at which the credential expires. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SelfSignedCertificate) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SelfSignedCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["customKeyIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomKeyIdentifier(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["thumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbprint(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["usage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsage(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The value for the key credential. Should be a base-64 encoded value.
func (m *SelfSignedCertificate) GetKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// GetKeyId gets the keyId property value. The unique identifier (GUID) for the key.
func (m *SelfSignedCertificate) GetKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyId
    }
}
// GetStartDateTime gets the startDateTime property value. The date and time at which the credential becomes valid. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SelfSignedCertificate) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetThumbprint gets the thumbprint property value. The thumbprint value for the key.
func (m *SelfSignedCertificate) GetThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbprint
    }
}
// GetType gets the type property value. The type of key credential. 'AsymmetricX509Cert'.
func (m *SelfSignedCertificate) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUsage gets the usage property value. A string that describes the purpose for which the key can be used. For example, 'Verify'.
func (m *SelfSignedCertificate) GetUsage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usage
    }
}
// Serialize serializes information the current object
func (m *SelfSignedCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("customKeyIdentifier", m.GetCustomKeyIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("keyId", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbprint", m.GetThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("usage", m.GetUsage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SelfSignedCertificate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCustomKeyIdentifier sets the customKeyIdentifier property value. Custom key identifier.
func (m *SelfSignedCertificate) SetCustomKeyIdentifier(value []byte)() {
    if m != nil {
        m.customKeyIdentifier = value
    }
}
// SetDisplayName sets the displayName property value. The friendly name for the key.
func (m *SelfSignedCertificate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEndDateTime sets the endDateTime property value. The date and time at which the credential expires. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SelfSignedCertificate) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetKey sets the key property value. The value for the key credential. Should be a base-64 encoded value.
func (m *SelfSignedCertificate) SetKey(value []byte)() {
    if m != nil {
        m.key = value
    }
}
// SetKeyId sets the keyId property value. The unique identifier (GUID) for the key.
func (m *SelfSignedCertificate) SetKeyId(value *string)() {
    if m != nil {
        m.keyId = value
    }
}
// SetStartDateTime sets the startDateTime property value. The date and time at which the credential becomes valid. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SelfSignedCertificate) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetThumbprint sets the thumbprint property value. The thumbprint value for the key.
func (m *SelfSignedCertificate) SetThumbprint(value *string)() {
    if m != nil {
        m.thumbprint = value
    }
}
// SetType sets the type property value. The type of key credential. 'AsymmetricX509Cert'.
func (m *SelfSignedCertificate) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUsage sets the usage property value. A string that describes the purpose for which the key can be used. For example, 'Verify'.
func (m *SelfSignedCertificate) SetUsage(value *string)() {
    if m != nil {
        m.usage = value
    }
}
