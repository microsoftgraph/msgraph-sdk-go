package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MailTipsable 
type MailTipsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAutomaticReplies()(AutomaticRepliesMailTipsable)
    GetCustomMailTip()(*string)
    GetDeliveryRestricted()(*bool)
    GetEmailAddress()(EmailAddressable)
    GetError()(MailTipsErrorable)
    GetExternalMemberCount()(*int32)
    GetIsModerated()(*bool)
    GetMailboxFull()(*bool)
    GetMaxMessageSize()(*int32)
    GetRecipientScope()(*RecipientScopeType)
    GetRecipientSuggestions()([]Recipientable)
    GetTotalMemberCount()(*int32)
    SetAutomaticReplies(value AutomaticRepliesMailTipsable)()
    SetCustomMailTip(value *string)()
    SetDeliveryRestricted(value *bool)()
    SetEmailAddress(value EmailAddressable)()
    SetError(value MailTipsErrorable)()
    SetExternalMemberCount(value *int32)()
    SetIsModerated(value *bool)()
    SetMailboxFull(value *bool)()
    SetMaxMessageSize(value *int32)()
    SetRecipientScope(value *RecipientScopeType)()
    SetRecipientSuggestions(value []Recipientable)()
    SetTotalMemberCount(value *int32)()
}
