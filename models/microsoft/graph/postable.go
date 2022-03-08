package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Postable 
type Postable interface {
    OutlookItemable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAttachments()([]Attachmentable)
    GetBody()(ItemBodyable)
    GetConversationId()(*string)
    GetConversationThreadId()(*string)
    GetExtensions()([]Extensionable)
    GetFrom()(Recipientable)
    GetHasAttachments()(*bool)
    GetInReplyTo()(Postable)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetNewParticipants()([]Recipientable)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSender()(Recipientable)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    SetAttachments(value []Attachmentable)()
    SetBody(value ItemBodyable)()
    SetConversationId(value *string)()
    SetConversationThreadId(value *string)()
    SetExtensions(value []Extensionable)()
    SetFrom(value Recipientable)()
    SetHasAttachments(value *bool)()
    SetInReplyTo(value Postable)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetNewParticipants(value []Recipientable)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSender(value Recipientable)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
}
