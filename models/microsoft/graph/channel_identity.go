package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChannelIdentity 
type ChannelIdentity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identity of the channel in which the message was posted.
    channelId *string;
    // The identity of the team in which the message was posted.
    teamId *string;
}
// NewChannelIdentity instantiates a new channelIdentity and sets the default values.
func NewChannelIdentity()(*ChannelIdentity) {
    m := &ChannelIdentity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChannelIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChannelId gets the channelId property value. The identity of the channel in which the message was posted.
func (m *ChannelIdentity) GetChannelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.channelId
    }
}
// GetTeamId gets the teamId property value. The identity of the team in which the message was posted.
func (m *ChannelIdentity) GetTeamId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["channelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelId(val)
        }
        return nil
    }
    res["teamId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    return res
}
func (m *ChannelIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChannelIdentity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChannelId sets the channelId property value. The identity of the channel in which the message was posted.
func (m *ChannelIdentity) SetChannelId(value *string)() {
    if m != nil {
        m.channelId = value
    }
}
// SetTeamId sets the teamId property value. The identity of the team in which the message was posted.
func (m *ChannelIdentity) SetTeamId(value *string)() {
    if m != nil {
        m.teamId = value
    }
}
