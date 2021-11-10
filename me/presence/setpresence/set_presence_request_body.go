package setpresence

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new setPresenceRequestBody and sets the default values.
func NewSetPresenceRequestBody()(*SetPresenceRequestBody) {
    m := &SetPresenceRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the activity property value. 
func (m *SetPresenceRequestBody) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetPresenceRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the availability property value. 
func (m *SetPresenceRequestBody) GetAvailability()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availability
    }
}
// Gets the expirationDuration property value. 
func (m *SetPresenceRequestBody) GetExpirationDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expirationDuration
    }
}
// Gets the sessionId property value. 
func (m *SetPresenceRequestBody) GetSessionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sessionId
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activity property value. 
// Parameters:
//  - value : Value to set for the activity property.
func (m *SetPresenceRequestBody) SetActivity(value *string)() {
    m.activity = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SetPresenceRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the availability property value. 
// Parameters:
//  - value : Value to set for the availability property.
func (m *SetPresenceRequestBody) SetAvailability(value *string)() {
    m.availability = value
}
// Sets the expirationDuration property value. 
// Parameters:
//  - value : Value to set for the expirationDuration property.
func (m *SetPresenceRequestBody) SetExpirationDuration(value *string)() {
    m.expirationDuration = value
}
// Sets the sessionId property value. 
// Parameters:
//  - value : Value to set for the sessionId property.
func (m *SetPresenceRequestBody) SetSessionId(value *string)() {
    m.sessionId = value
}
