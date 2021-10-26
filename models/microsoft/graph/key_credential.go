package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type KeyCredential struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Custom key identifier
    customKeyIdentifier []byte;
    // Friendly name for the key. Optional.
    displayName *string;
    // The date and time at which the credential expires.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The certificate's raw data in byte array converted to Base64 string; for example, [System.Convert]::ToBase64String($Cert.GetRawCertData()).
    key []byte;
    // The unique identifier (GUID) for the key.
    keyId *string;
    // The date and time at which the credential becomes valid.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The type of key credential; for example, Symmetric.
    type_escaped *string;
    // A string that describes the purpose for which the key can be used; for example, Verify.
    usage *string;
}
// Instantiates a new keyCredential and sets the default values.
func NewKeyCredential()(*KeyCredential) {
    m := &KeyCredential{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyCredential) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the customKeyIdentifier property value. Custom key identifier
func (m *KeyCredential) GetCustomKeyIdentifier()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.customKeyIdentifier
    }
}
// Gets the displayName property value. Friendly name for the key. Optional.
func (m *KeyCredential) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the endDateTime property value. The date and time at which the credential expires.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *KeyCredential) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the key property value. The certificate's raw data in byte array converted to Base64 string; for example, [System.Convert]::ToBase64String($Cert.GetRawCertData()).
func (m *KeyCredential) GetKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// Gets the keyId property value. The unique identifier (GUID) for the key.
func (m *KeyCredential) GetKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyId
    }
}
// Gets the startDateTime property value. The date and time at which the credential becomes valid.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *KeyCredential) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the type_escaped property value. The type of key credential; for example, Symmetric.
func (m *KeyCredential) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the usage property value. A string that describes the purpose for which the key can be used; for example, Verify.
func (m *KeyCredential) GetUsage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usage
    }
}
// The deserialization information for the current model
func (m *KeyCredential) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["customKeyIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetCustomKeyIdentifier(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["keyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKeyId(val)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    res["usage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUsage(val)
        return nil
    }
    return res
}
func (m *KeyCredential) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *KeyCredential) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *KeyCredential) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the customKeyIdentifier property value. Custom key identifier
// Parameters:
//  - value : Value to set for the customKeyIdentifier property.
func (m *KeyCredential) SetCustomKeyIdentifier(value []byte)() {
    m.customKeyIdentifier = value
}
// Sets the displayName property value. Friendly name for the key. Optional.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *KeyCredential) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the endDateTime property value. The date and time at which the credential expires.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *KeyCredential) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the key property value. The certificate's raw data in byte array converted to Base64 string; for example, [System.Convert]::ToBase64String($Cert.GetRawCertData()).
// Parameters:
//  - value : Value to set for the key property.
func (m *KeyCredential) SetKey(value []byte)() {
    m.key = value
}
// Sets the keyId property value. The unique identifier (GUID) for the key.
// Parameters:
//  - value : Value to set for the keyId property.
func (m *KeyCredential) SetKeyId(value *string)() {
    m.keyId = value
}
// Sets the startDateTime property value. The date and time at which the credential becomes valid.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *KeyCredential) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the type_escaped property value. The type of key credential; for example, Symmetric.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *KeyCredential) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the usage property value. A string that describes the purpose for which the key can be used; for example, Verify.
// Parameters:
//  - value : Value to set for the usage property.
func (m *KeyCredential) SetUsage(value *string)() {
    m.usage = value
}
