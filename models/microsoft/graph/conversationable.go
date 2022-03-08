package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Conversationable 
type Conversationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetHasAttachments()(*bool)
    GetLastDeliveredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPreview()(*string)
    GetThreads()([]ConversationThreadable)
    GetTopic()(*string)
    GetUniqueSenders()([]string)
    SetHasAttachments(value *bool)()
    SetLastDeliveredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPreview(value *string)()
    SetThreads(value []ConversationThreadable)()
    SetTopic(value *string)()
    SetUniqueSenders(value []string)()
}
