package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingParticipantInfo provides operations to manage the cloudCommunications singleton.
type MeetingParticipantInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identity information of the participant.
    identity IdentitySetable;
    // Specifies the participant's role in the meeting.  Possible values are attendee, presenter, producer, and unknownFutureValue.
    role *OnlineMeetingRole;
    // User principal name of the participant.
    upn *string;
}
// NewMeetingParticipantInfo instantiates a new meetingParticipantInfo and sets the default values.
func NewMeetingParticipantInfo()(*MeetingParticipantInfo) {
    m := &MeetingParticipantInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeetingParticipantInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingParticipantInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMeetingParticipantInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingParticipantInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingParticipantInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["identity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(IdentitySetable))
        }
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*OnlineMeetingRole))
        }
        return nil
    }
    res["upn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpn(val)
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. Identity information of the participant.
func (m *MeetingParticipantInfo) GetIdentity()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
// GetRole gets the role property value. Specifies the participant's role in the meeting.  Possible values are attendee, presenter, producer, and unknownFutureValue.
func (m *MeetingParticipantInfo) GetRole()(*OnlineMeetingRole) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// GetUpn gets the upn property value. User principal name of the participant.
func (m *MeetingParticipantInfo) GetUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upn
    }
}
func (m *MeetingParticipantInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MeetingParticipantInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err := writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("upn", m.GetUpn())
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
func (m *MeetingParticipantInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIdentity sets the identity property value. Identity information of the participant.
func (m *MeetingParticipantInfo) SetIdentity(value IdentitySetable)() {
    if m != nil {
        m.identity = value
    }
}
// SetRole sets the role property value. Specifies the participant's role in the meeting.  Possible values are attendee, presenter, producer, and unknownFutureValue.
func (m *MeetingParticipantInfo) SetRole(value *OnlineMeetingRole)() {
    if m != nil {
        m.role = value
    }
}
// SetUpn sets the upn property value. User principal name of the participant.
func (m *MeetingParticipantInfo) SetUpn(value *string)() {
    if m != nil {
        m.upn = value
    }
}
