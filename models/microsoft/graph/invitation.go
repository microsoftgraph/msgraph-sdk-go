package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Invitation struct {
    Entity
    invitedUser *User;
    invitedUserDisplayName *string;
    invitedUserEmailAddress *string;
    invitedUserMessageInfo *InvitedUserMessageInfo;
    invitedUserType *string;
    inviteRedeemUrl *string;
    inviteRedirectUrl *string;
    sendInvitationMessage *bool;
    status *string;
}
func NewInvitation()(*Invitation) {
    m := &Invitation{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Invitation) GetInvitedUser()(*User) {
    if m == nil {
        return nil
    } else {
        return m.invitedUser
    }
}
func (m *Invitation) GetInvitedUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitedUserDisplayName
    }
}
func (m *Invitation) GetInvitedUserEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitedUserEmailAddress
    }
}
func (m *Invitation) GetInvitedUserMessageInfo()(*InvitedUserMessageInfo) {
    if m == nil {
        return nil
    } else {
        return m.invitedUserMessageInfo
    }
}
func (m *Invitation) GetInvitedUserType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitedUserType
    }
}
func (m *Invitation) GetInviteRedeemUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inviteRedeemUrl
    }
}
func (m *Invitation) GetInviteRedirectUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inviteRedirectUrl
    }
}
func (m *Invitation) GetSendInvitationMessage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendInvitationMessage
    }
}
func (m *Invitation) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *Invitation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["invitedUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUser() })
        if err != nil {
            return err
        }
        m.SetInvitedUser(val.(*User))
        return nil
    }
    res["invitedUserDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvitedUserDisplayName(val)
        return nil
    }
    res["invitedUserEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvitedUserEmailAddress(val)
        return nil
    }
    res["invitedUserMessageInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInvitedUserMessageInfo() })
        if err != nil {
            return err
        }
        m.SetInvitedUserMessageInfo(val.(*InvitedUserMessageInfo))
        return nil
    }
    res["invitedUserType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvitedUserType(val)
        return nil
    }
    res["inviteRedeemUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInviteRedeemUrl(val)
        return nil
    }
    res["inviteRedirectUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInviteRedirectUrl(val)
        return nil
    }
    res["sendInvitationMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSendInvitationMessage(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    return res
}
func (m *Invitation) IsNil()(bool) {
    return m == nil
}
func (m *Invitation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("invitedUser", m.GetInvitedUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invitedUserDisplayName", m.GetInvitedUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invitedUserEmailAddress", m.GetInvitedUserEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("invitedUserMessageInfo", m.GetInvitedUserMessageInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invitedUserType", m.GetInvitedUserType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("inviteRedeemUrl", m.GetInviteRedeemUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("inviteRedirectUrl", m.GetInviteRedirectUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("sendInvitationMessage", m.GetSendInvitationMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Invitation) SetInvitedUser(value *User)() {
    m.invitedUser = value
}
func (m *Invitation) SetInvitedUserDisplayName(value *string)() {
    m.invitedUserDisplayName = value
}
func (m *Invitation) SetInvitedUserEmailAddress(value *string)() {
    m.invitedUserEmailAddress = value
}
func (m *Invitation) SetInvitedUserMessageInfo(value *InvitedUserMessageInfo)() {
    m.invitedUserMessageInfo = value
}
func (m *Invitation) SetInvitedUserType(value *string)() {
    m.invitedUserType = value
}
func (m *Invitation) SetInviteRedeemUrl(value *string)() {
    m.inviteRedeemUrl = value
}
func (m *Invitation) SetInviteRedirectUrl(value *string)() {
    m.inviteRedirectUrl = value
}
func (m *Invitation) SetSendInvitationMessage(value *bool)() {
    m.sendInvitationMessage = value
}
func (m *Invitation) SetStatus(value *string)() {
    m.status = value
}
