package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Invitation provides operations to manage the collection of invitation entities.
type Invitation struct {
    Entity
    // The user created as part of the invitation creation. Read-Only
    invitedUser Userable;
    // The display name of the user being invited.
    invitedUserDisplayName *string;
    // The email address of the user being invited. Required. The following special characters are not permitted in the email address:Tilde (~)Exclamation point (!)Number sign (#)Dollar sign ($)Percent (%)Circumflex (^)Ampersand (&)Asterisk (*)Parentheses (( ))Plus sign (+)Equal sign (=)Brackets ([ ])Braces ({ })Backslash (/)Slash mark (/)Pipe (/|)Semicolon (;)Colon (:)Quotation marks (')Angle brackets (< >)Question mark (?)Comma (,)However, the following exceptions apply:A period (.) or a hyphen (-) is permitted anywhere in the user name, except at the beginning or end of the name.An underscore (_) is permitted anywhere in the user name. This includes at the beginning or end of the name.
    invitedUserEmailAddress *string;
    // Additional configuration for the message being sent to the invited user, including customizing message text, language and cc recipient list.
    invitedUserMessageInfo InvitedUserMessageInfoable;
    // The userType of the user being invited. By default, this is Guest. You can invite as Member if you are a company administrator.
    invitedUserType *string;
    // The URL the user can use to redeem their invitation. Read-only.
    inviteRedeemUrl *string;
    // The URL the user should be redirected to once the invitation is redeemed. Required.
    inviteRedirectUrl *string;
    // Indicates whether an email should be sent to the user being invited. The default is false.
    sendInvitationMessage *bool;
    // The status of the invitation. Possible values are: PendingAcceptance, Completed, InProgress, and Error.
    status *string;
}
// NewInvitation instantiates a new invitation and sets the default values.
func NewInvitation()(*Invitation) {
    m := &Invitation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInvitationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInvitationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInvitation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Invitation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["invitedUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitedUser(val.(Userable))
        }
        return nil
    }
    res["invitedUserDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitedUserDisplayName(val)
        }
        return nil
    }
    res["invitedUserEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitedUserEmailAddress(val)
        }
        return nil
    }
    res["invitedUserMessageInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateInvitedUserMessageInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitedUserMessageInfo(val.(InvitedUserMessageInfoable))
        }
        return nil
    }
    res["invitedUserType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitedUserType(val)
        }
        return nil
    }
    res["inviteRedeemUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInviteRedeemUrl(val)
        }
        return nil
    }
    res["inviteRedirectUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInviteRedirectUrl(val)
        }
        return nil
    }
    res["sendInvitationMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendInvitationMessage(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetInvitedUser gets the invitedUser property value. The user created as part of the invitation creation. Read-Only
func (m *Invitation) GetInvitedUser()(Userable) {
    if m == nil {
        return nil
    } else {
        return m.invitedUser
    }
}
// GetInvitedUserDisplayName gets the invitedUserDisplayName property value. The display name of the user being invited.
func (m *Invitation) GetInvitedUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitedUserDisplayName
    }
}
// GetInvitedUserEmailAddress gets the invitedUserEmailAddress property value. The email address of the user being invited. Required. The following special characters are not permitted in the email address:Tilde (~)Exclamation point (!)Number sign (#)Dollar sign ($)Percent (%)Circumflex (^)Ampersand (&)Asterisk (*)Parentheses (( ))Plus sign (+)Equal sign (=)Brackets ([ ])Braces ({ })Backslash (/)Slash mark (/)Pipe (/|)Semicolon (;)Colon (:)Quotation marks (')Angle brackets (< >)Question mark (?)Comma (,)However, the following exceptions apply:A period (.) or a hyphen (-) is permitted anywhere in the user name, except at the beginning or end of the name.An underscore (_) is permitted anywhere in the user name. This includes at the beginning or end of the name.
func (m *Invitation) GetInvitedUserEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitedUserEmailAddress
    }
}
// GetInvitedUserMessageInfo gets the invitedUserMessageInfo property value. Additional configuration for the message being sent to the invited user, including customizing message text, language and cc recipient list.
func (m *Invitation) GetInvitedUserMessageInfo()(InvitedUserMessageInfoable) {
    if m == nil {
        return nil
    } else {
        return m.invitedUserMessageInfo
    }
}
// GetInvitedUserType gets the invitedUserType property value. The userType of the user being invited. By default, this is Guest. You can invite as Member if you are a company administrator.
func (m *Invitation) GetInvitedUserType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitedUserType
    }
}
// GetInviteRedeemUrl gets the inviteRedeemUrl property value. The URL the user can use to redeem their invitation. Read-only.
func (m *Invitation) GetInviteRedeemUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inviteRedeemUrl
    }
}
// GetInviteRedirectUrl gets the inviteRedirectUrl property value. The URL the user should be redirected to once the invitation is redeemed. Required.
func (m *Invitation) GetInviteRedirectUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inviteRedirectUrl
    }
}
// GetSendInvitationMessage gets the sendInvitationMessage property value. Indicates whether an email should be sent to the user being invited. The default is false.
func (m *Invitation) GetSendInvitationMessage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendInvitationMessage
    }
}
// GetStatus gets the status property value. The status of the invitation. Possible values are: PendingAcceptance, Completed, InProgress, and Error.
func (m *Invitation) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *Invitation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetInvitedUser sets the invitedUser property value. The user created as part of the invitation creation. Read-Only
func (m *Invitation) SetInvitedUser(value Userable)() {
    if m != nil {
        m.invitedUser = value
    }
}
// SetInvitedUserDisplayName sets the invitedUserDisplayName property value. The display name of the user being invited.
func (m *Invitation) SetInvitedUserDisplayName(value *string)() {
    if m != nil {
        m.invitedUserDisplayName = value
    }
}
// SetInvitedUserEmailAddress sets the invitedUserEmailAddress property value. The email address of the user being invited. Required. The following special characters are not permitted in the email address:Tilde (~)Exclamation point (!)Number sign (#)Dollar sign ($)Percent (%)Circumflex (^)Ampersand (&)Asterisk (*)Parentheses (( ))Plus sign (+)Equal sign (=)Brackets ([ ])Braces ({ })Backslash (/)Slash mark (/)Pipe (/|)Semicolon (;)Colon (:)Quotation marks (')Angle brackets (< >)Question mark (?)Comma (,)However, the following exceptions apply:A period (.) or a hyphen (-) is permitted anywhere in the user name, except at the beginning or end of the name.An underscore (_) is permitted anywhere in the user name. This includes at the beginning or end of the name.
func (m *Invitation) SetInvitedUserEmailAddress(value *string)() {
    if m != nil {
        m.invitedUserEmailAddress = value
    }
}
// SetInvitedUserMessageInfo sets the invitedUserMessageInfo property value. Additional configuration for the message being sent to the invited user, including customizing message text, language and cc recipient list.
func (m *Invitation) SetInvitedUserMessageInfo(value InvitedUserMessageInfoable)() {
    if m != nil {
        m.invitedUserMessageInfo = value
    }
}
// SetInvitedUserType sets the invitedUserType property value. The userType of the user being invited. By default, this is Guest. You can invite as Member if you are a company administrator.
func (m *Invitation) SetInvitedUserType(value *string)() {
    if m != nil {
        m.invitedUserType = value
    }
}
// SetInviteRedeemUrl sets the inviteRedeemUrl property value. The URL the user can use to redeem their invitation. Read-only.
func (m *Invitation) SetInviteRedeemUrl(value *string)() {
    if m != nil {
        m.inviteRedeemUrl = value
    }
}
// SetInviteRedirectUrl sets the inviteRedirectUrl property value. The URL the user should be redirected to once the invitation is redeemed. Required.
func (m *Invitation) SetInviteRedirectUrl(value *string)() {
    if m != nil {
        m.inviteRedirectUrl = value
    }
}
// SetSendInvitationMessage sets the sendInvitationMessage property value. Indicates whether an email should be sent to the user being invited. The default is false.
func (m *Invitation) SetSendInvitationMessage(value *bool)() {
    if m != nil {
        m.sendInvitationMessage = value
    }
}
// SetStatus sets the status property value. The status of the invitation. Possible values are: PendingAcceptance, Completed, InProgress, and Error.
func (m *Invitation) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
