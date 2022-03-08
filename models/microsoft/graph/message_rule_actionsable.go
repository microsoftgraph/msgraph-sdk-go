package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageRuleActionsable 
type MessageRuleActionsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAssignCategories()([]string)
    GetCopyToFolder()(*string)
    GetDelete()(*bool)
    GetForwardAsAttachmentTo()([]Recipientable)
    GetForwardTo()([]Recipientable)
    GetMarkAsRead()(*bool)
    GetMarkImportance()(*Importance)
    GetMoveToFolder()(*string)
    GetPermanentDelete()(*bool)
    GetRedirectTo()([]Recipientable)
    GetStopProcessingRules()(*bool)
    SetAssignCategories(value []string)()
    SetCopyToFolder(value *string)()
    SetDelete(value *bool)()
    SetForwardAsAttachmentTo(value []Recipientable)()
    SetForwardTo(value []Recipientable)()
    SetMarkAsRead(value *bool)()
    SetMarkImportance(value *Importance)()
    SetMoveToFolder(value *string)()
    SetPermanentDelete(value *bool)()
    SetRedirectTo(value []Recipientable)()
    SetStopProcessingRules(value *bool)()
}
