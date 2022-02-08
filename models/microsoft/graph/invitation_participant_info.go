package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InvitationParticipantInfo 
type InvitationParticipantInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    hidden *bool;
    // 
    identity *IdentitySet;
    // 
    participantId *string;
    // 
    removeFromDefaultAudioRoutingGroup *bool;
    // Optional. The call which the target identity is currently a part of. This call will be dropped once the participant is added.
    replacesCallId *string;
}
// NewInvitationParticipantInfo instantiates a new invitationParticipantInfo and sets the default values.
func NewInvitationParticipantInfo()(*InvitationParticipantInfo) {
    m := &InvitationParticipantInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InvitationParticipantInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetHidden gets the hidden property value. 
func (m *InvitationParticipantInfo) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
// GetIdentity gets the identity property value. 
func (m *InvitationParticipantInfo) GetIdentity()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
// GetParticipantId gets the participantId property value. 
func (m *InvitationParticipantInfo) GetParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participantId
    }
}
// GetRemoveFromDefaultAudioRoutingGroup gets the removeFromDefaultAudioRoutingGroup property value. 
func (m *InvitationParticipantInfo) GetRemoveFromDefaultAudioRoutingGroup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.removeFromDefaultAudioRoutingGroup
    }
}
// GetReplacesCallId gets the replacesCallId property value. Optional. The call which the target identity is currently a part of. This call will be dropped once the participant is added.
func (m *InvitationParticipantInfo) GetReplacesCallId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.replacesCallId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InvitationParticipantInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidden(val)
        }
        return nil
    }
    res["identity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(*IdentitySet))
        }
        return nil
    }
    res["participantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantId(val)
        }
        return nil
    }
    res["removeFromDefaultAudioRoutingGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveFromDefaultAudioRoutingGroup(val)
        }
        return nil
    }
    res["replacesCallId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplacesCallId(val)
        }
        return nil
    }
    return res
}
func (m *InvitationParticipantInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InvitationParticipantInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hidden", m.GetHidden())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("participantId", m.GetParticipantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("removeFromDefaultAudioRoutingGroup", m.GetRemoveFromDefaultAudioRoutingGroup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("replacesCallId", m.GetReplacesCallId())
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
func (m *InvitationParticipantInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHidden sets the hidden property value. 
func (m *InvitationParticipantInfo) SetHidden(value *bool)() {
    if m != nil {
        m.hidden = value
    }
}
// SetIdentity sets the identity property value. 
func (m *InvitationParticipantInfo) SetIdentity(value *IdentitySet)() {
    if m != nil {
        m.identity = value
    }
}
// SetParticipantId sets the participantId property value. 
func (m *InvitationParticipantInfo) SetParticipantId(value *string)() {
    if m != nil {
        m.participantId = value
    }
}
// SetRemoveFromDefaultAudioRoutingGroup sets the removeFromDefaultAudioRoutingGroup property value. 
func (m *InvitationParticipantInfo) SetRemoveFromDefaultAudioRoutingGroup(value *bool)() {
    if m != nil {
        m.removeFromDefaultAudioRoutingGroup = value
    }
}
// SetReplacesCallId sets the replacesCallId property value. Optional. The call which the target identity is currently a part of. This call will be dropped once the participant is added.
func (m *InvitationParticipantInfo) SetReplacesCallId(value *string)() {
    if m != nil {
        m.replacesCallId = value
    }
}
