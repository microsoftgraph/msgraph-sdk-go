package setpresence

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SetPresenceRequestBody 
type SetPresenceRequestBody struct {
    // 
    activity *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    availability *string;
    // 
    expirationDuration *string;
    // 
    sessionId *string;
}
// NewSetPresenceRequestBody instantiates a new setPresenceRequestBody and sets the default values.
func NewSetPresenceRequestBody()(*SetPresenceRequestBody) {
    m := &SetPresenceRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetActivity gets the activity property value. 
func (m *SetPresenceRequestBody) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetPresenceRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAvailability gets the availability property value. 
func (m *SetPresenceRequestBody) GetAvailability()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availability
    }
}
// GetExpirationDuration gets the expirationDuration property value. 
func (m *SetPresenceRequestBody) GetExpirationDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expirationDuration
    }
}
// GetSessionId gets the sessionId property value. 
func (m *SetPresenceRequestBody) GetSessionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sessionId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetPresenceRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["availability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailability(val)
        }
        return nil
    }
    res["expirationDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDuration(val)
        }
        return nil
    }
    res["sessionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionId(val)
        }
        return nil
    }
    return res
}
func (m *SetPresenceRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SetPresenceRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("availability", m.GetAvailability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("expirationDuration", m.GetExpirationDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionId", m.GetSessionId())
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
// SetActivity sets the activity property value. 
func (m *SetPresenceRequestBody) SetActivity(value *string)() {
    if m != nil {
        m.activity = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetPresenceRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAvailability sets the availability property value. 
func (m *SetPresenceRequestBody) SetAvailability(value *string)() {
    if m != nil {
        m.availability = value
    }
}
// SetExpirationDuration sets the expirationDuration property value. 
func (m *SetPresenceRequestBody) SetExpirationDuration(value *string)() {
    if m != nil {
        m.expirationDuration = value
    }
}
// SetSessionId sets the sessionId property value. 
func (m *SetPresenceRequestBody) SetSessionId(value *string)() {
    if m != nil {
        m.sessionId = value
    }
}
