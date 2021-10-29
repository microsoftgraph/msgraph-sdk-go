package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamGuestSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If set to true, guests can add and update channels.
    allowCreateUpdateChannels *bool;
    // If set to true, guests can delete channels.
    allowDeleteChannels *bool;
}
// Instantiates a new teamGuestSettings and sets the default values.
func NewTeamGuestSettings()(*TeamGuestSettings) {
    m := &TeamGuestSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamGuestSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowCreateUpdateChannels property value. If set to true, guests can add and update channels.
func (m *TeamGuestSettings) GetAllowCreateUpdateChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreateUpdateChannels
    }
}
// Gets the allowDeleteChannels property value. If set to true, guests can delete channels.
func (m *TeamGuestSettings) GetAllowDeleteChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteChannels
    }
}
// The deserialization information for the current model
func (m *TeamGuestSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowCreateUpdateChannels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowCreateUpdateChannels(val)
        return nil
    }
    res["allowDeleteChannels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowDeleteChannels(val)
        return nil
    }
    return res
}
func (m *TeamGuestSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeamGuestSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowCreateUpdateChannels", m.GetAllowCreateUpdateChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeleteChannels", m.GetAllowDeleteChannels())
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
func (m *TeamGuestSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowCreateUpdateChannels property value. If set to true, guests can add and update channels.
// Parameters:
//  - value : Value to set for the allowCreateUpdateChannels property.
func (m *TeamGuestSettings) SetAllowCreateUpdateChannels(value *bool)() {
    m.allowCreateUpdateChannels = value
}
// Sets the allowDeleteChannels property value. If set to true, guests can delete channels.
// Parameters:
//  - value : Value to set for the allowDeleteChannels property.
func (m *TeamGuestSettings) SetAllowDeleteChannels(value *bool)() {
    m.allowDeleteChannels = value
}
