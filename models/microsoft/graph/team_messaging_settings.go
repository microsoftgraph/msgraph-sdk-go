package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamMessagingSettings struct {
    additionalData map[string]interface{};
    allowChannelMentions *bool;
    allowOwnerDeleteMessages *bool;
    allowTeamMentions *bool;
    allowUserDeleteMessages *bool;
    allowUserEditMessages *bool;
}
func NewTeamMessagingSettings()(*TeamMessagingSettings) {
    m := &TeamMessagingSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TeamMessagingSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TeamMessagingSettings) GetAllowChannelMentions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowChannelMentions
    }
}
func (m *TeamMessagingSettings) GetAllowOwnerDeleteMessages()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowOwnerDeleteMessages
    }
}
func (m *TeamMessagingSettings) GetAllowTeamMentions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowTeamMentions
    }
}
func (m *TeamMessagingSettings) GetAllowUserDeleteMessages()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUserDeleteMessages
    }
}
func (m *TeamMessagingSettings) GetAllowUserEditMessages()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUserEditMessages
    }
}
func (m *TeamMessagingSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowChannelMentions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowChannelMentions(val)
        return nil
    }
    res["allowOwnerDeleteMessages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowOwnerDeleteMessages(val)
        return nil
    }
    res["allowTeamMentions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowTeamMentions(val)
        return nil
    }
    res["allowUserDeleteMessages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowUserDeleteMessages(val)
        return nil
    }
    res["allowUserEditMessages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowUserEditMessages(val)
        return nil
    }
    return res
}
func (m *TeamMessagingSettings) IsNil()(bool) {
    return m == nil
}
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
func (m *TeamMessagingSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TeamMessagingSettings) SetAllowChannelMentions(value *bool)() {
    m.allowChannelMentions = value
}
func (m *TeamMessagingSettings) SetAllowOwnerDeleteMessages(value *bool)() {
    m.allowOwnerDeleteMessages = value
}
func (m *TeamMessagingSettings) SetAllowTeamMentions(value *bool)() {
    m.allowTeamMentions = value
}
func (m *TeamMessagingSettings) SetAllowUserDeleteMessages(value *bool)() {
    m.allowUserDeleteMessages = value
}
func (m *TeamMessagingSettings) SetAllowUserEditMessages(value *bool)() {
    m.allowUserEditMessages = value
}
