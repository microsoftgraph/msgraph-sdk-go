package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConversationThreadable 
type ConversationThreadable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCcRecipients()([]Recipientable)
    GetHasAttachments()(*bool)
    GetIsLocked()(*bool)
    GetLastDeliveredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPosts()([]Postable)
    GetPreview()(*string)
    GetTopic()(*string)
    GetToRecipients()([]Recipientable)
    GetUniqueSenders()([]string)
    SetCcRecipients(value []Recipientable)()
    SetHasAttachments(value *bool)()
    SetIsLocked(value *bool)()
    SetLastDeliveredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPosts(value []Postable)()
    SetPreview(value *string)()
    SetTopic(value *string)()
    SetToRecipients(value []Recipientable)()
    SetUniqueSenders(value []string)()
}
