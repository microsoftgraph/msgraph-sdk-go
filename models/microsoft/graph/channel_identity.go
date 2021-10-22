package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChannelIdentity struct {
    additionalData map[string]interface{};
    channelId *string;
    teamId *string;
}
func NewChannelIdentity()(*ChannelIdentity) {
    m := &ChannelIdentity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChannelIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChannelIdentity) GetChannelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.channelId
    }
}
func (m *ChannelIdentity) GetTeamId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamId
    }
}
func (m *ChannelIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["channelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChannelId(val)
        return nil
    }
    res["teamId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTeamId(val)
        return nil
    }
    return res
}
func (m *ChannelIdentity) IsNil()(bool) {
    return m == nil
}
func (m *ChannelIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("channelId", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teamId", m.GetTeamId())
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
func (m *ChannelIdentity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChannelIdentity) SetChannelId(value *string)() {
    m.channelId = value
}
func (m *ChannelIdentity) SetTeamId(value *string)() {
    m.teamId = value
}
