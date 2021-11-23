package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// teamGuestSettings 
type TeamGuestSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If set to true, guests can add and update channels.
    allowCreateUpdateChannels *bool;
    // If set to true, guests can delete channels.
    allowDeleteChannels *bool;
}
// NewTeamGuestSettings instantiates a new teamGuestSettings and sets the default values.
func NewTeamGuestSettings()(*TeamGuestSettings) {
    m := &TeamGuestSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamGuestSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowCreateUpdateChannels gets the allowCreateUpdateChannels property value. If set to true, guests can add and update channels.
func (m *TeamGuestSettings) GetAllowCreateUpdateChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreateUpdateChannels
    }
}
// GetAllowDeleteChannels gets the allowDeleteChannels property value. If set to true, guests can delete channels.
func (m *TeamGuestSettings) GetAllowDeleteChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteChannels
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamGuestSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowCreateUpdateChannels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCreateUpdateChannels(val)
        }
        return nil
    }
    res["allowDeleteChannels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeleteChannels(val)
        }
        return nil
    }
    return res
}
func (m *TeamGuestSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamGuestSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowCreateUpdateChannels sets the allowCreateUpdateChannels property value. If set to true, guests can add and update channels.
func (m *TeamGuestSettings) SetAllowCreateUpdateChannels(value *bool)() {
    m.allowCreateUpdateChannels = value
}
// SetAllowDeleteChannels sets the allowDeleteChannels property value. If set to true, guests can delete channels.
func (m *TeamGuestSettings) SetAllowDeleteChannels(value *bool)() {
    m.allowDeleteChannels = value
}
