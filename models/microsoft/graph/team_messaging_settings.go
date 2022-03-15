package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamMessagingSettings provides operations to manage the drive singleton.
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
// NewTeamMessagingSettings instantiates a new teamMessagingSettings and sets the default values.
func NewTeamMessagingSettings()(*TeamMessagingSettings) {
    m := &TeamMessagingSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamMessagingSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamMessagingSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamMessagingSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamMessagingSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowChannelMentions gets the allowChannelMentions property value. If set to true, @channel mentions are allowed.
func (m *TeamMessagingSettings) GetAllowChannelMentions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowChannelMentions
    }
}
// GetAllowOwnerDeleteMessages gets the allowOwnerDeleteMessages property value. If set to true, owners can delete any message.
func (m *TeamMessagingSettings) GetAllowOwnerDeleteMessages()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowOwnerDeleteMessages
    }
}
// GetAllowTeamMentions gets the allowTeamMentions property value. If set to true, @team mentions are allowed.
func (m *TeamMessagingSettings) GetAllowTeamMentions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowTeamMentions
    }
}
// GetAllowUserDeleteMessages gets the allowUserDeleteMessages property value. If set to true, users can delete their messages.
func (m *TeamMessagingSettings) GetAllowUserDeleteMessages()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUserDeleteMessages
    }
}
// GetAllowUserEditMessages gets the allowUserEditMessages property value. If set to true, users can edit their messages.
func (m *TeamMessagingSettings) GetAllowUserEditMessages()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUserEditMessages
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamMessagingSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowChannelMentions sets the allowChannelMentions property value. If set to true, @channel mentions are allowed.
func (m *TeamMessagingSettings) SetAllowChannelMentions(value *bool)() {
    if m != nil {
        m.allowChannelMentions = value
    }
}
// SetAllowOwnerDeleteMessages sets the allowOwnerDeleteMessages property value. If set to true, owners can delete any message.
func (m *TeamMessagingSettings) SetAllowOwnerDeleteMessages(value *bool)() {
    if m != nil {
        m.allowOwnerDeleteMessages = value
    }
}
// SetAllowTeamMentions sets the allowTeamMentions property value. If set to true, @team mentions are allowed.
func (m *TeamMessagingSettings) SetAllowTeamMentions(value *bool)() {
    if m != nil {
        m.allowTeamMentions = value
    }
}
// SetAllowUserDeleteMessages sets the allowUserDeleteMessages property value. If set to true, users can delete their messages.
func (m *TeamMessagingSettings) SetAllowUserDeleteMessages(value *bool)() {
    if m != nil {
        m.allowUserDeleteMessages = value
    }
}
// SetAllowUserEditMessages sets the allowUserEditMessages property value. If set to true, users can edit their messages.
func (m *TeamMessagingSettings) SetAllowUserEditMessages(value *bool)() {
    if m != nil {
        m.allowUserEditMessages = value
    }
}
