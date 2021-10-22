package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InviteParticipantsOperation struct {
    CommsOperation
    participants []InvitationParticipantInfo;
}
func NewInviteParticipantsOperation()(*InviteParticipantsOperation) {
    m := &InviteParticipantsOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
func (m *InviteParticipantsOperation) GetParticipants()([]InvitationParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
func (m *InviteParticipantsOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInvitationParticipantInfo() })
        if err != nil {
            return err
        }
        res := make([]InvitationParticipantInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*InvitationParticipantInfo))
        }
        m.SetParticipants(res)
        return nil
    }
    return res
}
func (m *InviteParticipantsOperation) IsNil()(bool) {
    return m == nil
}
func (m *InviteParticipantsOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParticipants()))
        for i, v := range m.GetParticipants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("participants", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *InviteParticipantsOperation) SetParticipants(value []InvitationParticipantInfo)() {
    m.participants = value
}
