package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MeetingParticipantInfo struct {
    additionalData map[string]interface{};
    identity *IdentitySet;
    role *OnlineMeetingRole;
    upn *string;
}
func NewMeetingParticipantInfo()(*MeetingParticipantInfo) {
    m := &MeetingParticipantInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MeetingParticipantInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MeetingParticipantInfo) GetIdentity()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
func (m *MeetingParticipantInfo) GetRole()(*OnlineMeetingRole) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
func (m *MeetingParticipantInfo) GetUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upn
    }
}
func (m *MeetingParticipantInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["identity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetIdentity(val.(*IdentitySet))
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingRole)
        if err != nil {
            return err
        }
        cast := val.(OnlineMeetingRole)
        m.SetRole(&cast)
        return nil
    }
    res["upn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUpn(val)
        return nil
    }
    return res
}
func (m *MeetingParticipantInfo) IsNil()(bool) {
    return m == nil
}
func (m *MeetingParticipantInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := m.GetRole().String()
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
func (m *MeetingParticipantInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MeetingParticipantInfo) SetIdentity(value *IdentitySet)() {
    m.identity = value
}
func (m *MeetingParticipantInfo) SetRole(value *OnlineMeetingRole)() {
    m.role = value
}
func (m *MeetingParticipantInfo) SetUpn(value *string)() {
    m.upn = value
}
