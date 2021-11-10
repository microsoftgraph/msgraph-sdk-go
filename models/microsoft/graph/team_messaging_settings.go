package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamMessagingSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If set to true, @channel mentions are allowed.
    allowChannelMentions *bool;
    // If set to true, owners can delete any message.
    allowOwnerDeleteMessages *bool;
    // If set to true, @team mentions are allowed.
    allowTeamMentions *bool;
    // If set to true, users can delete their messages.
    allowUserDeleteMessages *bool;
    // If set to true, users can edit their messages.
    allowUserEditMessages *bool;
}
// Instantiates a new teamMessagingSettings and sets the default values.
func NewTeamMessagingSettings()(*TeamMessagingSettings) {
    m := &TeamMessagingSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamMessagingSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowChannelMentions property value. If set to true, @channel mentions are allowed.
func (m *TeamMessagingSettings) GetAllowChannelMentions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowChannelMentions
    }
}
// Gets the allowOwnerDeleteMessages property value. If set to true, owners can delete any message.
func (m *TeamMessagingSettings) GetAllowOwnerDeleteMessages()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowOwnerDeleteMessages
    }
}
// Gets the allowTeamMentions property value. If set to true, @team mentions are allowed.
func (m *TeamMessagingSettings) GetAllowTeamMentions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowTeamMentions
    }
}
// Gets the allowUserDeleteMessages property value. If set to true, users can delete their messages.
func (m *TeamMessagingSettings) GetAllowUserDeleteMessages()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUserDeleteMessages
    }
}
// Gets the allowUserEditMessages property value. If set to true, users can edit their messages.
func (m *TeamMessagingSettings) GetAllowUserEditMessages()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUserEditMessages
    }
}
// The deserialization information for the current model
func (m *TeamMessagingSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowChannelMentions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowChannelMentions(val)
        }
        return nil
    }
    res["allowOwnerDeleteMessages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowOwnerDeleteMessages(val)
        }
        return nil
    }
    res["allowTeamMentions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTeamMentions(val)
        }
        return nil
    }
    res["allowUserDeleteMessages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUserDeleteMessages(val)
        }
        return nil
    }
    res["allowUserEditMessages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUserEditMessages(val)
        }
        return nil
    }
    return res
}
func (m *TeamMessagingSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeamMessagingSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowChannelMentions", m.GetAllowChannelMentions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowOwnerDeleteMessages", m.GetAllowOwnerDeleteMessages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowTeamMentions", m.GetAllowTeamMentions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowUserDeleteMessages", m.GetAllowUserDeleteMessages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowUserEditMessages", m.GetAllowUserEditMessages())
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
func (m *TeamMessagingSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowChannelMentions property value. If set to true, @channel mentions are allowed.
// Parameters:
//  - value : Value to set for the allowChannelMentions property.
func (m *TeamMessagingSettings) SetAllowChannelMentions(value *bool)() {
    m.allowChannelMentions = value
}
// Sets the allowOwnerDeleteMessages property value. If set to true, owners can delete any message.
// Parameters:
//  - value : Value to set for the allowOwnerDeleteMessages property.
func (m *TeamMessagingSettings) SetAllowOwnerDeleteMessages(value *bool)() {
    m.allowOwnerDeleteMessages = value
}
// Sets the allowTeamMentions property value. If set to true, @team mentions are allowed.
// Parameters:
//  - value : Value to set for the allowTeamMentions property.
func (m *TeamMessagingSettings) SetAllowTeamMentions(value *bool)() {
    m.allowTeamMentions = value
}
// Sets the allowUserDeleteMessages property value. If set to true, users can delete their messages.
// Parameters:
//  - value : Value to set for the allowUserDeleteMessages property.
func (m *TeamMessagingSettings) SetAllowUserDeleteMessages(value *bool)() {
    m.allowUserDeleteMessages = value
}
// Sets the allowUserEditMessages property value. If set to true, users can edit their messages.
// Parameters:
//  - value : Value to set for the allowUserEditMessages property.
func (m *TeamMessagingSettings) SetAllowUserEditMessages(value *bool)() {
    m.allowUserEditMessages = value
}
