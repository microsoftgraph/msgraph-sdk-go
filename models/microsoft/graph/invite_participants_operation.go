package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InviteParticipantsOperation 
type InviteParticipantsOperation struct {
    CommsOperation
    // The participants to invite.
    participants []InvitationParticipantInfoable;
}
// NewInviteParticipantsOperation instantiates a new inviteParticipantsOperation and sets the default values.
func NewInviteParticipantsOperation()(*InviteParticipantsOperation) {
    m := &InviteParticipantsOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateInviteParticipantsOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInviteParticipantsOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInviteParticipantsOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InviteParticipantsOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInvitationParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InvitationParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(InvitationParticipantInfoable)
            }
            m.SetParticipants(res)
        }
        return nil
    }
    return res
}
// GetParticipants gets the participants property value. The participants to invite.
func (m *InviteParticipantsOperation) GetParticipants()([]InvitationParticipantInfoable) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// Serialize serializes information the current object
func (m *InviteParticipantsOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetParticipants() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParticipants()))
        for i, v := range m.GetParticipants() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("participants", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetParticipants sets the participants property value. The participants to invite.
func (m *InviteParticipantsOperation) SetParticipants(value []InvitationParticipantInfoable)() {
    if m != nil {
        m.participants = value
    }
}
