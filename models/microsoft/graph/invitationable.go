package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Invitationable 
type Invitationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetInvitedUser()(Userable)
    GetInvitedUserDisplayName()(*string)
    GetInvitedUserEmailAddress()(*string)
    GetInvitedUserMessageInfo()(InvitedUserMessageInfoable)
    GetInvitedUserType()(*string)
    GetInviteRedeemUrl()(*string)
    GetInviteRedirectUrl()(*string)
    GetSendInvitationMessage()(*bool)
    GetStatus()(*string)
    SetInvitedUser(value Userable)()
    SetInvitedUserDisplayName(value *string)()
    SetInvitedUserEmailAddress(value *string)()
    SetInvitedUserMessageInfo(value InvitedUserMessageInfoable)()
    SetInvitedUserType(value *string)()
    SetInviteRedeemUrl(value *string)()
    SetInviteRedirectUrl(value *string)()
    SetSendInvitationMessage(value *bool)()
    SetStatus(value *string)()
}
